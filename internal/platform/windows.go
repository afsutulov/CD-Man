//go:build windows

package platform

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"cdman2/internal/game"
)

var user = syscall.NewLazyDLL("user32.dll")
var gdi = syscall.NewLazyDLL("gdi32.dll")
var winmm = syscall.NewLazyDLL("winmm.dll")
var defProc = user.NewProc("DefWindowProcW")
var destroy = user.NewProc("DestroyWindow")
var postQuit = user.NewProc("PostQuitMessage")
var setCursor = user.NewProc("SetCursor")
var keyboard = user.NewProc("GetKeyboardState")
var toAscii = user.NewProc("ToAscii")
var setStyle = user.NewProc("SetWindowLongPtrW")
var setPos = user.NewProc("SetWindowPos")
var getMetric = user.NewProc("GetSystemMetrics")
var current *game.Machine
var full = true

type point struct{ X, Y int32 }
type rect struct{ Left, Top, Right, Bottom int32 }
type message struct {
	Hwnd           uintptr
	Message        uint32
	WParam, LParam uintptr
	Time           uint32
	Pt             point
	Private        uint32
}
type windowClass struct {
	Size, Style                        uint32
	Proc                               uintptr
	ClassExtra, WindowExtra            int32
	Instance, Icon, Cursor, Background uintptr
	MenuName, ClassName                *uint16
	IconSmall                          uintptr
}
type bitmapHeader struct {
	Size                   uint32
	Width, Height          int32
	Planes, BitCount       uint16
	Compression, ImageSize uint32
	XPels, YPels           int32
	Used, Important        uint32
}
type waveFormat struct {
	Format, Channels        uint16
	Rate, BytesPerSecond    uint32
	BlockAlign, Bits, Extra uint16
}
type waveHeader struct {
	Data             uintptr
	Length, Recorded uint32
	User             uintptr
	Flags, Loops     uint32
	Next, Reserved   uintptr
}
type waveSlot struct {
	header   waveHeader
	data     [4096]int16
	pinner   runtime.Pinner
	prepared bool
}
type waveOutput struct {
	handle uintptr
	slots  [8]waveSlot
	index  int
}

func utf(s string) *uint16 { p, _ := syscall.UTF16PtrFromString(s); return p }
func ShowError(s string) {
	user.NewProc("MessageBoxW").Call(0, uintptr(unsafe.Pointer(utf(s))), uintptr(unsafe.Pointer(utf("CD-Man"))), 0x10)
}
func windowProc(hwnd uintptr, msg uint32, w, l uintptr) uintptr {
	switch msg {
	case 0x0f:
		paintWindow(hwnd)
		return 0 // WM_PAINT: restore the completed frame.
	case 0x14:
		return 1 // WM_ERASEBKGND: background is part of the back buffer.
	case 0x10:
		destroy.Call(hwnd)
		return 0
	case 2:
		postQuit.Call(0)
		return 0
	case 0x20:
		setCursor.Call(0)
		return 1
	case 0x104:
		if w == 13 {
			full = !full
			setWindowMode(hwnd)
			return 0
		}
	case 0x100:
		if current != nil {
			var state [256]byte
			keyboard.Call(uintptr(unsafe.Pointer(&state[0])))
			var translated uint16
			scan := byte((l >> 16) & 255)
			n, _, _ := toAscii.Call(w, uintptr(scan), uintptr(unsafe.Pointer(&state[0])), uintptr(unsafe.Pointer(&translated)), 0)
			ascii := byte(0)
			if int32(n) > 0 {
				ascii = byte(translated)
			}
			current.Key(scan, ascii)
		}
		return 0
	}
	v, _, _ := defProc.Call(hwnd, uintptr(msg), w, l)
	return v
}
func setWindowMode(hwnd uintptr) {
	style := uintptr(0x00cf0000)
	w, h := uintptr(976), uintptr(759)
	x, y := uintptr(50), uintptr(50)
	if full {
		style = 0x80000000
		w, _, _ = getMetric.Call(0)
		h, _, _ = getMetric.Call(1)
		x = 0
		y = 0
	}
	setStyle.Call(hwnd, ^uintptr(15), style) // GWL_STYLE = -16
	setPos.Call(hwnd, 0, x, y, w, h, 0x0020|0x0040)
}
func newWave() (*waveOutput, error) {
	a := &waveOutput{}
	f := waveFormat{Format: 1, Channels: 1, Rate: game.SampleRate, BytesPerSecond: game.SampleRate * 2, BlockAlign: 2, Bits: 16}
	r, _, _ := winmm.NewProc("waveOutOpen").Call(uintptr(unsafe.Pointer(&a.handle)), 0xffffffff, uintptr(unsafe.Pointer(&f)), 0, 0, 0)
	if r != 0 {
		return nil, fmt.Errorf("waveOutOpen: %d", r)
	}
	for i := range a.slots {
		s := &a.slots[i]
		s.pinner.Pin(&s.header)
		s.pinner.Pin(&s.data[0])
		s.header.Data = uintptr(unsafe.Pointer(&s.data[0]))
	}
	return a, nil
}
func (a *waveOutput) write(samples []int16) error {
	for len(samples) > 0 {
		s := &a.slots[a.index%len(a.slots)]
		if s.prepared {
			if atomic.LoadUint32(&s.header.Flags)&1 == 0 {
				return nil
			}
			r, _, _ := winmm.NewProc("waveOutUnprepareHeader").Call(a.handle, uintptr(unsafe.Pointer(&s.header)), unsafe.Sizeof(s.header))
			if r != 0 {
				return fmt.Errorf("waveOutUnprepareHeader: %d", r)
			}
			s.prepared = false
		}
		n := len(samples)
		if n > len(s.data) {
			n = len(s.data)
		}
		copy(s.data[:], samples[:n])
		s.header.Length = uint32(n * 2)
		s.header.Flags = 0
		r, _, _ := winmm.NewProc("waveOutPrepareHeader").Call(a.handle, uintptr(unsafe.Pointer(&s.header)), unsafe.Sizeof(s.header))
		if r != 0 {
			return fmt.Errorf("waveOutPrepareHeader: %d", r)
		}
		s.prepared = true
		r, _, _ = winmm.NewProc("waveOutWrite").Call(a.handle, uintptr(unsafe.Pointer(&s.header)), unsafe.Sizeof(s.header))
		if r != 0 {
			return fmt.Errorf("waveOutWrite: %d", r)
		}
		a.index++
		samples = samples[n:]
	}
	return nil
}
func (a *waveOutput) close() {
	winmm.NewProc("waveOutReset").Call(a.handle)
	for i := range a.slots {
		s := &a.slots[i]
		if s.prepared {
			winmm.NewProc("waveOutUnprepareHeader").Call(a.handle, uintptr(unsafe.Pointer(&s.header)), unsafe.Sizeof(s.header))
		}
		s.pinner.Unpin()
	}
	winmm.NewProc("waveOutClose").Call(a.handle)
}

type joyInfo struct{ Size, Flags, X, Y, Z, R, U, V, Buttons, ButtonNumber, POV, Reserved1, Reserved2 uint32 }

func Run(m *game.Machine, save func() error) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	current = m
	defer func() { current = nil }()
	surface := &windowSurface{}
	currentSurface = surface
	defer func() { currentSurface = nil; surface.close() }()
	// Keep one source pixel grid; GDI scaling below explicitly disables filtering.
	user.NewProc("SetProcessDPIAware").Call()
	instance, _, _ := syscall.NewLazyDLL("kernel32.dll").NewProc("GetModuleHandleW").Call(0)
	name := utf("CDManGoWindow")
	wc := windowClass{Proc: syscall.NewCallback(windowProc), Instance: instance, ClassName: name}
	wc.Size = uint32(unsafe.Sizeof(wc))
	atom, _, err := user.NewProc("RegisterClassExW").Call(uintptr(unsafe.Pointer(&wc)))
	if atom == 0 {
		return fmt.Errorf("RegisterClassEx: %v", err)
	}
	defer user.NewProc("UnregisterClassW").Call(uintptr(unsafe.Pointer(name)), instance)
	hwnd, _, err := user.NewProc("CreateWindowExW").Call(0, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(utf("CD-Man 2.0"))), 0x80000000, 0, 0, 960, 720, 0, 0, instance, 0)
	if hwnd == 0 {
		return fmt.Errorf("CreateWindow: %v", err)
	}
	defer destroy.Call(hwnd)
	setWindowMode(hwnd)
	user.NewProc("ShowWindow").Call(hwnd, 5)
	audio, err := newWave()
	if err != nil {
		return err
	}
	defer audio.close()
	peek := user.NewProc("PeekMessageW")
	dispatch := user.NewProc("DispatchMessageW")
	translate := user.NewProc("TranslateMessage")
	deadline := time.Now()
	var pixels []byte
	for !m.Halted() {
		var msg message
		for {
			ok, _, _ := peek.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0, 1)
			if ok == 0 {
				break
			}
			if msg.Message == 0x12 {
				return nil
			}
			translate.Call(uintptr(unsafe.Pointer(&msg)))
			dispatch.Call(uintptr(unsafe.Pointer(&msg)))
		}
		var joy joyInfo
		joy.Size = uint32(unsafe.Sizeof(joy))
		joy.Flags = 0xff
		r, _, _ := winmm.NewProc("joyGetPosEx").Call(0, uintptr(unsafe.Pointer(&joy)))
		m.Joystick(r == 0, uint16(joy.X), uint16(joy.Y), byte(joy.Buttons&3))
		m.RunCycles(game.ClockHz / 70)
		if err := audio.write(m.Audio()); err != nil {
			return err
		}
		if err := save(); err != nil {
			return err
		}
		im := m.Frame()
		w, h := im.Bounds().Dx(), im.Bounds().Dy()
		if len(pixels) != w*h*4 {
			pixels = make([]byte, w*h*4)
		}
		for i := 0; i < len(pixels); i += 4 {
			pixels[i] = im.Pix[i+2]
			pixels[i+1] = im.Pix[i+1]
			pixels[i+2] = im.Pix[i]
			pixels[i+3] = 0
		}
		if err := surface.draw(hwnd, pixels, w, h); err != nil {
			return err
		}
		deadline = deadline.Add(time.Second / 70)
		if delay := time.Until(deadline); delay > 0 {
			time.Sleep(delay)
		} else if delay < -time.Second/4 {
			deadline = time.Now()
		}
	}
	return m.Error()
}
