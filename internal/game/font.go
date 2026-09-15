package game

import _ "embed"

// Standard EGA BIOS bitmap, int10_font_14 from DOSBox-X int10_memory.cpp.
// Copyright (C) 2002-2020 The DOSBox Team; GPL-2.0-or-later.
// Source: https://dosbox-x.com/doxygen/html/int10__memory_8cpp_source.html
// License text: LICENSE-font.txt.
//
//go:embed ega14.bin
var biosFont14 []byte
