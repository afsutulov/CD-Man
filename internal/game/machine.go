// Package game implements game state, drawing, input and sound.
package game

import (
	"fmt"
	"image"
	"image/color"
	"time"
)

const (
	loadSegment uint16 = 0x1000
	dataSegment uint16 = loadSegment + 0x20
	codeSegment uint16 = loadSegment + 0x376
	ClockHz     uint64 = 8_000_000
)

type fileHandle struct {
	name     string
	pos      int
	writable bool
	changed  bool
}

type Machine struct {
	r                                      [12]uint16
	ip                                     uint16
	cf, zf, sf, of, pf, af, df, interrupts bool
	mem                                    [1 << 20]byte
	video                                  [4][65536]byte
	latch                                  [4]byte
	gc                                     [16]byte
	gcIndex                                byte
	palette                                [16]byte
	attrIndex                              byte
	attrData                               bool
	mode                                   byte
	font                                   [256][32]byte
	fontHeight                             int
	dualFont                               bool
	cursorX, cursorY                       int
	cursorStart, cursorEnd                 byte
	blink                                  bool
	cycles                                 uint64
	instructions                           uint64
	halted                                 bool
	err                                    error
	files                                  map[string][]byte
	handles                                map[uint16]*fileHandle
	nextHandle                             uint16
	closedWrites                           map[string][]byte
	keys                                   []uint16
	date                                   time.Time
	speaker                                byte
	pitDivider                             uint16
	pitLow                                 byte
	pitHighNext                            bool
	audioCycle                             uint64
	audioPhase                             uint64
	audio                                  []int16
	joystick                               bool
	joyX, joyY                             uint16
	joyButtons                             byte
	joyLatch                               uint64
	trace                                  []uint16
	TraceEnabled                           bool
}

func New(initial []byte, resources map[string][]byte, epoch time.Time) (*Machine, error) {
	if len(initial) != 14176 {
		return nil, fmt.Errorf("unexpected game state size %d", len(initial))
	}
	m := &Machine{files: map[string][]byte{}, handles: map[uint16]*fileHandle{}, closedWrites: map[string][]byte{}, nextHandle: 5, date: epoch, interrupts: true, fontHeight: 14}
	copy(m.mem[uint32(loadSegment)<<4:], initial)
	// Plane write masks used by the graphics decoder.
	copy(m.mem[(uint32(codeSegment)<<4)+0x477a:], []byte{1, 7, 1, 11, 1, 13, 1, 14})
	// BIOS data area: color CRTC base used by the palette routine.
	m.mem[0x463] = 0xd4
	m.mem[0x464] = 0x03
	m.r[9] = codeSegment
	m.r[10] = loadSegment
	m.r[11] = loadSegment - 16
	m.r[8] = loadSegment - 16
	m.r[4] = 0x200
	for n, b := range resources {
		m.files[n] = append([]byte(nil), b...)
	}
	if b, ok := m.files["CHARTAB.CDM"]; ok {
		for ch := 0; ch < 254; ch++ {
			copy(m.font[ch][:14], b[ch*14:(ch+1)*14])
		}
	}
	m.setMode(3, false)
	return m, nil
}

func (m *Machine) fail(s string) {
	if m.err == nil {
		m.err = fmt.Errorf("%s at CS:%04x DS:%04x, AX:%04x BX:%04x CX:%04x DX:%04x", s, m.ip, m.r[11], m.r[0], m.r[3], m.r[1], m.r[2])
	}
	m.halted = true
}
func (m *Machine) Error() error                { return m.err }
func (m *Machine) Halted() bool                { return m.halted }
func (m *Machine) Cycles() uint64              { return m.cycles }
func (m *Machine) Address() uint16             { return m.ip }
func (m *Machine) Instructions() uint64        { return m.instructions }
func (m *Machine) Mode() byte                  { return m.mode }
func (m *Machine) Data8(offset uint16) byte    { return byte(m.rd8(dataSegment, offset)) }
func (m *Machine) Data16(offset uint16) uint16 { return m.rd16(dataSegment, offset) }
func (m *Machine) Trace() []uint16             { return append([]uint16(nil), m.trace...) }
func (m *Machine) Key(scan, ascii byte) {
	if len(m.keys) < 64 {
		m.keys = append(m.keys, uint16(scan)<<8|uint16(ascii))
	}
}
func (m *Machine) Joystick(present bool, x, y uint16, buttons byte) {
	m.joystick = present
	m.joyX = x
	m.joyY = y
	m.joyButtons = buttons
}

// RunCycles advances deterministic virtual time. Host frame rate does not set
// instruction speed. Cycle costs are an approximation, not an exact PC model.
func (m *Machine) RunCycles(budget uint64) {
	target := m.cycles + budget
	for !m.halted && m.cycles < target {
		if m.TraceEnabled {
			m.trace = append(m.trace, m.ip)
		}
		m.step()
		m.instructions++
	}
}

func (m *Machine) DrainWrites() map[string][]byte {
	r := m.closedWrites
	m.closedWrites = map[string][]byte{}
	return r
}
func (m *Machine) set8(reg, shift int, v uint16) {
	mask := uint16(255) << uint(shift)
	m.r[reg] = (m.r[reg] &^ mask) | ((v & 255) << uint(shift))
}
func (m *Machine) rd8(seg, off uint16) uint16 {
	a := (uint32(seg)*16 + uint32(off)) & 0xfffff
	if a >= 0xa0000 && a < 0xb0000 {
		o := uint16(a - 0xa0000)
		for p := 0; p < 4; p++ {
			m.latch[p] = m.video[p][o]
		}
		if m.gc[5]&8 == 0 {
			return uint16(m.latch[m.gc[4]&3])
		}
		v := byte(255)
		for p := 0; p < 4; p++ {
			if m.gc[7]&(1<<uint(p)) != 0 {
				c := byte(0)
				if m.gc[2]&(1<<uint(p)) != 0 {
					c = 255
				}
				v &^= m.latch[p] ^ c
			}
		}
		return uint16(v)
	}
	return uint16(m.mem[a])
}
func (m *Machine) rd16(seg, off uint16) uint16 {
	a := m.rd8(seg, off)
	b := m.rd8(seg, off+1)
	return a | (b << 8)
}
func (m *Machine) wr8(seg, off, v uint16) {
	a := (uint32(seg)*16 + uint32(off)) & 0xfffff
	if a >= 0xa0000 && a < 0xb0000 {
		o := uint16(a - 0xa0000)
		value := byte(v)
		rotation := m.gc[3] & 7
		rotated := (value >> rotation) | (value << ((8 - rotation) & 7))
		for p := 0; p < 4; p++ {
			mask := m.gc[8]
			src := rotated
			switch m.gc[5] & 3 {
			case 0:
				if m.gc[1]&(1<<uint(p)) != 0 {
					src = 0
					if m.gc[0]&(1<<uint(p)) != 0 {
						src = 255
					}
				}
			case 1:
				m.video[p][o] = m.latch[p]
				continue
			case 2:
				src = 0
				if value&(1<<uint(p)) != 0 {
					src = 255
				}
			case 3:
				mask &= rotated
				src = 0
				if m.gc[0]&(1<<uint(p)) != 0 {
					src = 255
				}
			}
			switch (m.gc[3] >> 3) & 3 {
			case 1:
				src &= m.latch[p]
			case 2:
				src |= m.latch[p]
			case 3:
				src ^= m.latch[p]
			}
			m.video[p][o] = (src & mask) | (m.latch[p] &^ mask)
		}
		return
	}
	m.mem[a] = byte(v)
}
func (m *Machine) wr16(seg, off, v uint16) { m.wr8(seg, off, v); m.wr8(seg, off+1, v>>8) }
func (m *Machine) push(v uint16)           { m.r[4] -= 2; m.wr16(m.r[10], m.r[4], v) }
func (m *Machine) pop() uint16             { v := m.rd16(m.r[10], m.r[4]); m.r[4] += 2; return v }

func (m *Machine) flags(v uint32, width int) {
	mask := uint32(0xffff)
	sign := uint32(0x8000)
	if width == 8 {
		mask = 255
		sign = 128
	}
	v &= mask
	m.zf = v == 0
	m.sf = v&sign != 0
	b := byte(v)
	b ^= b >> 4
	b ^= b >> 2
	b ^= b >> 1
	m.pf = b&1 == 0
}
func (m *Machine) alu(op string, width int, av, bv uint16) uint16 {
	mask := uint32(65535)
	sign := uint32(32768)
	if width == 8 {
		mask = 255
		sign = 128
	}
	a, b := uint32(av)&mask, uint32(bv)&mask
	var v uint32
	switch op {
	case "add":
		v = a + b
		m.cf = v > mask
		m.of = (^(a ^ b) & (a ^ v) & sign) != 0
		m.af = (a^b^v)&16 != 0
	case "sub":
		v = (a - b) & mask
		m.cf = a < b
		m.of = ((a ^ b) & (a ^ v) & sign) != 0
		m.af = (a^b^v)&16 != 0
	case "and":
		v = a & b
		m.cf = false
		m.of = false
	case "or":
		v = a | b
		m.cf = false
		m.of = false
	case "xor":
		v = a ^ b
		m.cf = false
		m.of = false
	default:
		m.fail("unknown ALU operation")
	}
	m.flags(v, width)
	return uint16(v & mask)
}
func (m *Machine) unary(op string, width int, v uint16) uint16 {
	carry := m.cf
	var r uint16
	switch op {
	case "inc":
		r = m.alu("add", width, v, 1)
		m.cf = carry
	case "dec":
		r = m.alu("sub", width, v, 1)
		m.cf = carry
	case "neg":
		r = m.alu("sub", width, 0, v)
	}
	return r
}
func (m *Machine) shift(op string, width int, value, count uint16) uint16 {
	mask := uint32(65535)
	sign := uint32(32768)
	if width == 8 {
		mask = 255
		sign = 128
	}
	v := uint32(value) & mask
	original := v
	for i := uint16(0); i < count; i++ {
		switch op {
		case "shl":
			m.cf = v&sign != 0
			v = (v << 1) & mask
		case "shr":
			m.cf = v&1 != 0
			v >>= 1
		case "rol":
			m.cf = v&sign != 0
			v = (v << 1) & mask
			if m.cf {
				v |= 1
			}
		case "ror":
			m.cf = v&1 != 0
			v >>= 1
			if m.cf {
				v |= sign
			}
		case "rcl":
			old := m.cf
			m.cf = v&sign != 0
			v = (v << 1) & mask
			if old {
				v |= 1
			}
		}
	}
	if count != 0 && (op == "shr" || op == "shl") {
		m.flags(v, width)
	}
	if count == 1 {
		switch op {
		case "shl", "rol", "rcl":
			m.of = (v&sign != 0) != m.cf
		case "shr":
			m.of = original&sign != 0
		case "ror":
			m.of = (v&sign != 0) != (v&(sign>>1) != 0)
		}
	}
	return uint16(v)
}
func (m *Machine) multiplyDivide(op string, v uint16) {
	if op == "mul" {
		p := uint32(m.r[0]) * uint32(v)
		m.r[0] = uint16(p)
		m.r[2] = uint16(p >> 16)
		m.cf = m.r[2] != 0
		m.of = m.cf
		return
	}
	n := (uint32(m.r[2]) << 16) | uint32(m.r[0])
	if v == 0 || n/uint32(v) > 65535 {
		m.fail("division overflow")
		return
	}
	m.r[0] = uint16(n / uint32(v))
	m.r[2] = uint16(n % uint32(v))
}
func (m *Machine) stringOp(op string, width int) {
	step := uint16(width / 8)
	if m.df {
		step = 0 - step
	}
	for m.r[1] != 0 {
		v := m.r[0]
		if op == "movs" {
			if width == 8 {
				v = m.rd8(m.r[11], m.r[6])
			} else {
				v = m.rd16(m.r[11], m.r[6])
			}
			m.r[6] += step
		}
		if width == 8 {
			m.wr8(m.r[8], m.r[7], v)
		} else {
			m.wr16(m.r[8], m.r[7], v)
		}
		m.r[7] += step
		m.r[1]--
		m.cycles += 18
	}
}

func (m *Machine) setMode(v byte, preserve bool) {
	m.mode = v
	m.gc = [16]byte{}
	m.gc[7] = 15
	m.gc[8] = 255
	m.latch = [4]byte{}
	m.dualFont = false
	m.palette = [16]byte{0, 1, 2, 3, 4, 5, 20, 7, 56, 57, 58, 59, 60, 61, 62, 63}
	if !preserve {
		if v == 0x10 {
			m.video = [4][65536]byte{}
		} else {
			for i := 0; i < 80*25; i++ {
				m.mem[0xb8000+2*i] = 32
				m.mem[0xb8000+2*i+1] = 7
			}
		}
	}
	m.cursorX = 0
	m.cursorY = 0
	m.cursorStart = 12
	m.cursorEnd = 13
}
func ega(v byte) color.RGBA {
	return color.RGBA{((v>>2)&1)*170 + ((v>>5)&1)*85, ((v>>1)&1)*170 + ((v>>4)&1)*85, (v&1)*170 + ((v>>3)&1)*85, 255}
}

// Frame returns the graphics raster or the 9-dot EGA text raster.
// The custom text font is loaded by the video service.
func (m *Machine) Frame() *image.RGBA {
	w, h := 640, 350
	if m.mode == 3 {
		w = 720
		h = 25 * m.fontHeight
	}
	im := image.NewRGBA(image.Rect(0, 0, w, h))
	var colors [16]color.RGBA
	for i, v := range m.palette {
		colors[i] = ega(v)
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			var idx byte
			if m.mode == 0x10 {
				o := y*80 + x/8
				bit := uint(7 - x%8)
				for p := 0; p < 4; p++ {
					idx |= ((m.video[p][o] >> bit) & 1) << uint(p)
				}
			} else {
				col, row := x/9, y/m.fontHeight
				a := 0xb8000 + 2*(row*80+col)
				ch, attr := m.mem[a], m.mem[a+1]
				px, py := x%9, y%m.fontHeight
				glyph := biosFont14[int(ch)*14+py]
				if m.dualFont && attr&8 != 0 {
					glyph = m.font[ch][py]
				}
				on := false
				if px < 8 {
					on = glyph&(1<<uint(7-px)) != 0
				} else if ch >= 0xc0 && ch <= 0xdf {
					on = glyph&1 != 0
				}
				if m.blink && attr&128 != 0 && (m.cycles/(ClockHz/2))&1 != 0 {
					on = false
				}
				if row == m.cursorY && col == m.cursorX && m.cursorStart&0x20 == 0 && py >= int(m.cursorStart&31) && py <= int(m.cursorEnd&31) && (m.cycles/(ClockHz/2))&1 == 0 {
					on = true
				}
				if on {
					idx = attr & 15
				} else {
					idx = (attr >> 4) & 15
					if m.blink {
						idx &= 7
					}
				}
			}
			c := colors[idx]
			o := y*im.Stride + x*4
			im.Pix[o] = c.R
			im.Pix[o+1] = c.G
			im.Pix[o+2] = c.B
			im.Pix[o+3] = 255
		}
	}
	return im
}
