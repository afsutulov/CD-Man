//go:build windows

package platform

import (
	"fmt"
	"unsafe"
)

var createMemoryDC = gdi.NewProc("CreateCompatibleDC")
var createBitmap = gdi.NewProc("CreateCompatibleBitmap")
var selectObject = gdi.NewProc("SelectObject")
var deleteObject = gdi.NewProc("DeleteObject")
var deleteDC = gdi.NewProc("DeleteDC")
var bitBlt = gdi.NewProc("BitBlt")
var stretchDIBits = gdi.NewProc("StretchDIBits")
var setStretchMode = gdi.NewProc("SetStretchBltMode")
var blackBlt = gdi.NewProc("PatBlt")
var beginPaint = user.NewProc("BeginPaint")
var endPaint = user.NewProc("EndPaint")
var getClientRect = user.NewProc("GetClientRect")
var getWindowDC = user.NewProc("GetDC")
var releaseWindowDC = user.NewProc("ReleaseDC")

const sourceCopy = 0x00cc0020
const blackness = 0x00000042

// PAINTSTRUCT uses 32-bit BOOLs, including on 64-bit Windows.
type paintStruct struct {
	DC                 uintptr
	Erase              int32
	Rect               rect
	Restore, IncUpdate int32
	Reserved           [32]byte
}

// All operations on the surface and window run on the window's OS thread.
// The desktop sees only BitBlt of a complete frame. In particular, clearing
// letterbox borders and scaling the game never modify the visible window.
type windowSurface struct {
	dc, bitmap, previous uintptr
	width, height        int
	ready                bool
	paintError           error
}

var currentSurface *windowSurface

func (s *windowSurface) close() {
	if s.dc != 0 {
		selectObject.Call(s.dc, s.previous)
		deleteObject.Call(s.bitmap)
		deleteDC.Call(s.dc)
	}
	s.dc, s.bitmap, s.previous = 0, 0, 0
	s.width, s.height = 0, 0
	s.ready = false
}

func (s *windowSurface) resize(windowDC uintptr, width, height int) error {
	if s.dc != 0 && s.width == width && s.height == height {
		return nil
	}
	dc, _, err := createMemoryDC.Call(windowDC)
	if dc == 0 {
		return fmt.Errorf("CreateCompatibleDC: %v", err)
	}
	// Use the window DC: a newly created memory DC has a monochrome bitmap.
	bitmap, _, err := createBitmap.Call(windowDC, uintptr(width), uintptr(height))
	if bitmap == 0 {
		deleteDC.Call(dc)
		return fmt.Errorf("CreateCompatibleBitmap: %v", err)
	}
	previous, _, err := selectObject.Call(dc, bitmap)
	if previous == 0 || previous == ^uintptr(0) {
		deleteObject.Call(bitmap)
		deleteDC.Call(dc)
		return fmt.Errorf("SelectObject: %v", err)
	}
	s.close()
	s.dc, s.bitmap, s.previous = dc, bitmap, previous
	s.width, s.height = width, height
	return nil
}

func (s *windowSurface) blit(destination uintptr) error {
	ok, _, err := bitBlt.Call(destination, 0, 0, uintptr(s.width), uintptr(s.height), s.dc, 0, 0, sourceCopy)
	if ok == 0 {
		return fmt.Errorf("BitBlt: %v", err)
	}
	return nil
}

func (s *windowSurface) draw(hwnd uintptr, pixels []byte, width, height int) error {
	if s.paintError != nil {
		return s.paintError
	}
	var rc rect
	ok, _, err := getClientRect.Call(hwnd, uintptr(unsafe.Pointer(&rc)))
	if ok == 0 {
		return fmt.Errorf("GetClientRect: %v", err)
	}
	cw, ch := int(rc.Right-rc.Left), int(rc.Bottom-rc.Top)
	if cw <= 0 || ch <= 0 {
		return nil // Minimized: retain the last complete frame.
	}
	dc, _, err := getWindowDC.Call(hwnd)
	if dc == 0 {
		return fmt.Errorf("GetDC: %v", err)
	}
	defer releaseWindowDC.Call(hwnd, dc)
	if err := s.resize(dc, cw, ch); err != nil {
		return err
	}
	dw, dh := cw, cw*3/4
	if dh > ch {
		dh, dw = ch, ch*4/3
	}
	// These two drawing operations affect only the off-screen bitmap.
	ok, _, err = blackBlt.Call(s.dc, 0, 0, uintptr(cw), uintptr(ch), blackness)
	if ok == 0 {
		return fmt.Errorf("PatBlt (back buffer): %v", err)
	}
	if dw > 0 && dh > 0 {
		setStretchMode.Call(s.dc, 3) // COLORONCOLOR: preserve pixel art.
		bi := bitmapHeader{Size: 40, Width: int32(width), Height: -int32(height), Planes: 1, BitCount: 32}
		lines, _, err := stretchDIBits.Call(s.dc, uintptr((cw-dw)/2), uintptr((ch-dh)/2), uintptr(dw), uintptr(dh),
			0, 0, uintptr(width), uintptr(height), uintptr(unsafe.Pointer(&pixels[0])), uintptr(unsafe.Pointer(&bi)), 0, sourceCopy)
		if lines == 0 || int32(lines) == -1 {
			return fmt.Errorf("StretchDIBits (back buffer): %v", err)
		}
	}
	s.ready = true
	return s.blit(dc)
}

func paintWindow(hwnd uintptr) {
	var ps paintStruct
	dc, _, _ := beginPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
	defer endPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
	if dc == 0 {
		return
	}
	if currentSurface != nil && currentSurface.ready {
		// Restore the cached frame; never erase an already rendered frame.
		if err := currentSurface.blit(dc); err != nil {
			currentSurface.paintError = err
		}
	} else {
		// There is no frame yet, e.g. during initial window creation.
		blackBlt.Call(dc, uintptr(ps.Rect.Left), uintptr(ps.Rect.Top),
			uintptr(ps.Rect.Right-ps.Rect.Left), uintptr(ps.Rect.Bottom-ps.Rect.Top), blackness)
	}
}
