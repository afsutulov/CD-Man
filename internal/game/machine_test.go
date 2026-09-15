package game

import (
	"cdman-go/internal/assets"
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"
)

func testMachine(t *testing.T) *Machine {
	t.Helper()
	files, err := assets.Files()
	if err != nil {
		t.Fatal(err)
	}
	m, err := New(assets.Initial, files, time.Date(1992, 7, 20, 12, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	return m
}

// Exhaustive byte arithmetic checks distinguish signed overflow from carry.
// Those flags decide branches in collision and menu routines.
func TestByteArithmetic(t *testing.T) {
	m := &Machine{}
	for a := 0; a < 256; a++ {
		for b := 0; b < 256; b++ {
			v := m.alu("add", 8, uint16(a), uint16(b))
			signed := int(int8(a)) + int(int8(b))
			if int(v) != (a+b)&255 || m.cf != (a+b > 255) || m.of != (signed < -128 || signed > 127) || m.zf != (v == 0) {
				t.Fatalf("add %d %d", a, b)
			}
			v = m.alu("sub", 8, uint16(a), uint16(b))
			signed = int(int8(a)) - int(int8(b))
			if int(v) != (a-b)&255 || m.cf != (a < b) || m.of != (signed < -128 || signed > 127) || m.zf != (v == 0) {
				t.Fatalf("sub %d %d", a, b)
			}
		}
	}
}

// Verify that all resource images retain their exact pixel values.
func TestGraphicsResources(t *testing.T) {
	images := []struct{ Source, Hash string }{
		{"TITLE.CDM", "5d378e158a5a3482f7a4ddf2708ffcb5807f73679d5084c5580e6a6772e35f8c"},
		{"WOSHI.CDM", "9887612860c49e15b27128d640c1911d4596996446f402d1e0611ea26cd32852"},
		{"SCREEN1.CDM", "622e8be1b6131aec0a47e9c67c657c0408632057be3ad12a2d79a58229459906"},
		{"SCREEN2.CDM", "84e69013f2f7f7c3ff7712ad4cafaa0d659df7060a10df526ea9eec6d9b0c7ba"},
		{"SCREEN3.CDM", "bfb1b9d741349c89ad642284eb45b6ff5500ee1ee808bf390d05e4a357465bf1"},
		{"SCREEN4.CDM", "17e4003199a9b71d272979a5f9a81f0216904e70a2b910dc67dd20597ba83813"},
		{"SCREEN5.CDM", "7e131fbf4e49f5649d9625e61d6fa13c876c1686e725e56e8af7cb5ecffb3735"},
		{"GRAPHIC1.CDM", "e2e5364b93bb441d900290076df77fc59cc20163a52754c7b2386454ae0645b6"},
		{"GRAPHIC2.CDM", "74c0d43a5aaf3a0d1d51cd803362027c11c6a3c1302388e01bc18bafdfc55537"},
		{"GRAPHIC3.CDM", "c1c9a1780557cc6eb1adf9ca8ca7d444a17f3c0d9520fa32af14618850a10fda"},
		{"GRAPHIC4.CDM", "1cafe4c40e5ace4ebdbd3a72509c4aad845f9e02b7b4d21cc148e4980ddab591"},
		{"GRAPHIC5.CDM", "f1f12766a3a91b28366951db160e3a863d23a11f6175997dc2a832e6f5cb9b17"},
		{"GRAFIX1.CDM", "ac6a3ef04edcd3b9b16bf759aa75b6ac79210df13819604b03f62f8b8d34aeba"},
		{"GRAFIX2.CDM", "4ab545d00ae3a1dc5b3ad08d1f1560c4e385885de11a0e3b85243d161b768f3d"},
		{"GRAFIX3.CDM", "a288a71254dae2033c8c82ff5598736be82e3628a5873c109a103e8b6a82561e"},
		{"GRAFIX4.CDM", "58210ac2337f3aa0aa9149001f74325937389fe66a30809c734ffcaa65075e92"},
		{"GRAFIX5.CDM", "36b681a5b0ab1ca8a8235589c531c6b06e0971fd79db27a0d37bc88a17fb5aba"},
	}
	for _, item := range images {
		t.Run(item.Source, func(t *testing.T) {
			m := testMachine(t)
			data := m.files[item.Source]
			m.setMode(0x10, false)
			copy(m.palette[:], data[:16])
			m.r[11] = dataSegment
			m.r[10] = loadSegment
			m.r[4] = 0x200
			m.r[0] = 0xa000
			m.wr16(dataSegment, 0x294, 0x3000)
			copy(m.mem[0x30000:], data[16:])
			m.push(0xffff)
			m.ip = 0x46d9
			for m.ip != 0xffff && !m.halted && m.cycles < 100_000_000 {
				m.step()
			}
			if m.err != nil {
				t.Fatal(m.err)
			}
			if m.ip != 0xffff {
				t.Fatal("graphics decoder did not return")
			}
			h := sha256.Sum256(m.Frame().Pix)
			if hex.EncodeToString(h[:]) != item.Hash {
				t.Fatal("graphics pixels differ from the expected image")
			}
		})
	}
}

func TestStartAndContinueWithoutPrompt(t *testing.T) {
	m := testMachine(t)
	keyAt := func(seconds float64, scan, ascii byte) {
		target := uint64(seconds * float64(ClockHz))
		if target > m.cycles {
			m.RunCycles(target - m.cycles)
		}
		if m.err != nil {
			t.Fatal(m.err)
		}
		m.Key(scan, ascii)
		m.Audio()
	}
	keyAt(17, 0x39, 32)
	m.RunCycles(19*ClockHz - m.cycles)
	if m.mode != 3 {
		t.Fatalf("menu not reached: mode=%x address=%x", m.mode, m.ip)
	}
	m.Key(0x1c, 13)
	// Starting a game must require no answer or extra confirmation.
	m.RunCycles(25*ClockHz - m.cycles)
	if m.err != nil {
		t.Fatal(m.err)
	}
	if m.halted || m.mode != 16 || m.Data16(0x78) == 0 {
		t.Fatal("game did not start directly")
	}
	keyAt(29, 0x4d, 0)
	m.RunCycles(31*ClockHz - m.cycles)
	if m.err != nil {
		t.Fatal(m.err)
	}
	if m.halted || m.mode != 16 {
		t.Fatal("game did not start")
	}
	if n := m.Data16(0x78); n == 0 || n >= 127 {
		t.Fatalf("no dots collected: %d", n)
	}
	m.Key(1, 27)
	m.RunCycles(2 * ClockHz)
	if m.err != nil {
		t.Fatal(m.err)
	}
	if m.mode != 3 {
		t.Fatal("Esc did not return to menu")
	}
	if !m.dualFont {
		t.Fatal("menu custom font bank not selected")
	}
	m.Key(0x1c, 13)
	m.RunCycles(2 * ClockHz)
	if m.err != nil {
		t.Fatal(m.err)
	}
	if m.halted || m.mode != 16 {
		t.Fatal("Continue did not resume the game")
	}

}
