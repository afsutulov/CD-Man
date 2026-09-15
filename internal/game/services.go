package game

import (
	"fmt"
	"strings"
	"time"
)

func (m *Machine) now() time.Time { return m.date.Add(time.Duration(m.cycles/8) * time.Microsecond) }
func (m *Machine) cstring(seg, off uint16, terminator byte) string {
	b := make([]byte, 0, 64)
	for n := 0; n < 65536; n++ {
		v := byte(m.rd8(seg, off))
		off++
		if v == terminator {
			return string(b)
		}
		b = append(b, v)
	}
	m.fail("unterminated DOS string")
	return ""
}
func canonical(s string) string {
	p := strings.Split(strings.ToUpper(strings.TrimSpace(s)), ".")
	for i := range p {
		p[i] = strings.TrimSpace(p[i])
	}
	return strings.Join(p, ".")
}
func (m *Machine) dosError(code uint16) { m.cf = true; m.r[0] = code }
func (m *Machine) interrupt(n uint16) {
	switch n {
	case 0x12:
		m.r[0] = 640
	case 0x1a:
		t := m.now()
		secs := uint64(t.Hour()*3600 + t.Minute()*60 + t.Second())
		us := secs*1_000_000 + uint64(t.Nanosecond()/1000)
		ticks := us * 1193182 / (65536 * 1_000_000)
		m.r[1] = uint16(ticks >> 16)
		m.r[2] = uint16(ticks)
		m.set8(0, 0, 0)
	case 0x16:
		switch m.r[0] >> 8 {
		case 1:
			m.zf = len(m.keys) == 0
			if len(m.keys) > 0 {
				m.r[0] = m.keys[0]
			}
		case 0:
			if len(m.keys) > 0 {
				m.r[0] = m.keys[0]
				m.keys = m.keys[1:]
			} else {
				m.ip -= 2
				m.cycles += ClockHz / 1000
			}
		default:
			m.fail("unsupported keyboard BIOS service")
		}
	case 0x10:
		m.videoBIOS()
	case 0x21:
		m.dos()
	default:
		m.fail(fmt.Sprintf("unsupported interrupt %02x", n))
	}
}
func (m *Machine) videoBIOS() {
	ah, al := byte(m.r[0]>>8), byte(m.r[0])
	bl, bh := byte(m.r[3]), byte(m.r[3]>>8)
	switch ah {
	case 0:
		mode := al & 127
		if mode != 3 && mode != 0x10 {
			m.fail("unsupported video mode")
			return
		}
		m.setMode(mode, al&128 != 0)
	case 1:
		m.cursorStart = byte(m.r[1] >> 8)
		m.cursorEnd = byte(m.r[1])
	case 2:
		m.cursorX = int(m.r[2] & 255)
		m.cursorY = int(m.r[2] >> 8)
	case 9:
		start := m.cursorY*80 + m.cursorX
		for i := 0; i < int(m.r[1]) && start+i < 2000; i++ {
			a := 0xb8000 + 2*(start+i)
			if a >= 0xb8000 {
				m.mem[a] = al
				m.mem[a+1] = bl
			}
		}
	case 0x10:
		switch al {
		case 0:
			if bl < 16 {
				m.palette[bl] = bh & 63
			}
		case 3:
			m.blink = bl != 0
		default:
			m.fail("unsupported palette BIOS service")
		}
	case 0x11:
		switch al {
		case 0x10:
			height := int(bh)
			if height < 1 || height > 32 {
				m.fail("invalid custom font height")
				return
			}
			m.fontHeight = height
			for i := 0; i < int(m.r[1]); i++ {
				ch := int(m.r[2]) + i
				if ch >= 256 {
					break
				}
				for y := 0; y < height; y++ {
					m.font[ch][y] = byte(m.rd8(m.r[8], m.r[5]+uint16(i*height+y)))
				}
			}
		case 3:
			if bl != 4 {
				m.fail("unsupported font bank selection")
				return
			}
			m.dualFont = true
		default:
			m.fail("unsupported font BIOS service")
		}
	case 0x12:
		if bl != 0x30 || al != 1 {
			m.fail("unsupported scanline BIOS service")
		}
	default:
		m.fail(fmt.Sprintf("unsupported video BIOS AX=%04x", m.r[0]))
	}
}
func (m *Machine) text(s string) {
	for _, b := range []byte(s) {
		switch b {
		case 13:
			m.cursorX = 0
		case 10:
			m.cursorY++
		default:
			if m.cursorY < 25 && m.cursorX < 80 {
				a := 0xb8000 + 2*(m.cursorY*80+m.cursorX)
				m.mem[a] = b
				m.mem[a+1] = 7
			}
			m.cursorX++
			if m.cursorX >= 80 {
				m.cursorX = 0
				m.cursorY++
			}
		}
		if m.cursorY >= 25 {
			copy(m.mem[0xb8000:0xb8000+24*160], m.mem[0xb8000+160:0xb8000+25*160])
			for x := 0; x < 80; x++ {
				m.mem[0xb8000+24*160+2*x] = 32
				m.mem[0xb8000+24*160+2*x+1] = 7
			}
			m.cursorY = 24
		}
	}
}
func (m *Machine) dos() {
	ah := byte(m.r[0] >> 8)
	al := byte(m.r[0])
	m.cf = false
	switch ah {
	case 0x09:
		m.text(m.cstring(m.r[11], m.r[2], '$'))
	case 0x2a:
		t := m.now()
		m.r[1] = uint16(t.Year())
		m.r[2] = uint16(t.Month())<<8 | uint16(t.Day())
		m.set8(0, 0, uint16(t.Weekday()))
	case 0x2c:
		t := m.now()
		m.r[1] = uint16(t.Hour())<<8 | uint16(t.Minute())
		m.r[2] = uint16(t.Second())<<8 | uint16(t.Nanosecond()/10_000_000)
	case 0x3c, 0x3d:
		name := canonical(m.cstring(m.r[11], m.r[2], 0))
		if name == "" || strings.ContainsAny(name, "/\\:") {
			m.dosError(3)
			return
		}
		if ah == 0x3c {
			m.files[name] = []byte{}
		} else if _, ok := m.files[name]; !ok {
			m.dosError(2)
			return
		}
		h := m.nextHandle
		m.nextHandle++
		m.handles[h] = &fileHandle{name: name, writable: ah == 0x3c || al&3 != 0, changed: ah == 0x3c}
		m.r[0] = h
	case 0x3e:
		h, ok := m.handles[m.r[3]]
		if !ok {
			m.dosError(6)
			return
		}
		if h.changed {
			m.closedWrites[h.name] = append([]byte(nil), m.files[h.name]...)
		}
		delete(m.handles, m.r[3])
	case 0x3f:
		h, ok := m.handles[m.r[3]]
		if !ok {
			m.dosError(6)
			return
		}
		b := m.files[h.name]
		n := int(m.r[1])
		if h.pos >= len(b) {
			n = 0
		} else if n > len(b)-h.pos {
			n = len(b) - h.pos
		}
		for i := 0; i < n; i++ {
			m.wr8(m.r[11], m.r[2]+uint16(i), uint16(b[h.pos+i]))
		}
		h.pos += n
		m.r[0] = uint16(n)
	case 0x40:
		h, ok := m.handles[m.r[3]]
		if !ok {
			m.dosError(6)
			return
		}
		if !h.writable {
			m.dosError(5)
			return
		}
		n := int(m.r[1])
		if h.pos+n > 2*1024*1024 {
			m.dosError(8)
			return
		}
		b := m.files[h.name]
		if n == 0 {
			if h.pos < len(b) {
				b = b[:h.pos]
			} else {
				b = append(b, make([]byte, h.pos-len(b))...)
			}
		} else {
			if h.pos+n > len(b) {
				b = append(b, make([]byte, h.pos+n-len(b))...)
			}
			for i := 0; i < n; i++ {
				b[h.pos+i] = byte(m.rd8(m.r[11], m.r[2]+uint16(i)))
			}
		}
		m.files[h.name] = b
		h.pos += n
		h.changed = true
		m.r[0] = uint16(n)
	case 0x42:
		h, ok := m.handles[m.r[3]]
		if !ok {
			m.dosError(6)
			return
		}
		delta := int64(int32(uint32(m.r[1])<<16 | uint32(m.r[2])))
		base := int64(0)
		if al == 1 {
			base = int64(h.pos)
		} else if al == 2 {
			base = int64(len(m.files[h.name]))
		} else if al != 0 {
			m.dosError(1)
			return
		}
		pos := base + delta
		if pos < 0 || pos > 2*1024*1024 {
			m.dosError(1)
			return
		}
		h.pos = int(pos)
		m.r[0] = uint16(pos)
		m.r[2] = uint16(pos >> 16)
	case 0x4c:
		m.halted = true
	default:
		m.fail(fmt.Sprintf("unsupported DOS service AH=%02x", ah))
	}
}
func (m *Machine) input(port uint16) uint16 {
	switch port {
	case 0x61:
		return uint16(m.speaker)
	case 0x3da:
		m.attrData = false
		return 8
	case 0x201:
		if !m.joystick {
			return 255
		}
		v := byte(0xf0) &^ (m.joyButtons << 4)
		elapsed := m.cycles - m.joyLatch
		if elapsed < uint64(1000+uint32(m.joyX)*12000/65535) {
			v |= 1
		}
		if elapsed < uint64(1000+uint32(m.joyY)*12000/65535) {
			v |= 2
		}
		return uint16(v)
	default:
		m.fail(fmt.Sprintf("unsupported input port %04x", port))
		return 0
	}
}
func (m *Machine) output(port, v uint16, width int) {
	if width == 16 {
		m.output(port, v&255, 8)
		m.output(port+1, v>>8, 8)
		return
	}
	switch port {
	case 0x3ce:
		m.gcIndex = byte(v) & 15
	case 0x3cf:
		m.gc[m.gcIndex] = byte(v)
	case 0x3c0:
		if !m.attrData {
			m.attrIndex = byte(v) & 31
			m.attrData = true
		} else {
			if m.attrIndex < 16 {
				m.palette[m.attrIndex] = byte(v) & 63
			}
			m.attrData = false
		}
	case 0x61:
		m.renderAudio()
		m.speaker = byte(v)
	case 0x43:
		m.renderAudio()
		if byte(v) != 0xb6 {
			m.fail("unsupported PIT mode")
			return
		}
		m.pitHighNext = false
	case 0x42:
		m.renderAudio()
		if !m.pitHighNext {
			m.pitLow = byte(v)
			m.pitHighNext = true
		} else {
			m.pitDivider = uint16(m.pitLow) | (v << 8)
			m.pitHighNext = false
			m.audioPhase = 0
		}
	case 0x201:
		m.joyLatch = m.cycles
	default:
		m.fail(fmt.Sprintf("unsupported output port %04x", port))
	}
}

const SampleRate = 48000

func (m *Machine) renderAudio() {
	end := m.cycles * SampleRate / ClockHz
	start := m.audioCycle
	if end-start > SampleRate*10 {
		start = end - SampleRate*10
	}
	divider := uint64(m.pitDivider)
	if divider == 0 {
		divider = 65536
	}
	period := divider * SampleRate
	for i := start; i < end; i++ {
		var sample int16
		if m.speaker&3 == 3 {
			if m.audioPhase < period/2 {
				sample = 6000
			} else {
				sample = -6000
			}
		}
		m.audio = append(m.audio, sample)
		m.audioPhase = (m.audioPhase + 1193182) % period
	}
	m.audioCycle = end
}
func (m *Machine) Audio() []int16 { m.renderAudio(); a := m.audio; m.audio = nil; return a }
