// Game state transitions and rendering operations.
package game

func (m *Machine) step() {
	switch m.ip {
	case 0x0000:
		{
			m.ip = 0x0001
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x0001:
		{
			m.ip = 0x0004
			m.r[0] = uint16(0)
			m.cycles += 2
		}
	case 0x0004:
		{
			m.ip = 0x0005
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x0005:
		{
			m.ip = 0x0008
			m.r[0] = dataSegment
			m.cycles += 2
		}
	case 0x0008:
		{
			m.ip = 0x000a
			m.r[11] = m.r[0]
			m.cycles += 2
		}
	case 0x000a:
		{
			m.ip = 0x000c
			m.interrupt(uint16(18))
			m.cycles += 50
		}
	case 0x000c:
		{
			m.ip = 0x000e
			m.set8(1, 0, uint16(6))
			m.cycles += 2
		}
	case 0x000e:
		{
			m.ip = 0x0010
			m.r[0] = m.shift("shl", 16, m.r[0], ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0010:
		{
			m.ip = 0x0012
			m.r[3] = m.r[9]
			m.cycles += 2
		}
	case 0x0012:
		{
			m.ip = 0x0016
			m.r[3] = m.alu("add", 16, m.r[3], uint16(4096))
			m.cycles += 4
		}
	case 0x0016:
		{
			m.ip = 0x0018
			m.r[0] = m.alu("sub", 16, m.r[0], m.r[3])
			m.cycles += 4
		}
	case 0x0018:
		{
			m.ip = 0x001b
			m.alu("sub", 16, m.r[0], uint16(4096))
			m.cycles += 4
		}
	case 0x001b:
		{
			m.ip = 0x001d
			if !m.cf && !m.zf {
				m.ip = uint16(29)
			}
			m.cycles += 8
		}
	case 0x001d:
		{
			m.ip = 0x0021
			m.wr16(m.r[11], uint16(0x28a), m.r[3])
			m.cycles += 8
		}
	case 0x0021:
		{
			m.ip = 0x0024
			m.r[0] = uint16(625)
			m.cycles += 2
		}
	case 0x0024:
		{
			m.ip = 0x0026
			m.set8(1, 0, uint16(4))
			m.cycles += 2
		}
	case 0x0026:
		{
			m.ip = 0x0028
			m.r[0] = m.shift("shr", 16, m.r[0], ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0028:
		{
			m.ip = 0x0029
			m.r[0] = m.unary("inc", 16, m.r[0])
			m.cycles += 3
		}
	case 0x0029:
		{
			m.ip = 0x002b
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x002b:
		{
			m.ip = 0x002f
			m.wr16(m.r[11], uint16(0x28c), m.r[3])
			m.cycles += 8
		}
	case 0x002f:
		{
			m.ip = 0x0032
			m.r[0] = uint16(559)
			m.cycles += 2
		}
	case 0x0032:
		{
			m.ip = 0x0034
			m.set8(1, 0, uint16(4))
			m.cycles += 2
		}
	case 0x0034:
		{
			m.ip = 0x0036
			m.r[0] = m.shift("shr", 16, m.r[0], ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0036:
		{
			m.ip = 0x0037
			m.r[0] = m.unary("inc", 16, m.r[0])
			m.cycles += 3
		}
	case 0x0037:
		{
			m.ip = 0x0039
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x0039:
		{
			m.ip = 0x003d
			m.wr16(m.r[11], uint16(0x28e), m.r[3])
			m.cycles += 8
		}
	case 0x003d:
		{
			m.ip = 0x003f
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x003f:
		{
			m.ip = 0x0043
			m.wr16(m.r[11], uint16(0x290), m.r[3])
			m.cycles += 8
		}
	case 0x0043:
		{
			m.ip = 0x0045
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x0045:
		{
			m.ip = 0x0049
			m.wr16(m.r[11], uint16(0x292), m.r[3])
			m.cycles += 8
		}
	case 0x0049:
		{
			m.ip = 0x004b
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x004b:
		{
			m.ip = 0x004f
			m.wr16(m.r[11], uint16(0x294), m.r[3])
			m.cycles += 8
		}
	case 0x004f:
		{
			m.ip = 0x0052
			target := uint16(18923)
			m.push(0x0052)
			m.ip = target
			m.cycles += 19
		}
	case 0x0052:
		{
			m.ip = 0x0057
			m.wr8(m.r[11], uint16(0x1069), uint16(0))
			m.cycles += 8
		}
	case 0x0057:
		{
			m.ip = 0x005a
			target := uint16(104)
			m.push(0x005a)
			m.ip = target
			m.cycles += 19
		}
	case 0x005a:
		{
			m.ip = 0x005d
			m.r[0] = uint16(16)
			m.cycles += 2
		}
	case 0x005d:
		{
			m.ip = 0x005f
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x005f:
		{
			m.ip = 0x0062
			target := uint16(152)
			m.push(0x0062)
			m.ip = target
			m.cycles += 19
		}
	case 0x0062:
		{
			m.ip = 0x0065
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x0065:
		{
			m.ip = 0x0067
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0067:
		{
			m.ip = 0x0068
			m.ip = m.pop()
			m.r[9] = m.pop()
			if m.r[9] != codeSegment {
				m.halted = true
			}
			m.cycles += 18
		}
	case 0x0068:
		{
			m.ip = 0x006c
			m.r[2] = uint16(0x390)
			m.cycles += 3
		}
	case 0x006c:
		{
			m.ip = 0x006e
			m.set8(0, 8, uint16(61))
			m.cycles += 2
		}
	case 0x006e:
		{
			m.ip = 0x0070
			m.set8(0, 0, uint16(2))
			m.cycles += 2
		}
	case 0x0070:
		{
			m.ip = 0x0072
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x0072:
		{
			m.ip = 0x0074
			if !m.cf {
				m.ip = uint16(134)
			}
			m.cycles += 8
		}
	case 0x0074:
		{
			m.ip = 0x0077
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x0077:
		{
			m.ip = 0x0079
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0079:
		{
			m.ip = 0x007d
			m.r[2] = uint16(0x3b3)
			m.cycles += 3
		}
	case 0x007d:
		{
			m.ip = 0x007f
			m.set8(0, 8, uint16(9))
			m.cycles += 2
		}
	case 0x007f:
		{
			m.ip = 0x0081
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x0081:
		{
			m.ip = 0x0083
			m.set8(0, 8, uint16(76))
			m.cycles += 2
		}
	case 0x0083:
		{
			m.ip = 0x0085
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x0085:
		{
			m.ip = 0x0086
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0086:
		{
			m.ip = 0x0088
			m.r[3] = m.r[0]
			m.cycles += 2
		}
	case 0x0088:
		{
			m.ip = 0x008b
			m.r[1] = uint16(3556)
			m.cycles += 2
		}
	case 0x008b:
		{
			m.ip = 0x008f
			m.r[2] = uint16(0x254d)
			m.cycles += 3
		}
	case 0x008f:
		{
			m.ip = 0x0091
			m.set8(0, 8, uint16(63))
			m.cycles += 2
		}
	case 0x0091:
		{
			m.ip = 0x0093
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x0093:
		{
			m.ip = 0x0095
			m.set8(0, 8, uint16(62))
			m.cycles += 2
		}
	case 0x0095:
		{
			m.ip = 0x0097
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x0097:
		{
			m.ip = 0x0098
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0098:
		{
			m.ip = 0x009c
			m.wr16(m.r[11], uint16(0x278), m.r[4])
			m.cycles += 8
		}
	case 0x009c:
		{
			m.ip = 0x009f
			target := uint16(9161)
			m.push(0x009f)
			m.ip = target
			m.cycles += 19
		}
	case 0x009f:
		{
			m.ip = 0x00a3
			m.wr8(m.r[11], uint16(0x139e), ((m.r[1] >> 8) & 255))
			m.cycles += 8
		}
	case 0x00a3:
		{
			m.ip = 0x00a7
			m.wr8(m.r[11], uint16(0x139f), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x00a7:
		{
			m.ip = 0x00ab
			m.wr8(m.r[11], uint16(0x13a0), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x00ab:
		{
			m.ip = 0x00ae
			m.r[0] = m.rd16(m.r[11], uint16(0x28a))
			m.cycles += 8
		}
	case 0x00ae:
		{
			m.ip = 0x00b1
			target := uint16(3035)
			m.push(0x00b1)
			m.ip = target
			m.cycles += 19
		}
	case 0x00b1:
		{
			m.ip = 0x00b4
			target := uint16(610)
			m.push(0x00b4)
			m.ip = target
			m.cycles += 19
		}
	case 0x00b4:
		{
			m.ip = 0x00b7
			target := uint16(9166)
			m.push(0x00b7)
			m.ip = target
			m.cycles += 19
		}
	case 0x00b7:
		{
			m.ip = 0x00ba
			target := uint16(348)
			m.push(0x00ba)
			m.ip = target
			m.cycles += 19
		}
	case 0x00ba:
		{
			m.ip = 0x00bd
			target := uint16(19440)
			m.push(0x00bd)
			m.ip = target
			m.cycles += 19
		}
	case 0x00bd:
		{
			m.ip = 0x00c0
			target := uint16(3488)
			m.push(0x00c0)
			m.ip = target
			m.cycles += 19
		}
	case 0x00c0:
		{
			m.ip = 0x00c3
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x00c3:
		{
			m.ip = 0x00c5
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x00c5:
		{
			m.ip = 0x00c8
			m.r[1] = uint16(8192)
			m.cycles += 2
		}
	case 0x00c8:
		{
			m.ip = 0x00ca
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x00ca:
		{
			m.ip = 0x00cc
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x00cc:
		{
			m.ip = 0x00ce
			m.set8(3, 0, uint16(0))
			m.cycles += 2
		}
	case 0x00ce:
		{
			m.ip = 0x00d1
			m.r[0] = uint16(4099)
			m.cycles += 2
		}
	case 0x00d1:
		{
			m.ip = 0x00d3
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x00d3:
		{
			m.ip = 0x00d7
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x00d7:
		{
			m.ip = 0x00da
			target := uint16(3488)
			m.push(0x00da)
			m.ip = target
			m.cycles += 19
		}
	case 0x00da:
		{
			m.ip = 0x00df
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2077)), uint16(1))
			m.cycles += 16
		}
	case 0x00df:
		{
			m.ip = 0x00e1
			if !m.zf {
				m.ip = uint16(228)
			}
			m.cycles += 8
		}
	case 0x00e1:
		{
			m.ip = 0x00e4
			target := uint16(6660)
			m.push(0x00e4)
			m.ip = target
			m.cycles += 19
		}
	case 0x00e4:
		{
			m.ip = 0x00e7
			target := uint16(3676)
			m.push(0x00e7)
			m.ip = target
			m.cycles += 19
		}
	case 0x00e7:
		{
			m.ip = 0x00ec
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x77)), uint16(1))
			m.cycles += 16
		}
	case 0x00ec:
		{
			m.ip = 0x00ee
			if m.zf {
				m.ip = uint16(341)
			}
			m.cycles += 8
		}
	case 0x00ee:
		{
			m.ip = 0x00f1
			m.r[0] = m.rd16(m.r[11], uint16(0x28a))
			m.cycles += 8
		}
	case 0x00f1:
		{
			m.ip = 0x00f4
			target := uint16(3061)
			m.push(0x00f4)
			m.ip = target
			m.cycles += 19
		}
	case 0x00f4:
		{
			m.ip = 0x00f7
			m.r[0] = uint16(16)
			m.cycles += 2
		}
	case 0x00f7:
		{
			m.ip = 0x00f9
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x00f9:
		{
			m.ip = 0x00fc
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x617)))
			m.cycles += 8
		}
	case 0x00fc:
		{
			m.ip = 0x00ff
			m.wr16(m.r[11], uint16(0x616), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x00ff:
		{
			m.ip = 0x0102
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x619)))
			m.cycles += 8
		}
	case 0x0102:
		{
			m.ip = 0x0105
			m.wr16(m.r[11], uint16(0x618), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0105:
		{
			m.ip = 0x010a
			m.wr8(m.r[11], uint16(0x615), uint16(1))
			m.cycles += 8
		}
	case 0x010a:
		{
			m.ip = 0x010f
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x010f:
		{
			m.ip = 0x0111
			if !m.zf {
				m.ip = uint16(279)
			}
			m.cycles += 8
		}
	case 0x0111:
		{
			m.ip = 0x0114
			target := uint16(1537)
			m.push(0x0114)
			m.ip = target
			m.cycles += 19
		}
	case 0x0114:
		{
			m.ip = 0x0116
			m.ip = uint16(295)
			m.cycles += 15
		}
	case 0x0116:
		{
			m.ip = 0x0117
			m.cycles += 3
		}
	case 0x0117:
		{
			m.ip = 0x011c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(1))
			m.cycles += 16
		}
	case 0x011c:
		{
			m.ip = 0x011e
			if !m.zf {
				m.ip = uint16(292)
			}
			m.cycles += 8
		}
	case 0x011e:
		{
			m.ip = 0x0121
			target := uint16(645)
			m.push(0x0121)
			m.ip = target
			m.cycles += 19
		}
	case 0x0121:
		{
			m.ip = 0x0123
			m.ip = uint16(295)
			m.cycles += 15
		}
	case 0x0123:
		{
			m.ip = 0x0124
			m.cycles += 3
		}
	case 0x0124:
		{
			m.ip = 0x0127
			target := uint16(875)
			m.push(0x0127)
			m.ip = target
			m.cycles += 19
		}
	case 0x0127:
		{
			m.ip = 0x012b
			m.r[4] = m.rd16(m.r[11], uint16(0x278))
			m.cycles += 8
		}
	case 0x012b:
		{
			m.ip = 0x012d
			m.set8(0, 0, m.input(uint16(97)))
			m.cycles += 8
		}
	case 0x012d:
		{
			m.ip = 0x012f
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(252)))
			m.cycles += 4
		}
	case 0x012f:
		{
			m.ip = 0x0131
			m.output(uint16(97), ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x0131:
		{
			m.ip = 0x0137
			m.wr16(m.r[11], uint16(0x55b), uint16(0))
			m.cycles += 8
		}
	case 0x0137:
		{
			m.ip = 0x013b
			m.r[6] = uint16(0x55f)
			m.cycles += 3
		}
	case 0x013b:
		{
			m.ip = 0x013f
			m.wr16(m.r[11], uint16(0x55d), m.r[6])
			m.cycles += 8
		}
	case 0x013f:
		{
			m.ip = 0x0144
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x77)), uint16(1))
			m.cycles += 16
		}
	case 0x0144:
		{
			m.ip = 0x0146
			if m.zf {
				m.ip = uint16(341)
			}
			m.cycles += 8
		}
	case 0x0146:
		{
			m.ip = 0x014b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x75)), uint16(1))
			m.cycles += 16
		}
	case 0x014b:
		{
			m.ip = 0x014d
			if m.zf {
				m.ip = uint16(238)
			}
			m.cycles += 8
		}
	case 0x014d:
		{
			m.ip = 0x0152
			m.wr8(m.r[11], uint16(0x615), uint16(0))
			m.cycles += 8
		}
	case 0x0152:
		{
			m.ip = 0x0155
			m.ip = uint16(192)
			m.cycles += 15
		}
	case 0x0155:
		{
			m.ip = 0x0158
			target := uint16(8813)
			m.push(0x0158)
			m.ip = target
			m.cycles += 19
		}
	case 0x0158:
		{
			m.ip = 0x015b
			target := uint16(9309)
			m.push(0x015b)
			m.ip = target
			m.cycles += 19
		}
	case 0x015b:
		{
			m.ip = 0x015c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x015c:
		{
			m.ip = 0x0160
			m.r[2] = uint16(0x24f9)
			m.cycles += 3
		}
	case 0x0160:
		{
			m.ip = 0x0162
			m.set8(0, 8, uint16(9))
			m.cycles += 2
		}
	case 0x0162:
		{
			m.ip = 0x0164
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x0164:
		{
			m.ip = 0x0167
			target := uint16(9161)
			m.push(0x0167)
			m.ip = target
			m.cycles += 19
		}
	case 0x0167:
		{
			m.ip = 0x016b
			m.wr16(m.r[11], uint16(0x24f7), m.r[2])
			m.cycles += 8
		}
	case 0x016b:
		{
			m.ip = 0x016f
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x016f:
		{
			m.ip = 0x0172
			m.r[1] = uint16(16)
			m.cycles += 2
		}
	case 0x0172:
		{
			m.ip = 0x0173
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x0173:
		{
			m.ip = 0x0176
			m.r[1] = uint16(40000)
			m.cycles += 2
		}
	case 0x0176:
		{
			m.ip = 0x0179
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x0179:
		{
			m.ip = 0x017b
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(374)
			}
			m.cycles += 17
		}
	case 0x017b:
		{
			m.ip = 0x017c
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x017c:
		{
			m.ip = 0x017e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(370)
			}
			m.cycles += 17
		}
	case 0x017e:
		{
			m.ip = 0x0181
			target := uint16(557)
			m.push(0x0181)
			m.ip = target
			m.cycles += 19
		}
	case 0x0181:
		{
			m.ip = 0x0184
			m.alu("sub", 16, m.r[3], uint16(110))
			m.cycles += 4
		}
	case 0x0184:
		{
			m.ip = 0x0186
			if !m.cf && !m.zf {
				m.ip = uint16(406)
			}
			m.cycles += 8
		}
	case 0x0186:
		{
			m.ip = 0x0188
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x0188:
		{
			m.ip = 0x018a
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x018a:
		{
			m.ip = 0x018c
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x018c:
		{
			m.ip = 0x018e
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x018e:
		{
			m.ip = 0x0191
			m.r[2] = uint16(2448)
			m.cycles += 2
		}
	case 0x0191:
		{
			m.ip = 0x0193
			m.r[2] = m.alu("sub", 16, m.r[2], m.r[3])
			m.cycles += 4
		}
	case 0x0193:
		{
			m.ip = 0x0195
			m.ip = uint16(482)
			m.cycles += 15
		}
	case 0x0195:
		{
			m.ip = 0x0196
			m.cycles += 3
		}
	case 0x0196:
		{
			m.ip = 0x019a
			m.alu("sub", 16, m.r[3], uint16(182))
			m.cycles += 4
		}
	case 0x019a:
		{
			m.ip = 0x019c
			if !m.cf && !m.zf {
				m.ip = uint16(424)
			}
			m.cycles += 8
		}
	case 0x019c:
		{
			m.ip = 0x019e
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x019e:
		{
			m.ip = 0x01a0
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x01a0:
		{
			m.ip = 0x01a3
			m.r[2] = uint16(1128)
			m.cycles += 2
		}
	case 0x01a3:
		{
			m.ip = 0x01a5
			m.r[2] = m.alu("sub", 16, m.r[2], m.r[3])
			m.cycles += 4
		}
	case 0x01a5:
		{
			m.ip = 0x01a7
			m.ip = uint16(482)
			m.cycles += 15
		}
	case 0x01a7:
		{
			m.ip = 0x01a8
			m.cycles += 3
		}
	case 0x01a8:
		{
			m.ip = 0x01ac
			m.alu("sub", 16, m.r[3], uint16(246))
			m.cycles += 4
		}
	case 0x01ac:
		{
			m.ip = 0x01ae
			if !m.cf && !m.zf {
				m.ip = uint16(440)
			}
			m.cycles += 8
		}
	case 0x01ae:
		{
			m.ip = 0x01b0
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x01b0:
		{
			m.ip = 0x01b3
			m.r[2] = uint16(765)
			m.cycles += 2
		}
	case 0x01b3:
		{
			m.ip = 0x01b5
			m.r[2] = m.alu("sub", 16, m.r[2], m.r[3])
			m.cycles += 4
		}
	case 0x01b5:
		{
			m.ip = 0x01b7
			m.ip = uint16(482)
			m.cycles += 15
		}
	case 0x01b7:
		{
			m.ip = 0x01b8
			m.cycles += 3
		}
	case 0x01b8:
		{
			m.ip = 0x01bc
			m.alu("sub", 16, m.r[3], uint16(332))
			m.cycles += 4
		}
	case 0x01bc:
		{
			m.ip = 0x01be
			if !m.cf && !m.zf {
				m.ip = uint16(454)
			}
			m.cycles += 8
		}
	case 0x01be:
		{
			m.ip = 0x01c1
			m.r[2] = uint16(519)
			m.cycles += 2
		}
	case 0x01c1:
		{
			m.ip = 0x01c3
			m.r[2] = m.alu("sub", 16, m.r[2], m.r[3])
			m.cycles += 4
		}
	case 0x01c3:
		{
			m.ip = 0x01c5
			m.ip = uint16(482)
			m.cycles += 15
		}
	case 0x01c5:
		{
			m.ip = 0x01c6
			m.cycles += 3
		}
	case 0x01c6:
		{
			m.ip = 0x01ca
			m.alu("sub", 16, m.r[3], uint16(456))
			m.cycles += 4
		}
	case 0x01ca:
		{
			m.ip = 0x01cc
			if !m.cf && !m.zf {
				m.ip = uint16(470)
			}
			m.cycles += 8
		}
	case 0x01cc:
		{
			m.ip = 0x01ce
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x01ce:
		{
			m.ip = 0x01d1
			m.r[2] = uint16(353)
			m.cycles += 2
		}
	case 0x01d1:
		{
			m.ip = 0x01d3
			m.r[2] = m.alu("sub", 16, m.r[2], m.r[3])
			m.cycles += 4
		}
	case 0x01d3:
		{
			m.ip = 0x01d5
			m.ip = uint16(482)
			m.cycles += 15
		}
	case 0x01d5:
		{
			m.ip = 0x01d6
			m.cycles += 3
		}
	case 0x01d6:
		{
			m.ip = 0x01d8
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x01d8:
		{
			m.ip = 0x01da
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x01da:
		{
			m.ip = 0x01dd
			m.r[2] = uint16(239)
			m.cycles += 2
		}
	case 0x01dd:
		{
			m.ip = 0x01df
			m.r[2] = m.alu("sub", 16, m.r[2], m.r[3])
			m.cycles += 4
		}
	case 0x01df:
		{
			m.ip = 0x01e1
			m.ip = uint16(482)
			m.cycles += 15
		}
	case 0x01e1:
		{
			m.ip = 0x01e2
			m.cycles += 3
		}
	case 0x01e2:
		{
			m.ip = 0x01e6
			m.wr16(m.r[11], uint16(0x106c), m.r[2])
			m.cycles += 8
		}
	case 0x01e6:
		{
			m.ip = 0x01e8
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x01e8:
		{
			m.ip = 0x01ea
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x01ea:
		{
			m.ip = 0x01ec
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x01ec:
		{
			m.ip = 0x01ee
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x01ee:
		{
			m.ip = 0x01f0
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x01f0:
		{
			m.ip = 0x01f2
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x01f2:
		{
			m.ip = 0x01f6
			m.wr16(m.r[11], uint16(0x106e), m.r[2])
			m.cycles += 8
		}
	case 0x01f6:
		{
			m.ip = 0x01fb
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x106e)), uint16(0))
			m.cycles += 16
		}
	case 0x01fb:
		{
			m.ip = 0x01fd
			if !m.zf {
				m.ip = uint16(515)
			}
			m.cycles += 8
		}
	case 0x01fd:
		{
			m.ip = 0x0203
			m.wr16(m.r[11], uint16(0x106e), uint16(1))
			m.cycles += 8
		}
	case 0x0203:
		{
			m.ip = 0x0208
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x106c)), uint16(100))
			m.cycles += 16
		}
	case 0x0208:
		{
			m.ip = 0x020a
			if m.cf {
				m.ip = uint16(550)
			}
			m.cycles += 8
		}
	case 0x020a:
		{
			m.ip = 0x020e
			m.wr8(m.r[11], uint16(0x49d), m.shift("shr", 8, m.rd8(m.r[11], uint16(0x49d)), uint16(1)))
			m.cycles += 8
		}
	case 0x020e:
		{
			m.ip = 0x0214
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x106c)), uint16(200))
			m.cycles += 16
		}
	case 0x0214:
		{
			m.ip = 0x0216
			if m.cf {
				m.ip = uint16(550)
			}
			m.cycles += 8
		}
	case 0x0216:
		{
			m.ip = 0x021a
			m.wr8(m.r[11], uint16(0x49d), m.shift("shr", 8, m.rd8(m.r[11], uint16(0x49d)), uint16(1)))
			m.cycles += 8
		}
	case 0x021a:
		{
			m.ip = 0x0220
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x106c)), uint16(300))
			m.cycles += 16
		}
	case 0x0220:
		{
			m.ip = 0x0222
			if m.cf {
				m.ip = uint16(550)
			}
			m.cycles += 8
		}
	case 0x0222:
		{
			m.ip = 0x0226
			m.wr8(m.r[11], uint16(0x49d), m.shift("shr", 8, m.rd8(m.r[11], uint16(0x49d)), uint16(1)))
			m.cycles += 8
		}
	case 0x0226:
		{
			m.ip = 0x0229
			m.r[0] = m.rd16(m.r[11], uint16(0x106c))
			m.cycles += 8
		}
	case 0x0229:
		{
			m.ip = 0x022c
			m.wr16(m.r[11], uint16(0x497), m.r[0])
			m.cycles += 8
		}
	case 0x022c:
		{
			m.ip = 0x022d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x022d:
		{
			m.ip = 0x0230
			target := uint16(9161)
			m.push(0x0230)
			m.ip = target
			m.cycles += 19
		}
	case 0x0230:
		{
			m.ip = 0x0233
			m.r[0] = m.rd16(m.r[11], uint16(0x24f7))
			m.cycles += 8
		}
	case 0x0233:
		{
			m.ip = 0x0235
			m.set8(3, 8, uint16(0))
			m.cycles += 2
		}
	case 0x0235:
		{
			m.ip = 0x0237
			m.set8(3, 0, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x0237:
		{
			m.ip = 0x0239
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), ((m.r[2] >> 8) & 255))
			m.cycles += 4
		}
	case 0x0239:
		{
			m.ip = 0x023b
			if m.zf {
				m.ip = uint16(605)
			}
			m.cycles += 8
		}
	case 0x023b:
		{
			m.ip = 0x023d
			if m.sf != m.of {
				m.ip = uint16(582)
			}
			m.cycles += 8
		}
	case 0x023d:
		{
			m.ip = 0x023f
			m.set8(0, 8, m.unary("neg", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x023f:
		{
			m.ip = 0x0242
			m.set8(0, 8, m.alu("add", 8, ((m.r[0]>>8)&255), uint16(60)))
			m.cycles += 4
		}
	case 0x0242:
		{
			m.ip = 0x0244
			m.set8(2, 8, m.alu("add", 8, ((m.r[2]>>8)&255), ((m.r[0]>>8)&255)))
			m.cycles += 4
		}
	case 0x0244:
		{
			m.ip = 0x0246
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x0246:
		{
			m.ip = 0x0248
			m.set8(2, 8, m.unary("dec", 8, ((m.r[2]>>8)&255)))
			m.cycles += 3
		}
	case 0x0248:
		{
			m.ip = 0x024a
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), ((m.r[2] >> 8) & 255))
			m.cycles += 4
		}
	case 0x024a:
		{
			m.ip = 0x024c
			if m.zf {
				m.ip = uint16(593)
			}
			m.cycles += 8
		}
	case 0x024c:
		{
			m.ip = 0x024f
			m.r[3] = m.alu("add", 16, m.r[3], uint16(100))
			m.cycles += 4
		}
	case 0x024f:
		{
			m.ip = 0x0251
			m.ip = uint16(582)
			m.cycles += 15
		}
	case 0x0251:
		{
			m.ip = 0x0253
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x0253:
		{
			m.ip = 0x0255
			m.r[0] = m.unary("neg", 16, m.r[0])
			m.cycles += 3
		}
	case 0x0255:
		{
			m.ip = 0x0258
			m.r[0] = m.alu("add", 16, m.r[0], uint16(99))
			m.cycles += 4
		}
	case 0x0258:
		{
			m.ip = 0x025a
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x025a:
		{
			m.ip = 0x025c
			m.ip = uint16(609)
			m.cycles += 15
		}
	case 0x025c:
		{
			m.ip = 0x025d
			m.cycles += 3
		}
	case 0x025d:
		{
			m.ip = 0x025f
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x025f:
		{
			m.ip = 0x0261
			m.r[3] = m.alu("sub", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x0261:
		{
			m.ip = 0x0262
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0262:
		{
			m.ip = 0x0265
			m.r[0] = uint16(4642)
			m.cycles += 2
		}
	case 0x0265:
		{
			m.ip = 0x0268
			m.r[7] = uint16(44)
			m.cycles += 2
		}
	case 0x0268:
		{
			m.ip = 0x026b
			m.r[1] = uint16(11)
			m.cycles += 2
		}
	case 0x026b:
		{
			m.ip = 0x026c
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x026c:
		{
			m.ip = 0x026f
			m.r[1] = uint16(19)
			m.cycles += 2
		}
	case 0x026f:
		{
			m.ip = 0x0273
			m.wr16(m.r[11], uint16(m.r[7]+0x61c), m.r[0])
			m.cycles += 8
		}
	case 0x0273:
		{
			m.ip = 0x0276
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4))
			m.cycles += 4
		}
	case 0x0276:
		{
			m.ip = 0x0279
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x0279:
		{
			m.ip = 0x027b
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(623)
			}
			m.cycles += 17
		}
	case 0x027b:
		{
			m.ip = 0x027e
			m.r[0] = m.alu("add", 16, m.r[0], uint16(1844))
			m.cycles += 4
		}
	case 0x027e:
		{
			m.ip = 0x0281
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x0281:
		{
			m.ip = 0x0282
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x0282:
		{
			m.ip = 0x0284
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(619)
			}
			m.cycles += 17
		}
	case 0x0284:
		{
			m.ip = 0x0285
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0285:
		{
			m.ip = 0x028b
			m.wr16(m.r[11], uint16(0x466), uint16(0))
			m.cycles += 8
		}
	case 0x028b:
		{
			m.ip = 0x0291
			m.wr16(m.r[11], uint16(0x468), uint16(1))
			m.cycles += 8
		}
	case 0x0291:
		{
			m.ip = 0x0296
			m.wr8(m.r[11], uint16(0x46a), uint16(0))
			m.cycles += 8
		}
	case 0x0296:
		{
			m.ip = 0x029b
			m.wr8(m.r[11], uint16(0x2074), uint16(0))
			m.cycles += 8
		}
	case 0x029b:
		{
			m.ip = 0x02a0
			m.wr8(m.r[11], uint16(0x27a), uint16(0))
			m.cycles += 8
		}
	case 0x02a0:
		{
			m.ip = 0x02a3
			m.r[0] = m.rd16(m.r[11], uint16(0x27b))
			m.cycles += 8
		}
	case 0x02a3:
		{
			m.ip = 0x02a6
			m.wr16(m.r[11], uint16(0x466), m.r[0])
			m.cycles += 8
		}
	case 0x02a6:
		{
			m.ip = 0x02a9
			target := uint16(1805)
			m.push(0x02a9)
			m.ip = target
			m.cycles += 19
		}
	case 0x02a9:
		{
			m.ip = 0x02ac
			m.r[0] = m.rd16(m.r[11], uint16(0x28c))
			m.cycles += 8
		}
	case 0x02ac:
		{
			m.ip = 0x02af
			target := uint16(3035)
			m.push(0x02af)
			m.ip = target
			m.cycles += 19
		}
	case 0x02af:
		{
			m.ip = 0x02b4
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x468)), uint16(3))
			m.cycles += 16
		}
	case 0x02b4:
		{
			m.ip = 0x02b6
			if m.zf {
				m.ip = uint16(700)
			}
			m.cycles += 8
		}
	case 0x02b6:
		{
			m.ip = 0x02bc
			m.wr16(m.r[11], uint16(0x468), uint16(1))
			m.cycles += 8
		}
	case 0x02bc:
		{
			m.ip = 0x02bf
			target := uint16(2656)
			m.push(0x02bf)
			m.ip = target
			m.cycles += 19
		}
	case 0x02bf:
		{
			m.ip = 0x02c2
			target := uint16(2087)
			m.push(0x02c2)
			m.ip = target
			m.cycles += 19
		}
	case 0x02c2:
		{
			m.ip = 0x02c5
			target := uint16(8099)
			m.push(0x02c5)
			m.ip = target
			m.cycles += 19
		}
	case 0x02c5:
		{
			m.ip = 0x02c8
			target := uint16(2732)
			m.push(0x02c8)
			m.ip = target
			m.cycles += 19
		}
	case 0x02c8:
		{
			m.ip = 0x02cd
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x75)), uint16(1))
			m.cycles += 16
		}
	case 0x02cd:
		{
			m.ip = 0x02cf
			if !m.zf {
				m.ip = uint16(722)
			}
			m.cycles += 8
		}
	case 0x02cf:
		{
			m.ip = 0x02d2
			m.ip = uint16(869)
			m.cycles += 15
		}
	case 0x02d2:
		{
			m.ip = 0x02d5
			target := uint16(10769)
			m.push(0x02d5)
			m.ip = target
			m.cycles += 19
		}
	case 0x02d5:
		{
			m.ip = 0x02da
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x78)), uint16(0))
			m.cycles += 16
		}
	case 0x02da:
		{
			m.ip = 0x02dc
			if !m.zf {
				m.ip = uint16(709)
			}
			m.cycles += 8
		}
	case 0x02dc:
		{
			m.ip = 0x02df
			target := uint16(8099)
			m.push(0x02df)
			m.ip = target
			m.cycles += 19
		}
	case 0x02df:
		{
			m.ip = 0x02e4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x77)), uint16(1))
			m.cycles += 16
		}
	case 0x02e4:
		{
			m.ip = 0x02e6
			if m.zf {
				m.ip = uint16(869)
			}
			m.cycles += 8
		}
	case 0x02e6:
		{
			m.ip = 0x02eb
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x26b)), uint16(1))
			m.cycles += 16
		}
	case 0x02eb:
		{
			m.ip = 0x02ed
			if m.zf {
				m.ip = uint16(869)
			}
			m.cycles += 8
		}
	case 0x02ed:
		{
			m.ip = 0x02f2
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x246)), uint16(0))
			m.cycles += 16
		}
	case 0x02f2:
		{
			m.ip = 0x02f4
			if m.zf {
				m.ip = uint16(864)
			}
			m.cycles += 8
		}
	case 0x02f4:
		{
			m.ip = 0x02f9
			m.wr8(m.r[11], uint16(0x24e), m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24e)), uint16(16)))
			m.cycles += 16
		}
	case 0x02f9:
		{
			m.ip = 0x02fc
			m.r[0] = m.rd16(m.r[11], uint16(0x28c))
			m.cycles += 8
		}
	case 0x02fc:
		{
			m.ip = 0x02ff
			target := uint16(3061)
			m.push(0x02ff)
			m.ip = target
			m.cycles += 19
		}
	case 0x02ff:
		{
			m.ip = 0x0304
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(1))
			m.cycles += 16
		}
	case 0x0304:
		{
			m.ip = 0x0306
			if m.zf {
				m.ip = uint16(794)
			}
			m.cycles += 8
		}
	case 0x0306:
		{
			m.ip = 0x030b
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x468)), uint16(1))
			m.cycles += 16
		}
	case 0x030b:
		{
			m.ip = 0x030d
			if !m.zf {
				m.ip = uint16(794)
			}
			m.cycles += 8
		}
	case 0x030d:
		{
			m.ip = 0x0311
			m.wr16(m.r[11], uint16(0x468), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x468))))
			m.cycles += 15
		}
	case 0x0311:
		{
			m.ip = 0x0315
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x0315:
		{
			m.ip = 0x0318
			target := uint16(9400)
			m.push(0x0318)
			m.ip = target
			m.cycles += 19
		}
	case 0x0318:
		{
			m.ip = 0x031a
			m.ip = uint16(700)
			m.cycles += 15
		}
	case 0x031a:
		{
			m.ip = 0x031f
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(1))
			m.cycles += 16
		}
	case 0x031f:
		{
			m.ip = 0x0321
			if !m.zf {
				m.ip = uint16(811)
			}
			m.cycles += 8
		}
	case 0x0321:
		{
			m.ip = 0x0326
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x466)), uint16(2))
			m.cycles += 16
		}
	case 0x0326:
		{
			m.ip = 0x0328
			if !m.zf {
				m.ip = uint16(811)
			}
			m.cycles += 8
		}
	case 0x0328:
		{
			m.ip = 0x032b
			m.ip = uint16(5684)
			m.cycles += 15
		}
	case 0x032b:
		{
			m.ip = 0x0330
			m.wr16(m.r[11], uint16(0x466), m.alu("add", 16, m.rd16(m.r[11], uint16(0x466)), uint16(2)))
			m.cycles += 16
		}
	case 0x0330:
		{
			m.ip = 0x0335
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x466)), uint16(8))
			m.cycles += 16
		}
	case 0x0335:
		{
			m.ip = 0x0337
			if !m.zf {
				m.ip = uint16(835)
			}
			m.cycles += 8
		}
	case 0x0337:
		{
			m.ip = 0x033c
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x468)), uint16(3))
			m.cycles += 16
		}
	case 0x033c:
		{
			m.ip = 0x033e
			if !m.zf {
				m.ip = uint16(835)
			}
			m.cycles += 8
		}
	case 0x033e:
		{
			m.ip = 0x0343
			m.wr8(m.r[11], uint16(0x46a), uint16(1))
			m.cycles += 8
		}
	case 0x0343:
		{
			m.ip = 0x0348
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x466)), uint16(10))
			m.cycles += 16
		}
	case 0x0348:
		{
			m.ip = 0x034a
			if !m.zf {
				m.ip = uint16(861)
			}
			m.cycles += 8
		}
	case 0x034a:
		{
			m.ip = 0x034f
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x468)), uint16(3))
			m.cycles += 16
		}
	case 0x034f:
		{
			m.ip = 0x0351
			if m.zf {
				m.ip = uint16(864)
			}
			m.cycles += 8
		}
	case 0x0351:
		{
			m.ip = 0x0357
			m.wr16(m.r[11], uint16(0x466), uint16(0))
			m.cycles += 8
		}
	case 0x0357:
		{
			m.ip = 0x035d
			m.wr16(m.r[11], uint16(0x468), uint16(3))
			m.cycles += 8
		}
	case 0x035d:
		{
			m.ip = 0x0360
			m.ip = uint16(678)
			m.cycles += 15
		}
	case 0x0360:
		{
			m.ip = 0x0365
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x0365:
		{
			m.ip = 0x036a
			m.wr8(m.r[11], uint16(0x2074), uint16(1))
			m.cycles += 8
		}
	case 0x036a:
		{
			m.ip = 0x036b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x036b:
		{
			m.ip = 0x0370
			m.wr8(m.r[11], uint16(0x2074), uint16(0))
			m.cycles += 8
		}
	case 0x0370:
		{
			m.ip = 0x0375
			m.wr8(m.r[11], uint16(0x23b), uint16(1))
			m.cycles += 8
		}
	case 0x0375:
		{
			m.ip = 0x037b
			m.wr16(m.r[11], uint16(0x466), uint16(0))
			m.cycles += 8
		}
	case 0x037b:
		{
			m.ip = 0x037e
			target := uint16(1805)
			m.push(0x037e)
			m.ip = target
			m.cycles += 19
		}
	case 0x037e:
		{
			m.ip = 0x0381
			target := uint16(2656)
			m.push(0x0381)
			m.ip = target
			m.cycles += 19
		}
	case 0x0381:
		{
			m.ip = 0x0384
			m.r[0] = m.rd16(m.r[11], uint16(0x292))
			m.cycles += 8
		}
	case 0x0384:
		{
			m.ip = 0x0387
			target := uint16(3035)
			m.push(0x0387)
			m.ip = target
			m.cycles += 19
		}
	case 0x0387:
		{
			m.ip = 0x038a
			m.r[0] = m.rd16(m.r[11], uint16(0x28e))
			m.cycles += 8
		}
	case 0x038a:
		{
			m.ip = 0x038d
			target := uint16(3035)
			m.push(0x038d)
			m.ip = target
			m.cycles += 19
		}
	case 0x038d:
		{
			m.ip = 0x038f
			m.ip = uint16(915)
			m.cycles += 15
		}
	case 0x038f:
		{
			m.ip = 0x0390
			m.cycles += 3
		}
	case 0x0390:
		{
			m.ip = 0x0393
			target := uint16(1805)
			m.push(0x0393)
			m.ip = target
			m.cycles += 19
		}
	case 0x0393:
		{
			m.ip = 0x0396
			m.r[0] = m.rd16(m.r[11], uint16(0x28c))
			m.cycles += 8
		}
	case 0x0396:
		{
			m.ip = 0x0399
			target := uint16(3035)
			m.push(0x0399)
			m.ip = target
			m.cycles += 19
		}
	case 0x0399:
		{
			m.ip = 0x039e
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x23e)), uint16(3))
			m.cycles += 16
		}
	case 0x039e:
		{
			m.ip = 0x03a0
			if m.zf {
				m.ip = uint16(934)
			}
			m.cycles += 8
		}
	case 0x03a0:
		{
			m.ip = 0x03a6
			m.wr16(m.r[11], uint16(0x23e), uint16(1))
			m.cycles += 8
		}
	case 0x03a6:
		{
			m.ip = 0x03a9
			target := uint16(2656)
			m.push(0x03a9)
			m.ip = target
			m.cycles += 19
		}
	case 0x03a9:
		{
			m.ip = 0x03ac
			target := uint16(2087)
			m.push(0x03ac)
			m.ip = target
			m.cycles += 19
		}
	case 0x03ac:
		{
			m.ip = 0x03af
			target := uint16(8099)
			m.push(0x03af)
			m.ip = target
			m.cycles += 19
		}
	case 0x03af:
		{
			m.ip = 0x03b2
			m.r[0] = m.rd16(m.r[11], uint16(0x23e))
			m.cycles += 8
		}
	case 0x03b2:
		{
			m.ip = 0x03b5
			m.wr16(m.r[11], uint16(0x468), m.r[0])
			m.cycles += 8
		}
	case 0x03b5:
		{
			m.ip = 0x03b8
			target := uint16(2732)
			m.push(0x03b8)
			m.ip = target
			m.cycles += 19
		}
	case 0x03b8:
		{
			m.ip = 0x03bd
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x75)), uint16(1))
			m.cycles += 16
		}
	case 0x03bd:
		{
			m.ip = 0x03bf
			if !m.zf {
				m.ip = uint16(962)
			}
			m.cycles += 8
		}
	case 0x03bf:
		{
			m.ip = 0x03c2
			m.ip = uint16(1332)
			m.cycles += 15
		}
	case 0x03c2:
		{
			m.ip = 0x03c7
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x76)), uint16(1))
			m.cycles += 16
		}
	case 0x03c7:
		{
			m.ip = 0x03c9
			if !m.zf {
				m.ip = uint16(975)
			}
			m.cycles += 8
		}
	case 0x03c9:
		{
			m.ip = 0x03cc
			target := uint16(1333)
			m.push(0x03cc)
			m.ip = target
			m.cycles += 19
		}
	case 0x03cc:
		{
			m.ip = 0x03cf
			m.ip = uint16(1140)
			m.cycles += 15
		}
	case 0x03cf:
		{
			m.ip = 0x03d2
			target := uint16(10769)
			m.push(0x03d2)
			m.ip = target
			m.cycles += 19
		}
	case 0x03d2:
		{
			m.ip = 0x03d7
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x78)), uint16(0))
			m.cycles += 16
		}
	case 0x03d7:
		{
			m.ip = 0x03d9
			if !m.zf {
				m.ip = uint16(949)
			}
			m.cycles += 8
		}
	case 0x03d9:
		{
			m.ip = 0x03dc
			target := uint16(8099)
			m.push(0x03dc)
			m.ip = target
			m.cycles += 19
		}
	case 0x03dc:
		{
			m.ip = 0x03e1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x77)), uint16(1))
			m.cycles += 16
		}
	case 0x03e1:
		{
			m.ip = 0x03e3
			if !m.zf {
				m.ip = uint16(998)
			}
			m.cycles += 8
		}
	case 0x03e3:
		{
			m.ip = 0x03e6
			m.ip = uint16(1332)
			m.cycles += 15
		}
	case 0x03e6:
		{
			m.ip = 0x03e9
			m.r[0] = m.rd16(m.r[11], uint16(0x28c))
			m.cycles += 8
		}
	case 0x03e9:
		{
			m.ip = 0x03ec
			target := uint16(3061)
			m.push(0x03ec)
			m.ip = target
			m.cycles += 19
		}
	case 0x03ec:
		{
			m.ip = 0x03f1
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x23e)), uint16(1))
			m.cycles += 16
		}
	case 0x03f1:
		{
			m.ip = 0x03f3
			if !m.zf {
				m.ip = uint16(1024)
			}
			m.cycles += 8
		}
	case 0x03f3:
		{
			m.ip = 0x03f7
			m.wr16(m.r[11], uint16(0x23e), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x23e))))
			m.cycles += 15
		}
	case 0x03f7:
		{
			m.ip = 0x03fb
			m.r[3] = m.rd16(m.r[11], uint16(0x23c))
			m.cycles += 8
		}
	case 0x03fb:
		{
			m.ip = 0x03fe
			target := uint16(9400)
			m.push(0x03fe)
			m.ip = target
			m.cycles += 19
		}
	case 0x03fe:
		{
			m.ip = 0x0400
			m.ip = uint16(934)
			m.cycles += 15
		}
	case 0x0400:
		{
			m.ip = 0x0405
			m.wr16(m.r[11], uint16(0x23c), m.alu("add", 16, m.rd16(m.r[11], uint16(0x23c)), uint16(2)))
			m.cycles += 16
		}
	case 0x0405:
		{
			m.ip = 0x040a
			m.wr16(m.r[11], uint16(0x466), m.alu("add", 16, m.rd16(m.r[11], uint16(0x466)), uint16(2)))
			m.cycles += 16
		}
	case 0x040a:
		{
			m.ip = 0x040f
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x23c)), uint16(8))
			m.cycles += 16
		}
	case 0x040f:
		{
			m.ip = 0x0411
			if !m.zf {
				m.ip = uint16(1058)
			}
			m.cycles += 8
		}
	case 0x0411:
		{
			m.ip = 0x0416
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x23e)), uint16(3))
			m.cycles += 16
		}
	case 0x0416:
		{
			m.ip = 0x0418
			if !m.zf {
				m.ip = uint16(1058)
			}
			m.cycles += 8
		}
	case 0x0418:
		{
			m.ip = 0x041d
			m.wr8(m.r[11], uint16(0x244), uint16(1))
			m.cycles += 8
		}
	case 0x041d:
		{
			m.ip = 0x0422
			m.wr8(m.r[11], uint16(0x46a), uint16(1))
			m.cycles += 8
		}
	case 0x0422:
		{
			m.ip = 0x0427
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x23c)), uint16(10))
			m.cycles += 16
		}
	case 0x0427:
		{
			m.ip = 0x0429
			if !m.zf {
				m.ip = uint16(1096)
			}
			m.cycles += 8
		}
	case 0x0429:
		{
			m.ip = 0x042e
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x23e)), uint16(3))
			m.cycles += 16
		}
	case 0x042e:
		{
			m.ip = 0x0430
			if m.zf {
				m.ip = uint16(1099)
			}
			m.cycles += 8
		}
	case 0x0430:
		{
			m.ip = 0x0436
			m.wr16(m.r[11], uint16(0x23c), uint16(0))
			m.cycles += 8
		}
	case 0x0436:
		{
			m.ip = 0x043c
			m.wr16(m.r[11], uint16(0x466), uint16(0))
			m.cycles += 8
		}
	case 0x043c:
		{
			m.ip = 0x0442
			m.wr16(m.r[11], uint16(0x23e), uint16(3))
			m.cycles += 8
		}
	case 0x0442:
		{
			m.ip = 0x0448
			m.wr16(m.r[11], uint16(0x468), uint16(3))
			m.cycles += 8
		}
	case 0x0448:
		{
			m.ip = 0x044b
			m.ip = uint16(912)
			m.cycles += 15
		}
	case 0x044b:
		{
			m.ip = 0x044e
			target := uint16(1302)
			m.push(0x044e)
			m.ip = target
			m.cycles += 19
		}
	case 0x044e:
		{
			m.ip = 0x044f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x044f:
		{
			m.ip = 0x0452
			target := uint16(1805)
			m.push(0x0452)
			m.ip = target
			m.cycles += 19
		}
	case 0x0452:
		{
			m.ip = 0x0455
			m.r[0] = m.rd16(m.r[11], uint16(0x28e))
			m.cycles += 8
		}
	case 0x0455:
		{
			m.ip = 0x0458
			target := uint16(3035)
			m.push(0x0458)
			m.ip = target
			m.cycles += 19
		}
	case 0x0458:
		{
			m.ip = 0x045d
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x242)), uint16(3))
			m.cycles += 16
		}
	case 0x045d:
		{
			m.ip = 0x045f
			if m.zf {
				m.ip = uint16(1125)
			}
			m.cycles += 8
		}
	case 0x045f:
		{
			m.ip = 0x0465
			m.wr16(m.r[11], uint16(0x242), uint16(1))
			m.cycles += 8
		}
	case 0x0465:
		{
			m.ip = 0x0468
			target := uint16(2656)
			m.push(0x0468)
			m.ip = target
			m.cycles += 19
		}
	case 0x0468:
		{
			m.ip = 0x046b
			target := uint16(2087)
			m.push(0x046b)
			m.ip = target
			m.cycles += 19
		}
	case 0x046b:
		{
			m.ip = 0x046e
			target := uint16(8099)
			m.push(0x046e)
			m.ip = target
			m.cycles += 19
		}
	case 0x046e:
		{
			m.ip = 0x0471
			m.r[0] = m.rd16(m.r[11], uint16(0x242))
			m.cycles += 8
		}
	case 0x0471:
		{
			m.ip = 0x0474
			m.wr16(m.r[11], uint16(0x468), m.r[0])
			m.cycles += 8
		}
	case 0x0474:
		{
			m.ip = 0x0477
			target := uint16(2732)
			m.push(0x0477)
			m.ip = target
			m.cycles += 19
		}
	case 0x0477:
		{
			m.ip = 0x047c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x75)), uint16(1))
			m.cycles += 16
		}
	case 0x047c:
		{
			m.ip = 0x047e
			if !m.zf {
				m.ip = uint16(1153)
			}
			m.cycles += 8
		}
	case 0x047e:
		{
			m.ip = 0x0481
			m.ip = uint16(1332)
			m.cycles += 15
		}
	case 0x0481:
		{
			m.ip = 0x0486
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x76)), uint16(1))
			m.cycles += 16
		}
	case 0x0486:
		{
			m.ip = 0x0488
			if !m.zf {
				m.ip = uint16(1174)
			}
			m.cycles += 8
		}
	case 0x0488:
		{
			m.ip = 0x048b
			target := uint16(1435)
			m.push(0x048b)
			m.ip = target
			m.cycles += 19
		}
	case 0x048b:
		{
			m.ip = 0x0490
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2074)), uint16(1))
			m.cycles += 16
		}
	case 0x0490:
		{
			m.ip = 0x0492
			if m.zf {
				m.ip = uint16(1173)
			}
			m.cycles += 8
		}
	case 0x0492:
		{
			m.ip = 0x0495
			m.ip = uint16(949)
			m.cycles += 15
		}
	case 0x0495:
		{
			m.ip = 0x0496
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0496:
		{
			m.ip = 0x0499
			target := uint16(10769)
			m.push(0x0499)
			m.ip = target
			m.cycles += 19
		}
	case 0x0499:
		{
			m.ip = 0x049e
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x78)), uint16(0))
			m.cycles += 16
		}
	case 0x049e:
		{
			m.ip = 0x04a0
			if !m.zf {
				m.ip = uint16(1140)
			}
			m.cycles += 8
		}
	case 0x04a0:
		{
			m.ip = 0x04a3
			target := uint16(8099)
			m.push(0x04a3)
			m.ip = target
			m.cycles += 19
		}
	case 0x04a3:
		{
			m.ip = 0x04a8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x77)), uint16(1))
			m.cycles += 16
		}
	case 0x04a8:
		{
			m.ip = 0x04aa
			if !m.zf {
				m.ip = uint16(1197)
			}
			m.cycles += 8
		}
	case 0x04aa:
		{
			m.ip = 0x04ad
			m.ip = uint16(1332)
			m.cycles += 15
		}
	case 0x04ad:
		{
			m.ip = 0x04b0
			m.r[0] = m.rd16(m.r[11], uint16(0x28e))
			m.cycles += 8
		}
	case 0x04b0:
		{
			m.ip = 0x04b3
			target := uint16(3061)
			m.push(0x04b3)
			m.ip = target
			m.cycles += 19
		}
	case 0x04b3:
		{
			m.ip = 0x04b8
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x242)), uint16(1))
			m.cycles += 16
		}
	case 0x04b8:
		{
			m.ip = 0x04ba
			if !m.zf {
				m.ip = uint16(1223)
			}
			m.cycles += 8
		}
	case 0x04ba:
		{
			m.ip = 0x04be
			m.wr16(m.r[11], uint16(0x242), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x242))))
			m.cycles += 15
		}
	case 0x04be:
		{
			m.ip = 0x04c2
			m.r[3] = m.rd16(m.r[11], uint16(0x240))
			m.cycles += 8
		}
	case 0x04c2:
		{
			m.ip = 0x04c5
			target := uint16(9400)
			m.push(0x04c5)
			m.ip = target
			m.cycles += 19
		}
	case 0x04c5:
		{
			m.ip = 0x04c7
			m.ip = uint16(1125)
			m.cycles += 15
		}
	case 0x04c7:
		{
			m.ip = 0x04cc
			m.wr16(m.r[11], uint16(0x240), m.alu("add", 16, m.rd16(m.r[11], uint16(0x240)), uint16(2)))
			m.cycles += 16
		}
	case 0x04cc:
		{
			m.ip = 0x04d1
			m.wr16(m.r[11], uint16(0x466), m.alu("add", 16, m.rd16(m.r[11], uint16(0x466)), uint16(2)))
			m.cycles += 16
		}
	case 0x04d1:
		{
			m.ip = 0x04d6
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x240)), uint16(8))
			m.cycles += 16
		}
	case 0x04d6:
		{
			m.ip = 0x04d8
			if !m.zf {
				m.ip = uint16(1257)
			}
			m.cycles += 8
		}
	case 0x04d8:
		{
			m.ip = 0x04dd
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x242)), uint16(3))
			m.cycles += 16
		}
	case 0x04dd:
		{
			m.ip = 0x04df
			if !m.zf {
				m.ip = uint16(1257)
			}
			m.cycles += 8
		}
	case 0x04df:
		{
			m.ip = 0x04e4
			m.wr8(m.r[11], uint16(0x245), uint16(1))
			m.cycles += 8
		}
	case 0x04e4:
		{
			m.ip = 0x04e9
			m.wr8(m.r[11], uint16(0x46a), uint16(1))
			m.cycles += 8
		}
	case 0x04e9:
		{
			m.ip = 0x04ee
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x240)), uint16(10))
			m.cycles += 16
		}
	case 0x04ee:
		{
			m.ip = 0x04f0
			if !m.zf {
				m.ip = uint16(1295)
			}
			m.cycles += 8
		}
	case 0x04f0:
		{
			m.ip = 0x04f5
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x242)), uint16(3))
			m.cycles += 16
		}
	case 0x04f5:
		{
			m.ip = 0x04f7
			if m.zf {
				m.ip = uint16(1298)
			}
			m.cycles += 8
		}
	case 0x04f7:
		{
			m.ip = 0x04fd
			m.wr16(m.r[11], uint16(0x240), uint16(0))
			m.cycles += 8
		}
	case 0x04fd:
		{
			m.ip = 0x0503
			m.wr16(m.r[11], uint16(0x466), uint16(0))
			m.cycles += 8
		}
	case 0x0503:
		{
			m.ip = 0x0509
			m.wr16(m.r[11], uint16(0x242), uint16(3))
			m.cycles += 8
		}
	case 0x0509:
		{
			m.ip = 0x050f
			m.wr16(m.r[11], uint16(0x468), uint16(3))
			m.cycles += 8
		}
	case 0x050f:
		{
			m.ip = 0x0512
			m.ip = uint16(1103)
			m.cycles += 15
		}
	case 0x0512:
		{
			m.ip = 0x0514
			m.ip = uint16(1319)
			m.cycles += 15
		}
	case 0x0514:
		{
			m.ip = 0x0515
			m.cycles += 3
		}
	case 0x0515:
		{
			m.ip = 0x0516
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0516:
		{
			m.ip = 0x0519
			target := uint16(6912)
			m.push(0x0519)
			m.ip = target
			m.cycles += 19
		}
	case 0x0519:
		{
			m.ip = 0x051e
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x051e:
		{
			m.ip = 0x0523
			m.wr8(m.r[11], uint16(0x2078), uint16(1))
			m.cycles += 8
		}
	case 0x0523:
		{
			m.ip = 0x0526
			target := uint16(3091)
			m.push(0x0526)
			m.ip = target
			m.cycles += 19
		}
	case 0x0526:
		{
			m.ip = 0x0527
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0527:
		{
			m.ip = 0x052a
			target := uint16(6912)
			m.push(0x052a)
			m.ip = target
			m.cycles += 19
		}
	case 0x052a:
		{
			m.ip = 0x052f
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x052f:
		{
			m.ip = 0x0534
			m.wr8(m.r[11], uint16(0x2074), uint16(1))
			m.cycles += 8
		}
	case 0x0534:
		{
			m.ip = 0x0535
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0535:
		{
			m.ip = 0x053a
			m.wr8(m.r[11], uint16(0x76), uint16(0))
			m.cycles += 8
		}
	case 0x053a:
		{
			m.ip = 0x053d
			target := uint16(8099)
			m.push(0x053d)
			m.ip = target
			m.cycles += 19
		}
	case 0x053d:
		{
			m.ip = 0x0540
			m.r[0] = m.rd16(m.r[11], uint16(0x290))
			m.cycles += 8
		}
	case 0x0540:
		{
			m.ip = 0x0543
			target := uint16(3035)
			m.push(0x0543)
			m.ip = target
			m.cycles += 19
		}
	case 0x0543:
		{
			m.ip = 0x0548
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24a)), uint16(1))
			m.cycles += 16
		}
	case 0x0548:
		{
			m.ip = 0x054a
			if !m.zf {
				m.ip = uint16(1357)
			}
			m.cycles += 8
		}
	case 0x054a:
		{
			m.ip = 0x054d
			target := uint16(1302)
			m.push(0x054d)
			m.ip = target
			m.cycles += 19
		}
	case 0x054d:
		{
			m.ip = 0x0551
			m.wr16(m.r[11], uint16(0x24a), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x24a))))
			m.cycles += 15
		}
	case 0x0551:
		{
			m.ip = 0x0554
			m.r[0] = m.rd16(m.r[11], uint16(0x24c))
			m.cycles += 8
		}
	case 0x0554:
		{
			m.ip = 0x0557
			m.wr16(m.r[11], uint16(0x246), m.r[0])
			m.cycles += 8
		}
	case 0x0557:
		{
			m.ip = 0x055a
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x245)))
			m.cycles += 8
		}
	case 0x055a:
		{
			m.ip = 0x055d
			m.wr16(m.r[11], uint16(0x46a), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x055d:
		{
			m.ip = 0x0562
			m.wr8(m.r[11], uint16(0x24f), m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f)), uint16(16)))
			m.cycles += 16
		}
	case 0x0562:
		{
			m.ip = 0x0565
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x250)))
			m.cycles += 8
		}
	case 0x0565:
		{
			m.ip = 0x0568
			m.wr16(m.r[11], uint16(0x24e), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0568:
		{
			m.ip = 0x056b
			m.r[0] = m.rd16(m.r[11], uint16(0x233))
			m.cycles += 8
		}
	case 0x056b:
		{
			m.ip = 0x056e
			m.wr16(m.r[11], uint16(0x237), m.r[0])
			m.cycles += 8
		}
	case 0x056e:
		{
			m.ip = 0x0571
			m.r[0] = m.rd16(m.r[11], uint16(0x239))
			m.cycles += 8
		}
	case 0x0571:
		{
			m.ip = 0x0574
			m.wr16(m.r[11], uint16(0x233), m.r[0])
			m.cycles += 8
		}
	case 0x0574:
		{
			m.ip = 0x0579
			m.wr8(m.r[11], uint16(0x23b), uint16(2))
			m.cycles += 8
		}
	case 0x0579:
		{
			m.ip = 0x057c
			target := uint16(8099)
			m.push(0x057c)
			m.ip = target
			m.cycles += 19
		}
	case 0x057c:
		{
			m.ip = 0x057f
			m.r[0] = m.rd16(m.r[11], uint16(0x240))
			m.cycles += 8
		}
	case 0x057f:
		{
			m.ip = 0x0582
			m.wr16(m.r[11], uint16(0x466), m.r[0])
			m.cycles += 8
		}
	case 0x0582:
		{
			m.ip = 0x0586
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(0x23c)))
			m.cycles += 16
		}
	case 0x0586:
		{
			m.ip = 0x0588
			if m.zf {
				m.ip = uint16(1419)
			}
			m.cycles += 8
		}
	case 0x0588:
		{
			m.ip = 0x058b
			target := uint16(1805)
			m.push(0x058b)
			m.ip = target
			m.cycles += 19
		}
	case 0x058b:
		{
			m.ip = 0x058e
			target := uint16(1919)
			m.push(0x058e)
			m.ip = target
			m.cycles += 19
		}
	case 0x058e:
		{
			m.ip = 0x0591
			m.r[0] = m.rd16(m.r[11], uint16(0x292))
			m.cycles += 8
		}
	case 0x0591:
		{
			m.ip = 0x0594
			target := uint16(3061)
			m.push(0x0594)
			m.ip = target
			m.cycles += 19
		}
	case 0x0594:
		{
			m.ip = 0x0597
			target := uint16(9487)
			m.push(0x0597)
			m.ip = target
			m.cycles += 19
		}
	case 0x0597:
		{
			m.ip = 0x059a
			target := uint16(2087)
			m.push(0x059a)
			m.ip = target
			m.cycles += 19
		}
	case 0x059a:
		{
			m.ip = 0x059b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x059b:
		{
			m.ip = 0x05a0
			m.wr8(m.r[11], uint16(0x76), uint16(0))
			m.cycles += 8
		}
	case 0x05a0:
		{
			m.ip = 0x05a3
			target := uint16(8099)
			m.push(0x05a3)
			m.ip = target
			m.cycles += 19
		}
	case 0x05a3:
		{
			m.ip = 0x05a6
			m.r[0] = m.rd16(m.r[11], uint16(0x292))
			m.cycles += 8
		}
	case 0x05a6:
		{
			m.ip = 0x05a9
			target := uint16(3035)
			m.push(0x05a9)
			m.ip = target
			m.cycles += 19
		}
	case 0x05a9:
		{
			m.ip = 0x05ae
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24c)), uint16(1))
			m.cycles += 16
		}
	case 0x05ae:
		{
			m.ip = 0x05b0
			if !m.zf {
				m.ip = uint16(1459)
			}
			m.cycles += 8
		}
	case 0x05b0:
		{
			m.ip = 0x05b3
			m.ip = uint16(1319)
			m.cycles += 15
		}
	case 0x05b3:
		{
			m.ip = 0x05b7
			m.wr16(m.r[11], uint16(0x24c), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x24c))))
			m.cycles += 15
		}
	case 0x05b7:
		{
			m.ip = 0x05ba
			m.r[0] = m.rd16(m.r[11], uint16(0x24a))
			m.cycles += 8
		}
	case 0x05ba:
		{
			m.ip = 0x05bd
			m.wr16(m.r[11], uint16(0x246), m.r[0])
			m.cycles += 8
		}
	case 0x05bd:
		{
			m.ip = 0x05c0
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x244)))
			m.cycles += 8
		}
	case 0x05c0:
		{
			m.ip = 0x05c3
			m.wr16(m.r[11], uint16(0x46a), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x05c3:
		{
			m.ip = 0x05c8
			m.wr8(m.r[11], uint16(0x250), m.alu("sub", 8, m.rd8(m.r[11], uint16(0x250)), uint16(16)))
			m.cycles += 16
		}
	case 0x05c8:
		{
			m.ip = 0x05cb
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x24f)))
			m.cycles += 8
		}
	case 0x05cb:
		{
			m.ip = 0x05ce
			m.wr16(m.r[11], uint16(0x24e), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x05ce:
		{
			m.ip = 0x05d1
			m.r[0] = m.rd16(m.r[11], uint16(0x233))
			m.cycles += 8
		}
	case 0x05d1:
		{
			m.ip = 0x05d4
			m.wr16(m.r[11], uint16(0x239), m.r[0])
			m.cycles += 8
		}
	case 0x05d4:
		{
			m.ip = 0x05d7
			m.r[0] = m.rd16(m.r[11], uint16(0x237))
			m.cycles += 8
		}
	case 0x05d7:
		{
			m.ip = 0x05da
			m.wr16(m.r[11], uint16(0x233), m.r[0])
			m.cycles += 8
		}
	case 0x05da:
		{
			m.ip = 0x05df
			m.wr8(m.r[11], uint16(0x23b), uint16(1))
			m.cycles += 8
		}
	case 0x05df:
		{
			m.ip = 0x05e2
			target := uint16(8099)
			m.push(0x05e2)
			m.ip = target
			m.cycles += 19
		}
	case 0x05e2:
		{
			m.ip = 0x05e5
			m.r[0] = m.rd16(m.r[11], uint16(0x23c))
			m.cycles += 8
		}
	case 0x05e5:
		{
			m.ip = 0x05e8
			m.wr16(m.r[11], uint16(0x466), m.r[0])
			m.cycles += 8
		}
	case 0x05e8:
		{
			m.ip = 0x05ec
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(0x240)))
			m.cycles += 16
		}
	case 0x05ec:
		{
			m.ip = 0x05ee
			if m.zf {
				m.ip = uint16(1521)
			}
			m.cycles += 8
		}
	case 0x05ee:
		{
			m.ip = 0x05f1
			target := uint16(1805)
			m.push(0x05f1)
			m.ip = target
			m.cycles += 19
		}
	case 0x05f1:
		{
			m.ip = 0x05f4
			target := uint16(1919)
			m.push(0x05f4)
			m.ip = target
			m.cycles += 19
		}
	case 0x05f4:
		{
			m.ip = 0x05f7
			m.r[0] = m.rd16(m.r[11], uint16(0x290))
			m.cycles += 8
		}
	case 0x05f7:
		{
			m.ip = 0x05fa
			target := uint16(3061)
			m.push(0x05fa)
			m.ip = target
			m.cycles += 19
		}
	case 0x05fa:
		{
			m.ip = 0x05fd
			target := uint16(9487)
			m.push(0x05fd)
			m.ip = target
			m.cycles += 19
		}
	case 0x05fd:
		{
			m.ip = 0x0600
			target := uint16(2087)
			m.push(0x0600)
			m.ip = target
			m.cycles += 19
		}
	case 0x0600:
		{
			m.ip = 0x0601
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0601:
		{
			m.ip = 0x0607
			m.wr16(m.r[11], uint16(0x466), uint16(0))
			m.cycles += 8
		}
	case 0x0607:
		{
			m.ip = 0x060d
			m.wr16(m.r[11], uint16(0x468), uint16(1))
			m.cycles += 8
		}
	case 0x060d:
		{
			m.ip = 0x0612
			m.wr8(m.r[11], uint16(0x46a), uint16(0))
			m.cycles += 8
		}
	case 0x0612:
		{
			m.ip = 0x0617
			m.wr8(m.r[11], uint16(0x2074), uint16(0))
			m.cycles += 8
		}
	case 0x0617:
		{
			m.ip = 0x061c
			m.wr8(m.r[11], uint16(0x61a), uint16(2))
			m.cycles += 8
		}
	case 0x061c:
		{
			m.ip = 0x0621
			m.wr8(m.r[11], uint16(0x3f), uint16(0))
			m.cycles += 8
		}
	case 0x0621:
		{
			m.ip = 0x0626
			m.wr8(m.r[11], uint16(0x27a), uint16(0))
			m.cycles += 8
		}
	case 0x0626:
		{
			m.ip = 0x0629
			m.r[0] = m.rd16(m.r[11], uint16(0x27b))
			m.cycles += 8
		}
	case 0x0629:
		{
			m.ip = 0x062c
			m.wr16(m.r[11], uint16(0x466), m.r[0])
			m.cycles += 8
		}
	case 0x062c:
		{
			m.ip = 0x062f
			target := uint16(1805)
			m.push(0x062f)
			m.ip = target
			m.cycles += 19
		}
	case 0x062f:
		{
			m.ip = 0x0632
			m.r[0] = m.rd16(m.r[11], uint16(0x28c))
			m.cycles += 8
		}
	case 0x0632:
		{
			m.ip = 0x0635
			target := uint16(3035)
			m.push(0x0635)
			m.ip = target
			m.cycles += 19
		}
	case 0x0635:
		{
			m.ip = 0x063a
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x468)), uint16(3))
			m.cycles += 16
		}
	case 0x063a:
		{
			m.ip = 0x063c
			if m.zf {
				m.ip = uint16(1602)
			}
			m.cycles += 8
		}
	case 0x063c:
		{
			m.ip = 0x0642
			m.wr16(m.r[11], uint16(0x468), uint16(1))
			m.cycles += 8
		}
	case 0x0642:
		{
			m.ip = 0x0645
			target := uint16(2656)
			m.push(0x0645)
			m.ip = target
			m.cycles += 19
		}
	case 0x0645:
		{
			m.ip = 0x0648
			target := uint16(2087)
			m.push(0x0648)
			m.ip = target
			m.cycles += 19
		}
	case 0x0648:
		{
			m.ip = 0x064b
			target := uint16(8099)
			m.push(0x064b)
			m.ip = target
			m.cycles += 19
		}
	case 0x064b:
		{
			m.ip = 0x064e
			target := uint16(2732)
			m.push(0x064e)
			m.ip = target
			m.cycles += 19
		}
	case 0x064e:
		{
			m.ip = 0x0653
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x75)), uint16(1))
			m.cycles += 16
		}
	case 0x0653:
		{
			m.ip = 0x0655
			if !m.zf {
				m.ip = uint16(1624)
			}
			m.cycles += 8
		}
	case 0x0655:
		{
			m.ip = 0x0658
			m.ip = uint16(1799)
			m.cycles += 15
		}
	case 0x0658:
		{
			m.ip = 0x065b
			target := uint16(10769)
			m.push(0x065b)
			m.ip = target
			m.cycles += 19
		}
	case 0x065b:
		{
			m.ip = 0x065e
			m.r[0] = m.rd16(m.r[11], uint16(0x246))
			m.cycles += 8
		}
	case 0x065e:
		{
			m.ip = 0x0662
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(0x248)))
			m.cycles += 16
		}
	case 0x0662:
		{
			m.ip = 0x0665
			m.alu("sub", 16, m.r[0], uint16(0))
			m.cycles += 4
		}
	case 0x0665:
		{
			m.ip = 0x0667
			if !m.zf {
				m.ip = uint16(1642)
			}
			m.cycles += 8
		}
	case 0x0667:
		{
			m.ip = 0x0669
			m.ip = uint16(1753)
			m.cycles += 15
		}
	case 0x0669:
		{
			m.ip = 0x066a
			m.cycles += 3
		}
	case 0x066a:
		{
			m.ip = 0x066f
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x78)), uint16(0))
			m.cycles += 16
		}
	case 0x066f:
		{
			m.ip = 0x0671
			if !m.zf {
				m.ip = uint16(1611)
			}
			m.cycles += 8
		}
	case 0x0671:
		{
			m.ip = 0x0674
			target := uint16(8099)
			m.push(0x0674)
			m.ip = target
			m.cycles += 19
		}
	case 0x0674:
		{
			m.ip = 0x0679
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x77)), uint16(1))
			m.cycles += 16
		}
	case 0x0679:
		{
			m.ip = 0x067b
			if !m.zf {
				m.ip = uint16(1662)
			}
			m.cycles += 8
		}
	case 0x067b:
		{
			m.ip = 0x067e
			m.ip = uint16(1799)
			m.cycles += 15
		}
	case 0x067e:
		{
			m.ip = 0x0683
			m.wr8(m.r[11], uint16(0x24e), m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24e)), uint16(16)))
			m.cycles += 16
		}
	case 0x0683:
		{
			m.ip = 0x0686
			m.r[0] = m.rd16(m.r[11], uint16(0x28c))
			m.cycles += 8
		}
	case 0x0686:
		{
			m.ip = 0x0689
			target := uint16(3061)
			m.push(0x0689)
			m.ip = target
			m.cycles += 19
		}
	case 0x0689:
		{
			m.ip = 0x068e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(1))
			m.cycles += 16
		}
	case 0x068e:
		{
			m.ip = 0x0690
			if m.zf {
				m.ip = uint16(1700)
			}
			m.cycles += 8
		}
	case 0x0690:
		{
			m.ip = 0x0695
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x468)), uint16(1))
			m.cycles += 16
		}
	case 0x0695:
		{
			m.ip = 0x0697
			if !m.zf {
				m.ip = uint16(1700)
			}
			m.cycles += 8
		}
	case 0x0697:
		{
			m.ip = 0x069b
			m.wr16(m.r[11], uint16(0x468), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x468))))
			m.cycles += 15
		}
	case 0x069b:
		{
			m.ip = 0x069f
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x069f:
		{
			m.ip = 0x06a2
			target := uint16(9400)
			m.push(0x06a2)
			m.ip = target
			m.cycles += 19
		}
	case 0x06a2:
		{
			m.ip = 0x06a4
			m.ip = uint16(1602)
			m.cycles += 15
		}
	case 0x06a4:
		{
			m.ip = 0x06a9
			m.wr16(m.r[11], uint16(0x466), m.alu("add", 16, m.rd16(m.r[11], uint16(0x466)), uint16(2)))
			m.cycles += 16
		}
	case 0x06a9:
		{
			m.ip = 0x06ae
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x466)), uint16(8))
			m.cycles += 16
		}
	case 0x06ae:
		{
			m.ip = 0x06b0
			if !m.zf {
				m.ip = uint16(1724)
			}
			m.cycles += 8
		}
	case 0x06b0:
		{
			m.ip = 0x06b5
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x468)), uint16(3))
			m.cycles += 16
		}
	case 0x06b5:
		{
			m.ip = 0x06b7
			if !m.zf {
				m.ip = uint16(1724)
			}
			m.cycles += 8
		}
	case 0x06b7:
		{
			m.ip = 0x06bc
			m.wr8(m.r[11], uint16(0x46a), uint16(1))
			m.cycles += 8
		}
	case 0x06bc:
		{
			m.ip = 0x06c1
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x466)), uint16(10))
			m.cycles += 16
		}
	case 0x06c1:
		{
			m.ip = 0x06c3
			if !m.zf {
				m.ip = uint16(1750)
			}
			m.cycles += 8
		}
	case 0x06c3:
		{
			m.ip = 0x06c8
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x468)), uint16(3))
			m.cycles += 16
		}
	case 0x06c8:
		{
			m.ip = 0x06ca
			if m.zf {
				m.ip = uint16(1753)
			}
			m.cycles += 8
		}
	case 0x06ca:
		{
			m.ip = 0x06d0
			m.wr16(m.r[11], uint16(0x466), uint16(0))
			m.cycles += 8
		}
	case 0x06d0:
		{
			m.ip = 0x06d6
			m.wr16(m.r[11], uint16(0x468), uint16(3))
			m.cycles += 8
		}
	case 0x06d6:
		{
			m.ip = 0x06d9
			m.ip = uint16(1580)
			m.cycles += 15
		}
	case 0x06d9:
		{
			m.ip = 0x06de
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x06de:
		{
			m.ip = 0x06e3
			m.wr8(m.r[11], uint16(0x1984), uint16(0))
			m.cycles += 8
		}
	case 0x06e3:
		{
			m.ip = 0x06e8
			m.wr8(m.r[11], uint16(0x616), uint16(2))
			m.cycles += 8
		}
	case 0x06e8:
		{
			m.ip = 0x06ee
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x06ee:
		{
			m.ip = 0x06f0
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x06f0:
		{
			m.ip = 0x06f3
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x61a)))
			m.cycles += 8
		}
	case 0x06f3:
		{
			m.ip = 0x06f7
			m.wr16(m.r[11], uint16(0x223), m.alu("sub", 16, m.rd16(m.r[11], uint16(0x223)), m.r[0]))
			m.cycles += 16
		}
	case 0x06f7:
		{
			m.ip = 0x06f9
			m.r[0] = m.shift("shl", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x06f9:
		{
			m.ip = 0x06fd
			m.wr16(m.r[11], uint16(0x225), m.alu("sub", 16, m.rd16(m.r[11], uint16(0x225)), m.r[0]))
			m.cycles += 16
		}
	case 0x06fd:
		{
			m.ip = 0x0700
			m.r[0] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x0700:
		{
			m.ip = 0x0701
			m.r[0] = m.unary("inc", 16, m.r[0])
			m.cycles += 3
		}
	case 0x0701:
		{
			m.ip = 0x0704
			m.wr16(m.r[11], uint16(0x23b), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0704:
		{
			m.ip = 0x0707
			target := uint16(8099)
			m.push(0x0707)
			m.ip = target
			m.cycles += 19
		}
	case 0x0707:
		{
			m.ip = 0x070c
			m.wr8(m.r[11], uint16(0x2074), uint16(1))
			m.cycles += 8
		}
	case 0x070c:
		{
			m.ip = 0x070d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x070d:
		{
			m.ip = 0x0712
			m.wr8(m.r[11], uint16(0x270), uint16(0))
			m.cycles += 8
		}
	case 0x0712:
		{
			m.ip = 0x0716
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x0716:
		{
			m.ip = 0x071a
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0xd93))
			m.cycles += 8
		}
	case 0x071a:
		{
			m.ip = 0x071d
			m.wr16(m.r[11], uint16(0x21d), m.r[0])
			m.cycles += 8
		}
	case 0x071d:
		{
			m.ip = 0x0721
			m.r[2] = m.rd16(m.r[11], uint16(m.r[3]+0x479))
			m.cycles += 8
		}
	case 0x0721:
		{
			m.ip = 0x0724
			target := uint16(17942)
			m.push(0x0724)
			m.ip = target
			m.cycles += 19
		}
	case 0x0724:
		{
			m.ip = 0x0727
			m.r[0] = uint16(44460)
			m.cycles += 2
		}
	case 0x0727:
		{
			m.ip = 0x072a
			target := uint16(18137)
			m.push(0x072a)
			m.ip = target
			m.cycles += 19
		}
	case 0x072a:
		{
			m.ip = 0x072e
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x072e:
		{
			m.ip = 0x0732
			m.r[2] = m.rd16(m.r[11], uint16(m.r[3]+0x46f))
			m.cycles += 8
		}
	case 0x0732:
		{
			m.ip = 0x0735
			target := uint16(17942)
			m.push(0x0735)
			m.ip = target
			m.cycles += 19
		}
	case 0x0735:
		{
			m.ip = 0x0738
			m.r[0] = uint16(42710)
			m.cycles += 2
		}
	case 0x0738:
		{
			m.ip = 0x073b
			target := uint16(18137)
			m.push(0x073b)
			m.ip = target
			m.cycles += 19
		}
	case 0x073b:
		{
			m.ip = 0x073f
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x073f:
		{
			m.ip = 0x0743
			m.r[2] = m.rd16(m.r[11], uint16(m.r[3]+0x483))
			m.cycles += 8
		}
	case 0x0743:
		{
			m.ip = 0x0746
			target := uint16(17942)
			m.push(0x0746)
			m.ip = target
			m.cycles += 19
		}
	case 0x0746:
		{
			m.ip = 0x074b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(1))
			m.cycles += 16
		}
	case 0x074b:
		{
			m.ip = 0x074d
			if m.zf {
				m.ip = uint16(1884)
			}
			m.cycles += 8
		}
	case 0x074d:
		{
			m.ip = 0x0752
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x271)), uint16(1))
			m.cycles += 16
		}
	case 0x0752:
		{
			m.ip = 0x0754
			if !m.zf {
				m.ip = uint16(1884)
			}
			m.cycles += 8
		}
	case 0x0754:
		{
			m.ip = 0x0759
			m.wr8(m.r[11], uint16(0x271), uint16(0))
			m.cycles += 8
		}
	case 0x0759:
		{
			// Continue directly into world setup.
			m.ip = 0x075c
			m.cycles += 19
		}
	case 0x075c:
		{
			m.ip = 0x0760
			m.r[6] = uint16(0x1056)
			m.cycles += 3
		}
	case 0x0760:
		{
			m.ip = 0x0763
			target := uint16(18306)
			m.push(0x0763)
			m.ip = target
			m.cycles += 19
		}
	case 0x0763:
		{
			m.ip = 0x0766
			m.r[0] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x0766:
		{
			m.ip = 0x0769
			target := uint16(18137)
			m.push(0x0769)
			m.ip = target
			m.cycles += 19
		}
	case 0x0769:
		{
			m.ip = 0x076d
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x076d:
		{
			m.ip = 0x0770
			target := uint16(9400)
			m.push(0x0770)
			m.ip = target
			m.cycles += 19
		}
	case 0x0770:
		{
			m.ip = 0x0774
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x0774:
		{
			m.ip = 0x0777
			target := uint16(9656)
			m.push(0x0777)
			m.ip = target
			m.cycles += 19
		}
	case 0x0777:
		{
			m.ip = 0x077a
			m.r[0] = uint16(261)
			m.cycles += 2
		}
	case 0x077a:
		{
			m.ip = 0x077d
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x077d:
		{
			m.ip = 0x077e
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x077e:
		{
			m.ip = 0x077f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x077f:
		{
			m.ip = 0x0784
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x270)), uint16(1))
			m.cycles += 16
		}
	case 0x0784:
		{
			m.ip = 0x0786
			if m.zf {
				m.ip = uint16(1937)
			}
			m.cycles += 8
		}
	case 0x0786:
		{
			m.ip = 0x078b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x078b:
		{
			m.ip = 0x078d
			if !m.zf {
				m.ip = uint16(1936)
			}
			m.cycles += 8
		}
	case 0x078d:
		{
			m.ip = 0x0790
			m.ip = uint16(2064)
			m.cycles += 15
		}
	case 0x0790:
		{
			m.ip = 0x0791
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0791:
		{
			m.ip = 0x0796
			m.wr8(m.r[11], uint16(0x270), uint16(0))
			m.cycles += 8
		}
	case 0x0796:
		{
			m.ip = 0x079a
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x079a:
		{
			m.ip = 0x079e
			m.r[2] = m.rd16(m.r[11], uint16(m.r[3]+0x483))
			m.cycles += 8
		}
	case 0x079e:
		{
			m.ip = 0x07a1
			target := uint16(17942)
			m.push(0x07a1)
			m.ip = target
			m.cycles += 19
		}
	case 0x07a1:
		{
			m.ip = 0x07a5
			m.r[6] = uint16(0x1056)
			m.cycles += 3
		}
	case 0x07a5:
		{
			m.ip = 0x07a8
			target := uint16(18306)
			m.push(0x07a8)
			m.ip = target
			m.cycles += 19
		}
	case 0x07a8:
		{
			m.ip = 0x07ab
			m.r[0] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x07ab:
		{
			m.ip = 0x07ae
			target := uint16(18137)
			m.push(0x07ae)
			m.ip = target
			m.cycles += 19
		}
	case 0x07ae:
		{
			m.ip = 0x07b3
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x07b3:
		{
			m.ip = 0x07b5
			if !m.zf {
				m.ip = uint16(2064)
			}
			m.cycles += 8
		}
	case 0x07b5:
		{
			m.ip = 0x07b9
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x07b9:
		{
			m.ip = 0x07bc
			target := uint16(9487)
			m.push(0x07bc)
			m.ip = target
			m.cycles += 19
		}
	case 0x07bc:
		{
			m.ip = 0x07c2
			m.wr16(m.r[11], uint16(0x223), uint16(0))
			m.cycles += 8
		}
	case 0x07c2:
		{
			m.ip = 0x07c8
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x07c8:
		{
			m.ip = 0x07cb
			m.r[7] = uint16(1120)
			m.cycles += 2
		}
	case 0x07cb:
		{
			m.ip = 0x07ce
			target := uint16(16936)
			m.push(0x07ce)
			m.ip = target
			m.cycles += 19
		}
	case 0x07ce:
		{
			m.ip = 0x07d1
			target := uint16(12473)
			m.push(0x07d1)
			m.ip = target
			m.cycles += 19
		}
	case 0x07d1:
		{
			m.ip = 0x07d4
			m.r[6] = uint16(55460)
			m.cycles += 2
		}
	case 0x07d4:
		{
			m.ip = 0x07d7
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x07d7:
		{
			m.ip = 0x07da
			target := uint16(16909)
			m.push(0x07da)
			m.ip = target
			m.cycles += 19
		}
	case 0x07da:
		{
			m.ip = 0x07e0
			m.wr16(m.r[11], uint16(0x102b), uint16(55440))
			m.cycles += 8
		}
	case 0x07e0:
		{
			m.ip = 0x07e6
			m.wr16(m.r[11], uint16(0x102d), uint16(3128))
			m.cycles += 8
		}
	case 0x07e6:
		{
			m.ip = 0x07e9
			target := uint16(17038)
			m.push(0x07e9)
			m.ip = target
			m.cycles += 19
		}
	case 0x07e9:
		{
			m.ip = 0x07ef
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x07ef:
		{
			m.ip = 0x07f5
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x07f5:
		{
			m.ip = 0x07f8
			m.r[7] = uint16(1180)
			m.cycles += 2
		}
	case 0x07f8:
		{
			m.ip = 0x07fb
			target := uint16(16936)
			m.push(0x07fb)
			m.ip = target
			m.cycles += 19
		}
	case 0x07fb:
		{
			m.ip = 0x07fe
			target := uint16(12473)
			m.push(0x07fe)
			m.ip = target
			m.cycles += 19
		}
	case 0x07fe:
		{
			m.ip = 0x0801
			m.r[6] = uint16(55476)
			m.cycles += 2
		}
	case 0x0801:
		{
			m.ip = 0x0804
			m.r[7] = uint16(60)
			m.cycles += 2
		}
	case 0x0804:
		{
			m.ip = 0x0807
			target := uint16(16909)
			m.push(0x0807)
			m.ip = target
			m.cycles += 19
		}
	case 0x0807:
		{
			m.ip = 0x080d
			m.wr16(m.r[11], uint16(0x102f), uint16(3180))
			m.cycles += 8
		}
	case 0x080d:
		{
			m.ip = 0x0810
			target := uint16(17038)
			m.push(0x0810)
			m.ip = target
			m.cycles += 19
		}
	case 0x0810:
		{
			m.ip = 0x0814
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x0814:
		{
			m.ip = 0x0818
			m.r[6] = uint16(0x1045)
			m.cycles += 3
		}
	case 0x0818:
		{
			m.ip = 0x081b
			target := uint16(18306)
			m.push(0x081b)
			m.ip = target
			m.cycles += 19
		}
	case 0x081b:
		{
			m.ip = 0x081f
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x081f:
		{
			m.ip = 0x0822
			m.r[0] = uint16(261)
			m.cycles += 2
		}
	case 0x0822:
		{
			m.ip = 0x0825
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x0825:
		{
			m.ip = 0x0826
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0826:
		{
			m.ip = 0x0827
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0827:
		{
			m.ip = 0x082b
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x082b:
		{
			m.ip = 0x0830
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(2))
			m.cycles += 16
		}
	case 0x0830:
		{
			m.ip = 0x0832
			if m.zf {
				m.ip = uint16(2154)
			}
			m.cycles += 8
		}
	case 0x0832:
		{
			m.ip = 0x0837
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x0837:
		{
			m.ip = 0x0839
			if !m.zf {
				m.ip = uint16(2108)
			}
			m.cycles += 8
		}
	case 0x0839:
		{
			m.ip = 0x083c
			m.ip = uint16(2310)
			m.cycles += 15
		}
	case 0x083c:
		{
			m.ip = 0x083f
			target := uint16(16844)
			m.push(0x083f)
			m.ip = target
			m.cycles += 19
		}
	case 0x083f:
		{
			m.ip = 0x0842
			target := uint16(12473)
			m.push(0x0842)
			m.ip = target
			m.cycles += 19
		}
	case 0x0842:
		{
			m.ip = 0x0845
			m.r[7] = uint16(1180)
			m.cycles += 2
		}
	case 0x0845:
		{
			m.ip = 0x0848
			target := uint16(16936)
			m.push(0x0848)
			m.ip = target
			m.cycles += 19
		}
	case 0x0848:
		{
			m.ip = 0x084e
			m.wr16(m.r[11], uint16(0x102d), uint16(3180))
			m.cycles += 8
		}
	case 0x084e:
		{
			m.ip = 0x0851
			target := uint16(13361)
			m.push(0x0851)
			m.ip = target
			m.cycles += 19
		}
	case 0x0851:
		{
			m.ip = 0x0854
			target := uint16(11821)
			m.push(0x0854)
			m.ip = target
			m.cycles += 19
		}
	case 0x0854:
		{
			m.ip = 0x0857
			m.r[7] = uint16(60)
			m.cycles += 2
		}
	case 0x0857:
		{
			m.ip = 0x085a
			m.r[1] = uint16(5)
			m.cycles += 2
		}
	case 0x085a:
		{
			m.ip = 0x085d
			m.r[2] = uint16(20)
			m.cycles += 2
		}
	case 0x085d:
		{
			m.ip = 0x0863
			m.wr16(m.r[11], uint16(0x102d), uint16(3180))
			m.cycles += 8
		}
	case 0x0863:
		{
			m.ip = 0x0866
			target := uint16(2450)
			m.push(0x0866)
			m.ip = target
			m.cycles += 19
		}
	case 0x0866:
		{
			m.ip = 0x0869
			target := uint16(15936)
			m.push(0x0869)
			m.ip = target
			m.cycles += 19
		}
	case 0x0869:
		{
			m.ip = 0x086a
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x086a:
		{
			m.ip = 0x086d
			m.r[6] = uint16(55440)
			m.cycles += 2
		}
	case 0x086d:
		{
			m.ip = 0x0870
			m.r[6] = m.alu("add", 16, m.r[6], uint16(20))
			m.cycles += 4
		}
	case 0x0870:
		{
			m.ip = 0x0873
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x0873:
		{
			m.ip = 0x0878
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x0878:
		{
			m.ip = 0x087a
			if !m.zf {
				m.ip = uint16(2176)
			}
			m.cycles += 8
		}
	case 0x087a:
		{
			m.ip = 0x087d
			m.r[6] = m.alu("add", 16, m.r[6], uint16(16))
			m.cycles += 4
		}
	case 0x087d:
		{
			m.ip = 0x0880
			m.r[7] = m.alu("add", 16, m.r[7], uint16(60))
			m.cycles += 4
		}
	case 0x0880:
		{
			m.ip = 0x0883
			target := uint16(16909)
			m.push(0x0883)
			m.ip = target
			m.cycles += 19
		}
	case 0x0883:
		{
			m.ip = 0x0886
			target := uint16(12473)
			m.push(0x0886)
			m.ip = target
			m.cycles += 19
		}
	case 0x0886:
		{
			m.ip = 0x088a
			m.push(m.rd16(m.r[11], uint16(0x23b)))
			m.cycles += 11
		}
	case 0x088a:
		{
			m.ip = 0x088f
			m.wr8(m.r[11], uint16(0x23b), uint16(1))
			m.cycles += 8
		}
	case 0x088f:
		{
			m.ip = 0x0892
			m.r[7] = uint16(1120)
			m.cycles += 2
		}
	case 0x0892:
		{
			m.ip = 0x0895
			target := uint16(16936)
			m.push(0x0895)
			m.ip = target
			m.cycles += 19
		}
	case 0x0895:
		{
			m.ip = 0x089a
			m.wr8(m.r[11], uint16(0x23b), uint16(2))
			m.cycles += 8
		}
	case 0x089a:
		{
			m.ip = 0x089d
			m.r[7] = uint16(1180)
			m.cycles += 2
		}
	case 0x089d:
		{
			m.ip = 0x08a0
			target := uint16(16936)
			m.push(0x08a0)
			m.ip = target
			m.cycles += 19
		}
	case 0x08a0:
		{
			m.ip = 0x08a4
			m.wr16(m.r[11], uint16(0x23b), m.pop())
			m.cycles += 8
		}
	case 0x08a4:
		{
			m.ip = 0x08a8
			m.push(m.rd16(m.r[11], uint16(0x233)))
			m.cycles += 11
		}
	case 0x08a8:
		{
			m.ip = 0x08ae
			m.wr16(m.r[11], uint16(0x102b), uint16(55440))
			m.cycles += 8
		}
	case 0x08ae:
		{
			m.ip = 0x08b4
			m.wr16(m.r[11], uint16(0x102d), uint16(3128))
			m.cycles += 8
		}
	case 0x08b4:
		{
			m.ip = 0x08b7
			m.r[0] = m.rd16(m.r[11], uint16(0x237))
			m.cycles += 8
		}
	case 0x08b7:
		{
			m.ip = 0x08bc
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x08bc:
		{
			m.ip = 0x08be
			if !m.zf {
				m.ip = uint16(2246)
			}
			m.cycles += 8
		}
	case 0x08be:
		{
			m.ip = 0x08c3
			m.wr16(m.r[11], uint16(0x102d), m.alu("add", 16, m.rd16(m.r[11], uint16(0x102d)), uint16(52)))
			m.cycles += 16
		}
	case 0x08c3:
		{
			m.ip = 0x08c6
			m.r[0] = m.rd16(m.r[11], uint16(0x239))
			m.cycles += 8
		}
	case 0x08c6:
		{
			m.ip = 0x08c9
			m.wr16(m.r[11], uint16(0x233), m.r[0])
			m.cycles += 8
		}
	case 0x08c9:
		{
			m.ip = 0x08cd
			m.push(m.rd16(m.r[11], uint16(0x233)))
			m.cycles += 11
		}
	case 0x08cd:
		{
			m.ip = 0x08d0
			target := uint16(13361)
			m.push(0x08d0)
			m.ip = target
			m.cycles += 19
		}
	case 0x08d0:
		{
			m.ip = 0x08d4
			m.wr16(m.r[11], uint16(0x233), m.pop())
			m.cycles += 8
		}
	case 0x08d4:
		{
			m.ip = 0x08d7
			target := uint16(17038)
			m.push(0x08d7)
			m.ip = target
			m.cycles += 19
		}
	case 0x08d7:
		{
			m.ip = 0x08db
			m.wr16(m.r[11], uint16(0x233), m.pop())
			m.cycles += 8
		}
	case 0x08db:
		{
			m.ip = 0x08de
			target := uint16(11821)
			m.push(0x08de)
			m.ip = target
			m.cycles += 19
		}
	case 0x08de:
		{
			m.ip = 0x08e1
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x08e1:
		{
			m.ip = 0x08e4
			m.r[1] = uint16(200)
			m.cycles += 2
		}
	case 0x08e4:
		{
			m.ip = 0x08e7
			m.r[2] = uint16(20)
			m.cycles += 2
		}
	case 0x08e7:
		{
			m.ip = 0x08ed
			m.wr16(m.r[11], uint16(0x102d), uint16(3128))
			m.cycles += 8
		}
	case 0x08ed:
		{
			m.ip = 0x08f2
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x08f2:
		{
			m.ip = 0x08f4
			if m.zf {
				m.ip = uint16(2303)
			}
			m.cycles += 8
		}
	case 0x08f4:
		{
			m.ip = 0x08f7
			m.r[7] = m.alu("add", 16, m.r[7], uint16(60))
			m.cycles += 4
		}
	case 0x08f7:
		{
			m.ip = 0x08fa
			m.r[2] = m.alu("add", 16, m.r[2], uint16(16))
			m.cycles += 4
		}
	case 0x08fa:
		{
			m.ip = 0x08ff
			m.wr16(m.r[11], uint16(0x102d), m.alu("add", 16, m.rd16(m.r[11], uint16(0x102d)), uint16(52)))
			m.cycles += 16
		}
	case 0x08ff:
		{
			m.ip = 0x0902
			target := uint16(2450)
			m.push(0x0902)
			m.ip = target
			m.cycles += 19
		}
	case 0x0902:
		{
			m.ip = 0x0905
			target := uint16(15936)
			m.push(0x0905)
			m.ip = target
			m.cycles += 19
		}
	case 0x0905:
		{
			m.ip = 0x0906
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0906:
		{
			m.ip = 0x090c
			m.wr16(m.r[11], uint16(0x223), uint16(0))
			m.cycles += 8
		}
	case 0x090c:
		{
			m.ip = 0x090f
			m.r[7] = uint16(1120)
			m.cycles += 2
		}
	case 0x090f:
		{
			m.ip = 0x0912
			target := uint16(16936)
			m.push(0x0912)
			m.ip = target
			m.cycles += 19
		}
	case 0x0912:
		{
			m.ip = 0x0915
			target := uint16(12473)
			m.push(0x0915)
			m.ip = target
			m.cycles += 19
		}
	case 0x0915:
		{
			m.ip = 0x091b
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x091b:
		{
			m.ip = 0x091e
			m.r[7] = uint16(1180)
			m.cycles += 2
		}
	case 0x091e:
		{
			m.ip = 0x0921
			target := uint16(16936)
			m.push(0x0921)
			m.ip = target
			m.cycles += 19
		}
	case 0x0921:
		{
			m.ip = 0x0924
			target := uint16(12473)
			m.push(0x0924)
			m.ip = target
			m.cycles += 19
		}
	case 0x0924:
		{
			m.ip = 0x0927
			target := uint16(11821)
			m.push(0x0927)
			m.ip = target
			m.cycles += 19
		}
	case 0x0927:
		{
			m.ip = 0x092c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(0))
			m.cycles += 16
		}
	case 0x092c:
		{
			m.ip = 0x092e
			if m.zf {
				m.ip = uint16(2365)
			}
			m.cycles += 8
		}
	case 0x092e:
		{
			m.ip = 0x0934
			m.wr16(m.r[11], uint16(0x223), uint16(0))
			m.cycles += 8
		}
	case 0x0934:
		{
			m.ip = 0x093a
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x093a:
		{
			m.ip = 0x093d
			target := uint16(13361)
			m.push(0x093d)
			m.ip = target
			m.cycles += 19
		}
	case 0x093d:
		{
			m.ip = 0x0942
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(1))
			m.cycles += 16
		}
	case 0x0942:
		{
			m.ip = 0x0944
			if m.zf {
				m.ip = uint16(2387)
			}
			m.cycles += 8
		}
	case 0x0944:
		{
			m.ip = 0x094a
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x094a:
		{
			m.ip = 0x0950
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x0950:
		{
			m.ip = 0x0953
			target := uint16(13361)
			m.push(0x0953)
			m.ip = target
			m.cycles += 19
		}
	case 0x0953:
		{
			m.ip = 0x0959
			m.wr16(m.r[11], uint16(0x102d), uint16(3128))
			m.cycles += 8
		}
	case 0x0959:
		{
			m.ip = 0x095f
			m.wr16(m.r[11], uint16(0x102f), uint16(3180))
			m.cycles += 8
		}
	case 0x095f:
		{
			m.ip = 0x0962
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x0962:
		{
			m.ip = 0x0965
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x0965:
		{
			m.ip = 0x0968
			m.r[2] = uint16(20)
			m.cycles += 2
		}
	case 0x0968:
		{
			m.ip = 0x096e
			m.wr16(m.r[11], uint16(0x102d), uint16(3128))
			m.cycles += 8
		}
	case 0x096e:
		{
			m.ip = 0x0971
			target := uint16(2450)
			m.push(0x0971)
			m.ip = target
			m.cycles += 19
		}
	case 0x0971:
		{
			m.ip = 0x0976
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(0))
			m.cycles += 16
		}
	case 0x0976:
		{
			m.ip = 0x0978
			if m.zf {
				m.ip = uint16(2433)
			}
			m.cycles += 8
		}
	case 0x0978:
		{
			m.ip = 0x097e
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x097e:
		{
			m.ip = 0x0981
			target := uint16(15936)
			m.push(0x0981)
			m.ip = target
			m.cycles += 19
		}
	case 0x0981:
		{
			m.ip = 0x0986
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(1))
			m.cycles += 16
		}
	case 0x0986:
		{
			m.ip = 0x0988
			if m.zf {
				m.ip = uint16(2449)
			}
			m.cycles += 8
		}
	case 0x0988:
		{
			m.ip = 0x098e
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x098e:
		{
			m.ip = 0x0991
			target := uint16(15936)
			m.push(0x0991)
			m.ip = target
			m.cycles += 19
		}
	case 0x0991:
		{
			m.ip = 0x0992
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0992:
		{
			m.ip = 0x0993
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x0993:
		{
			m.ip = 0x0994
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x0994:
		{
			m.ip = 0x0999
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x0999:
		{
			m.ip = 0x099b
			if !m.zf {
				m.ip = uint16(2466)
			}
			m.cycles += 8
		}
	case 0x099b:
		{
			m.ip = 0x09a0
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(0))
			m.cycles += 16
		}
	case 0x09a0:
		{
			m.ip = 0x09a2
			if m.zf {
				m.ip = uint16(2497)
			}
			m.cycles += 8
		}
	case 0x09a2:
		{
			m.ip = 0x09a5
			m.r[6] = uint16(55440)
			m.cycles += 2
		}
	case 0x09a5:
		{
			m.ip = 0x09a7
			m.r[6] = m.alu("add", 16, m.r[6], m.r[2])
			m.cycles += 4
		}
	case 0x09a7:
		{
			m.ip = 0x09aa
			target := uint16(16909)
			m.push(0x09aa)
			m.ip = target
			m.cycles += 19
		}
	case 0x09aa:
		{
			m.ip = 0x09b0
			m.wr16(m.r[11], uint16(0x102b), uint16(55440))
			m.cycles += 8
		}
	case 0x09b0:
		{
			m.ip = 0x09b3
			target := uint16(17038)
			m.push(0x09b3)
			m.ip = target
			m.cycles += 19
		}
	case 0x09b3:
		{
			m.ip = 0x09b8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x09b8:
		{
			m.ip = 0x09ba
			if !m.zf {
				m.ip = uint16(2530)
			}
			m.cycles += 8
		}
	case 0x09ba:
		{
			m.ip = 0x09bf
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(1))
			m.cycles += 16
		}
	case 0x09bf:
		{
			m.ip = 0x09c1
			if m.zf {
				m.ip = uint16(2530)
			}
			m.cycles += 8
		}
	case 0x09c1:
		{
			m.ip = 0x09c4
			m.r[6] = uint16(55440)
			m.cycles += 2
		}
	case 0x09c4:
		{
			m.ip = 0x09c7
			m.r[6] = m.alu("add", 16, m.r[6], uint16(36))
			m.cycles += 4
		}
	case 0x09c7:
		{
			m.ip = 0x09ca
			m.r[7] = uint16(60)
			m.cycles += 2
		}
	case 0x09ca:
		{
			m.ip = 0x09cd
			target := uint16(16909)
			m.push(0x09cd)
			m.ip = target
			m.cycles += 19
		}
	case 0x09cd:
		{
			m.ip = 0x09d3
			m.wr16(m.r[11], uint16(0x102b), uint16(55440))
			m.cycles += 8
		}
	case 0x09d3:
		{
			m.ip = 0x09d9
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x09d9:
		{
			m.ip = 0x09dc
			target := uint16(17038)
			m.push(0x09dc)
			m.ip = target
			m.cycles += 19
		}
	case 0x09dc:
		{
			m.ip = 0x09e2
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x09e2:
		{
			m.ip = 0x09e8
			m.wr16(m.r[11], uint16(0x106a), uint16(150))
			m.cycles += 8
		}
	case 0x09e8:
		{
			m.ip = 0x09eb
			target := uint16(19311)
			m.push(0x09eb)
			m.ip = target
			m.cycles += 19
		}
	case 0x09eb:
		{
			m.ip = 0x09ec
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x09ec:
		{
			m.ip = 0x09ed
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x09ed:
		{
			m.ip = 0x09ee
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x09ee:
		{
			m.ip = 0x09f3
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x09f3:
		{
			m.ip = 0x09f5
			if !m.zf {
				m.ip = uint16(2556)
			}
			m.cycles += 8
		}
	case 0x09f5:
		{
			m.ip = 0x09fa
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(0))
			m.cycles += 16
		}
	case 0x09fa:
		{
			m.ip = 0x09fc
			if m.zf {
				m.ip = uint16(2587)
			}
			m.cycles += 8
		}
	case 0x09fc:
		{
			m.ip = 0x09ff
			m.r[6] = uint16(54880)
			m.cycles += 2
		}
	case 0x09ff:
		{
			m.ip = 0x0a01
			m.r[6] = m.alu("add", 16, m.r[6], m.r[2])
			m.cycles += 4
		}
	case 0x0a01:
		{
			m.ip = 0x0a04
			target := uint16(16909)
			m.push(0x0a04)
			m.ip = target
			m.cycles += 19
		}
	case 0x0a04:
		{
			m.ip = 0x0a0a
			m.wr16(m.r[11], uint16(0x102b), uint16(54880))
			m.cycles += 8
		}
	case 0x0a0a:
		{
			m.ip = 0x0a0d
			target := uint16(17038)
			m.push(0x0a0d)
			m.ip = target
			m.cycles += 19
		}
	case 0x0a0d:
		{
			m.ip = 0x0a12
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x0a12:
		{
			m.ip = 0x0a14
			if !m.zf {
				m.ip = uint16(2620)
			}
			m.cycles += 8
		}
	case 0x0a14:
		{
			m.ip = 0x0a19
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(1))
			m.cycles += 16
		}
	case 0x0a19:
		{
			m.ip = 0x0a1b
			if m.zf {
				m.ip = uint16(2620)
			}
			m.cycles += 8
		}
	case 0x0a1b:
		{
			m.ip = 0x0a1e
			m.r[6] = uint16(54880)
			m.cycles += 2
		}
	case 0x0a1e:
		{
			m.ip = 0x0a21
			m.r[6] = m.alu("add", 16, m.r[6], uint16(36))
			m.cycles += 4
		}
	case 0x0a21:
		{
			m.ip = 0x0a24
			m.r[7] = uint16(60)
			m.cycles += 2
		}
	case 0x0a24:
		{
			m.ip = 0x0a27
			target := uint16(16909)
			m.push(0x0a27)
			m.ip = target
			m.cycles += 19
		}
	case 0x0a27:
		{
			m.ip = 0x0a2d
			m.wr16(m.r[11], uint16(0x102b), uint16(54880))
			m.cycles += 8
		}
	case 0x0a2d:
		{
			m.ip = 0x0a33
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x0a33:
		{
			m.ip = 0x0a36
			target := uint16(17038)
			m.push(0x0a36)
			m.ip = target
			m.cycles += 19
		}
	case 0x0a36:
		{
			m.ip = 0x0a3c
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x0a3c:
		{
			m.ip = 0x0a3d
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x0a3d:
		{
			m.ip = 0x0a3e
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x0a3e:
		{
			m.ip = 0x0a3f
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x0a3f:
		{
			m.ip = 0x0a44
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1070)), uint16(1))
			m.cycles += 16
		}
	case 0x0a44:
		{
			m.ip = 0x0a46
			if m.zf {
				m.ip = uint16(2655)
			}
			m.cycles += 8
		}
	case 0x0a46:
		{
			m.ip = 0x0a4c
			m.wr16(m.r[11], uint16(0x106a), uint16(150))
			m.cycles += 8
		}
	case 0x0a4c:
		{
			m.ip = 0x0a4f
			target := uint16(19311)
			m.push(0x0a4f)
			m.ip = target
			m.cycles += 19
		}
	case 0x0a4f:
		{
			m.ip = 0x0a54
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1070)), uint16(1))
			m.cycles += 16
		}
	case 0x0a54:
		{
			m.ip = 0x0a56
			if m.zf {
				m.ip = uint16(2655)
			}
			m.cycles += 8
		}
	case 0x0a56:
		{
			m.ip = 0x0a57
			m.r[1] = m.unary("dec", 16, m.r[1])
			m.cycles += 3
		}
	case 0x0a57:
		{
			m.ip = 0x0a5a
			m.alu("sub", 16, m.r[1], uint16(0))
			m.cycles += 4
		}
	case 0x0a5a:
		{
			m.ip = 0x0a5c
			if m.zf {
				m.ip = uint16(2655)
			}
			m.cycles += 8
		}
	case 0x0a5c:
		{
			m.ip = 0x0a5f
			m.ip = uint16(2450)
			m.cycles += 15
		}
	case 0x0a5f:
		{
			m.ip = 0x0a60
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0a60:
		{
			m.ip = 0x0a64
			m.r[6] = m.rd16(m.r[11], uint16(0x21d))
			m.cycles += 8
		}
	case 0x0a64:
		{
			m.ip = 0x0a68
			m.r[7] = uint16(0x19f)
			m.cycles += 3
		}
	case 0x0a68:
		{
			m.ip = 0x0a6b
			m.r[1] = uint16(102)
			m.cycles += 2
		}
	case 0x0a6b:
		{
			m.ip = 0x0a6e
			target := uint16(9392)
			m.push(0x0a6e)
			m.ip = target
			m.cycles += 19
		}
	case 0x0a6e:
		{
			m.ip = 0x0a72
			m.r[6] = uint16(0xff5)
			m.cycles += 3
		}
	case 0x0a72:
		{
			m.ip = 0x0a75
			m.r[3] = uint16(65535)
			m.cycles += 2
		}
	case 0x0a75:
		{
			m.ip = 0x0a78
			target := uint16(9346)
			m.push(0x0a78)
			m.ip = target
			m.cycles += 19
		}
	case 0x0a78:
		{
			m.ip = 0x0a7d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(0))
			m.cycles += 16
		}
	case 0x0a7d:
		{
			m.ip = 0x0a7f
			if m.zf {
				m.ip = uint16(2731)
			}
			m.cycles += 8
		}
	case 0x0a7f:
		{
			m.ip = 0x0a83
			m.r[6] = m.rd16(m.r[11], uint16(0x21d))
			m.cycles += 8
		}
	case 0x0a83:
		{
			m.ip = 0x0a86
			m.r[6] = m.alu("add", 16, m.r[6], uint16(102))
			m.cycles += 4
		}
	case 0x0a86:
		{
			m.ip = 0x0a8a
			m.r[7] = uint16(0x19f)
			m.cycles += 3
		}
	case 0x0a8a:
		{
			m.ip = 0x0a8d
			m.r[1] = uint16(18)
			m.cycles += 2
		}
	case 0x0a8d:
		{
			m.ip = 0x0a90
			target := uint16(9392)
			m.push(0x0a90)
			m.ip = target
			m.cycles += 19
		}
	case 0x0a90:
		{
			m.ip = 0x0a95
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(2))
			m.cycles += 16
		}
	case 0x0a95:
		{
			m.ip = 0x0a97
			if m.zf {
				m.ip = uint16(2731)
			}
			m.cycles += 8
		}
	case 0x0a97:
		{
			m.ip = 0x0a99
			m.r[3] = m.alu("xor", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x0a99:
		{
			m.ip = 0x0a9d
			m.set8(3, 0, m.rd8(m.r[11], uint16(0x61a)))
			m.cycles += 8
		}
	case 0x0a9d:
		{
			m.ip = 0x0a9f
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x0a9f:
		{
			m.ip = 0x0aa5
			m.wr16(m.r[11], uint16(m.r[3]+0x1ad), uint16(0))
			m.cycles += 8
		}
	case 0x0aa5:
		{
			m.ip = 0x0aab
			m.wr16(m.r[11], uint16(m.r[3]+0x1a5), uint16(0))
			m.cycles += 8
		}
	case 0x0aab:
		{
			m.ip = 0x0aac
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0aac:
		{
			m.ip = 0x0ab2
			m.wr16(m.r[11], uint16(0x223), uint16(0))
			m.cycles += 8
		}
	case 0x0ab2:
		{
			m.ip = 0x0ab8
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x0ab8:
		{
			m.ip = 0x0abd
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x246)), uint16(0))
			m.cycles += 16
		}
	case 0x0abd:
		{
			m.ip = 0x0abf
			if m.zf {
				m.ip = uint16(2784)
			}
			m.cycles += 8
		}
	case 0x0abf:
		{
			m.ip = 0x0ac4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x68)), uint16(20))
			m.cycles += 16
		}
	case 0x0ac4:
		{
			m.ip = 0x0ac6
			if !m.cf {
				m.ip = uint16(2771)
			}
			m.cycles += 8
		}
	case 0x0ac6:
		{
			m.ip = 0x0acb
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x68)), uint16(8))
			m.cycles += 16
		}
	case 0x0acb:
		{
			m.ip = 0x0acd
			if m.cf || m.zf {
				m.ip = uint16(2771)
			}
			m.cycles += 8
		}
	case 0x0acd:
		{
			m.ip = 0x0ad0
			target := uint16(13112)
			m.push(0x0ad0)
			m.ip = target
			m.cycles += 19
		}
	case 0x0ad0:
		{
			m.ip = 0x0ad2
			m.ip = uint16(2784)
			m.cycles += 15
		}
	case 0x0ad2:
		{
			m.ip = 0x0ad3
			m.cycles += 3
		}
	case 0x0ad3:
		{
			m.ip = 0x0ad6
			target := uint16(13361)
			m.push(0x0ad6)
			m.ip = target
			m.cycles += 19
		}
	case 0x0ad6:
		{
			m.ip = 0x0ad9
			target := uint16(12940)
			m.push(0x0ad9)
			m.ip = target
			m.cycles += 19
		}
	case 0x0ad9:
		{
			m.ip = 0x0ade
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x78)), uint16(0))
			m.cycles += 16
		}
	case 0x0ade:
		{
			m.ip = 0x0ae0
			if m.zf {
				m.ip = uint16(2836)
			}
			m.cycles += 8
		}
	case 0x0ae0:
		{
			m.ip = 0x0ae5
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x248)), uint16(0))
			m.cycles += 16
		}
	case 0x0ae5:
		{
			m.ip = 0x0ae7
			if m.zf {
				m.ip = uint16(2836)
			}
			m.cycles += 8
		}
	case 0x0ae7:
		{
			m.ip = 0x0aec
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x0aec:
		{
			m.ip = 0x0aee
			if !m.zf {
				m.ip = uint16(2836)
			}
			m.cycles += 8
		}
	case 0x0aee:
		{
			m.ip = 0x0af4
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x0af4:
		{
			m.ip = 0x0afa
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x0afa:
		{
			m.ip = 0x0aff
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x69)), uint16(20))
			m.cycles += 16
		}
	case 0x0aff:
		{
			m.ip = 0x0b01
			if !m.cf {
				m.ip = uint16(2830)
			}
			m.cycles += 8
		}
	case 0x0b01:
		{
			m.ip = 0x0b06
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x69)), uint16(8))
			m.cycles += 16
		}
	case 0x0b06:
		{
			m.ip = 0x0b08
			if m.cf || m.zf {
				m.ip = uint16(2830)
			}
			m.cycles += 8
		}
	case 0x0b08:
		{
			m.ip = 0x0b0b
			target := uint16(13112)
			m.push(0x0b0b)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b0b:
		{
			m.ip = 0x0b0d
			m.ip = uint16(2836)
			m.cycles += 15
		}
	case 0x0b0d:
		{
			m.ip = 0x0b0e
			m.cycles += 3
		}
	case 0x0b0e:
		{
			m.ip = 0x0b11
			target := uint16(13361)
			m.push(0x0b11)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b11:
		{
			m.ip = 0x0b14
			target := uint16(12940)
			m.push(0x0b14)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b14:
		{
			m.ip = 0x0b18
			m.r[0] = uint16(0x26d9)
			m.cycles += 3
		}
	case 0x0b18:
		{
			m.ip = 0x0b1d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x0b1d:
		{
			m.ip = 0x0b1f
			if !m.zf {
				m.ip = uint16(2851)
			}
			m.cycles += 8
		}
	case 0x0b1f:
		{
			m.ip = 0x0b23
			m.r[0] = uint16(0x2732)
			m.cycles += 3
		}
	case 0x0b23:
		{
			m.ip = 0x0b25
			target := m.r[0]
			m.push(0x0b25)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b25:
		{
			m.ip = 0x0b2a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x74)), uint16(1))
			m.cycles += 16
		}
	case 0x0b2a:
		{
			m.ip = 0x0b2c
			if m.zf {
				m.ip = uint16(2964)
			}
			m.cycles += 8
		}
	case 0x0b2c:
		{
			m.ip = 0x0b2f
			target := uint16(9749)
			m.push(0x0b2f)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b2f:
		{
			m.ip = 0x0b32
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x0b32:
		{
			m.ip = 0x0b35
			target := uint16(10173)
			m.push(0x0b35)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b35:
		{
			m.ip = 0x0b3a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x0b3a:
		{
			m.ip = 0x0b3c
			if !m.zf {
				m.ip = uint16(2882)
			}
			m.cycles += 8
		}
	case 0x0b3c:
		{
			m.ip = 0x0b3f
			m.r[6] = uint16(2)
			m.cycles += 2
		}
	case 0x0b3f:
		{
			m.ip = 0x0b42
			target := uint16(10173)
			m.push(0x0b42)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b42:
		{
			m.ip = 0x0b45
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x0b45:
		{
			m.ip = 0x0b49
			m.wr16(m.r[11], uint16(0x221), m.r[3])
			m.cycles += 8
		}
	case 0x0b49:
		{
			m.ip = 0x0b4e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x7a)), uint16(77))
			m.cycles += 16
		}
	case 0x0b4e:
		{
			m.ip = 0x0b50
			if m.zf {
				m.ip = uint16(2930)
			}
			m.cycles += 8
		}
	case 0x0b50:
		{
			m.ip = 0x0b54
			m.wr8(m.r[11], uint16(m.r[3]+0x227), m.unary("inc", 8, m.rd8(m.r[11], uint16(m.r[3]+0x227))))
			m.cycles += 15
		}
	case 0x0b54:
		{
			m.ip = 0x0b59
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x227)), uint16(15))
			m.cycles += 16
		}
	case 0x0b59:
		{
			m.ip = 0x0b5b
			if !m.zf {
				m.ip = uint16(2921)
			}
			m.cycles += 8
		}
	case 0x0b5b:
		{
			m.ip = 0x0b60
			m.wr8(m.r[11], uint16(m.r[3]+0x227), uint16(0))
			m.cycles += 8
		}
	case 0x0b60:
		{
			m.ip = 0x0b63
			target := uint16(15058)
			m.push(0x0b63)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b63:
		{
			m.ip = 0x0b66
			target := uint16(10221)
			m.push(0x0b66)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b66:
		{
			m.ip = 0x0b69
			target := uint16(9858)
			m.push(0x0b69)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b69:
		{
			m.ip = 0x0b6c
			target := uint16(15058)
			m.push(0x0b6c)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b6c:
		{
			m.ip = 0x0b6f
			target := uint16(10221)
			m.push(0x0b6f)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b6f:
		{
			m.ip = 0x0b71
			m.ip = uint16(2955)
			m.cycles += 15
		}
	case 0x0b71:
		{
			m.ip = 0x0b72
			m.cycles += 3
		}
	case 0x0b72:
		{
			m.ip = 0x0b76
			m.wr8(m.r[11], uint16(m.r[3]+0x22b), m.unary("inc", 8, m.rd8(m.r[11], uint16(m.r[3]+0x22b))))
			m.cycles += 15
		}
	case 0x0b76:
		{
			m.ip = 0x0b7b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x22b)), uint16(8))
			m.cycles += 16
		}
	case 0x0b7b:
		{
			m.ip = 0x0b7d
			if !m.zf {
				m.ip = uint16(2949)
			}
			m.cycles += 8
		}
	case 0x0b7d:
		{
			m.ip = 0x0b82
			m.wr8(m.r[11], uint16(m.r[3]+0x22b), uint16(0))
			m.cycles += 8
		}
	case 0x0b82:
		{
			m.ip = 0x0b84
			m.ip = uint16(2955)
			m.cycles += 15
		}
	case 0x0b84:
		{
			m.ip = 0x0b85
			m.cycles += 3
		}
	case 0x0b85:
		{
			m.ip = 0x0b88
			target := uint16(15058)
			m.push(0x0b88)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b88:
		{
			m.ip = 0x0b8b
			target := uint16(10221)
			m.push(0x0b8b)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b8b:
		{
			m.ip = 0x0b8e
			target := uint16(9858)
			m.push(0x0b8e)
			m.ip = target
			m.cycles += 19
		}
	case 0x0b8e:
		{
			m.ip = 0x0b8f
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x0b8f:
		{
			m.ip = 0x0b92
			m.alu("sub", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x0b92:
		{
			m.ip = 0x0b94
			if !m.zf {
				m.ip = uint16(2885)
			}
			m.cycles += 8
		}
	case 0x0b94:
		{
			m.ip = 0x0b95
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0b95:
		{
			m.ip = 0x0b9a
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x223)), uint16(0))
			m.cycles += 16
		}
	case 0x0b9a:
		{
			m.ip = 0x0b9c
			if !m.zf {
				m.ip = uint16(3001)
			}
			m.cycles += 8
		}
	case 0x0b9c:
		{
			m.ip = 0x0ba1
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x246)), uint16(0))
			m.cycles += 16
		}
	case 0x0ba1:
		{
			m.ip = 0x0ba3
			if m.zf {
				m.ip = uint16(3034)
			}
			m.cycles += 8
		}
	case 0x0ba3:
		{
			m.ip = 0x0ba8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x68)), uint16(20))
			m.cycles += 16
		}
	case 0x0ba8:
		{
			m.ip = 0x0baa
			if !m.cf {
				m.ip = uint16(2997)
			}
			m.cycles += 8
		}
	case 0x0baa:
		{
			m.ip = 0x0baf
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x68)), uint16(8))
			m.cycles += 16
		}
	case 0x0baf:
		{
			m.ip = 0x0bb1
			if m.cf || m.zf {
				m.ip = uint16(2997)
			}
			m.cycles += 8
		}
	case 0x0bb1:
		{
			m.ip = 0x0bb4
			target := uint16(13112)
			m.push(0x0bb4)
			m.ip = target
			m.cycles += 19
		}
	case 0x0bb4:
		{
			m.ip = 0x0bb5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0bb5:
		{
			m.ip = 0x0bb8
			target := uint16(13361)
			m.push(0x0bb8)
			m.ip = target
			m.cycles += 19
		}
	case 0x0bb8:
		{
			m.ip = 0x0bb9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0bb9:
		{
			m.ip = 0x0bbf
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x0bbf:
		{
			m.ip = 0x0bc5
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x0bc5:
		{
			m.ip = 0x0bca
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x69)), uint16(20))
			m.cycles += 16
		}
	case 0x0bca:
		{
			m.ip = 0x0bcc
			if !m.cf {
				m.ip = uint16(3031)
			}
			m.cycles += 8
		}
	case 0x0bcc:
		{
			m.ip = 0x0bd1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x69)), uint16(8))
			m.cycles += 16
		}
	case 0x0bd1:
		{
			m.ip = 0x0bd3
			if m.cf || m.zf {
				m.ip = uint16(3031)
			}
			m.cycles += 8
		}
	case 0x0bd3:
		{
			m.ip = 0x0bd6
			target := uint16(13112)
			m.push(0x0bd6)
			m.ip = target
			m.cycles += 19
		}
	case 0x0bd6:
		{
			m.ip = 0x0bd7
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0bd7:
		{
			m.ip = 0x0bda
			target := uint16(13361)
			m.push(0x0bda)
			m.ip = target
			m.cycles += 19
		}
	case 0x0bda:
		{
			m.ip = 0x0bdb
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0bdb:
		{
			m.ip = 0x0bdc
			m.push(m.r[8])
			m.cycles += 11
		}
	case 0x0bdc:
		{
			m.ip = 0x0bde
			m.r[8] = m.r[0]
			m.cycles += 2
		}
	case 0x0bde:
		{
			m.ip = 0x0be1
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x0be1:
		{
			m.ip = 0x0be4
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x0be4:
		{
			m.ip = 0x0be7
			m.r[1] = uint16(559)
			m.cycles += 2
		}
	case 0x0be7:
		{
			m.ip = 0x0beb
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(0x28a)))
			m.cycles += 16
		}
	case 0x0beb:
		{
			m.ip = 0x0bed
			if !m.zf {
				m.ip = uint16(3056)
			}
			m.cycles += 8
		}
	case 0x0bed:
		{
			m.ip = 0x0bf0
			m.r[1] = uint16(625)
			m.cycles += 2
		}
	case 0x0bf0:
		{
			m.ip = 0x0bf1
			m.df = false
			m.cycles += 2
		}
	case 0x0bf1:
		{
			m.ip = 0x0bf3
			m.stringOp("movs", 8)
			m.cycles += 2
		}
	case 0x0bf3:
		{
			m.ip = 0x0bf4
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x0bf4:
		{
			m.ip = 0x0bf5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0bf5:
		{
			m.ip = 0x0bf6
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x0bf6:
		{
			m.ip = 0x0bf7
			m.push(m.r[8])
			m.cycles += 11
		}
	case 0x0bf7:
		{
			m.ip = 0x0bfa
			m.r[1] = uint16(559)
			m.cycles += 2
		}
	case 0x0bfa:
		{
			m.ip = 0x0bfe
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(0x28a)))
			m.cycles += 16
		}
	case 0x0bfe:
		{
			m.ip = 0x0c00
			if !m.zf {
				m.ip = uint16(3075)
			}
			m.cycles += 8
		}
	case 0x0c00:
		{
			m.ip = 0x0c03
			m.r[1] = uint16(625)
			m.cycles += 2
		}
	case 0x0c03:
		{
			m.ip = 0x0c04
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x0c04:
		{
			m.ip = 0x0c05
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x0c05:
		{
			m.ip = 0x0c07
			m.r[11] = m.r[0]
			m.cycles += 2
		}
	case 0x0c07:
		{
			m.ip = 0x0c0a
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x0c0a:
		{
			m.ip = 0x0c0d
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x0c0d:
		{
			m.ip = 0x0c0e
			m.df = false
			m.cycles += 2
		}
	case 0x0c0e:
		{
			m.ip = 0x0c10
			m.stringOp("movs", 8)
			m.cycles += 2
		}
	case 0x0c10:
		{
			m.ip = 0x0c11
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x0c11:
		{
			m.ip = 0x0c12
			m.r[11] = m.pop()
			m.cycles += 8
		}
	case 0x0c12:
		{
			m.ip = 0x0c13
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0c13:
		{
			m.ip = 0x0c14
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x0c14:
		{
			m.ip = 0x0c15
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x0c15:
		{
			m.ip = 0x0c16
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x0c16:
		{
			m.ip = 0x0c17
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x0c17:
		{
			m.ip = 0x0c18
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x0c18:
		{
			m.ip = 0x0c19
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x0c19:
		{
			m.ip = 0x0c1a
			m.push(m.r[5])
			m.cycles += 11
		}
	case 0x0c1a:
		{
			m.ip = 0x0c1b
			m.push(m.r[8])
			m.cycles += 11
		}
	case 0x0c1b:
		{
			m.ip = 0x0c1c
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x0c1c:
		{
			m.ip = 0x0c21
			m.wr8(m.r[11], uint16(0x2074), uint16(1))
			m.cycles += 8
		}
	case 0x0c21:
		{
			m.ip = 0x0c24
			target := uint16(8099)
			m.push(0x0c24)
			m.ip = target
			m.cycles += 19
		}
	case 0x0c24:
		{
			m.ip = 0x0c26
			m.set8(0, 0, m.input(uint16(97)))
			m.cycles += 8
		}
	case 0x0c26:
		{
			m.ip = 0x0c28
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(252)))
			m.cycles += 4
		}
	case 0x0c28:
		{
			m.ip = 0x0c2a
			m.output(uint16(97), ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x0c2a:
		{
			m.ip = 0x0c2e
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x0c2e:
		{
			m.ip = 0x0c31
			m.r[0] = uint16(40960)
			m.cycles += 2
		}
	case 0x0c31:
		{
			m.ip = 0x0c32
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x0c32:
		{
			m.ip = 0x0c34
			m.r[11] = m.r[0]
			m.cycles += 2
		}
	case 0x0c34:
		{
			m.ip = 0x0c36
			m.r[6] = m.alu("xor", 16, m.r[6], m.r[6])
			m.cycles += 4
		}
	case 0x0c36:
		{
			m.ip = 0x0c38
			m.r[7] = m.alu("xor", 16, m.r[7], m.r[7])
			m.cycles += 4
		}
	case 0x0c38:
		{
			m.ip = 0x0c3b
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x0c3b:
		{
			m.ip = 0x0c3e
			m.r[1] = uint16(8192)
			m.cycles += 2
		}
	case 0x0c3e:
		{
			m.ip = 0x0c41
			m.r[0] = uint16(772)
			m.cycles += 2
		}
	case 0x0c41:
		{
			m.ip = 0x0c42
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0c42:
		{
			m.ip = 0x0c44
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x0c44:
		{
			m.ip = 0x0c47
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0c47:
		{
			m.ip = 0x0c49
			m.set8(0, 8, m.unary("dec", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x0c49:
		{
			m.ip = 0x0c4a
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0c4a:
		{
			m.ip = 0x0c4c
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x0c4c:
		{
			m.ip = 0x0c50
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0c50:
		{
			m.ip = 0x0c52
			m.set8(0, 8, m.unary("dec", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x0c52:
		{
			m.ip = 0x0c53
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0c53:
		{
			m.ip = 0x0c55
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x0c55:
		{
			m.ip = 0x0c59
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0c59:
		{
			m.ip = 0x0c5b
			m.set8(0, 8, m.unary("dec", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x0c5b:
		{
			m.ip = 0x0c5c
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0c5c:
		{
			m.ip = 0x0c5e
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x0c5e:
		{
			m.ip = 0x0c62
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0c62:
		{
			m.ip = 0x0c63
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x0c63:
		{
			m.ip = 0x0c66
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x0c66:
		{
			m.ip = 0x0c68
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3134)
			}
			m.cycles += 17
		}
	case 0x0c68:
		{
			m.ip = 0x0c6c
			m.alu("sub", 16, m.r[6], uint16(16384))
			m.cycles += 4
		}
	case 0x0c6c:
		{
			m.ip = 0x0c6e
			if !m.cf && !m.zf {
				m.ip = uint16(3190)
			}
			m.cycles += 8
		}
	case 0x0c6e:
		{
			m.ip = 0x0c71
			m.r[6] = uint16(16384)
			m.cycles += 2
		}
	case 0x0c71:
		{
			m.ip = 0x0c74
			m.r[1] = uint16(8200)
			m.cycles += 2
		}
	case 0x0c74:
		{
			m.ip = 0x0c76
			m.ip = uint16(3134)
			m.cycles += 15
		}
	case 0x0c76:
		{
			m.ip = 0x0c77
			m.r[11] = m.pop()
			m.cycles += 8
		}
	case 0x0c77:
		{
			m.ip = 0x0c7d
			m.wr16(m.r[11], uint16(0x294), m.alu("add", 16, m.rd16(m.r[11], uint16(0x294)), uint16(3632)))
			m.cycles += 16
		}
	case 0x0c7d:
		{
			m.ip = 0x0c80
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x0c80:
		{
			m.ip = 0x0c82
			m.set8(0, 0, m.alu("or", 8, ((m.r[0]>>0)&255), uint16(128)))
			m.cycles += 4
		}
	case 0x0c82:
		{
			m.ip = 0x0c84
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0c84:
		{
			m.ip = 0x0c87
			m.r[1] = uint16(8192)
			m.cycles += 2
		}
	case 0x0c87:
		{
			m.ip = 0x0c89
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x0c89:
		{
			m.ip = 0x0c8b
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0c8b:
		{
			m.ip = 0x0c8d
			m.set8(3, 0, uint16(0))
			m.cycles += 2
		}
	case 0x0c8d:
		{
			m.ip = 0x0c90
			m.r[0] = uint16(4099)
			m.cycles += 2
		}
	case 0x0c90:
		{
			m.ip = 0x0c92
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0c92:
		{
			m.ip = 0x0c95
			target := uint16(3488)
			m.push(0x0c95)
			m.ip = target
			m.cycles += 19
		}
	case 0x0c95:
		{
			m.ip = 0x0c99
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x0c99:
		{
			m.ip = 0x0c9e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2077)), uint16(1))
			m.cycles += 16
		}
	case 0x0c9e:
		{
			m.ip = 0x0ca0
			if !m.zf {
				m.ip = uint16(3241)
			}
			m.cycles += 8
		}
	case 0x0ca0:
		{
			m.ip = 0x0ca3
			target := uint16(3447)
			m.push(0x0ca3)
			m.ip = target
			m.cycles += 19
		}
	case 0x0ca3:
		{
			m.ip = 0x0ca6
			target := uint16(6660)
			m.push(0x0ca6)
			m.ip = target
			m.cycles += 19
		}
	case 0x0ca6:
		{
			m.ip = 0x0ca8
			m.ip = uint16(3254)
			m.cycles += 15
		}
	case 0x0ca8:
		{
			m.ip = 0x0ca9
			m.cycles += 3
		}
	case 0x0ca9:
		{
			m.ip = 0x0cac
			target := uint16(3676)
			m.push(0x0cac)
			m.ip = target
			m.cycles += 19
		}
	case 0x0cac:
		{
			m.ip = 0x0cb1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x74)), uint16(1))
			m.cycles += 16
		}
	case 0x0cb1:
		{
			m.ip = 0x0cb3
			if !m.zf {
				m.ip = uint16(3254)
			}
			m.cycles += 8
		}
	case 0x0cb3:
		{
			m.ip = 0x0cb6
			m.ip = uint16(3424)
			m.cycles += 15
		}
	case 0x0cb6:
		{
			m.ip = 0x0cb9
			m.r[0] = uint16(16)
			m.cycles += 2
		}
	case 0x0cb9:
		{
			m.ip = 0x0cbb
			m.set8(0, 0, m.alu("or", 8, ((m.r[0]>>0)&255), uint16(128)))
			m.cycles += 4
		}
	case 0x0cbb:
		{
			m.ip = 0x0cbd
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0cbd:
		{
			m.ip = 0x0cc1
			m.r[6] = uint16(0x1056)
			m.cycles += 3
		}
	case 0x0cc1:
		{
			m.ip = 0x0cc4
			target := uint16(18306)
			m.push(0x0cc4)
			m.ip = target
			m.cycles += 19
		}
	case 0x0cc4:
		{
			m.ip = 0x0cca
			m.wr16(m.r[11], uint16(0x294), m.alu("sub", 16, m.rd16(m.r[11], uint16(0x294)), uint16(3632)))
			m.cycles += 16
		}
	case 0x0cca:
		{
			m.ip = 0x0cce
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x0cce:
		{
			m.ip = 0x0cd1
			m.r[0] = uint16(40960)
			m.cycles += 2
		}
	case 0x0cd1:
		{
			m.ip = 0x0cd2
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x0cd2:
		{
			m.ip = 0x0cd4
			m.r[11] = m.r[0]
			m.cycles += 2
		}
	case 0x0cd4:
		{
			m.ip = 0x0cd6
			m.r[6] = m.alu("xor", 16, m.r[6], m.r[6])
			m.cycles += 4
		}
	case 0x0cd6:
		{
			m.ip = 0x0cd8
			m.r[7] = m.alu("xor", 16, m.r[7], m.r[7])
			m.cycles += 4
		}
	case 0x0cd8:
		{
			m.ip = 0x0cdb
			m.r[0] = uint16(5)
			m.cycles += 2
		}
	case 0x0cdb:
		{
			m.ip = 0x0cde
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x0cde:
		{
			m.ip = 0x0cdf
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0cdf:
		{
			m.ip = 0x0ce2
			m.r[0] = uint16(65288)
			m.cycles += 2
		}
	case 0x0ce2:
		{
			m.ip = 0x0ce3
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0ce3:
		{
			m.ip = 0x0ce6
			m.r[0] = uint16(3840)
			m.cycles += 2
		}
	case 0x0ce6:
		{
			m.ip = 0x0ce7
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0ce7:
		{
			m.ip = 0x0cea
			m.r[1] = uint16(8192)
			m.cycles += 2
		}
	case 0x0cea:
		{
			m.ip = 0x0ced
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x0ced:
		{
			m.ip = 0x0cee
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0cee:
		{
			m.ip = 0x0cf1
			m.r[0] = uint16(3841)
			m.cycles += 2
		}
	case 0x0cf1:
		{
			m.ip = 0x0cf2
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0cf2:
		{
			m.ip = 0x0cf4
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0cf4:
		{
			m.ip = 0x0cf7
			m.r[0] = uint16(2051)
			m.cycles += 2
		}
	case 0x0cf7:
		{
			m.ip = 0x0cf8
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0cf8:
		{
			m.ip = 0x0cfb
			m.r[0] = uint16(3585)
			m.cycles += 2
		}
	case 0x0cfb:
		{
			m.ip = 0x0cfc
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0cfc:
		{
			m.ip = 0x0cfe
			m.set8(0, 8, m.rd8(m.r[11], uint16(m.r[7])))
			m.cycles += 8
		}
	case 0x0cfe:
		{
			m.ip = 0x0d02
			m.r[3] = m.rd16(m.r[8], uint16(m.r[6]+0x2))
			m.cycles += 8
		}
	case 0x0d02:
		{
			m.ip = 0x0d04
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[3] >> 8) & 255))
			m.cycles += 8
		}
	case 0x0d04:
		{
			m.ip = 0x0d06
			m.set8(0, 8, uint16(13))
			m.cycles += 2
		}
	case 0x0d06:
		{
			m.ip = 0x0d07
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0d07:
		{
			m.ip = 0x0d09
			m.set8(0, 8, m.rd8(m.r[11], uint16(m.r[7])))
			m.cycles += 8
		}
	case 0x0d09:
		{
			m.ip = 0x0d0b
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0d0b:
		{
			m.ip = 0x0d0d
			m.set8(0, 8, uint16(11))
			m.cycles += 2
		}
	case 0x0d0d:
		{
			m.ip = 0x0d0e
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0d0e:
		{
			m.ip = 0x0d10
			m.set8(0, 8, m.rd8(m.r[11], uint16(m.r[7])))
			m.cycles += 8
		}
	case 0x0d10:
		{
			m.ip = 0x0d13
			m.r[3] = m.rd16(m.r[8], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x0d13:
		{
			m.ip = 0x0d15
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[3] >> 8) & 255))
			m.cycles += 8
		}
	case 0x0d15:
		{
			m.ip = 0x0d17
			m.set8(0, 8, uint16(7))
			m.cycles += 2
		}
	case 0x0d17:
		{
			m.ip = 0x0d18
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0d18:
		{
			m.ip = 0x0d1a
			m.set8(0, 8, m.rd8(m.r[11], uint16(m.r[7])))
			m.cycles += 8
		}
	case 0x0d1a:
		{
			m.ip = 0x0d1c
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0d1c:
		{
			m.ip = 0x0d1f
			m.r[6] = m.alu("add", 16, m.r[6], uint16(4))
			m.cycles += 4
		}
	case 0x0d1f:
		{
			m.ip = 0x0d20
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x0d20:
		{
			m.ip = 0x0d22
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3306)
			}
			m.cycles += 17
		}
	case 0x0d22:
		{
			m.ip = 0x0d26
			m.alu("sub", 16, m.r[7], uint16(16384))
			m.cycles += 4
		}
	case 0x0d26:
		{
			m.ip = 0x0d28
			if !m.cf && !m.zf {
				m.ip = uint16(3376)
			}
			m.cycles += 8
		}
	case 0x0d28:
		{
			m.ip = 0x0d2b
			m.r[7] = uint16(16384)
			m.cycles += 2
		}
	case 0x0d2b:
		{
			m.ip = 0x0d2e
			m.r[1] = uint16(8200)
			m.cycles += 2
		}
	case 0x0d2e:
		{
			m.ip = 0x0d30
			m.ip = uint16(3306)
			m.cycles += 15
		}
	case 0x0d30:
		{
			m.ip = 0x0d31
			m.r[11] = m.pop()
			m.cycles += 8
		}
	case 0x0d31:
		{
			m.ip = 0x0d36
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2078)), uint16(1))
			m.cycles += 16
		}
	case 0x0d36:
		{
			m.ip = 0x0d38
			if m.zf {
				m.ip = uint16(3395)
			}
			m.cycles += 8
		}
	case 0x0d38:
		{
			m.ip = 0x0d3c
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x0d3c:
		{
			m.ip = 0x0d40
			m.r[6] = uint16(0x1045)
			m.cycles += 3
		}
	case 0x0d40:
		{
			m.ip = 0x0d43
			target := uint16(18306)
			m.push(0x0d43)
			m.ip = target
			m.cycles += 19
		}
	case 0x0d43:
		{
			m.ip = 0x0d48
			m.wr8(m.r[11], uint16(0x2078), uint16(0))
			m.cycles += 8
		}
	case 0x0d48:
		{
			m.ip = 0x0d4b
			m.r[0] = uint16(261)
			m.cycles += 2
		}
	case 0x0d4b:
		{
			m.ip = 0x0d4e
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x0d4e:
		{
			m.ip = 0x0d4f
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x0d4f:
		{
			m.ip = 0x0d53
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x0d53:
		{
			m.ip = 0x0d58
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2076)), uint16(1))
			m.cycles += 16
		}
	case 0x0d58:
		{
			m.ip = 0x0d5a
			if !m.zf {
				m.ip = uint16(3424)
			}
			m.cycles += 8
		}
	case 0x0d5a:
		{
			m.ip = 0x0d5d
			target := uint16(18893)
			m.push(0x0d5d)
			m.ip = target
			m.cycles += 19
		}
	case 0x0d5d:
		{
			m.ip = 0x0d60
			m.ip = uint16(3114)
			m.cycles += 15
		}
	case 0x0d60:
		{
			m.ip = 0x0d63
			target := uint16(8099)
			m.push(0x0d63)
			m.ip = target
			m.cycles += 19
		}
	case 0x0d63:
		{
			m.ip = 0x0d68
			m.wr8(m.r[11], uint16(0x2074), uint16(0))
			m.cycles += 8
		}
	case 0x0d68:
		{
			m.ip = 0x0d6d
			m.wr8(m.r[11], uint16(0x2077), uint16(0))
			m.cycles += 8
		}
	case 0x0d6d:
		{
			m.ip = 0x0d6e
			m.r[11] = m.pop()
			m.cycles += 8
		}
	case 0x0d6e:
		{
			m.ip = 0x0d6f
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x0d6f:
		{
			m.ip = 0x0d70
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x0d70:
		{
			m.ip = 0x0d71
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x0d71:
		{
			m.ip = 0x0d72
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x0d72:
		{
			m.ip = 0x0d73
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x0d73:
		{
			m.ip = 0x0d74
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x0d74:
		{
			m.ip = 0x0d75
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x0d75:
		{
			m.ip = 0x0d76
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x0d76:
		{
			m.ip = 0x0d77
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0d77:
		{
			m.ip = 0x0d78
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x0d78:
		{
			m.ip = 0x0d79
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x0d79:
		{
			m.ip = 0x0d7a
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x0d7a:
		{
			m.ip = 0x0d7e
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x0d7e:
		{
			m.ip = 0x0d80
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x0d80:
		{
			m.ip = 0x0d82
			m.set8(0, 0, uint16(32))
			m.cycles += 2
		}
	case 0x0d82:
		{
			m.ip = 0x0d84
			m.r[7] = m.alu("xor", 16, m.r[7], m.r[7])
			m.cycles += 4
		}
	case 0x0d84:
		{
			m.ip = 0x0d87
			m.r[1] = uint16(2000)
			m.cycles += 2
		}
	case 0x0d87:
		{
			m.ip = 0x0d8c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2077)), uint16(1))
			m.cycles += 16
		}
	case 0x0d8c:
		{
			m.ip = 0x0d8e
			if !m.zf {
				m.ip = uint16(3476)
			}
			m.cycles += 8
		}
	case 0x0d8e:
		{
			m.ip = 0x0d91
			m.r[7] = uint16(800)
			m.cycles += 2
		}
	case 0x0d91:
		{
			m.ip = 0x0d94
			m.r[1] = uint16(1600)
			m.cycles += 2
		}
	case 0x0d94:
		{
			m.ip = 0x0d97
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x0d97:
		{
			m.ip = 0x0d9a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x0d9a:
		{
			m.ip = 0x0d9c
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3476)
			}
			m.cycles += 17
		}
	case 0x0d9c:
		{
			m.ip = 0x0d9d
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x0d9d:
		{
			m.ip = 0x0d9e
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x0d9e:
		{
			m.ip = 0x0d9f
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x0d9f:
		{
			m.ip = 0x0da0
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0da0:
		{
			m.ip = 0x0da3
			m.r[0] = uint16(4368)
			m.cycles += 2
		}
	case 0x0da3:
		{
			m.ip = 0x0da5
			m.set8(3, 8, uint16(14))
			m.cycles += 2
		}
	case 0x0da5:
		{
			m.ip = 0x0da7
			m.set8(3, 0, uint16(1))
			m.cycles += 2
		}
	case 0x0da7:
		{
			m.ip = 0x0daa
			m.r[1] = uint16(254)
			m.cycles += 2
		}
	case 0x0daa:
		{
			m.ip = 0x0dad
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x0dad:
		{
			m.ip = 0x0dae
			m.push(m.r[8])
			m.cycles += 11
		}
	case 0x0dae:
		{
			m.ip = 0x0daf
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x0daf:
		{
			m.ip = 0x0db0
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x0db0:
		{
			m.ip = 0x0db4
			m.r[5] = uint16(0x254d)
			m.cycles += 3
		}
	case 0x0db4:
		{
			m.ip = 0x0db6
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0db6:
		{
			m.ip = 0x0db7
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x0db7:
		{
			m.ip = 0x0dba
			m.r[0] = uint16(4609)
			m.cycles += 2
		}
	case 0x0dba:
		{
			m.ip = 0x0dbc
			m.set8(3, 0, uint16(48))
			m.cycles += 2
		}
	case 0x0dbc:
		{
			m.ip = 0x0dbe
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0dbe:
		{
			m.ip = 0x0dc2
			m.r[6] = uint16(0x1034)
			m.cycles += 3
		}
	case 0x0dc2:
		{
			m.ip = 0x0dc5
			target := uint16(18306)
			m.push(0x0dc5)
			m.ip = target
			m.cycles += 19
		}
	case 0x0dc5:
		{
			m.ip = 0x0dc8
			m.r[0] = uint16(4355)
			m.cycles += 2
		}
	case 0x0dc8:
		{
			m.ip = 0x0dca
			m.set8(3, 0, uint16(4))
			m.cycles += 2
		}
	case 0x0dca:
		{
			m.ip = 0x0dcc
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x0dcc:
		{
			m.ip = 0x0dcd
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0dcd:
		{
			m.ip = 0x0dce
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x0dce:
		{
			m.ip = 0x0dcf
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x0dcf:
		{
			m.ip = 0x0dd3
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x0dd3:
		{
			m.ip = 0x0dd8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2077)), uint16(1))
			m.cycles += 16
		}
	case 0x0dd8:
		{
			m.ip = 0x0dda
			if m.zf {
				m.ip = uint16(3563)
			}
			m.cycles += 8
		}
	case 0x0dda:
		{
			m.ip = 0x0ddd
			m.r[0] = uint16(32)
			m.cycles += 2
		}
	case 0x0ddd:
		{
			m.ip = 0x0de0
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x0de0:
		{
			m.ip = 0x0de3
			m.r[1] = uint16(400)
			m.cycles += 2
		}
	case 0x0de3:
		{
			m.ip = 0x0de6
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x0de6:
		{
			m.ip = 0x0de9
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x0de9:
		{
			m.ip = 0x0deb
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3555)
			}
			m.cycles += 17
		}
	case 0x0deb:
		{
			m.ip = 0x0dee
			m.r[7] = uint16(10)
			m.cycles += 2
		}
	case 0x0dee:
		{
			m.ip = 0x0df1
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x0df1:
		{
			m.ip = 0x0df3
			m.set8(0, 0, uint16(15))
			m.cycles += 2
		}
	case 0x0df3:
		{
			m.ip = 0x0df6
			m.r[1] = uint16(4)
			m.cycles += 2
		}
	case 0x0df6:
		{
			m.ip = 0x0df7
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x0df7:
		{
			m.ip = 0x0dfb
			m.set8(0, 8, m.rd8(m.r[11], uint16(m.r[3]+0x2549)))
			m.cycles += 8
		}
	case 0x0dfb:
		{
			m.ip = 0x0dfc
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x0dfc:
		{
			m.ip = 0x0dff
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x0dff:
		{
			m.ip = 0x0e02
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x0e02:
		{
			m.ip = 0x0e05
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x0e05:
		{
			m.ip = 0x0e07
			m.set8(0, 0, m.unary("inc", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x0e07:
		{
			m.ip = 0x0e09
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3583)
			}
			m.cycles += 17
		}
	case 0x0e09:
		{
			m.ip = 0x0e0d
			m.r[7] = m.alu("add", 16, m.r[7], uint16(132))
			m.cycles += 4
		}
	case 0x0e0d:
		{
			m.ip = 0x0e0e
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x0e0e:
		{
			m.ip = 0x0e10
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3574)
			}
			m.cycles += 17
		}
	case 0x0e10:
		{
			m.ip = 0x0e13
			m.r[7] = uint16(46)
			m.cycles += 2
		}
	case 0x0e13:
		{
			m.ip = 0x0e16
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x0e16:
		{
			m.ip = 0x0e18
			m.set8(0, 0, uint16(122))
			m.cycles += 2
		}
	case 0x0e18:
		{
			m.ip = 0x0e1b
			m.r[1] = uint16(4)
			m.cycles += 2
		}
	case 0x0e1b:
		{
			m.ip = 0x0e1c
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x0e1c:
		{
			m.ip = 0x0e20
			m.set8(0, 8, m.rd8(m.r[11], uint16(m.r[3]+0x2545)))
			m.cycles += 8
		}
	case 0x0e20:
		{
			m.ip = 0x0e21
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x0e21:
		{
			m.ip = 0x0e24
			m.r[1] = uint16(33)
			m.cycles += 2
		}
	case 0x0e24:
		{
			m.ip = 0x0e27
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x0e27:
		{
			m.ip = 0x0e2a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x0e2a:
		{
			m.ip = 0x0e2c
			m.set8(0, 0, m.unary("inc", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x0e2c:
		{
			m.ip = 0x0e2e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3620)
			}
			m.cycles += 17
		}
	case 0x0e2e:
		{
			m.ip = 0x0e31
			m.r[7] = m.alu("add", 16, m.r[7], uint16(94))
			m.cycles += 4
		}
	case 0x0e31:
		{
			m.ip = 0x0e32
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x0e32:
		{
			m.ip = 0x0e34
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3611)
			}
			m.cycles += 17
		}
	case 0x0e34:
		{
			m.ip = 0x0e37
			m.r[7] = uint16(120)
			m.cycles += 2
		}
	case 0x0e37:
		{
			m.ip = 0x0e3a
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x0e3a:
		{
			m.ip = 0x0e3c
			m.set8(0, 0, uint16(15))
			m.cycles += 2
		}
	case 0x0e3c:
		{
			m.ip = 0x0e3f
			m.r[1] = uint16(4)
			m.cycles += 2
		}
	case 0x0e3f:
		{
			m.ip = 0x0e40
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x0e40:
		{
			m.ip = 0x0e44
			m.set8(0, 8, m.rd8(m.r[11], uint16(m.r[3]+0x2549)))
			m.cycles += 8
		}
	case 0x0e44:
		{
			m.ip = 0x0e45
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x0e45:
		{
			m.ip = 0x0e48
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x0e48:
		{
			m.ip = 0x0e4b
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x0e4b:
		{
			m.ip = 0x0e4e
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x0e4e:
		{
			m.ip = 0x0e50
			m.set8(0, 0, m.unary("inc", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x0e50:
		{
			m.ip = 0x0e52
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3656)
			}
			m.cycles += 17
		}
	case 0x0e52:
		{
			m.ip = 0x0e56
			m.r[7] = m.alu("add", 16, m.r[7], uint16(132))
			m.cycles += 4
		}
	case 0x0e56:
		{
			m.ip = 0x0e57
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x0e57:
		{
			m.ip = 0x0e59
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3647)
			}
			m.cycles += 17
		}
	case 0x0e59:
		{
			m.ip = 0x0e5a
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x0e5a:
		{
			m.ip = 0x0e5b
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x0e5b:
		{
			m.ip = 0x0e5c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x0e5c:
		{
			m.ip = 0x0e61
			m.wr8(m.r[11], uint16(0x74), uint16(0))
			m.cycles += 8
		}
	case 0x0e61:
		{
			m.ip = 0x0e66
			m.wr8(m.r[11], uint16(0x2075), uint16(0))
			m.cycles += 8
		}
	case 0x0e66:
		{
			m.ip = 0x0e6b
			m.wr8(m.r[11], uint16(0x2076), uint16(0))
			m.cycles += 8
		}
	case 0x0e6b:
		{
			m.ip = 0x0e70
			m.wr8(m.r[11], uint16(0x2077), uint16(0))
			m.cycles += 8
		}
	case 0x0e70:
		{
			m.ip = 0x0e73
			target := uint16(3447)
			m.push(0x0e73)
			m.ip = target
			m.cycles += 19
		}
	case 0x0e73:
		{
			m.ip = 0x0e76
			target := uint16(3533)
			m.push(0x0e76)
			m.ip = target
			m.cycles += 19
		}
	case 0x0e76:
		{
			m.ip = 0x0e7c
			m.wr16(m.r[11], uint16(0x2091), uint16(320))
			m.cycles += 8
		}
	case 0x0e7c:
		{
			m.ip = 0x0e82
			m.wr16(m.r[11], uint16(0x20af), uint16(0))
			m.cycles += 8
		}
	case 0x0e82:
		{
			m.ip = 0x0e87
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x615)), uint16(1))
			m.cycles += 16
		}
	case 0x0e87:
		{
			m.ip = 0x0e89
			if !m.zf {
				m.ip = uint16(3733)
			}
			m.cycles += 8
		}
	case 0x0e89:
		{
			m.ip = 0x0e8f
			m.wr16(m.r[11], uint16(0x2091), uint16(640))
			m.cycles += 8
		}
	case 0x0e8f:
		{
			m.ip = 0x0e95
			m.wr16(m.r[11], uint16(0x20af), uint16(4))
			m.cycles += 8
		}
	case 0x0e95:
		{
			m.ip = 0x0e98
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x0e98:
		{
			m.ip = 0x0e9b
			m.r[1] = uint16(5)
			m.cycles += 2
		}
	case 0x0e9b:
		{
			m.ip = 0x0e9c
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x0e9c:
		{
			m.ip = 0x0e9e
			m.set8(0, 8, uint16(7))
			m.cycles += 2
		}
	case 0x0e9e:
		{
			m.ip = 0x0ea1
			m.alu("sub", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x0ea1:
		{
			m.ip = 0x0ea3
			if !m.zf {
				m.ip = uint16(3749)
			}
			m.cycles += 8
		}
	case 0x0ea3:
		{
			m.ip = 0x0ea5
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x0ea5:
		{
			m.ip = 0x0ea8
			target := uint16(5203)
			m.push(0x0ea8)
			m.ip = target
			m.cycles += 19
		}
	case 0x0ea8:
		{
			m.ip = 0x0eac
			m.r[6] = m.rd16(m.r[11], uint16(m.r[3]+0x2079))
			m.cycles += 8
		}
	case 0x0eac:
		{
			m.ip = 0x0eb0
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x2083))
			m.cycles += 8
		}
	case 0x0eb0:
		{
			m.ip = 0x0eb3
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x0eb3:
		{
			m.ip = 0x0eb7
			m.r[1] = m.rd16(m.r[11], uint16(m.r[3]+0x20a1))
			m.cycles += 8
		}
	case 0x0eb7:
		{
			m.ip = 0x0eb8
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x0eb8:
		{
			m.ip = 0x0eba
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x0eba:
		{
			m.ip = 0x0ebd
			m.alu("sub", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x0ebd:
		{
			m.ip = 0x0ebf
			if !m.zf {
				m.ip = uint16(3777)
			}
			m.cycles += 8
		}
	case 0x0ebf:
		{
			m.ip = 0x0ec1
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x0ec1:
		{
			m.ip = 0x0ec4
			m.alu("sub", 16, m.r[2], uint16(0))
			m.cycles += 4
		}
	case 0x0ec4:
		{
			m.ip = 0x0ec6
			if m.zf {
				m.ip = uint16(3790)
			}
			m.cycles += 8
		}
	case 0x0ec6:
		{
			m.ip = 0x0eca
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3]+0x208d)))
			m.cycles += 16
		}
	case 0x0eca:
		{
			m.ip = 0x0ecc
			if m.zf {
				m.ip = uint16(3790)
			}
			m.cycles += 8
		}
	case 0x0ecc:
		{
			m.ip = 0x0ece
			m.set8(0, 8, uint16(7))
			m.cycles += 2
		}
	case 0x0ece:
		{
			m.ip = 0x0ed1
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x0ed1:
		{
			m.ip = 0x0ed3
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x0ed3:
		{
			m.ip = 0x0ed6
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x0ed6:
		{
			m.ip = 0x0ed7
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x0ed7:
		{
			m.ip = 0x0eda
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x0eda:
		{
			m.ip = 0x0edc
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3793)
			}
			m.cycles += 17
		}
	case 0x0edc:
		{
			m.ip = 0x0ee0
			m.r[7] = m.alu("add", 16, m.r[7], uint16(300))
			m.cycles += 4
		}
	case 0x0ee0:
		{
			m.ip = 0x0ee4
			m.r[2] = m.alu("add", 16, m.r[2], uint16(320))
			m.cycles += 4
		}
	case 0x0ee4:
		{
			m.ip = 0x0ee5
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x0ee5:
		{
			m.ip = 0x0ee7
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3767)
			}
			m.cycles += 17
		}
	case 0x0ee7:
		{
			m.ip = 0x0eea
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x0eea:
		{
			m.ip = 0x0eeb
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x0eeb:
		{
			m.ip = 0x0eed
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3739)
			}
			m.cycles += 17
		}
	case 0x0eed:
		{
			m.ip = 0x0eef
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x0eef:
		{
			m.ip = 0x0ef2
			target := uint16(4519)
			m.push(0x0ef2)
			m.ip = target
			m.cycles += 19
		}
	case 0x0ef2:
		{
			m.ip = 0x0ef5
			m.r[3] = uint16(4)
			m.cycles += 2
		}
	case 0x0ef5:
		{
			m.ip = 0x0efa
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(1))
			m.cycles += 16
		}
	case 0x0efa:
		{
			m.ip = 0x0efc
			if !m.zf {
				m.ip = uint16(3839)
			}
			m.cycles += 8
		}
	case 0x0efc:
		{
			m.ip = 0x0eff
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x0eff:
		{
			m.ip = 0x0f02
			target := uint16(4761)
			m.push(0x0f02)
			m.ip = target
			m.cycles += 19
		}
	case 0x0f02:
		{
			m.ip = 0x0f05
			target := uint16(18893)
			m.push(0x0f05)
			m.ip = target
			m.cycles += 19
		}
	case 0x0f05:
		{
			m.ip = 0x0f07
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x0f07:
		{
			m.ip = 0x0f09
			if m.zf {
				m.ip = uint16(3852)
			}
			m.cycles += 8
		}
	case 0x0f09:
		{
			m.ip = 0x0f0c
			m.ip = uint16(4268)
			m.cycles += 15
		}
	case 0x0f0c:
		{
			m.ip = 0x0f0f
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(72))
			m.cycles += 4
		}
	case 0x0f0f:
		{
			m.ip = 0x0f11
			if !m.zf {
				m.ip = uint16(3871)
			}
			m.cycles += 8
		}
	case 0x0f11:
		{
			m.ip = 0x0f14
			m.alu("sub", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x0f14:
		{
			m.ip = 0x0f16
			if m.zf {
				m.ip = uint16(3874)
			}
			m.cycles += 8
		}
	case 0x0f16:
		{
			m.ip = 0x0f18
			m.set8(2, 0, uint16(7))
			m.cycles += 2
		}
	case 0x0f18:
		{
			m.ip = 0x0f1c
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x208d))
			m.cycles += 8
		}
	case 0x0f1c:
		{
			m.ip = 0x0f1e
			m.ip = uint16(3945)
			m.cycles += 15
		}
	case 0x0f1e:
		{
			m.ip = 0x0f1f
			m.cycles += 3
		}
	case 0x0f1f:
		{
			m.ip = 0x0f22
			m.ip = uint16(4008)
			m.cycles += 15
		}
	case 0x0f22:
		{
			m.ip = 0x0f27
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f6)), uint16(14))
			m.cycles += 16
		}
	case 0x0f27:
		{
			m.ip = 0x0f29
			if !m.cf {
				m.ip = uint16(3900)
			}
			m.cycles += 8
		}
	case 0x0f29:
		{
			m.ip = 0x0f2d
			m.wr8(m.r[11], uint16(0x24f6), m.unary("inc", 8, m.rd8(m.r[11], uint16(0x24f6))))
			m.cycles += 15
		}
	case 0x0f2d:
		{
			m.ip = 0x0f30
			m.r[0] = m.rd16(m.r[11], uint16(0x106e))
			m.cycles += 8
		}
	case 0x0f30:
		{
			m.ip = 0x0f34
			m.wr16(m.r[11], uint16(0x106c), m.alu("sub", 16, m.rd16(m.r[11], uint16(0x106c)), m.r[0]))
			m.cycles += 16
		}
	case 0x0f34:
		{
			m.ip = 0x0f36
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x0f36:
		{
			m.ip = 0x0f39
			target := uint16(4519)
			m.push(0x0f39)
			m.ip = target
			m.cycles += 19
		}
	case 0x0f39:
		{
			m.ip = 0x0f3b
			m.ip = uint16(3928)
			m.cycles += 15
		}
	case 0x0f3b:
		{
			m.ip = 0x0f3c
			m.cycles += 3
		}
	case 0x0f3c:
		{
			m.ip = 0x0f41
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(7))
			m.cycles += 16
		}
	case 0x0f41:
		{
			m.ip = 0x0f43
			if m.cf || m.zf {
				m.ip = uint16(3842)
			}
			m.cycles += 8
		}
	case 0x0f43:
		{
			m.ip = 0x0f48
			m.wr8(m.r[11], uint16(0x24f6), uint16(1))
			m.cycles += 8
		}
	case 0x0f48:
		{
			m.ip = 0x0f4c
			m.wr16(m.r[11], uint16(0x24f2), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x24f2))))
			m.cycles += 15
		}
	case 0x0f4c:
		{
			m.ip = 0x0f4f
			m.r[0] = m.rd16(m.r[11], uint16(0x106e))
			m.cycles += 8
		}
	case 0x0f4f:
		{
			m.ip = 0x0f53
			m.wr16(m.r[11], uint16(0x106c), m.alu("sub", 16, m.rd16(m.r[11], uint16(0x106c)), m.r[0]))
			m.cycles += 16
		}
	case 0x0f53:
		{
			m.ip = 0x0f55
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x0f55:
		{
			m.ip = 0x0f58
			target := uint16(4519)
			m.push(0x0f58)
			m.ip = target
			m.cycles += 19
		}
	case 0x0f58:
		{
			m.ip = 0x0f5e
			m.alu("and", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(1))
			m.cycles += 16
		}
	case 0x0f5e:
		{
			m.ip = 0x0f60
			if !m.zf {
				m.ip = uint16(3943)
			}
			m.cycles += 8
		}
	case 0x0f60:
		{
			m.ip = 0x0f65
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f6)), uint16(8))
			m.cycles += 16
		}
	case 0x0f65:
		{
			m.ip = 0x0f67
			if m.zf {
				m.ip = uint16(3862)
			}
			m.cycles += 8
		}
	case 0x0f67:
		{
			m.ip = 0x0f69
			m.ip = uint16(3842)
			m.cycles += 15
		}
	case 0x0f69:
		{
			m.ip = 0x0f6d
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x2083))
			m.cycles += 8
		}
	case 0x0f6d:
		{
			m.ip = 0x0f6f
			m.r[7] = m.alu("add", 16, m.r[7], m.r[0])
			m.cycles += 4
		}
	case 0x0f6f:
		{
			m.ip = 0x0f72
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x0f72:
		{
			m.ip = 0x0f75
			m.alu("sub", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x0f75:
		{
			m.ip = 0x0f77
			if !m.zf {
				m.ip = uint16(3962)
			}
			m.cycles += 8
		}
	case 0x0f77:
		{
			m.ip = 0x0f7a
			m.r[1] = uint16(9)
			m.cycles += 2
		}
	case 0x0f7a:
		{
			m.ip = 0x0f7e
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[2] >> 0) & 255))
			m.cycles += 8
		}
	case 0x0f7e:
		{
			m.ip = 0x0f81
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x0f81:
		{
			m.ip = 0x0f83
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(3962)
			}
			m.cycles += 17
		}
	case 0x0f83:
		{
			m.ip = 0x0f86
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(1))
			m.cycles += 4
		}
	case 0x0f86:
		{
			m.ip = 0x0f88
			if !m.zf {
				m.ip = uint16(3979)
			}
			m.cycles += 8
		}
	case 0x0f88:
		{
			m.ip = 0x0f8b
			m.ip = uint16(3839)
			m.cycles += 15
		}
	case 0x0f8b:
		{
			m.ip = 0x0f8e
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(320))
			m.cycles += 4
		}
	case 0x0f8e:
		{
			m.ip = 0x0f93
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x20ab)), uint16(4)))
			m.cycles += 16
		}
	case 0x0f93:
		{
			m.ip = 0x0f96
			m.alu("sub", 16, m.r[0], uint16(320))
			m.cycles += 4
		}
	case 0x0f96:
		{
			m.ip = 0x0f98
			if m.sf == m.of {
				m.ip = uint16(4000)
			}
			m.cycles += 8
		}
	case 0x0f98:
		{
			m.ip = 0x0f9d
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x20ab)), uint16(4)))
			m.cycles += 16
		}
	case 0x0f9d:
		{
			m.ip = 0x0fa0
			m.r[0] = m.alu("add", 16, m.r[0], uint16(320))
			m.cycles += 4
		}
	case 0x0fa0:
		{
			m.ip = 0x0fa4
			m.wr16(m.r[11], uint16(m.r[3]+0x208d), m.r[0])
			m.cycles += 8
		}
	case 0x0fa4:
		{
			m.ip = 0x0fa6
			m.set8(2, 0, uint16(1))
			m.cycles += 2
		}
	case 0x0fa6:
		{
			m.ip = 0x0fa8
			m.ip = uint16(3945)
			m.cycles += 15
		}
	case 0x0fa8:
		{
			m.ip = 0x0fab
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(80))
			m.cycles += 4
		}
	case 0x0fab:
		{
			m.ip = 0x0fad
			if !m.zf {
				m.ip = uint16(4027)
			}
			m.cycles += 8
		}
	case 0x0fad:
		{
			m.ip = 0x0fb0
			m.alu("sub", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x0fb0:
		{
			m.ip = 0x0fb2
			if m.zf {
				m.ip = uint16(4030)
			}
			m.cycles += 8
		}
	case 0x0fb2:
		{
			m.ip = 0x0fb4
			m.set8(2, 0, uint16(7))
			m.cycles += 2
		}
	case 0x0fb4:
		{
			m.ip = 0x0fb8
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x208d))
			m.cycles += 8
		}
	case 0x0fb8:
		{
			m.ip = 0x0fba
			m.ip = uint16(4102)
			m.cycles += 15
		}
	case 0x0fba:
		{
			m.ip = 0x0fbb
			m.cycles += 3
		}
	case 0x0fbb:
		{
			m.ip = 0x0fbe
			m.ip = uint16(4158)
			m.cycles += 15
		}
	case 0x0fbe:
		{
			m.ip = 0x0fc3
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f6)), uint16(1))
			m.cycles += 16
		}
	case 0x0fc3:
		{
			m.ip = 0x0fc5
			if m.cf || m.zf {
				m.ip = uint16(4056)
			}
			m.cycles += 8
		}
	case 0x0fc5:
		{
			m.ip = 0x0fc9
			m.wr8(m.r[11], uint16(0x24f6), m.unary("dec", 8, m.rd8(m.r[11], uint16(0x24f6))))
			m.cycles += 15
		}
	case 0x0fc9:
		{
			m.ip = 0x0fcc
			m.r[0] = m.rd16(m.r[11], uint16(0x106e))
			m.cycles += 8
		}
	case 0x0fcc:
		{
			m.ip = 0x0fd0
			m.wr16(m.r[11], uint16(0x106c), m.alu("add", 16, m.rd16(m.r[11], uint16(0x106c)), m.r[0]))
			m.cycles += 16
		}
	case 0x0fd0:
		{
			m.ip = 0x0fd2
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x0fd2:
		{
			m.ip = 0x0fd5
			target := uint16(4519)
			m.push(0x0fd5)
			m.ip = target
			m.cycles += 19
		}
	case 0x0fd5:
		{
			m.ip = 0x0fd7
			m.ip = uint16(4084)
			m.cycles += 15
		}
	case 0x0fd7:
		{
			m.ip = 0x0fd8
			m.cycles += 3
		}
	case 0x0fd8:
		{
			m.ip = 0x0fdd
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(15))
			m.cycles += 16
		}
	case 0x0fdd:
		{
			m.ip = 0x0fdf
			if !m.cf {
				m.ip = uint16(4099)
			}
			m.cycles += 8
		}
	case 0x0fdf:
		{
			m.ip = 0x0fe4
			m.wr8(m.r[11], uint16(0x24f6), uint16(14))
			m.cycles += 8
		}
	case 0x0fe4:
		{
			m.ip = 0x0fe7
			m.r[0] = m.rd16(m.r[11], uint16(0x106e))
			m.cycles += 8
		}
	case 0x0fe7:
		{
			m.ip = 0x0feb
			m.wr16(m.r[11], uint16(0x106c), m.alu("add", 16, m.rd16(m.r[11], uint16(0x106c)), m.r[0]))
			m.cycles += 16
		}
	case 0x0feb:
		{
			m.ip = 0x0fef
			m.wr16(m.r[11], uint16(0x24f2), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x24f2))))
			m.cycles += 15
		}
	case 0x0fef:
		{
			m.ip = 0x0ff1
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x0ff1:
		{
			m.ip = 0x0ff4
			target := uint16(4519)
			m.push(0x0ff4)
			m.ip = target
			m.cycles += 19
		}
	case 0x0ff4:
		{
			m.ip = 0x0ffa
			m.alu("and", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(1))
			m.cycles += 16
		}
	case 0x0ffa:
		{
			m.ip = 0x0ffc
			if !m.zf {
				m.ip = uint16(4099)
			}
			m.cycles += 8
		}
	case 0x0ffc:
		{
			m.ip = 0x1001
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f6)), uint16(7))
			m.cycles += 16
		}
	case 0x1001:
		{
			m.ip = 0x1003
			if m.zf {
				m.ip = uint16(4018)
			}
			m.cycles += 8
		}
	case 0x1003:
		{
			m.ip = 0x1006
			m.ip = uint16(3842)
			m.cycles += 15
		}
	case 0x1006:
		{
			m.ip = 0x100a
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x2083))
			m.cycles += 8
		}
	case 0x100a:
		{
			m.ip = 0x100c
			m.r[7] = m.alu("add", 16, m.r[7], m.r[0])
			m.cycles += 4
		}
	case 0x100c:
		{
			m.ip = 0x100f
			m.r[1] = uint16(9)
			m.cycles += 2
		}
	case 0x100f:
		{
			m.ip = 0x1013
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[2] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1013:
		{
			m.ip = 0x1016
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1016:
		{
			m.ip = 0x1018
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4111)
			}
			m.cycles += 17
		}
	case 0x1018:
		{
			m.ip = 0x101b
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(1))
			m.cycles += 4
		}
	case 0x101b:
		{
			m.ip = 0x101d
			if !m.zf {
				m.ip = uint16(4128)
			}
			m.cycles += 8
		}
	case 0x101d:
		{
			m.ip = 0x1020
			m.ip = uint16(3839)
			m.cycles += 15
		}
	case 0x1020:
		{
			m.ip = 0x1023
			m.r[0] = m.alu("add", 16, m.r[0], uint16(320))
			m.cycles += 4
		}
	case 0x1023:
		{
			m.ip = 0x1028
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x20ab)), uint16(4)))
			m.cycles += 16
		}
	case 0x1028:
		{
			m.ip = 0x102c
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(m.r[3]+0x2097)))
			m.cycles += 16
		}
	case 0x102c:
		{
			m.ip = 0x102e
			if m.zf || m.sf != m.of {
				m.ip = uint16(4150)
			}
			m.cycles += 8
		}
	case 0x102e:
		{
			m.ip = 0x1031
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(320))
			m.cycles += 4
		}
	case 0x1031:
		{
			m.ip = 0x1036
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x20ab)), uint16(4)))
			m.cycles += 16
		}
	case 0x1036:
		{
			m.ip = 0x103a
			m.wr16(m.r[11], uint16(m.r[3]+0x208d), m.r[0])
			m.cycles += 8
		}
	case 0x103a:
		{
			m.ip = 0x103c
			m.set8(2, 0, uint16(1))
			m.cycles += 2
		}
	case 0x103c:
		{
			m.ip = 0x103e
			m.ip = uint16(4102)
			m.cycles += 15
		}
	case 0x103e:
		{
			m.ip = 0x1041
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(75))
			m.cycles += 4
		}
	case 0x1041:
		{
			m.ip = 0x1043
			if !m.zf {
				m.ip = uint16(4177)
			}
			m.cycles += 8
		}
	case 0x1043:
		{
			m.ip = 0x1046
			m.alu("sub", 16, m.r[3], uint16(0))
			m.cycles += 4
		}
	case 0x1046:
		{
			m.ip = 0x1048
			if m.zf {
				m.ip = uint16(4265)
			}
			m.cycles += 8
		}
	case 0x1048:
		{
			m.ip = 0x104e
			m.wr16(m.r[11], uint16(0x297), uint16(65534))
			m.cycles += 8
		}
	case 0x104e:
		{
			m.ip = 0x1050
			m.ip = uint16(4193)
			m.cycles += 15
		}
	case 0x1050:
		{
			m.ip = 0x1051
			m.cycles += 3
		}
	case 0x1051:
		{
			m.ip = 0x1054
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(77))
			m.cycles += 4
		}
	case 0x1054:
		{
			m.ip = 0x1056
			if !m.zf {
				m.ip = uint16(4265)
			}
			m.cycles += 8
		}
	case 0x1056:
		{
			m.ip = 0x1059
			m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x1059:
		{
			m.ip = 0x105b
			if m.zf {
				m.ip = uint16(4265)
			}
			m.cycles += 8
		}
	case 0x105b:
		{
			m.ip = 0x1061
			m.wr16(m.r[11], uint16(0x297), uint16(2))
			m.cycles += 8
		}
	case 0x1061:
		{
			m.ip = 0x1063
			m.set8(2, 0, uint16(4))
			m.cycles += 2
		}
	case 0x1063:
		{
			m.ip = 0x1065
			m.set8(0, 8, uint16(7))
			m.cycles += 2
		}
	case 0x1065:
		{
			m.ip = 0x1068
			target := uint16(5203)
			m.push(0x1068)
			m.ip = target
			m.cycles += 19
		}
	case 0x1068:
		{
			m.ip = 0x106b
			m.alu("sub", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x106b:
		{
			m.ip = 0x106d
			if !m.zf {
				m.ip = uint16(4210)
			}
			m.cycles += 8
		}
	case 0x106d:
		{
			m.ip = 0x106f
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x106f:
		{
			m.ip = 0x1072
			target := uint16(4519)
			m.push(0x1072)
			m.ip = target
			m.cycles += 19
		}
	case 0x1072:
		{
			m.ip = 0x1076
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x2083))
			m.cycles += 8
		}
	case 0x1076:
		{
			m.ip = 0x107a
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x208d))
			m.cycles += 8
		}
	case 0x107a:
		{
			m.ip = 0x107d
			m.r[1] = uint16(9)
			m.cycles += 2
		}
	case 0x107d:
		{
			m.ip = 0x1081
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[2] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1081:
		{
			m.ip = 0x1085
			m.wr8(m.r[8], uint16(m.r[5]+m.r[7]+0x1), ((m.r[2] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1085:
		{
			m.ip = 0x1088
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1088:
		{
			m.ip = 0x108a
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4221)
			}
			m.cycles += 17
		}
	case 0x108a:
		{
			m.ip = 0x108d
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(1))
			m.cycles += 4
		}
	case 0x108d:
		{
			m.ip = 0x108f
			if !m.zf {
				m.ip = uint16(4242)
			}
			m.cycles += 8
		}
	case 0x108f:
		{
			m.ip = 0x1092
			m.ip = uint16(3839)
			m.cycles += 15
		}
	case 0x1092:
		{
			m.ip = 0x1094
			m.set8(2, 0, uint16(1))
			m.cycles += 2
		}
	case 0x1094:
		{
			m.ip = 0x1098
			m.r[3] = m.alu("add", 16, m.r[3], m.rd16(m.r[11], uint16(0x297)))
			m.cycles += 16
		}
	case 0x1098:
		{
			m.ip = 0x109a
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x109a:
		{
			m.ip = 0x109d
			target := uint16(5203)
			m.push(0x109d)
			m.ip = target
			m.cycles += 19
		}
	case 0x109d:
		{
			m.ip = 0x10a0
			m.alu("sub", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x10a0:
		{
			m.ip = 0x10a2
			if !m.zf {
				m.ip = uint16(4210)
			}
			m.cycles += 8
		}
	case 0x10a2:
		{
			m.ip = 0x10a4
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x10a4:
		{
			m.ip = 0x10a7
			target := uint16(4519)
			m.push(0x10a7)
			m.ip = target
			m.cycles += 19
		}
	case 0x10a7:
		{
			m.ip = 0x10a9
			m.ip = uint16(4210)
			m.cycles += 15
		}
	case 0x10a9:
		{
			m.ip = 0x10ac
			m.ip = uint16(3839)
			m.cycles += 15
		}
	case 0x10ac:
		{
			m.ip = 0x10ae
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x10ae:
		{
			m.ip = 0x10b0
			if !m.zf {
				m.ip = uint16(4320)
			}
			m.cycles += 8
		}
	case 0x10b0:
		{
			m.ip = 0x10b3
			target := uint16(4929)
			m.push(0x10b3)
			m.ip = target
			m.cycles += 19
		}
	case 0x10b3:
		{
			m.ip = 0x10b8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x3b2)), uint16(1))
			m.cycles += 16
		}
	case 0x10b8:
		{
			m.ip = 0x10ba
			if !m.zf {
				m.ip = uint16(4295)
			}
			m.cycles += 8
		}
	case 0x10ba:
		{
			m.ip = 0x10bf
			m.wr8(m.r[11], uint16(0x3b2), uint16(0))
			m.cycles += 8
		}
	case 0x10bf:
		{
			m.ip = 0x10c4
			m.wr8(m.r[11], uint16(0x2077), uint16(0))
			m.cycles += 8
		}
	case 0x10c4:
		{
			m.ip = 0x10c7
			m.ip = uint16(3842)
			m.cycles += 15
		}
	case 0x10c7:
		{
			m.ip = 0x10cc
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2077)), uint16(1))
			m.cycles += 16
		}
	case 0x10cc:
		{
			m.ip = 0x10ce
			if !m.zf {
				m.ip = uint16(4310)
			}
			m.cycles += 8
		}
	case 0x10ce:
		{
			m.ip = 0x10d3
			m.wr8(m.r[11], uint16(0x2077), uint16(0))
			m.cycles += 8
		}
	case 0x10d3:
		{
			m.ip = 0x10d6
			m.ip = uint16(3699)
			m.cycles += 15
		}
	case 0x10d6:
		{
			m.ip = 0x10db
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2075)), uint16(1))
			m.cycles += 16
		}
	case 0x10db:
		{
			m.ip = 0x10dd
			if !m.zf {
				m.ip = uint16(4265)
			}
			m.cycles += 8
		}
	case 0x10dd:
		{
			m.ip = 0x10e0
			m.ip = uint16(4518)
			m.cycles += 15
		}
	case 0x10e0:
		{
			m.ip = 0x10e3
			m.alu("sub", 16, m.r[0], uint16(4113))
			m.cycles += 4
		}
	case 0x10e3:
		{
			m.ip = 0x10e5
			if !m.zf {
				m.ip = uint16(4328)
			}
			m.cycles += 8
		}
	case 0x10e5:
		{
			m.ip = 0x10e8
			m.ip = uint16(17160)
			m.cycles += 15
		}
	case 0x10e8:
		{
			m.ip = 0x10ea
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(97))
			m.cycles += 4
		}
	case 0x10ea:
		{
			m.ip = 0x10ec
			if m.sf != m.of {
				m.ip = uint16(4334)
			}
			m.cycles += 8
		}
	case 0x10ec:
		{
			m.ip = 0x10ee
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), uint16(32)))
			m.cycles += 4
		}
	case 0x10ee:
		{
			m.ip = 0x10f1
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x10f1:
		{
			m.ip = 0x10f3
			m.r[6] = m.r[3]
			m.cycles += 2
		}
	case 0x10f3:
		{
			m.ip = 0x10f5
			m.r[6] = m.shift("shr", 16, m.r[6], uint16(1))
			m.cycles += 8
		}
	case 0x10f5:
		{
			m.ip = 0x10f7
			m.r[6] = m.alu("add", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x10f7:
		{
			m.ip = 0x10f9
			m.r[6] = m.alu("add", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x10f9:
		{
			m.ip = 0x10fc
			m.r[5] = uint16(320)
			m.cycles += 2
		}
	case 0x10fc:
		{
			m.ip = 0x10ff
			m.r[1] = uint16(5)
			m.cycles += 2
		}
	case 0x10ff:
		{
			m.ip = 0x1103
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), m.rd8(m.r[11], uint16(m.r[6]+0x20bf)))
			m.cycles += 16
		}
	case 0x1103:
		{
			m.ip = 0x1105
			if m.zf {
				m.ip = uint16(4370)
			}
			m.cycles += 8
		}
	case 0x1105:
		{
			m.ip = 0x1106
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1106:
		{
			m.ip = 0x1109
			m.r[2] = m.alu("add", 16, m.r[2], uint16(4))
			m.cycles += 4
		}
	case 0x1109:
		{
			m.ip = 0x110d
			m.r[5] = m.alu("add", 16, m.r[5], uint16(320))
			m.cycles += 4
		}
	case 0x110d:
		{
			m.ip = 0x110f
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4351)
			}
			m.cycles += 17
		}
	case 0x110f:
		{
			m.ip = 0x1112
			m.ip = uint16(3839)
			m.cycles += 15
		}
	case 0x1112:
		{
			m.ip = 0x1115
			m.alu("sub", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x1115:
		{
			m.ip = 0x1117
			if !m.zf {
				m.ip = uint16(4378)
			}
			m.cycles += 8
		}
	case 0x1117:
		{
			m.ip = 0x111a
			m.ip = uint16(3839)
			m.cycles += 15
		}
	case 0x111a:
		{
			m.ip = 0x111e
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), m.r[2])
			m.cycles += 8
		}
	case 0x111e:
		{
			m.ip = 0x1121
			m.alu("sub", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x1121:
		{
			m.ip = 0x1123
			if !m.zf {
				m.ip = uint16(4387)
			}
			m.cycles += 8
		}
	case 0x1123:
		{
			m.ip = 0x1126
			target := uint16(4432)
			m.push(0x1126)
			m.ip = target
			m.cycles += 19
		}
	case 0x1126:
		{
			m.ip = 0x1128
			m.set8(2, 0, uint16(7))
			m.cycles += 2
		}
	case 0x1128:
		{
			m.ip = 0x112c
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x208d))
			m.cycles += 8
		}
	case 0x112c:
		{
			m.ip = 0x1130
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x2083))
			m.cycles += 8
		}
	case 0x1130:
		{
			m.ip = 0x1132
			m.r[7] = m.alu("add", 16, m.r[7], m.r[0])
			m.cycles += 4
		}
	case 0x1132:
		{
			m.ip = 0x1135
			m.r[1] = uint16(9)
			m.cycles += 2
		}
	case 0x1135:
		{
			m.ip = 0x1139
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[2] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1139:
		{
			m.ip = 0x113c
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x113c:
		{
			m.ip = 0x113e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4405)
			}
			m.cycles += 17
		}
	case 0x113e:
		{
			m.ip = 0x1141
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(1))
			m.cycles += 4
		}
	case 0x1141:
		{
			m.ip = 0x1143
			if !m.zf {
				m.ip = uint16(4422)
			}
			m.cycles += 8
		}
	case 0x1143:
		{
			m.ip = 0x1146
			m.ip = uint16(3839)
			m.cycles += 15
		}
	case 0x1146:
		{
			m.ip = 0x1148
			m.r[0] = m.r[5]
			m.cycles += 2
		}
	case 0x1148:
		{
			m.ip = 0x114c
			m.wr16(m.r[11], uint16(m.r[3]+0x208d), m.r[5])
			m.cycles += 8
		}
	case 0x114c:
		{
			m.ip = 0x114e
			m.set8(2, 0, uint16(1))
			m.cycles += 2
		}
	case 0x114e:
		{
			m.ip = 0x1150
			m.ip = uint16(4396)
			m.cycles += 15
		}
	case 0x1150:
		{
			m.ip = 0x1153
			m.alu("sub", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x1153:
		{
			m.ip = 0x1155
			if !m.zf {
				m.ip = uint16(4477)
			}
			m.cycles += 8
		}
	case 0x1155:
		{
			m.ip = 0x1157
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(67))
			m.cycles += 4
		}
	case 0x1157:
		{
			m.ip = 0x1159
			if !m.zf {
				m.ip = uint16(4517)
			}
			m.cycles += 8
		}
	case 0x1159:
		{
			m.ip = 0x115f
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(640))
			m.cycles += 16
		}
	case 0x115f:
		{
			m.ip = 0x1161
			if !m.zf {
				m.ip = uint16(4459)
			}
			m.cycles += 8
		}
	case 0x1161:
		{
			m.ip = 0x1164
			m.r[5] = uint16(960)
			m.cycles += 2
		}
	case 0x1164:
		{
			m.ip = 0x116a
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), uint16(8))
			m.cycles += 8
		}
	case 0x116a:
		{
			m.ip = 0x116b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x116b:
		{
			m.ip = 0x1171
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(960))
			m.cycles += 16
		}
	case 0x1171:
		{
			m.ip = 0x1173
			if !m.zf {
				m.ip = uint16(4517)
			}
			m.cycles += 8
		}
	case 0x1173:
		{
			m.ip = 0x1176
			m.r[5] = uint16(640)
			m.cycles += 2
		}
	case 0x1176:
		{
			m.ip = 0x117c
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), uint16(4))
			m.cycles += 8
		}
	case 0x117c:
		{
			m.ip = 0x117d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x117d:
		{
			m.ip = 0x1180
			m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x1180:
		{
			m.ip = 0x1182
			if !m.zf {
				m.ip = uint16(4517)
			}
			m.cycles += 8
		}
	case 0x1182:
		{
			m.ip = 0x1188
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(320))
			m.cycles += 16
		}
	case 0x1188:
		{
			m.ip = 0x118a
			if !m.zf {
				m.ip = uint16(4500)
			}
			m.cycles += 8
		}
	case 0x118a:
		{
			m.ip = 0x118d
			m.r[5] = uint16(640)
			m.cycles += 2
		}
	case 0x118d:
		{
			m.ip = 0x1193
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), uint16(4))
			m.cycles += 8
		}
	case 0x1193:
		{
			m.ip = 0x1194
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1194:
		{
			m.ip = 0x119a
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(640))
			m.cycles += 16
		}
	case 0x119a:
		{
			m.ip = 0x119c
			if !m.zf {
				m.ip = uint16(4517)
			}
			m.cycles += 8
		}
	case 0x119c:
		{
			m.ip = 0x119f
			m.r[5] = uint16(320)
			m.cycles += 2
		}
	case 0x119f:
		{
			m.ip = 0x11a5
			m.wr16(m.r[11], uint16(m.r[3]+0x20ab), uint16(0))
			m.cycles += 8
		}
	case 0x11a5:
		{
			m.ip = 0x11a6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x11a6:
		{
			m.ip = 0x11a7
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x11a7:
		{
			m.ip = 0x11a8
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x11a8:
		{
			m.ip = 0x11a9
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x11a9:
		{
			m.ip = 0x11aa
			m.push(m.r[8])
			m.cycles += 11
		}
	case 0x11aa:
		{
			m.ip = 0x11ab
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x11ab:
		{
			m.ip = 0x11af
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x11af:
		{
			m.ip = 0x11b0
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x11b0:
		{
			m.ip = 0x11b1
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x11b1:
		{
			m.ip = 0x11b2
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x11b2:
		{
			m.ip = 0x11b6
			m.r[2] = m.rd16(m.r[11], uint16(0x24f2))
			m.cycles += 8
		}
	case 0x11b6:
		{
			m.ip = 0x11b8
			m.set8(2, 0, m.unary("neg", 8, ((m.r[2]>>0)&255)))
			m.cycles += 3
		}
	case 0x11b8:
		{
			m.ip = 0x11bb
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), uint16(15)))
			m.cycles += 4
		}
	case 0x11bb:
		{
			m.ip = 0x11bd
			m.set8(2, 0, m.shift("shl", 8, ((m.r[2]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x11bd:
		{
			m.ip = 0x11bf
			m.set8(3, 0, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x11bf:
		{
			m.ip = 0x11c1
			m.set8(2, 0, m.shift("shl", 8, ((m.r[2]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x11c1:
		{
			m.ip = 0x11c3
			m.set8(2, 0, m.shift("shl", 8, ((m.r[2]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x11c3:
		{
			m.ip = 0x11c5
			m.set8(2, 0, m.shift("shl", 8, ((m.r[2]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x11c5:
		{
			m.ip = 0x11c7
			m.set8(2, 0, m.alu("sub", 8, ((m.r[2]>>0)&255), ((m.r[3]>>0)&255)))
			m.cycles += 4
		}
	case 0x11c7:
		{
			m.ip = 0x11cb
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), m.rd8(m.r[11], uint16(0x24f6))))
			m.cycles += 16
		}
	case 0x11cb:
		{
			m.ip = 0x11ce
			m.r[7] = uint16(1012)
			m.cycles += 2
		}
	case 0x11ce:
		{
			m.ip = 0x11d2
			m.r[3] = uint16(0x1027)
			m.cycles += 3
		}
	case 0x11d2:
		{
			m.ip = 0x11d4
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x11d4:
		{
			m.ip = 0x11d6
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3])))
			m.cycles += 16
		}
	case 0x11d6:
		{
			m.ip = 0x11d8
			if m.sf != m.of {
				m.ip = uint16(4574)
			}
			m.cycles += 8
		}
	case 0x11d8:
		{
			m.ip = 0x11da
			m.set8(0, 0, m.unary("inc", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x11da:
		{
			m.ip = 0x11dc
			m.r[2] = m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3])))
			m.cycles += 16
		}
	case 0x11dc:
		{
			m.ip = 0x11de
			m.ip = uint16(4564)
			m.cycles += 15
		}
	case 0x11de:
		{
			m.ip = 0x11e0
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), uint16(48)))
			m.cycles += 4
		}
	case 0x11e0:
		{
			m.ip = 0x11e3
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x11e3:
		{
			m.ip = 0x11e6
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x11e6:
		{
			m.ip = 0x11e9
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x11e9:
		{
			m.ip = 0x11eb
			m.set8(0, 0, uint16(10))
			m.cycles += 2
		}
	case 0x11eb:
		{
			m.ip = 0x11ee
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0xfffe)), ((m.r[0] >> 0) & 255))
			m.cycles += 16
		}
	case 0x11ee:
		{
			m.ip = 0x11f0
			if !m.zf {
				m.ip = uint16(4562)
			}
			m.cycles += 8
		}
	case 0x11f0:
		{
			m.ip = 0x11f2
			m.set8(0, 0, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x11f2:
		{
			m.ip = 0x11f4
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), uint16(48)))
			m.cycles += 4
		}
	case 0x11f4:
		{
			m.ip = 0x11f7
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x11f7:
		{
			m.ip = 0x11f8
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x11f8:
		{
			m.ip = 0x11f9
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x11f9:
		{
			m.ip = 0x11fa
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x11fa:
		{
			m.ip = 0x11fd
			m.set8(0, 8, m.alu("add", 8, ((m.r[0]>>8)&255), uint16(4)))
			m.cycles += 4
		}
	case 0x11fd:
		{
			m.ip = 0x1200
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(8))
			m.cycles += 4
		}
	case 0x1200:
		{
			m.ip = 0x1202
			if !m.cf {
				m.ip = uint16(4613)
			}
			m.cycles += 8
		}
	case 0x1202:
		{
			m.ip = 0x1205
			m.set8(0, 8, m.alu("add", 8, ((m.r[0]>>8)&255), uint16(4)))
			m.cycles += 4
		}
	case 0x1205:
		{
			m.ip = 0x120b
			m.wr16(m.r[11], uint16(0x24f4), uint16(0))
			m.cycles += 8
		}
	case 0x120b:
		{
			m.ip = 0x1210
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(15))
			m.cycles += 16
		}
	case 0x1210:
		{
			m.ip = 0x1212
			if m.cf || m.zf {
				m.ip = uint16(4632)
			}
			m.cycles += 8
		}
	case 0x1212:
		{
			m.ip = 0x1218
			m.wr16(m.r[11], uint16(0x24f2), uint16(15))
			m.cycles += 8
		}
	case 0x1218:
		{
			m.ip = 0x121d
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(7))
			m.cycles += 16
		}
	case 0x121d:
		{
			m.ip = 0x121f
			if m.sf == m.of {
				m.ip = uint16(4645)
			}
			m.cycles += 8
		}
	case 0x121f:
		{
			m.ip = 0x1225
			m.wr16(m.r[11], uint16(0x24f2), uint16(7))
			m.cycles += 8
		}
	case 0x1225:
		{
			m.ip = 0x122a
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(7))
			m.cycles += 16
		}
	case 0x122a:
		{
			m.ip = 0x122c
			if m.zf || m.sf != m.of {
				m.ip = uint16(4677)
			}
			m.cycles += 8
		}
	case 0x122c:
		{
			m.ip = 0x1230
			m.r[1] = m.rd16(m.r[11], uint16(0x24f2))
			m.cycles += 8
		}
	case 0x1230:
		{
			m.ip = 0x1233
			m.r[1] = m.alu("sub", 16, m.r[1], uint16(7))
			m.cycles += 4
		}
	case 0x1233:
		{
			m.ip = 0x1235
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x1235:
		{
			m.ip = 0x1238
			m.r[7] = uint16(1336)
			m.cycles += 2
		}
	case 0x1238:
		{
			m.ip = 0x123b
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x123b:
		{
			m.ip = 0x123f
			m.wr16(m.r[8], uint16(m.r[7]+0x2), m.r[0])
			m.cycles += 8
		}
	case 0x123f:
		{
			m.ip = 0x1243
			m.r[7] = m.alu("add", 16, m.r[7], uint16(160))
			m.cycles += 4
		}
	case 0x1243:
		{
			m.ip = 0x1245
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4664)
			}
			m.cycles += 17
		}
	case 0x1245:
		{
			m.ip = 0x1249
			m.r[1] = m.rd16(m.r[11], uint16(0x24f2))
			m.cycles += 8
		}
	case 0x1249:
		{
			m.ip = 0x124c
			m.r[1] = m.alu("add", 16, m.r[1], uint16(1))
			m.cycles += 4
		}
	case 0x124c:
		{
			m.ip = 0x1252
			m.wr16(m.r[11], uint16(0x24f4), uint16(0))
			m.cycles += 8
		}
	case 0x1252:
		{
			m.ip = 0x1258
			m.wr16(m.r[11], uint16(0x24f4), m.alu("add", 16, m.rd16(m.r[11], uint16(0x24f4)), uint16(160)))
			m.cycles += 16
		}
	case 0x1258:
		{
			m.ip = 0x125a
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4690)
			}
			m.cycles += 17
		}
	case 0x125a:
		{
			m.ip = 0x125f
			m.wr16(m.r[11], uint16(0x24f4), m.alu("add", 16, m.rd16(m.r[11], uint16(0x24f4)), uint16(56)))
			m.cycles += 16
		}
	case 0x125f:
		{
			m.ip = 0x1263
			m.r[7] = m.rd16(m.r[11], uint16(0x24f4))
			m.cycles += 8
		}
	case 0x1263:
		{
			m.ip = 0x1266
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x24f6)))
			m.cycles += 8
		}
	case 0x1266:
		{
			m.ip = 0x1269
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1269:
		{
			m.ip = 0x126d
			m.wr16(m.r[8], uint16(m.r[7]+0x2), m.r[0])
			m.cycles += 8
		}
	case 0x126d:
		{
			m.ip = 0x1272
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(15))
			m.cycles += 16
		}
	case 0x1272:
		{
			m.ip = 0x1274
			if !m.cf {
				m.ip = uint16(4756)
			}
			m.cycles += 8
		}
	case 0x1274:
		{
			m.ip = 0x1277
			m.r[1] = uint16(15)
			m.cycles += 2
		}
	case 0x1277:
		{
			m.ip = 0x127b
			m.r[1] = m.alu("sub", 16, m.r[1], m.rd16(m.r[11], uint16(0x24f2)))
			m.cycles += 16
		}
	case 0x127b:
		{
			m.ip = 0x127d
			m.set8(0, 0, uint16(14))
			m.cycles += 2
		}
	case 0x127d:
		{
			m.ip = 0x1283
			m.wr16(m.r[11], uint16(0x24f4), m.alu("add", 16, m.rd16(m.r[11], uint16(0x24f4)), uint16(160)))
			m.cycles += 16
		}
	case 0x1283:
		{
			m.ip = 0x1287
			m.r[7] = m.rd16(m.r[11], uint16(0x24f4))
			m.cycles += 8
		}
	case 0x1287:
		{
			m.ip = 0x128a
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x128a:
		{
			m.ip = 0x128e
			m.wr16(m.r[8], uint16(m.r[7]+0x2), m.r[0])
			m.cycles += 8
		}
	case 0x128e:
		{
			m.ip = 0x1292
			m.r[7] = m.alu("add", 16, m.r[7], uint16(160))
			m.cycles += 4
		}
	case 0x1292:
		{
			m.ip = 0x1294
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4743)
			}
			m.cycles += 17
		}
	case 0x1294:
		{
			m.ip = 0x1295
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x1295:
		{
			m.ip = 0x1296
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x1296:
		{
			m.ip = 0x1297
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1297:
		{
			m.ip = 0x1298
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x1298:
		{
			m.ip = 0x1299
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1299:
		{
			m.ip = 0x129b
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x129b:
		{
			m.ip = 0x129d
			m.set8(0, 0, uint16(32))
			m.cycles += 2
		}
	case 0x129d:
		{
			m.ip = 0x12a0
			m.r[7] = uint16(3360)
			m.cycles += 2
		}
	case 0x12a0:
		{
			m.ip = 0x12a3
			m.r[1] = uint16(240)
			m.cycles += 2
		}
	case 0x12a3:
		{
			m.ip = 0x12a6
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x12a6:
		{
			m.ip = 0x12a9
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x12a9:
		{
			m.ip = 0x12ab
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4771)
			}
			m.cycles += 17
		}
	case 0x12ab:
		{
			m.ip = 0x12af
			m.r[6] = m.rd16(m.r[11], uint16(m.r[3]+0x221e))
			m.cycles += 8
		}
	case 0x12af:
		{
			m.ip = 0x12b4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x615)), uint16(1))
			m.cycles += 16
		}
	case 0x12b4:
		{
			m.ip = 0x12b6
			if !m.zf {
				m.ip = uint16(4794)
			}
			m.cycles += 8
		}
	case 0x12b6:
		{
			m.ip = 0x12ba
			m.r[6] = m.rd16(m.r[11], uint16(m.r[3]+0x2214))
			m.cycles += 8
		}
	case 0x12ba:
		{
			m.ip = 0x12be
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(m.r[3]+0x20ab)))
			m.cycles += 16
		}
	case 0x12be:
		{
			m.ip = 0x12bf
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x12bf:
		{
			m.ip = 0x12c0
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x12c0:
		{
			m.ip = 0x12c1
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x12c1:
		{
			m.ip = 0x12c3
			m.r[6] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x12c3:
		{
			m.ip = 0x12c5
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x12c5:
		{
			m.ip = 0x12c6
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x12c6:
		{
			m.ip = 0x12c9
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x12c9:
		{
			m.ip = 0x12cb
			m.r[6] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x12cb:
		{
			m.ip = 0x12ce
			m.r[7] = uint16(3600)
			m.cycles += 2
		}
	case 0x12ce:
		{
			m.ip = 0x12d1
			m.alu("sub", 16, m.r[6], uint16(0))
			m.cycles += 4
		}
	case 0x12d1:
		{
			m.ip = 0x12d3
			if m.zf {
				m.ip = uint16(4821)
			}
			m.cycles += 8
		}
	case 0x12d3:
		{
			m.ip = 0x12d5
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(m.r[6]))))
			m.cycles += 16
		}
	case 0x12d5:
		{
			m.ip = 0x12d7
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x12d7:
		{
			m.ip = 0x12d9
			m.r[7] = m.alu("sub", 16, m.r[7], m.r[0])
			m.cycles += 4
		}
	case 0x12d9:
		{
			m.ip = 0x12db
			m.r[7] = m.shift("shr", 16, m.r[7], uint16(1))
			m.cycles += 8
		}
	case 0x12db:
		{
			m.ip = 0x12dd
			m.r[7] = m.shift("shl", 16, m.r[7], uint16(1))
			m.cycles += 8
		}
	case 0x12dd:
		{
			m.ip = 0x12de
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x12de:
		{
			m.ip = 0x12e0
			m.r[5] = m.r[0]
			m.cycles += 2
		}
	case 0x12e0:
		{
			m.ip = 0x12e2
			m.r[5] = m.shift("shl", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x12e2:
		{
			m.ip = 0x12e4
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x12e4:
		{
			m.ip = 0x12e8
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(162))
			m.cycles += 4
		}
	case 0x12e8:
		{
			m.ip = 0x12ed
			m.wr8(m.r[8], uint16(m.r[7]+0xfffe), uint16(218))
			m.cycles += 8
		}
	case 0x12ed:
		{
			m.ip = 0x12f3
			m.wr8(m.r[8], uint16(m.r[7]+0x9e), uint16(179))
			m.cycles += 8
		}
	case 0x12f3:
		{
			m.ip = 0x12f9
			m.wr8(m.r[8], uint16(m.r[7]+0x13e), uint16(192))
			m.cycles += 8
		}
	case 0x12f9:
		{
			m.ip = 0x12fe
			m.wr8(m.r[8], uint16(m.r[5]+m.r[7]+0x4), uint16(191))
			m.cycles += 8
		}
	case 0x12fe:
		{
			m.ip = 0x1304
			m.wr8(m.r[8], uint16(m.r[5]+m.r[7]+0xa4), uint16(179))
			m.cycles += 8
		}
	case 0x1304:
		{
			m.ip = 0x130a
			m.wr8(m.r[8], uint16(m.r[5]+m.r[7]+0x144), uint16(217))
			m.cycles += 8
		}
	case 0x130a:
		{
			m.ip = 0x130c
			m.r[5] = m.shift("shr", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x130c:
		{
			m.ip = 0x130f
			m.r[5] = m.alu("add", 16, m.r[5], uint16(2))
			m.cycles += 4
		}
	case 0x130f:
		{
			m.ip = 0x1312
			m.r[1] = uint16(2)
			m.cycles += 2
		}
	case 0x1312:
		{
			m.ip = 0x1313
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x1313:
		{
			m.ip = 0x1315
			m.r[1] = m.r[5]
			m.cycles += 2
		}
	case 0x1315:
		{
			m.ip = 0x1319
			m.wr8(m.r[8], uint16(m.r[7]), uint16(196))
			m.cycles += 8
		}
	case 0x1319:
		{
			m.ip = 0x131c
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x131c:
		{
			m.ip = 0x131e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4885)
			}
			m.cycles += 17
		}
	case 0x131e:
		{
			m.ip = 0x1322
			m.r[7] = m.alu("add", 16, m.r[7], uint16(320))
			m.cycles += 4
		}
	case 0x1322:
		{
			m.ip = 0x1324
			m.r[7] = m.alu("sub", 16, m.r[7], m.r[5])
			m.cycles += 4
		}
	case 0x1324:
		{
			m.ip = 0x1326
			m.r[7] = m.alu("sub", 16, m.r[7], m.r[5])
			m.cycles += 4
		}
	case 0x1326:
		{
			m.ip = 0x1327
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1327:
		{
			m.ip = 0x1329
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(4882)
			}
			m.cycles += 17
		}
	case 0x1329:
		{
			m.ip = 0x132a
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x132a:
		{
			m.ip = 0x132b
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x132b:
		{
			m.ip = 0x132d
			m.r[6] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x132d:
		{
			m.ip = 0x132e
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x132e:
		{
			m.ip = 0x1331
			target := uint16(9087)
			m.push(0x1331)
			m.ip = target
			m.cycles += 19
		}
	case 0x1331:
		{
			m.ip = 0x1332
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x1332:
		{
			m.ip = 0x1335
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x1335:
		{
			m.ip = 0x1337
			m.r[6] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x1337:
		{
			m.ip = 0x133a
			m.alu("sub", 16, m.r[6], uint16(0))
			m.cycles += 4
		}
	case 0x133a:
		{
			m.ip = 0x133c
			if m.zf {
				m.ip = uint16(4928)
			}
			m.cycles += 8
		}
	case 0x133c:
		{
			m.ip = 0x133d
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x133d:
		{
			m.ip = 0x1340
			target := uint16(9087)
			m.push(0x1340)
			m.ip = target
			m.cycles += 19
		}
	case 0x1340:
		{
			m.ip = 0x1341
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1341:
		{
			m.ip = 0x1344
			m.alu("sub", 16, m.r[3], uint16(0))
			m.cycles += 4
		}
	case 0x1344:
		{
			m.ip = 0x1346
			if m.zf {
				m.ip = uint16(4942)
			}
			m.cycles += 8
		}
	case 0x1346:
		{
			m.ip = 0x1348
			m.r[2] = m.r[3]
			m.cycles += 2
		}
	case 0x1348:
		{
			m.ip = 0x134b
			m.r[3] = uint16(4)
			m.cycles += 2
		}
	case 0x134b:
		{
			m.ip = 0x134d
			m.ip = uint16(5029)
			m.cycles += 15
		}
	case 0x134d:
		{
			m.ip = 0x134e
			m.cycles += 3
		}
	case 0x134e:
		{
			m.ip = 0x1354
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(320))
			m.cycles += 16
		}
	case 0x1354:
		{
			m.ip = 0x1356
			if !m.zf {
				m.ip = uint16(4972)
			}
			m.cycles += 8
		}
	case 0x1356:
		{
			m.ip = 0x135b
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x135b:
		{
			m.ip = 0x135e
			target := uint16(3447)
			m.push(0x135e)
			m.ip = target
			m.cycles += 19
		}
	case 0x135e:
		{
			m.ip = 0x1360
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x1360:
		{
			m.ip = 0x1363
			target := uint16(7032)
			m.push(0x1363)
			m.ip = target
			m.cycles += 19
		}
	case 0x1363:
		{
			m.ip = 0x1366
			target := uint16(3447)
			m.push(0x1366)
			m.ip = target
			m.cycles += 19
		}
	case 0x1366:
		{
			m.ip = 0x136b
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x136b:
		{
			m.ip = 0x136c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x136c:
		{
			m.ip = 0x1372
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(640))
			m.cycles += 16
		}
	case 0x1372:
		{
			m.ip = 0x1374
			if !m.zf {
				m.ip = uint16(4999)
			}
			m.cycles += 8
		}
	case 0x1374:
		{
			m.ip = 0x1377
			target := uint16(5320)
			m.push(0x1377)
			m.ip = target
			m.cycles += 19
		}
	case 0x1377:
		{
			m.ip = 0x137c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x3b2)), uint16(1))
			m.cycles += 16
		}
	case 0x137c:
		{
			m.ip = 0x137e
			if m.zf {
				m.ip = uint16(4993)
			}
			m.cycles += 8
		}
	case 0x137e:
		{
			m.ip = 0x1381
			target := uint16(3447)
			m.push(0x1381)
			m.ip = target
			m.cycles += 19
		}
	case 0x1381:
		{
			m.ip = 0x1386
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x1386:
		{
			m.ip = 0x1387
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1387:
		{
			m.ip = 0x138d
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(960))
			m.cycles += 16
		}
	case 0x138d:
		{
			m.ip = 0x138f
			if !m.zf {
				m.ip = uint16(5011)
			}
			m.cycles += 8
		}
	case 0x138f:
		{
			m.ip = 0x1392
			target := uint16(5556)
			m.push(0x1392)
			m.ip = target
			m.cycles += 19
		}
	case 0x1392:
		{
			m.ip = 0x1393
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1393:
		{
			m.ip = 0x1398
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x615)), uint16(1))
			m.cycles += 16
		}
	case 0x1398:
		{
			m.ip = 0x139a
			if !m.zf {
				m.ip = uint16(5028)
			}
			m.cycles += 8
		}
	case 0x139a:
		{
			m.ip = 0x139f
			m.wr8(m.r[11], uint16(0x2076), uint16(1))
			m.cycles += 8
		}
	case 0x139f:
		{
			m.ip = 0x13a4
			m.wr8(m.r[11], uint16(0x2075), uint16(1))
			m.cycles += 8
		}
	case 0x13a4:
		{
			m.ip = 0x13a5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x13a5:
		{
			m.ip = 0x13ab
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(320))
			m.cycles += 16
		}
	case 0x13ab:
		{
			m.ip = 0x13ad
			if !m.zf {
				m.ip = uint16(5054)
			}
			m.cycles += 8
		}
	case 0x13ad:
		{
			m.ip = 0x13b0
			target := uint16(5118)
			m.push(0x13b0)
			m.ip = target
			m.cycles += 19
		}
	case 0x13b0:
		{
			m.ip = 0x13b5
			m.wr8(m.r[11], uint16(0x615), uint16(0))
			m.cycles += 8
		}
	case 0x13b5:
		{
			m.ip = 0x13ba
			m.wr8(m.r[11], uint16(0x75), uint16(1))
			m.cycles += 8
		}
	case 0x13ba:
		{
			m.ip = 0x13bd
			m.ip = uint16(295)
			m.cycles += 15
		}
	case 0x13bd:
		{
			m.ip = 0x13be
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x13be:
		{
			m.ip = 0x13c4
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(640))
			m.cycles += 16
		}
	case 0x13c4:
		{
			m.ip = 0x13c6
			if !m.zf {
				m.ip = uint16(5081)
			}
			m.cycles += 8
		}
	case 0x13c6:
		{
			m.ip = 0x13cb
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x615)), uint16(1))
			m.cycles += 16
		}
	case 0x13cb:
		{
			m.ip = 0x13cd
			if m.zf {
				m.ip = uint16(5072)
			}
			m.cycles += 8
		}
	case 0x13cd:
		{
			m.ip = 0x13cf
			m.r[3] = m.r[2]
			m.cycles += 2
		}
	case 0x13cf:
		{
			m.ip = 0x13d0
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x13d0:
		{
			m.ip = 0x13d5
			m.wr8(m.r[11], uint16(0x2075), uint16(1))
			m.cycles += 8
		}
	case 0x13d5:
		{
			m.ip = 0x13d8
			target := uint16(5118)
			m.push(0x13d8)
			m.ip = target
			m.cycles += 19
		}
	case 0x13d8:
		{
			m.ip = 0x13d9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x13d9:
		{
			m.ip = 0x13df
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(960))
			m.cycles += 16
		}
	case 0x13df:
		{
			m.ip = 0x13e1
			if !m.zf {
				m.ip = uint16(5093)
			}
			m.cycles += 8
		}
	case 0x13e1:
		{
			m.ip = 0x13e4
			target := uint16(6363)
			m.push(0x13e4)
			m.ip = target
			m.cycles += 19
		}
	case 0x13e4:
		{
			m.ip = 0x13e5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x13e5:
		{
			m.ip = 0x13eb
			m.wr16(m.r[11], uint16(0x78), uint16(0))
			m.cycles += 8
		}
	case 0x13eb:
		{
			m.ip = 0x13f0
			m.wr8(m.r[11], uint16(0x74), uint16(1))
			m.cycles += 8
		}
	case 0x13f0:
		{
			m.ip = 0x13f5
			m.wr8(m.r[11], uint16(0x77), uint16(1))
			m.cycles += 8
		}
	case 0x13f5:
		{
			m.ip = 0x13fa
			m.wr8(m.r[11], uint16(0x2075), uint16(1))
			m.cycles += 8
		}
	case 0x13fa:
		{
			m.ip = 0x13fd
			m.ip = uint16(295)
			m.cycles += 15
		}
	case 0x13fd:
		{
			m.ip = 0x13fe
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x13fe:
		{
			m.ip = 0x1401
			m.r[3] = uint16(6)
			m.cycles += 2
		}
	case 0x1401:
		{
			m.ip = 0x1406
			m.wr8(m.r[11], uint16(0x617), uint16(1))
			m.cycles += 8
		}
	case 0x1406:
		{
			m.ip = 0x140b
			m.wr8(m.r[11], uint16(0x619), uint16(0))
			m.cycles += 8
		}
	case 0x140b:
		{
			m.ip = 0x1410
			m.wr8(m.r[11], uint16(0x61b), uint16(0))
			m.cycles += 8
		}
	case 0x1410:
		{
			m.ip = 0x1416
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(320))
			m.cycles += 16
		}
	case 0x1416:
		{
			m.ip = 0x1418
			if m.zf {
				m.ip = uint16(5181)
			}
			m.cycles += 8
		}
	case 0x1418:
		{
			m.ip = 0x141e
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(640))
			m.cycles += 16
		}
	case 0x141e:
		{
			m.ip = 0x1420
			if m.zf {
				m.ip = uint16(5176)
			}
			m.cycles += 8
		}
	case 0x1420:
		{
			m.ip = 0x1426
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(960))
			m.cycles += 16
		}
	case 0x1426:
		{
			m.ip = 0x1428
			if !m.zf {
				m.ip = uint16(5168)
			}
			m.cycles += 8
		}
	case 0x1428:
		{
			m.ip = 0x142d
			m.wr8(m.r[11], uint16(0x619), uint16(1))
			m.cycles += 8
		}
	case 0x142d:
		{
			m.ip = 0x142f
			m.ip = uint16(5181)
			m.cycles += 15
		}
	case 0x142f:
		{
			m.ip = 0x1430
			m.cycles += 3
		}
	case 0x1430:
		{
			m.ip = 0x1435
			m.wr8(m.r[11], uint16(0x61b), uint16(1))
			m.cycles += 8
		}
	case 0x1435:
		{
			m.ip = 0x1437
			m.ip = uint16(5181)
			m.cycles += 15
		}
	case 0x1437:
		{
			m.ip = 0x1438
			m.cycles += 3
		}
	case 0x1438:
		{
			m.ip = 0x143d
			m.wr8(m.r[11], uint16(0x617), uint16(2))
			m.cycles += 8
		}
	case 0x143d:
		{
			m.ip = 0x1440
			m.r[3] = uint16(8)
			m.cycles += 2
		}
	case 0x1440:
		{
			m.ip = 0x1445
			m.wr8(m.r[11], uint16(0x55a), uint16(1))
			m.cycles += 8
		}
	case 0x1445:
		{
			m.ip = 0x144b
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x208d)), uint16(320))
			m.cycles += 16
		}
	case 0x144b:
		{
			m.ip = 0x144d
			if m.zf {
				m.ip = uint16(5202)
			}
			m.cycles += 8
		}
	case 0x144d:
		{
			m.ip = 0x1452
			m.wr8(m.r[11], uint16(0x55a), uint16(0))
			m.cycles += 8
		}
	case 0x1452:
		{
			m.ip = 0x1453
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1453:
		{
			m.ip = 0x1457
			m.r[6] = uint16(0x21be)
			m.cycles += 3
		}
	case 0x1457:
		{
			m.ip = 0x145a
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(4))
			m.cycles += 4
		}
	case 0x145a:
		{
			m.ip = 0x145c
			if !m.zf {
				m.ip = uint16(5216)
			}
			m.cycles += 8
		}
	case 0x145c:
		{
			m.ip = 0x1460
			m.r[6] = uint16(0x21e9)
			m.cycles += 3
		}
	case 0x1460:
		{
			m.ip = 0x1464
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x2083))
			m.cycles += 8
		}
	case 0x1464:
		{
			m.ip = 0x1468
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(164))
			m.cycles += 4
		}
	case 0x1468:
		{
			m.ip = 0x146b
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x146b:
		{
			m.ip = 0x146d
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x146d:
		{
			m.ip = 0x1470
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1470:
		{
			m.ip = 0x1471
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1471:
		{
			m.ip = 0x1474
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1474:
		{
			m.ip = 0x1476
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(5227)
			}
			m.cycles += 17
		}
	case 0x1476:
		{
			m.ip = 0x147a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(132))
			m.cycles += 4
		}
	case 0x147a:
		{
			m.ip = 0x147d
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x1c)))
			m.cycles += 8
		}
	case 0x147d:
		{
			m.ip = 0x1480
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1480:
		{
			m.ip = 0x1483
			m.r[7] = m.alu("add", 16, m.r[7], uint16(26))
			m.cycles += 4
		}
	case 0x1483:
		{
			m.ip = 0x1486
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1486:
		{
			m.ip = 0x148a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(134))
			m.cycles += 4
		}
	case 0x148a:
		{
			m.ip = 0x148d
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x148d:
		{
			m.ip = 0x148f
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x148f:
		{
			m.ip = 0x1492
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1492:
		{
			m.ip = 0x1493
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1493:
		{
			m.ip = 0x1496
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1496:
		{
			m.ip = 0x1498
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(5261)
			}
			m.cycles += 17
		}
	case 0x1498:
		{
			m.ip = 0x149b
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0xe)))
			m.cycles += 8
		}
	case 0x149b:
		{
			m.ip = 0x149f
			m.r[1] = m.rd16(m.r[11], uint16(m.r[3]+0x20a1))
			m.cycles += 8
		}
	case 0x149f:
		{
			m.ip = 0x14a0
			m.r[1] = m.unary("dec", 16, m.r[1])
			m.cycles += 3
		}
	case 0x14a0:
		{
			m.ip = 0x14a2
			m.r[1] = m.shift("shl", 16, m.r[1], uint16(1))
			m.cycles += 8
		}
	case 0x14a2:
		{
			m.ip = 0x14a3
			m.r[1] = m.unary("dec", 16, m.r[1])
			m.cycles += 3
		}
	case 0x14a3:
		{
			m.ip = 0x14a6
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x14a6:
		{
			m.ip = 0x14aa
			m.r[7] = m.alu("add", 16, m.r[7], uint16(134))
			m.cycles += 4
		}
	case 0x14aa:
		{
			m.ip = 0x14ad
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x14ad:
		{
			m.ip = 0x14b0
			m.r[7] = m.alu("add", 16, m.r[7], uint16(26))
			m.cycles += 4
		}
	case 0x14b0:
		{
			m.ip = 0x14b3
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x14b3:
		{
			m.ip = 0x14b5
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(5286)
			}
			m.cycles += 17
		}
	case 0x14b5:
		{
			m.ip = 0x14b9
			m.r[7] = m.alu("add", 16, m.r[7], uint16(134))
			m.cycles += 4
		}
	case 0x14b9:
		{
			m.ip = 0x14bc
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x14bc:
		{
			m.ip = 0x14be
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x14be:
		{
			m.ip = 0x14c1
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x14c1:
		{
			m.ip = 0x14c2
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x14c2:
		{
			m.ip = 0x14c5
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x14c5:
		{
			m.ip = 0x14c7
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(5308)
			}
			m.cycles += 17
		}
	case 0x14c7:
		{
			m.ip = 0x14c8
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x14c8:
		{
			m.ip = 0x14cc
			m.r[2] = uint16(0x376)
			m.cycles += 3
		}
	case 0x14cc:
		{
			m.ip = 0x14cf
			target := uint16(17942)
			m.push(0x14cf)
			m.ip = target
			m.cycles += 19
		}
	case 0x14cf:
		{
			m.ip = 0x14d4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x3b2)), uint16(0))
			m.cycles += 16
		}
	case 0x14d4:
		{
			m.ip = 0x14d6
			if m.zf {
				m.ip = uint16(5342)
			}
			m.cycles += 8
		}
	case 0x14d6:
		{
			m.ip = 0x14da
			m.r[6] = uint16(0x12ff)
			m.cycles += 3
		}
	case 0x14da:
		{
			m.ip = 0x14dd
			target := uint16(7416)
			m.push(0x14dd)
			m.ip = target
			m.cycles += 19
		}
	case 0x14dd:
		{
			m.ip = 0x14de
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x14de:
		{
			m.ip = 0x14e0
			m.r[6] = m.alu("xor", 16, m.r[6], m.r[6])
			m.cycles += 4
		}
	case 0x14e0:
		{
			m.ip = 0x14e2
			m.r[7] = m.alu("xor", 16, m.r[7], m.r[7])
			m.cycles += 4
		}
	case 0x14e2:
		{
			m.ip = 0x14e5
			target := uint16(5425)
			m.push(0x14e5)
			m.ip = target
			m.cycles += 19
		}
	case 0x14e5:
		{
			m.ip = 0x14e8
			target := uint16(18893)
			m.push(0x14e8)
			m.ip = target
			m.cycles += 19
		}
	case 0x14e8:
		{
			m.ip = 0x14ea
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x14ea:
		{
			m.ip = 0x14ec
			if m.zf {
				m.ip = uint16(5360)
			}
			m.cycles += 8
		}
	case 0x14ec:
		{
			m.ip = 0x14ee
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x14ee:
		{
			m.ip = 0x14f0
			if !m.zf {
				m.ip = uint16(5365)
			}
			m.cycles += 8
		}
	case 0x14f0:
		{
			m.ip = 0x14f4
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x14f4:
		{
			m.ip = 0x14f5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x14f5:
		{
			m.ip = 0x14f7
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x14f7:
		{
			m.ip = 0x14f9
			if !m.zf {
				m.ip = uint16(5349)
			}
			m.cycles += 8
		}
	case 0x14f9:
		{
			m.ip = 0x14fc
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(72))
			m.cycles += 4
		}
	case 0x14fc:
		{
			m.ip = 0x14fe
			if !m.zf {
				m.ip = uint16(5377)
			}
			m.cycles += 8
		}
	case 0x14fe:
		{
			m.ip = 0x1500
			m.ip = uint16(5385)
			m.cycles += 15
		}
	case 0x1500:
		{
			m.ip = 0x1501
			m.cycles += 3
		}
	case 0x1501:
		{
			m.ip = 0x1504
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(80))
			m.cycles += 4
		}
	case 0x1504:
		{
			m.ip = 0x1506
			if !m.zf {
				m.ip = uint16(5349)
			}
			m.cycles += 8
		}
	case 0x1506:
		{
			m.ip = 0x1508
			m.ip = uint16(5414)
			m.cycles += 15
		}
	case 0x1508:
		{
			m.ip = 0x1509
			m.cycles += 3
		}
	case 0x1509:
		{
			m.ip = 0x150c
			m.alu("sub", 16, m.r[3], uint16(0))
			m.cycles += 4
		}
	case 0x150c:
		{
			m.ip = 0x150e
			if m.zf {
				m.ip = uint16(5349)
			}
			m.cycles += 8
		}
	case 0x150e:
		{
			m.ip = 0x1510
			m.r[6] = m.r[3]
			m.cycles += 2
		}
	case 0x1510:
		{
			m.ip = 0x1513
			m.r[1] = uint16(26)
			m.cycles += 2
		}
	case 0x1513:
		{
			m.ip = 0x1514
			m.r[6] = m.unary("dec", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1514:
		{
			m.ip = 0x1516
			if m.zf {
				m.ip = uint16(5409)
			}
			m.cycles += 8
		}
	case 0x1516:
		{
			m.ip = 0x151a
			m.alu("sub", 8, m.rd8(m.r[8], uint16(m.r[6])), uint16(13))
			m.cycles += 16
		}
	case 0x151a:
		{
			m.ip = 0x151c
			if !m.zf {
				m.ip = uint16(5395)
			}
			m.cycles += 8
		}
	case 0x151c:
		{
			m.ip = 0x151e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(5395)
			}
			m.cycles += 17
		}
	case 0x151e:
		{
			m.ip = 0x1521
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x1521:
		{
			m.ip = 0x1524
			target := uint16(5425)
			m.push(0x1524)
			m.ip = target
			m.cycles += 19
		}
	case 0x1524:
		{
			m.ip = 0x1526
			m.ip = uint16(5349)
			m.cycles += 15
		}
	case 0x1526:
		{
			m.ip = 0x152a
			m.alu("sub", 8, m.rd8(m.r[8], uint16(m.r[6])), uint16(26))
			m.cycles += 16
		}
	case 0x152a:
		{
			m.ip = 0x152c
			if m.zf {
				m.ip = uint16(5349)
			}
			m.cycles += 8
		}
	case 0x152c:
		{
			m.ip = 0x152f
			target := uint16(5425)
			m.push(0x152f)
			m.ip = target
			m.cycles += 19
		}
	case 0x152f:
		{
			m.ip = 0x1531
			m.ip = uint16(5349)
			m.cycles += 15
		}
	case 0x1531:
		{
			m.ip = 0x1535
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x1535:
		{
			m.ip = 0x1538
			target := uint16(3447)
			m.push(0x1538)
			m.ip = target
			m.cycles += 19
		}
	case 0x1538:
		{
			m.ip = 0x153a
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x153a:
		{
			m.ip = 0x153c
			m.r[3] = m.r[6]
			m.cycles += 2
		}
	case 0x153c:
		{
			m.ip = 0x153d
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x153d:
		{
			m.ip = 0x1541
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x1541:
		{
			m.ip = 0x1544
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x1544:
		{
			m.ip = 0x1545
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1545:
		{
			m.ip = 0x1547
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x1547:
		{
			m.ip = 0x1549
			if m.zf {
				m.ip = uint16(5461)
			}
			m.cycles += 8
		}
	case 0x1549:
		{
			m.ip = 0x154d
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x154d:
		{
			m.ip = 0x1550
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1550:
		{
			m.ip = 0x1553
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1553:
		{
			m.ip = 0x1555
			m.ip = uint16(5437)
			m.cycles += 15
		}
	case 0x1555:
		{
			m.ip = 0x1556
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1556:
		{
			m.ip = 0x1557
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x1557:
		{
			m.ip = 0x155b
			m.r[7] = m.alu("add", 16, m.r[7], uint16(160))
			m.cycles += 4
		}
	case 0x155b:
		{
			m.ip = 0x155d
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x155d:
		{
			m.ip = 0x1560
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x1560:
		{
			m.ip = 0x1561
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x1561:
		{
			m.ip = 0x1565
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x1565:
		{
			m.ip = 0x1568
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x1568:
		{
			m.ip = 0x1569
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1569:
		{
			m.ip = 0x156b
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x156b:
		{
			m.ip = 0x156d
			if m.zf {
				m.ip = uint16(5497)
			}
			m.cycles += 8
		}
	case 0x156d:
		{
			m.ip = 0x1571
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x1571:
		{
			m.ip = 0x1574
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1574:
		{
			m.ip = 0x1577
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1577:
		{
			m.ip = 0x1579
			m.ip = uint16(5473)
			m.cycles += 15
		}
	case 0x1579:
		{
			m.ip = 0x157a
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x157a:
		{
			m.ip = 0x157b
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x157b:
		{
			m.ip = 0x157f
			m.r[7] = m.alu("add", 16, m.r[7], uint16(160))
			m.cycles += 4
		}
	case 0x157f:
		{
			m.ip = 0x1581
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(5472)
			}
			m.cycles += 17
		}
	case 0x1581:
		{
			m.ip = 0x1585
			m.r[7] = uint16(0x1321)
			m.cycles += 3
		}
	case 0x1585:
		{
			m.ip = 0x1588
			m.alu("sub", 16, m.r[3], uint16(0))
			m.cycles += 4
		}
	case 0x1588:
		{
			m.ip = 0x158a
			if m.zf {
				m.ip = uint16(5528)
			}
			m.cycles += 8
		}
	case 0x158a:
		{
			m.ip = 0x158e
			m.r[7] = uint16(0x1347)
			m.cycles += 3
		}
	case 0x158e:
		{
			m.ip = 0x1592
			m.alu("sub", 8, m.rd8(m.r[8], uint16(m.r[6])), uint16(26))
			m.cycles += 16
		}
	case 0x1592:
		{
			m.ip = 0x1594
			if !m.zf {
				m.ip = uint16(5528)
			}
			m.cycles += 8
		}
	case 0x1594:
		{
			m.ip = 0x1598
			m.r[7] = uint16(0x1370)
			m.cycles += 3
		}
	case 0x1598:
		{
			m.ip = 0x1599
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x1599:
		{
			m.ip = 0x159b
			m.r[6] = m.r[7]
			m.cycles += 2
		}
	case 0x159b:
		{
			m.ip = 0x159f
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x159f:
		{
			m.ip = 0x15a2
			target := uint16(7416)
			m.push(0x15a2)
			m.ip = target
			m.cycles += 19
		}
	case 0x15a2:
		{
			m.ip = 0x15a6
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x15a6:
		{
			m.ip = 0x15a7
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x15a7:
		{
			m.ip = 0x15a9
			m.r[7] = m.alu("xor", 16, m.r[7], m.r[7])
			m.cycles += 4
		}
	case 0x15a9:
		{
			m.ip = 0x15aa
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x15aa:
		{
			m.ip = 0x15ab
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x15ab:
		{
			m.ip = 0x15ae
			m.r[2] = uint16(6656)
			m.cycles += 2
		}
	case 0x15ae:
		{
			m.ip = 0x15b1
			target := uint16(8092)
			m.push(0x15b1)
			m.ip = target
			m.cycles += 19
		}
	case 0x15b1:
		{
			m.ip = 0x15b2
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x15b2:
		{
			m.ip = 0x15b3
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x15b3:
		{
			m.ip = 0x15b4
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x15b4:
		{
			m.ip = 0x15b9
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x615)), uint16(1))
			m.cycles += 16
		}
	case 0x15b9:
		{
			m.ip = 0x15bb
			if !m.zf {
				m.ip = uint16(5581)
			}
			m.cycles += 8
		}
	case 0x15bb:
		{
			m.ip = 0x15bf
			m.r[6] = uint16(0x10c6)
			m.cycles += 3
		}
	case 0x15bf:
		{
			m.ip = 0x15c2
			target := uint16(7416)
			m.push(0x15c2)
			m.ip = target
			m.cycles += 19
		}
	case 0x15c2:
		{
			m.ip = 0x15c5
			m.r[7] = uint16(2092)
			m.cycles += 2
		}
	case 0x15c5:
		{
			m.ip = 0x15c8
			target := uint16(5985)
			m.push(0x15c8)
			m.ip = target
			m.cycles += 19
		}
	case 0x15c8:
		{
			m.ip = 0x15ca
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(89))
			m.cycles += 4
		}
	case 0x15ca:
		{
			m.ip = 0x15cc
			if m.zf {
				m.ip = uint16(5581)
			}
			m.cycles += 8
		}
	case 0x15cc:
		{
			m.ip = 0x15cd
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x15cd:
		{
			m.ip = 0x15d0
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x27d)))
			m.cycles += 8
		}
	case 0x15d0:
		{
			m.ip = 0x15d3
			m.wr16(m.r[11], uint16(0x27e), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x15d3:
		{
			m.ip = 0x15d8
			m.wr8(m.r[11], uint16(0x27d), uint16(0))
			m.cycles += 8
		}
	case 0x15d8:
		{
			m.ip = 0x15dc
			m.r[2] = uint16(0x39d)
			m.cycles += 3
		}
	case 0x15dc:
		{
			m.ip = 0x15df
			target := uint16(17966)
			m.push(0x15df)
			m.ip = target
			m.cycles += 19
		}
	case 0x15df:
		{
			m.ip = 0x15e2
			m.wr16(m.r[11], uint16(0x3ac), m.r[0])
			m.cycles += 8
		}
	case 0x15e2:
		{
			m.ip = 0x15e5
			m.r[0] = uint16(16896)
			m.cycles += 2
		}
	case 0x15e5:
		{
			m.ip = 0x15e8
			m.r[1] = uint16(0)
			m.cycles += 2
		}
	case 0x15e8:
		{
			m.ip = 0x15eb
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x15eb:
		{
			m.ip = 0x15ef
			m.r[3] = m.rd16(m.r[11], uint16(0x3ac))
			m.cycles += 8
		}
	case 0x15ef:
		{
			m.ip = 0x15f1
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x15f1:
		{
			m.ip = 0x15f5
			m.r[2] = uint16(0x26e)
			m.cycles += 3
		}
	case 0x15f5:
		{
			m.ip = 0x15f8
			target := uint16(5891)
			m.push(0x15f8)
			m.ip = target
			m.cycles += 19
		}
	case 0x15f8:
		{
			m.ip = 0x15fd
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x26e)), uint16(13))
			m.cycles += 16
		}
	case 0x15fd:
		{
			m.ip = 0x15ff
			if !m.zf {
				m.ip = uint16(5621)
			}
			m.cycles += 8
		}
	case 0x15ff:
		{
			m.ip = 0x1605
			m.wr16(m.r[11], uint16(0x1067), uint16(12345))
			m.cycles += 8
		}
	case 0x1605:
		{
			m.ip = 0x160a
			m.wr8(m.r[11], uint16(0x1069), uint16(0))
			m.cycles += 8
		}
	case 0x160a:
		{
			m.ip = 0x160f
			m.wr8(m.r[11], uint16(0x1396), uint16(1))
			m.cycles += 8
		}
	case 0x160f:
		{
			m.ip = 0x1615
			m.wr16(m.r[11], uint16(0x2093), uint16(320))
			m.cycles += 8
		}
	case 0x1615:
		{
			m.ip = 0x161b
			m.wr16(m.r[11], uint16(0x20b1), uint16(0))
			m.cycles += 8
		}
	case 0x161b:
		{
			m.ip = 0x161e
			target := uint16(5118)
			m.push(0x161e)
			m.ip = target
			m.cycles += 19
		}
	case 0x161e:
		{
			m.ip = 0x1624
			m.wr16(m.r[11], uint16(0x2091), uint16(320))
			m.cycles += 8
		}
	case 0x1624:
		{
			m.ip = 0x162a
			m.wr16(m.r[11], uint16(0x20af), uint16(0))
			m.cycles += 8
		}
	case 0x162a:
		{
			m.ip = 0x162d
			m.r[3] = uint16(4)
			m.cycles += 2
		}
	case 0x162d:
		{
			m.ip = 0x1630
			m.ip = uint16(4929)
			m.cycles += 15
		}
	case 0x1630:
		{
			m.ip = 0x1631
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1631:
		{
			m.ip = 0x1632
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x1632:
		{
			m.ip = 0x1633
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1633:
		{
			m.ip = 0x1634
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1634:
		{
			m.ip = 0x1639
			m.wr8(m.r[11], uint16(0x1396), uint16(0))
			m.cycles += 8
		}
	case 0x1639:
		{
			m.ip = 0x163c
			target := uint16(16773)
			m.push(0x163c)
			m.ip = target
			m.cycles += 19
		}
	case 0x163c:
		{
			m.ip = 0x1641
			m.wr8(m.r[11], uint16(0x1397), uint16(1))
			m.cycles += 8
		}
	case 0x1641:
		{
			m.ip = 0x1644
			target := uint16(19449)
			m.push(0x1644)
			m.ip = target
			m.cycles += 19
		}
	case 0x1644:
		{
			m.ip = 0x1649
			m.wr8(m.r[11], uint16(0x1397), uint16(0))
			m.cycles += 8
		}
	case 0x1649:
		{
			m.ip = 0x164c
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x164c:
		{
			m.ip = 0x164e
			if m.zf {
				m.ip = uint16(5733)
			}
			m.cycles += 8
		}
	case 0x164e:
		{
			m.ip = 0x1651
			target := uint16(5783)
			m.push(0x1651)
			m.ip = target
			m.cycles += 19
		}
	case 0x1651:
		{
			m.ip = 0x1657
			m.wr16(m.r[11], uint16(0x78), uint16(0))
			m.cycles += 8
		}
	case 0x1657:
		{
			m.ip = 0x165c
			m.wr8(m.r[11], uint16(0x74), uint16(1))
			m.cycles += 8
		}
	case 0x165c:
		{
			m.ip = 0x1661
			m.wr8(m.r[11], uint16(0x26b), uint16(1))
			m.cycles += 8
		}
	case 0x1661:
		{
			m.ip = 0x1664
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x1664:
		{
			m.ip = 0x1665
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1665:
		{
			m.ip = 0x1668
			m.ip = uint16(5602)
			m.cycles += 15
		}
	case 0x1668:
		{
			m.ip = 0x166d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(1))
			m.cycles += 16
		}
	case 0x166d:
		{
			m.ip = 0x166f
			if m.zf {
				m.ip = uint16(5773)
			}
			m.cycles += 8
		}
	case 0x166f:
		{
			m.ip = 0x1674
			m.wr8(m.r[11], uint16(0x1396), uint16(1))
			m.cycles += 8
		}
	case 0x1674:
		{
			m.ip = 0x1678
			m.r[2] = uint16(0x39d)
			m.cycles += 3
		}
	case 0x1678:
		{
			m.ip = 0x167a
			m.set8(0, 8, uint16(60))
			m.cycles += 2
		}
	case 0x167a:
		{
			m.ip = 0x167c
			m.set8(1, 0, uint16(32))
			m.cycles += 2
		}
	case 0x167c:
		{
			m.ip = 0x167e
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x167e:
		{
			m.ip = 0x1681
			m.wr16(m.r[11], uint16(0x3ac), m.r[0])
			m.cycles += 8
		}
	case 0x1681:
		{
			m.ip = 0x1687
			m.wr16(m.r[11], uint16(0x1067), uint16(12345))
			m.cycles += 8
		}
	case 0x1687:
		{
			m.ip = 0x168c
			m.wr8(m.r[11], uint16(0x1069), uint16(0))
			m.cycles += 8
		}
	case 0x168c:
		{
			m.ip = 0x168d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x168d:
		{
			m.ip = 0x1690
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x1690:
		{
			m.ip = 0x1692
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x1692:
		{
			m.ip = 0x1694
			m.set8(0, 8, uint16(76))
			m.cycles += 2
		}
	case 0x1694:
		{
			m.ip = 0x1696
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x1696:
		{
			m.ip = 0x1697
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1697:
		{
			m.ip = 0x169c
			m.wr8(m.r[11], uint16(0x1396), uint16(0))
			m.cycles += 8
		}
	case 0x169c:
		{
			m.ip = 0x169d
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x169d:
		{
			m.ip = 0x169e
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x169e:
		{
			m.ip = 0x16a2
			m.r[3] = m.rd16(m.r[11], uint16(0x3ac))
			m.cycles += 8
		}
	case 0x16a2:
		{
			m.ip = 0x16a4
			m.set8(0, 8, uint16(62))
			m.cycles += 2
		}
	case 0x16a4:
		{
			m.ip = 0x16a6
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x16a6:
		{
			m.ip = 0x16ac
			m.wr16(m.r[11], uint16(0x1398), uint16(0))
			m.cycles += 8
		}
	case 0x16ac:
		{
			m.ip = 0x16af
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x27e)))
			m.cycles += 8
		}
	case 0x16af:
		{
			m.ip = 0x16b2
			m.wr16(m.r[11], uint16(0x27d), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x16b2:
		{
			m.ip = 0x16b3
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x16b3:
		{
			m.ip = 0x16b4
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x16b4:
		{
			m.ip = 0x16b5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x16b5:
		{
			m.ip = 0x16b8
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x16b8:
		{
			m.ip = 0x16ba
			if m.zf {
				m.ip = uint16(5841)
			}
			m.cycles += 8
		}
	case 0x16ba:
		{
			m.ip = 0x16bd
			target := uint16(5783)
			m.push(0x16bd)
			m.ip = target
			m.cycles += 19
		}
	case 0x16bd:
		{
			m.ip = 0x16c3
			m.wr16(m.r[11], uint16(0x78), uint16(0))
			m.cycles += 8
		}
	case 0x16c3:
		{
			m.ip = 0x16c8
			m.wr8(m.r[11], uint16(0x74), uint16(1))
			m.cycles += 8
		}
	case 0x16c8:
		{
			m.ip = 0x16cd
			m.wr8(m.r[11], uint16(0x26b), uint16(1))
			m.cycles += 8
		}
	case 0x16cd:
		{
			m.ip = 0x16d0
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x16d0:
		{
			m.ip = 0x16d1
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x16d1:
		{
			m.ip = 0x16d2
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x16d2:
		{
			m.ip = 0x16d3
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x16d3:
		{
			m.ip = 0x16d4
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x16d4:
		{
			m.ip = 0x16d9
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x26c)), uint16(0))
			m.cycles += 16
		}
	case 0x16d9:
		{
			m.ip = 0x16db
			if !m.zf {
				m.ip = uint16(5873)
			}
			m.cycles += 8
		}
	case 0x16db:
		{
			m.ip = 0x16df
			m.r[2] = uint16(0x26e)
			m.cycles += 3
		}
	case 0x16df:
		{
			m.ip = 0x16e2
			target := uint16(5891)
			m.push(0x16e2)
			m.ip = target
			m.cycles += 19
		}
	case 0x16e2:
		{
			m.ip = 0x16e5
			m.r[0] = m.rd16(m.r[11], uint16(0x26e))
			m.cycles += 8
		}
	case 0x16e5:
		{
			m.ip = 0x16e8
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x16e8:
		{
			m.ip = 0x16ea
			if !m.zf {
				m.ip = uint16(5880)
			}
			m.cycles += 8
		}
	case 0x16ea:
		{
			m.ip = 0x16ee
			m.r[2] = uint16(0x26c)
			m.cycles += 3
		}
	case 0x16ee:
		{
			m.ip = 0x16f1
			target := uint16(5891)
			m.push(0x16f1)
			m.ip = target
			m.cycles += 19
		}
	case 0x16f1:
		{
			m.ip = 0x16f5
			m.wr16(m.r[11], uint16(0x26c), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x26c))))
			m.cycles += 15
		}
	case 0x16f5:
		{
			m.ip = 0x16f8
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x16f8:
		{
			m.ip = 0x16fa
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x16fa:
		{
			m.ip = 0x16fc
			if !m.zf {
				m.ip = uint16(5887)
			}
			m.cycles += 8
		}
	case 0x16fc:
		{
			m.ip = 0x16ff
			m.ip = uint16(5681)
			m.cycles += 15
		}
	case 0x16ff:
		{
			m.ip = 0x1700
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x1700:
		{
			m.ip = 0x1701
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1701:
		{
			m.ip = 0x1702
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1702:
		{
			m.ip = 0x1703
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1703:
		{
			m.ip = 0x1705
			m.set8(0, 8, uint16(63))
			m.cycles += 2
		}
	case 0x1705:
		{
			m.ip = 0x1709
			m.r[3] = m.rd16(m.r[11], uint16(0x3ac))
			m.cycles += 8
		}
	case 0x1709:
		{
			m.ip = 0x170c
			m.r[1] = uint16(2)
			m.cycles += 2
		}
	case 0x170c:
		{
			m.ip = 0x170e
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x170e:
		{
			m.ip = 0x170f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x170f:
		{
			m.ip = 0x1710
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x1710:
		{
			m.ip = 0x1711
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x1711:
		{
			m.ip = 0x1712
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x1712:
		{
			m.ip = 0x1715
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x1715:
		{
			m.ip = 0x1717
			if m.zf {
				m.ip = uint16(5951)
			}
			m.cycles += 8
		}
	case 0x1717:
		{
			m.ip = 0x171c
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x26c)), uint16(0))
			m.cycles += 16
		}
	case 0x171c:
		{
			m.ip = 0x171e
			if m.zf {
				m.ip = uint16(5932)
			}
			m.cycles += 8
		}
	case 0x171e:
		{
			m.ip = 0x1722
			m.r[2] = uint16(0x26e)
			m.cycles += 3
		}
	case 0x1722:
		{
			m.ip = 0x1725
			target := uint16(5971)
			m.push(0x1725)
			m.ip = target
			m.cycles += 19
		}
	case 0x1725:
		{
			m.ip = 0x1729
			m.r[2] = uint16(0x26c)
			m.cycles += 3
		}
	case 0x1729:
		{
			m.ip = 0x172c
			target := uint16(5971)
			m.push(0x172c)
			m.ip = target
			m.cycles += 19
		}
	case 0x172c:
		{
			m.ip = 0x172f
			m.wr16(m.r[11], uint16(0x26c), m.r[0])
			m.cycles += 8
		}
	case 0x172f:
		{
			m.ip = 0x1733
			m.r[2] = uint16(0x26c)
			m.cycles += 3
		}
	case 0x1733:
		{
			m.ip = 0x1736
			target := uint16(5971)
			m.push(0x1736)
			m.ip = target
			m.cycles += 19
		}
	case 0x1736:
		{
			m.ip = 0x173c
			m.wr16(m.r[11], uint16(0x26c), uint16(0))
			m.cycles += 8
		}
	case 0x173c:
		{
			m.ip = 0x173e
			m.ip = uint16(5967)
			m.cycles += 15
		}
	case 0x173e:
		{
			m.ip = 0x173f
			m.cycles += 3
		}
	case 0x173f:
		{
			m.ip = 0x1743
			m.wr16(m.r[11], uint16(0x26c), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x26c))))
			m.cycles += 15
		}
	case 0x1743:
		{
			m.ip = 0x1745
			if !m.cf {
				m.ip = uint16(5967)
			}
			m.cycles += 8
		}
	case 0x1745:
		{
			m.ip = 0x1749
			m.r[2] = uint16(0x26e)
			m.cycles += 3
		}
	case 0x1749:
		{
			m.ip = 0x174c
			target := uint16(5971)
			m.push(0x174c)
			m.ip = target
			m.cycles += 19
		}
	case 0x174c:
		{
			m.ip = 0x174f
			target := uint16(5971)
			m.push(0x174f)
			m.ip = target
			m.cycles += 19
		}
	case 0x174f:
		{
			m.ip = 0x1750
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x1750:
		{
			m.ip = 0x1751
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1751:
		{
			m.ip = 0x1752
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1752:
		{
			m.ip = 0x1753
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1753:
		{
			m.ip = 0x1754
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x1754:
		{
			m.ip = 0x1756
			m.set8(0, 8, uint16(64))
			m.cycles += 2
		}
	case 0x1756:
		{
			m.ip = 0x175a
			m.r[3] = m.rd16(m.r[11], uint16(0x3ac))
			m.cycles += 8
		}
	case 0x175a:
		{
			m.ip = 0x175d
			m.r[1] = uint16(2)
			m.cycles += 2
		}
	case 0x175d:
		{
			m.ip = 0x175f
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x175f:
		{
			m.ip = 0x1760
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x1760:
		{
			m.ip = 0x1761
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1761:
		{
			m.ip = 0x1765
			m.wr16(m.r[11], uint16(0x297), m.r[7])
			m.cycles += 8
		}
	case 0x1765:
		{
			m.ip = 0x1768
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x1768:
		{
			m.ip = 0x176b
			m.r[1] = uint16(4)
			m.cycles += 2
		}
	case 0x176b:
		{
			m.ip = 0x176c
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x176c:
		{
			m.ip = 0x176f
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x176f:
		{
			m.ip = 0x1772
			m.r[0] = m.rd16(m.r[8], uint16(m.r[7]))
			m.cycles += 8
		}
	case 0x1772:
		{
			m.ip = 0x1776
			m.wr16(m.r[11], uint16(m.r[6]+0x1072), m.r[0])
			m.cycles += 8
		}
	case 0x1776:
		{
			m.ip = 0x1779
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x1779:
		{
			m.ip = 0x177c
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x177c:
		{
			m.ip = 0x177e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(5999)
			}
			m.cycles += 17
		}
	case 0x177e:
		{
			m.ip = 0x1782
			m.r[7] = m.alu("add", 16, m.r[7], uint16(146))
			m.cycles += 4
		}
	case 0x1782:
		{
			m.ip = 0x1783
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1783:
		{
			m.ip = 0x1785
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(5995)
			}
			m.cycles += 17
		}
	case 0x1785:
		{
			m.ip = 0x1787
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x1787:
		{
			m.ip = 0x178a
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x178a:
		{
			m.ip = 0x178e
			m.r[7] = m.rd16(m.r[11], uint16(0x297))
			m.cycles += 8
		}
	case 0x178e:
		{
			m.ip = 0x1791
			m.r[1] = uint16(4)
			m.cycles += 2
		}
	case 0x1791:
		{
			m.ip = 0x1792
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x1792:
		{
			m.ip = 0x1795
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x1795:
		{
			m.ip = 0x1799
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x10aa)))
			m.cycles += 8
		}
	case 0x1799:
		{
			m.ip = 0x179c
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x179c:
		{
			m.ip = 0x179d
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x179d:
		{
			m.ip = 0x17a0
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x17a0:
		{
			m.ip = 0x17a2
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6037)
			}
			m.cycles += 17
		}
	case 0x17a2:
		{
			m.ip = 0x17a6
			m.r[7] = m.alu("add", 16, m.r[7], uint16(146))
			m.cycles += 4
		}
	case 0x17a6:
		{
			m.ip = 0x17a7
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x17a7:
		{
			m.ip = 0x17a9
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6033)
			}
			m.cycles += 17
		}
	case 0x17a9:
		{
			m.ip = 0x17ab
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x17ab:
		{
			m.ip = 0x17af
			m.r[7] = m.rd16(m.r[11], uint16(0x297))
			m.cycles += 8
		}
	case 0x17af:
		{
			m.ip = 0x17b4
			m.wr8(m.r[8], uint16(m.r[7]+0xa5), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x17b4:
		{
			m.ip = 0x17b9
			m.wr8(m.r[8], uint16(m.r[7]+0xa7), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x17b9:
		{
			m.ip = 0x17be
			m.wr8(m.r[8], uint16(m.r[7]+0xa9), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x17be:
		{
			m.ip = 0x17c1
			target := uint16(18893)
			m.push(0x17c1)
			m.ip = target
			m.cycles += 19
		}
	case 0x17c1:
		{
			m.ip = 0x17c3
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x17c3:
		{
			m.ip = 0x17c5
			if m.zf {
				m.ip = uint16(6111)
			}
			m.cycles += 8
		}
	case 0x17c5:
		{
			m.ip = 0x17c7
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x17c7:
		{
			m.ip = 0x17c9
			if m.zf {
				m.ip = uint16(6184)
			}
			m.cycles += 8
		}
	case 0x17c9:
		{
			m.ip = 0x17cb
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x17cb:
		{
			m.ip = 0x17cd
			if m.zf {
				m.ip = uint16(6189)
			}
			m.cycles += 8
		}
	case 0x17cd:
		{
			m.ip = 0x17cf
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(89))
			m.cycles += 4
		}
	case 0x17cf:
		{
			m.ip = 0x17d1
			if m.zf {
				m.ip = uint16(6123)
			}
			m.cycles += 8
		}
	case 0x17d1:
		{
			m.ip = 0x17d3
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(121))
			m.cycles += 4
		}
	case 0x17d3:
		{
			m.ip = 0x17d5
			if m.zf {
				m.ip = uint16(6123)
			}
			m.cycles += 8
		}
	case 0x17d5:
		{
			m.ip = 0x17d7
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(78))
			m.cycles += 4
		}
	case 0x17d7:
		{
			m.ip = 0x17d9
			if m.zf {
				m.ip = uint16(6138)
			}
			m.cycles += 8
		}
	case 0x17d9:
		{
			m.ip = 0x17db
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(110))
			m.cycles += 4
		}
	case 0x17db:
		{
			m.ip = 0x17dd
			if m.zf {
				m.ip = uint16(6138)
			}
			m.cycles += 8
		}
	case 0x17dd:
		{
			m.ip = 0x17df
			m.ip = uint16(6078)
			m.cycles += 15
		}
	case 0x17df:
		{
			m.ip = 0x17e2
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(72))
			m.cycles += 4
		}
	case 0x17e2:
		{
			m.ip = 0x17e4
			if m.zf {
				m.ip = uint16(6123)
			}
			m.cycles += 8
		}
	case 0x17e4:
		{
			m.ip = 0x17e7
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(80))
			m.cycles += 4
		}
	case 0x17e7:
		{
			m.ip = 0x17e9
			if m.zf {
				m.ip = uint16(6138)
			}
			m.cycles += 8
		}
	case 0x17e9:
		{
			m.ip = 0x17eb
			m.ip = uint16(6078)
			m.cycles += 15
		}
	case 0x17eb:
		{
			m.ip = 0x17f0
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1071)), uint16(89))
			m.cycles += 16
		}
	case 0x17f0:
		{
			m.ip = 0x17f2
			if m.zf {
				m.ip = uint16(6078)
			}
			m.cycles += 8
		}
	case 0x17f2:
		{
			m.ip = 0x17f7
			m.wr8(m.r[11], uint16(0x1071), uint16(89))
			m.cycles += 8
		}
	case 0x17f7:
		{
			m.ip = 0x17f9
			m.ip = uint16(6150)
			m.cycles += 15
		}
	case 0x17f9:
		{
			m.ip = 0x17fa
			m.cycles += 3
		}
	case 0x17fa:
		{
			m.ip = 0x17ff
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1071)), uint16(78))
			m.cycles += 16
		}
	case 0x17ff:
		{
			m.ip = 0x1801
			if m.zf {
				m.ip = uint16(6078)
			}
			m.cycles += 8
		}
	case 0x1801:
		{
			m.ip = 0x1806
			m.wr8(m.r[11], uint16(0x1071), uint16(78))
			m.cycles += 8
		}
	case 0x1806:
		{
			m.ip = 0x180a
			m.r[7] = m.rd16(m.r[11], uint16(0x297))
			m.cycles += 8
		}
	case 0x180a:
		{
			m.ip = 0x180d
			m.r[1] = uint16(3)
			m.cycles += 2
		}
	case 0x180d:
		{
			m.ip = 0x1812
			m.set8(0, 8, m.rd8(m.r[8], uint16(m.r[7]+0xa5)))
			m.cycles += 8
		}
	case 0x1812:
		{
			m.ip = 0x1817
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[7]+0x145)))
			m.cycles += 8
		}
	case 0x1817:
		{
			m.ip = 0x181c
			m.wr8(m.r[8], uint16(m.r[7]+0xa5), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x181c:
		{
			m.ip = 0x1821
			m.wr8(m.r[8], uint16(m.r[7]+0x145), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x1821:
		{
			m.ip = 0x1824
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1824:
		{
			m.ip = 0x1826
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6157)
			}
			m.cycles += 17
		}
	case 0x1826:
		{
			m.ip = 0x1828
			m.ip = uint16(6078)
			m.cycles += 15
		}
	case 0x1828:
		{
			m.ip = 0x182d
			m.wr8(m.r[11], uint16(0x1071), uint16(69))
			m.cycles += 8
		}
	case 0x182d:
		{
			m.ip = 0x1830
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x1071)))
			m.cycles += 8
		}
	case 0x1830:
		{
			m.ip = 0x1831
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x1831:
		{
			m.ip = 0x1836
			m.wr8(m.r[11], uint16(0x1071), uint16(89))
			m.cycles += 8
		}
	case 0x1836:
		{
			m.ip = 0x1839
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x1839:
		{
			m.ip = 0x183d
			m.r[7] = m.rd16(m.r[11], uint16(0x297))
			m.cycles += 8
		}
	case 0x183d:
		{
			m.ip = 0x1840
			m.r[1] = uint16(4)
			m.cycles += 2
		}
	case 0x1840:
		{
			m.ip = 0x1841
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x1841:
		{
			m.ip = 0x1844
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x1844:
		{
			m.ip = 0x1848
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]+0x1072))
			m.cycles += 8
		}
	case 0x1848:
		{
			m.ip = 0x184b
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x184b:
		{
			m.ip = 0x184e
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x184e:
		{
			m.ip = 0x1851
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1851:
		{
			m.ip = 0x1853
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6212)
			}
			m.cycles += 17
		}
	case 0x1853:
		{
			m.ip = 0x1857
			m.r[7] = m.alu("add", 16, m.r[7], uint16(146))
			m.cycles += 4
		}
	case 0x1857:
		{
			m.ip = 0x1858
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1858:
		{
			m.ip = 0x185a
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6208)
			}
			m.cycles += 17
		}
	case 0x185a:
		{
			m.ip = 0x185b
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x185b:
		{
			m.ip = 0x185c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x185c:
		{
			m.ip = 0x185e
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(69))
			m.cycles += 4
		}
	case 0x185e:
		{
			m.ip = 0x1860
			if !m.zf {
				m.ip = uint16(6241)
			}
			m.cycles += 8
		}
	case 0x1860:
		{
			m.ip = 0x1861
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1861:
		{
			m.ip = 0x1863
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(75))
			m.cycles += 4
		}
	case 0x1863:
		{
			m.ip = 0x1865
			if !m.zf {
				m.ip = uint16(6256)
			}
			m.cycles += 8
		}
	case 0x1865:
		{
			m.ip = 0x186a
			m.wr8(m.r[11], uint16(0x27d), uint16(0))
			m.cycles += 8
		}
	case 0x186a:
		{
			m.ip = 0x186f
			m.wr8(m.r[11], uint16(0x110f), uint16(75))
			m.cycles += 8
		}
	case 0x186f:
		{
			m.ip = 0x1870
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1870:
		{
			m.ip = 0x1873
			m.r[2] = uint16(513)
			m.cycles += 2
		}
	case 0x1873:
		{
			m.ip = 0x1875
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x1875:
		{
			m.ip = 0x1876
			m.output(m.r[2], ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x1876:
		{
			m.ip = 0x1879
			m.r[1] = uint16(10000)
			m.cycles += 2
		}
	case 0x1879:
		{
			m.ip = 0x187a
			m.set8(0, 0, m.input(m.r[2]))
			m.cycles += 8
		}
	case 0x187a:
		{
			m.ip = 0x187c
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6265)
			}
			m.cycles += 17
		}
	case 0x187c:
		{
			m.ip = 0x187e
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(255))
			m.cycles += 4
		}
	case 0x187e:
		{
			m.ip = 0x1880
			if !m.zf {
				m.ip = uint16(6300)
			}
			m.cycles += 8
		}
	case 0x1880:
		{
			m.ip = 0x1884
			m.r[6] = uint16(0x1237)
			m.cycles += 3
		}
	case 0x1884:
		{
			m.ip = 0x1887
			target := uint16(7416)
			m.push(0x1887)
			m.ip = target
			m.cycles += 19
		}
	case 0x1887:
		{
			m.ip = 0x188a
			m.r[7] = uint16(2804)
			m.cycles += 2
		}
	case 0x188a:
		{
			m.ip = 0x188d
			target := uint16(5985)
			m.push(0x188d)
			m.ip = target
			m.cycles += 19
		}
	case 0x188d:
		{
			m.ip = 0x188f
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(89))
			m.cycles += 4
		}
	case 0x188f:
		{
			m.ip = 0x1891
			if m.zf {
				m.ip = uint16(6300)
			}
			m.cycles += 8
		}
	case 0x1891:
		{
			m.ip = 0x1896
			m.wr8(m.r[11], uint16(0x27d), uint16(0))
			m.cycles += 8
		}
	case 0x1896:
		{
			m.ip = 0x189b
			m.wr8(m.r[11], uint16(0x110f), uint16(75))
			m.cycles += 8
		}
	case 0x189b:
		{
			m.ip = 0x189c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x189c:
		{
			m.ip = 0x18a0
			m.r[6] = uint16(0x1278)
			m.cycles += 3
		}
	case 0x18a0:
		{
			m.ip = 0x18a3
			target := uint16(7416)
			m.push(0x18a3)
			m.ip = target
			m.cycles += 19
		}
	case 0x18a3:
		{
			m.ip = 0x18a6
			target := uint16(18893)
			m.push(0x18a6)
			m.ip = target
			m.cycles += 19
		}
	case 0x18a6:
		{
			m.ip = 0x18a8
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x18a8:
		{
			m.ip = 0x18aa
			if m.zf {
				m.ip = uint16(6289)
			}
			m.cycles += 8
		}
	case 0x18aa:
		{
			m.ip = 0x18ac
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x18ac:
		{
			m.ip = 0x18ae
			if !m.zf {
				m.ip = uint16(6307)
			}
			m.cycles += 8
		}
	case 0x18ae:
		{
			m.ip = 0x18b0
			m.r[0] = m.alu("xor", 16, m.r[0], m.r[0])
			m.cycles += 4
		}
	case 0x18b0:
		{
			m.ip = 0x18b2
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x18b2:
		{
			m.ip = 0x18b5
			target := uint16(6344)
			m.push(0x18b5)
			m.ip = target
			m.cycles += 19
		}
	case 0x18b5:
		{
			m.ip = 0x18b9
			m.wr16(m.r[11], uint16(0x280), m.r[1])
			m.cycles += 8
		}
	case 0x18b9:
		{
			m.ip = 0x18bb
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x18bb:
		{
			m.ip = 0x18be
			target := uint16(6344)
			m.push(0x18be)
			m.ip = target
			m.cycles += 19
		}
	case 0x18be:
		{
			m.ip = 0x18c2
			m.wr16(m.r[11], uint16(0x282), m.r[1])
			m.cycles += 8
		}
	case 0x18c2:
		{
			m.ip = 0x18c7
			m.wr8(m.r[11], uint16(0x27d), uint16(1))
			m.cycles += 8
		}
	case 0x18c7:
		{
			m.ip = 0x18c8
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x18c8:
		{
			m.ip = 0x18cb
			m.r[2] = uint16(513)
			m.cycles += 2
		}
	case 0x18cb:
		{
			m.ip = 0x18ce
			m.r[1] = uint16(1024)
			m.cycles += 2
		}
	case 0x18ce:
		{
			m.ip = 0x18d0
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6350)
			}
			m.cycles += 17
		}
	case 0x18d0:
		{
			m.ip = 0x18d2
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x18d2:
		{
			m.ip = 0x18d3
			m.output(m.r[2], ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x18d3:
		{
			m.ip = 0x18d4
			m.set8(0, 0, m.input(m.r[2]))
			m.cycles += 8
		}
	case 0x18d4:
		{
			m.ip = 0x18d6
			m.alu("and", 8, ((m.r[0] >> 0) & 255), ((m.r[0] >> 8) & 255))
			m.cycles += 4
		}
	case 0x18d6:
		{
			m.ip = 0x18d8
			m.r[1]--
			if m.r[1] != 0 && !m.zf {
				m.ip = uint16(6355)
			}
			m.cycles += 17
		}
	case 0x18d8:
		{
			m.ip = 0x18da
			m.r[1] = m.unary("neg", 16, m.r[1])
			m.cycles += 3
		}
	case 0x18da:
		{
			m.ip = 0x18db
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x18db:
		{
			m.ip = 0x18de
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x110f)))
			m.cycles += 8
		}
	case 0x18de:
		{
			m.ip = 0x18e1
			m.wr16(m.r[11], uint16(0x1110), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x18e1:
		{
			m.ip = 0x18e5
			m.r[7] = m.rd16(m.r[11], uint16(0x110d))
			m.cycles += 8
		}
	case 0x18e5:
		{
			m.ip = 0x18e8
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x18e8:
		{
			m.ip = 0x18eb
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x18eb:
		{
			m.ip = 0x18ec
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x18ec:
		{
			m.ip = 0x18ef
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x18ef:
		{
			m.ip = 0x18f2
			m.r[0] = m.rd16(m.r[8], uint16(m.r[7]))
			m.cycles += 8
		}
	case 0x18f2:
		{
			m.ip = 0x18f6
			m.wr16(m.r[11], uint16(m.r[6]+0x1111), m.r[0])
			m.cycles += 8
		}
	case 0x18f6:
		{
			m.ip = 0x18f9
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x18f9:
		{
			m.ip = 0x18fc
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x18fc:
		{
			m.ip = 0x18fe
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6383)
			}
			m.cycles += 17
		}
	case 0x18fe:
		{
			m.ip = 0x1902
			m.r[7] = m.alu("add", 16, m.r[7], uint16(132))
			m.cycles += 4
		}
	case 0x1902:
		{
			m.ip = 0x1903
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1903:
		{
			m.ip = 0x1905
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6379)
			}
			m.cycles += 17
		}
	case 0x1905:
		{
			m.ip = 0x1908
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x1908:
		{
			m.ip = 0x190c
			m.r[7] = m.rd16(m.r[11], uint16(0x110d))
			m.cycles += 8
		}
	case 0x190c:
		{
			m.ip = 0x190f
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x190f:
		{
			m.ip = 0x1910
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x1910:
		{
			m.ip = 0x1913
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x1913:
		{
			m.ip = 0x1917
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x11d5)))
			m.cycles += 8
		}
	case 0x1917:
		{
			m.ip = 0x1919
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x1919:
		{
			m.ip = 0x191b
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(123))
			m.cycles += 4
		}
	case 0x191b:
		{
			m.ip = 0x191d
			if !m.cf && !m.zf {
				m.ip = uint16(6431)
			}
			m.cycles += 8
		}
	case 0x191d:
		{
			m.ip = 0x191f
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x191f:
		{
			m.ip = 0x1922
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1922:
		{
			m.ip = 0x1923
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1923:
		{
			m.ip = 0x1926
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1926:
		{
			m.ip = 0x1928
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6419)
			}
			m.cycles += 17
		}
	case 0x1928:
		{
			m.ip = 0x192c
			m.r[7] = m.alu("add", 16, m.r[7], uint16(132))
			m.cycles += 4
		}
	case 0x192c:
		{
			m.ip = 0x192d
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x192d:
		{
			m.ip = 0x192f
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6415)
			}
			m.cycles += 17
		}
	case 0x192f:
		{
			m.ip = 0x1932
			m.r[1] = uint16(8)
			m.cycles += 2
		}
	case 0x1932:
		{
			m.ip = 0x1936
			m.r[7] = m.rd16(m.r[11], uint16(0x110d))
			m.cycles += 8
		}
	case 0x1936:
		{
			m.ip = 0x193a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(487))
			m.cycles += 4
		}
	case 0x193a:
		{
			m.ip = 0x193f
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x110f)), uint16(74))
			m.cycles += 16
		}
	case 0x193f:
		{
			m.ip = 0x1941
			if m.zf {
				m.ip = uint16(6469)
			}
			m.cycles += 8
		}
	case 0x1941:
		{
			m.ip = 0x1945
			m.r[7] = m.alu("add", 16, m.r[7], uint16(320))
			m.cycles += 4
		}
	case 0x1945:
		{
			m.ip = 0x1948
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x1948:
		{
			m.ip = 0x194b
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x194b:
		{
			m.ip = 0x194d
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6469)
			}
			m.cycles += 17
		}
	case 0x194d:
		{
			m.ip = 0x1951
			m.r[6] = uint16(0x12af)
			m.cycles += 3
		}
	case 0x1951:
		{
			m.ip = 0x1954
			target := uint16(7416)
			m.push(0x1954)
			m.ip = target
			m.cycles += 19
		}
	case 0x1954:
		{
			m.ip = 0x1957
			target := uint16(18893)
			m.push(0x1957)
			m.ip = target
			m.cycles += 19
		}
	case 0x1957:
		{
			m.ip = 0x1959
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x1959:
		{
			m.ip = 0x195b
			if m.zf {
				m.ip = uint16(6517)
			}
			m.cycles += 8
		}
	case 0x195b:
		{
			m.ip = 0x195d
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x195d:
		{
			m.ip = 0x195f
			if m.zf {
				m.ip = uint16(6604)
			}
			m.cycles += 8
		}
	case 0x195f:
		{
			m.ip = 0x1961
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x1961:
		{
			m.ip = 0x1963
			if m.zf {
				m.ip = uint16(6615)
			}
			m.cycles += 8
		}
	case 0x1963:
		{
			m.ip = 0x1965
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(75))
			m.cycles += 4
		}
	case 0x1965:
		{
			m.ip = 0x1967
			if m.zf {
				m.ip = uint16(6529)
			}
			m.cycles += 8
		}
	case 0x1967:
		{
			m.ip = 0x1969
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(107))
			m.cycles += 4
		}
	case 0x1969:
		{
			m.ip = 0x196b
			if m.zf {
				m.ip = uint16(6529)
			}
			m.cycles += 8
		}
	case 0x196b:
		{
			m.ip = 0x196d
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(74))
			m.cycles += 4
		}
	case 0x196d:
		{
			m.ip = 0x196f
			if m.zf {
				m.ip = uint16(6551)
			}
			m.cycles += 8
		}
	case 0x196f:
		{
			m.ip = 0x1971
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(106))
			m.cycles += 4
		}
	case 0x1971:
		{
			m.ip = 0x1973
			if m.zf {
				m.ip = uint16(6551)
			}
			m.cycles += 8
		}
	case 0x1973:
		{
			m.ip = 0x1975
			m.ip = uint16(6484)
			m.cycles += 15
		}
	case 0x1975:
		{
			m.ip = 0x1978
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(72))
			m.cycles += 4
		}
	case 0x1978:
		{
			m.ip = 0x197a
			if m.zf {
				m.ip = uint16(6529)
			}
			m.cycles += 8
		}
	case 0x197a:
		{
			m.ip = 0x197d
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(80))
			m.cycles += 4
		}
	case 0x197d:
		{
			m.ip = 0x197f
			if m.zf {
				m.ip = uint16(6551)
			}
			m.cycles += 8
		}
	case 0x197f:
		{
			m.ip = 0x1981
			m.ip = uint16(6484)
			m.cycles += 15
		}
	case 0x1981:
		{
			m.ip = 0x1986
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x110f)), uint16(75))
			m.cycles += 16
		}
	case 0x1986:
		{
			m.ip = 0x1988
			if m.zf {
				m.ip = uint16(6484)
			}
			m.cycles += 8
		}
	case 0x1988:
		{
			m.ip = 0x198d
			m.wr8(m.r[11], uint16(0x110f), uint16(75))
			m.cycles += 8
		}
	case 0x198d:
		{
			m.ip = 0x1991
			m.r[6] = uint16(0x12af)
			m.cycles += 3
		}
	case 0x1991:
		{
			m.ip = 0x1994
			target := uint16(7416)
			m.push(0x1994)
			m.ip = target
			m.cycles += 19
		}
	case 0x1994:
		{
			m.ip = 0x1996
			m.ip = uint16(6570)
			m.cycles += 15
		}
	case 0x1996:
		{
			m.ip = 0x1997
			m.cycles += 3
		}
	case 0x1997:
		{
			m.ip = 0x199c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x110f)), uint16(74))
			m.cycles += 16
		}
	case 0x199c:
		{
			m.ip = 0x199e
			if m.zf {
				m.ip = uint16(6484)
			}
			m.cycles += 8
		}
	case 0x199e:
		{
			m.ip = 0x19a3
			m.wr8(m.r[11], uint16(0x110f), uint16(74))
			m.cycles += 8
		}
	case 0x19a3:
		{
			m.ip = 0x19a7
			m.r[6] = uint16(0x12d7)
			m.cycles += 3
		}
	case 0x19a7:
		{
			m.ip = 0x19aa
			target := uint16(7416)
			m.push(0x19aa)
			m.ip = target
			m.cycles += 19
		}
	case 0x19aa:
		{
			m.ip = 0x19ae
			m.r[7] = m.rd16(m.r[11], uint16(0x110d))
			m.cycles += 8
		}
	case 0x19ae:
		{
			m.ip = 0x19b1
			m.r[1] = uint16(8)
			m.cycles += 2
		}
	case 0x19b1:
		{
			m.ip = 0x19b6
			m.set8(0, 8, m.rd8(m.r[8], uint16(m.r[7]+0x1e7)))
			m.cycles += 8
		}
	case 0x19b6:
		{
			m.ip = 0x19bb
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[7]+0x327)))
			m.cycles += 8
		}
	case 0x19bb:
		{
			m.ip = 0x19c0
			m.wr8(m.r[8], uint16(m.r[7]+0x1e7), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x19c0:
		{
			m.ip = 0x19c5
			m.wr8(m.r[8], uint16(m.r[7]+0x327), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x19c5:
		{
			m.ip = 0x19c8
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x19c8:
		{
			m.ip = 0x19ca
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6577)
			}
			m.cycles += 17
		}
	case 0x19ca:
		{
			m.ip = 0x19cc
			m.ip = uint16(6484)
			m.cycles += 15
		}
	case 0x19cc:
		{
			m.ip = 0x19cf
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x1110)))
			m.cycles += 8
		}
	case 0x19cf:
		{
			m.ip = 0x19d2
			m.wr16(m.r[11], uint16(0x110f), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x19d2:
		{
			m.ip = 0x19d4
			m.set8(0, 0, uint16(69))
			m.cycles += 2
		}
	case 0x19d4:
		{
			m.ip = 0x19d6
			m.ip = uint16(6618)
			m.cycles += 15
		}
	case 0x19d6:
		{
			m.ip = 0x19d7
			m.cycles += 3
		}
	case 0x19d7:
		{
			m.ip = 0x19da
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x110f)))
			m.cycles += 8
		}
	case 0x19da:
		{
			m.ip = 0x19db
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x19db:
		{
			m.ip = 0x19de
			target := uint16(6236)
			m.push(0x19de)
			m.ip = target
			m.cycles += 19
		}
	case 0x19de:
		{
			m.ip = 0x19e1
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x19e1:
		{
			m.ip = 0x19e5
			m.r[7] = m.rd16(m.r[11], uint16(0x110d))
			m.cycles += 8
		}
	case 0x19e5:
		{
			m.ip = 0x19e8
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x19e8:
		{
			m.ip = 0x19e9
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x19e9:
		{
			m.ip = 0x19ec
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x19ec:
		{
			m.ip = 0x19f0
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]+0x1111))
			m.cycles += 8
		}
	case 0x19f0:
		{
			m.ip = 0x19f3
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x19f3:
		{
			m.ip = 0x19f6
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x19f6:
		{
			m.ip = 0x19f9
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x19f9:
		{
			m.ip = 0x19fb
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6636)
			}
			m.cycles += 17
		}
	case 0x19fb:
		{
			m.ip = 0x19ff
			m.r[7] = m.alu("add", 16, m.r[7], uint16(132))
			m.cycles += 4
		}
	case 0x19ff:
		{
			m.ip = 0x1a00
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1a00:
		{
			m.ip = 0x1a02
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(6632)
			}
			m.cycles += 17
		}
	case 0x1a02:
		{
			m.ip = 0x1a03
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x1a03:
		{
			m.ip = 0x1a04
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1a04:
		{
			m.ip = 0x1a09
			m.wr8(m.r[11], uint16(0x2077), uint16(0))
			m.cycles += 8
		}
	case 0x1a09:
		{
			m.ip = 0x1a0e
			m.wr8(m.r[11], uint16(0x206d), uint16(8))
			m.cycles += 8
		}
	case 0x1a0e:
		{
			m.ip = 0x1a14
			m.wr16(m.r[11], uint16(0x206b), uint16(1280))
			m.cycles += 8
		}
	case 0x1a14:
		{
			m.ip = 0x1a18
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x1a18:
		{
			m.ip = 0x1a1c
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x233))
			m.cycles += 8
		}
	case 0x1a1c:
		{
			m.ip = 0x1a1f
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x1a1f:
		{
			m.ip = 0x1a23
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(m.r[3]+0x1985)))
			m.cycles += 16
		}
	case 0x1a23:
		{
			m.ip = 0x1a25
			if m.zf || m.sf != m.of {
				m.ip = uint16(6696)
			}
			m.cycles += 8
		}
	case 0x1a25:
		{
			m.ip = 0x1a28
			m.ip = uint16(6892)
			m.cycles += 15
		}
	case 0x1a28:
		{
			m.ip = 0x1a2c
			m.wr8(m.r[11], uint16(0x206d), m.unary("inc", 8, m.rd8(m.r[11], uint16(0x206d))))
			m.cycles += 15
		}
	case 0x1a2c:
		{
			m.ip = 0x1a32
			m.wr16(m.r[11], uint16(0x206b), m.alu("add", 16, m.rd16(m.r[11], uint16(0x206b)), uint16(160)))
			m.cycles += 16
		}
	case 0x1a32:
		{
			m.ip = 0x1a35
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x1a35:
		{
			m.ip = 0x1a38
			m.alu("sub", 16, m.r[3], uint16(20))
			m.cycles += 4
		}
	case 0x1a38:
		{
			m.ip = 0x1a3a
			if !m.zf {
				m.ip = uint16(6687)
			}
			m.cycles += 8
		}
	case 0x1a3a:
		{
			m.ip = 0x1a3f
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x1a3f:
		{
			m.ip = 0x1a41
			if !m.zf {
				m.ip = uint16(6775)
			}
			m.cycles += 8
		}
	case 0x1a41:
		{
			m.ip = 0x1a46
			m.wr8(m.r[11], uint16(0x1983), uint16(2))
			m.cycles += 8
		}
	case 0x1a46:
		{
			m.ip = 0x1a4b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1984)), uint16(0))
			m.cycles += 16
		}
	case 0x1a4b:
		{
			m.ip = 0x1a4d
			if m.zf {
				m.ip = uint16(6754)
			}
			m.cycles += 8
		}
	case 0x1a4d:
		{
			m.ip = 0x1a52
			m.wr8(m.r[11], uint16(0x1e2f), uint16(32))
			m.cycles += 8
		}
	case 0x1a52:
		{
			m.ip = 0x1a57
			m.wr8(m.r[11], uint16(0x1e30), uint16(32))
			m.cycles += 8
		}
	case 0x1a57:
		{
			m.ip = 0x1a5b
			m.r[7] = uint16(0x1e5a)
			m.cycles += 3
		}
	case 0x1a5b:
		{
			m.ip = 0x1a5f
			m.r[5] = uint16(0x18d0)
			m.cycles += 3
		}
	case 0x1a5f:
		{
			m.ip = 0x1a61
			m.ip = uint16(6869)
			m.cycles += 15
		}
	case 0x1a61:
		{
			m.ip = 0x1a62
			m.cycles += 3
		}
	case 0x1a62:
		{
			m.ip = 0x1a67
			m.wr8(m.r[11], uint16(0x1e79), uint16(32))
			m.cycles += 8
		}
	case 0x1a67:
		{
			m.ip = 0x1a6c
			m.wr8(m.r[11], uint16(0x1e7a), uint16(32))
			m.cycles += 8
		}
	case 0x1a6c:
		{
			m.ip = 0x1a70
			m.r[7] = uint16(0x1ea4)
			m.cycles += 3
		}
	case 0x1a70:
		{
			m.ip = 0x1a74
			m.r[5] = uint16(0x1905)
			m.cycles += 3
		}
	case 0x1a74:
		{
			m.ip = 0x1a76
			m.ip = uint16(6869)
			m.cycles += 15
		}
	case 0x1a76:
		{
			m.ip = 0x1a77
			m.cycles += 3
		}
	case 0x1a77:
		{
			m.ip = 0x1a7c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(2))
			m.cycles += 16
		}
	case 0x1a7c:
		{
			m.ip = 0x1a7e
			if m.zf {
				m.ip = uint16(6808)
			}
			m.cycles += 8
		}
	case 0x1a7e:
		{
			m.ip = 0x1a83
			m.wr8(m.r[11], uint16(0x1983), uint16(1))
			m.cycles += 8
		}
	case 0x1a83:
		{
			m.ip = 0x1a88
			m.wr8(m.r[11], uint16(0x1de5), uint16(32))
			m.cycles += 8
		}
	case 0x1a88:
		{
			m.ip = 0x1a8d
			m.wr8(m.r[11], uint16(0x1de6), uint16(32))
			m.cycles += 8
		}
	case 0x1a8d:
		{
			m.ip = 0x1a91
			m.r[7] = uint16(0x1e10)
			m.cycles += 3
		}
	case 0x1a91:
		{
			m.ip = 0x1a95
			m.r[5] = uint16(0x189b)
			m.cycles += 3
		}
	case 0x1a95:
		{
			m.ip = 0x1a97
			m.ip = uint16(6869)
			m.cycles += 15
		}
	case 0x1a97:
		{
			m.ip = 0x1a98
			m.cycles += 3
		}
	case 0x1a98:
		{
			m.ip = 0x1a9d
			m.wr8(m.r[11], uint16(0x1983), uint16(2))
			m.cycles += 8
		}
	case 0x1a9d:
		{
			m.ip = 0x1aa2
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(2))
			m.cycles += 16
		}
	case 0x1aa2:
		{
			m.ip = 0x1aa4
			if m.zf {
				m.ip = uint16(6846)
			}
			m.cycles += 8
		}
	case 0x1aa4:
		{
			m.ip = 0x1aa9
			m.wr8(m.r[11], uint16(0x1984), uint16(1))
			m.cycles += 8
		}
	case 0x1aa9:
		{
			m.ip = 0x1aae
			m.wr8(m.r[11], uint16(0x1e2f), uint16(32))
			m.cycles += 8
		}
	case 0x1aae:
		{
			m.ip = 0x1ab3
			m.wr8(m.r[11], uint16(0x1e30), uint16(32))
			m.cycles += 8
		}
	case 0x1ab3:
		{
			m.ip = 0x1ab7
			m.r[7] = uint16(0x1e5a)
			m.cycles += 3
		}
	case 0x1ab7:
		{
			m.ip = 0x1abb
			m.r[5] = uint16(0x18d0)
			m.cycles += 3
		}
	case 0x1abb:
		{
			m.ip = 0x1abd
			m.ip = uint16(6869)
			m.cycles += 15
		}
	case 0x1abd:
		{
			m.ip = 0x1abe
			m.cycles += 3
		}
	case 0x1abe:
		{
			m.ip = 0x1ac3
			m.wr8(m.r[11], uint16(0x1984), uint16(0))
			m.cycles += 8
		}
	case 0x1ac3:
		{
			m.ip = 0x1ac8
			m.wr8(m.r[11], uint16(0x1e79), uint16(32))
			m.cycles += 8
		}
	case 0x1ac8:
		{
			m.ip = 0x1acd
			m.wr8(m.r[11], uint16(0x1e7a), uint16(32))
			m.cycles += 8
		}
	case 0x1acd:
		{
			m.ip = 0x1ad1
			m.r[7] = uint16(0x1ea4)
			m.cycles += 3
		}
	case 0x1ad1:
		{
			m.ip = 0x1ad5
			m.r[5] = uint16(0x1905)
			m.cycles += 3
		}
	case 0x1ad5:
		{
			m.ip = 0x1ada
			m.wr8(m.r[11], uint16(0x1980), uint16(1))
			m.cycles += 8
		}
	case 0x1ada:
		{
			m.ip = 0x1ae0
			m.wr16(m.r[11], uint16(0x1981), uint16(100))
			m.cycles += 8
		}
	case 0x1ae0:
		{
			m.ip = 0x1ae3
			target := uint16(7545)
			m.push(0x1ae3)
			m.ip = target
			m.cycles += 19
		}
	case 0x1ae3:
		{
			m.ip = 0x1ae5
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x1ae5:
		{
			m.ip = 0x1ae8
			target := uint16(7032)
			m.push(0x1ae8)
			m.ip = target
			m.cycles += 19
		}
	case 0x1ae8:
		{
			m.ip = 0x1aeb
			target := uint16(18515)
			m.push(0x1aeb)
			m.ip = target
			m.cycles += 19
		}
	case 0x1aeb:
		{
			m.ip = 0x1aec
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1aec:
		{
			m.ip = 0x1af1
			m.wr8(m.r[11], uint16(0x1980), uint16(0))
			m.cycles += 8
		}
	case 0x1af1:
		{
			m.ip = 0x1af4
			target := uint16(7465)
			m.push(0x1af4)
			m.ip = target
			m.cycles += 19
		}
	case 0x1af4:
		{
			m.ip = 0x1af7
			target := uint16(8322)
			m.push(0x1af7)
			m.ip = target
			m.cycles += 19
		}
	case 0x1af7:
		{
			m.ip = 0x1afa
			target := uint16(3447)
			m.push(0x1afa)
			m.ip = target
			m.cycles += 19
		}
	case 0x1afa:
		{
			m.ip = 0x1afc
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x1afc:
		{
			m.ip = 0x1aff
			target := uint16(7032)
			m.push(0x1aff)
			m.ip = target
			m.cycles += 19
		}
	case 0x1aff:
		{
			m.ip = 0x1b00
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1b00:
		{
			m.ip = 0x1b01
			m.push(m.r[8])
			m.cycles += 11
		}
	case 0x1b01:
		{
			m.ip = 0x1b05
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x1b05:
		{
			m.ip = 0x1b09
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x233))
			m.cycles += 8
		}
	case 0x1b09:
		{
			m.ip = 0x1b0d
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(0x1985)))
			m.cycles += 16
		}
	case 0x1b0d:
		{
			m.ip = 0x1b0f
			if m.cf {
				m.ip = uint16(6970)
			}
			m.cycles += 8
		}
	case 0x1b0f:
		{
			m.ip = 0x1b12
			m.alu("sub", 16, m.r[0], uint16(1000))
			m.cycles += 4
		}
	case 0x1b12:
		{
			m.ip = 0x1b14
			if m.cf {
				m.ip = uint16(6970)
			}
			m.cycles += 8
		}
	case 0x1b14:
		{
			m.ip = 0x1b18
			m.r[2] = uint16(0x369)
			m.cycles += 3
		}
	case 0x1b18:
		{
			m.ip = 0x1b1b
			target := uint16(17942)
			m.push(0x1b1b)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b1b:
		{
			m.ip = 0x1b1f
			m.r[6] = uint16(0x1056)
			m.cycles += 3
		}
	case 0x1b1f:
		{
			m.ip = 0x1b22
			target := uint16(18306)
			m.push(0x1b22)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b22:
		{
			m.ip = 0x1b25
			m.r[0] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x1b25:
		{
			m.ip = 0x1b28
			target := uint16(18137)
			m.push(0x1b28)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b28:
		{
			m.ip = 0x1b2c
			m.r[6] = uint16(0x1045)
			m.cycles += 3
		}
	case 0x1b2c:
		{
			m.ip = 0x1b2f
			target := uint16(18306)
			m.push(0x1b2f)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b2f:
		{
			m.ip = 0x1b34
			m.wr8(m.r[11], uint16(0x270), uint16(1))
			m.cycles += 8
		}
	case 0x1b34:
		{
			m.ip = 0x1b37
			target := uint16(18515)
			m.push(0x1b37)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b37:
		{
			m.ip = 0x1b3a
			target := uint16(18893)
			m.push(0x1b3a)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b3a:
		{
			m.ip = 0x1b3b
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x1b3b:
		{
			m.ip = 0x1b3c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1b3c:
		{
			m.ip = 0x1b3f
			target := uint16(3533)
			m.push(0x1b3f)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b3f:
		{
			m.ip = 0x1b43
			m.r[6] = uint16(0x199b)
			m.cycles += 3
		}
	case 0x1b43:
		{
			m.ip = 0x1b46
			m.r[7] = uint16(648)
			m.cycles += 2
		}
	case 0x1b46:
		{
			m.ip = 0x1b49
			target := uint16(9087)
			m.push(0x1b49)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b49:
		{
			m.ip = 0x1b4d
			m.r[6] = uint16(0x1de3)
			m.cycles += 3
		}
	case 0x1b4d:
		{
			m.ip = 0x1b52
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1983)), uint16(1))
			m.cycles += 16
		}
	case 0x1b52:
		{
			m.ip = 0x1b54
			if m.zf {
				m.ip = uint16(7014)
			}
			m.cycles += 8
		}
	case 0x1b54:
		{
			m.ip = 0x1b58
			m.r[6] = uint16(0x1e2d)
			m.cycles += 3
		}
	case 0x1b58:
		{
			m.ip = 0x1b5d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1984)), uint16(1))
			m.cycles += 16
		}
	case 0x1b5d:
		{
			m.ip = 0x1b5f
			if m.zf {
				m.ip = uint16(7014)
			}
			m.cycles += 8
		}
	case 0x1b5f:
		{
			m.ip = 0x1b62
			target := uint16(9087)
			m.push(0x1b62)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b62:
		{
			m.ip = 0x1b66
			m.r[6] = uint16(0x1e77)
			m.cycles += 3
		}
	case 0x1b66:
		{
			m.ip = 0x1b69
			target := uint16(9087)
			m.push(0x1b69)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b69:
		{
			m.ip = 0x1b6d
			m.r[6] = uint16(0x1ec1)
			m.cycles += 3
		}
	case 0x1b6d:
		{
			m.ip = 0x1b70
			target := uint16(9087)
			m.push(0x1b70)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b70:
		{
			m.ip = 0x1b74
			m.r[6] = uint16(0x24dd)
			m.cycles += 3
		}
	case 0x1b74:
		{
			m.ip = 0x1b77
			target := uint16(7390)
			m.push(0x1b77)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b77:
		{
			m.ip = 0x1b78
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1b78:
		{
			m.ip = 0x1b7a
			m.set8(3, 8, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x1b7a:
		{
			m.ip = 0x1b7c
			m.set8(0, 8, ((m.r[3] >> 8) & 255))
			m.cycles += 2
		}
	case 0x1b7c:
		{
			m.ip = 0x1b7f
			target := uint16(6972)
			m.push(0x1b7f)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b7f:
		{
			m.ip = 0x1b84
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x1b84:
		{
			m.ip = 0x1b88
			m.r[6] = uint16(0x1fe9)
			m.cycles += 3
		}
	case 0x1b88:
		{
			m.ip = 0x1b8b
			target := uint16(7416)
			m.push(0x1b8b)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b8b:
		{
			m.ip = 0x1b8e
			target := uint16(7165)
			m.push(0x1b8e)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b8e:
		{
			m.ip = 0x1b90
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x1b90:
		{
			m.ip = 0x1b92
			if m.zf {
				m.ip = uint16(7164)
			}
			m.cycles += 8
		}
	case 0x1b92:
		{
			m.ip = 0x1b94
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(32))
			m.cycles += 4
		}
	case 0x1b94:
		{
			m.ip = 0x1b96
			if !m.zf {
				m.ip = uint16(7164)
			}
			m.cycles += 8
		}
	case 0x1b96:
		{
			m.ip = 0x1b97
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x1b97:
		{
			m.ip = 0x1b9a
			target := uint16(8813)
			m.push(0x1b9a)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b9a:
		{
			m.ip = 0x1b9d
			target := uint16(7259)
			m.push(0x1b9d)
			m.ip = target
			m.cycles += 19
		}
	case 0x1b9d:
		{
			m.ip = 0x1b9e
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1b9e:
		{
			m.ip = 0x1ba0
			m.set8(0, 8, ((m.r[3] >> 8) & 255))
			m.cycles += 2
		}
	case 0x1ba0:
		{
			m.ip = 0x1ba4
			m.r[6] = uint16(0x1412)
			m.cycles += 3
		}
	case 0x1ba4:
		{
			m.ip = 0x1ba7
			m.r[7] = uint16(8)
			m.cycles += 2
		}
	case 0x1ba7:
		{
			m.ip = 0x1baa
			target := uint16(9087)
			m.push(0x1baa)
			m.ip = target
			m.cycles += 19
		}
	case 0x1baa:
		{
			m.ip = 0x1bae
			m.r[6] = uint16(0x165b)
			m.cycles += 3
		}
	case 0x1bae:
		{
			m.ip = 0x1bb1
			m.r[7] = uint16(1330)
			m.cycles += 2
		}
	case 0x1bb1:
		{
			m.ip = 0x1bb4
			target := uint16(9087)
			m.push(0x1bb4)
			m.ip = target
			m.cycles += 19
		}
	case 0x1bb4:
		{
			m.ip = 0x1bb8
			m.r[6] = uint16(0x1898)
			m.cycles += 3
		}
	case 0x1bb8:
		{
			m.ip = 0x1bbd
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1983)), uint16(1))
			m.cycles += 16
		}
	case 0x1bbd:
		{
			m.ip = 0x1bbf
			if m.zf {
				m.ip = uint16(7121)
			}
			m.cycles += 8
		}
	case 0x1bbf:
		{
			m.ip = 0x1bc3
			m.r[6] = uint16(0x18cd)
			m.cycles += 3
		}
	case 0x1bc3:
		{
			m.ip = 0x1bc8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1984)), uint16(1))
			m.cycles += 16
		}
	case 0x1bc8:
		{
			m.ip = 0x1bca
			if m.zf {
				m.ip = uint16(7121)
			}
			m.cycles += 8
		}
	case 0x1bca:
		{
			m.ip = 0x1bcd
			target := uint16(9087)
			m.push(0x1bcd)
			m.ip = target
			m.cycles += 19
		}
	case 0x1bcd:
		{
			m.ip = 0x1bd1
			m.r[6] = uint16(0x1902)
			m.cycles += 3
		}
	case 0x1bd1:
		{
			m.ip = 0x1bd4
			target := uint16(9087)
			m.push(0x1bd4)
			m.ip = target
			m.cycles += 19
		}
	case 0x1bd4:
		{
			m.ip = 0x1bd8
			m.r[6] = uint16(0x1937)
			m.cycles += 3
		}
	case 0x1bd8:
		{
			m.ip = 0x1bdb
			target := uint16(9087)
			m.push(0x1bdb)
			m.ip = target
			m.cycles += 19
		}
	case 0x1bdb:
		{
			m.ip = 0x1bdf
			m.r[6] = uint16(0x24e7)
			m.cycles += 3
		}
	case 0x1bdf:
		{
			m.ip = 0x1be2
			target := uint16(7390)
			m.push(0x1be2)
			m.ip = target
			m.cycles += 19
		}
	case 0x1be2:
		{
			m.ip = 0x1be6
			m.r[6] = uint16(0x202b)
			m.cycles += 3
		}
	case 0x1be6:
		{
			m.ip = 0x1be9
			target := uint16(7416)
			m.push(0x1be9)
			m.ip = target
			m.cycles += 19
		}
	case 0x1be9:
		{
			m.ip = 0x1bec
			target := uint16(7165)
			m.push(0x1bec)
			m.ip = target
			m.cycles += 19
		}
	case 0x1bec:
		{
			m.ip = 0x1bf1
			m.wr8(m.r[11], uint16(0x2077), uint16(0))
			m.cycles += 8
		}
	case 0x1bf1:
		{
			m.ip = 0x1bf3
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x1bf3:
		{
			m.ip = 0x1bf5
			if m.zf {
				m.ip = uint16(7164)
			}
			m.cycles += 8
		}
	case 0x1bf5:
		{
			m.ip = 0x1bf7
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(32))
			m.cycles += 4
		}
	case 0x1bf7:
		{
			m.ip = 0x1bf9
			if !m.zf {
				m.ip = uint16(7164)
			}
			m.cycles += 8
		}
	case 0x1bf9:
		{
			m.ip = 0x1bfc
			m.ip = uint16(7034)
			m.cycles += 15
		}
	case 0x1bfc:
		{
			m.ip = 0x1bfd
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1bfd:
		{
			m.ip = 0x1bfe
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x1bfe:
		{
			m.ip = 0x1c00
			m.set8(3, 0, uint16(2))
			m.cycles += 2
		}
	case 0x1c00:
		{
			m.ip = 0x1c04
			m.wr8(m.r[11], uint16(0x24f1), m.unary("inc", 8, m.rd8(m.r[11], uint16(0x24f1))))
			m.cycles += 15
		}
	case 0x1c04:
		{
			m.ip = 0x1c09
			m.wr8(m.r[11], uint16(0x24f1), m.alu("and", 8, m.rd8(m.r[11], uint16(0x24f1)), uint16(63)))
			m.cycles += 16
		}
	case 0x1c09:
		{
			m.ip = 0x1c0e
			m.wr8(m.r[11], uint16(0x24f1), m.alu("or", 8, m.rd8(m.r[11], uint16(0x24f1)), uint16(56)))
			m.cycles += 16
		}
	case 0x1c0e:
		{
			m.ip = 0x1c13
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f1)), uint16(56))
			m.cycles += 16
		}
	case 0x1c13:
		{
			m.ip = 0x1c15
			if !m.zf {
				m.ip = uint16(7193)
			}
			m.cycles += 8
		}
	case 0x1c15:
		{
			m.ip = 0x1c19
			m.wr8(m.r[11], uint16(0x24f1), m.unary("inc", 8, m.rd8(m.r[11], uint16(0x24f1))))
			m.cycles += 15
		}
	case 0x1c19:
		{
			m.ip = 0x1c1d
			m.set8(3, 8, m.rd8(m.r[11], uint16(0x24f1)))
			m.cycles += 8
		}
	case 0x1c1d:
		{
			m.ip = 0x1c20
			m.r[0] = uint16(4096)
			m.cycles += 2
		}
	case 0x1c20:
		{
			m.ip = 0x1c22
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x1c22:
		{
			m.ip = 0x1c23
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1c23:
		{
			m.ip = 0x1c27
			m.push(m.rd16(m.r[11], uint16(0x106c)))
			m.cycles += 11
		}
	case 0x1c27:
		{
			m.ip = 0x1c2b
			m.push(m.rd16(m.r[11], uint16(0x497)))
			m.cycles += 11
		}
	case 0x1c2b:
		{
			m.ip = 0x1c2f
			m.wr16(m.r[11], uint16(0x106c), m.pop())
			m.cycles += 8
		}
	case 0x1c2f:
		{
			m.ip = 0x1c35
			m.wr16(m.r[11], uint16(0x106a), uint16(500))
			m.cycles += 8
		}
	case 0x1c35:
		{
			m.ip = 0x1c38
			target := uint16(19311)
			m.push(0x1c38)
			m.ip = target
			m.cycles += 19
		}
	case 0x1c38:
		{
			m.ip = 0x1c3c
			m.wr16(m.r[11], uint16(0x106c), m.pop())
			m.cycles += 8
		}
	case 0x1c3c:
		{
			m.ip = 0x1c41
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1070)), uint16(0))
			m.cycles += 16
		}
	case 0x1c41:
		{
			m.ip = 0x1c43
			if m.zf {
				m.ip = uint16(7165)
			}
			m.cycles += 8
		}
	case 0x1c43:
		{
			m.ip = 0x1c48
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x27f)), uint16(0))
			m.cycles += 16
		}
	case 0x1c48:
		{
			m.ip = 0x1c4a
			if m.zf {
				m.ip = uint16(7255)
			}
			m.cycles += 8
		}
	case 0x1c4a:
		{
			m.ip = 0x1c4c
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x1c4c:
		{
			m.ip = 0x1c4e
			if m.zf {
				m.ip = uint16(7258)
			}
			m.cycles += 8
		}
	case 0x1c4e:
		{
			m.ip = 0x1c50
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x1c50:
		{
			m.ip = 0x1c52
			if !m.zf {
				m.ip = uint16(7165)
			}
			m.cycles += 8
		}
	case 0x1c52:
		{
			m.ip = 0x1c54
			m.set8(0, 0, uint16(32))
			m.cycles += 2
		}
	case 0x1c54:
		{
			m.ip = 0x1c56
			m.ip = uint16(7258)
			m.cycles += 15
		}
	case 0x1c56:
		{
			m.ip = 0x1c57
			m.cycles += 3
		}
	case 0x1c57:
		{
			m.ip = 0x1c5a
			target := uint16(18515)
			m.push(0x1c5a)
			m.ip = target
			m.cycles += 19
		}
	case 0x1c5a:
		{
			m.ip = 0x1c5b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1c5b:
		{
			m.ip = 0x1c5c
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x1c5c:
		{
			m.ip = 0x1c60
			m.r[5] = uint16(0x196c)
			m.cycles += 3
		}
	case 0x1c60:
		{
			m.ip = 0x1c64
			m.r[7] = uint16(0x1472)
			m.cycles += 3
		}
	case 0x1c64:
		{
			m.ip = 0x1c67
			target := uint16(7284)
			m.push(0x1c67)
			m.ip = target
			m.cycles += 19
		}
	case 0x1c67:
		{
			m.ip = 0x1c6b
			m.r[5] = uint16(0x1976)
			m.cycles += 3
		}
	case 0x1c6b:
		{
			m.ip = 0x1c6f
			m.r[7] = uint16(0x14bb)
			m.cycles += 3
		}
	case 0x1c6f:
		{
			m.ip = 0x1c72
			target := uint16(7284)
			m.push(0x1c72)
			m.ip = target
			m.cycles += 19
		}
	case 0x1c72:
		{
			m.ip = 0x1c73
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1c73:
		{
			m.ip = 0x1c74
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1c74:
		{
			m.ip = 0x1c76
			m.r[3] = m.alu("xor", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x1c76:
		{
			m.ip = 0x1c78
			m.r[6] = m.alu("xor", 16, m.r[6], m.r[6])
			m.cycles += 4
		}
	case 0x1c78:
		{
			m.ip = 0x1c7d
			m.wr8(m.r[11], uint16(0x13a3), uint16(0))
			m.cycles += 8
		}
	case 0x1c7d:
		{
			m.ip = 0x1c81
			m.r[1] = m.rd16(m.r[11], uint16(m.r[3]+0x13a4))
			m.cycles += 8
		}
	case 0x1c81:
		{
			m.ip = 0x1c83
			m.r[3] = m.alu("add", 16, m.r[3], m.r[5])
			m.cycles += 4
		}
	case 0x1c83:
		{
			m.ip = 0x1c85
			m.r[2] = m.rd16(m.r[11], uint16(m.r[3]))
			m.cycles += 8
		}
	case 0x1c85:
		{
			m.ip = 0x1c87
			m.r[3] = m.alu("sub", 16, m.r[3], m.r[5])
			m.cycles += 4
		}
	case 0x1c87:
		{
			m.ip = 0x1c8a
			m.alu("sub", 16, m.r[2], uint16(0))
			m.cycles += 4
		}
	case 0x1c8a:
		{
			m.ip = 0x1c8c
			if !m.zf {
				m.ip = uint16(7329)
			}
			m.cycles += 8
		}
	case 0x1c8c:
		{
			m.ip = 0x1c91
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x13a3)), uint16(1))
			m.cycles += 16
		}
	case 0x1c91:
		{
			m.ip = 0x1c93
			if m.zf {
				m.ip = uint16(7329)
			}
			m.cycles += 8
		}
	case 0x1c93:
		{
			m.ip = 0x1c97
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(m.r[3]+0x13a4)))
			m.cycles += 16
		}
	case 0x1c97:
		{
			m.ip = 0x1c99
			m.set8(0, 0, uint16(32))
			m.cycles += 2
		}
	case 0x1c99:
		{
			m.ip = 0x1c9b
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1c9b:
		{
			m.ip = 0x1c9c
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x1c9c:
		{
			m.ip = 0x1c9e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7321)
			}
			m.cycles += 17
		}
	case 0x1c9e:
		{
			m.ip = 0x1ca0
			m.ip = uint16(7381)
			m.cycles += 15
		}
	case 0x1ca0:
		{
			m.ip = 0x1ca1
			m.cycles += 3
		}
	case 0x1ca1:
		{
			m.ip = 0x1ca2
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x1ca2:
		{
			m.ip = 0x1ca7
			m.wr8(m.r[11], uint16(0x13a3), uint16(1))
			m.cycles += 8
		}
	case 0x1ca7:
		{
			m.ip = 0x1cab
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x13b4))
			m.cycles += 8
		}
	case 0x1cab:
		{
			m.ip = 0x1caf
			m.r[3] = uint16(0x1023)
			m.cycles += 3
		}
	case 0x1caf:
		{
			m.ip = 0x1cb1
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x1cb1:
		{
			m.ip = 0x1cb3
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x1cb3:
		{
			m.ip = 0x1cb4
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x1cb4:
		{
			m.ip = 0x1cb7
			target := uint16(9036)
			m.push(0x1cb7)
			m.ip = target
			m.cycles += 19
		}
	case 0x1cb7:
		{
			m.ip = 0x1cb8
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x1cb8:
		{
			m.ip = 0x1cb9
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1cb9:
		{
			m.ip = 0x1cbd
			m.r[1] = m.alu("sub", 16, m.r[1], m.rd16(m.r[11], uint16(m.r[3]+0x13ac)))
			m.cycles += 16
		}
	case 0x1cbd:
		{
			m.ip = 0x1cc1
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(m.r[3]+0x13ac)))
			m.cycles += 16
		}
	case 0x1cc1:
		{
			m.ip = 0x1cc2
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x1cc2:
		{
			m.ip = 0x1cc6
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x13e7)))
			m.cycles += 8
		}
	case 0x1cc6:
		{
			m.ip = 0x1cc9
			m.alu("sub", 16, m.r[2], uint16(1))
			m.cycles += 4
		}
	case 0x1cc9:
		{
			m.ip = 0x1ccb
			if !m.zf {
				m.ip = uint16(7375)
			}
			m.cycles += 8
		}
	case 0x1ccb:
		{
			m.ip = 0x1ccf
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x13bc)))
			m.cycles += 8
		}
	case 0x1ccf:
		{
			m.ip = 0x1cd1
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1cd1:
		{
			m.ip = 0x1cd2
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1cd2:
		{
			m.ip = 0x1cd3
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x1cd3:
		{
			m.ip = 0x1cd5
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7362)
			}
			m.cycles += 17
		}
	case 0x1cd5:
		{
			m.ip = 0x1cd8
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x1cd8:
		{
			m.ip = 0x1cdb
			m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x1cdb:
		{
			m.ip = 0x1cdd
			if !m.zf {
				m.ip = uint16(7293)
			}
			m.cycles += 8
		}
	case 0x1cdd:
		{
			m.ip = 0x1cde
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1cde:
		{
			m.ip = 0x1ce0
			m.set8(0, 8, uint16(2))
			m.cycles += 2
		}
	case 0x1ce0:
		{
			m.ip = 0x1ce3
			m.r[7] = uint16(1442)
			m.cycles += 2
		}
	case 0x1ce3:
		{
			m.ip = 0x1ce6
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x1ce6:
		{
			m.ip = 0x1ce8
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x1ce8:
		{
			m.ip = 0x1ceb
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1ceb:
		{
			m.ip = 0x1cf0
			m.wr16(m.r[8], uint16(m.r[7]+0x9a), m.r[0])
			m.cycles += 8
		}
	case 0x1cf0:
		{
			m.ip = 0x1cf1
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1cf1:
		{
			m.ip = 0x1cf5
			m.r[7] = m.alu("add", 16, m.r[7], uint16(160))
			m.cycles += 4
		}
	case 0x1cf5:
		{
			m.ip = 0x1cf7
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7398)
			}
			m.cycles += 17
		}
	case 0x1cf7:
		{
			m.ip = 0x1cf8
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1cf8:
		{
			m.ip = 0x1cfa
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x1cfa:
		{
			m.ip = 0x1cfc
			m.set8(0, 0, uint16(32))
			m.cycles += 2
		}
	case 0x1cfc:
		{
			m.ip = 0x1cff
			m.r[1] = uint16(240)
			m.cycles += 2
		}
	case 0x1cff:
		{
			m.ip = 0x1d02
			m.r[7] = uint16(3520)
			m.cycles += 2
		}
	case 0x1d02:
		{
			m.ip = 0x1d05
			m.r[5] = uint16(3760)
			m.cycles += 2
		}
	case 0x1d05:
		{
			m.ip = 0x1d09
			m.alu("sub", 16, m.r[6], uint16(4897))
			m.cycles += 4
		}
	case 0x1d09:
		{
			m.ip = 0x1d0b
			if !m.cf {
				m.ip = uint16(7441)
			}
			m.cycles += 8
		}
	case 0x1d0b:
		{
			m.ip = 0x1d0e
			m.r[7] = uint16(3360)
			m.cycles += 2
		}
	case 0x1d0e:
		{
			m.ip = 0x1d11
			m.r[5] = uint16(3600)
			m.cycles += 2
		}
	case 0x1d11:
		{
			m.ip = 0x1d14
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1d14:
		{
			m.ip = 0x1d17
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1d17:
		{
			m.ip = 0x1d19
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7441)
			}
			m.cycles += 17
		}
	case 0x1d19:
		{
			m.ip = 0x1d1c
			target := uint16(7453)
			m.push(0x1d1c)
			m.ip = target
			m.cycles += 19
		}
	case 0x1d1c:
		{
			m.ip = 0x1d1d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1d1d:
		{
			m.ip = 0x1d1f
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x1d1f:
		{
			m.ip = 0x1d22
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x1d22:
		{
			m.ip = 0x1d23
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x1d23:
		{
			m.ip = 0x1d24
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x1d24:
		{
			m.ip = 0x1d26
			m.r[7] = m.r[5]
			m.cycles += 2
		}
	case 0x1d26:
		{
			m.ip = 0x1d29
			m.ip = uint16(4821)
			m.cycles += 15
		}
	case 0x1d29:
		{
			m.ip = 0x1d2a
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x1d2a:
		{
			m.ip = 0x1d2b
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x1d2b:
		{
			m.ip = 0x1d2c
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x1d2c:
		{
			m.ip = 0x1d30
			m.r[6] = uint16(0x1997)
			m.cycles += 3
		}
	case 0x1d30:
		{
			m.ip = 0x1d32
			m.r[7] = m.r[6]
			m.cycles += 2
		}
	case 0x1d32:
		{
			m.ip = 0x1d35
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1d35:
		{
			m.ip = 0x1d38
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x1d38:
		{
			m.ip = 0x1d3a
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x1d3a:
		{
			m.ip = 0x1d3c
			m.r[1] = m.alu("sub", 16, m.r[1], m.r[3])
			m.cycles += 4
		}
	case 0x1d3c:
		{
			m.ip = 0x1d3d
			m.df = true
			m.cycles += 2
		}
	case 0x1d3d:
		{
			m.ip = 0x1d3f
			m.stringOp("movs", 16)
			m.cycles += 2
		}
	case 0x1d3f:
		{
			m.ip = 0x1d41
			m.wr16(m.r[11], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1d41:
		{
			m.ip = 0x1d45
			m.r[6] = uint16(0x1d0f)
			m.cycles += 3
		}
	case 0x1d45:
		{
			m.ip = 0x1d47
			m.r[7] = m.r[6]
			m.cycles += 2
		}
	case 0x1d47:
		{
			m.ip = 0x1d4a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(73))
			m.cycles += 4
		}
	case 0x1d4a:
		{
			m.ip = 0x1d4d
			m.r[1] = uint16(9)
			m.cycles += 2
		}
	case 0x1d4d:
		{
			m.ip = 0x1d4f
			m.r[1] = m.alu("sub", 16, m.r[1], m.r[3])
			m.cycles += 4
		}
	case 0x1d4f:
		{
			m.ip = 0x1d52
			m.alu("sub", 16, m.r[1], uint16(0))
			m.cycles += 4
		}
	case 0x1d52:
		{
			m.ip = 0x1d54
			if m.zf {
				m.ip = uint16(7528)
			}
			m.cycles += 8
		}
	case 0x1d54:
		{
			m.ip = 0x1d55
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x1d55:
		{
			m.ip = 0x1d56
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x1d56:
		{
			m.ip = 0x1d57
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x1d57:
		{
			m.ip = 0x1d5a
			m.r[1] = uint16(63)
			m.cycles += 2
		}
	case 0x1d5a:
		{
			m.ip = 0x1d5b
			m.df = false
			m.cycles += 2
		}
	case 0x1d5b:
		{
			m.ip = 0x1d5d
			m.stringOp("movs", 8)
			m.cycles += 2
		}
	case 0x1d5d:
		{
			m.ip = 0x1d5e
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x1d5e:
		{
			m.ip = 0x1d5f
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x1d5f:
		{
			m.ip = 0x1d62
			m.r[6] = m.alu("sub", 16, m.r[6], uint16(73))
			m.cycles += 4
		}
	case 0x1d62:
		{
			m.ip = 0x1d65
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(73))
			m.cycles += 4
		}
	case 0x1d65:
		{
			m.ip = 0x1d66
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x1d66:
		{
			m.ip = 0x1d68
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7508)
			}
			m.cycles += 17
		}
	case 0x1d68:
		{
			m.ip = 0x1d6c
			m.wr16(m.r[11], uint16(0x2069), m.r[7])
			m.cycles += 8
		}
	case 0x1d6c:
		{
			m.ip = 0x1d6e
			m.set8(2, 0, uint16(32))
			m.cycles += 2
		}
	case 0x1d6e:
		{
			m.ip = 0x1d71
			m.r[1] = uint16(34)
			m.cycles += 2
		}
	case 0x1d71:
		{
			m.ip = 0x1d73
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[2] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1d73:
		{
			m.ip = 0x1d74
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x1d74:
		{
			m.ip = 0x1d76
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7537)
			}
			m.cycles += 17
		}
	case 0x1d76:
		{
			m.ip = 0x1d79
			m.r[7] = m.alu("add", 16, m.r[7], uint16(3))
			m.cycles += 4
		}
	case 0x1d79:
		{
			m.ip = 0x1d7b
			m.r[2] = m.r[0]
			m.cycles += 2
		}
	case 0x1d7b:
		{
			m.ip = 0x1d7f
			m.r[3] = uint16(0x1023)
			m.cycles += 3
		}
	case 0x1d7f:
		{
			m.ip = 0x1d81
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x1d81:
		{
			m.ip = 0x1d83
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3])))
			m.cycles += 16
		}
	case 0x1d83:
		{
			m.ip = 0x1d85
			if m.sf != m.of {
				m.ip = uint16(7563)
			}
			m.cycles += 8
		}
	case 0x1d85:
		{
			m.ip = 0x1d87
			m.set8(0, 0, m.unary("inc", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x1d87:
		{
			m.ip = 0x1d89
			m.r[2] = m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3])))
			m.cycles += 16
		}
	case 0x1d89:
		{
			m.ip = 0x1d8b
			m.ip = uint16(7553)
			m.cycles += 15
		}
	case 0x1d8b:
		{
			m.ip = 0x1d8d
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), uint16(48)))
			m.cycles += 4
		}
	case 0x1d8d:
		{
			m.ip = 0x1d8f
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1d8f:
		{
			m.ip = 0x1d90
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x1d90:
		{
			m.ip = 0x1d93
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x1d93:
		{
			m.ip = 0x1d95
			m.set8(0, 0, uint16(10))
			m.cycles += 2
		}
	case 0x1d95:
		{
			m.ip = 0x1d98
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0xfffe)), ((m.r[0] >> 0) & 255))
			m.cycles += 16
		}
	case 0x1d98:
		{
			m.ip = 0x1d9a
			if !m.zf {
				m.ip = uint16(7551)
			}
			m.cycles += 8
		}
	case 0x1d9a:
		{
			m.ip = 0x1d9c
			m.set8(0, 0, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x1d9c:
		{
			m.ip = 0x1d9e
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), uint16(48)))
			m.cycles += 4
		}
	case 0x1d9e:
		{
			m.ip = 0x1da0
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1da0:
		{
			m.ip = 0x1da3
			m.r[7] = m.alu("add", 16, m.r[7], uint16(5))
			m.cycles += 4
		}
	case 0x1da3:
		{
			m.ip = 0x1da6
			target := uint16(9144)
			m.push(0x1da6)
			m.ip = target
			m.cycles += 19
		}
	case 0x1da6:
		{
			m.ip = 0x1da8
			m.set8(0, 0, ((m.r[1] >> 0) & 255))
			m.cycles += 2
		}
	case 0x1da8:
		{
			m.ip = 0x1dab
			target := uint16(9120)
			m.push(0x1dab)
			m.ip = target
			m.cycles += 19
		}
	case 0x1dab:
		{
			m.ip = 0x1dad
			m.set8(0, 0, ((m.r[2] >> 8) & 255))
			m.cycles += 2
		}
	case 0x1dad:
		{
			m.ip = 0x1db0
			target := uint16(9120)
			m.push(0x1db0)
			m.ip = target
			m.cycles += 19
		}
	case 0x1db0:
		{
			m.ip = 0x1db2
			m.set8(0, 0, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x1db2:
		{
			m.ip = 0x1db5
			target := uint16(9120)
			m.push(0x1db5)
			m.ip = target
			m.cycles += 19
		}
	case 0x1db5:
		{
			m.ip = 0x1db8
			target := uint16(9161)
			m.push(0x1db8)
			m.ip = target
			m.cycles += 19
		}
	case 0x1db8:
		{
			m.ip = 0x1dbb
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1dbb:
		{
			m.ip = 0x1dbd
			m.set8(0, 0, ((m.r[1] >> 8) & 255))
			m.cycles += 2
		}
	case 0x1dbd:
		{
			m.ip = 0x1dc0
			target := uint16(9120)
			m.push(0x1dc0)
			m.ip = target
			m.cycles += 19
		}
	case 0x1dc0:
		{
			m.ip = 0x1dc2
			m.set8(0, 0, ((m.r[1] >> 0) & 255))
			m.cycles += 2
		}
	case 0x1dc2:
		{
			m.ip = 0x1dc5
			target := uint16(9120)
			m.push(0x1dc5)
			m.ip = target
			m.cycles += 19
		}
	case 0x1dc5:
		{
			m.ip = 0x1dca
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1980)), uint16(1))
			m.cycles += 16
		}
	case 0x1dca:
		{
			m.ip = 0x1dcc
			if !m.zf {
				m.ip = uint16(7631)
			}
			m.cycles += 8
		}
	case 0x1dcc:
		{
			m.ip = 0x1dcf
			m.ip = uint16(8371)
			m.cycles += 15
		}
	case 0x1dcf:
		{
			m.ip = 0x1dd3
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x1dd3:
		{
			m.ip = 0x1dd5
			m.set8(0, 8, uint16(7))
			m.cycles += 2
		}
	case 0x1dd5:
		{
			m.ip = 0x1dd8
			target := uint16(6972)
			m.push(0x1dd8)
			m.ip = target
			m.cycles += 19
		}
	case 0x1dd8:
		{
			m.ip = 0x1ddc
			m.r[7] = m.rd16(m.r[11], uint16(0x206b))
			m.cycles += 8
		}
	case 0x1ddc:
		{
			m.ip = 0x1de1
			m.wr8(m.r[8], uint16(m.r[7]+0xd), uint16(5))
			m.cycles += 8
		}
	case 0x1de1:
		{
			m.ip = 0x1de6
			m.wr8(m.r[8], uint16(m.r[7]+0xf), uint16(5))
			m.cycles += 8
		}
	case 0x1de6:
		{
			m.ip = 0x1de9
			m.r[7] = m.alu("add", 16, m.r[7], uint16(25))
			m.cycles += 4
		}
	case 0x1de9:
		{
			m.ip = 0x1dec
			m.r[1] = uint16(34)
			m.cycles += 2
		}
	case 0x1dec:
		{
			m.ip = 0x1df0
			m.wr8(m.r[8], uint16(m.r[7]), uint16(4))
			m.cycles += 8
		}
	case 0x1df0:
		{
			m.ip = 0x1df3
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1df3:
		{
			m.ip = 0x1df5
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7660)
			}
			m.cycles += 17
		}
	case 0x1df5:
		{
			m.ip = 0x1df8
			m.r[7] = m.alu("add", 16, m.r[7], uint16(6))
			m.cycles += 4
		}
	case 0x1df8:
		{
			m.ip = 0x1dfb
			m.r[1] = uint16(6)
			m.cycles += 2
		}
	case 0x1dfb:
		{
			m.ip = 0x1dff
			m.wr8(m.r[8], uint16(m.r[7]), uint16(5))
			m.cycles += 8
		}
	case 0x1dff:
		{
			m.ip = 0x1e02
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1e02:
		{
			m.ip = 0x1e04
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7675)
			}
			m.cycles += 17
		}
	case 0x1e04:
		{
			m.ip = 0x1e07
			m.r[7] = m.alu("add", 16, m.r[7], uint16(6))
			m.cycles += 4
		}
	case 0x1e07:
		{
			m.ip = 0x1e0a
			m.r[1] = uint16(8)
			m.cycles += 2
		}
	case 0x1e0a:
		{
			m.ip = 0x1e0e
			m.wr8(m.r[8], uint16(m.r[7]), uint16(4))
			m.cycles += 8
		}
	case 0x1e0e:
		{
			m.ip = 0x1e11
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1e11:
		{
			m.ip = 0x1e13
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7690)
			}
			m.cycles += 17
		}
	case 0x1e13:
		{
			m.ip = 0x1e16
			m.r[7] = m.alu("add", 16, m.r[7], uint16(6))
			m.cycles += 4
		}
	case 0x1e16:
		{
			m.ip = 0x1e19
			m.r[1] = uint16(5)
			m.cycles += 2
		}
	case 0x1e19:
		{
			m.ip = 0x1e1d
			m.wr8(m.r[8], uint16(m.r[7]), uint16(4))
			m.cycles += 8
		}
	case 0x1e1d:
		{
			m.ip = 0x1e20
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x1e20:
		{
			m.ip = 0x1e22
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7705)
			}
			m.cycles += 17
		}
	case 0x1e22:
		{
			m.ip = 0x1e26
			m.r[6] = uint16(0x1f0b)
			m.cycles += 3
		}
	case 0x1e26:
		{
			m.ip = 0x1e2b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(1))
			m.cycles += 16
		}
	case 0x1e2b:
		{
			m.ip = 0x1e2d
			if m.zf {
				m.ip = uint16(7740)
			}
			m.cycles += 8
		}
	case 0x1e2d:
		{
			m.ip = 0x1e31
			m.r[6] = uint16(0x1f2f)
			m.cycles += 3
		}
	case 0x1e31:
		{
			m.ip = 0x1e36
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x1e36:
		{
			m.ip = 0x1e38
			if m.zf {
				m.ip = uint16(7740)
			}
			m.cycles += 8
		}
	case 0x1e38:
		{
			m.ip = 0x1e3c
			m.r[6] = uint16(0x1f5e)
			m.cycles += 3
		}
	case 0x1e3c:
		{
			m.ip = 0x1e3f
			target := uint16(7416)
			m.push(0x1e3f)
			m.ip = target
			m.cycles += 19
		}
	case 0x1e3f:
		{
			m.ip = 0x1e42
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x1e42:
		{
			m.ip = 0x1e46
			m.r[7] = m.rd16(m.r[11], uint16(0x2069))
			m.cycles += 8
		}
	case 0x1e46:
		{
			m.ip = 0x1e4a
			m.set8(2, 8, m.rd8(m.r[11], uint16(0x206d)))
			m.cycles += 8
		}
	case 0x1e4a:
		{
			m.ip = 0x1e4c
			m.set8(2, 0, uint16(12))
			m.cycles += 2
		}
	case 0x1e4c:
		{
			m.ip = 0x1e4f
			target := uint16(8092)
			m.push(0x1e4f)
			m.ip = target
			m.cycles += 19
		}
	case 0x1e4f:
		{
			m.ip = 0x1e52
			m.r[1] = uint16(1543)
			m.cycles += 2
		}
	case 0x1e52:
		{
			m.ip = 0x1e54
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x1e54:
		{
			m.ip = 0x1e56
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x1e56:
		{
			m.ip = 0x1e59
			target := uint16(18893)
			m.push(0x1e59)
			m.ip = target
			m.cycles += 19
		}
	case 0x1e59:
		{
			m.ip = 0x1e5b
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x1e5b:
		{
			m.ip = 0x1e5d
			if m.zf {
				m.ip = uint16(7870)
			}
			m.cycles += 8
		}
	case 0x1e5d:
		{
			m.ip = 0x1e5f
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(9))
			m.cycles += 4
		}
	case 0x1e5f:
		{
			m.ip = 0x1e61
			if m.zf {
				m.ip = uint16(7839)
			}
			m.cycles += 8
		}
	case 0x1e61:
		{
			m.ip = 0x1e63
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(8))
			m.cycles += 4
		}
	case 0x1e63:
		{
			m.ip = 0x1e65
			if m.zf {
				m.ip = uint16(7793)
			}
			m.cycles += 8
		}
	case 0x1e65:
		{
			m.ip = 0x1e67
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x1e67:
		{
			m.ip = 0x1e69
			if m.zf {
				m.ip = uint16(7788)
			}
			m.cycles += 8
		}
	case 0x1e69:
		{
			m.ip = 0x1e6c
			m.ip = uint16(8045)
			m.cycles += 15
		}
	case 0x1e6c:
		{
			m.ip = 0x1e6f
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(83))
			m.cycles += 4
		}
	case 0x1e6f:
		{
			m.ip = 0x1e71
			if !m.zf {
				m.ip = uint16(7766)
			}
			m.cycles += 8
		}
	case 0x1e71:
		{
			m.ip = 0x1e73
			m.set8(0, 0, uint16(32))
			m.cycles += 2
		}
	case 0x1e73:
		{
			m.ip = 0x1e77
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), m.rd8(m.r[11], uint16(m.r[6]+0x206e)))
			m.cycles += 16
		}
	case 0x1e77:
		{
			m.ip = 0x1e79
			if m.zf {
				m.ip = uint16(7832)
			}
			m.cycles += 8
		}
	case 0x1e79:
		{
			m.ip = 0x1e7d
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), m.rd8(m.r[11], uint16(m.r[6]+0x2071)))
			m.cycles += 16
		}
	case 0x1e7d:
		{
			m.ip = 0x1e7f
			if !m.zf {
				m.ip = uint16(7811)
			}
			m.cycles += 8
		}
	case 0x1e7f:
		{
			m.ip = 0x1e81
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[7])), ((m.r[0] >> 0) & 255))
			m.cycles += 16
		}
	case 0x1e81:
		{
			m.ip = 0x1e83
			if !m.zf {
				m.ip = uint16(7832)
			}
			m.cycles += 8
		}
	case 0x1e83:
		{
			m.ip = 0x1e84
			m.r[7] = m.unary("dec", 16, m.r[7])
			m.cycles += 3
		}
	case 0x1e84:
		{
			m.ip = 0x1e86
			m.set8(2, 0, m.unary("dec", 8, ((m.r[2]>>0)&255)))
			m.cycles += 3
		}
	case 0x1e86:
		{
			m.ip = 0x1e89
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(60))
			m.cycles += 4
		}
	case 0x1e89:
		{
			m.ip = 0x1e8b
			if m.zf {
				m.ip = uint16(7811)
			}
			m.cycles += 8
		}
	case 0x1e8b:
		{
			m.ip = 0x1e8e
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(63))
			m.cycles += 4
		}
	case 0x1e8e:
		{
			m.ip = 0x1e90
			if m.zf {
				m.ip = uint16(7811)
			}
			m.cycles += 8
		}
	case 0x1e90:
		{
			m.ip = 0x1e93
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(71))
			m.cycles += 4
		}
	case 0x1e93:
		{
			m.ip = 0x1e95
			if m.zf {
				m.ip = uint16(7811)
			}
			m.cycles += 8
		}
	case 0x1e95:
		{
			m.ip = 0x1e98
			target := uint16(8092)
			m.push(0x1e98)
			m.ip = target
			m.cycles += 19
		}
	case 0x1e98:
		{
			m.ip = 0x1e9a
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1e9a:
		{
			m.ip = 0x1e9d
			target := uint16(8080)
			m.push(0x1e9d)
			m.ip = target
			m.cycles += 19
		}
	case 0x1e9d:
		{
			m.ip = 0x1e9f
			m.ip = uint16(7766)
			m.cycles += 15
		}
	case 0x1e9f:
		{
			m.ip = 0x1ea0
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1ea0:
		{
			m.ip = 0x1ea3
			m.alu("sub", 16, m.r[6], uint16(3))
			m.cycles += 4
		}
	case 0x1ea3:
		{
			m.ip = 0x1ea5
			if !m.zf {
				m.ip = uint16(7848)
			}
			m.cycles += 8
		}
	case 0x1ea5:
		{
			m.ip = 0x1ea8
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x1ea8:
		{
			m.ip = 0x1eac
			m.set8(2, 0, m.rd8(m.r[11], uint16(m.r[6]+0x206e)))
			m.cycles += 8
		}
	case 0x1eac:
		{
			m.ip = 0x1eaf
			target := uint16(8092)
			m.push(0x1eaf)
			m.ip = target
			m.cycles += 19
		}
	case 0x1eaf:
		{
			m.ip = 0x1eb0
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x1eb0:
		{
			m.ip = 0x1eb2
			m.set8(2, 8, uint16(0))
			m.cycles += 2
		}
	case 0x1eb2:
		{
			m.ip = 0x1eb6
			m.r[7] = m.rd16(m.r[11], uint16(0x2069))
			m.cycles += 8
		}
	case 0x1eb6:
		{
			m.ip = 0x1eb9
			m.r[2] = m.alu("sub", 16, m.r[2], uint16(12))
			m.cycles += 4
		}
	case 0x1eb9:
		{
			m.ip = 0x1ebb
			m.r[7] = m.alu("add", 16, m.r[7], m.r[2])
			m.cycles += 4
		}
	case 0x1ebb:
		{
			m.ip = 0x1ebc
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x1ebc:
		{
			m.ip = 0x1ebe
			m.ip = uint16(7766)
			m.cycles += 15
		}
	case 0x1ebe:
		{
			m.ip = 0x1ec1
			m.r[1] = uint16(8192)
			m.cycles += 2
		}
	case 0x1ec1:
		{
			m.ip = 0x1ec3
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x1ec3:
		{
			m.ip = 0x1ec5
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x1ec5:
		{
			m.ip = 0x1ec9
			m.r[6] = m.rd16(m.r[11], uint16(0x2069))
			m.cycles += 8
		}
	case 0x1ec9:
		{
			m.ip = 0x1ecc
			m.r[6] = m.alu("add", 16, m.r[6], uint16(37))
			m.cycles += 4
		}
	case 0x1ecc:
		{
			m.ip = 0x1ecf
			m.r[1] = uint16(25)
			m.cycles += 2
		}
	case 0x1ecf:
		{
			m.ip = 0x1ed4
			m.wr8(m.r[11], uint16(0x1983), uint16(1))
			m.cycles += 8
		}
	case 0x1ed4:
		{
			m.ip = 0x1ed8
			m.r[7] = uint16(0x1e10)
			m.cycles += 3
		}
	case 0x1ed8:
		{
			m.ip = 0x1edd
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(1))
			m.cycles += 16
		}
	case 0x1edd:
		{
			m.ip = 0x1edf
			if m.zf {
				m.ip = uint16(7963)
			}
			m.cycles += 8
		}
	case 0x1edf:
		{
			m.ip = 0x1ee4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x1ee4:
		{
			m.ip = 0x1ee6
			if !m.zf {
				m.ip = uint16(7933)
			}
			m.cycles += 8
		}
	case 0x1ee6:
		{
			m.ip = 0x1eeb
			m.wr8(m.r[11], uint16(0x1983), uint16(2))
			m.cycles += 8
		}
	case 0x1eeb:
		{
			m.ip = 0x1eef
			m.r[7] = uint16(0x1e5a)
			m.cycles += 3
		}
	case 0x1eef:
		{
			m.ip = 0x1ef4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1984)), uint16(1))
			m.cycles += 16
		}
	case 0x1ef4:
		{
			m.ip = 0x1ef6
			if m.zf {
				m.ip = uint16(7963)
			}
			m.cycles += 8
		}
	case 0x1ef6:
		{
			m.ip = 0x1efa
			m.r[7] = uint16(0x1ea4)
			m.cycles += 3
		}
	case 0x1efa:
		{
			m.ip = 0x1efc
			m.ip = uint16(7963)
			m.cycles += 15
		}
	case 0x1efc:
		{
			m.ip = 0x1efd
			m.cycles += 3
		}
	case 0x1efd:
		{
			m.ip = 0x1f02
			m.wr8(m.r[11], uint16(0x1984), uint16(0))
			m.cycles += 8
		}
	case 0x1f02:
		{
			m.ip = 0x1f07
			m.wr8(m.r[11], uint16(0x1983), uint16(2))
			m.cycles += 8
		}
	case 0x1f07:
		{
			m.ip = 0x1f0b
			m.r[7] = uint16(0x1ea4)
			m.cycles += 3
		}
	case 0x1f0b:
		{
			m.ip = 0x1f10
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(2))
			m.cycles += 16
		}
	case 0x1f10:
		{
			m.ip = 0x1f12
			if m.zf {
				m.ip = uint16(7963)
			}
			m.cycles += 8
		}
	case 0x1f12:
		{
			m.ip = 0x1f17
			m.wr8(m.r[11], uint16(0x1984), uint16(1))
			m.cycles += 8
		}
	case 0x1f17:
		{
			m.ip = 0x1f1b
			m.r[7] = uint16(0x1e5a)
			m.cycles += 3
		}
	case 0x1f1b:
		{
			m.ip = 0x1f1d
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x1f1d:
		{
			m.ip = 0x1f1f
			m.wr16(m.r[11], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x1f1f:
		{
			m.ip = 0x1f20
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x1f20:
		{
			m.ip = 0x1f21
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x1f21:
		{
			m.ip = 0x1f23
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(7963)
			}
			m.cycles += 17
		}
	case 0x1f23:
		{
			m.ip = 0x1f24
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x1f24:
		{
			m.ip = 0x1f25
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x1f25:
		{
			m.ip = 0x1f27
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x1f27:
		{
			m.ip = 0x1f2a
			m.r[2] = m.alu("add", 16, m.r[2], uint16(1))
			m.cycles += 4
		}
	case 0x1f2a:
		{
			m.ip = 0x1f2e
			m.r[3] = uint16(0x1029)
			m.cycles += 3
		}
	case 0x1f2e:
		{
			m.ip = 0x1f31
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(68))
			m.cycles += 4
		}
	case 0x1f31:
		{
			m.ip = 0x1f36
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x1f36:
		{
			m.ip = 0x1f38
			if m.zf {
				m.ip = uint16(8036)
			}
			m.cycles += 8
		}
	case 0x1f38:
		{
			m.ip = 0x1f3c
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(0x1981)))
			m.cycles += 16
		}
	case 0x1f3c:
		{
			m.ip = 0x1f3e
			if !m.cf && !m.zf {
				m.ip = uint16(8036)
			}
			m.cycles += 8
		}
	case 0x1f3e:
		{
			m.ip = 0x1f43
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x1981)), uint16(10))
			m.cycles += 16
		}
	case 0x1f43:
		{
			m.ip = 0x1f45
			if m.cf || m.zf {
				m.ip = uint16(8018)
			}
			m.cycles += 8
		}
	case 0x1f45:
		{
			m.ip = 0x1f4a
			m.wr8(m.r[11], uint16(0x1e2f), uint16(32))
			m.cycles += 8
		}
	case 0x1f4a:
		{
			m.ip = 0x1f4f
			m.wr8(m.r[11], uint16(0x1e30), uint16(32))
			m.cycles += 8
		}
	case 0x1f4f:
		{
			m.ip = 0x1f51
			m.ip = uint16(8036)
			m.cycles += 15
		}
	case 0x1f51:
		{
			m.ip = 0x1f52
			m.cycles += 3
		}
	case 0x1f52:
		{
			m.ip = 0x1f53
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x1f53:
		{
			m.ip = 0x1f54
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x1f54:
		{
			m.ip = 0x1f55
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x1f55:
		{
			m.ip = 0x1f59
			m.r[2] = m.rd16(m.r[11], uint16(0x1981))
			m.cycles += 8
		}
	case 0x1f59:
		{
			m.ip = 0x1f5a
			m.r[2] = m.unary("inc", 16, m.r[2])
			m.cycles += 3
		}
	case 0x1f5a:
		{
			m.ip = 0x1f5e
			m.r[7] = uint16(0x1e2f)
			m.cycles += 3
		}
	case 0x1f5e:
		{
			m.ip = 0x1f61
			target := uint16(9036)
			m.push(0x1f61)
			m.ip = target
			m.cycles += 19
		}
	case 0x1f61:
		{
			m.ip = 0x1f62
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x1f62:
		{
			m.ip = 0x1f63
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1f63:
		{
			m.ip = 0x1f64
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x1f64:
		{
			m.ip = 0x1f68
			m.wr16(m.r[11], uint16(0x1981), m.r[2])
			m.cycles += 8
		}
	case 0x1f68:
		{
			m.ip = 0x1f6b
			target := uint16(9036)
			m.push(0x1f6b)
			m.ip = target
			m.cycles += 19
		}
	case 0x1f6b:
		{
			m.ip = 0x1f6c
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x1f6c:
		{
			m.ip = 0x1f6d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1f6d:
		{
			m.ip = 0x1f6f
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1f6f:
		{
			m.ip = 0x1f72
			target := uint16(8080)
			m.push(0x1f72)
			m.ip = target
			m.cycles += 19
		}
	case 0x1f72:
		{
			m.ip = 0x1f76
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), m.rd8(m.r[11], uint16(m.r[6]+0x2071)))
			m.cycles += 16
		}
	case 0x1f76:
		{
			m.ip = 0x1f78
			if m.zf {
				m.ip = uint16(8077)
			}
			m.cycles += 8
		}
	case 0x1f78:
		{
			m.ip = 0x1f79
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x1f79:
		{
			m.ip = 0x1f7b
			m.set8(2, 0, m.unary("inc", 8, ((m.r[2]>>0)&255)))
			m.cycles += 3
		}
	case 0x1f7b:
		{
			m.ip = 0x1f7e
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(60))
			m.cycles += 4
		}
	case 0x1f7e:
		{
			m.ip = 0x1f80
			if m.zf {
				m.ip = uint16(8056)
			}
			m.cycles += 8
		}
	case 0x1f80:
		{
			m.ip = 0x1f83
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(63))
			m.cycles += 4
		}
	case 0x1f83:
		{
			m.ip = 0x1f85
			if m.zf {
				m.ip = uint16(8056)
			}
			m.cycles += 8
		}
	case 0x1f85:
		{
			m.ip = 0x1f88
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(71))
			m.cycles += 4
		}
	case 0x1f88:
		{
			m.ip = 0x1f8a
			if m.zf {
				m.ip = uint16(8056)
			}
			m.cycles += 8
		}
	case 0x1f8a:
		{
			m.ip = 0x1f8d
			target := uint16(8092)
			m.push(0x1f8d)
			m.ip = target
			m.cycles += 19
		}
	case 0x1f8d:
		{
			m.ip = 0x1f90
			m.ip = uint16(7766)
			m.cycles += 15
		}
	case 0x1f90:
		{
			m.ip = 0x1f92
			m.set8(0, 8, uint16(9))
			m.cycles += 2
		}
	case 0x1f92:
		{
			m.ip = 0x1f94
			m.set8(3, 8, uint16(0))
			m.cycles += 2
		}
	case 0x1f94:
		{
			m.ip = 0x1f96
			m.set8(3, 0, uint16(4))
			m.cycles += 2
		}
	case 0x1f96:
		{
			m.ip = 0x1f99
			m.r[1] = uint16(1)
			m.cycles += 2
		}
	case 0x1f99:
		{
			m.ip = 0x1f9b
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x1f9b:
		{
			m.ip = 0x1f9c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1f9c:
		{
			m.ip = 0x1f9e
			m.set8(0, 8, uint16(2))
			m.cycles += 2
		}
	case 0x1f9e:
		{
			m.ip = 0x1fa0
			m.set8(3, 8, uint16(0))
			m.cycles += 2
		}
	case 0x1fa0:
		{
			m.ip = 0x1fa2
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x1fa2:
		{
			m.ip = 0x1fa3
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1fa3:
		{
			m.ip = 0x1fa8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x251)), uint16(1))
			m.cycles += 16
		}
	case 0x1fa8:
		{
			m.ip = 0x1faa
			if m.zf {
				m.ip = uint16(8127)
			}
			m.cycles += 8
		}
	case 0x1faa:
		{
			m.ip = 0x1faf
			m.wr8(m.r[11], uint16(0x251), uint16(1))
			m.cycles += 8
		}
	case 0x1faf:
		{
			m.ip = 0x1fb2
			target := uint16(9161)
			m.push(0x1fb2)
			m.ip = target
			m.cycles += 19
		}
	case 0x1fb2:
		{
			m.ip = 0x1fb6
			m.wr8(m.r[11], uint16(0x254), ((m.r[1] >> 8) & 255))
			m.cycles += 8
		}
	case 0x1fb6:
		{
			m.ip = 0x1fba
			m.wr8(m.r[11], uint16(0x255), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x1fba:
		{
			m.ip = 0x1fbe
			m.wr8(m.r[11], uint16(0x256), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x1fbe:
		{
			m.ip = 0x1fbf
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x1fbf:
		{
			m.ip = 0x1fc4
			m.wr8(m.r[11], uint16(0x251), uint16(0))
			m.cycles += 8
		}
	case 0x1fc4:
		{
			m.ip = 0x1fca
			m.wr16(m.r[11], uint16(0x252), uint16(0))
			m.cycles += 8
		}
	case 0x1fca:
		{
			m.ip = 0x1fcd
			target := uint16(9161)
			m.push(0x1fcd)
			m.ip = target
			m.cycles += 19
		}
	case 0x1fcd:
		{
			m.ip = 0x1fd1
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), m.rd8(m.r[11], uint16(0x256)))
			m.cycles += 16
		}
	case 0x1fd1:
		{
			m.ip = 0x1fd3
			if !m.cf {
				m.ip = uint16(8163)
			}
			m.cycles += 8
		}
	case 0x1fd3:
		{
			m.ip = 0x1fd8
			m.wr8(m.r[11], uint16(0x252), uint16(1))
			m.cycles += 8
		}
	case 0x1fd8:
		{
			m.ip = 0x1fda
			m.set8(0, 0, uint16(60))
			m.cycles += 2
		}
	case 0x1fda:
		{
			m.ip = 0x1fde
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(0x256))))
			m.cycles += 16
		}
	case 0x1fde:
		{
			m.ip = 0x1fe0
			m.set8(2, 8, m.alu("add", 8, ((m.r[2]>>8)&255), ((m.r[0]>>0)&255)))
			m.cycles += 4
		}
	case 0x1fe0:
		{
			m.ip = 0x1fe2
			m.ip = uint16(8167)
			m.cycles += 15
		}
	case 0x1fe2:
		{
			m.ip = 0x1fe3
			m.cycles += 3
		}
	case 0x1fe3:
		{
			m.ip = 0x1fe7
			m.set8(2, 8, m.alu("sub", 8, ((m.r[2]>>8)&255), m.rd8(m.r[11], uint16(0x256))))
			m.cycles += 16
		}
	case 0x1fe7:
		{
			m.ip = 0x1feb
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), m.rd8(m.r[11], uint16(0x255)))
			m.cycles += 16
		}
	case 0x1feb:
		{
			m.ip = 0x1fed
			if !m.cf {
				m.ip = uint16(8189)
			}
			m.cycles += 8
		}
	case 0x1fed:
		{
			m.ip = 0x1ff2
			m.wr8(m.r[11], uint16(0x253), uint16(1))
			m.cycles += 8
		}
	case 0x1ff2:
		{
			m.ip = 0x1ff4
			m.set8(0, 0, uint16(60))
			m.cycles += 2
		}
	case 0x1ff4:
		{
			m.ip = 0x1ff8
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(0x255))))
			m.cycles += 16
		}
	case 0x1ff8:
		{
			m.ip = 0x1ffa
			m.set8(1, 0, m.alu("add", 8, ((m.r[1]>>0)&255), ((m.r[0]>>0)&255)))
			m.cycles += 4
		}
	case 0x1ffa:
		{
			m.ip = 0x1ffc
			m.ip = uint16(8202)
			m.cycles += 15
		}
	case 0x1ffc:
		{
			m.ip = 0x1ffd
			m.cycles += 3
		}
	case 0x1ffd:
		{
			m.ip = 0x2001
			m.set8(1, 0, m.alu("sub", 8, ((m.r[1]>>0)&255), m.rd8(m.r[11], uint16(0x255))))
			m.cycles += 16
		}
	case 0x2001:
		{
			m.ip = 0x2006
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x252)), uint16(0))
			m.cycles += 16
		}
	case 0x2006:
		{
			m.ip = 0x2008
			if m.zf {
				m.ip = uint16(8202)
			}
			m.cycles += 8
		}
	case 0x2008:
		{
			m.ip = 0x200a
			m.set8(1, 0, m.unary("dec", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x200a:
		{
			m.ip = 0x200e
			m.alu("sub", 8, ((m.r[1] >> 8) & 255), m.rd8(m.r[11], uint16(0x254)))
			m.cycles += 16
		}
	case 0x200e:
		{
			m.ip = 0x2010
			if !m.cf {
				m.ip = uint16(8228)
			}
			m.cycles += 8
		}
	case 0x2010:
		{
			m.ip = 0x2012
			m.set8(0, 0, uint16(24))
			m.cycles += 2
		}
	case 0x2012:
		{
			m.ip = 0x2017
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x252)), uint16(0))
			m.cycles += 16
		}
	case 0x2017:
		{
			m.ip = 0x2019
			if m.zf {
				m.ip = uint16(8219)
			}
			m.cycles += 8
		}
	case 0x2019:
		{
			m.ip = 0x201b
			m.set8(0, 0, m.unary("dec", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x201b:
		{
			m.ip = 0x201f
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(0x254))))
			m.cycles += 16
		}
	case 0x201f:
		{
			m.ip = 0x2021
			m.set8(1, 8, m.alu("add", 8, ((m.r[1]>>8)&255), ((m.r[0]>>0)&255)))
			m.cycles += 4
		}
	case 0x2021:
		{
			m.ip = 0x2023
			m.ip = uint16(8241)
			m.cycles += 15
		}
	case 0x2023:
		{
			m.ip = 0x2024
			m.cycles += 3
		}
	case 0x2024:
		{
			m.ip = 0x2028
			m.set8(1, 8, m.alu("sub", 8, ((m.r[1]>>8)&255), m.rd8(m.r[11], uint16(0x254))))
			m.cycles += 16
		}
	case 0x2028:
		{
			m.ip = 0x202d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x253)), uint16(0))
			m.cycles += 16
		}
	case 0x202d:
		{
			m.ip = 0x202f
			if m.zf {
				m.ip = uint16(8241)
			}
			m.cycles += 8
		}
	case 0x202f:
		{
			m.ip = 0x2031
			m.set8(1, 8, m.unary("dec", 8, ((m.r[1]>>8)&255)))
			m.cycles += 3
		}
	case 0x2031:
		{
			m.ip = 0x2035
			m.r[6] = uint16(0x25f)
			m.cycles += 3
		}
	case 0x2035:
		{
			m.ip = 0x203a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x203a:
		{
			m.ip = 0x203c
			if m.zf {
				m.ip = uint16(8263)
			}
			m.cycles += 8
		}
	case 0x203c:
		{
			m.ip = 0x2041
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x2041:
		{
			m.ip = 0x2043
			if m.zf {
				m.ip = uint16(8263)
			}
			m.cycles += 8
		}
	case 0x2043:
		{
			m.ip = 0x2047
			m.r[6] = uint16(0x268)
			m.cycles += 3
		}
	case 0x2047:
		{
			m.ip = 0x204a
			m.set8(2, 8, m.alu("add", 8, ((m.r[2]>>8)&255), m.rd8(m.r[11], uint16(m.r[6]+0x2))))
			m.cycles += 16
		}
	case 0x204a:
		{
			m.ip = 0x204d
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(60))
			m.cycles += 4
		}
	case 0x204d:
		{
			m.ip = 0x204f
			if m.cf {
				m.ip = uint16(8276)
			}
			m.cycles += 8
		}
	case 0x204f:
		{
			m.ip = 0x2051
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x2051:
		{
			m.ip = 0x2054
			m.set8(2, 8, m.alu("sub", 8, ((m.r[2]>>8)&255), uint16(60)))
			m.cycles += 4
		}
	case 0x2054:
		{
			m.ip = 0x2057
			m.wr8(m.r[11], uint16(m.r[6]+0x2), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x2057:
		{
			m.ip = 0x205a
			m.set8(1, 0, m.alu("add", 8, ((m.r[1]>>0)&255), m.rd8(m.r[11], uint16(m.r[6]+0x1))))
			m.cycles += 16
		}
	case 0x205a:
		{
			m.ip = 0x205d
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), uint16(60))
			m.cycles += 4
		}
	case 0x205d:
		{
			m.ip = 0x205f
			if m.cf {
				m.ip = uint16(8292)
			}
			m.cycles += 8
		}
	case 0x205f:
		{
			m.ip = 0x2061
			m.set8(1, 8, m.unary("inc", 8, ((m.r[1]>>8)&255)))
			m.cycles += 3
		}
	case 0x2061:
		{
			m.ip = 0x2064
			m.set8(1, 0, m.alu("sub", 8, ((m.r[1]>>0)&255), uint16(60)))
			m.cycles += 4
		}
	case 0x2064:
		{
			m.ip = 0x2067
			m.wr8(m.r[11], uint16(m.r[6]+0x1), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x2067:
		{
			m.ip = 0x2069
			m.wr8(m.r[11], uint16(m.r[6]), m.alu("add", 8, m.rd8(m.r[11], uint16(m.r[6])), ((m.r[1]>>8)&255)))
			m.cycles += 16
		}
	case 0x2069:
		{
			m.ip = 0x206e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x206e:
		{
			m.ip = 0x2070
			if !m.zf {
				m.ip = uint16(8321)
			}
			m.cycles += 8
		}
	case 0x2070:
		{
			m.ip = 0x2072
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x2072:
		{
			m.ip = 0x2075
			m.wr8(m.r[11], uint16(m.r[6]+0x9), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x2075:
		{
			m.ip = 0x2078
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x2078:
		{
			m.ip = 0x207b
			m.wr8(m.r[11], uint16(m.r[6]+0xa), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x207b:
		{
			m.ip = 0x207e
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x207e:
		{
			m.ip = 0x2081
			m.wr8(m.r[11], uint16(m.r[6]+0xb), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x2081:
		{
			m.ip = 0x2082
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2082:
		{
			m.ip = 0x2083
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x2083:
		{
			m.ip = 0x2084
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x2084:
		{
			m.ip = 0x2088
			m.r[6] = uint16(0x17fe)
			m.cycles += 3
		}
	case 0x2088:
		{
			m.ip = 0x208a
			m.r[7] = m.r[6]
			m.cycles += 2
		}
	case 0x208a:
		{
			m.ip = 0x208d
			m.r[7] = m.alu("add", 16, m.r[7], uint16(52))
			m.cycles += 4
		}
	case 0x208d:
		{
			m.ip = 0x2090
			m.r[1] = uint16(9)
			m.cycles += 2
		}
	case 0x2090:
		{
			m.ip = 0x2092
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2092:
		{
			m.ip = 0x2094
			m.r[1] = m.alu("sub", 16, m.r[1], m.r[3])
			m.cycles += 4
		}
	case 0x2094:
		{
			m.ip = 0x2097
			m.alu("sub", 16, m.r[1], uint16(0))
			m.cycles += 4
		}
	case 0x2097:
		{
			m.ip = 0x2099
			if m.zf {
				m.ip = uint16(8365)
			}
			m.cycles += 8
		}
	case 0x2099:
		{
			m.ip = 0x209a
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x209a:
		{
			m.ip = 0x209b
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x209b:
		{
			m.ip = 0x209c
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x209c:
		{
			m.ip = 0x209f
			m.r[1] = uint16(46)
			m.cycles += 2
		}
	case 0x209f:
		{
			m.ip = 0x20a0
			m.df = false
			m.cycles += 2
		}
	case 0x20a0:
		{
			m.ip = 0x20a2
			m.stringOp("movs", 8)
			m.cycles += 2
		}
	case 0x20a2:
		{
			m.ip = 0x20a3
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x20a3:
		{
			m.ip = 0x20a4
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x20a4:
		{
			m.ip = 0x20a7
			m.r[6] = m.alu("sub", 16, m.r[6], uint16(52))
			m.cycles += 4
		}
	case 0x20a7:
		{
			m.ip = 0x20aa
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(52))
			m.cycles += 4
		}
	case 0x20aa:
		{
			m.ip = 0x20ab
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x20ab:
		{
			m.ip = 0x20ad
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(8345)
			}
			m.cycles += 17
		}
	case 0x20ad:
		{
			m.ip = 0x20b1
			m.r[8] = m.rd16(m.r[11], uint16(0x276))
			m.cycles += 8
		}
	case 0x20b1:
		{
			m.ip = 0x20b3
			m.r[5] = m.r[7]
			m.cycles += 2
		}
	case 0x20b3:
		{
			m.ip = 0x20b7
			m.r[6] = uint16(0x262)
			m.cycles += 3
		}
	case 0x20b7:
		{
			m.ip = 0x20bc
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(2))
			m.cycles += 16
		}
	case 0x20bc:
		{
			m.ip = 0x20be
			if m.zf {
				m.ip = uint16(8393)
			}
			m.cycles += 8
		}
	case 0x20be:
		{
			m.ip = 0x20c3
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x223)), uint16(1))
			m.cycles += 16
		}
	case 0x20c3:
		{
			m.ip = 0x20c5
			if m.zf {
				m.ip = uint16(8393)
			}
			m.cycles += 8
		}
	case 0x20c5:
		{
			m.ip = 0x20c9
			m.r[6] = uint16(0x259)
			m.cycles += 3
		}
	case 0x20c9:
		{
			m.ip = 0x20cc
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x6)))
			m.cycles += 8
		}
	case 0x20cc:
		{
			m.ip = 0x20ce
			m.r[7] = m.r[5]
			m.cycles += 2
		}
	case 0x20ce:
		{
			m.ip = 0x20d1
			target := uint16(9120)
			m.push(0x20d1)
			m.ip = target
			m.cycles += 19
		}
	case 0x20d1:
		{
			m.ip = 0x20d4
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x7)))
			m.cycles += 8
		}
	case 0x20d4:
		{
			m.ip = 0x20d7
			target := uint16(9120)
			m.push(0x20d7)
			m.ip = target
			m.cycles += 19
		}
	case 0x20d7:
		{
			m.ip = 0x20da
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x8)))
			m.cycles += 8
		}
	case 0x20da:
		{
			m.ip = 0x20dd
			target := uint16(9120)
			m.push(0x20dd)
			m.ip = target
			m.cycles += 19
		}
	case 0x20dd:
		{
			m.ip = 0x20de
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x20de:
		{
			m.ip = 0x20e2
			m.r[6] = uint16(0x466)
			m.cycles += 3
		}
	case 0x20e2:
		{
			m.ip = 0x20e7
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(1))
			m.cycles += 16
		}
	case 0x20e7:
		{
			m.ip = 0x20e9
			if m.zf {
				m.ip = uint16(8447)
			}
			m.cycles += 8
		}
	case 0x20e9:
		{
			m.ip = 0x20ee
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x20ee:
		{
			m.ip = 0x20f0
			if m.zf {
				m.ip = uint16(8447)
			}
			m.cycles += 8
		}
	case 0x20f0:
		{
			m.ip = 0x20f4
			m.r[6] = uint16(0x23c)
			m.cycles += 3
		}
	case 0x20f4:
		{
			m.ip = 0x20f9
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x20f9:
		{
			m.ip = 0x20fb
			if m.zf {
				m.ip = uint16(8447)
			}
			m.cycles += 8
		}
	case 0x20fb:
		{
			m.ip = 0x20ff
			m.r[6] = uint16(0x240)
			m.cycles += 3
		}
	case 0x20ff:
		{
			m.ip = 0x2101
			m.r[7] = m.r[5]
			m.cycles += 2
		}
	case 0x2101:
		{
			m.ip = 0x2103
			m.r[2] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x2103:
		{
			m.ip = 0x2105
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x2105:
		{
			m.ip = 0x2106
			m.r[2] = m.unary("inc", 16, m.r[2])
			m.cycles += 3
		}
	case 0x2106:
		{
			m.ip = 0x2109
			m.r[7] = m.alu("add", 16, m.r[7], uint16(11))
			m.cycles += 4
		}
	case 0x2109:
		{
			m.ip = 0x210d
			m.r[3] = uint16(0x1029)
			m.cycles += 3
		}
	case 0x210d:
		{
			m.ip = 0x2110
			target := uint16(9036)
			m.push(0x2110)
			m.ip = target
			m.cycles += 19
		}
	case 0x2110:
		{
			m.ip = 0x2113
			m.r[2] = m.rd16(m.r[11], uint16(m.r[6]+0x2))
			m.cycles += 8
		}
	case 0x2113:
		{
			m.ip = 0x2116
			m.r[7] = m.alu("add", 16, m.r[7], uint16(3))
			m.cycles += 4
		}
	case 0x2116:
		{
			m.ip = 0x211a
			m.r[3] = uint16(0x1029)
			m.cycles += 3
		}
	case 0x211a:
		{
			m.ip = 0x211d
			target := uint16(9036)
			m.push(0x211d)
			m.ip = target
			m.cycles += 19
		}
	case 0x211d:
		{
			m.ip = 0x211e
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x211e:
		{
			m.ip = 0x2120
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x2120:
		{
			m.ip = 0x2123
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x6)))
			m.cycles += 8
		}
	case 0x2123:
		{
			m.ip = 0x2125
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(17))
			m.cycles += 4
		}
	case 0x2125:
		{
			m.ip = 0x2127
			if m.cf || m.zf {
				m.ip = uint16(8489)
			}
			m.cycles += 8
		}
	case 0x2127:
		{
			m.ip = 0x2129
			m.set8(0, 0, uint16(17))
			m.cycles += 2
		}
	case 0x2129:
		{
			m.ip = 0x212c
			m.r[1] = uint16(3600)
			m.cycles += 2
		}
	case 0x212c:
		{
			m.ip = 0x212e
			m.multiplyDivide("mul", m.r[1])
			m.cycles += 100
		}
	case 0x212e:
		{
			m.ip = 0x2130
			m.r[2] = m.r[0]
			m.cycles += 2
		}
	case 0x2130:
		{
			m.ip = 0x2132
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x2132:
		{
			m.ip = 0x2135
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x7)))
			m.cycles += 8
		}
	case 0x2135:
		{
			m.ip = 0x2138
			m.r[1] = uint16(60)
			m.cycles += 2
		}
	case 0x2138:
		{
			m.ip = 0x213a
			m.multiplyDivide("mul", m.r[1])
			m.cycles += 100
		}
	case 0x213a:
		{
			m.ip = 0x213c
			m.r[2] = m.alu("add", 16, m.r[2], m.r[0])
			m.cycles += 4
		}
	case 0x213c:
		{
			m.ip = 0x213e
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x213e:
		{
			m.ip = 0x2141
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x8)))
			m.cycles += 8
		}
	case 0x2141:
		{
			m.ip = 0x2143
			m.r[2] = m.alu("add", 16, m.r[2], m.r[0])
			m.cycles += 4
		}
	case 0x2143:
		{
			m.ip = 0x2145
			m.r[1] = m.r[2]
			m.cycles += 2
		}
	case 0x2145:
		{
			m.ip = 0x2148
			m.r[0] = uint16(0)
			m.cycles += 2
		}
	case 0x2148:
		{
			m.ip = 0x214a
			m.r[2] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x214a:
		{
			m.ip = 0x214c
			m.set8(0, 8, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x214c:
		{
			m.ip = 0x214e
			m.set8(2, 0, ((m.r[2] >> 8) & 255))
			m.cycles += 2
		}
	case 0x214e:
		{
			m.ip = 0x2150
			m.set8(2, 8, uint16(0))
			m.cycles += 2
		}
	case 0x2150:
		{
			m.ip = 0x2152
			m.multiplyDivide("div", m.r[1])
			m.cycles += 100
		}
	case 0x2152:
		{
			m.ip = 0x2154
			m.r[2] = m.r[0]
			m.cycles += 2
		}
	case 0x2154:
		{
			m.ip = 0x2156
			m.r[7] = m.r[5]
			m.cycles += 2
		}
	case 0x2156:
		{
			m.ip = 0x2159
			m.r[7] = m.alu("add", 16, m.r[7], uint16(21))
			m.cycles += 4
		}
	case 0x2159:
		{
			m.ip = 0x215d
			m.r[3] = uint16(0x1023)
			m.cycles += 3
		}
	case 0x215d:
		{
			m.ip = 0x2160
			target := uint16(9036)
			m.push(0x2160)
			m.ip = target
			m.cycles += 19
		}
	case 0x2160:
		{
			m.ip = 0x2162
			m.r[7] = m.r[5]
			m.cycles += 2
		}
	case 0x2162:
		{
			m.ip = 0x2164
			m.r[2] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x2164:
		{
			m.ip = 0x2167
			m.r[7] = m.alu("add", 16, m.r[7], uint16(30))
			m.cycles += 4
		}
	case 0x2167:
		{
			m.ip = 0x216b
			m.r[3] = uint16(0x1023)
			m.cycles += 3
		}
	case 0x216b:
		{
			m.ip = 0x216e
			target := uint16(9036)
			m.push(0x216e)
			m.ip = target
			m.cycles += 19
		}
	case 0x216e:
		{
			m.ip = 0x2170
			m.r[7] = m.r[5]
			m.cycles += 2
		}
	case 0x2170:
		{
			m.ip = 0x2173
			m.r[2] = m.rd16(m.r[11], uint16(m.r[6]+0x2))
			m.cycles += 8
		}
	case 0x2173:
		{
			m.ip = 0x2177
			m.alu("sub", 16, m.r[2], uint16(999))
			m.cycles += 4
		}
	case 0x2177:
		{
			m.ip = 0x2179
			if m.cf || m.zf {
				m.ip = uint16(8572)
			}
			m.cycles += 8
		}
	case 0x2179:
		{
			m.ip = 0x217c
			m.r[2] = uint16(999)
			m.cycles += 2
		}
	case 0x217c:
		{
			m.ip = 0x217f
			m.r[7] = m.alu("add", 16, m.r[7], uint16(38))
			m.cycles += 4
		}
	case 0x217f:
		{
			m.ip = 0x2183
			m.r[3] = uint16(0x1027)
			m.cycles += 3
		}
	case 0x2183:
		{
			m.ip = 0x2186
			target := uint16(9036)
			m.push(0x2186)
			m.ip = target
			m.cycles += 19
		}
	case 0x2186:
		{
			m.ip = 0x2189
			m.r[2] = m.rd16(m.r[11], uint16(m.r[6]+0x4))
			m.cycles += 8
		}
	case 0x2189:
		{
			m.ip = 0x218c
			m.alu("sub", 16, m.r[2], uint16(99))
			m.cycles += 4
		}
	case 0x218c:
		{
			m.ip = 0x218e
			if m.cf || m.zf {
				m.ip = uint16(8593)
			}
			m.cycles += 8
		}
	case 0x218e:
		{
			m.ip = 0x2191
			m.r[2] = uint16(99)
			m.cycles += 2
		}
	case 0x2191:
		{
			m.ip = 0x2194
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x2194:
		{
			m.ip = 0x2198
			m.r[3] = uint16(0x1029)
			m.cycles += 3
		}
	case 0x2198:
		{
			m.ip = 0x219b
			target := uint16(9036)
			m.push(0x219b)
			m.ip = target
			m.cycles += 19
		}
	case 0x219b:
		{
			m.ip = 0x21a0
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1980)), uint16(1))
			m.cycles += 16
		}
	case 0x21a0:
		{
			m.ip = 0x21a2
			if m.zf {
				m.ip = uint16(8702)
			}
			m.cycles += 8
		}
	case 0x21a2:
		{
			m.ip = 0x21a4
			m.r[6] = m.r[5]
			m.cycles += 2
		}
	case 0x21a4:
		{
			m.ip = 0x21a7
			m.r[5] = m.alu("add", 16, m.r[5], uint16(3))
			m.cycles += 4
		}
	case 0x21a7:
		{
			m.ip = 0x21aa
			m.r[1] = uint16(47)
			m.cycles += 2
		}
	case 0x21aa:
		{
			m.ip = 0x21af
			m.wr8(m.r[11], uint16(0x1983), uint16(1))
			m.cycles += 8
		}
	case 0x21af:
		{
			m.ip = 0x21b3
			m.r[7] = uint16(0x189b)
			m.cycles += 3
		}
	case 0x21b3:
		{
			m.ip = 0x21b8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(1))
			m.cycles += 16
		}
	case 0x21b8:
		{
			m.ip = 0x21ba
			if m.zf {
				m.ip = uint16(8694)
			}
			m.cycles += 8
		}
	case 0x21ba:
		{
			m.ip = 0x21bf
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x21bf:
		{
			m.ip = 0x21c1
			if !m.zf {
				m.ip = uint16(8664)
			}
			m.cycles += 8
		}
	case 0x21c1:
		{
			m.ip = 0x21c6
			m.wr8(m.r[11], uint16(0x1983), uint16(2))
			m.cycles += 8
		}
	case 0x21c6:
		{
			m.ip = 0x21ca
			m.r[7] = uint16(0x18d0)
			m.cycles += 3
		}
	case 0x21ca:
		{
			m.ip = 0x21cf
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1984)), uint16(1))
			m.cycles += 16
		}
	case 0x21cf:
		{
			m.ip = 0x21d1
			if m.zf {
				m.ip = uint16(8694)
			}
			m.cycles += 8
		}
	case 0x21d1:
		{
			m.ip = 0x21d5
			m.r[7] = uint16(0x1905)
			m.cycles += 3
		}
	case 0x21d5:
		{
			m.ip = 0x21d7
			m.ip = uint16(8694)
			m.cycles += 15
		}
	case 0x21d7:
		{
			m.ip = 0x21d8
			m.cycles += 3
		}
	case 0x21d8:
		{
			m.ip = 0x21dd
			m.wr8(m.r[11], uint16(0x1984), uint16(0))
			m.cycles += 8
		}
	case 0x21dd:
		{
			m.ip = 0x21e2
			m.wr8(m.r[11], uint16(0x1983), uint16(2))
			m.cycles += 8
		}
	case 0x21e2:
		{
			m.ip = 0x21e6
			m.r[7] = uint16(0x1905)
			m.cycles += 3
		}
	case 0x21e6:
		{
			m.ip = 0x21eb
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(2))
			m.cycles += 16
		}
	case 0x21eb:
		{
			m.ip = 0x21ed
			if m.zf {
				m.ip = uint16(8694)
			}
			m.cycles += 8
		}
	case 0x21ed:
		{
			m.ip = 0x21f2
			m.wr8(m.r[11], uint16(0x1984), uint16(1))
			m.cycles += 8
		}
	case 0x21f2:
		{
			m.ip = 0x21f6
			m.r[7] = uint16(0x18d0)
			m.cycles += 3
		}
	case 0x21f6:
		{
			m.ip = 0x21f8
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x21f8:
		{
			m.ip = 0x21fa
			m.wr16(m.r[11], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x21fa:
		{
			m.ip = 0x21fb
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x21fb:
		{
			m.ip = 0x21fc
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x21fc:
		{
			m.ip = 0x21fe
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(8694)
			}
			m.cycles += 17
		}
	case 0x21fe:
		{
			m.ip = 0x2201
			m.r[0] = m.rd16(m.r[11], uint16(0x1978))
			m.cycles += 8
		}
	case 0x2201:
		{
			m.ip = 0x2205
			m.r[3] = m.rd16(m.r[11], uint16(0x197a))
			m.cycles += 8
		}
	case 0x2205:
		{
			m.ip = 0x2209
			m.r[1] = m.rd16(m.r[11], uint16(0x197c))
			m.cycles += 8
		}
	case 0x2209:
		{
			m.ip = 0x220d
			m.r[2] = m.rd16(m.r[11], uint16(0x197e))
			m.cycles += 8
		}
	case 0x220d:
		{
			m.ip = 0x2211
			m.r[6] = uint16(0x25f)
			m.cycles += 3
		}
	case 0x2211:
		{
			m.ip = 0x2216
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x2216:
		{
			m.ip = 0x2218
			if !m.zf {
				m.ip = uint16(8735)
			}
			m.cycles += 8
		}
	case 0x2218:
		{
			m.ip = 0x221d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1984)), uint16(0))
			m.cycles += 16
		}
	case 0x221d:
		{
			m.ip = 0x221f
			if m.zf {
				m.ip = uint16(8746)
			}
			m.cycles += 8
		}
	case 0x221f:
		{
			m.ip = 0x2224
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x2224:
		{
			m.ip = 0x2226
			if m.zf {
				m.ip = uint16(8746)
			}
			m.cycles += 8
		}
	case 0x2226:
		{
			m.ip = 0x222a
			m.r[6] = uint16(0x268)
			m.cycles += 3
		}
	case 0x222a:
		{
			m.ip = 0x222d
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), m.rd8(m.r[11], uint16(m.r[6]+0x2))))
			m.cycles += 16
		}
	case 0x222d:
		{
			m.ip = 0x2230
			m.alu("sub", 16, m.r[2], uint16(60))
			m.cycles += 4
		}
	case 0x2230:
		{
			m.ip = 0x2232
			if m.cf {
				m.ip = uint16(8759)
			}
			m.cycles += 8
		}
	case 0x2232:
		{
			m.ip = 0x2234
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x2234:
		{
			m.ip = 0x2237
			m.r[2] = m.alu("sub", 16, m.r[2], uint16(60))
			m.cycles += 4
		}
	case 0x2237:
		{
			m.ip = 0x223b
			m.wr16(m.r[11], uint16(0x197e), m.r[2])
			m.cycles += 8
		}
	case 0x223b:
		{
			m.ip = 0x223e
			m.set8(1, 0, m.alu("add", 8, ((m.r[1]>>0)&255), m.rd8(m.r[11], uint16(m.r[6]+0x1))))
			m.cycles += 16
		}
	case 0x223e:
		{
			m.ip = 0x2241
			m.alu("sub", 16, m.r[1], uint16(60))
			m.cycles += 4
		}
	case 0x2241:
		{
			m.ip = 0x2243
			if m.cf {
				m.ip = uint16(8776)
			}
			m.cycles += 8
		}
	case 0x2243:
		{
			m.ip = 0x2245
			m.set8(3, 0, m.unary("inc", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x2245:
		{
			m.ip = 0x2248
			m.r[1] = m.alu("sub", 16, m.r[1], uint16(60))
			m.cycles += 4
		}
	case 0x2248:
		{
			m.ip = 0x224c
			m.wr16(m.r[11], uint16(0x197c), m.r[1])
			m.cycles += 8
		}
	case 0x224c:
		{
			m.ip = 0x224e
			m.set8(3, 0, m.alu("add", 8, ((m.r[3]>>0)&255), m.rd8(m.r[11], uint16(m.r[6]))))
			m.cycles += 16
		}
	case 0x224e:
		{
			m.ip = 0x2251
			m.alu("sub", 16, m.r[3], uint16(24))
			m.cycles += 4
		}
	case 0x2251:
		{
			m.ip = 0x2253
			if m.cf {
				m.ip = uint16(8793)
			}
			m.cycles += 8
		}
	case 0x2253:
		{
			m.ip = 0x2256
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(24))
			m.cycles += 4
		}
	case 0x2256:
		{
			m.ip = 0x2257
			m.r[0] = m.unary("inc", 16, m.r[0])
			m.cycles += 3
		}
	case 0x2257:
		{
			m.ip = 0x2259
			m.ip = uint16(8782)
			m.cycles += 15
		}
	case 0x2259:
		{
			m.ip = 0x225d
			m.wr16(m.r[11], uint16(0x197a), m.r[3])
			m.cycles += 8
		}
	case 0x225d:
		{
			m.ip = 0x2260
			m.alu("sub", 16, m.r[0], uint16(7))
			m.cycles += 4
		}
	case 0x2260:
		{
			m.ip = 0x2262
			if m.cf || m.zf {
				m.ip = uint16(8809)
			}
			m.cycles += 8
		}
	case 0x2262:
		{
			m.ip = 0x2265
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(7))
			m.cycles += 4
		}
	case 0x2265:
		{
			m.ip = 0x2269
			m.wr16(m.r[11], uint16(0x1976), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x1976))))
			m.cycles += 15
		}
	case 0x2269:
		{
			m.ip = 0x226c
			m.wr16(m.r[11], uint16(0x1978), m.r[0])
			m.cycles += 8
		}
	case 0x226c:
		{
			m.ip = 0x226d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x226d:
		{
			m.ip = 0x2270
			target := uint16(9161)
			m.push(0x2270)
			m.ip = target
			m.cycles += 19
		}
	case 0x2270:
		{
			m.ip = 0x2271
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x2271:
		{
			m.ip = 0x2272
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x2272:
		{
			m.ip = 0x2278
			m.wr16(m.r[11], uint16(0x13a1), uint16(0))
			m.cycles += 8
		}
	case 0x2278:
		{
			m.ip = 0x227c
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), m.rd8(m.r[11], uint16(0x13a0)))
			m.cycles += 16
		}
	case 0x227c:
		{
			m.ip = 0x227e
			if !m.cf {
				m.ip = uint16(8846)
			}
			m.cycles += 8
		}
	case 0x227e:
		{
			m.ip = 0x2283
			m.wr8(m.r[11], uint16(0x13a1), uint16(1))
			m.cycles += 8
		}
	case 0x2283:
		{
			m.ip = 0x2285
			m.set8(0, 0, uint16(60))
			m.cycles += 2
		}
	case 0x2285:
		{
			m.ip = 0x2289
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(0x13a0))))
			m.cycles += 16
		}
	case 0x2289:
		{
			m.ip = 0x228b
			m.set8(2, 8, m.alu("add", 8, ((m.r[2]>>8)&255), ((m.r[0]>>0)&255)))
			m.cycles += 4
		}
	case 0x228b:
		{
			m.ip = 0x228d
			m.ip = uint16(8850)
			m.cycles += 15
		}
	case 0x228d:
		{
			m.ip = 0x228e
			m.cycles += 3
		}
	case 0x228e:
		{
			m.ip = 0x2292
			m.set8(2, 8, m.alu("sub", 8, ((m.r[2]>>8)&255), m.rd8(m.r[11], uint16(0x13a0))))
			m.cycles += 16
		}
	case 0x2292:
		{
			m.ip = 0x2296
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), m.rd8(m.r[11], uint16(0x139f)))
			m.cycles += 16
		}
	case 0x2296:
		{
			m.ip = 0x2298
			if !m.cf {
				m.ip = uint16(8872)
			}
			m.cycles += 8
		}
	case 0x2298:
		{
			m.ip = 0x229d
			m.wr8(m.r[11], uint16(0x13a2), uint16(1))
			m.cycles += 8
		}
	case 0x229d:
		{
			m.ip = 0x229f
			m.set8(0, 0, uint16(60))
			m.cycles += 2
		}
	case 0x229f:
		{
			m.ip = 0x22a3
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(0x139f))))
			m.cycles += 16
		}
	case 0x22a3:
		{
			m.ip = 0x22a5
			m.set8(1, 0, m.alu("add", 8, ((m.r[1]>>0)&255), ((m.r[0]>>0)&255)))
			m.cycles += 4
		}
	case 0x22a5:
		{
			m.ip = 0x22a7
			m.ip = uint16(8885)
			m.cycles += 15
		}
	case 0x22a7:
		{
			m.ip = 0x22a8
			m.cycles += 3
		}
	case 0x22a8:
		{
			m.ip = 0x22ac
			m.set8(1, 0, m.alu("sub", 8, ((m.r[1]>>0)&255), m.rd8(m.r[11], uint16(0x139f))))
			m.cycles += 16
		}
	case 0x22ac:
		{
			m.ip = 0x22b1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x13a1)), uint16(0))
			m.cycles += 16
		}
	case 0x22b1:
		{
			m.ip = 0x22b3
			if m.zf {
				m.ip = uint16(8885)
			}
			m.cycles += 8
		}
	case 0x22b3:
		{
			m.ip = 0x22b5
			m.set8(1, 0, m.unary("dec", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x22b5:
		{
			m.ip = 0x22b9
			m.alu("sub", 8, ((m.r[1] >> 8) & 255), m.rd8(m.r[11], uint16(0x139e)))
			m.cycles += 16
		}
	case 0x22b9:
		{
			m.ip = 0x22bb
			if !m.cf {
				m.ip = uint16(8911)
			}
			m.cycles += 8
		}
	case 0x22bb:
		{
			m.ip = 0x22bd
			m.set8(0, 0, uint16(24))
			m.cycles += 2
		}
	case 0x22bd:
		{
			m.ip = 0x22c2
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x13a1)), uint16(0))
			m.cycles += 16
		}
	case 0x22c2:
		{
			m.ip = 0x22c4
			if m.zf {
				m.ip = uint16(8902)
			}
			m.cycles += 8
		}
	case 0x22c4:
		{
			m.ip = 0x22c6
			m.set8(0, 0, m.unary("dec", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x22c6:
		{
			m.ip = 0x22ca
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(0x139e))))
			m.cycles += 16
		}
	case 0x22ca:
		{
			m.ip = 0x22cc
			m.set8(1, 8, m.alu("add", 8, ((m.r[1]>>8)&255), ((m.r[0]>>0)&255)))
			m.cycles += 4
		}
	case 0x22cc:
		{
			m.ip = 0x22ce
			m.ip = uint16(8924)
			m.cycles += 15
		}
	case 0x22ce:
		{
			m.ip = 0x22cf
			m.cycles += 3
		}
	case 0x22cf:
		{
			m.ip = 0x22d3
			m.set8(1, 8, m.alu("sub", 8, ((m.r[1]>>8)&255), m.rd8(m.r[11], uint16(0x139e))))
			m.cycles += 16
		}
	case 0x22d3:
		{
			m.ip = 0x22d8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x13a2)), uint16(0))
			m.cycles += 16
		}
	case 0x22d8:
		{
			m.ip = 0x22da
			if m.zf {
				m.ip = uint16(8924)
			}
			m.cycles += 8
		}
	case 0x22da:
		{
			m.ip = 0x22dc
			m.set8(1, 8, m.unary("dec", 8, ((m.r[1]>>8)&255)))
			m.cycles += 3
		}
	case 0x22dc:
		{
			m.ip = 0x22e0
			m.wr8(m.r[11], uint16(0x139e), ((m.r[1] >> 8) & 255))
			m.cycles += 8
		}
	case 0x22e0:
		{
			m.ip = 0x22e4
			m.wr8(m.r[11], uint16(0x139f), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x22e4:
		{
			m.ip = 0x22e8
			m.wr8(m.r[11], uint16(0x13a0), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x22e8:
		{
			m.ip = 0x22eb
			m.r[0] = m.rd16(m.r[11], uint16(0x196e))
			m.cycles += 8
		}
	case 0x22eb:
		{
			m.ip = 0x22ef
			m.r[3] = m.rd16(m.r[11], uint16(0x1970))
			m.cycles += 8
		}
	case 0x22ef:
		{
			m.ip = 0x22f3
			m.r[1] = m.rd16(m.r[11], uint16(0x1972))
			m.cycles += 8
		}
	case 0x22f3:
		{
			m.ip = 0x22f7
			m.r[2] = m.rd16(m.r[11], uint16(0x1974))
			m.cycles += 8
		}
	case 0x22f7:
		{
			m.ip = 0x22fb
			m.r[6] = uint16(0x139e)
			m.cycles += 3
		}
	case 0x22fb:
		{
			m.ip = 0x22fe
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), m.rd8(m.r[11], uint16(m.r[6]+0x2))))
			m.cycles += 16
		}
	case 0x22fe:
		{
			m.ip = 0x2301
			m.alu("sub", 16, m.r[2], uint16(60))
			m.cycles += 4
		}
	case 0x2301:
		{
			m.ip = 0x2303
			if m.cf {
				m.ip = uint16(8968)
			}
			m.cycles += 8
		}
	case 0x2303:
		{
			m.ip = 0x2305
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x2305:
		{
			m.ip = 0x2308
			m.r[2] = m.alu("sub", 16, m.r[2], uint16(60))
			m.cycles += 4
		}
	case 0x2308:
		{
			m.ip = 0x230c
			m.wr16(m.r[11], uint16(0x1974), m.r[2])
			m.cycles += 8
		}
	case 0x230c:
		{
			m.ip = 0x230f
			m.set8(1, 0, m.alu("add", 8, ((m.r[1]>>0)&255), m.rd8(m.r[11], uint16(m.r[6]+0x1))))
			m.cycles += 16
		}
	case 0x230f:
		{
			m.ip = 0x2312
			m.alu("sub", 16, m.r[1], uint16(60))
			m.cycles += 4
		}
	case 0x2312:
		{
			m.ip = 0x2314
			if m.cf {
				m.ip = uint16(8985)
			}
			m.cycles += 8
		}
	case 0x2314:
		{
			m.ip = 0x2316
			m.set8(3, 0, m.unary("inc", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x2316:
		{
			m.ip = 0x2319
			m.r[1] = m.alu("sub", 16, m.r[1], uint16(60))
			m.cycles += 4
		}
	case 0x2319:
		{
			m.ip = 0x231d
			m.wr16(m.r[11], uint16(0x1972), m.r[1])
			m.cycles += 8
		}
	case 0x231d:
		{
			m.ip = 0x231f
			m.set8(3, 0, m.alu("add", 8, ((m.r[3]>>0)&255), m.rd8(m.r[11], uint16(m.r[6]))))
			m.cycles += 16
		}
	case 0x231f:
		{
			m.ip = 0x2322
			m.alu("sub", 16, m.r[3], uint16(24))
			m.cycles += 4
		}
	case 0x2322:
		{
			m.ip = 0x2324
			if m.cf {
				m.ip = uint16(9002)
			}
			m.cycles += 8
		}
	case 0x2324:
		{
			m.ip = 0x2327
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(24))
			m.cycles += 4
		}
	case 0x2327:
		{
			m.ip = 0x2328
			m.r[0] = m.unary("inc", 16, m.r[0])
			m.cycles += 3
		}
	case 0x2328:
		{
			m.ip = 0x232a
			m.ip = uint16(8991)
			m.cycles += 15
		}
	case 0x232a:
		{
			m.ip = 0x232e
			m.wr16(m.r[11], uint16(0x1970), m.r[3])
			m.cycles += 8
		}
	case 0x232e:
		{
			m.ip = 0x2331
			m.alu("sub", 16, m.r[0], uint16(7))
			m.cycles += 4
		}
	case 0x2331:
		{
			m.ip = 0x2333
			if m.cf || m.zf {
				m.ip = uint16(9018)
			}
			m.cycles += 8
		}
	case 0x2333:
		{
			m.ip = 0x2336
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(7))
			m.cycles += 4
		}
	case 0x2336:
		{
			m.ip = 0x233a
			m.wr16(m.r[11], uint16(0x196c), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x196c))))
			m.cycles += 15
		}
	case 0x233a:
		{
			m.ip = 0x233d
			m.wr16(m.r[11], uint16(0x196e), m.r[0])
			m.cycles += 8
		}
	case 0x233d:
		{
			m.ip = 0x233e
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x233e:
		{
			m.ip = 0x233f
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x233f:
		{
			m.ip = 0x2343
			m.wr8(m.r[11], uint16(0x139e), ((m.r[1] >> 8) & 255))
			m.cycles += 8
		}
	case 0x2343:
		{
			m.ip = 0x2347
			m.wr8(m.r[11], uint16(0x139f), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x2347:
		{
			m.ip = 0x234b
			m.wr8(m.r[11], uint16(0x13a0), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x234b:
		{
			m.ip = 0x234c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x234c:
		{
			m.ip = 0x234e
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x234e:
		{
			m.ip = 0x2350
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x2350:
		{
			m.ip = 0x2352
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3])))
			m.cycles += 16
		}
	case 0x2352:
		{
			m.ip = 0x2354
			if m.sf != m.of {
				m.ip = uint16(9050)
			}
			m.cycles += 8
		}
	case 0x2354:
		{
			m.ip = 0x2356
			m.set8(0, 0, m.unary("inc", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x2356:
		{
			m.ip = 0x2358
			m.r[2] = m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3])))
			m.cycles += 16
		}
	case 0x2358:
		{
			m.ip = 0x235a
			m.ip = uint16(9040)
			m.cycles += 15
		}
	case 0x235a:
		{
			m.ip = 0x235c
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), uint16(48)))
			m.cycles += 4
		}
	case 0x235c:
		{
			m.ip = 0x235e
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(48))
			m.cycles += 4
		}
	case 0x235e:
		{
			m.ip = 0x2360
			if !m.zf {
				m.ip = uint16(9066)
			}
			m.cycles += 8
		}
	case 0x2360:
		{
			m.ip = 0x2363
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x2363:
		{
			m.ip = 0x2365
			if !m.zf {
				m.ip = uint16(9068)
			}
			m.cycles += 8
		}
	case 0x2365:
		{
			m.ip = 0x2367
			m.set8(0, 0, uint16(32))
			m.cycles += 2
		}
	case 0x2367:
		{
			m.ip = 0x2369
			m.ip = uint16(9068)
			m.cycles += 15
		}
	case 0x2369:
		{
			m.ip = 0x236a
			m.cycles += 3
		}
	case 0x236a:
		{
			m.ip = 0x236c
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x236c:
		{
			m.ip = 0x236e
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x236e:
		{
			m.ip = 0x236f
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x236f:
		{
			m.ip = 0x2372
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x2372:
		{
			m.ip = 0x2374
			m.set8(0, 0, uint16(10))
			m.cycles += 2
		}
	case 0x2374:
		{
			m.ip = 0x2377
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0xfffe)), ((m.r[0] >> 0) & 255))
			m.cycles += 16
		}
	case 0x2377:
		{
			m.ip = 0x2379
			if !m.zf {
				m.ip = uint16(9038)
			}
			m.cycles += 8
		}
	case 0x2379:
		{
			m.ip = 0x237c
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), uint16(48)))
			m.cycles += 4
		}
	case 0x237c:
		{
			m.ip = 0x237e
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[2] >> 0) & 255))
			m.cycles += 8
		}
	case 0x237e:
		{
			m.ip = 0x237f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x237f:
		{
			m.ip = 0x2380
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x2380:
		{
			m.ip = 0x2382
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x2382:
		{
			m.ip = 0x2383
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x2383:
		{
			m.ip = 0x2385
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(255))
			m.cycles += 4
		}
	case 0x2385:
		{
			m.ip = 0x2387
			if !m.zf {
				m.ip = uint16(9100)
			}
			m.cycles += 8
		}
	case 0x2387:
		{
			m.ip = 0x238b
			m.wr16(m.r[11], uint16(0x297), m.pop())
			m.cycles += 8
		}
	case 0x238b:
		{
			m.ip = 0x238c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x238c:
		{
			m.ip = 0x238e
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(13))
			m.cycles += 4
		}
	case 0x238e:
		{
			m.ip = 0x2390
			if !m.zf {
				m.ip = uint16(9112)
			}
			m.cycles += 8
		}
	case 0x2390:
		{
			m.ip = 0x2391
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x2391:
		{
			m.ip = 0x2395
			m.r[7] = m.alu("add", 16, m.r[7], uint16(160))
			m.cycles += 4
		}
	case 0x2395:
		{
			m.ip = 0x2396
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x2396:
		{
			m.ip = 0x2398
			m.ip = uint16(9088)
			m.cycles += 15
		}
	case 0x2398:
		{
			m.ip = 0x239b
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x239b:
		{
			m.ip = 0x239e
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x239e:
		{
			m.ip = 0x23a0
			m.ip = uint16(9088)
			m.cycles += 15
		}
	case 0x23a0:
		{
			m.ip = 0x23a2
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x23a2:
		{
			m.ip = 0x23a4
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(10))
			m.cycles += 4
		}
	case 0x23a4:
		{
			m.ip = 0x23a6
			if m.cf {
				m.ip = uint16(9132)
			}
			m.cycles += 8
		}
	case 0x23a6:
		{
			m.ip = 0x23a8
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), uint16(10)))
			m.cycles += 4
		}
	case 0x23a8:
		{
			m.ip = 0x23aa
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x23aa:
		{
			m.ip = 0x23ac
			m.ip = uint16(9122)
			m.cycles += 15
		}
	case 0x23ac:
		{
			m.ip = 0x23af
			m.r[0] = m.alu("add", 16, m.r[0], uint16(12336))
			m.cycles += 4
		}
	case 0x23af:
		{
			m.ip = 0x23b1
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x23b1:
		{
			m.ip = 0x23b4
			m.wr8(m.r[11], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x23b4:
		{
			m.ip = 0x23b7
			m.r[7] = m.alu("add", 16, m.r[7], uint16(3))
			m.cycles += 4
		}
	case 0x23b7:
		{
			m.ip = 0x23b8
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x23b8:
		{
			m.ip = 0x23ba
			m.set8(0, 8, uint16(42))
			m.cycles += 2
		}
	case 0x23ba:
		{
			m.ip = 0x23bc
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x23bc:
		{
			m.ip = 0x23c0
			m.r[1] = m.alu("sub", 16, m.r[1], uint16(1800))
			m.cycles += 4
		}
	case 0x23c0:
		{
			m.ip = 0x23c3
			m.r[1] = m.alu("sub", 16, m.r[1], uint16(100))
			m.cycles += 4
		}
	case 0x23c3:
		{
			m.ip = 0x23c6
			m.alu("sub", 16, m.r[1], uint16(100))
			m.cycles += 4
		}
	case 0x23c6:
		{
			m.ip = 0x23c8
			if !m.cf {
				m.ip = uint16(9152)
			}
			m.cycles += 8
		}
	case 0x23c8:
		{
			m.ip = 0x23c9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x23c9:
		{
			m.ip = 0x23cb
			m.set8(0, 8, uint16(44))
			m.cycles += 2
		}
	case 0x23cb:
		{
			m.ip = 0x23cd
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x23cd:
		{
			m.ip = 0x23ce
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x23ce:
		{
			m.ip = 0x23d2
			m.r[2] = uint16(0x383)
			m.cycles += 3
		}
	case 0x23d2:
		{
			m.ip = 0x23d4
			m.set8(0, 8, uint16(61))
			m.cycles += 2
		}
	case 0x23d4:
		{
			m.ip = 0x23d6
			m.set8(0, 0, uint16(2))
			m.cycles += 2
		}
	case 0x23d6:
		{
			m.ip = 0x23d8
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x23d8:
		{
			m.ip = 0x23da
			if !m.cf {
				m.ip = uint16(9282)
			}
			m.cycles += 8
		}
	case 0x23da:
		{
			m.ip = 0x23dd
			target := uint16(9144)
			m.push(0x23dd)
			m.ip = target
			m.cycles += 19
		}
	case 0x23dd:
		{
			m.ip = 0x23e1
			m.wr8(m.r[11], uint16(0x296), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x23e1:
		{
			m.ip = 0x23e5
			m.r[7] = uint16(0x199b)
			m.cycles += 3
		}
	case 0x23e5:
		{
			m.ip = 0x23e9
			m.r[7] = m.alu("add", 16, m.r[7], uint16(346))
			m.cycles += 4
		}
	case 0x23e9:
		{
			m.ip = 0x23ec
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x23ec:
		{
			m.ip = 0x23ef
			target := uint16(9265)
			m.push(0x23ef)
			m.ip = target
			m.cycles += 19
		}
	case 0x23ef:
		{
			m.ip = 0x23f2
			m.r[7] = m.alu("add", 16, m.r[7], uint16(64))
			m.cycles += 4
		}
	case 0x23f2:
		{
			m.ip = 0x23f4
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9196)
			}
			m.cycles += 17
		}
	case 0x23f4:
		{
			m.ip = 0x23f7
			m.r[7] = m.alu("add", 16, m.r[7], uint16(74))
			m.cycles += 4
		}
	case 0x23f7:
		{
			m.ip = 0x23fa
			target := uint16(9265)
			m.push(0x23fa)
			m.ip = target
			m.cycles += 19
		}
	case 0x23fa:
		{
			m.ip = 0x23fd
			m.r[7] = m.alu("add", 16, m.r[7], uint16(65))
			m.cycles += 4
		}
	case 0x23fd:
		{
			m.ip = 0x2400
			target := uint16(9265)
			m.push(0x2400)
			m.ip = target
			m.cycles += 19
		}
	case 0x2400:
		{
			m.ip = 0x2403
			m.r[7] = m.alu("add", 16, m.r[7], uint16(65))
			m.cycles += 4
		}
	case 0x2403:
		{
			m.ip = 0x2406
			target := uint16(9265)
			m.push(0x2406)
			m.ip = target
			m.cycles += 19
		}
	case 0x2406:
		{
			m.ip = 0x2409
			target := uint16(9161)
			m.push(0x2409)
			m.ip = target
			m.cycles += 19
		}
	case 0x2409:
		{
			m.ip = 0x240b
			m.r[2] = m.r[1]
			m.cycles += 2
		}
	case 0x240b:
		{
			m.ip = 0x240f
			m.r[7] = uint16(0x199b)
			m.cycles += 3
		}
	case 0x240f:
		{
			m.ip = 0x2413
			m.r[7] = m.alu("add", 16, m.r[7], uint16(357))
			m.cycles += 4
		}
	case 0x2413:
		{
			m.ip = 0x2416
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x2416:
		{
			m.ip = 0x2419
			target := uint16(9271)
			m.push(0x2419)
			m.ip = target
			m.cycles += 19
		}
	case 0x2419:
		{
			m.ip = 0x241c
			m.r[7] = m.alu("add", 16, m.r[7], uint16(67))
			m.cycles += 4
		}
	case 0x241c:
		{
			m.ip = 0x241e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9238)
			}
			m.cycles += 17
		}
	case 0x241e:
		{
			m.ip = 0x2421
			m.r[7] = m.alu("add", 16, m.r[7], uint16(74))
			m.cycles += 4
		}
	case 0x2421:
		{
			m.ip = 0x2424
			target := uint16(9271)
			m.push(0x2424)
			m.ip = target
			m.cycles += 19
		}
	case 0x2424:
		{
			m.ip = 0x2427
			m.r[7] = m.alu("add", 16, m.r[7], uint16(68))
			m.cycles += 4
		}
	case 0x2427:
		{
			m.ip = 0x242a
			target := uint16(9271)
			m.push(0x242a)
			m.ip = target
			m.cycles += 19
		}
	case 0x242a:
		{
			m.ip = 0x242d
			m.r[7] = m.alu("add", 16, m.r[7], uint16(68))
			m.cycles += 4
		}
	case 0x242d:
		{
			m.ip = 0x2430
			target := uint16(9271)
			m.push(0x2430)
			m.ip = target
			m.cycles += 19
		}
	case 0x2430:
		{
			m.ip = 0x2431
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2431:
		{
			m.ip = 0x2434
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x296)))
			m.cycles += 8
		}
	case 0x2434:
		{
			m.ip = 0x2437
			target := uint16(9120)
			m.push(0x2437)
			m.ip = target
			m.cycles += 19
		}
	case 0x2437:
		{
			m.ip = 0x2439
			m.set8(0, 0, ((m.r[2] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2439:
		{
			m.ip = 0x243c
			target := uint16(9120)
			m.push(0x243c)
			m.ip = target
			m.cycles += 19
		}
	case 0x243c:
		{
			m.ip = 0x243e
			m.set8(0, 0, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x243e:
		{
			m.ip = 0x2441
			target := uint16(9120)
			m.push(0x2441)
			m.ip = target
			m.cycles += 19
		}
	case 0x2441:
		{
			m.ip = 0x2442
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2442:
		{
			m.ip = 0x2445
			m.wr16(m.r[11], uint16(0x3aa), m.r[0])
			m.cycles += 8
		}
	case 0x2445:
		{
			m.ip = 0x2449
			m.r[3] = m.rd16(m.r[11], uint16(0x3aa))
			m.cycles += 8
		}
	case 0x2449:
		{
			m.ip = 0x244c
			m.r[1] = uint16(2221)
			m.cycles += 2
		}
	case 0x244c:
		{
			m.ip = 0x2450
			m.r[2] = uint16(0x165b)
			m.cycles += 3
		}
	case 0x2450:
		{
			m.ip = 0x2452
			m.set8(0, 8, uint16(63))
			m.cycles += 2
		}
	case 0x2452:
		{
			m.ip = 0x2454
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x2454:
		{
			m.ip = 0x2458
			m.r[3] = m.rd16(m.r[11], uint16(0x3aa))
			m.cycles += 8
		}
	case 0x2458:
		{
			m.ip = 0x245a
			m.set8(0, 8, uint16(62))
			m.cycles += 2
		}
	case 0x245a:
		{
			m.ip = 0x245c
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x245c:
		{
			m.ip = 0x245d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x245d:
		{
			m.ip = 0x2461
			m.r[2] = uint16(0x383)
			m.cycles += 3
		}
	case 0x2461:
		{
			m.ip = 0x2463
			m.set8(0, 8, uint16(60))
			m.cycles += 2
		}
	case 0x2463:
		{
			m.ip = 0x2465
			m.set8(1, 0, uint16(32))
			m.cycles += 2
		}
	case 0x2465:
		{
			m.ip = 0x2467
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x2467:
		{
			m.ip = 0x246a
			m.wr16(m.r[11], uint16(0x3aa), m.r[0])
			m.cycles += 8
		}
	case 0x246a:
		{
			m.ip = 0x246e
			m.r[3] = m.rd16(m.r[11], uint16(0x3aa))
			m.cycles += 8
		}
	case 0x246e:
		{
			m.ip = 0x2471
			m.r[1] = uint16(2221)
			m.cycles += 2
		}
	case 0x2471:
		{
			m.ip = 0x2475
			m.r[2] = uint16(0x165b)
			m.cycles += 3
		}
	case 0x2475:
		{
			m.ip = 0x2477
			m.set8(0, 8, uint16(64))
			m.cycles += 2
		}
	case 0x2477:
		{
			m.ip = 0x2479
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x2479:
		{
			m.ip = 0x247d
			m.r[3] = m.rd16(m.r[11], uint16(0x3aa))
			m.cycles += 8
		}
	case 0x247d:
		{
			m.ip = 0x247f
			m.set8(0, 8, uint16(62))
			m.cycles += 2
		}
	case 0x247f:
		{
			m.ip = 0x2481
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x2481:
		{
			m.ip = 0x2482
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2482:
		{
			m.ip = 0x2485
			m.alu("sub", 16, m.r[3], uint16(65535))
			m.cycles += 4
		}
	case 0x2485:
		{
			m.ip = 0x2487
			if m.zf {
				m.ip = uint16(9357)
			}
			m.cycles += 8
		}
	case 0x2487:
		{
			m.ip = 0x248a
			m.r[1] = uint16(1)
			m.cycles += 2
		}
	case 0x248a:
		{
			m.ip = 0x248c
			m.ip = uint16(9363)
			m.cycles += 15
		}
	case 0x248c:
		{
			m.ip = 0x248d
			m.cycles += 3
		}
	case 0x248d:
		{
			m.ip = 0x2490
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x2490:
		{
			m.ip = 0x2493
			m.r[1] = uint16(4)
			m.cycles += 2
		}
	case 0x2493:
		{
			m.ip = 0x2495
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x2495:
		{
			m.ip = 0x2499
			m.wr16(m.r[11], uint16(m.r[3]+0x205), m.r[0])
			m.cycles += 8
		}
	case 0x2499:
		{
			m.ip = 0x249c
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]+0x2))
			m.cycles += 8
		}
	case 0x249c:
		{
			m.ip = 0x24a0
			m.wr16(m.r[11], uint16(m.r[3]+0x20d), m.r[0])
			m.cycles += 8
		}
	case 0x24a0:
		{
			m.ip = 0x24a3
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]+0x4))
			m.cycles += 8
		}
	case 0x24a3:
		{
			m.ip = 0x24a7
			m.wr16(m.r[11], uint16(m.r[3]+0x215), m.r[0])
			m.cycles += 8
		}
	case 0x24a7:
		{
			m.ip = 0x24aa
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x24aa:
		{
			m.ip = 0x24ac
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9363)
			}
			m.cycles += 17
		}
	case 0x24ac:
		{
			m.ip = 0x24af
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x24af:
		{
			m.ip = 0x24b0
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x24b0:
		{
			m.ip = 0x24b1
			m.push(m.r[8])
			m.cycles += 11
		}
	case 0x24b1:
		{
			m.ip = 0x24b2
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x24b2:
		{
			m.ip = 0x24b3
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x24b3:
		{
			m.ip = 0x24b4
			m.df = false
			m.cycles += 2
		}
	case 0x24b4:
		{
			m.ip = 0x24b6
			m.stringOp("movs", 8)
			m.cycles += 2
		}
	case 0x24b6:
		{
			m.ip = 0x24b7
			m.r[8] = m.pop()
			m.cycles += 8
		}
	case 0x24b7:
		{
			m.ip = 0x24b8
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x24b8:
		{
			m.ip = 0x24bc
			m.r[3] = m.rd16(m.r[11], uint16(m.r[3]+0x48d))
			m.cycles += 8
		}
	case 0x24bc:
		{
			m.ip = 0x24c0
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x24c0:
		{
			m.ip = 0x24c3
			m.r[7] = uint16(4642)
			m.cycles += 2
		}
	case 0x24c3:
		{
			m.ip = 0x24c6
			m.r[3] = m.alu("add", 16, m.r[3], uint16(22))
			m.cycles += 4
		}
	case 0x24c6:
		{
			m.ip = 0x24c9
			m.r[1] = uint16(11)
			m.cycles += 2
		}
	case 0x24c9:
		{
			m.ip = 0x24ca
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x24ca:
		{
			m.ip = 0x24cb
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x24cb:
		{
			m.ip = 0x24ce
			m.r[1] = uint16(19)
			m.cycles += 2
		}
	case 0x24ce:
		{
			m.ip = 0x24d1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3])), uint16(32))
			m.cycles += 16
		}
	case 0x24d1:
		{
			m.ip = 0x24d3
			if !m.zf {
				m.ip = uint16(9454)
			}
			m.cycles += 8
		}
	case 0x24d3:
		{
			m.ip = 0x24d4
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x24d4:
		{
			m.ip = 0x24d7
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x24d7:
		{
			m.ip = 0x24da
			target := uint16(10677)
			m.push(0x24da)
			m.ip = target
			m.cycles += 19
		}
	case 0x24da:
		{
			m.ip = 0x24de
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x24de:
		{
			m.ip = 0x24df
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x24df:
		{
			m.ip = 0x24e1
			m.ip = uint16(9454)
			m.cycles += 15
		}
	case 0x24e1:
		{
			m.ip = 0x24e2
			m.cycles += 3
		}
	case 0x24e2:
		{
			m.ip = 0x24e3
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x24e3:
		{
			m.ip = 0x24e6
			m.r[6] = uint16(28072)
			m.cycles += 2
		}
	case 0x24e6:
		{
			m.ip = 0x24e9
			target := uint16(10677)
			m.push(0x24e9)
			m.ip = target
			m.cycles += 19
		}
	case 0x24e9:
		{
			m.ip = 0x24ed
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x24ed:
		{
			m.ip = 0x24ee
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x24ee:
		{
			m.ip = 0x24f1
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x24f1:
		{
			m.ip = 0x24f2
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x24f2:
		{
			m.ip = 0x24f4
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9422)
			}
			m.cycles += 17
		}
	case 0x24f4:
		{
			m.ip = 0x24f7
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x24f7:
		{
			m.ip = 0x24f8
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x24f8:
		{
			m.ip = 0x24fc
			m.r[7] = m.alu("add", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x24fc:
		{
			m.ip = 0x24fd
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x24fd:
		{
			m.ip = 0x24ff
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9417)
			}
			m.cycles += 17
		}
	case 0x24ff:
		{
			m.ip = 0x2503
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x2503:
		{
			m.ip = 0x2507
			m.r[6] = uint16(0x1045)
			m.cycles += 3
		}
	case 0x2507:
		{
			m.ip = 0x250a
			target := uint16(18306)
			m.push(0x250a)
			m.ip = target
			m.cycles += 19
		}
	case 0x250a:
		{
			m.ip = 0x250e
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x250e:
		{
			m.ip = 0x250f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x250f:
		{
			m.ip = 0x2512
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x2512:
		{
			m.ip = 0x2515
			m.r[7] = uint16(4642)
			m.cycles += 2
		}
	case 0x2515:
		{
			m.ip = 0x2518
			m.r[3] = m.alu("add", 16, m.r[3], uint16(22))
			m.cycles += 4
		}
	case 0x2518:
		{
			m.ip = 0x251b
			m.r[1] = uint16(11)
			m.cycles += 2
		}
	case 0x251b:
		{
			m.ip = 0x251c
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x251c:
		{
			m.ip = 0x251d
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x251d:
		{
			m.ip = 0x2520
			m.r[1] = uint16(19)
			m.cycles += 2
		}
	case 0x2520:
		{
			m.ip = 0x2523
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x2523:
		{
			m.ip = 0x2528
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x2528:
		{
			m.ip = 0x252a
			if !m.zf {
				m.ip = uint16(9524)
			}
			m.cycles += 8
		}
	case 0x252a:
		{
			m.ip = 0x252f
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(1))
			m.cycles += 16
		}
	case 0x252f:
		{
			m.ip = 0x2531
			if m.zf {
				m.ip = uint16(9533)
			}
			m.cycles += 8
		}
	case 0x2531:
		{
			m.ip = 0x2534
			m.r[6] = uint16(28064)
			m.cycles += 2
		}
	case 0x2534:
		{
			m.ip = 0x2535
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x2535:
		{
			m.ip = 0x2538
			target := uint16(10677)
			m.push(0x2538)
			m.ip = target
			m.cycles += 19
		}
	case 0x2538:
		{
			m.ip = 0x253c
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x253c:
		{
			m.ip = 0x253d
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x253d:
		{
			m.ip = 0x2540
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x2540:
		{
			m.ip = 0x2541
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x2541:
		{
			m.ip = 0x2543
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9504)
			}
			m.cycles += 17
		}
	case 0x2543:
		{
			m.ip = 0x2546
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x2546:
		{
			m.ip = 0x2547
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x2547:
		{
			m.ip = 0x254b
			m.r[7] = m.alu("add", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x254b:
		{
			m.ip = 0x254c
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x254c:
		{
			m.ip = 0x254e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9499)
			}
			m.cycles += 17
		}
	case 0x254e:
		{
			m.ip = 0x2553
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x37)), uint16(0))
			m.cycles += 16
		}
	case 0x2553:
		{
			m.ip = 0x2555
			if m.zf {
				m.ip = uint16(9568)
			}
			m.cycles += 8
		}
	case 0x2555:
		{
			m.ip = 0x2559
			m.r[6] = m.rd16(m.r[11], uint16(0x12))
			m.cycles += 8
		}
	case 0x2559:
		{
			m.ip = 0x255d
			m.r[7] = m.rd16(m.r[11], uint16(0x3a))
			m.cycles += 8
		}
	case 0x255d:
		{
			m.ip = 0x2560
			target := uint16(10677)
			m.push(0x2560)
			m.ip = target
			m.cycles += 19
		}
	case 0x2560:
		{
			m.ip = 0x2565
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x40)), uint16(1))
			m.cycles += 16
		}
	case 0x2565:
		{
			m.ip = 0x2567
			if !m.zf {
				m.ip = uint16(9586)
			}
			m.cycles += 8
		}
	case 0x2567:
		{
			m.ip = 0x256b
			m.r[6] = m.rd16(m.r[11], uint16(0x16))
			m.cycles += 8
		}
	case 0x256b:
		{
			m.ip = 0x256f
			m.r[7] = m.rd16(m.r[11], uint16(0x43))
			m.cycles += 8
		}
	case 0x256f:
		{
			m.ip = 0x2572
			target := uint16(10677)
			m.push(0x2572)
			m.ip = target
			m.cycles += 19
		}
	case 0x2572:
		{
			m.ip = 0x2577
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x466)), uint16(0))
			m.cycles += 16
		}
	case 0x2577:
		{
			m.ip = 0x2579
			if !m.zf {
				m.ip = uint16(9596)
			}
			m.cycles += 8
		}
	case 0x2579:
		{
			m.ip = 0x257c
			target := uint16(12901)
			m.push(0x257c)
			m.ip = target
			m.cycles += 19
		}
	case 0x257c:
		{
			m.ip = 0x2581
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x2581:
		{
			m.ip = 0x2583
			if !m.zf {
				m.ip = uint16(9640)
			}
			m.cycles += 8
		}
	case 0x2583:
		{
			m.ip = 0x2588
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(2))
			m.cycles += 16
		}
	case 0x2588:
		{
			m.ip = 0x258a
			if !m.zf {
				m.ip = uint16(9640)
			}
			m.cycles += 8
		}
	case 0x258a:
		{
			m.ip = 0x258d
			m.r[0] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x258d:
		{
			m.ip = 0x258e
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x258e:
		{
			m.ip = 0x2591
			m.r[0] = m.alu("add", 16, m.r[0], uint16(2))
			m.cycles += 4
		}
	case 0x2591:
		{
			m.ip = 0x2594
			m.r[0] = m.alu("and", 16, m.r[0], uint16(2))
			m.cycles += 4
		}
	case 0x2594:
		{
			m.ip = 0x2597
			m.wr16(m.r[11], uint16(0x225), m.r[0])
			m.cycles += 8
		}
	case 0x2597:
		{
			m.ip = 0x2599
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x2599:
		{
			m.ip = 0x259c
			m.wr16(m.r[11], uint16(0x223), m.r[0])
			m.cycles += 8
		}
	case 0x259c:
		{
			m.ip = 0x259f
			target := uint16(2965)
			m.push(0x259f)
			m.ip = target
			m.cycles += 19
		}
	case 0x259f:
		{
			m.ip = 0x25a0
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x25a0:
		{
			m.ip = 0x25a3
			m.wr16(m.r[11], uint16(0x225), m.r[0])
			m.cycles += 8
		}
	case 0x25a3:
		{
			m.ip = 0x25a5
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x25a5:
		{
			m.ip = 0x25a8
			m.wr16(m.r[11], uint16(0x223), m.r[0])
			m.cycles += 8
		}
	case 0x25a8:
		{
			m.ip = 0x25ac
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x25ac:
		{
			m.ip = 0x25b0
			m.r[6] = uint16(0x1045)
			m.cycles += 3
		}
	case 0x25b0:
		{
			m.ip = 0x25b3
			target := uint16(18306)
			m.push(0x25b3)
			m.ip = target
			m.cycles += 19
		}
	case 0x25b3:
		{
			m.ip = 0x25b7
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x25b7:
		{
			m.ip = 0x25b8
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x25b8:
		{
			m.ip = 0x25bc
			m.r[3] = m.rd16(m.r[11], uint16(m.r[3]+0x48d))
			m.cycles += 8
		}
	case 0x25bc:
		{
			m.ip = 0x25c2
			m.wr16(m.r[11], uint16(0x78), uint16(209))
			m.cycles += 8
		}
	case 0x25c2:
		{
			m.ip = 0x25c6
			m.r[7] = uint16(0xa4)
			m.cycles += 3
		}
	case 0x25c6:
		{
			m.ip = 0x25c9
			m.r[3] = m.alu("add", 16, m.r[3], uint16(22))
			m.cycles += 4
		}
	case 0x25c9:
		{
			m.ip = 0x25cc
			m.r[1] = uint16(11)
			m.cycles += 2
		}
	case 0x25cc:
		{
			m.ip = 0x25cd
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x25cd:
		{
			m.ip = 0x25d0
			m.r[1] = uint16(19)
			m.cycles += 2
		}
	case 0x25d0:
		{
			m.ip = 0x25d2
			m.set8(0, 0, uint16(249))
			m.cycles += 2
		}
	case 0x25d2:
		{
			m.ip = 0x25d5
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3])), uint16(32))
			m.cycles += 16
		}
	case 0x25d5:
		{
			m.ip = 0x25d7
			if m.zf {
				m.ip = uint16(9693)
			}
			m.cycles += 8
		}
	case 0x25d7:
		{
			m.ip = 0x25d9
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(246)))
			m.cycles += 4
		}
	case 0x25d9:
		{
			m.ip = 0x25dd
			m.wr16(m.r[11], uint16(0x78), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x78))))
			m.cycles += 15
		}
	case 0x25dd:
		{
			m.ip = 0x25e1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0xffeb)), uint16(32))
			m.cycles += 16
		}
	case 0x25e1:
		{
			m.ip = 0x25e3
			if m.zf {
				m.ip = uint16(9701)
			}
			m.cycles += 8
		}
	case 0x25e3:
		{
			m.ip = 0x25e5
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(127)))
			m.cycles += 4
		}
	case 0x25e5:
		{
			m.ip = 0x25e9
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x15)), uint16(32))
			m.cycles += 16
		}
	case 0x25e9:
		{
			m.ip = 0x25eb
			if m.zf {
				m.ip = uint16(9709)
			}
			m.cycles += 8
		}
	case 0x25eb:
		{
			m.ip = 0x25ed
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(191)))
			m.cycles += 4
		}
	case 0x25ed:
		{
			m.ip = 0x25f1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0xffff)), uint16(32))
			m.cycles += 16
		}
	case 0x25f1:
		{
			m.ip = 0x25f3
			if m.zf {
				m.ip = uint16(9717)
			}
			m.cycles += 8
		}
	case 0x25f3:
		{
			m.ip = 0x25f5
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(223)))
			m.cycles += 4
		}
	case 0x25f5:
		{
			m.ip = 0x25f9
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1)), uint16(32))
			m.cycles += 16
		}
	case 0x25f9:
		{
			m.ip = 0x25fb
			if m.zf {
				m.ip = uint16(9725)
			}
			m.cycles += 8
		}
	case 0x25fb:
		{
			m.ip = 0x25fd
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(239)))
			m.cycles += 4
		}
	case 0x25fd:
		{
			m.ip = 0x25ff
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x25ff:
		{
			m.ip = 0x2600
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x2600:
		{
			m.ip = 0x2601
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x2601:
		{
			m.ip = 0x2603
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9680)
			}
			m.cycles += 17
		}
	case 0x2603:
		{
			m.ip = 0x2606
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x2606:
		{
			m.ip = 0x2609
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x2609:
		{
			m.ip = 0x260a
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x260a:
		{
			m.ip = 0x260c
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9676)
			}
			m.cycles += 17
		}
	case 0x260c:
		{
			m.ip = 0x260f
			m.r[7] = uint16(115)
			m.cycles += 2
		}
	case 0x260f:
		{
			m.ip = 0x2614
			m.wr8(m.r[11], uint16(m.r[7]+0x8e), m.alu("or", 8, m.rd8(m.r[11], uint16(m.r[7]+0x8e)), uint16(64)))
			m.cycles += 16
		}
	case 0x2614:
		{
			m.ip = 0x2615
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2615:
		{
			m.ip = 0x2619
			m.r[6] = m.rd16(m.r[11], uint16(0x55d))
			m.cycles += 8
		}
	case 0x2619:
		{
			m.ip = 0x261d
			m.alu("sub", 16, m.r[6], uint16(1375))
			m.cycles += 4
		}
	case 0x261d:
		{
			m.ip = 0x261f
			if m.zf {
				m.ip = uint16(9780)
			}
			m.cycles += 8
		}
	case 0x261f:
		{
			m.ip = 0x2622
			m.r[0] = m.rd16(m.r[11], uint16(0x55b))
			m.cycles += 8
		}
	case 0x2622:
		{
			m.ip = 0x2624
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(m.r[6])))
			m.cycles += 16
		}
	case 0x2624:
		{
			m.ip = 0x2626
			if !m.zf {
				m.ip = uint16(9843)
			}
			m.cycles += 8
		}
	case 0x2626:
		{
			m.ip = 0x262c
			m.wr16(m.r[11], uint16(0x55b), uint16(0))
			m.cycles += 8
		}
	case 0x262c:
		{
			m.ip = 0x2630
			m.r[6] = uint16(0x55f)
			m.cycles += 3
		}
	case 0x2630:
		{
			m.ip = 0x2634
			m.wr16(m.r[11], uint16(0x55d), m.r[6])
			m.cycles += 8
		}
	case 0x2634:
		{
			m.ip = 0x2637
			m.r[0] = uint16(0)
			m.cycles += 2
		}
	case 0x2637:
		{
			m.ip = 0x263b
			m.r[5] = m.rd16(m.r[11], uint16(0x1a5))
			m.cycles += 8
		}
	case 0x263b:
		{
			m.ip = 0x2641
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x2641:
		{
			m.ip = 0x2643
			if m.zf {
				m.ip = uint16(9804)
			}
			m.cycles += 8
		}
	case 0x2643:
		{
			m.ip = 0x2646
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x1a3)))
			m.cycles += 8
		}
	case 0x2646:
		{
			m.ip = 0x2648
			m.set8(0, 0, m.shift("shl", 8, ((m.r[0]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x2648:
		{
			m.ip = 0x264a
			m.r[6] = m.alu("add", 16, m.r[6], m.r[0])
			m.cycles += 4
		}
	case 0x264a:
		{
			m.ip = 0x264c
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x264c:
		{
			m.ip = 0x2651
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x2651:
		{
			m.ip = 0x2653
			if !m.zf {
				m.ip = uint16(9855)
			}
			m.cycles += 8
		}
	case 0x2653:
		{
			m.ip = 0x2657
			m.r[6] = uint16(0x55f)
			m.cycles += 3
		}
	case 0x2657:
		{
			m.ip = 0x265a
			m.r[3] = uint16(1)
			m.cycles += 2
		}
	case 0x265a:
		{
			m.ip = 0x265e
			m.r[5] = m.rd16(m.r[11], uint16(0x1a7))
			m.cycles += 8
		}
	case 0x265e:
		{
			m.ip = 0x2664
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x2664:
		{
			m.ip = 0x2666
			if m.zf {
				m.ip = uint16(9855)
			}
			m.cycles += 8
		}
	case 0x2666:
		{
			m.ip = 0x266a
			m.set8(3, 0, m.rd8(m.r[11], uint16(0x1a4)))
			m.cycles += 8
		}
	case 0x266a:
		{
			m.ip = 0x266c
			m.set8(3, 0, m.shift("shl", 8, ((m.r[3]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x266c:
		{
			m.ip = 0x266e
			m.r[6] = m.alu("add", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x266e:
		{
			m.ip = 0x2670
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(m.r[6])))
			m.cycles += 16
		}
	case 0x2670:
		{
			m.ip = 0x2672
			m.ip = uint16(9855)
			m.cycles += 15
		}
	case 0x2672:
		{
			m.ip = 0x2673
			m.cycles += 3
		}
	case 0x2673:
		{
			m.ip = 0x2676
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x2676:
		{
			m.ip = 0x2678
			m.r[6] = m.alu("add", 16, m.r[6], m.r[0])
			m.cycles += 4
		}
	case 0x2678:
		{
			m.ip = 0x267d
			m.wr16(m.r[11], uint16(0x55b), m.alu("add", 16, m.rd16(m.r[11], uint16(0x55b)), uint16(2)))
			m.cycles += 16
		}
	case 0x267d:
		{
			m.ip = 0x267f
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x267f:
		{
			m.ip = 0x2682
			target := uint16(9905)
			m.push(0x2682)
			m.ip = target
			m.cycles += 19
		}
	case 0x2682:
		{
			m.ip = 0x2687
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(7))
			m.cycles += 16
		}
	case 0x2687:
		{
			m.ip = 0x2689
			if !m.zf {
				m.ip = uint16(9873)
			}
			m.cycles += 8
		}
	case 0x2689:
		{
			m.ip = 0x268e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f6)), uint16(14))
			m.cycles += 16
		}
	case 0x268e:
		{
			m.ip = 0x2690
			if !m.zf {
				m.ip = uint16(9873)
			}
			m.cycles += 8
		}
	case 0x2690:
		{
			m.ip = 0x2691
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2691:
		{
			m.ip = 0x2692
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x2692:
		{
			m.ip = 0x2693
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x2693:
		{
			m.ip = 0x2697
			m.r[3] = m.rd16(m.r[11], uint16(0x106c))
			m.cycles += 8
		}
	case 0x2697:
		{
			m.ip = 0x269a
			m.r[1] = uint16(8)
			m.cycles += 2
		}
	case 0x269a:
		{
			m.ip = 0x269f
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x27d)), uint16(0))
			m.cycles += 16
		}
	case 0x269f:
		{
			m.ip = 0x26a1
			if m.zf {
				m.ip = uint16(9892)
			}
			m.cycles += 8
		}
	case 0x26a1:
		{
			m.ip = 0x26a4
			m.r[1] = uint16(5)
			m.cycles += 2
		}
	case 0x26a4:
		{
			m.ip = 0x26a7
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x26a7:
		{
			m.ip = 0x26a9
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9892)
			}
			m.cycles += 17
		}
	case 0x26a9:
		{
			m.ip = 0x26ab
			m.r[1] = m.r[3]
			m.cycles += 2
		}
	case 0x26ab:
		{
			m.ip = 0x26ac
			m.r[3] = m.unary("dec", 16, m.r[3])
			m.cycles += 3
		}
	case 0x26ac:
		{
			m.ip = 0x26ae
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(9879)
			}
			m.cycles += 17
		}
	case 0x26ae:
		{
			m.ip = 0x26af
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x26af:
		{
			m.ip = 0x26b0
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x26b0:
		{
			m.ip = 0x26b1
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x26b1:
		{
			m.ip = 0x26b6
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x55a)), uint16(0))
			m.cycles += 16
		}
	case 0x26b6:
		{
			m.ip = 0x26b8
			if m.zf {
				m.ip = uint16(9938)
			}
			m.cycles += 8
		}
	case 0x26b8:
		{
			m.ip = 0x26bb
			m.alu("sub", 16, m.r[0], uint16(0))
			m.cycles += 4
		}
	case 0x26bb:
		{
			m.ip = 0x26bd
			if m.zf {
				m.ip = uint16(9938)
			}
			m.cycles += 8
		}
	case 0x26bd:
		{
			m.ip = 0x26be
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x26be:
		{
			m.ip = 0x26c0
			m.set8(0, 0, uint16(182))
			m.cycles += 2
		}
	case 0x26c0:
		{
			m.ip = 0x26c2
			m.output(uint16(67), ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x26c2:
		{
			m.ip = 0x26c3
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x26c3:
		{
			m.ip = 0x26c5
			m.output(uint16(66), ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x26c5:
		{
			m.ip = 0x26c7
			m.ip = uint16(9927)
			m.cycles += 15
		}
	case 0x26c7:
		{
			m.ip = 0x26c9
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x26c9:
		{
			m.ip = 0x26cb
			m.output(uint16(66), ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x26cb:
		{
			m.ip = 0x26cd
			m.set8(0, 0, m.input(uint16(97)))
			m.cycles += 8
		}
	case 0x26cd:
		{
			m.ip = 0x26cf
			m.set8(0, 0, m.alu("or", 8, ((m.r[0]>>0)&255), uint16(3)))
			m.cycles += 4
		}
	case 0x26cf:
		{
			m.ip = 0x26d1
			m.output(uint16(97), ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x26d1:
		{
			m.ip = 0x26d2
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x26d2:
		{
			m.ip = 0x26d4
			m.set8(0, 0, m.input(uint16(97)))
			m.cycles += 8
		}
	case 0x26d4:
		{
			m.ip = 0x26d6
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(252)))
			m.cycles += 4
		}
	case 0x26d6:
		{
			m.ip = 0x26d8
			m.output(uint16(97), ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x26d8:
		{
			m.ip = 0x26d9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x26d9:
		{
			m.ip = 0x26dc
			target := uint16(18515)
			m.push(0x26dc)
			m.ip = target
			m.cycles += 19
		}
	case 0x26dc:
		{
			m.ip = 0x26de
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x26de:
		{
			m.ip = 0x26e0
			if !m.zf {
				m.ip = uint16(10018)
			}
			m.cycles += 8
		}
	case 0x26e0:
		{
			m.ip = 0x26e3
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(72))
			m.cycles += 4
		}
	case 0x26e3:
		{
			m.ip = 0x26e5
			if !m.zf {
				m.ip = uint16(9963)
			}
			m.cycles += 8
		}
	case 0x26e5:
		{
			m.ip = 0x26ea
			m.wr8(m.r[11], uint16(0x1a1), uint16(24))
			m.cycles += 8
		}
	case 0x26ea:
		{
			m.ip = 0x26eb
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x26eb:
		{
			m.ip = 0x26ee
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(80))
			m.cycles += 4
		}
	case 0x26ee:
		{
			m.ip = 0x26f0
			if !m.zf {
				m.ip = uint16(9974)
			}
			m.cycles += 8
		}
	case 0x26f0:
		{
			m.ip = 0x26f5
			m.wr8(m.r[11], uint16(0x1a1), uint16(16))
			m.cycles += 8
		}
	case 0x26f5:
		{
			m.ip = 0x26f6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x26f6:
		{
			m.ip = 0x26f9
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(75))
			m.cycles += 4
		}
	case 0x26f9:
		{
			m.ip = 0x26fb
			if !m.zf {
				m.ip = uint16(9985)
			}
			m.cycles += 8
		}
	case 0x26fb:
		{
			m.ip = 0x2700
			m.wr8(m.r[11], uint16(0x1a1), uint16(8))
			m.cycles += 8
		}
	case 0x2700:
		{
			m.ip = 0x2701
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2701:
		{
			m.ip = 0x2704
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(77))
			m.cycles += 4
		}
	case 0x2704:
		{
			m.ip = 0x2706
			if !m.zf {
				m.ip = uint16(10033)
			}
			m.cycles += 8
		}
	case 0x2706:
		{
			m.ip = 0x270b
			m.wr8(m.r[11], uint16(0x1a1), uint16(0))
			m.cycles += 8
		}
	case 0x270b:
		{
			m.ip = 0x270c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x270c:
		{
			m.ip = 0x270f
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(71))
			m.cycles += 4
		}
	case 0x270f:
		{
			m.ip = 0x2711
			if !m.zf {
				m.ip = uint16(10007)
			}
			m.cycles += 8
		}
	case 0x2711:
		{
			m.ip = 0x2716
			m.wr8(m.r[11], uint16(0x1a1), uint16(255))
			m.cycles += 8
		}
	case 0x2716:
		{
			m.ip = 0x2717
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2717:
		{
			m.ip = 0x271a
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(59))
			m.cycles += 4
		}
	case 0x271a:
		{
			m.ip = 0x271c
			if !m.zf {
				m.ip = uint16(10033)
			}
			m.cycles += 8
		}
	case 0x271c:
		{
			m.ip = 0x2721
			m.wr8(m.r[11], uint16(0x27a), m.alu("xor", 8, m.rd8(m.r[11], uint16(0x27a)), uint16(1)))
			m.cycles += 16
		}
	case 0x2721:
		{
			m.ip = 0x2722
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2722:
		{
			m.ip = 0x2724
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x2724:
		{
			m.ip = 0x2726
			if !m.zf {
				m.ip = uint16(10025)
			}
			m.cycles += 8
		}
	case 0x2726:
		{
			m.ip = 0x2729
			target := uint16(3091)
			m.push(0x2729)
			m.ip = target
			m.cycles += 19
		}
	case 0x2729:
		{
			m.ip = 0x272c
			m.alu("sub", 16, m.r[0], uint16(4113))
			m.cycles += 4
		}
	case 0x272c:
		{
			m.ip = 0x272e
			if !m.zf {
				m.ip = uint16(10033)
			}
			m.cycles += 8
		}
	case 0x272e:
		{
			m.ip = 0x2731
			m.ip = uint16(17160)
			m.cycles += 15
		}
	case 0x2731:
		{
			m.ip = 0x2732
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2732:
		{
			m.ip = 0x2735
			target := uint16(18710)
			m.push(0x2735)
			m.ip = target
			m.cycles += 19
		}
	case 0x2735:
		{
			m.ip = 0x2737
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x2737:
		{
			m.ip = 0x2739
			if !m.zf {
				m.ip = uint16(10107)
			}
			m.cycles += 8
		}
	case 0x2739:
		{
			m.ip = 0x273c
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(72))
			m.cycles += 4
		}
	case 0x273c:
		{
			m.ip = 0x273e
			if !m.zf {
				m.ip = uint16(10052)
			}
			m.cycles += 8
		}
	case 0x273e:
		{
			m.ip = 0x2743
			m.wr8(m.r[11], uint16(0x1a2), uint16(24))
			m.cycles += 8
		}
	case 0x2743:
		{
			m.ip = 0x2744
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2744:
		{
			m.ip = 0x2747
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(80))
			m.cycles += 4
		}
	case 0x2747:
		{
			m.ip = 0x2749
			if !m.zf {
				m.ip = uint16(10063)
			}
			m.cycles += 8
		}
	case 0x2749:
		{
			m.ip = 0x274e
			m.wr8(m.r[11], uint16(0x1a2), uint16(16))
			m.cycles += 8
		}
	case 0x274e:
		{
			m.ip = 0x274f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x274f:
		{
			m.ip = 0x2752
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(75))
			m.cycles += 4
		}
	case 0x2752:
		{
			m.ip = 0x2754
			if !m.zf {
				m.ip = uint16(10074)
			}
			m.cycles += 8
		}
	case 0x2754:
		{
			m.ip = 0x2759
			m.wr8(m.r[11], uint16(0x1a2), uint16(8))
			m.cycles += 8
		}
	case 0x2759:
		{
			m.ip = 0x275a
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x275a:
		{
			m.ip = 0x275d
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(77))
			m.cycles += 4
		}
	case 0x275d:
		{
			m.ip = 0x275f
			if !m.zf {
				m.ip = uint16(10172)
			}
			m.cycles += 8
		}
	case 0x275f:
		{
			m.ip = 0x2764
			m.wr8(m.r[11], uint16(0x1a2), uint16(0))
			m.cycles += 8
		}
	case 0x2764:
		{
			m.ip = 0x2765
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2765:
		{
			m.ip = 0x2768
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(71))
			m.cycles += 4
		}
	case 0x2768:
		{
			m.ip = 0x276a
			if !m.zf {
				m.ip = uint16(10096)
			}
			m.cycles += 8
		}
	case 0x276a:
		{
			m.ip = 0x276f
			m.wr8(m.r[11], uint16(0x1a2), uint16(255))
			m.cycles += 8
		}
	case 0x276f:
		{
			m.ip = 0x2770
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2770:
		{
			m.ip = 0x2773
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(59))
			m.cycles += 4
		}
	case 0x2773:
		{
			m.ip = 0x2775
			if !m.zf {
				m.ip = uint16(10172)
			}
			m.cycles += 8
		}
	case 0x2775:
		{
			m.ip = 0x277a
			m.wr8(m.r[11], uint16(0x27a), m.alu("xor", 8, m.rd8(m.r[11], uint16(0x27a)), uint16(1)))
			m.cycles += 16
		}
	case 0x277a:
		{
			m.ip = 0x277b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x277b:
		{
			m.ip = 0x277d
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(97))
			m.cycles += 4
		}
	case 0x277d:
		{
			m.ip = 0x277f
			if m.cf {
				m.ip = uint16(10113)
			}
			m.cycles += 8
		}
	case 0x277f:
		{
			m.ip = 0x2781
			m.set8(0, 0, m.alu("sub", 8, ((m.r[0]>>0)&255), uint16(32)))
			m.cycles += 4
		}
	case 0x2781:
		{
			m.ip = 0x2783
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(87))
			m.cycles += 4
		}
	case 0x2783:
		{
			m.ip = 0x2785
			if !m.zf {
				m.ip = uint16(10123)
			}
			m.cycles += 8
		}
	case 0x2785:
		{
			m.ip = 0x278a
			m.wr8(m.r[11], uint16(0x1a1), uint16(24))
			m.cycles += 8
		}
	case 0x278a:
		{
			m.ip = 0x278b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x278b:
		{
			m.ip = 0x278d
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(88))
			m.cycles += 4
		}
	case 0x278d:
		{
			m.ip = 0x278f
			if m.zf {
				m.ip = uint16(10131)
			}
			m.cycles += 8
		}
	case 0x278f:
		{
			m.ip = 0x2791
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(90))
			m.cycles += 4
		}
	case 0x2791:
		{
			m.ip = 0x2793
			if !m.zf {
				m.ip = uint16(10137)
			}
			m.cycles += 8
		}
	case 0x2793:
		{
			m.ip = 0x2798
			m.wr8(m.r[11], uint16(0x1a1), uint16(16))
			m.cycles += 8
		}
	case 0x2798:
		{
			m.ip = 0x2799
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2799:
		{
			m.ip = 0x279b
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(65))
			m.cycles += 4
		}
	case 0x279b:
		{
			m.ip = 0x279d
			if !m.zf {
				m.ip = uint16(10147)
			}
			m.cycles += 8
		}
	case 0x279d:
		{
			m.ip = 0x27a2
			m.wr8(m.r[11], uint16(0x1a1), uint16(8))
			m.cycles += 8
		}
	case 0x27a2:
		{
			m.ip = 0x27a3
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x27a3:
		{
			m.ip = 0x27a5
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(68))
			m.cycles += 4
		}
	case 0x27a5:
		{
			m.ip = 0x27a7
			if !m.zf {
				m.ip = uint16(10157)
			}
			m.cycles += 8
		}
	case 0x27a7:
		{
			m.ip = 0x27ac
			m.wr8(m.r[11], uint16(0x1a1), uint16(0))
			m.cycles += 8
		}
	case 0x27ac:
		{
			m.ip = 0x27ad
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x27ad:
		{
			m.ip = 0x27af
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x27af:
		{
			m.ip = 0x27b1
			if !m.zf {
				m.ip = uint16(10164)
			}
			m.cycles += 8
		}
	case 0x27b1:
		{
			m.ip = 0x27b4
			target := uint16(3091)
			m.push(0x27b4)
			m.ip = target
			m.cycles += 19
		}
	case 0x27b4:
		{
			m.ip = 0x27b7
			m.alu("sub", 16, m.r[0], uint16(4113))
			m.cycles += 4
		}
	case 0x27b7:
		{
			m.ip = 0x27b9
			if !m.zf {
				m.ip = uint16(10172)
			}
			m.cycles += 8
		}
	case 0x27b9:
		{
			m.ip = 0x27bc
			m.ip = uint16(17160)
			m.cycles += 15
		}
	case 0x27bc:
		{
			m.ip = 0x27bd
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x27bd:
		{
			m.ip = 0x27bf
			m.r[3] = m.r[6]
			m.cycles += 2
		}
	case 0x27bf:
		{
			m.ip = 0x27c4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(2))
			m.cycles += 16
		}
	case 0x27c4:
		{
			m.ip = 0x27c6
			if m.zf {
				m.ip = uint16(10191)
			}
			m.cycles += 8
		}
	case 0x27c6:
		{
			m.ip = 0x27c9
			m.r[3] = uint16(1)
			m.cycles += 2
		}
	case 0x27c9:
		{
			m.ip = 0x27cd
			m.set8(3, 0, m.alu("sub", 8, ((m.r[3]>>0)&255), m.rd8(m.r[11], uint16(0x61a))))
			m.cycles += 16
		}
	case 0x27cd:
		{
			m.ip = 0x27cf
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x27cf:
		{
			m.ip = 0x27d3
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x1a5))
			m.cycles += 8
		}
	case 0x27d3:
		{
			m.ip = 0x27d6
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x27d6:
		{
			m.ip = 0x27d9
			m.alu("sub", 16, m.r[0], uint16(40))
			m.cycles += 4
		}
	case 0x27d9:
		{
			m.ip = 0x27db
			if m.cf || m.zf {
				m.ip = uint16(10209)
			}
			m.cycles += 8
		}
	case 0x27db:
		{
			m.ip = 0x27de
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(21))
			m.cycles += 4
		}
	case 0x27de:
		{
			m.ip = 0x27df
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x27df:
		{
			m.ip = 0x27e1
			m.ip = uint16(10198)
			m.cycles += 15
		}
	case 0x27e1:
		{
			m.ip = 0x27e4
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(22))
			m.cycles += 4
		}
	case 0x27e4:
		{
			m.ip = 0x27e8
			m.wr16(m.r[11], uint16(m.r[6]+0x101b), m.r[0])
			m.cycles += 8
		}
	case 0x27e8:
		{
			m.ip = 0x27ec
			m.wr16(m.r[11], uint16(m.r[6]+0x101f), m.r[3])
			m.cycles += 8
		}
	case 0x27ec:
		{
			m.ip = 0x27ed
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x27ed:
		{
			m.ip = 0x27f1
			m.r[3] = m.rd16(m.r[11], uint16(0x221))
			m.cycles += 8
		}
	case 0x27f1:
		{
			m.ip = 0x27f6
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)), uint16(7))
			m.cycles += 16
		}
	case 0x27f6:
		{
			m.ip = 0x27f8
			if !m.zf {
				m.ip = uint16(10355)
			}
			m.cycles += 8
		}
	case 0x27f8:
		{
			m.ip = 0x27fa
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x27fa:
		{
			m.ip = 0x27ff
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5)), uint16(17))
			m.cycles += 16
		}
	case 0x27ff:
		{
			m.ip = 0x2801
			if m.cf {
				m.ip = uint16(10244)
			}
			m.cycles += 8
		}
	case 0x2801:
		{
			m.ip = 0x2803
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2803:
		{
			m.ip = 0x2804
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2804:
		{
			m.ip = 0x2806
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2806:
		{
			m.ip = 0x2809
			target := uint16(19095)
			m.push(0x2809)
			m.ip = target
			m.cycles += 19
		}
	case 0x2809:
		{
			m.ip = 0x280d
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), m.rd8(m.r[11], uint16(0x24e)))
			m.cycles += 16
		}
	case 0x280d:
		{
			m.ip = 0x280f
			if !m.cf && !m.zf {
				m.ip = uint16(10356)
			}
			m.cycles += 8
		}
	case 0x280f:
		{
			m.ip = 0x2814
			m.wr8(m.r[11], uint16(0x100f), uint16(0))
			m.cycles += 8
		}
	case 0x2814:
		{
			m.ip = 0x2819
			m.wr8(m.r[11], uint16(0x1010), uint16(1))
			m.cycles += 8
		}
	case 0x2819:
		{
			m.ip = 0x281e
			m.wr8(m.r[11], uint16(0x1011), uint16(2))
			m.cycles += 8
		}
	case 0x281e:
		{
			m.ip = 0x2823
			m.wr8(m.r[11], uint16(0x1012), uint16(3))
			m.cycles += 8
		}
	case 0x2823:
		{
			m.ip = 0x2826
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x2826:
		{
			m.ip = 0x2828
			m.set8(0, 8, m.shift("shr", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x2828:
		{
			m.ip = 0x282a
			m.r[6] = m.shift("rcl", 16, m.r[6], uint16(1))
			m.cycles += 8
		}
	case 0x282a:
		{
			m.ip = 0x282c
			m.set8(0, 8, m.shift("shr", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x282c:
		{
			m.ip = 0x282e
			m.r[6] = m.shift("rcl", 16, m.r[6], uint16(1))
			m.cycles += 8
		}
	case 0x282e:
		{
			m.ip = 0x2832
			m.set8(2, 8, m.rd8(m.r[11], uint16(m.r[6]+0x100f)))
			m.cycles += 8
		}
	case 0x2832:
		{
			m.ip = 0x2835
			target := uint16(10341)
			m.push(0x2835)
			m.ip = target
			m.cycles += 19
		}
	case 0x2835:
		{
			m.ip = 0x2838
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x2838:
		{
			m.ip = 0x283a
			m.set8(0, 8, m.shift("shr", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x283a:
		{
			m.ip = 0x283c
			m.r[6] = m.shift("rcl", 16, m.r[6], uint16(1))
			m.cycles += 8
		}
	case 0x283c:
		{
			m.ip = 0x283f
			m.alu("and", 8, ((m.r[0] >> 8) & 255), uint16(1))
			m.cycles += 4
		}
	case 0x283f:
		{
			m.ip = 0x2841
			if m.zf {
				m.ip = uint16(10306)
			}
			m.cycles += 8
		}
	case 0x2841:
		{
			m.ip = 0x2842
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x2842:
		{
			m.ip = 0x2846
			m.set8(2, 0, m.rd8(m.r[11], uint16(m.r[6]+0x100f)))
			m.cycles += 8
		}
	case 0x2846:
		{
			m.ip = 0x2849
			target := uint16(10341)
			m.push(0x2849)
			m.ip = target
			m.cycles += 19
		}
	case 0x2849:
		{
			m.ip = 0x284d
			m.set8(3, 8, m.rd8(m.r[11], uint16(0x100f)))
			m.cycles += 8
		}
	case 0x284d:
		{
			m.ip = 0x2851
			m.set8(3, 0, m.rd8(m.r[11], uint16(0x1010)))
			m.cycles += 8
		}
	case 0x2851:
		{
			m.ip = 0x2854
			m.alu("and", 8, ((m.r[0] >> 8) & 255), uint16(2))
			m.cycles += 4
		}
	case 0x2854:
		{
			m.ip = 0x2856
			if m.zf {
				m.ip = uint16(10336)
			}
			m.cycles += 8
		}
	case 0x2856:
		{
			m.ip = 0x285a
			m.wr8(m.r[11], uint16(0x296), ((m.r[3] >> 8) & 255))
			m.cycles += 8
		}
	case 0x285a:
		{
			m.ip = 0x285c
			m.set8(3, 8, ((m.r[3] >> 0) & 255))
			m.cycles += 2
		}
	case 0x285c:
		{
			m.ip = 0x2860
			m.set8(3, 0, m.rd8(m.r[11], uint16(0x296)))
			m.cycles += 8
		}
	case 0x2860:
		{
			m.ip = 0x2862
			m.r[0] = m.r[2]
			m.cycles += 2
		}
	case 0x2862:
		{
			m.ip = 0x2865
			m.ip = uint16(10471)
			m.cycles += 15
		}
	case 0x2865:
		{
			m.ip = 0x2869
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6]+0x1010)))
			m.cycles += 8
		}
	case 0x2869:
		{
			m.ip = 0x286d
			m.wr8(m.r[11], uint16(m.r[6]+0x100f), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x286d:
		{
			m.ip = 0x286e
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x286e:
		{
			m.ip = 0x2871
			m.alu("sub", 16, m.r[6], uint16(4))
			m.cycles += 4
		}
	case 0x2871:
		{
			m.ip = 0x2873
			if m.cf {
				m.ip = uint16(10341)
			}
			m.cycles += 8
		}
	case 0x2873:
		{
			m.ip = 0x2874
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2874:
		{
			m.ip = 0x2876
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2876:
		{
			m.ip = 0x287a
			m.r[1] = m.rd16(m.r[11], uint16(m.r[3]+0x1bd))
			m.cycles += 8
		}
	case 0x287a:
		{
			m.ip = 0x287d
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x287d:
		{
			m.ip = 0x2880
			m.alu("sub", 16, m.r[1], uint16(40))
			m.cycles += 4
		}
	case 0x2880:
		{
			m.ip = 0x2882
			if m.cf || m.zf {
				m.ip = uint16(10376)
			}
			m.cycles += 8
		}
	case 0x2882:
		{
			m.ip = 0x2885
			m.r[1] = m.alu("sub", 16, m.r[1], uint16(21))
			m.cycles += 4
		}
	case 0x2885:
		{
			m.ip = 0x2886
			m.r[2] = m.unary("inc", 16, m.r[2])
			m.cycles += 3
		}
	case 0x2886:
		{
			m.ip = 0x2888
			m.ip = uint16(10365)
			m.cycles += 15
		}
	case 0x2888:
		{
			m.ip = 0x288b
			m.r[1] = m.alu("sub", 16, m.r[1], uint16(22))
			m.cycles += 4
		}
	case 0x288b:
		{
			m.ip = 0x288e
			m.r[0] = uint16(259)
			m.cycles += 2
		}
	case 0x288e:
		{
			m.ip = 0x2891
			m.r[3] = uint16(2)
			m.cycles += 2
		}
	case 0x2891:
		{
			m.ip = 0x2895
			m.r[6] = m.rd16(m.r[11], uint16(0x221))
			m.cycles += 8
		}
	case 0x2895:
		{
			m.ip = 0x289a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x7a)), uint16(71))
			m.cycles += 16
		}
	case 0x289a:
		{
			m.ip = 0x289c
			if m.zf {
				m.ip = uint16(10405)
			}
			m.cycles += 8
		}
	case 0x289c:
		{
			m.ip = 0x289f
			m.wr16(m.r[11], uint16(0x297), m.r[0])
			m.cycles += 8
		}
	case 0x289f:
		{
			m.ip = 0x28a1
			m.r[0] = m.r[3]
			m.cycles += 2
		}
	case 0x28a1:
		{
			m.ip = 0x28a5
			m.r[3] = m.rd16(m.r[11], uint16(0x297))
			m.cycles += 8
		}
	case 0x28a5:
		{
			m.ip = 0x28a8
			m.r[6] = m.alu("and", 16, m.r[6], uint16(1))
			m.cycles += 4
		}
	case 0x28a8:
		{
			m.ip = 0x28aa
			m.r[6] = m.shift("shl", 16, m.r[6], uint16(1))
			m.cycles += 8
		}
	case 0x28aa:
		{
			m.ip = 0x28ae
			m.r[1] = m.alu("sub", 16, m.r[1], m.rd16(m.r[11], uint16(m.r[6]+0x101b)))
			m.cycles += 16
		}
	case 0x28ae:
		{
			m.ip = 0x28b0
			if !m.sf {
				m.ip = uint16(10428)
			}
			m.cycles += 8
		}
	case 0x28b0:
		{
			m.ip = 0x28b4
			m.wr8(m.r[11], uint16(0x296), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x28b4:
		{
			m.ip = 0x28b6
			m.set8(0, 8, ((m.r[3] >> 8) & 255))
			m.cycles += 2
		}
	case 0x28b6:
		{
			m.ip = 0x28ba
			m.set8(3, 8, m.rd8(m.r[11], uint16(0x296)))
			m.cycles += 8
		}
	case 0x28ba:
		{
			m.ip = 0x28bc
			m.r[1] = m.unary("neg", 16, m.r[1])
			m.cycles += 3
		}
	case 0x28bc:
		{
			m.ip = 0x28c0
			m.r[2] = m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[6]+0x101f)))
			m.cycles += 16
		}
	case 0x28c0:
		{
			m.ip = 0x28c2
			if !m.sf {
				m.ip = uint16(10445)
			}
			m.cycles += 8
		}
	case 0x28c2:
		{
			m.ip = 0x28c5
			m.wr16(m.r[11], uint16(0x296), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x28c5:
		{
			m.ip = 0x28c7
			m.set8(0, 0, ((m.r[3] >> 0) & 255))
			m.cycles += 2
		}
	case 0x28c7:
		{
			m.ip = 0x28cb
			m.set8(3, 0, m.rd8(m.r[11], uint16(0x296)))
			m.cycles += 8
		}
	case 0x28cb:
		{
			m.ip = 0x28cd
			m.r[2] = m.unary("neg", 16, m.r[2])
			m.cycles += 3
		}
	case 0x28cd:
		{
			m.ip = 0x28cf
			m.alu("sub", 16, m.r[1], m.r[2])
			m.cycles += 4
		}
	case 0x28cf:
		{
			m.ip = 0x28d1
			if !m.cf && !m.zf {
				m.ip = uint16(10461)
			}
			m.cycles += 8
		}
	case 0x28d1:
		{
			m.ip = 0x28d5
			m.wr8(m.r[11], uint16(0x296), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x28d5:
		{
			m.ip = 0x28d7
			m.set8(0, 8, ((m.r[0] >> 0) & 255))
			m.cycles += 2
		}
	case 0x28d7:
		{
			m.ip = 0x28da
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x296)))
			m.cycles += 8
		}
	case 0x28da:
		{
			m.ip = 0x28dc
			m.ip = uint16(10471)
			m.cycles += 15
		}
	case 0x28dc:
		{
			m.ip = 0x28dd
			m.cycles += 3
		}
	case 0x28dd:
		{
			m.ip = 0x28e1
			m.wr8(m.r[11], uint16(0x296), ((m.r[3] >> 8) & 255))
			m.cycles += 8
		}
	case 0x28e1:
		{
			m.ip = 0x28e3
			m.set8(3, 8, ((m.r[3] >> 0) & 255))
			m.cycles += 2
		}
	case 0x28e3:
		{
			m.ip = 0x28e7
			m.set8(3, 0, m.rd8(m.r[11], uint16(0x296)))
			m.cycles += 8
		}
	case 0x28e7:
		{
			m.ip = 0x28eb
			m.wr8(m.r[11], uint16(0x100f), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x28eb:
		{
			m.ip = 0x28ee
			m.wr16(m.r[11], uint16(0x1010), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x28ee:
		{
			m.ip = 0x28f2
			m.wr8(m.r[11], uint16(0x1011), ((m.r[3] >> 8) & 255))
			m.cycles += 8
		}
	case 0x28f2:
		{
			m.ip = 0x28f6
			m.wr8(m.r[11], uint16(0x1012), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x28f6:
		{
			m.ip = 0x28fa
			m.r[3] = m.rd16(m.r[11], uint16(0x221))
			m.cycles += 8
		}
	case 0x28fa:
		{
			m.ip = 0x28fc
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x28fc:
		{
			m.ip = 0x2900
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x1bd))
			m.cycles += 8
		}
	case 0x2900:
		{
			m.ip = 0x2902
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2902:
		{
			m.ip = 0x2905
			m.r[6] = uint16(0)
			m.cycles += 2
		}
	case 0x2905:
		{
			m.ip = 0x2909
			m.set8(0, 8, m.rd8(m.r[11], uint16(m.r[6]+0x100f)))
			m.cycles += 8
		}
	case 0x2909:
		{
			m.ip = 0x290b
			m.set8(1, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x290b:
		{
			m.ip = 0x290d
			m.set8(0, 0, uint16(16))
			m.cycles += 2
		}
	case 0x290d:
		{
			m.ip = 0x290f
			m.set8(0, 0, m.shift("shl", 8, ((m.r[0]>>0)&255), ((m.r[1]>>0)&255)))
			m.cycles += 8
		}
	case 0x290f:
		{
			m.ip = 0x2914
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(m.r[5]+0x8e))))
			m.cycles += 16
		}
	case 0x2914:
		{
			m.ip = 0x2916
			if !m.zf {
				m.ip = uint16(10521)
			}
			m.cycles += 8
		}
	case 0x2916:
		{
			m.ip = 0x2919
			m.ip = uint16(10661)
			m.cycles += 15
		}
	case 0x2919:
		{
			m.ip = 0x291d
			m.r[7] = uint16(0x1013)
			m.cycles += 3
		}
	case 0x291d:
		{
			m.ip = 0x291f
			m.set8(1, 8, m.alu("xor", 8, ((m.r[1]>>8)&255), ((m.r[1]>>8)&255)))
			m.cycles += 4
		}
	case 0x291f:
		{
			m.ip = 0x2921
			m.set8(1, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2921:
		{
			m.ip = 0x2923
			m.r[1] = m.shift("shl", 16, m.r[1], uint16(1))
			m.cycles += 8
		}
	case 0x2923:
		{
			m.ip = 0x2925
			m.r[7] = m.alu("add", 16, m.r[7], m.r[1])
			m.cycles += 4
		}
	case 0x2925:
		{
			m.ip = 0x2927
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]))
			m.cycles += 8
		}
	case 0x2927:
		{
			m.ip = 0x292d
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+m.r[7]+0x8e)), uint16(2))
			m.cycles += 16
		}
	case 0x292d:
		{
			m.ip = 0x292f
			if !m.zf {
				m.ip = uint16(10661)
			}
			m.cycles += 8
		}
	case 0x292f:
		{
			m.ip = 0x2931
			m.r[1] = m.r[7]
			m.cycles += 2
		}
	case 0x2931:
		{
			m.ip = 0x2933
			m.r[1] = m.alu("add", 16, m.r[1], m.r[5])
			m.cycles += 4
		}
	case 0x2933:
		{
			m.ip = 0x2936
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x2936:
		{
			m.ip = 0x293a
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[7]+0x1bd)), m.r[1])
			m.cycles += 16
		}
	case 0x293a:
		{
			m.ip = 0x293c
			if m.zf {
				m.ip = uint16(10661)
			}
			m.cycles += 8
		}
	case 0x293c:
		{
			m.ip = 0x293f
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x293f:
		{
			m.ip = 0x2942
			m.alu("sub", 16, m.r[7], uint16(8))
			m.cycles += 4
		}
	case 0x2942:
		{
			m.ip = 0x2944
			if !m.zf {
				m.ip = uint16(10550)
			}
			m.cycles += 8
		}
	case 0x2944:
		{
			m.ip = 0x2946
			m.set8(0, 8, m.shift("shl", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x2946:
		{
			m.ip = 0x2948
			m.set8(0, 8, m.shift("shl", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x2948:
		{
			m.ip = 0x294a
			m.set8(0, 8, m.shift("shl", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x294a:
		{
			m.ip = 0x294c
			m.set8(1, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x294c:
		{
			m.ip = 0x2950
			m.set8(1, 0, m.alu("xor", 8, ((m.r[1]>>0)&255), m.rd8(m.r[11], uint16(m.r[3]+0x1b1))))
			m.cycles += 16
		}
	case 0x2950:
		{
			m.ip = 0x2953
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), uint16(8))
			m.cycles += 4
		}
	case 0x2953:
		{
			m.ip = 0x2955
			if !m.zf {
				m.ip = uint16(10672)
			}
			m.cycles += 8
		}
	case 0x2955:
		{
			m.ip = 0x295a
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[5]+0x8e)))
			m.cycles += 8
		}
	case 0x295a:
		{
			m.ip = 0x295e
			m.r[7] = uint16(0x1013)
			m.cycles += 3
		}
	case 0x295e:
		{
			m.ip = 0x2960
			m.set8(1, 0, uint16(239))
			m.cycles += 2
		}
	case 0x2960:
		{
			m.ip = 0x2961
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x2961:
		{
			m.ip = 0x2963
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]))
			m.cycles += 8
		}
	case 0x2963:
		{
			m.ip = 0x2965
			m.r[2] = m.r[7]
			m.cycles += 2
		}
	case 0x2965:
		{
			m.ip = 0x2967
			m.r[2] = m.alu("add", 16, m.r[2], m.r[5])
			m.cycles += 4
		}
	case 0x2967:
		{
			m.ip = 0x296d
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+m.r[7]+0x8e)), uint16(2))
			m.cycles += 16
		}
	case 0x296d:
		{
			m.ip = 0x296e
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x296e:
		{
			m.ip = 0x2970
			if !m.zf {
				m.ip = uint16(10632)
			}
			m.cycles += 8
		}
	case 0x2970:
		{
			m.ip = 0x2974
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(0x1bd)))
			m.cycles += 16
		}
	case 0x2974:
		{
			m.ip = 0x2976
			if m.zf {
				m.ip = uint16(10632)
			}
			m.cycles += 8
		}
	case 0x2976:
		{
			m.ip = 0x297a
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(0x1bf)))
			m.cycles += 16
		}
	case 0x297a:
		{
			m.ip = 0x297c
			if m.zf {
				m.ip = uint16(10632)
			}
			m.cycles += 8
		}
	case 0x297c:
		{
			m.ip = 0x2980
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(0x1c1)))
			m.cycles += 16
		}
	case 0x2980:
		{
			m.ip = 0x2982
			if m.zf {
				m.ip = uint16(10632)
			}
			m.cycles += 8
		}
	case 0x2982:
		{
			m.ip = 0x2986
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(0x1c3)))
			m.cycles += 16
		}
	case 0x2986:
		{
			m.ip = 0x2988
			if !m.zf {
				m.ip = uint16(10634)
			}
			m.cycles += 8
		}
	case 0x2988:
		{
			m.ip = 0x298a
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), ((m.r[1]>>0)&255)))
			m.cycles += 4
		}
	case 0x298a:
		{
			m.ip = 0x298c
			m.set8(1, 0, m.shift("rol", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x298c:
		{
			m.ip = 0x298f
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x298f:
		{
			m.ip = 0x2992
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), uint16(254))
			m.cycles += 4
		}
	case 0x2992:
		{
			m.ip = 0x2994
			if !m.zf {
				m.ip = uint16(10592)
			}
			m.cycles += 8
		}
	case 0x2994:
		{
			m.ip = 0x2996
			m.set8(2, 0, uint16(0))
			m.cycles += 2
		}
	case 0x2996:
		{
			m.ip = 0x2998
			m.set8(1, 0, uint16(4))
			m.cycles += 2
		}
	case 0x2998:
		{
			m.ip = 0x299a
			m.set8(0, 0, m.shift("rol", 8, ((m.r[0]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x299a:
		{
			m.ip = 0x299c
			if !m.cf {
				m.ip = uint16(10654)
			}
			m.cycles += 8
		}
	case 0x299c:
		{
			m.ip = 0x299e
			m.set8(2, 0, m.unary("inc", 8, ((m.r[2]>>0)&255)))
			m.cycles += 3
		}
	case 0x299e:
		{
			m.ip = 0x29a0
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(10648)
			}
			m.cycles += 17
		}
	case 0x29a0:
		{
			m.ip = 0x29a3
			m.alu("sub", 8, ((m.r[2] >> 0) & 255), uint16(1))
			m.cycles += 4
		}
	case 0x29a3:
		{
			m.ip = 0x29a5
			if m.zf {
				m.ip = uint16(10672)
			}
			m.cycles += 8
		}
	case 0x29a5:
		{
			m.ip = 0x29a6
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x29a6:
		{
			m.ip = 0x29a9
			m.alu("sub", 16, m.r[6], uint16(4))
			m.cycles += 4
		}
	case 0x29a9:
		{
			m.ip = 0x29ab
			if m.zf {
				m.ip = uint16(10670)
			}
			m.cycles += 8
		}
	case 0x29ab:
		{
			m.ip = 0x29ae
			m.ip = uint16(10501)
			m.cycles += 15
		}
	case 0x29ae:
		{
			m.ip = 0x29b0
			m.set8(0, 8, uint16(255))
			m.cycles += 2
		}
	case 0x29b0:
		{
			m.ip = 0x29b4
			m.wr8(m.r[11], uint16(m.r[3]+0x1b5), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x29b4:
		{
			m.ip = 0x29b5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x29b5:
		{
			m.ip = 0x29b8
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x29b8:
		{
			m.ip = 0x29bb
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x29bb:
		{
			m.ip = 0x29be
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x29be:
		{
			m.ip = 0x29c2
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x29c2:
		{
			m.ip = 0x29c6
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x29c6:
		{
			m.ip = 0x29ca
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x29ca:
		{
			m.ip = 0x29ce
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x29ce:
		{
			m.ip = 0x29d2
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x29d2:
		{
			m.ip = 0x29d6
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x29d6:
		{
			m.ip = 0x29d9
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x29d9:
		{
			m.ip = 0x29dc
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x29dc:
		{
			m.ip = 0x29de
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(10680)
			}
			m.cycles += 17
		}
	case 0x29de:
		{
			m.ip = 0x29df
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x29df:
		{
			m.ip = 0x29e2
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x29e2:
		{
			m.ip = 0x29e5
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x29e5:
		{
			m.ip = 0x29e8
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x29e8:
		{
			m.ip = 0x29ec
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x29ec:
		{
			m.ip = 0x29f0
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x29f0:
		{
			m.ip = 0x29f4
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x29f4:
		{
			m.ip = 0x29f8
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x29f8:
		{
			m.ip = 0x29fc
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x29fc:
		{
			m.ip = 0x2a00
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x2a00:
		{
			m.ip = 0x2a04
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x2a04:
		{
			m.ip = 0x2a08
			m.wr8(m.r[8], uint16(m.r[7]+0x4), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x2a08:
		{
			m.ip = 0x2a0b
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x2a0b:
		{
			m.ip = 0x2a0e
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x2a0e:
		{
			m.ip = 0x2a10
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(10722)
			}
			m.cycles += 17
		}
	case 0x2a10:
		{
			m.ip = 0x2a11
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2a11:
		{
			m.ip = 0x2a14
			target := uint16(11471)
			m.push(0x2a14)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a14:
		{
			m.ip = 0x2a17
			target := uint16(11325)
			m.push(0x2a17)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a17:
		{
			m.ip = 0x2a1a
			target := uint16(11821)
			m.push(0x2a1a)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a1a:
		{
			m.ip = 0x2a1e
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x2a1e:
		{
			m.ip = 0x2a22
			m.ip = m.rd16(m.r[11], uint16(m.r[3]+0x49e))
			m.cycles += 15
		}
	case 0x2a22:
		{
			m.ip = 0x2a25
			target := uint16(12120)
			m.push(0x2a25)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a25:
		{
			m.ip = 0x2a2b
			m.wr16(m.r[11], uint16(0x223), uint16(0))
			m.cycles += 8
		}
	case 0x2a2b:
		{
			m.ip = 0x2a31
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x2a31:
		{
			m.ip = 0x2a34
			target := uint16(12567)
			m.push(0x2a34)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a34:
		{
			m.ip = 0x2a39
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x2a39:
		{
			m.ip = 0x2a3b
			if !m.zf {
				m.ip = uint16(10826)
			}
			m.cycles += 8
		}
	case 0x2a3b:
		{
			m.ip = 0x2a41
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x2a41:
		{
			m.ip = 0x2a47
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x2a47:
		{
			m.ip = 0x2a4a
			target := uint16(12567)
			m.push(0x2a4a)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a4a:
		{
			m.ip = 0x2a4d
			target := uint16(11077)
			m.push(0x2a4d)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a4d:
		{
			m.ip = 0x2a50
			target := uint16(11230)
			m.push(0x2a50)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a50:
		{
			m.ip = 0x2a51
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2a51:
		{
			m.ip = 0x2a54
			target := uint16(10982)
			m.push(0x2a54)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a54:
		{
			m.ip = 0x2a57
			target := uint16(11128)
			m.push(0x2a57)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a57:
		{
			m.ip = 0x2a5a
			target := uint16(11179)
			m.push(0x2a5a)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a5a:
		{
			m.ip = 0x2a5b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2a5b:
		{
			m.ip = 0x2a5c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2a5c:
		{
			m.ip = 0x2a5d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2a5d:
		{
			m.ip = 0x2a60
			target := uint16(10928)
			m.push(0x2a60)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a60:
		{
			m.ip = 0x2a63
			target := uint16(10853)
			m.push(0x2a63)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a63:
		{
			m.ip = 0x2a64
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2a64:
		{
			m.ip = 0x2a65
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2a65:
		{
			m.ip = 0x2a6a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x46a)), uint16(1))
			m.cycles += 16
		}
	case 0x2a6a:
		{
			m.ip = 0x2a6c
			if !m.zf {
				m.ip = uint16(10927)
			}
			m.cycles += 8
		}
	case 0x2a6c:
		{
			m.ip = 0x2a6f
			m.r[0] = m.rd16(m.r[11], uint16(0x5e))
			m.cycles += 8
		}
	case 0x2a6f:
		{
			m.ip = 0x2a73
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(0x60)))
			m.cycles += 16
		}
	case 0x2a73:
		{
			m.ip = 0x2a77
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(0x62)))
			m.cycles += 16
		}
	case 0x2a77:
		{
			m.ip = 0x2a7b
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(0x64)))
			m.cycles += 16
		}
	case 0x2a7b:
		{
			m.ip = 0x2a7e
			m.alu("sub", 16, m.r[0], uint16(1020))
			m.cycles += 4
		}
	case 0x2a7e:
		{
			m.ip = 0x2a80
			if !m.zf {
				m.ip = uint16(10927)
			}
			m.cycles += 8
		}
	case 0x2a80:
		{
			m.ip = 0x2a85
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x35)), uint16(0))
			m.cycles += 16
		}
	case 0x2a85:
		{
			m.ip = 0x2a87
			if m.zf {
				m.ip = uint16(10892)
			}
			m.cycles += 8
		}
	case 0x2a87:
		{
			m.ip = 0x2a8b
			m.wr16(m.r[11], uint16(0x35), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x35))))
			m.cycles += 15
		}
	case 0x2a8b:
		{
			m.ip = 0x2a8c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2a8c:
		{
			m.ip = 0x2a8f
			target := uint16(19095)
			m.push(0x2a8f)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a8f:
		{
			m.ip = 0x2a91
			m.set8(3, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2a91:
		{
			m.ip = 0x2a94
			m.r[3] = m.alu("and", 16, m.r[3], uint16(3))
			m.cycles += 4
		}
	case 0x2a94:
		{
			m.ip = 0x2a96
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2a96:
		{
			m.ip = 0x2a9c
			m.wr16(m.r[11], uint16(m.r[3]+0x5e), uint16(0))
			m.cycles += 8
		}
	case 0x2a9c:
		{
			m.ip = 0x2a9f
			target := uint16(19095)
			m.push(0x2a9f)
			m.ip = target
			m.cycles += 19
		}
	case 0x2a9f:
		{
			m.ip = 0x2aa1
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2aa1:
		{
			m.ip = 0x2aa3
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x2aa3:
		{
			m.ip = 0x2aa7
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(0x33)))
			m.cycles += 16
		}
	case 0x2aa7:
		{
			m.ip = 0x2aaa
			m.wr16(m.r[11], uint16(0x35), m.r[0])
			m.cycles += 8
		}
	case 0x2aaa:
		{
			m.ip = 0x2aaf
			m.wr16(m.r[11], uint16(0x33), m.alu("add", 16, m.rd16(m.r[11], uint16(0x33)), uint16(15)))
			m.cycles += 16
		}
	case 0x2aaf:
		{
			m.ip = 0x2ab0
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2ab0:
		{
			m.ip = 0x2ab5
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x30)), uint16(0))
			m.cycles += 16
		}
	case 0x2ab5:
		{
			m.ip = 0x2ab7
			if m.zf {
				m.ip = uint16(10940)
			}
			m.cycles += 8
		}
	case 0x2ab7:
		{
			m.ip = 0x2abb
			m.wr8(m.r[11], uint16(0x30), m.unary("dec", 8, m.rd8(m.r[11], uint16(0x30))))
			m.cycles += 15
		}
	case 0x2abb:
		{
			m.ip = 0x2abc
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2abc:
		{
			m.ip = 0x2ac1
			m.wr8(m.r[11], uint16(0x30), uint16(1))
			m.cycles += 8
		}
	case 0x2ac1:
		{
			m.ip = 0x2ac5
			m.r[6] = m.rd16(m.r[11], uint16(0x4ec))
			m.cycles += 8
		}
	case 0x2ac5:
		{
			m.ip = 0x2ac9
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x31)))
			m.cycles += 16
		}
	case 0x2ac9:
		{
			m.ip = 0x2acd
			m.r[7] = m.rd16(m.r[11], uint16(0x4ee))
			m.cycles += 8
		}
	case 0x2acd:
		{
			m.ip = 0x2ad0
			m.r[1] = uint16(48)
			m.cycles += 2
		}
	case 0x2ad0:
		{
			m.ip = 0x2ad3
			target := uint16(10680)
			m.push(0x2ad3)
			m.ip = target
			m.cycles += 19
		}
	case 0x2ad3:
		{
			m.ip = 0x2ad8
			m.wr16(m.r[11], uint16(0x31), m.alu("add", 16, m.rd16(m.r[11], uint16(0x31)), uint16(4)))
			m.cycles += 16
		}
	case 0x2ad8:
		{
			m.ip = 0x2add
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x31)), uint16(16))
			m.cycles += 16
		}
	case 0x2add:
		{
			m.ip = 0x2adf
			if !m.zf {
				m.ip = uint16(10981)
			}
			m.cycles += 8
		}
	case 0x2adf:
		{
			m.ip = 0x2ae5
			m.wr16(m.r[11], uint16(0x31), uint16(0))
			m.cycles += 8
		}
	case 0x2ae5:
		{
			m.ip = 0x2ae6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2ae6:
		{
			m.ip = 0x2aeb
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1b)), uint16(0))
			m.cycles += 16
		}
	case 0x2aeb:
		{
			m.ip = 0x2aed
			if m.zf {
				m.ip = uint16(10994)
			}
			m.cycles += 8
		}
	case 0x2aed:
		{
			m.ip = 0x2af1
			m.wr8(m.r[11], uint16(0x1b), m.unary("dec", 8, m.rd8(m.r[11], uint16(0x1b))))
			m.cycles += 15
		}
	case 0x2af1:
		{
			m.ip = 0x2af2
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2af2:
		{
			m.ip = 0x2af7
			m.wr8(m.r[11], uint16(0x1b), uint16(1))
			m.cycles += 8
		}
	case 0x2af7:
		{
			m.ip = 0x2afb
			m.r[6] = m.rd16(m.r[11], uint16(0x4ae))
			m.cycles += 8
		}
	case 0x2afb:
		{
			m.ip = 0x2aff
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x1d)))
			m.cycles += 16
		}
	case 0x2aff:
		{
			m.ip = 0x2b03
			m.r[7] = m.rd16(m.r[11], uint16(0x4b0))
			m.cycles += 8
		}
	case 0x2b03:
		{
			m.ip = 0x2b06
			m.r[1] = uint16(12)
			m.cycles += 2
		}
	case 0x2b06:
		{
			m.ip = 0x2b07
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x2b07:
		{
			m.ip = 0x2b0a
			m.r[1] = uint16(8)
			m.cycles += 2
		}
	case 0x2b0a:
		{
			m.ip = 0x2b0d
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x2b0d:
		{
			m.ip = 0x2b10
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x2b10:
		{
			m.ip = 0x2b11
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x2b11:
		{
			m.ip = 0x2b12
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x2b12:
		{
			m.ip = 0x2b14
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(11018)
			}
			m.cycles += 17
		}
	case 0x2b14:
		{
			m.ip = 0x2b17
			m.r[6] = m.alu("add", 16, m.r[6], uint16(72))
			m.cycles += 4
		}
	case 0x2b17:
		{
			m.ip = 0x2b1a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(72))
			m.cycles += 4
		}
	case 0x2b1a:
		{
			m.ip = 0x2b1b
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x2b1b:
		{
			m.ip = 0x2b1d
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(11014)
			}
			m.cycles += 17
		}
	case 0x2b1d:
		{
			m.ip = 0x2b20
			m.r[0] = uint16(960)
			m.cycles += 2
		}
	case 0x2b20:
		{
			m.ip = 0x2b25
			m.alu("and", 8, m.rd8(m.r[11], uint16(0x1c)), uint16(1))
			m.cycles += 16
		}
	case 0x2b25:
		{
			m.ip = 0x2b27
			if m.zf {
				m.ip = uint16(11050)
			}
			m.cycles += 8
		}
	case 0x2b27:
		{
			m.ip = 0x2b2a
			m.r[0] = uint16(64584)
			m.cycles += 2
		}
	case 0x2b2a:
		{
			m.ip = 0x2b2e
			m.wr16(m.r[11], uint16(0x1d), m.alu("add", 16, m.rd16(m.r[11], uint16(0x1d)), m.r[0]))
			m.cycles += 16
		}
	case 0x2b2e:
		{
			m.ip = 0x2b33
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1c)), uint16(10))
			m.cycles += 16
		}
	case 0x2b33:
		{
			m.ip = 0x2b35
			if !m.zf {
				m.ip = uint16(11072)
			}
			m.cycles += 8
		}
	case 0x2b35:
		{
			m.ip = 0x2b3a
			m.wr8(m.r[11], uint16(0x1c), uint16(255))
			m.cycles += 8
		}
	case 0x2b3a:
		{
			m.ip = 0x2b40
			m.wr16(m.r[11], uint16(0x1d), uint16(0))
			m.cycles += 8
		}
	case 0x2b40:
		{
			m.ip = 0x2b44
			m.wr8(m.r[11], uint16(0x1c), m.unary("inc", 8, m.rd8(m.r[11], uint16(0x1c))))
			m.cycles += 15
		}
	case 0x2b44:
		{
			m.ip = 0x2b45
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2b45:
		{
			m.ip = 0x2b4a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x18)), uint16(0))
			m.cycles += 16
		}
	case 0x2b4a:
		{
			m.ip = 0x2b4c
			if m.zf {
				m.ip = uint16(11089)
			}
			m.cycles += 8
		}
	case 0x2b4c:
		{
			m.ip = 0x2b50
			m.wr8(m.r[11], uint16(0x18), m.unary("dec", 8, m.rd8(m.r[11], uint16(0x18))))
			m.cycles += 15
		}
	case 0x2b50:
		{
			m.ip = 0x2b51
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2b51:
		{
			m.ip = 0x2b56
			m.wr8(m.r[11], uint16(0x18), uint16(1))
			m.cycles += 8
		}
	case 0x2b56:
		{
			m.ip = 0x2b5a
			m.r[6] = m.rd16(m.r[11], uint16(0x4aa))
			m.cycles += 8
		}
	case 0x2b5a:
		{
			m.ip = 0x2b5e
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x19)))
			m.cycles += 16
		}
	case 0x2b5e:
		{
			m.ip = 0x2b62
			m.r[7] = m.rd16(m.r[11], uint16(0x4ac))
			m.cycles += 8
		}
	case 0x2b62:
		{
			m.ip = 0x2b65
			target := uint16(10677)
			m.push(0x2b65)
			m.ip = target
			m.cycles += 19
		}
	case 0x2b65:
		{
			m.ip = 0x2b6a
			m.wr16(m.r[11], uint16(0x19), m.alu("add", 16, m.rd16(m.r[11], uint16(0x19)), uint16(4)))
			m.cycles += 16
		}
	case 0x2b6a:
		{
			m.ip = 0x2b6f
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x19)), uint16(20))
			m.cycles += 16
		}
	case 0x2b6f:
		{
			m.ip = 0x2b71
			if !m.zf {
				m.ip = uint16(11127)
			}
			m.cycles += 8
		}
	case 0x2b71:
		{
			m.ip = 0x2b77
			m.wr16(m.r[11], uint16(0x19), uint16(0))
			m.cycles += 8
		}
	case 0x2b77:
		{
			m.ip = 0x2b78
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2b78:
		{
			m.ip = 0x2b7d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1f)), uint16(0))
			m.cycles += 16
		}
	case 0x2b7d:
		{
			m.ip = 0x2b7f
			if m.zf {
				m.ip = uint16(11140)
			}
			m.cycles += 8
		}
	case 0x2b7f:
		{
			m.ip = 0x2b83
			m.wr8(m.r[11], uint16(0x1f), m.unary("dec", 8, m.rd8(m.r[11], uint16(0x1f))))
			m.cycles += 15
		}
	case 0x2b83:
		{
			m.ip = 0x2b84
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2b84:
		{
			m.ip = 0x2b89
			m.wr8(m.r[11], uint16(0x1f), uint16(1))
			m.cycles += 8
		}
	case 0x2b89:
		{
			m.ip = 0x2b8d
			m.r[6] = m.rd16(m.r[11], uint16(0x4b2))
			m.cycles += 8
		}
	case 0x2b8d:
		{
			m.ip = 0x2b91
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x20)))
			m.cycles += 16
		}
	case 0x2b91:
		{
			m.ip = 0x2b95
			m.r[7] = m.rd16(m.r[11], uint16(0x4b4))
			m.cycles += 8
		}
	case 0x2b95:
		{
			m.ip = 0x2b98
			target := uint16(10677)
			m.push(0x2b98)
			m.ip = target
			m.cycles += 19
		}
	case 0x2b98:
		{
			m.ip = 0x2b9d
			m.wr16(m.r[11], uint16(0x20), m.alu("add", 16, m.rd16(m.r[11], uint16(0x20)), uint16(4)))
			m.cycles += 16
		}
	case 0x2b9d:
		{
			m.ip = 0x2ba2
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x20)), uint16(80))
			m.cycles += 16
		}
	case 0x2ba2:
		{
			m.ip = 0x2ba4
			if !m.zf {
				m.ip = uint16(11178)
			}
			m.cycles += 8
		}
	case 0x2ba4:
		{
			m.ip = 0x2baa
			m.wr16(m.r[11], uint16(0x20), uint16(0))
			m.cycles += 8
		}
	case 0x2baa:
		{
			m.ip = 0x2bab
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2bab:
		{
			m.ip = 0x2bb0
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x22)), uint16(0))
			m.cycles += 16
		}
	case 0x2bb0:
		{
			m.ip = 0x2bb2
			if m.zf {
				m.ip = uint16(11191)
			}
			m.cycles += 8
		}
	case 0x2bb2:
		{
			m.ip = 0x2bb6
			m.wr8(m.r[11], uint16(0x22), m.unary("dec", 8, m.rd8(m.r[11], uint16(0x22))))
			m.cycles += 15
		}
	case 0x2bb6:
		{
			m.ip = 0x2bb7
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2bb7:
		{
			m.ip = 0x2bbc
			m.wr8(m.r[11], uint16(0x22), uint16(5))
			m.cycles += 8
		}
	case 0x2bbc:
		{
			m.ip = 0x2bc0
			m.r[6] = m.rd16(m.r[11], uint16(0x4b6))
			m.cycles += 8
		}
	case 0x2bc0:
		{
			m.ip = 0x2bc4
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x23)))
			m.cycles += 16
		}
	case 0x2bc4:
		{
			m.ip = 0x2bc8
			m.r[7] = m.rd16(m.r[11], uint16(0x4b8))
			m.cycles += 8
		}
	case 0x2bc8:
		{
			m.ip = 0x2bcb
			target := uint16(10677)
			m.push(0x2bcb)
			m.ip = target
			m.cycles += 19
		}
	case 0x2bcb:
		{
			m.ip = 0x2bd0
			m.wr16(m.r[11], uint16(0x23), m.alu("add", 16, m.rd16(m.r[11], uint16(0x23)), uint16(4)))
			m.cycles += 16
		}
	case 0x2bd0:
		{
			m.ip = 0x2bd5
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x23)), uint16(24))
			m.cycles += 16
		}
	case 0x2bd5:
		{
			m.ip = 0x2bd7
			if !m.zf {
				m.ip = uint16(11229)
			}
			m.cycles += 8
		}
	case 0x2bd7:
		{
			m.ip = 0x2bdd
			m.wr16(m.r[11], uint16(0x23), uint16(0))
			m.cycles += 8
		}
	case 0x2bdd:
		{
			m.ip = 0x2bde
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2bde:
		{
			m.ip = 0x2be3
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x25)), uint16(0))
			m.cycles += 16
		}
	case 0x2be3:
		{
			m.ip = 0x2be5
			if m.zf {
				m.ip = uint16(11242)
			}
			m.cycles += 8
		}
	case 0x2be5:
		{
			m.ip = 0x2be9
			m.wr8(m.r[11], uint16(0x25), m.unary("dec", 8, m.rd8(m.r[11], uint16(0x25))))
			m.cycles += 15
		}
	case 0x2be9:
		{
			m.ip = 0x2bea
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2bea:
		{
			m.ip = 0x2bef
			m.wr8(m.r[11], uint16(0x25), uint16(1))
			m.cycles += 8
		}
	case 0x2bef:
		{
			m.ip = 0x2bf2
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x2bf2:
		{
			m.ip = 0x2bf6
			m.r[6] = m.rd16(m.r[11], uint16(m.r[3]+0x4ba))
			m.cycles += 8
		}
	case 0x2bf6:
		{
			m.ip = 0x2bfa
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(m.r[3]+0x26)))
			m.cycles += 16
		}
	case 0x2bfa:
		{
			m.ip = 0x2bfe
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x4c4))
			m.cycles += 8
		}
	case 0x2bfe:
		{
			m.ip = 0x2c02
			m.r[2] = m.rd16(m.r[11], uint16(m.r[3]+0x4d8))
			m.cycles += 8
		}
	case 0x2c02:
		{
			m.ip = 0x2c06
			m.r[1] = m.rd16(m.r[11], uint16(m.r[3]+0x4ce))
			m.cycles += 8
		}
	case 0x2c06:
		{
			m.ip = 0x2c07
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x2c07:
		{
			m.ip = 0x2c09
			m.r[1] = m.r[2]
			m.cycles += 2
		}
	case 0x2c09:
		{
			m.ip = 0x2c0c
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x2c0c:
		{
			m.ip = 0x2c0f
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x2c0f:
		{
			m.ip = 0x2c10
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x2c10:
		{
			m.ip = 0x2c11
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x2c11:
		{
			m.ip = 0x2c13
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(11273)
			}
			m.cycles += 17
		}
	case 0x2c13:
		{
			m.ip = 0x2c14
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x2c14:
		{
			m.ip = 0x2c17
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x2c17:
		{
			m.ip = 0x2c1a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x2c1a:
		{
			m.ip = 0x2c1c
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[2])
			m.cycles += 4
		}
	case 0x2c1c:
		{
			m.ip = 0x2c1e
			m.r[7] = m.alu("sub", 16, m.r[7], m.r[2])
			m.cycles += 4
		}
	case 0x2c1e:
		{
			m.ip = 0x2c20
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(11270)
			}
			m.cycles += 17
		}
	case 0x2c20:
		{
			m.ip = 0x2c24
			m.wr16(m.r[11], uint16(m.r[3]+0x26), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x26)), m.r[2]))
			m.cycles += 16
		}
	case 0x2c24:
		{
			m.ip = 0x2c28
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x4e2))
			m.cycles += 8
		}
	case 0x2c28:
		{
			m.ip = 0x2c2c
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x26)), m.r[0])
			m.cycles += 16
		}
	case 0x2c2c:
		{
			m.ip = 0x2c2e
			if !m.zf {
				m.ip = uint16(11316)
			}
			m.cycles += 8
		}
	case 0x2c2e:
		{
			m.ip = 0x2c34
			m.wr16(m.r[11], uint16(m.r[3]+0x26), uint16(0))
			m.cycles += 8
		}
	case 0x2c34:
		{
			m.ip = 0x2c37
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x2c37:
		{
			m.ip = 0x2c3a
			m.alu("sub", 16, m.r[3], uint16(10))
			m.cycles += 4
		}
	case 0x2c3a:
		{
			m.ip = 0x2c3c
			if !m.zf {
				m.ip = uint16(11250)
			}
			m.cycles += 8
		}
	case 0x2c3c:
		{
			m.ip = 0x2c3d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2c3d:
		{
			m.ip = 0x2c43
			m.wr16(m.r[11], uint16(0x5c), uint16(68))
			m.cycles += 8
		}
	case 0x2c43:
		{
			m.ip = 0x2c46
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x2c46:
		{
			m.ip = 0x2c4b
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5)), uint16(0))
			m.cycles += 16
		}
	case 0x2c4b:
		{
			m.ip = 0x2c4d
			if m.zf {
				m.ip = uint16(11396)
			}
			m.cycles += 8
		}
	case 0x2c4d:
		{
			m.ip = 0x2c52
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5)), uint16(8))
			m.cycles += 16
		}
	case 0x2c52:
		{
			m.ip = 0x2c54
			if m.cf {
				m.ip = uint16(11383)
			}
			m.cycles += 8
		}
	case 0x2c54:
		{
			m.ip = 0x2c59
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5)), uint16(46))
			m.cycles += 16
		}
	case 0x2c59:
		{
			m.ip = 0x2c5b
			if !m.cf && !m.zf {
				m.ip = uint16(11396)
			}
			m.cycles += 8
		}
	case 0x2c5b:
		{
			m.ip = 0x2c61
			m.alu("and", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5)), uint16(1))
			m.cycles += 16
		}
	case 0x2c61:
		{
			m.ip = 0x2c63
			if m.zf {
				m.ip = uint16(11470)
			}
			m.cycles += 8
		}
	case 0x2c63:
		{
			m.ip = 0x2c65
			if !m.zf {
				m.ip = uint16(11374)
			}
			m.cycles += 8
		}
	case 0x2c65:
		{
			m.ip = 0x2c6b
			m.wr16(m.r[11], uint16(0x5c), uint16(78))
			m.cycles += 8
		}
	case 0x2c6b:
		{
			m.ip = 0x2c6d
			m.ip = uint16(11396)
			m.cycles += 15
		}
	case 0x2c6d:
		{
			m.ip = 0x2c6e
			m.cycles += 3
		}
	case 0x2c6e:
		{
			m.ip = 0x2c74
			m.wr16(m.r[11], uint16(0x5c), uint16(85))
			m.cycles += 8
		}
	case 0x2c74:
		{
			m.ip = 0x2c76
			m.ip = uint16(11396)
			m.cycles += 15
		}
	case 0x2c76:
		{
			m.ip = 0x2c77
			m.cycles += 3
		}
	case 0x2c77:
		{
			m.ip = 0x2c7c
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x5c)), uint16(85))
			m.cycles += 16
		}
	case 0x2c7c:
		{
			m.ip = 0x2c7e
			if m.zf {
				m.ip = uint16(11396)
			}
			m.cycles += 8
		}
	case 0x2c7e:
		{
			m.ip = 0x2c84
			m.wr16(m.r[11], uint16(0x5c), uint16(68))
			m.cycles += 8
		}
	case 0x2c84:
		{
			m.ip = 0x2c87
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x2c87:
		{
			m.ip = 0x2c8a
			m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x2c8a:
		{
			m.ip = 0x2c8c
			if !m.zf {
				m.ip = uint16(11334)
			}
			m.cycles += 8
		}
	case 0x2c8c:
		{
			m.ip = 0x2c90
			m.r[3] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x2c90:
		{
			m.ip = 0x2c95
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x5c)), uint16(85))
			m.cycles += 16
		}
	case 0x2c95:
		{
			m.ip = 0x2c97
			if !m.zf {
				m.ip = uint16(11443)
			}
			m.cycles += 8
		}
	case 0x2c97:
		{
			m.ip = 0x2c9c
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x5a)), uint16(32))
			m.cycles += 16
		}
	case 0x2c9c:
		{
			m.ip = 0x2c9e
			if m.zf {
				m.ip = uint16(11470)
			}
			m.cycles += 8
		}
	case 0x2c9e:
		{
			m.ip = 0x2ca3
			m.wr16(m.r[11], uint16(0x5a), m.alu("add", 16, m.rd16(m.r[11], uint16(0x5a)), uint16(4)))
			m.cycles += 16
		}
	case 0x2ca3:
		{
			m.ip = 0x2ca7
			m.r[6] = m.rd16(m.r[11], uint16(0x4e))
			m.cycles += 8
		}
	case 0x2ca7:
		{
			m.ip = 0x2cab
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x5a)))
			m.cycles += 16
		}
	case 0x2cab:
		{
			m.ip = 0x2caf
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x50))
			m.cycles += 8
		}
	case 0x2caf:
		{
			m.ip = 0x2cb2
			target := uint16(10677)
			m.push(0x2cb2)
			m.ip = target
			m.cycles += 19
		}
	case 0x2cb2:
		{
			m.ip = 0x2cb3
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2cb3:
		{
			m.ip = 0x2cb8
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x5a)), uint16(65532))
			m.cycles += 16
		}
	case 0x2cb8:
		{
			m.ip = 0x2cba
			if m.zf {
				m.ip = uint16(11470)
			}
			m.cycles += 8
		}
	case 0x2cba:
		{
			m.ip = 0x2cbe
			m.r[6] = m.rd16(m.r[11], uint16(0x4e))
			m.cycles += 8
		}
	case 0x2cbe:
		{
			m.ip = 0x2cc2
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x5a)))
			m.cycles += 16
		}
	case 0x2cc2:
		{
			m.ip = 0x2cc6
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x50))
			m.cycles += 8
		}
	case 0x2cc6:
		{
			m.ip = 0x2cc9
			target := uint16(10677)
			m.push(0x2cc9)
			m.ip = target
			m.cycles += 19
		}
	case 0x2cc9:
		{
			m.ip = 0x2cce
			m.wr16(m.r[11], uint16(0x5a), m.alu("sub", 16, m.rd16(m.r[11], uint16(0x5a)), uint16(4)))
			m.cycles += 16
		}
	case 0x2cce:
		{
			m.ip = 0x2ccf
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2ccf:
		{
			m.ip = 0x2cd4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x37)), uint16(0))
			m.cycles += 16
		}
	case 0x2cd4:
		{
			m.ip = 0x2cd6
			if !m.zf {
				m.ip = uint16(11481)
			}
			m.cycles += 8
		}
	case 0x2cd6:
		{
			m.ip = 0x2cd9
			m.ip = uint16(11711)
			m.cycles += 15
		}
	case 0x2cd9:
		{
			m.ip = 0x2cdd
			m.wr16(m.r[11], uint16(0x38), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x38))))
			m.cycles += 15
		}
	case 0x2cdd:
		{
			m.ip = 0x2cdf
			if !m.zf {
				m.ip = uint16(11490)
			}
			m.cycles += 8
		}
	case 0x2cdf:
		{
			m.ip = 0x2ce2
			m.ip = uint16(11657)
			m.cycles += 15
		}
	case 0x2ce2:
		{
			m.ip = 0x2ce6
			m.r[5] = m.rd16(m.r[11], uint16(0x10))
			m.cycles += 8
		}
	case 0x2ce6:
		{
			m.ip = 0x2ce9
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x2ce9:
		{
			m.ip = 0x2cef
			m.wr16(m.r[11], uint16(0x223), uint16(0))
			m.cycles += 8
		}
	case 0x2cef:
		{
			m.ip = 0x2cf5
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x2cf5:
		{
			m.ip = 0x2cf9
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(m.r[7]+0x1a5)))
			m.cycles += 16
		}
	case 0x2cf9:
		{
			m.ip = 0x2cfb
			if m.zf {
				m.ip = uint16(11546)
			}
			m.cycles += 8
		}
	case 0x2cfb:
		{
			m.ip = 0x2d00
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x2d00:
		{
			m.ip = 0x2d02
			if m.zf {
				m.ip = uint16(11525)
			}
			m.cycles += 8
		}
	case 0x2d02:
		{
			m.ip = 0x2d05
			m.ip = uint16(11656)
			m.cycles += 15
		}
	case 0x2d05:
		{
			m.ip = 0x2d08
			m.r[7] = uint16(2)
			m.cycles += 2
		}
	case 0x2d08:
		{
			m.ip = 0x2d0e
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x2d0e:
		{
			m.ip = 0x2d14
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x2d14:
		{
			m.ip = 0x2d18
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(m.r[7]+0x1a5)))
			m.cycles += 16
		}
	case 0x2d18:
		{
			m.ip = 0x2d1a
			if !m.zf {
				m.ip = uint16(11656)
			}
			m.cycles += 8
		}
	case 0x2d1a:
		{
			m.ip = 0x2d1f
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x37)), uint16(1))
			m.cycles += 16
		}
	case 0x2d1f:
		{
			m.ip = 0x2d21
			if !m.zf {
				m.ip = uint16(11656)
			}
			m.cycles += 8
		}
	case 0x2d21:
		{
			m.ip = 0x2d26
			m.wr8(m.r[11], uint16(0x37), uint16(2))
			m.cycles += 8
		}
	case 0x2d26:
		{
			m.ip = 0x2d2c
			m.wr16(m.r[11], uint16(0x38), uint16(50))
			m.cycles += 8
		}
	case 0x2d2c:
		{
			m.ip = 0x2d2f
			target := uint16(19095)
			m.push(0x2d2f)
			m.ip = target
			m.cycles += 19
		}
	case 0x2d2f:
		{
			m.ip = 0x2d31
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2d31:
		{
			m.ip = 0x2d34
			m.r[0] = m.alu("and", 16, m.r[0], uint16(6))
			m.cycles += 4
		}
	case 0x2d34:
		{
			m.ip = 0x2d36
			m.r[6] = m.r[0]
			m.cycles += 2
		}
	case 0x2d36:
		{
			m.ip = 0x2d3a
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]+0x520))
			m.cycles += 8
		}
	case 0x2d3a:
		{
			m.ip = 0x2d3e
			m.r[3] = m.rd16(m.r[11], uint16(m.r[6]+0x528))
			m.cycles += 8
		}
	case 0x2d3e:
		{
			m.ip = 0x2d42
			m.r[6] = m.rd16(m.r[11], uint16(0x3c))
			m.cycles += 8
		}
	case 0x2d42:
		{
			m.ip = 0x2d46
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(m.r[6]+0x500)))
			m.cycles += 16
		}
	case 0x2d46:
		{
			m.ip = 0x2d4a
			m.r[3] = m.alu("add", 16, m.r[3], m.rd16(m.r[11], uint16(m.r[6]+0x510)))
			m.cycles += 16
		}
	case 0x2d4a:
		{
			m.ip = 0x2d4d
			m.wr16(m.r[11], uint16(0x12), m.r[0])
			m.cycles += 8
		}
	case 0x2d4d:
		{
			m.ip = 0x2d51
			m.wr16(m.r[11], uint16(m.r[7]+0x233), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[7]+0x233)), m.r[3]))
			m.cycles += 16
		}
	case 0x2d51:
		{
			m.ip = 0x2d54
			target := uint16(17038)
			m.push(0x2d54)
			m.ip = target
			m.cycles += 19
		}
	case 0x2d54:
		{
			m.ip = 0x2d58
			m.r[6] = m.rd16(m.r[11], uint16(0x12))
			m.cycles += 8
		}
	case 0x2d58:
		{
			m.ip = 0x2d5c
			m.wr16(m.r[11], uint16(0x21f), m.r[6])
			m.cycles += 8
		}
	case 0x2d5c:
		{
			m.ip = 0x2d60
			m.r[7] = m.rd16(m.r[11], uint16(0x3a))
			m.cycles += 8
		}
	case 0x2d60:
		{
			m.ip = 0x2d63
			target := uint16(10677)
			m.push(0x2d63)
			m.ip = target
			m.cycles += 19
		}
	case 0x2d63:
		{
			m.ip = 0x2d69
			m.wr16(m.r[11], uint16(0x55b), uint16(0))
			m.cycles += 8
		}
	case 0x2d69:
		{
			m.ip = 0x2d6d
			m.r[0] = uint16(0x583)
			m.cycles += 3
		}
	case 0x2d6d:
		{
			m.ip = 0x2d70
			m.wr16(m.r[11], uint16(0x55d), m.r[0])
			m.cycles += 8
		}
	case 0x2d70:
		{
			m.ip = 0x2d74
			m.r[7] = uint16(0x266)
			m.cycles += 3
		}
	case 0x2d74:
		{
			m.ip = 0x2d79
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(2))
			m.cycles += 16
		}
	case 0x2d79:
		{
			m.ip = 0x2d7b
			if m.zf {
				m.ip = uint16(11654)
			}
			m.cycles += 8
		}
	case 0x2d7b:
		{
			m.ip = 0x2d80
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x223)), uint16(1))
			m.cycles += 16
		}
	case 0x2d80:
		{
			m.ip = 0x2d82
			if m.zf {
				m.ip = uint16(11654)
			}
			m.cycles += 8
		}
	case 0x2d82:
		{
			m.ip = 0x2d86
			m.r[7] = uint16(0x25d)
			m.cycles += 3
		}
	case 0x2d86:
		{
			m.ip = 0x2d88
			m.wr16(m.r[11], uint16(m.r[7]), m.unary("inc", 16, m.rd16(m.r[11], uint16(m.r[7]))))
			m.cycles += 15
		}
	case 0x2d88:
		{
			m.ip = 0x2d89
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2d89:
		{
			m.ip = 0x2d8e
			m.wr8(m.r[11], uint16(0x37), uint16(0))
			m.cycles += 8
		}
	case 0x2d8e:
		{
			m.ip = 0x2d91
			target := uint16(19095)
			m.push(0x2d91)
			m.ip = target
			m.cycles += 19
		}
	case 0x2d91:
		{
			m.ip = 0x2d93
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2d93:
		{
			m.ip = 0x2d96
			m.r[0] = m.alu("and", 16, m.r[0], uint16(127))
			m.cycles += 4
		}
	case 0x2d96:
		{
			m.ip = 0x2d99
			m.r[0] = m.alu("add", 16, m.r[0], uint16(336))
			m.cycles += 4
		}
	case 0x2d99:
		{
			m.ip = 0x2d9c
			m.wr16(m.r[11], uint16(0x38), m.r[0])
			m.cycles += 8
		}
	case 0x2d9c:
		{
			m.ip = 0x2d9f
			m.r[6] = uint16(28064)
			m.cycles += 2
		}
	case 0x2d9f:
		{
			m.ip = 0x2da3
			m.r[3] = m.rd16(m.r[11], uint16(0x10))
			m.cycles += 8
		}
	case 0x2da3:
		{
			m.ip = 0x2da8
			m.wr8(m.r[11], uint16(m.r[3]+0x8e), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(251)))
			m.cycles += 16
		}
	case 0x2da8:
		{
			m.ip = 0x2dad
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x2dad:
		{
			m.ip = 0x2daf
			if m.zf {
				m.ip = uint16(11698)
			}
			m.cycles += 8
		}
	case 0x2daf:
		{
			m.ip = 0x2db2
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x2db2:
		{
			m.ip = 0x2db6
			m.r[7] = m.rd16(m.r[11], uint16(0x3a))
			m.cycles += 8
		}
	case 0x2db6:
		{
			m.ip = 0x2db9
			target := uint16(10677)
			m.push(0x2db9)
			m.ip = target
			m.cycles += 19
		}
	case 0x2db9:
		{
			m.ip = 0x2dbf
			m.wr16(m.r[11], uint16(0x10), uint16(0))
			m.cycles += 8
		}
	case 0x2dbf:
		{
			m.ip = 0x2dc3
			m.wr16(m.r[11], uint16(0x38), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x38))))
			m.cycles += 15
		}
	case 0x2dc3:
		{
			m.ip = 0x2dc5
			if m.zf {
				m.ip = uint16(11718)
			}
			m.cycles += 8
		}
	case 0x2dc5:
		{
			m.ip = 0x2dc6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2dc6:
		{
			m.ip = 0x2dcb
			m.wr8(m.r[11], uint16(0x37), uint16(1))
			m.cycles += 8
		}
	case 0x2dcb:
		{
			m.ip = 0x2dce
			target := uint16(19095)
			m.push(0x2dce)
			m.ip = target
			m.cycles += 19
		}
	case 0x2dce:
		{
			m.ip = 0x2dd0
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2dd0:
		{
			m.ip = 0x2dd3
			m.r[0] = m.alu("and", 16, m.r[0], uint16(63))
			m.cycles += 4
		}
	case 0x2dd3:
		{
			m.ip = 0x2dd6
			m.r[0] = m.alu("add", 16, m.r[0], uint16(168))
			m.cycles += 4
		}
	case 0x2dd6:
		{
			m.ip = 0x2dd9
			m.wr16(m.r[11], uint16(0x38), m.r[0])
			m.cycles += 8
		}
	case 0x2dd9:
		{
			m.ip = 0x2ddc
			target := uint16(19095)
			m.push(0x2ddc)
			m.ip = target
			m.cycles += 19
		}
	case 0x2ddc:
		{
			m.ip = 0x2dde
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2dde:
		{
			m.ip = 0x2de1
			m.r[0] = m.alu("and", 16, m.r[0], uint16(14))
			m.cycles += 4
		}
	case 0x2de1:
		{
			m.ip = 0x2de4
			m.wr16(m.r[11], uint16(0x3c), m.r[0])
			m.cycles += 8
		}
	case 0x2de4:
		{
			m.ip = 0x2de6
			m.r[6] = m.r[0]
			m.cycles += 2
		}
	case 0x2de6:
		{
			m.ip = 0x2dea
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]+0x4f0))
			m.cycles += 8
		}
	case 0x2dea:
		{
			m.ip = 0x2dee
			m.r[6] = m.rd16(m.r[11], uint16(0x468))
			m.cycles += 8
		}
	case 0x2dee:
		{
			m.ip = 0x2def
			m.r[6] = m.unary("dec", 16, m.r[6])
			m.cycles += 3
		}
	case 0x2def:
		{
			m.ip = 0x2df1
			m.r[6] = m.shift("shl", 16, m.r[6], uint16(1))
			m.cycles += 8
		}
	case 0x2df1:
		{
			m.ip = 0x2df5
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(m.r[6]+0x520)))
			m.cycles += 16
		}
	case 0x2df5:
		{
			m.ip = 0x2df8
			m.wr16(m.r[11], uint16(0x12), m.r[0])
			m.cycles += 8
		}
	case 0x2df8:
		{
			m.ip = 0x2dfb
			target := uint16(19095)
			m.push(0x2dfb)
			m.ip = target
			m.cycles += 19
		}
	case 0x2dfb:
		{
			m.ip = 0x2dfd
			m.set8(3, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2dfd:
		{
			m.ip = 0x2dff
			m.set8(3, 8, uint16(0))
			m.cycles += 2
		}
	case 0x2dff:
		{
			m.ip = 0x2e02
			m.r[3] = m.alu("add", 16, m.r[3], uint16(11))
			m.cycles += 4
		}
	case 0x2e02:
		{
			m.ip = 0x2e07
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(1))
			m.cycles += 16
		}
	case 0x2e07:
		{
			m.ip = 0x2e09
			if m.zf {
				m.ip = uint16(11768)
			}
			m.cycles += 8
		}
	case 0x2e09:
		{
			m.ip = 0x2e0e
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(2))
			m.cycles += 16
		}
	case 0x2e0e:
		{
			m.ip = 0x2e10
			if m.zf {
				m.ip = uint16(11794)
			}
			m.cycles += 8
		}
	case 0x2e10:
		{
			m.ip = 0x2e12
			m.ip = uint16(11768)
			m.cycles += 15
		}
	case 0x2e12:
		{
			m.ip = 0x2e16
			m.wr16(m.r[11], uint16(0x10), m.r[3])
			m.cycles += 8
		}
	case 0x2e16:
		{
			m.ip = 0x2e1b
			m.wr8(m.r[11], uint16(m.r[3]+0x8e), m.alu("or", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(4)))
			m.cycles += 16
		}
	case 0x2e1b:
		{
			m.ip = 0x2e1d
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2e1d:
		{
			m.ip = 0x2e21
			m.r[6] = m.rd16(m.r[11], uint16(0x12))
			m.cycles += 8
		}
	case 0x2e21:
		{
			m.ip = 0x2e25
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x61c))
			m.cycles += 8
		}
	case 0x2e25:
		{
			m.ip = 0x2e29
			m.wr16(m.r[11], uint16(0x3a), m.r[7])
			m.cycles += 8
		}
	case 0x2e29:
		{
			m.ip = 0x2e2c
			target := uint16(10677)
			m.push(0x2e2c)
			m.ip = target
			m.cycles += 19
		}
	case 0x2e2c:
		{
			m.ip = 0x2e2d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2e2d:
		{
			m.ip = 0x2e30
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x2e30:
		{
			m.ip = 0x2e35
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x7e)), uint16(0))
			m.cycles += 16
		}
	case 0x2e35:
		{
			m.ip = 0x2e37
			if m.zf {
				m.ip = uint16(11875)
			}
			m.cycles += 8
		}
	case 0x2e37:
		{
			m.ip = 0x2e3b
			m.wr16(m.r[11], uint16(m.r[3]+0x7e), m.unary("dec", 16, m.rd16(m.r[11], uint16(m.r[3]+0x7e))))
			m.cycles += 15
		}
	case 0x2e3b:
		{
			m.ip = 0x2e40
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x7e)), uint16(50))
			m.cycles += 16
		}
	case 0x2e40:
		{
			m.ip = 0x2e42
			if !m.zf {
				m.ip = uint16(11852)
			}
			m.cycles += 8
		}
	case 0x2e42:
		{
			m.ip = 0x2e46
			m.r[6] = uint16(0x1001)
			m.cycles += 3
		}
	case 0x2e46:
		{
			m.ip = 0x2e49
			target := uint16(9346)
			m.push(0x2e49)
			m.ip = target
			m.cycles += 19
		}
	case 0x2e49:
		{
			m.ip = 0x2e4b
			m.ip = uint16(11875)
			m.cycles += 15
		}
	case 0x2e4b:
		{
			m.ip = 0x2e4c
			m.cycles += 3
		}
	case 0x2e4c:
		{
			m.ip = 0x2e51
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x7e)), uint16(0))
			m.cycles += 16
		}
	case 0x2e51:
		{
			m.ip = 0x2e53
			if !m.zf {
				m.ip = uint16(11875)
			}
			m.cycles += 8
		}
	case 0x2e53:
		{
			m.ip = 0x2e55
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2e55:
		{
			m.ip = 0x2e5a
			m.wr8(m.r[11], uint16(m.r[3]+0x7a), uint16(71))
			m.cycles += 8
		}
	case 0x2e5a:
		{
			m.ip = 0x2e5c
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2e5c:
		{
			m.ip = 0x2e60
			m.r[6] = uint16(0xff5)
			m.cycles += 3
		}
	case 0x2e60:
		{
			m.ip = 0x2e63
			target := uint16(9346)
			m.push(0x2e63)
			m.ip = target
			m.cycles += 19
		}
	case 0x2e63:
		{
			m.ip = 0x2e68
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x86)), uint16(0))
			m.cycles += 16
		}
	case 0x2e68:
		{
			m.ip = 0x2e6a
			if m.zf {
				m.ip = uint16(11940)
			}
			m.cycles += 8
		}
	case 0x2e6a:
		{
			m.ip = 0x2e6e
			m.wr16(m.r[11], uint16(m.r[3]+0x86), m.unary("dec", 16, m.rd16(m.r[11], uint16(m.r[3]+0x86))))
			m.cycles += 15
		}
	case 0x2e6e:
		{
			m.ip = 0x2e73
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x86)), uint16(0))
			m.cycles += 16
		}
	case 0x2e73:
		{
			m.ip = 0x2e75
			if !m.zf {
				m.ip = uint16(11940)
			}
			m.cycles += 8
		}
	case 0x2e75:
		{
			m.ip = 0x2e77
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2e77:
		{
			m.ip = 0x2e7b
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x0))
			m.cycles += 8
		}
	case 0x2e7b:
		{
			m.ip = 0x2e81
			m.wr16(m.r[11], uint16(m.r[3]+0x0), uint16(0))
			m.cycles += 8
		}
	case 0x2e81:
		{
			m.ip = 0x2e87
			m.wr16(m.r[11], uint16(m.r[3]+0x2), uint16(0))
			m.cycles += 8
		}
	case 0x2e87:
		{
			m.ip = 0x2e89
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2e89:
		{
			m.ip = 0x2e8c
			m.r[6] = uint16(28064)
			m.cycles += 2
		}
	case 0x2e8c:
		{
			m.ip = 0x2e91
			m.wr8(m.r[11], uint16(m.r[7]+0x8e), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[7]+0x8e)), uint16(251)))
			m.cycles += 16
		}
	case 0x2e91:
		{
			m.ip = 0x2e96
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[7]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x2e96:
		{
			m.ip = 0x2e98
			if m.zf {
				m.ip = uint16(11931)
			}
			m.cycles += 8
		}
	case 0x2e98:
		{
			m.ip = 0x2e9b
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x2e9b:
		{
			m.ip = 0x2e9d
			m.r[7] = m.shift("shl", 16, m.r[7], uint16(1))
			m.cycles += 8
		}
	case 0x2e9d:
		{
			m.ip = 0x2ea1
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x61c))
			m.cycles += 8
		}
	case 0x2ea1:
		{
			m.ip = 0x2ea4
			target := uint16(10677)
			m.push(0x2ea4)
			m.ip = target
			m.cycles += 19
		}
	case 0x2ea4:
		{
			m.ip = 0x2ea7
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x2ea7:
		{
			m.ip = 0x2eaa
			m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x2eaa:
		{
			m.ip = 0x2eac
			if !m.zf {
				m.ip = uint16(11824)
			}
			m.cycles += 8
		}
	case 0x2eac:
		{
			m.ip = 0x2eaf
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x2eaf:
		{
			m.ip = 0x2eb5
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x5e)), uint16(255))
			m.cycles += 16
		}
	case 0x2eb5:
		{
			m.ip = 0x2eb7
			if !m.zf {
				m.ip = uint16(11962)
			}
			m.cycles += 8
		}
	case 0x2eb7:
		{
			m.ip = 0x2eba
			m.ip = uint16(12108)
			m.cycles += 15
		}
	case 0x2eba:
		{
			m.ip = 0x2ebe
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x1dd))
			m.cycles += 8
		}
	case 0x2ebe:
		{
			m.ip = 0x2ec2
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(0x1a5)))
			m.cycles += 16
		}
	case 0x2ec2:
		{
			m.ip = 0x2ec4
			if m.zf {
				m.ip = uint16(11985)
			}
			m.cycles += 8
		}
	case 0x2ec4:
		{
			m.ip = 0x2ec9
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x2ec9:
		{
			m.ip = 0x2ecb
			if !m.zf {
				m.ip = uint16(12077)
			}
			m.cycles += 8
		}
	case 0x2ecb:
		{
			m.ip = 0x2ecf
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(0x1a7)))
			m.cycles += 16
		}
	case 0x2ecf:
		{
			m.ip = 0x2ed1
			if !m.zf {
				m.ip = uint16(12077)
			}
			m.cycles += 8
		}
	case 0x2ed1:
		{
			m.ip = 0x2ed7
			m.wr8(m.r[11], uint16(m.r[5]+0x8e), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(253)))
			m.cycles += 16
		}
	case 0x2ed7:
		{
			m.ip = 0x2edd
			m.wr16(m.r[11], uint16(0x55b), uint16(0))
			m.cycles += 8
		}
	case 0x2edd:
		{
			m.ip = 0x2ee1
			m.r[0] = uint16(0x56f)
			m.cycles += 3
		}
	case 0x2ee1:
		{
			m.ip = 0x2ee4
			m.wr16(m.r[11], uint16(0x55d), m.r[0])
			m.cycles += 8
		}
	case 0x2ee4:
		{
			m.ip = 0x2eea
			m.wr16(m.r[11], uint16(m.r[3]+0x5e), uint16(255))
			m.cycles += 8
		}
	case 0x2eea:
		{
			m.ip = 0x2ef0
			m.wr16(m.r[11], uint16(0x66), uint16(0))
			m.cycles += 8
		}
	case 0x2ef0:
		{
			m.ip = 0x2ef3
			m.r[5] = uint16(0)
			m.cycles += 2
		}
	case 0x2ef3:
		{
			m.ip = 0x2ef9
			m.wr8(m.r[11], uint16(m.r[5]+0x7a), uint16(77))
			m.cycles += 8
		}
	case 0x2ef9:
		{
			m.ip = 0x2efc
			target := uint16(19095)
			m.push(0x2efc)
			m.ip = target
			m.cycles += 19
		}
	case 0x2efc:
		{
			m.ip = 0x2efe
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2efe:
		{
			m.ip = 0x2f01
			m.r[0] = m.alu("and", 16, m.r[0], uint16(127))
			m.cycles += 4
		}
	case 0x2f01:
		{
			m.ip = 0x2f04
			m.r[0] = m.alu("add", 16, m.r[0], uint16(235))
			m.cycles += 4
		}
	case 0x2f04:
		{
			m.ip = 0x2f06
			m.r[5] = m.shift("shl", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x2f06:
		{
			m.ip = 0x2f0b
			m.wr16(m.r[11], uint16(m.r[5]+0x7e), m.r[0])
			m.cycles += 8
		}
	case 0x2f0b:
		{
			m.ip = 0x2f0d
			m.r[5] = m.shift("shr", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x2f0d:
		{
			m.ip = 0x2f0e
			m.r[5] = m.unary("inc", 16, m.r[5])
			m.cycles += 3
		}
	case 0x2f0e:
		{
			m.ip = 0x2f11
			m.alu("sub", 16, m.r[5], uint16(4))
			m.cycles += 4
		}
	case 0x2f11:
		{
			m.ip = 0x2f13
			if !m.zf {
				m.ip = uint16(12019)
			}
			m.cycles += 8
		}
	case 0x2f13:
		{
			m.ip = 0x2f16
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x2f16:
		{
			m.ip = 0x2f1a
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1e5))
			m.cycles += 8
		}
	case 0x2f1a:
		{
			m.ip = 0x2f1d
			target := uint16(10677)
			m.push(0x2f1d)
			m.ip = target
			m.cycles += 19
		}
	case 0x2f1d:
		{
			m.ip = 0x2f21
			m.r[6] = uint16(0xffb)
			m.cycles += 3
		}
	case 0x2f21:
		{
			m.ip = 0x2f24
			m.r[3] = uint16(65535)
			m.cycles += 2
		}
	case 0x2f24:
		{
			m.ip = 0x2f27
			target := uint16(9346)
			m.push(0x2f27)
			m.ip = target
			m.cycles += 19
		}
	case 0x2f27:
		{
			m.ip = 0x2f2a
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x2f2a:
		{
			m.ip = 0x2f2c
			m.ip = uint16(12108)
			m.cycles += 15
		}
	case 0x2f2c:
		{
			m.ip = 0x2f2d
			m.cycles += 3
		}
	case 0x2f2d:
		{
			m.ip = 0x2f33
			m.wr8(m.r[11], uint16(m.r[5]+0x8e), m.alu("or", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(2)))
			m.cycles += 16
		}
	case 0x2f33:
		{
			m.ip = 0x2f38
			m.wr16(m.r[11], uint16(m.r[3]+0x5e), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x5e)), uint16(4)))
			m.cycles += 16
		}
	case 0x2f38:
		{
			m.ip = 0x2f3d
			m.wr16(m.r[11], uint16(m.r[3]+0x5e), m.alu("and", 16, m.rd16(m.r[11], uint16(m.r[3]+0x5e)), uint16(63)))
			m.cycles += 16
		}
	case 0x2f3d:
		{
			m.ip = 0x2f41
			m.r[6] = m.rd16(m.r[11], uint16(0x530))
			m.cycles += 8
		}
	case 0x2f41:
		{
			m.ip = 0x2f45
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(m.r[3]+0x5e)))
			m.cycles += 16
		}
	case 0x2f45:
		{
			m.ip = 0x2f49
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1e5))
			m.cycles += 8
		}
	case 0x2f49:
		{
			m.ip = 0x2f4c
			target := uint16(10677)
			m.push(0x2f4c)
			m.ip = target
			m.cycles += 19
		}
	case 0x2f4c:
		{
			m.ip = 0x2f4f
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x2f4f:
		{
			m.ip = 0x2f52
			m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x2f52:
		{
			m.ip = 0x2f54
			if m.zf {
				m.ip = uint16(12119)
			}
			m.cycles += 8
		}
	case 0x2f54:
		{
			m.ip = 0x2f57
			m.ip = uint16(11951)
			m.cycles += 15
		}
	case 0x2f57:
		{
			m.ip = 0x2f58
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2f58:
		{
			m.ip = 0x2f5e
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x3e)), uint16(257))
			m.cycles += 16
		}
	case 0x2f5e:
		{
			m.ip = 0x2f60
			if !m.zf {
				m.ip = uint16(12129)
			}
			m.cycles += 8
		}
	case 0x2f60:
		{
			m.ip = 0x2f61
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2f61:
		{
			m.ip = 0x2f66
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x40)), uint16(0))
			m.cycles += 16
		}
	case 0x2f66:
		{
			m.ip = 0x2f68
			if !m.zf {
				m.ip = uint16(12139)
			}
			m.cycles += 8
		}
	case 0x2f68:
		{
			m.ip = 0x2f6b
			m.ip = uint16(12347)
			m.cycles += 15
		}
	case 0x2f6b:
		{
			m.ip = 0x2f6f
			m.wr16(m.r[11], uint16(0x41), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x41))))
			m.cycles += 15
		}
	case 0x2f6f:
		{
			m.ip = 0x2f71
			if !m.zf {
				m.ip = uint16(12148)
			}
			m.cycles += 8
		}
	case 0x2f71:
		{
			m.ip = 0x2f74
			m.ip = uint16(12292)
			m.cycles += 15
		}
	case 0x2f74:
		{
			m.ip = 0x2f78
			m.r[5] = m.rd16(m.r[11], uint16(0x14))
			m.cycles += 8
		}
	case 0x2f78:
		{
			m.ip = 0x2f7e
			m.wr16(m.r[11], uint16(0x223), uint16(0))
			m.cycles += 8
		}
	case 0x2f7e:
		{
			m.ip = 0x2f84
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x2f84:
		{
			m.ip = 0x2f88
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(0x1a5)))
			m.cycles += 16
		}
	case 0x2f88:
		{
			m.ip = 0x2f8a
			if m.zf {
				m.ip = uint16(12195)
			}
			m.cycles += 8
		}
	case 0x2f8a:
		{
			m.ip = 0x2f8f
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x2f8f:
		{
			m.ip = 0x2f91
			if !m.zf {
				m.ip = uint16(12291)
			}
			m.cycles += 8
		}
	case 0x2f91:
		{
			m.ip = 0x2f97
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x2f97:
		{
			m.ip = 0x2f9d
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x2f9d:
		{
			m.ip = 0x2fa1
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(0x1a7)))
			m.cycles += 16
		}
	case 0x2fa1:
		{
			m.ip = 0x2fa3
			if !m.zf {
				m.ip = uint16(12291)
			}
			m.cycles += 8
		}
	case 0x2fa3:
		{
			m.ip = 0x2fa8
			m.wr8(m.r[11], uint16(0x40), uint16(0))
			m.cycles += 8
		}
	case 0x2fa8:
		{
			m.ip = 0x2fac
			m.r[3] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x2fac:
		{
			m.ip = 0x2fb2
			m.wr16(m.r[11], uint16(0x55b), uint16(0))
			m.cycles += 8
		}
	case 0x2fb2:
		{
			m.ip = 0x2fb6
			m.r[0] = uint16(0x583)
			m.cycles += 3
		}
	case 0x2fb6:
		{
			m.ip = 0x2fb9
			m.wr16(m.r[11], uint16(0x55d), m.r[0])
			m.cycles += 8
		}
	case 0x2fb9:
		{
			m.ip = 0x2fbe
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x3e)), uint16(1))
			m.cycles += 16
		}
	case 0x2fbe:
		{
			m.ip = 0x2fc0
			if !m.zf {
				m.ip = uint16(12280)
			}
			m.cycles += 8
		}
	case 0x2fc0:
		{
			m.ip = 0x2fc4
			m.r[6] = m.rd16(m.r[11], uint16(0x500))
			m.cycles += 8
		}
	case 0x2fc4:
		{
			m.ip = 0x2fc8
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x522)))
			m.cycles += 16
		}
	case 0x2fc8:
		{
			m.ip = 0x2fcc
			m.wr16(m.r[11], uint16(0x16), m.r[6])
			m.cycles += 8
		}
	case 0x2fcc:
		{
			m.ip = 0x2fd0
			m.r[7] = m.rd16(m.r[11], uint16(0x43))
			m.cycles += 8
		}
	case 0x2fd0:
		{
			m.ip = 0x2fd3
			target := uint16(10677)
			m.push(0x2fd3)
			m.ip = target
			m.cycles += 19
		}
	case 0x2fd3:
		{
			m.ip = 0x2fd5
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x2fd5:
		{
			m.ip = 0x2fd8
			m.r[0] = m.rd16(m.r[11], uint16(0x510))
			m.cycles += 8
		}
	case 0x2fd8:
		{
			m.ip = 0x2fdc
			m.r[0] = m.alu("add", 16, m.r[0], m.rd16(m.r[11], uint16(0x52a)))
			m.cycles += 16
		}
	case 0x2fdc:
		{
			m.ip = 0x2fe0
			m.wr16(m.r[11], uint16(m.r[3]+0x233), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x233)), m.r[0]))
			m.cycles += 16
		}
	case 0x2fe0:
		{
			m.ip = 0x2fe3
			target := uint16(17038)
			m.push(0x2fe3)
			m.ip = target
			m.cycles += 19
		}
	case 0x2fe3:
		{
			m.ip = 0x2fe6
			target := uint16(19095)
			m.push(0x2fe6)
			m.ip = target
			m.cycles += 19
		}
	case 0x2fe6:
		{
			m.ip = 0x2fe8
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x2fe8:
		{
			m.ip = 0x2feb
			m.r[0] = m.alu("and", 16, m.r[0], uint16(31))
			m.cycles += 4
		}
	case 0x2feb:
		{
			m.ip = 0x2fee
			m.r[0] = m.alu("add", 16, m.r[0], uint16(50))
			m.cycles += 4
		}
	case 0x2fee:
		{
			m.ip = 0x2ff1
			m.wr16(m.r[11], uint16(0x41), m.r[0])
			m.cycles += 8
		}
	case 0x2ff1:
		{
			m.ip = 0x2ff7
			m.wr16(m.r[11], uint16(0x45), uint16(50))
			m.cycles += 8
		}
	case 0x2ff7:
		{
			m.ip = 0x2ff8
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x2ff8:
		{
			m.ip = 0x2ffd
			m.wr8(m.r[11], uint16(m.r[3]+0x3e), uint16(1))
			m.cycles += 8
		}
	case 0x2ffd:
		{
			m.ip = 0x3000
			target := uint16(12473)
			m.push(0x3000)
			m.ip = target
			m.cycles += 19
		}
	case 0x3000:
		{
			m.ip = 0x3002
			m.ip = uint16(12311)
			m.cycles += 15
		}
	case 0x3002:
		{
			m.ip = 0x3003
			m.cycles += 3
		}
	case 0x3003:
		{
			m.ip = 0x3004
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3004:
		{
			m.ip = 0x3009
			m.wr8(m.r[11], uint16(0x40), uint16(0))
			m.cycles += 8
		}
	case 0x3009:
		{
			m.ip = 0x300c
			target := uint16(19095)
			m.push(0x300c)
			m.ip = target
			m.cycles += 19
		}
	case 0x300c:
		{
			m.ip = 0x300e
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x300e:
		{
			m.ip = 0x3011
			m.r[0] = m.alu("and", 16, m.r[0], uint16(31))
			m.cycles += 4
		}
	case 0x3011:
		{
			m.ip = 0x3014
			m.r[0] = m.alu("add", 16, m.r[0], uint16(50))
			m.cycles += 4
		}
	case 0x3014:
		{
			m.ip = 0x3017
			m.wr16(m.r[11], uint16(0x41), m.r[0])
			m.cycles += 8
		}
	case 0x3017:
		{
			m.ip = 0x301a
			m.r[6] = uint16(28064)
			m.cycles += 2
		}
	case 0x301a:
		{
			m.ip = 0x301e
			m.r[3] = m.rd16(m.r[11], uint16(0x14))
			m.cycles += 8
		}
	case 0x301e:
		{
			m.ip = 0x3023
			m.wr8(m.r[11], uint16(m.r[3]+0x8e), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(251)))
			m.cycles += 16
		}
	case 0x3023:
		{
			m.ip = 0x3028
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x3028:
		{
			m.ip = 0x302a
			if m.zf {
				m.ip = uint16(12333)
			}
			m.cycles += 8
		}
	case 0x302a:
		{
			m.ip = 0x302d
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x302d:
		{
			m.ip = 0x3031
			m.r[7] = m.rd16(m.r[11], uint16(0x43))
			m.cycles += 8
		}
	case 0x3031:
		{
			m.ip = 0x3034
			target := uint16(10677)
			m.push(0x3034)
			m.ip = target
			m.cycles += 19
		}
	case 0x3034:
		{
			m.ip = 0x303a
			m.wr16(m.r[11], uint16(0x14), uint16(0))
			m.cycles += 8
		}
	case 0x303a:
		{
			m.ip = 0x303b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x303b:
		{
			m.ip = 0x3040
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x3040:
		{
			m.ip = 0x3042
			if !m.zf {
				m.ip = uint16(12378)
			}
			m.cycles += 8
		}
	case 0x3042:
		{
			m.ip = 0x3047
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(2))
			m.cycles += 16
		}
	case 0x3047:
		{
			m.ip = 0x3049
			if m.zf {
				m.ip = uint16(12378)
			}
			m.cycles += 8
		}
	case 0x3049:
		{
			m.ip = 0x304b
			m.set8(3, 8, m.alu("xor", 8, ((m.r[3]>>8)&255), ((m.r[3]>>8)&255)))
			m.cycles += 4
		}
	case 0x304b:
		{
			m.ip = 0x304f
			m.set8(3, 0, m.rd8(m.r[11], uint16(0x61a)))
			m.cycles += 8
		}
	case 0x304f:
		{
			m.ip = 0x3052
			m.set8(3, 0, m.alu("xor", 8, ((m.r[3]>>0)&255), uint16(1)))
			m.cycles += 4
		}
	case 0x3052:
		{
			m.ip = 0x3057
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x3e)), uint16(1))
			m.cycles += 16
		}
	case 0x3057:
		{
			m.ip = 0x3059
			if !m.zf {
				m.ip = uint16(12378)
			}
			m.cycles += 8
		}
	case 0x3059:
		{
			m.ip = 0x305a
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x305a:
		{
			m.ip = 0x305e
			m.wr16(m.r[11], uint16(0x41), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x41))))
			m.cycles += 15
		}
	case 0x305e:
		{
			m.ip = 0x3060
			if m.zf {
				m.ip = uint16(12398)
			}
			m.cycles += 8
		}
	case 0x3060:
		{
			m.ip = 0x3065
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x45)), uint16(0))
			m.cycles += 16
		}
	case 0x3065:
		{
			m.ip = 0x3067
			if m.zf {
				m.ip = uint16(12472)
			}
			m.cycles += 8
		}
	case 0x3067:
		{
			m.ip = 0x306b
			m.wr16(m.r[11], uint16(0x45), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x45))))
			m.cycles += 15
		}
	case 0x306b:
		{
			m.ip = 0x306d
			if m.zf {
				m.ip = uint16(12311)
			}
			m.cycles += 8
		}
	case 0x306d:
		{
			m.ip = 0x306e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x306e:
		{
			m.ip = 0x3073
			m.wr8(m.r[11], uint16(0x40), uint16(1))
			m.cycles += 8
		}
	case 0x3073:
		{
			m.ip = 0x3076
			target := uint16(19095)
			m.push(0x3076)
			m.ip = target
			m.cycles += 19
		}
	case 0x3076:
		{
			m.ip = 0x3078
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x3078:
		{
			m.ip = 0x307b
			m.r[0] = m.alu("and", 16, m.r[0], uint16(63))
			m.cycles += 4
		}
	case 0x307b:
		{
			m.ip = 0x307e
			m.r[0] = m.alu("add", 16, m.r[0], uint16(168))
			m.cycles += 4
		}
	case 0x307e:
		{
			m.ip = 0x3081
			m.wr16(m.r[11], uint16(0x41), m.r[0])
			m.cycles += 8
		}
	case 0x3081:
		{
			m.ip = 0x3084
			target := uint16(19095)
			m.push(0x3084)
			m.ip = target
			m.cycles += 19
		}
	case 0x3084:
		{
			m.ip = 0x3086
			m.set8(3, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x3086:
		{
			m.ip = 0x3088
			m.set8(3, 8, uint16(0))
			m.cycles += 2
		}
	case 0x3088:
		{
			m.ip = 0x308b
			m.r[3] = m.alu("add", 16, m.r[3], uint16(11))
			m.cycles += 4
		}
	case 0x308b:
		{
			m.ip = 0x3090
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(1))
			m.cycles += 16
		}
	case 0x3090:
		{
			m.ip = 0x3092
			if m.zf {
				m.ip = uint16(12417)
			}
			m.cycles += 8
		}
	case 0x3092:
		{
			m.ip = 0x3097
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(2))
			m.cycles += 16
		}
	case 0x3097:
		{
			m.ip = 0x3099
			if m.zf {
				m.ip = uint16(12443)
			}
			m.cycles += 8
		}
	case 0x3099:
		{
			m.ip = 0x309b
			m.ip = uint16(12417)
			m.cycles += 15
		}
	case 0x309b:
		{
			m.ip = 0x309f
			m.wr16(m.r[11], uint16(0x14), m.r[3])
			m.cycles += 8
		}
	case 0x309f:
		{
			m.ip = 0x30a4
			m.wr8(m.r[11], uint16(m.r[3]+0x8e), m.alu("or", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(4)))
			m.cycles += 16
		}
	case 0x30a4:
		{
			m.ip = 0x30a6
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x30a6:
		{
			m.ip = 0x30a9
			m.r[6] = uint16(43414)
			m.cycles += 2
		}
	case 0x30a9:
		{
			m.ip = 0x30ad
			m.wr16(m.r[11], uint16(0x16), m.r[6])
			m.cycles += 8
		}
	case 0x30ad:
		{
			m.ip = 0x30b1
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x61c))
			m.cycles += 8
		}
	case 0x30b1:
		{
			m.ip = 0x30b5
			m.wr16(m.r[11], uint16(0x43), m.r[7])
			m.cycles += 8
		}
	case 0x30b5:
		{
			m.ip = 0x30b8
			target := uint16(10677)
			m.push(0x30b8)
			m.ip = target
			m.cycles += 19
		}
	case 0x30b8:
		{
			m.ip = 0x30b9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x30b9:
		{
			m.ip = 0x30be
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x466)), uint16(0))
			m.cycles += 16
		}
	case 0x30be:
		{
			m.ip = 0x30c0
			if m.zf {
				m.ip = uint16(12481)
			}
			m.cycles += 8
		}
	case 0x30c0:
		{
			m.ip = 0x30c1
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x30c1:
		{
			m.ip = 0x30c4
			m.r[7] = uint16(78)
			m.cycles += 2
		}
	case 0x30c4:
		{
			m.ip = 0x30c9
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(1))
			m.cycles += 16
		}
	case 0x30c9:
		{
			m.ip = 0x30cb
			if !m.zf {
				m.ip = uint16(12511)
			}
			m.cycles += 8
		}
	case 0x30cb:
		{
			m.ip = 0x30d0
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x30d0:
		{
			m.ip = 0x30d2
			if !m.zf {
				m.ip = uint16(12521)
			}
			m.cycles += 8
		}
	case 0x30d2:
		{
			m.ip = 0x30d7
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x223)), uint16(1))
			m.cycles += 16
		}
	case 0x30d7:
		{
			m.ip = 0x30d9
			if m.zf {
				m.ip = uint16(12521)
			}
			m.cycles += 8
		}
	case 0x30d9:
		{
			m.ip = 0x30dc
			m.r[7] = uint16(18)
			m.cycles += 2
		}
	case 0x30dc:
		{
			m.ip = 0x30de
			m.ip = uint16(12521)
			m.cycles += 15
		}
	case 0x30de:
		{
			m.ip = 0x30df
			m.cycles += 3
		}
	case 0x30df:
		{
			m.ip = 0x30e4
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(2))
			m.cycles += 16
		}
	case 0x30e4:
		{
			m.ip = 0x30e6
			if m.zf {
				m.ip = uint16(12521)
			}
			m.cycles += 8
		}
	case 0x30e6:
		{
			m.ip = 0x30e9
			m.r[7] = uint16(18)
			m.cycles += 2
		}
	case 0x30e9:
		{
			m.ip = 0x30ec
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x30ec:
		{
			m.ip = 0x30ef
			m.r[6] = uint16(54880)
			m.cycles += 2
		}
	case 0x30ef:
		{
			m.ip = 0x30f2
			m.r[6] = m.alu("add", 16, m.r[6], uint16(63))
			m.cycles += 4
		}
	case 0x30f2:
		{
			m.ip = 0x30f6
			m.r[3] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x30f6:
		{
			m.ip = 0x30fb
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x3e)), uint16(1))
			m.cycles += 16
		}
	case 0x30fb:
		{
			m.ip = 0x30fd
			if !m.zf {
				m.ip = uint16(12544)
			}
			m.cycles += 8
		}
	case 0x30fd:
		{
			m.ip = 0x3100
			m.r[6] = uint16(54933)
			m.cycles += 2
		}
	case 0x3100:
		{
			m.ip = 0x3103
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x3103:
		{
			m.ip = 0x3106
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3106:
		{
			m.ip = 0x310a
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x310a:
		{
			m.ip = 0x310e
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x310e:
		{
			m.ip = 0x3111
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x3111:
		{
			m.ip = 0x3114
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x3114:
		{
			m.ip = 0x3116
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(12544)
			}
			m.cycles += 17
		}
	case 0x3116:
		{
			m.ip = 0x3117
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3117:
		{
			m.ip = 0x311b
			m.r[3] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x311b:
		{
			m.ip = 0x3120
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x3e)), uint16(0))
			m.cycles += 16
		}
	case 0x3120:
		{
			m.ip = 0x3122
			if !m.zf {
				m.ip = uint16(12579)
			}
			m.cycles += 8
		}
	case 0x3122:
		{
			m.ip = 0x3123
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3123:
		{
			m.ip = 0x3128
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x49)), uint16(79))
			m.cycles += 16
		}
	case 0x3128:
		{
			m.ip = 0x312a
			if !m.zf {
				m.ip = uint16(12589)
			}
			m.cycles += 8
		}
	case 0x312a:
		{
			m.ip = 0x312d
			m.ip = uint16(12742)
			m.cycles += 15
		}
	case 0x312d:
		{
			m.ip = 0x3132
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x49)), uint16(67))
			m.cycles += 16
		}
	case 0x3132:
		{
			m.ip = 0x3134
			if !m.zf {
				m.ip = uint16(12599)
			}
			m.cycles += 8
		}
	case 0x3134:
		{
			m.ip = 0x3137
			m.ip = uint16(12819)
			m.cycles += 15
		}
	case 0x3137:
		{
			m.ip = 0x313b
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x313b:
		{
			m.ip = 0x313f
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x1a5))
			m.cycles += 8
		}
	case 0x313f:
		{
			m.ip = 0x3142
			m.r[1] = uint16(5)
			m.cycles += 2
		}
	case 0x3142:
		{
			m.ip = 0x3147
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x48)), uint16(79))
			m.cycles += 16
		}
	case 0x3147:
		{
			m.ip = 0x3149
			if m.zf {
				m.ip = uint16(12640)
			}
			m.cycles += 8
		}
	case 0x3149:
		{
			m.ip = 0x314d
			m.r[6] = uint16(0x546)
			m.cycles += 3
		}
	case 0x314d:
		{
			m.ip = 0x314f
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(m.r[6])))
			m.cycles += 16
		}
	case 0x314f:
		{
			m.ip = 0x3151
			if m.zf {
				m.ip = uint16(12631)
			}
			m.cycles += 8
		}
	case 0x3151:
		{
			m.ip = 0x3154
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x3154:
		{
			m.ip = 0x3156
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(12621)
			}
			m.cycles += 17
		}
	case 0x3156:
		{
			m.ip = 0x3157
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3157:
		{
			m.ip = 0x315d
			m.wr16(m.r[11], uint16(0x4c), uint16(1))
			m.cycles += 8
		}
	case 0x315d:
		{
			m.ip = 0x315f
			m.ip = uint16(12742)
			m.cycles += 15
		}
	case 0x315f:
		{
			m.ip = 0x3160
			m.cycles += 3
		}
	case 0x3160:
		{
			m.ip = 0x3163
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x3163:
		{
			m.ip = 0x3168
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x47)), uint16(1))
			m.cycles += 16
		}
	case 0x3168:
		{
			m.ip = 0x316a
			if m.zf {
				m.ip = uint16(12664)
			}
			m.cycles += 8
		}
	case 0x316a:
		{
			m.ip = 0x316e
			m.r[6] = uint16(0x550)
			m.cycles += 3
		}
	case 0x316e:
		{
			m.ip = 0x3170
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(m.r[6])))
			m.cycles += 16
		}
	case 0x3170:
		{
			m.ip = 0x3172
			if m.zf {
				m.ip = uint16(12664)
			}
			m.cycles += 8
		}
	case 0x3172:
		{
			m.ip = 0x3175
			m.r[6] = m.alu("add", 16, m.r[6], uint16(2))
			m.cycles += 4
		}
	case 0x3175:
		{
			m.ip = 0x3177
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(12654)
			}
			m.cycles += 17
		}
	case 0x3177:
		{
			m.ip = 0x3178
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3178:
		{
			m.ip = 0x317c
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x1bd))
			m.cycles += 8
		}
	case 0x317c:
		{
			m.ip = 0x317f
			m.alu("sub", 16, m.r[0], uint16(54))
			m.cycles += 4
		}
	case 0x317f:
		{
			m.ip = 0x3181
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x3181:
		{
			m.ip = 0x3184
			m.alu("sub", 16, m.r[0], uint16(75))
			m.cycles += 4
		}
	case 0x3184:
		{
			m.ip = 0x3186
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x3186:
		{
			m.ip = 0x3189
			m.alu("sub", 16, m.r[0], uint16(96))
			m.cycles += 4
		}
	case 0x3189:
		{
			m.ip = 0x318b
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x318b:
		{
			m.ip = 0x318e
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x318e:
		{
			m.ip = 0x3191
			m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x3191:
		{
			m.ip = 0x3193
			if !m.zf {
				m.ip = uint16(12664)
			}
			m.cycles += 8
		}
	case 0x3193:
		{
			m.ip = 0x3198
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x1a5)), uint16(54))
			m.cycles += 16
		}
	case 0x3198:
		{
			m.ip = 0x319a
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x319a:
		{
			m.ip = 0x319f
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x1a5)), uint16(75))
			m.cycles += 16
		}
	case 0x319f:
		{
			m.ip = 0x31a1
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x31a1:
		{
			m.ip = 0x31a6
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x1a5)), uint16(96))
			m.cycles += 16
		}
	case 0x31a6:
		{
			m.ip = 0x31a8
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x31a8:
		{
			m.ip = 0x31ad
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x1a7)), uint16(54))
			m.cycles += 16
		}
	case 0x31ad:
		{
			m.ip = 0x31af
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x31af:
		{
			m.ip = 0x31b4
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x1a7)), uint16(75))
			m.cycles += 16
		}
	case 0x31b4:
		{
			m.ip = 0x31b6
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x31b6:
		{
			m.ip = 0x31bb
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x1a7)), uint16(96))
			m.cycles += 16
		}
	case 0x31bb:
		{
			m.ip = 0x31bd
			if m.zf {
				m.ip = uint16(12736)
			}
			m.cycles += 8
		}
	case 0x31bd:
		{
			m.ip = 0x31bf
			m.ip = uint16(12819)
			m.cycles += 15
		}
	case 0x31bf:
		{
			m.ip = 0x31c0
			m.cycles += 3
		}
	case 0x31c0:
		{
			m.ip = 0x31c5
			m.wr8(m.r[11], uint16(0x47), uint16(1))
			m.cycles += 8
		}
	case 0x31c5:
		{
			m.ip = 0x31c6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x31c6:
		{
			m.ip = 0x31ca
			m.wr16(m.r[11], uint16(0x4c), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x4c))))
			m.cycles += 15
		}
	case 0x31ca:
		{
			m.ip = 0x31cc
			if !m.zf {
				m.ip = uint16(12818)
			}
			m.cycles += 8
		}
	case 0x31cc:
		{
			m.ip = 0x31d2
			m.wr16(m.r[11], uint16(0x4c), uint16(2))
			m.cycles += 8
		}
	case 0x31d2:
		{
			m.ip = 0x31d7
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x31d7:
		{
			m.ip = 0x31d9
			if !m.zf {
				m.ip = uint16(12775)
			}
			m.cycles += 8
		}
	case 0x31d9:
		{
			m.ip = 0x31df
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x3e)), uint16(257))
			m.cycles += 16
		}
	case 0x31df:
		{
			m.ip = 0x31e1
			if !m.zf {
				m.ip = uint16(12775)
			}
			m.cycles += 8
		}
	case 0x31e1:
		{
			m.ip = 0x31e7
			m.wr16(m.r[11], uint16(0x4c), uint16(4))
			m.cycles += 8
		}
	case 0x31e7:
		{
			m.ip = 0x31ec
			m.wr8(m.r[11], uint16(0x49), uint16(79))
			m.cycles += 8
		}
	case 0x31ec:
		{
			m.ip = 0x31f1
			m.wr16(m.r[11], uint16(0x4a), m.alu("add", 16, m.rd16(m.r[11], uint16(0x4a)), uint16(6)))
			m.cycles += 16
		}
	case 0x31f1:
		{
			m.ip = 0x31f4
			target := uint16(12901)
			m.push(0x31f4)
			m.ip = target
			m.cycles += 19
		}
	case 0x31f4:
		{
			m.ip = 0x31f9
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x4a)), uint16(48))
			m.cycles += 16
		}
	case 0x31f9:
		{
			m.ip = 0x31fb
			if !m.zf {
				m.ip = uint16(12818)
			}
			m.cycles += 8
		}
	case 0x31fb:
		{
			m.ip = 0x31fe
			m.r[3] = uint16(75)
			m.cycles += 2
		}
	case 0x31fe:
		{
			m.ip = 0x3203
			m.wr8(m.r[11], uint16(m.r[3]+0xa3), m.alu("or", 8, m.rd8(m.r[11], uint16(m.r[3]+0xa3)), uint16(128)))
			m.cycles += 16
		}
	case 0x3203:
		{
			m.ip = 0x3208
			m.wr8(m.r[11], uint16(m.r[3]+0x79), m.alu("or", 8, m.rd8(m.r[11], uint16(m.r[3]+0x79)), uint16(64)))
			m.cycles += 16
		}
	case 0x3208:
		{
			m.ip = 0x320d
			m.wr8(m.r[11], uint16(0x48), uint16(79))
			m.cycles += 8
		}
	case 0x320d:
		{
			m.ip = 0x3212
			m.wr8(m.r[11], uint16(0x49), uint16(78))
			m.cycles += 8
		}
	case 0x3212:
		{
			m.ip = 0x3213
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3213:
		{
			m.ip = 0x3217
			m.wr16(m.r[11], uint16(0x4c), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x4c))))
			m.cycles += 15
		}
	case 0x3217:
		{
			m.ip = 0x3219
			if !m.zf {
				m.ip = uint16(12818)
			}
			m.cycles += 8
		}
	case 0x3219:
		{
			m.ip = 0x321f
			m.wr16(m.r[11], uint16(0x4c), uint16(2))
			m.cycles += 8
		}
	case 0x321f:
		{
			m.ip = 0x3224
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x3224:
		{
			m.ip = 0x3226
			if !m.zf {
				m.ip = uint16(12852)
			}
			m.cycles += 8
		}
	case 0x3226:
		{
			m.ip = 0x322c
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x3e)), uint16(257))
			m.cycles += 16
		}
	case 0x322c:
		{
			m.ip = 0x322e
			if !m.zf {
				m.ip = uint16(12852)
			}
			m.cycles += 8
		}
	case 0x322e:
		{
			m.ip = 0x3234
			m.wr16(m.r[11], uint16(0x4c), uint16(4))
			m.cycles += 8
		}
	case 0x3234:
		{
			m.ip = 0x3239
			m.wr8(m.r[11], uint16(0x49), uint16(67))
			m.cycles += 8
		}
	case 0x3239:
		{
			m.ip = 0x323e
			m.wr8(m.r[11], uint16(0x47), uint16(0))
			m.cycles += 8
		}
	case 0x323e:
		{
			m.ip = 0x3243
			m.wr16(m.r[11], uint16(0x4a), m.alu("sub", 16, m.rd16(m.r[11], uint16(0x4a)), uint16(6)))
			m.cycles += 16
		}
	case 0x3243:
		{
			m.ip = 0x3246
			target := uint16(12901)
			m.push(0x3246)
			m.ip = target
			m.cycles += 19
		}
	case 0x3246:
		{
			m.ip = 0x3249
			m.r[3] = uint16(75)
			m.cycles += 2
		}
	case 0x3249:
		{
			m.ip = 0x324e
			m.wr8(m.r[11], uint16(m.r[3]+0xa3), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0xa3)), uint16(127)))
			m.cycles += 16
		}
	case 0x324e:
		{
			m.ip = 0x3253
			m.wr8(m.r[11], uint16(m.r[3]+0x79), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x79)), uint16(191)))
			m.cycles += 16
		}
	case 0x3253:
		{
			m.ip = 0x3258
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x4a)), uint16(0))
			m.cycles += 16
		}
	case 0x3258:
		{
			m.ip = 0x325a
			if !m.zf {
				m.ip = uint16(12818)
			}
			m.cycles += 8
		}
	case 0x325a:
		{
			m.ip = 0x325f
			m.wr8(m.r[11], uint16(0x48), uint16(67))
			m.cycles += 8
		}
	case 0x325f:
		{
			m.ip = 0x3264
			m.wr8(m.r[11], uint16(0x49), uint16(78))
			m.cycles += 8
		}
	case 0x3264:
		{
			m.ip = 0x3265
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3265:
		{
			m.ip = 0x3269
			m.r[6] = m.rd16(m.r[11], uint16(0x542))
			m.cycles += 8
		}
	case 0x3269:
		{
			m.ip = 0x326d
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(0x4a)))
			m.cycles += 16
		}
	case 0x326d:
		{
			m.ip = 0x3271
			m.r[7] = m.rd16(m.r[11], uint16(0x544))
			m.cycles += 8
		}
	case 0x3271:
		{
			m.ip = 0x3274
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x3274:
		{
			m.ip = 0x3275
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x3275:
		{
			m.ip = 0x3278
			m.r[1] = uint16(6)
			m.cycles += 2
		}
	case 0x3278:
		{
			m.ip = 0x327b
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x327b:
		{
			m.ip = 0x327e
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x327e:
		{
			m.ip = 0x327f
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x327f:
		{
			m.ip = 0x3280
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x3280:
		{
			m.ip = 0x3282
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(12920)
			}
			m.cycles += 17
		}
	case 0x3282:
		{
			m.ip = 0x3285
			m.r[6] = m.alu("add", 16, m.r[6], uint16(74))
			m.cycles += 4
		}
	case 0x3285:
		{
			m.ip = 0x3288
			m.r[7] = m.alu("add", 16, m.r[7], uint16(74))
			m.cycles += 4
		}
	case 0x3288:
		{
			m.ip = 0x3289
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x3289:
		{
			m.ip = 0x328b
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(12916)
			}
			m.cycles += 17
		}
	case 0x328b:
		{
			m.ip = 0x328c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x328c:
		{
			m.ip = 0x3291
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x78)), uint16(0))
			m.cycles += 16
		}
	case 0x3291:
		{
			m.ip = 0x3293
			if m.zf {
				m.ip = uint16(12950)
			}
			m.cycles += 8
		}
	case 0x3293:
		{
			m.ip = 0x3296
			m.ip = uint16(13111)
			m.cycles += 15
		}
	case 0x3296:
		{
			m.ip = 0x329a
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x329a:
		{
			m.ip = 0x329f
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x246)), uint16(0))
			m.cycles += 16
		}
	case 0x329f:
		{
			m.ip = 0x32a1
			if !m.zf {
				m.ip = uint16(12964)
			}
			m.cycles += 8
		}
	case 0x32a1:
		{
			m.ip = 0x32a4
			m.ip = uint16(13111)
			m.cycles += 15
		}
	case 0x32a4:
		{
			m.ip = 0x32a8
			m.r[3] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x32a8:
		{
			m.ip = 0x32ad
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x6a)), uint16(0))
			m.cycles += 16
		}
	case 0x32ad:
		{
			m.ip = 0x32af
			if m.zf {
				m.ip = uint16(13014)
			}
			m.cycles += 8
		}
	case 0x32af:
		{
			m.ip = 0x32b1
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x32b1:
		{
			m.ip = 0x32b5
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x6c))
			m.cycles += 8
		}
	case 0x32b5:
		{
			m.ip = 0x32b7
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x32b7:
		{
			m.ip = 0x32bc
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x68)), uint16(10))
			m.cycles += 16
		}
	case 0x32bc:
		{
			m.ip = 0x32be
			if !m.cf && !m.zf {
				m.ip = uint16(12996)
			}
			m.cycles += 8
		}
	case 0x32be:
		{
			m.ip = 0x32c1
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x32c1:
		{
			m.ip = 0x32c4
			m.r[7] = m.alu("and", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x32c4:
		{
			m.ip = 0x32c7
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x32c7:
		{
			m.ip = 0x32cb
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x1f5))
			m.cycles += 8
		}
	case 0x32cb:
		{
			m.ip = 0x32cf
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x32cf:
		{
			m.ip = 0x32d3
			m.wr16(m.r[11], uint16(m.r[3]+0x1ad), m.r[7])
			m.cycles += 8
		}
	case 0x32d3:
		{
			m.ip = 0x32d5
			m.ip = uint16(13036)
			m.cycles += 15
		}
	case 0x32d5:
		{
			m.ip = 0x32d6
			m.cycles += 3
		}
	case 0x32d6:
		{
			m.ip = 0x32da
			m.r[3] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x32da:
		{
			m.ip = 0x32df
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)), uint16(1))
			m.cycles += 16
		}
	case 0x32df:
		{
			m.ip = 0x32e1
			if !m.zf {
				m.ip = uint16(13036)
			}
			m.cycles += 8
		}
	case 0x32e1:
		{
			m.ip = 0x32e5
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[3]+0x19f)))
			m.cycles += 8
		}
	case 0x32e5:
		{
			m.ip = 0x32e9
			m.wr8(m.r[11], uint16(m.r[3]+0x1a1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x32e9:
		{
			m.ip = 0x32ec
			target := uint16(13361)
			m.push(0x32ec)
			m.ip = target
			m.cycles += 19
		}
	case 0x32ec:
		{
			m.ip = 0x32f2
			m.wr16(m.r[11], uint16(0x55b), uint16(0))
			m.cycles += 8
		}
	case 0x32f2:
		{
			m.ip = 0x32f6
			m.r[0] = uint16(0x597)
			m.cycles += 3
		}
	case 0x32f6:
		{
			m.ip = 0x32f9
			m.wr16(m.r[11], uint16(0x55d), m.r[0])
			m.cycles += 8
		}
	case 0x32f9:
		{
			m.ip = 0x32fc
			m.r[6] = uint16(47200)
			m.cycles += 2
		}
	case 0x32fc:
		{
			m.ip = 0x32ff
			m.r[1] = uint16(16)
			m.cycles += 2
		}
	case 0x32ff:
		{
			m.ip = 0x3300
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x3300:
		{
			m.ip = 0x3301
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x3301:
		{
			m.ip = 0x3305
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3305:
		{
			m.ip = 0x3309
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1ad))
			m.cycles += 8
		}
	case 0x3309:
		{
			m.ip = 0x330c
			target := uint16(10677)
			m.push(0x330c)
			m.ip = target
			m.cycles += 19
		}
	case 0x330c:
		{
			m.ip = 0x330f
			target := uint16(9749)
			m.push(0x330f)
			m.ip = target
			m.cycles += 19
		}
	case 0x330f:
		{
			m.ip = 0x3314
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x48)), uint16(67))
			m.cycles += 16
		}
	case 0x3314:
		{
			m.ip = 0x3316
			if m.zf {
				m.ip = uint16(13086)
			}
			m.cycles += 8
		}
	case 0x3316:
		{
			m.ip = 0x331b
			m.wr8(m.r[11], uint16(0x49), uint16(67))
			m.cycles += 8
		}
	case 0x331b:
		{
			m.ip = 0x331e
			target := uint16(12567)
			m.push(0x331e)
			m.ip = target
			m.cycles += 19
		}
	case 0x331e:
		{
			m.ip = 0x331f
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x331f:
		{
			m.ip = 0x3322
			m.r[6] = m.alu("add", 16, m.r[6], uint16(4))
			m.cycles += 4
		}
	case 0x3322:
		{
			m.ip = 0x3328
			m.wr16(m.r[11], uint16(0x106a), uint16(100))
			m.cycles += 8
		}
	case 0x3328:
		{
			m.ip = 0x332b
			target := uint16(19182)
			m.push(0x332b)
			m.ip = target
			m.cycles += 19
		}
	case 0x332b:
		{
			m.ip = 0x332c
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x332c:
		{
			m.ip = 0x332e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(13055)
			}
			m.cycles += 17
		}
	case 0x332e:
		{
			m.ip = 0x3334
			m.wr16(m.r[11], uint16(0x106a), uint16(1500))
			m.cycles += 8
		}
	case 0x3334:
		{
			m.ip = 0x3337
			target := uint16(19182)
			m.push(0x3337)
			m.ip = target
			m.cycles += 19
		}
	case 0x3337:
		{
			m.ip = 0x3338
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3338:
		{
			m.ip = 0x3339
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x3339:
		{
			m.ip = 0x333a
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x333a:
		{
			m.ip = 0x333e
			m.r[3] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x333e:
		{
			m.ip = 0x3342
			m.r[6] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3342:
		{
			m.ip = 0x3347
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x68)), uint16(0))
			m.cycles += 16
		}
	case 0x3347:
		{
			m.ip = 0x3349
			if !m.zf {
				m.ip = uint16(13254)
			}
			m.cycles += 8
		}
	case 0x3349:
		{
			m.ip = 0x334e
			m.wr8(m.r[11], uint16(m.r[3]+0x6a), uint16(0))
			m.cycles += 8
		}
	case 0x334e:
		{
			m.ip = 0x3352
			m.r[7] = m.rd16(m.r[11], uint16(m.r[6]+0x1ad))
			m.cycles += 8
		}
	case 0x3352:
		{
			m.ip = 0x3356
			m.r[5] = m.rd16(m.r[11], uint16(m.r[6]+0x1a5))
			m.cycles += 8
		}
	case 0x3356:
		{
			m.ip = 0x335a
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(0x1ed)))
			m.cycles += 16
		}
	case 0x335a:
		{
			m.ip = 0x335c
			if !m.zf {
				m.ip = uint16(13184)
			}
			m.cycles += 8
		}
	case 0x335c:
		{
			m.ip = 0x3360
			m.set8(1, 0, m.rd8(m.r[11], uint16(0x1fd)))
			m.cycles += 8
		}
	case 0x3360:
		{
			m.ip = 0x3364
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), m.rd8(m.r[11], uint16(m.r[3]+0x19f)))
			m.cycles += 16
		}
	case 0x3364:
		{
			m.ip = 0x3366
			if m.zf {
				m.ip = uint16(13168)
			}
			m.cycles += 8
		}
	case 0x3366:
		{
			m.ip = 0x336a
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), m.rd8(m.r[11], uint16(m.r[3]+0x1a1)))
			m.cycles += 16
		}
	case 0x336a:
		{
			m.ip = 0x336c
			if !m.zf {
				m.ip = uint16(13258)
			}
			m.cycles += 8
		}
	case 0x336c:
		{
			m.ip = 0x3370
			m.wr8(m.r[11], uint16(m.r[3]+0x19f), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3370:
		{
			m.ip = 0x3376
			m.wr16(m.r[11], uint16(m.r[6]+0x6c), uint16(0))
			m.cycles += 8
		}
	case 0x3376:
		{
			m.ip = 0x3379
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x203)))
			m.cycles += 8
		}
	case 0x3379:
		{
			m.ip = 0x337d
			m.wr8(m.r[11], uint16(m.r[3]+0x201), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x337d:
		{
			m.ip = 0x337f
			m.ip = uint16(13241)
			m.cycles += 15
		}
	case 0x337f:
		{
			m.ip = 0x3380
			m.cycles += 3
		}
	case 0x3380:
		{
			m.ip = 0x3384
			m.r[5] = m.rd16(m.r[11], uint16(m.r[6]+0x1a5))
			m.cycles += 8
		}
	case 0x3384:
		{
			m.ip = 0x3388
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(0x1ef)))
			m.cycles += 16
		}
	case 0x3388:
		{
			m.ip = 0x338a
			if !m.zf {
				m.ip = uint16(13258)
			}
			m.cycles += 8
		}
	case 0x338a:
		{
			m.ip = 0x338e
			m.set8(1, 0, m.rd8(m.r[11], uint16(0x1fe)))
			m.cycles += 8
		}
	case 0x338e:
		{
			m.ip = 0x3392
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), m.rd8(m.r[11], uint16(m.r[3]+0x19f)))
			m.cycles += 16
		}
	case 0x3392:
		{
			m.ip = 0x3394
			if m.zf {
				m.ip = uint16(13214)
			}
			m.cycles += 8
		}
	case 0x3394:
		{
			m.ip = 0x3398
			m.alu("sub", 8, ((m.r[1] >> 0) & 255), m.rd8(m.r[11], uint16(m.r[3]+0x1a1)))
			m.cycles += 16
		}
	case 0x3398:
		{
			m.ip = 0x339a
			if !m.zf {
				m.ip = uint16(13258)
			}
			m.cycles += 8
		}
	case 0x339a:
		{
			m.ip = 0x339e
			m.wr8(m.r[11], uint16(m.r[3]+0x19f), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x339e:
		{
			m.ip = 0x33a3
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x466)), uint16(8))
			m.cycles += 16
		}
	case 0x33a3:
		{
			m.ip = 0x33a5
			if !m.zf {
				m.ip = uint16(13228)
			}
			m.cycles += 8
		}
	case 0x33a5:
		{
			m.ip = 0x33aa
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1a1)), uint16(16))
			m.cycles += 16
		}
	case 0x33aa:
		{
			m.ip = 0x33ac
			if m.zf {
				m.ip = uint16(13258)
			}
			m.cycles += 8
		}
	case 0x33ac:
		{
			m.ip = 0x33b2
			m.wr16(m.r[11], uint16(m.r[6]+0x6c), uint16(2))
			m.cycles += 8
		}
	case 0x33b2:
		{
			m.ip = 0x33b5
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x204)))
			m.cycles += 8
		}
	case 0x33b5:
		{
			m.ip = 0x33b9
			m.wr8(m.r[11], uint16(m.r[3]+0x201), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x33b9:
		{
			m.ip = 0x33be
			m.wr8(m.r[11], uint16(m.r[3]+0x68), uint16(27))
			m.cycles += 8
		}
	case 0x33be:
		{
			m.ip = 0x33c3
			m.wr8(m.r[11], uint16(m.r[3]+0x6a), uint16(1))
			m.cycles += 8
		}
	case 0x33c3:
		{
			m.ip = 0x33c4
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x33c4:
		{
			m.ip = 0x33c5
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x33c5:
		{
			m.ip = 0x33c6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x33c6:
		{
			m.ip = 0x33ca
			m.wr8(m.r[11], uint16(m.r[3]+0x68), m.unary("dec", 8, m.rd8(m.r[11], uint16(m.r[3]+0x68))))
			m.cycles += 15
		}
	case 0x33ca:
		{
			m.ip = 0x33cb
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x33cb:
		{
			m.ip = 0x33cc
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x33cc:
		{
			m.ip = 0x33cd
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x33cd:
		{
			m.ip = 0x33ce
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x33ce:
		{
			m.ip = 0x33cf
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x33cf:
		{
			m.ip = 0x33d0
			m.push(m.r[5])
			m.cycles += 11
		}
	case 0x33d0:
		{
			m.ip = 0x33d4
			m.r[3] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x33d4:
		{
			m.ip = 0x33d8
			m.r[6] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x33d8:
		{
			m.ip = 0x33dc
			m.set8(1, 0, m.rd8(m.r[11], uint16(0x200)))
			m.cycles += 8
		}
	case 0x33dc:
		{
			m.ip = 0x33e0
			m.wr8(m.r[11], uint16(m.r[3]+0x1a1), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x33e0:
		{
			m.ip = 0x33e4
			m.wr8(m.r[11], uint16(m.r[3]+0x19f), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x33e4:
		{
			m.ip = 0x33e8
			m.r[5] = m.rd16(m.r[11], uint16(0x1f7))
			m.cycles += 8
		}
	case 0x33e8:
		{
			m.ip = 0x33ec
			m.wr16(m.r[11], uint16(m.r[6]+0x1ad), m.r[5])
			m.cycles += 8
		}
	case 0x33ec:
		{
			m.ip = 0x33f0
			m.r[5] = m.rd16(m.r[11], uint16(0x1f3))
			m.cycles += 8
		}
	case 0x33f0:
		{
			m.ip = 0x33f4
			m.wr16(m.r[11], uint16(m.r[6]+0x1a5), m.r[5])
			m.cycles += 8
		}
	case 0x33f4:
		{
			m.ip = 0x33f7
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x204)))
			m.cycles += 8
		}
	case 0x33f7:
		{
			m.ip = 0x33fb
			m.wr8(m.r[11], uint16(m.r[3]+0x201), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x33fb:
		{
			m.ip = 0x33fc
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x33fc:
		{
			m.ip = 0x33fd
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x33fd:
		{
			m.ip = 0x33fe
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x33fe:
		{
			m.ip = 0x33ff
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x33ff:
		{
			m.ip = 0x3400
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x3400:
		{
			m.ip = 0x3401
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x3401:
		{
			m.ip = 0x3402
			m.push(m.r[5])
			m.cycles += 11
		}
	case 0x3402:
		{
			m.ip = 0x3406
			m.r[3] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x3406:
		{
			m.ip = 0x340a
			m.r[6] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x340a:
		{
			m.ip = 0x340e
			m.set8(1, 0, m.rd8(m.r[11], uint16(0x1ff)))
			m.cycles += 8
		}
	case 0x340e:
		{
			m.ip = 0x3412
			m.wr8(m.r[11], uint16(m.r[3]+0x1a1), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3412:
		{
			m.ip = 0x3416
			m.wr8(m.r[11], uint16(m.r[3]+0x19f), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3416:
		{
			m.ip = 0x341a
			m.r[5] = m.rd16(m.r[11], uint16(0x1f5))
			m.cycles += 8
		}
	case 0x341a:
		{
			m.ip = 0x341e
			m.wr16(m.r[11], uint16(m.r[6]+0x1ad), m.r[5])
			m.cycles += 8
		}
	case 0x341e:
		{
			m.ip = 0x3422
			m.r[5] = m.rd16(m.r[11], uint16(0x1f1))
			m.cycles += 8
		}
	case 0x3422:
		{
			m.ip = 0x3426
			m.wr16(m.r[11], uint16(m.r[6]+0x1a5), m.r[5])
			m.cycles += 8
		}
	case 0x3426:
		{
			m.ip = 0x3429
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x203)))
			m.cycles += 8
		}
	case 0x3429:
		{
			m.ip = 0x342d
			m.wr8(m.r[11], uint16(m.r[3]+0x201), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x342d:
		{
			m.ip = 0x342e
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x342e:
		{
			m.ip = 0x342f
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x342f:
		{
			m.ip = 0x3430
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x3430:
		{
			m.ip = 0x3431
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3431:
		{
			m.ip = 0x3435
			m.r[6] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x3435:
		{
			m.ip = 0x343a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a1)), uint16(255))
			m.cycles += 16
		}
	case 0x343a:
		{
			m.ip = 0x343c
			if !m.zf {
				m.ip = uint16(13375)
			}
			m.cycles += 8
		}
	case 0x343c:
		{
			m.ip = 0x343f
			m.ip = uint16(14124)
			m.cycles += 15
		}
	case 0x343f:
		{
			m.ip = 0x3443
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3443:
		{
			m.ip = 0x3446
			target := uint16(14163)
			m.push(0x3446)
			m.ip = target
			m.cycles += 19
		}
	case 0x3446:
		{
			m.ip = 0x344a
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x1a5))
			m.cycles += 8
		}
	case 0x344a:
		{
			m.ip = 0x344e
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[6]+0x1a3)))
			m.cycles += 8
		}
	case 0x344e:
		{
			m.ip = 0x3450
			m.set8(3, 0, m.unary("inc", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x3450:
		{
			m.ip = 0x3453
			m.set8(3, 0, m.alu("and", 8, ((m.r[3]>>0)&255), uint16(7)))
			m.cycles += 4
		}
	case 0x3453:
		{
			m.ip = 0x3457
			m.wr8(m.r[11], uint16(m.r[6]+0x1a3), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3457:
		{
			m.ip = 0x345a
			m.alu("sub", 8, ((m.r[3] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x345a:
		{
			m.ip = 0x345c
			if !m.zf {
				m.ip = uint16(13407)
			}
			m.cycles += 8
		}
	case 0x345c:
		{
			m.ip = 0x345f
			m.ip = uint16(13609)
			m.cycles += 15
		}
	case 0x345f:
		{
			m.ip = 0x3464
			m.wr8(m.r[11], uint16(m.r[6]+0x257), uint16(0))
			m.cycles += 8
		}
	case 0x3464:
		{
			m.ip = 0x3468
			m.set8(2, 8, m.rd8(m.r[11], uint16(m.r[6]+0x19f)))
			m.cycles += 8
		}
	case 0x3468:
		{
			m.ip = 0x346c
			m.set8(2, 0, m.rd8(m.r[11], uint16(m.r[6]+0x1a1)))
			m.cycles += 8
		}
	case 0x346c:
		{
			m.ip = 0x346e
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), ((m.r[2] >> 0) & 255))
			m.cycles += 4
		}
	case 0x346e:
		{
			m.ip = 0x3470
			if m.zf {
				m.ip = uint16(13474)
			}
			m.cycles += 8
		}
	case 0x3470:
		{
			m.ip = 0x3475
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x3475:
		{
			m.ip = 0x3477
			if !m.zf {
				m.ip = uint16(13467)
			}
			m.cycles += 8
		}
	case 0x3477:
		{
			m.ip = 0x3478
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x3478:
		{
			m.ip = 0x347a
			m.r[3] = m.r[6]
			m.cycles += 2
		}
	case 0x347a:
		{
			m.ip = 0x347b
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x347b:
		{
			m.ip = 0x347e
			m.r[3] = m.alu("and", 16, m.r[3], uint16(1))
			m.cycles += 4
		}
	case 0x347e:
		{
			m.ip = 0x3480
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3480:
		{
			m.ip = 0x3482
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x3482:
		{
			m.ip = 0x3484
			m.set8(0, 0, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x3484:
		{
			m.ip = 0x3486
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x3486:
		{
			m.ip = 0x3488
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x3488:
		{
			m.ip = 0x348a
			m.r[7] = m.r[0]
			m.cycles += 2
		}
	case 0x348a:
		{
			m.ip = 0x348e
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x1013))
			m.cycles += 8
		}
	case 0x348e:
		{
			m.ip = 0x3490
			m.r[7] = m.alu("add", 16, m.r[7], m.r[5])
			m.cycles += 4
		}
	case 0x3490:
		{
			m.ip = 0x3494
			m.alu("sub", 16, m.r[7], m.rd16(m.r[11], uint16(m.r[3]+0x1a5)))
			m.cycles += 16
		}
	case 0x3494:
		{
			m.ip = 0x3495
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x3495:
		{
			m.ip = 0x3497
			if m.zf {
				m.ip = uint16(13474)
			}
			m.cycles += 8
		}
	case 0x3497:
		{
			m.ip = 0x349b
			m.wr8(m.r[11], uint16(m.r[6]+0x19f), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x349b:
		{
			m.ip = 0x349d
			m.set8(2, 0, m.alu("xor", 8, ((m.r[2]>>0)&255), ((m.r[2]>>8)&255)))
			m.cycles += 4
		}
	case 0x349d:
		{
			m.ip = 0x34a0
			m.alu("and", 8, ((m.r[2] >> 0) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x34a0:
		{
			m.ip = 0x34a2
			if m.zf {
				m.ip = uint16(13509)
			}
			m.cycles += 8
		}
	case 0x34a2:
		{
			m.ip = 0x34a6
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x34a6:
		{
			m.ip = 0x34aa
			m.r[6] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x34aa:
		{
			m.ip = 0x34ad
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x34ad:
		{
			m.ip = 0x34af
			if !m.zf {
				m.ip = uint16(13490)
			}
			m.cycles += 8
		}
	case 0x34af:
		{
			m.ip = 0x34b2
			m.ip = uint16(14005)
			m.cycles += 15
		}
	case 0x34b2:
		{
			m.ip = 0x34b5
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(8))
			m.cycles += 4
		}
	case 0x34b5:
		{
			m.ip = 0x34b7
			if !m.zf {
				m.ip = uint16(13498)
			}
			m.cycles += 8
		}
	case 0x34b7:
		{
			m.ip = 0x34ba
			m.ip = uint16(14068)
			m.cycles += 15
		}
	case 0x34ba:
		{
			m.ip = 0x34bd
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x34bd:
		{
			m.ip = 0x34bf
			if !m.zf {
				m.ip = uint16(13506)
			}
			m.cycles += 8
		}
	case 0x34bf:
		{
			m.ip = 0x34c2
			m.ip = uint16(13969)
			m.cycles += 15
		}
	case 0x34c2:
		{
			m.ip = 0x34c5
			m.ip = uint16(13925)
			m.cycles += 15
		}
	case 0x34c5:
		{
			m.ip = 0x34ca
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x6a)), uint16(0))
			m.cycles += 16
		}
	case 0x34ca:
		{
			m.ip = 0x34cc
			if m.zf {
				m.ip = uint16(13579)
			}
			m.cycles += 8
		}
	case 0x34cc:
		{
			m.ip = 0x34d0
			m.wr8(m.r[11], uint16(m.r[6]+0x68), m.unary("neg", 8, m.rd8(m.r[11], uint16(m.r[6]+0x68))))
			m.cycles += 15
		}
	case 0x34d0:
		{
			m.ip = 0x34d5
			m.wr8(m.r[11], uint16(m.r[6]+0x68), m.alu("add", 8, m.rd8(m.r[11], uint16(m.r[6]+0x68)), uint16(26)))
			m.cycles += 16
		}
	case 0x34d5:
		{
			m.ip = 0x34d9
			m.set8(2, 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a1)))
			m.cycles += 8
		}
	case 0x34d9:
		{
			m.ip = 0x34dd
			m.wr8(m.r[11], uint16(m.r[6]+0x19f), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x34dd:
		{
			m.ip = 0x34e2
			m.wr8(m.r[11], uint16(m.r[6]+0x1a3), m.alu("xor", 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a3)), uint16(255)))
			m.cycles += 16
		}
	case 0x34e2:
		{
			m.ip = 0x34e6
			m.wr8(m.r[11], uint16(m.r[6]+0x1a3), m.unary("inc", 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a3))))
			m.cycles += 15
		}
	case 0x34e6:
		{
			m.ip = 0x34eb
			m.wr8(m.r[11], uint16(m.r[6]+0x1a3), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a3)), uint16(7)))
			m.cycles += 16
		}
	case 0x34eb:
		{
			m.ip = 0x34ed
			m.r[6] = m.shift("shl", 16, m.r[6], uint16(1))
			m.cycles += 8
		}
	case 0x34ed:
		{
			m.ip = 0x34f1
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]+0x1a5))
			m.cycles += 8
		}
	case 0x34f1:
		{
			m.ip = 0x34f5
			m.r[3] = m.rd16(m.r[11], uint16(m.r[6]+0x1a9))
			m.cycles += 8
		}
	case 0x34f5:
		{
			m.ip = 0x34f9
			m.wr16(m.r[11], uint16(m.r[6]+0x1a5), m.r[3])
			m.cycles += 8
		}
	case 0x34f9:
		{
			m.ip = 0x34fd
			m.wr16(m.r[11], uint16(m.r[6]+0x1a9), m.r[0])
			m.cycles += 8
		}
	case 0x34fd:
		{
			m.ip = 0x3502
			m.wr16(m.r[11], uint16(m.r[6]+0x6c), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[6]+0x6c)), uint16(2)))
			m.cycles += 16
		}
	case 0x3502:
		{
			m.ip = 0x3507
			m.wr16(m.r[11], uint16(m.r[6]+0x6c), m.alu("and", 16, m.rd16(m.r[11], uint16(m.r[6]+0x6c)), uint16(2)))
			m.cycles += 16
		}
	case 0x3507:
		{
			m.ip = 0x3509
			m.r[6] = m.shift("shr", 16, m.r[6], uint16(1))
			m.cycles += 8
		}
	case 0x3509:
		{
			m.ip = 0x350b
			m.ip = uint16(13474)
			m.cycles += 15
		}
	case 0x350b:
		{
			m.ip = 0x350f
			m.set8(2, 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a1)))
			m.cycles += 8
		}
	case 0x350f:
		{
			m.ip = 0x3513
			m.wr8(m.r[11], uint16(m.r[6]+0x19f), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x3513:
		{
			m.ip = 0x3518
			m.wr8(m.r[11], uint16(m.r[6]+0x1a3), m.alu("xor", 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a3)), uint16(255)))
			m.cycles += 16
		}
	case 0x3518:
		{
			m.ip = 0x351c
			m.wr8(m.r[11], uint16(m.r[6]+0x1a3), m.unary("inc", 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a3))))
			m.cycles += 15
		}
	case 0x351c:
		{
			m.ip = 0x3521
			m.wr8(m.r[11], uint16(m.r[6]+0x1a3), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[6]+0x1a3)), uint16(7)))
			m.cycles += 16
		}
	case 0x3521:
		{
			m.ip = 0x3524
			m.alu("sub", 8, ((m.r[3] >> 0) & 255), uint16(6))
			m.cycles += 4
		}
	case 0x3524:
		{
			m.ip = 0x3526
			if !m.cf {
				m.ip = uint16(13609)
			}
			m.cycles += 8
		}
	case 0x3526:
		{
			m.ip = 0x3529
			m.ip = uint16(13856)
			m.cycles += 15
		}
	case 0x3529:
		{
			m.ip = 0x352e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x257)), uint16(1))
			m.cycles += 16
		}
	case 0x352e:
		{
			m.ip = 0x3530
			if m.zf {
				m.ip = uint16(13645)
			}
			m.cycles += 8
		}
	case 0x3530:
		{
			m.ip = 0x3535
			m.wr8(m.r[11], uint16(m.r[6]+0x257), uint16(1))
			m.cycles += 8
		}
	case 0x3535:
		{
			m.ip = 0x3539
			m.r[6] = uint16(0x262)
			m.cycles += 3
		}
	case 0x3539:
		{
			m.ip = 0x353e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(2))
			m.cycles += 16
		}
	case 0x353e:
		{
			m.ip = 0x3540
			if m.zf {
				m.ip = uint16(13643)
			}
			m.cycles += 8
		}
	case 0x3540:
		{
			m.ip = 0x3545
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x223)), uint16(1))
			m.cycles += 16
		}
	case 0x3545:
		{
			m.ip = 0x3547
			if m.zf {
				m.ip = uint16(13643)
			}
			m.cycles += 8
		}
	case 0x3547:
		{
			m.ip = 0x354b
			m.r[6] = uint16(0x259)
			m.cycles += 3
		}
	case 0x354b:
		{
			m.ip = 0x354d
			m.wr16(m.r[11], uint16(m.r[6]), m.unary("inc", 16, m.rd16(m.r[11], uint16(m.r[6]))))
			m.cycles += 15
		}
	case 0x354d:
		{
			m.ip = 0x3553
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x3553:
		{
			m.ip = 0x3555
			if m.zf {
				m.ip = uint16(13713)
			}
			m.cycles += 8
		}
	case 0x3555:
		{
			m.ip = 0x355a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x46a)), uint16(1))
			m.cycles += 16
		}
	case 0x355a:
		{
			m.ip = 0x355c
			if !m.zf {
				m.ip = uint16(13684)
			}
			m.cycles += 8
		}
	case 0x355c:
		{
			m.ip = 0x3560
			m.r[6] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3560:
		{
			m.ip = 0x3564
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(m.r[6]+0x46b)))
			m.cycles += 16
		}
	case 0x3564:
		{
			m.ip = 0x3566
			if m.zf {
				m.ip = uint16(13705)
			}
			m.cycles += 8
		}
	case 0x3566:
		{
			m.ip = 0x356a
			m.wr16(m.r[11], uint16(m.r[6]+0x46b), m.r[5])
			m.cycles += 8
		}
	case 0x356a:
		{
			m.ip = 0x356e
			m.wr16(m.r[11], uint16(m.r[6]+0x233), m.unary("inc", 16, m.rd16(m.r[11], uint16(m.r[6]+0x233))))
			m.cycles += 15
		}
	case 0x356e:
		{
			m.ip = 0x3571
			target := uint16(17038)
			m.push(0x3571)
			m.ip = target
			m.cycles += 19
		}
	case 0x3571:
		{
			m.ip = 0x3573
			m.ip = uint16(13705)
			m.cycles += 15
		}
	case 0x3573:
		{
			m.ip = 0x3574
			m.cycles += 3
		}
	case 0x3574:
		{
			m.ip = 0x3578
			m.r[6] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3578:
		{
			m.ip = 0x357c
			m.wr16(m.r[11], uint16(m.r[6]+0x233), m.unary("inc", 16, m.rd16(m.r[11], uint16(m.r[6]+0x233))))
			m.cycles += 15
		}
	case 0x357c:
		{
			m.ip = 0x357f
			target := uint16(17038)
			m.push(0x357f)
			m.ip = target
			m.cycles += 19
		}
	case 0x357f:
		{
			m.ip = 0x3583
			m.wr16(m.r[11], uint16(0x78), m.unary("dec", 16, m.rd16(m.r[11], uint16(0x78))))
			m.cycles += 15
		}
	case 0x3583:
		{
			m.ip = 0x3589
			m.wr8(m.r[11], uint16(m.r[5]+0x8e), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(247)))
			m.cycles += 16
		}
	case 0x3589:
		{
			m.ip = 0x358c
			m.alu("sub", 8, ((m.r[3] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x358c:
		{
			m.ip = 0x358e
			if m.zf {
				m.ip = uint16(13713)
			}
			m.cycles += 8
		}
	case 0x358e:
		{
			m.ip = 0x3591
			m.ip = uint16(13856)
			m.cycles += 15
		}
	case 0x3591:
		{
			m.ip = 0x3594
			target := uint16(13112)
			m.push(0x3594)
			m.ip = target
			m.cycles += 19
		}
	case 0x3594:
		{
			m.ip = 0x3598
			m.r[6] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x3598:
		{
			m.ip = 0x359d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x68)), uint16(0))
			m.cycles += 16
		}
	case 0x359d:
		{
			m.ip = 0x35a1
			m.set8(2, 8, m.rd8(m.r[11], uint16(m.r[6]+0x19f)))
			m.cycles += 8
		}
	case 0x35a1:
		{
			m.ip = 0x35a3
			if !m.zf {
				m.ip = uint16(13856)
			}
			m.cycles += 8
		}
	case 0x35a3:
		{
			m.ip = 0x35a7
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[6]+0x1a1)))
			m.cycles += 8
		}
	case 0x35a7:
		{
			m.ip = 0x35a9
			m.set8(2, 8, ((m.r[1] >> 0) & 255))
			m.cycles += 2
		}
	case 0x35a9:
		{
			m.ip = 0x35ab
			m.set8(1, 0, m.shift("ror", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x35ab:
		{
			m.ip = 0x35ad
			m.set8(1, 0, m.shift("ror", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x35ad:
		{
			m.ip = 0x35af
			m.set8(1, 0, m.shift("ror", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x35af:
		{
			m.ip = 0x35b1
			m.set8(0, 0, uint16(16))
			m.cycles += 2
		}
	case 0x35b1:
		{
			m.ip = 0x35b3
			m.set8(0, 0, m.shift("rol", 8, ((m.r[0]>>0)&255), ((m.r[1]>>0)&255)))
			m.cycles += 8
		}
	case 0x35b3:
		{
			m.ip = 0x35b8
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), ((m.r[0] >> 0) & 255))
			m.cycles += 16
		}
	case 0x35b8:
		{
			m.ip = 0x35ba
			if m.zf {
				m.ip = uint16(13798)
			}
			m.cycles += 8
		}
	case 0x35ba:
		{
			m.ip = 0x35bf
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x35bf:
		{
			m.ip = 0x35c1
			if !m.zf {
				m.ip = uint16(13791)
			}
			m.cycles += 8
		}
	case 0x35c1:
		{
			m.ip = 0x35c3
			m.r[3] = m.r[6]
			m.cycles += 2
		}
	case 0x35c3:
		{
			m.ip = 0x35c4
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x35c4:
		{
			m.ip = 0x35c7
			m.r[3] = m.alu("and", 16, m.r[3], uint16(1))
			m.cycles += 4
		}
	case 0x35c7:
		{
			m.ip = 0x35c9
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x35c9:
		{
			m.ip = 0x35cb
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x35cb:
		{
			m.ip = 0x35cd
			m.set8(0, 0, ((m.r[2] >> 8) & 255))
			m.cycles += 2
		}
	case 0x35cd:
		{
			m.ip = 0x35cf
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x35cf:
		{
			m.ip = 0x35d1
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x35d1:
		{
			m.ip = 0x35d3
			m.r[7] = m.r[0]
			m.cycles += 2
		}
	case 0x35d3:
		{
			m.ip = 0x35d7
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x1013))
			m.cycles += 8
		}
	case 0x35d7:
		{
			m.ip = 0x35d9
			m.r[7] = m.alu("add", 16, m.r[7], m.r[5])
			m.cycles += 4
		}
	case 0x35d9:
		{
			m.ip = 0x35dd
			m.alu("sub", 16, m.r[7], m.rd16(m.r[11], uint16(m.r[3]+0x1a5)))
			m.cycles += 16
		}
	case 0x35dd:
		{
			m.ip = 0x35df
			if m.zf {
				m.ip = uint16(13798)
			}
			m.cycles += 8
		}
	case 0x35df:
		{
			m.ip = 0x35e3
			m.wr8(m.r[11], uint16(m.r[6]+0x19f), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x35e3:
		{
			m.ip = 0x35e5
			m.ip = uint16(13856)
			m.cycles += 15
		}
	case 0x35e5:
		{
			m.ip = 0x35e6
			m.cycles += 3
		}
	case 0x35e6:
		{
			m.ip = 0x35ea
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[6]+0x19f)))
			m.cycles += 8
		}
	case 0x35ea:
		{
			m.ip = 0x35ec
			m.set8(2, 8, ((m.r[1] >> 0) & 255))
			m.cycles += 2
		}
	case 0x35ec:
		{
			m.ip = 0x35ee
			m.set8(1, 0, m.shift("ror", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x35ee:
		{
			m.ip = 0x35f0
			m.set8(1, 0, m.shift("ror", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x35f0:
		{
			m.ip = 0x35f2
			m.set8(1, 0, m.shift("ror", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x35f2:
		{
			m.ip = 0x35f4
			m.set8(0, 0, uint16(16))
			m.cycles += 2
		}
	case 0x35f4:
		{
			m.ip = 0x35f6
			m.set8(0, 0, m.shift("rol", 8, ((m.r[0]>>0)&255), ((m.r[1]>>0)&255)))
			m.cycles += 8
		}
	case 0x35f6:
		{
			m.ip = 0x35fb
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), ((m.r[0] >> 0) & 255))
			m.cycles += 16
		}
	case 0x35fb:
		{
			m.ip = 0x35fd
			if m.zf {
				m.ip = uint16(13848)
			}
			m.cycles += 8
		}
	case 0x35fd:
		{
			m.ip = 0x3602
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x3602:
		{
			m.ip = 0x3604
			if !m.zf {
				m.ip = uint16(13856)
			}
			m.cycles += 8
		}
	case 0x3604:
		{
			m.ip = 0x3606
			m.set8(0, 0, ((m.r[2] >> 8) & 255))
			m.cycles += 2
		}
	case 0x3606:
		{
			m.ip = 0x3608
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x3608:
		{
			m.ip = 0x360a
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x360a:
		{
			m.ip = 0x360c
			m.r[7] = m.r[0]
			m.cycles += 2
		}
	case 0x360c:
		{
			m.ip = 0x3610
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x1013))
			m.cycles += 8
		}
	case 0x3610:
		{
			m.ip = 0x3612
			m.r[7] = m.alu("add", 16, m.r[7], m.r[5])
			m.cycles += 4
		}
	case 0x3612:
		{
			m.ip = 0x3616
			m.alu("sub", 16, m.r[7], m.rd16(m.r[11], uint16(m.r[3]+0x1a5)))
			m.cycles += 16
		}
	case 0x3616:
		{
			m.ip = 0x3618
			if !m.zf {
				m.ip = uint16(13856)
			}
			m.cycles += 8
		}
	case 0x3618:
		{
			m.ip = 0x361d
			m.wr8(m.r[11], uint16(m.r[6]+0x1a3), uint16(7))
			m.cycles += 8
		}
	case 0x361d:
		{
			m.ip = 0x3620
			m.ip = uint16(14124)
			m.cycles += 15
		}
	case 0x3620:
		{
			m.ip = 0x3624
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3624:
		{
			m.ip = 0x3628
			m.r[6] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x3628:
		{
			m.ip = 0x362c
			m.wr16(m.r[11], uint16(m.r[3]+0x1a9), m.r[5])
			m.cycles += 8
		}
	case 0x362c:
		{
			m.ip = 0x362f
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x362f:
		{
			m.ip = 0x3631
			if !m.zf {
				m.ip = uint16(13883)
			}
			m.cycles += 8
		}
	case 0x3631:
		{
			m.ip = 0x3635
			m.wr16(m.r[11], uint16(m.r[3]+0x1a5), m.unary("inc", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1a5))))
			m.cycles += 15
		}
	case 0x3635:
		{
			m.ip = 0x3638
			target := uint16(14163)
			m.push(0x3638)
			m.ip = target
			m.cycles += 19
		}
	case 0x3638:
		{
			m.ip = 0x363a
			m.ip = uint16(14005)
			m.cycles += 15
		}
	case 0x363a:
		{
			m.ip = 0x363b
			m.cycles += 3
		}
	case 0x363b:
		{
			m.ip = 0x363e
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(8))
			m.cycles += 4
		}
	case 0x363e:
		{
			m.ip = 0x3640
			if !m.zf {
				m.ip = uint16(13898)
			}
			m.cycles += 8
		}
	case 0x3640:
		{
			m.ip = 0x3644
			m.wr16(m.r[11], uint16(m.r[3]+0x1a5), m.unary("dec", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1a5))))
			m.cycles += 15
		}
	case 0x3644:
		{
			m.ip = 0x3647
			target := uint16(14163)
			m.push(0x3647)
			m.ip = target
			m.cycles += 19
		}
	case 0x3647:
		{
			m.ip = 0x364a
			m.ip = uint16(14068)
			m.cycles += 15
		}
	case 0x364a:
		{
			m.ip = 0x364d
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x364d:
		{
			m.ip = 0x364f
			if !m.zf {
				m.ip = uint16(13914)
			}
			m.cycles += 8
		}
	case 0x364f:
		{
			m.ip = 0x3654
			m.wr16(m.r[11], uint16(m.r[3]+0x1a5), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1a5)), uint16(21)))
			m.cycles += 16
		}
	case 0x3654:
		{
			m.ip = 0x3657
			target := uint16(14163)
			m.push(0x3657)
			m.ip = target
			m.cycles += 19
		}
	case 0x3657:
		{
			m.ip = 0x3659
			m.ip = uint16(13969)
			m.cycles += 15
		}
	case 0x3659:
		{
			m.ip = 0x365a
			m.cycles += 3
		}
	case 0x365a:
		{
			m.ip = 0x365f
			m.wr16(m.r[11], uint16(m.r[3]+0x1a5), m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1a5)), uint16(21)))
			m.cycles += 16
		}
	case 0x365f:
		{
			m.ip = 0x3662
			target := uint16(14163)
			m.push(0x3662)
			m.ip = target
			m.cycles += 19
		}
	case 0x3662:
		{
			m.ip = 0x3664
			m.ip = uint16(13925)
			m.cycles += 15
		}
	case 0x3664:
		{
			m.ip = 0x3665
			m.cycles += 3
		}
	case 0x3665:
		{
			m.ip = 0x3669
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1ad))
			m.cycles += 8
		}
	case 0x3669:
		{
			m.ip = 0x366e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x68)), uint16(10))
			m.cycles += 16
		}
	case 0x366e:
		{
			m.ip = 0x3670
			if !m.cf && !m.zf {
				m.ip = uint16(13950)
			}
			m.cycles += 8
		}
	case 0x3670:
		{
			m.ip = 0x3675
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x6a)), uint16(1))
			m.cycles += 16
		}
	case 0x3675:
		{
			m.ip = 0x3677
			if !m.zf {
				m.ip = uint16(13950)
			}
			m.cycles += 8
		}
	case 0x3677:
		{
			m.ip = 0x367b
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(240))
			m.cycles += 4
		}
	case 0x367b:
		{
			m.ip = 0x367d
			m.ip = uint16(13961)
			m.cycles += 15
		}
	case 0x367d:
		{
			m.ip = 0x367e
			m.cycles += 3
		}
	case 0x367e:
		{
			m.ip = 0x3682
			m.r[7] = m.alu("add", 16, m.r[7], uint16(1680))
			m.cycles += 4
		}
	case 0x3682:
		{
			m.ip = 0x3685
			target := uint16(14846)
			m.push(0x3685)
			m.ip = target
			m.cycles += 19
		}
	case 0x3685:
		{
			m.ip = 0x3689
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(2160))
			m.cycles += 4
		}
	case 0x3689:
		{
			m.ip = 0x368d
			m.wr16(m.r[11], uint16(m.r[3]+0x1ad), m.r[7])
			m.cycles += 8
		}
	case 0x368d:
		{
			m.ip = 0x3690
			target := uint16(14200)
			m.push(0x3690)
			m.ip = target
			m.cycles += 19
		}
	case 0x3690:
		{
			m.ip = 0x3691
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3691:
		{
			m.ip = 0x3695
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1ad))
			m.cycles += 8
		}
	case 0x3695:
		{
			m.ip = 0x369a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x68)), uint16(8))
			m.cycles += 16
		}
	case 0x369a:
		{
			m.ip = 0x369c
			if !m.cf && !m.zf {
				m.ip = uint16(13994)
			}
			m.cycles += 8
		}
	case 0x369c:
		{
			m.ip = 0x36a1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x6a)), uint16(1))
			m.cycles += 16
		}
	case 0x36a1:
		{
			m.ip = 0x36a3
			if !m.zf {
				m.ip = uint16(13994)
			}
			m.cycles += 8
		}
	case 0x36a3:
		{
			m.ip = 0x36a7
			m.r[7] = m.alu("add", 16, m.r[7], uint16(240))
			m.cycles += 4
		}
	case 0x36a7:
		{
			m.ip = 0x36a9
			m.ip = uint16(13997)
			m.cycles += 15
		}
	case 0x36a9:
		{
			m.ip = 0x36aa
			m.cycles += 3
		}
	case 0x36aa:
		{
			m.ip = 0x36ad
			target := uint16(14914)
			m.push(0x36ad)
			m.ip = target
			m.cycles += 19
		}
	case 0x36ad:
		{
			m.ip = 0x36b1
			m.wr16(m.r[11], uint16(m.r[3]+0x1ad), m.r[7])
			m.cycles += 8
		}
	case 0x36b1:
		{
			m.ip = 0x36b4
			target := uint16(14200)
			m.push(0x36b4)
			m.ip = target
			m.cycles += 19
		}
	case 0x36b4:
		{
			m.ip = 0x36b5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x36b5:
		{
			m.ip = 0x36b9
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1ad))
			m.cycles += 8
		}
	case 0x36b9:
		{
			m.ip = 0x36bb
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x36bb:
		{
			m.ip = 0x36c0
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x68)), uint16(8))
			m.cycles += 16
		}
	case 0x36c0:
		{
			m.ip = 0x36c2
			if !m.cf && !m.zf {
				m.ip = uint16(14025)
			}
			m.cycles += 8
		}
	case 0x36c2:
		{
			m.ip = 0x36c7
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x6a)), uint16(1))
			m.cycles += 16
		}
	case 0x36c7:
		{
			m.ip = 0x36c9
			if m.zf {
				m.ip = uint16(14046)
			}
			m.cycles += 8
		}
	case 0x36c9:
		{
			m.ip = 0x36ce
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)), uint16(1))
			m.cycles += 16
		}
	case 0x36ce:
		{
			m.ip = 0x36d0
			if m.zf {
				m.ip = uint16(14064)
			}
			m.cycles += 8
		}
	case 0x36d0:
		{
			m.ip = 0x36d3
			target := uint16(14982)
			m.push(0x36d3)
			m.ip = target
			m.cycles += 19
		}
	case 0x36d3:
		{
			m.ip = 0x36d4
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x36d4:
		{
			m.ip = 0x36d6
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x36d6:
		{
			m.ip = 0x36da
			m.wr16(m.r[11], uint16(m.r[3]+0x1ad), m.r[7])
			m.cycles += 8
		}
	case 0x36da:
		{
			m.ip = 0x36dd
			target := uint16(14200)
			m.push(0x36dd)
			m.ip = target
			m.cycles += 19
		}
	case 0x36dd:
		{
			m.ip = 0x36de
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x36de:
		{
			m.ip = 0x36e3
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)), uint16(1))
			m.cycles += 16
		}
	case 0x36e3:
		{
			m.ip = 0x36e5
			if m.zf {
				m.ip = uint16(14064)
			}
			m.cycles += 8
		}
	case 0x36e5:
		{
			m.ip = 0x36e6
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x36e6:
		{
			m.ip = 0x36e8
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x36e8:
		{
			m.ip = 0x36ec
			m.wr16(m.r[11], uint16(m.r[3]+0x1ad), m.r[7])
			m.cycles += 8
		}
	case 0x36ec:
		{
			m.ip = 0x36ef
			target := uint16(14200)
			m.push(0x36ef)
			m.ip = target
			m.cycles += 19
		}
	case 0x36ef:
		{
			m.ip = 0x36f0
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x36f0:
		{
			m.ip = 0x36f3
			target := uint16(14579)
			m.push(0x36f3)
			m.ip = target
			m.cycles += 19
		}
	case 0x36f3:
		{
			m.ip = 0x36f4
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x36f4:
		{
			m.ip = 0x36f8
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1ad))
			m.cycles += 8
		}
	case 0x36f8:
		{
			m.ip = 0x36fa
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x36fa:
		{
			m.ip = 0x36ff
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x68)), uint16(8))
			m.cycles += 16
		}
	case 0x36ff:
		{
			m.ip = 0x3701
			if !m.cf && !m.zf {
				m.ip = uint16(14088)
			}
			m.cycles += 8
		}
	case 0x3701:
		{
			m.ip = 0x3706
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[6]+0x6a)), uint16(1))
			m.cycles += 16
		}
	case 0x3706:
		{
			m.ip = 0x3708
			if m.zf {
				m.ip = uint16(14102)
			}
			m.cycles += 8
		}
	case 0x3708:
		{
			m.ip = 0x370d
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)), uint16(1))
			m.cycles += 16
		}
	case 0x370d:
		{
			m.ip = 0x370f
			if m.zf {
				m.ip = uint16(14113)
			}
			m.cycles += 8
		}
	case 0x370f:
		{
			m.ip = 0x3712
			target := uint16(15018)
			m.push(0x3712)
			m.ip = target
			m.cycles += 19
		}
	case 0x3712:
		{
			m.ip = 0x3715
			target := uint16(14200)
			m.push(0x3715)
			m.ip = target
			m.cycles += 19
		}
	case 0x3715:
		{
			m.ip = 0x3716
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3716:
		{
			m.ip = 0x371b
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)), uint16(1))
			m.cycles += 16
		}
	case 0x371b:
		{
			m.ip = 0x371d
			if m.zf {
				m.ip = uint16(14113)
			}
			m.cycles += 8
		}
	case 0x371d:
		{
			m.ip = 0x3720
			target := uint16(14200)
			m.push(0x3720)
			m.ip = target
			m.cycles += 19
		}
	case 0x3720:
		{
			m.ip = 0x3721
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3721:
		{
			m.ip = 0x3722
			m.r[7] = m.unary("dec", 16, m.r[7])
			m.cycles += 3
		}
	case 0x3722:
		{
			m.ip = 0x3724
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3724:
		{
			m.ip = 0x3728
			m.wr16(m.r[11], uint16(m.r[3]+0x1ad), m.r[7])
			m.cycles += 8
		}
	case 0x3728:
		{
			m.ip = 0x372b
			target := uint16(14579)
			m.push(0x372b)
			m.ip = target
			m.cycles += 19
		}
	case 0x372b:
		{
			m.ip = 0x372c
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x372c:
		{
			m.ip = 0x3730
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3730:
		{
			m.ip = 0x3734
			m.r[6] = m.rd16(m.r[11], uint16(m.r[3]+0x22f))
			m.cycles += 8
		}
	case 0x3734:
		{
			m.ip = 0x3738
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1ad))
			m.cycles += 8
		}
	case 0x3738:
		{
			m.ip = 0x373a
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x373a:
		{
			m.ip = 0x373d
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x373d:
		{
			m.ip = 0x3742
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)), uint16(1))
			m.cycles += 16
		}
	case 0x3742:
		{
			m.ip = 0x3744
			if !m.zf {
				m.ip = uint16(14159)
			}
			m.cycles += 8
		}
	case 0x3744:
		{
			m.ip = 0x3749
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x19f)), uint16(16))
			m.cycles += 16
		}
	case 0x3749:
		{
			m.ip = 0x374b
			if !m.zf {
				m.ip = uint16(14159)
			}
			m.cycles += 8
		}
	case 0x374b:
		{
			m.ip = 0x374e
			target := uint16(14657)
			m.push(0x374e)
			m.ip = target
			m.cycles += 19
		}
	case 0x374e:
		{
			m.ip = 0x374f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x374f:
		{
			m.ip = 0x3752
			target := uint16(14299)
			m.push(0x3752)
			m.ip = target
			m.cycles += 19
		}
	case 0x3752:
		{
			m.ip = 0x3753
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3753:
		{
			m.ip = 0x3759
			m.wr16(m.r[11], uint16(0x21f), uint16(28064))
			m.cycles += 8
		}
	case 0x3759:
		{
			m.ip = 0x375d
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x1a9))
			m.cycles += 8
		}
	case 0x375d:
		{
			m.ip = 0x3763
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(4))
			m.cycles += 16
		}
	case 0x3763:
		{
			m.ip = 0x3765
			if m.zf {
				m.ip = uint16(14185)
			}
			m.cycles += 8
		}
	case 0x3765:
		{
			m.ip = 0x3768
			target := uint16(15549)
			m.push(0x3768)
			m.ip = target
			m.cycles += 19
		}
	case 0x3768:
		{
			m.ip = 0x3769
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3769:
		{
			m.ip = 0x376f
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x376f:
		{
			m.ip = 0x3771
			if m.zf {
				m.ip = uint16(14199)
			}
			m.cycles += 8
		}
	case 0x3771:
		{
			m.ip = 0x3777
			m.wr16(m.r[11], uint16(0x21f), uint16(28068))
			m.cycles += 8
		}
	case 0x3777:
		{
			m.ip = 0x3778
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3778:
		{
			m.ip = 0x3779
			m.push(m.r[5])
			m.cycles += 11
		}
	case 0x3779:
		{
			m.ip = 0x377d
			m.r[5] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x377d:
		{
			m.ip = 0x377f
			m.set8(3, 8, uint16(0))
			m.cycles += 2
		}
	case 0x377f:
		{
			m.ip = 0x3784
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[5]+0x19f)))
			m.cycles += 8
		}
	case 0x3784:
		{
			m.ip = 0x3786
			m.r[5] = m.shift("shl", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x3786:
		{
			m.ip = 0x378b
			m.r[6] = m.rd16(m.r[11], uint16(m.r[5]+0x1a5))
			m.cycles += 8
		}
	case 0x378b:
		{
			m.ip = 0x3790
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[6]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x3790:
		{
			m.ip = 0x3792
			if !m.zf {
				m.ip = uint16(14245)
			}
			m.cycles += 8
		}
	case 0x3792:
		{
			m.ip = 0x3795
			m.r[6] = uint16(56000)
			m.cycles += 2
		}
	case 0x3795:
		{
			m.ip = 0x3797
			m.r[5] = m.shift("shr", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x3797:
		{
			m.ip = 0x379d
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)), uint16(1))
			m.cycles += 16
		}
	case 0x379d:
		{
			m.ip = 0x379f
			if !m.zf {
				m.ip = uint16(14261)
			}
			m.cycles += 8
		}
	case 0x379f:
		{
			m.ip = 0x37a2
			m.r[6] = uint16(57888)
			m.cycles += 2
		}
	case 0x37a2:
		{
			m.ip = 0x37a4
			m.ip = uint16(14261)
			m.cycles += 15
		}
	case 0x37a4:
		{
			m.ip = 0x37a5
			m.cycles += 3
		}
	case 0x37a5:
		{
			m.ip = 0x37a8
			m.r[6] = uint16(28000)
			m.cycles += 2
		}
	case 0x37a8:
		{
			m.ip = 0x37aa
			m.r[5] = m.shift("shr", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x37aa:
		{
			m.ip = 0x37b0
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)), uint16(1))
			m.cycles += 16
		}
	case 0x37b0:
		{
			m.ip = 0x37b2
			if !m.zf {
				m.ip = uint16(14261)
			}
			m.cycles += 8
		}
	case 0x37b2:
		{
			m.ip = 0x37b5
			m.r[6] = uint16(29888)
			m.cycles += 2
		}
	case 0x37b5:
		{
			m.ip = 0x37b7
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x37b7:
		{
			m.ip = 0x37b9
			m.set8(2, 8, uint16(0))
			m.cycles += 2
		}
	case 0x37b9:
		{
			m.ip = 0x37be
			m.set8(2, 0, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)))
			m.cycles += 8
		}
	case 0x37be:
		{
			m.ip = 0x37c1
			m.set8(2, 0, m.alu("and", 8, ((m.r[2]>>0)&255), uint16(254)))
			m.cycles += 4
		}
	case 0x37c1:
		{
			m.ip = 0x37c3
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x37c3:
		{
			m.ip = 0x37c5
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x37c5:
		{
			m.ip = 0x37c7
			m.r[6] = m.alu("add", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x37c7:
		{
			m.ip = 0x37c9
			m.r[5] = m.shift("shl", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x37c9:
		{
			m.ip = 0x37ce
			m.wr16(m.r[11], uint16(m.r[5]+0x22f), m.r[6])
			m.cycles += 8
		}
	case 0x37ce:
		{
			m.ip = 0x37d0
			m.r[5] = m.shift("shr", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x37d0:
		{
			m.ip = 0x37d6
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68)), uint16(0))
			m.cycles += 16
		}
	case 0x37d6:
		{
			m.ip = 0x37d8
			if !m.zf {
				m.ip = uint16(14339)
			}
			m.cycles += 8
		}
	case 0x37d8:
		{
			m.ip = 0x37db
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x37db:
		{
			m.ip = 0x37de
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x37de:
		{
			m.ip = 0x37e1
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x37e1:
		{
			m.ip = 0x37e5
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x37e5:
		{
			m.ip = 0x37e9
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x37e9:
		{
			m.ip = 0x37ed
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x37ed:
		{
			m.ip = 0x37f1
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x37f1:
		{
			m.ip = 0x37f5
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x37f5:
		{
			m.ip = 0x37f9
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x37f9:
		{
			m.ip = 0x37fc
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x37fc:
		{
			m.ip = 0x37ff
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x37ff:
		{
			m.ip = 0x3801
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14299)
			}
			m.cycles += 17
		}
	case 0x3801:
		{
			m.ip = 0x3802
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x3802:
		{
			m.ip = 0x3803
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3803:
		{
			m.ip = 0x3808
			m.wr8(m.r[11], uint16(m.r[5]+0x68), m.unary("dec", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68))))
			m.cycles += 15
		}
	case 0x3808:
		{
			m.ip = 0x380e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)), uint16(7))
			m.cycles += 16
		}
	case 0x380e:
		{
			m.ip = 0x3810
			if !m.zf {
				m.ip = uint16(14367)
			}
			m.cycles += 8
		}
	case 0x3810:
		{
			m.ip = 0x3814
			m.r[6] = uint16(0x70)
			m.cycles += 3
		}
	case 0x3814:
		{
			m.ip = 0x3816
			m.r[5] = m.shift("shl", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x3816:
		{
			m.ip = 0x381b
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(m.r[5]+0x6c)))
			m.cycles += 16
		}
	case 0x381b:
		{
			m.ip = 0x381d
			target := m.rd16(m.r[11], uint16(m.r[6]))
			m.push(0x381d)
			m.ip = target
			m.cycles += 19
		}
	case 0x381d:
		{
			m.ip = 0x381e
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x381e:
		{
			m.ip = 0x381f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x381f:
		{
			m.ip = 0x3822
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x3822:
		{
			m.ip = 0x3825
			m.r[3] = uint16(4)
			m.cycles += 2
		}
	case 0x3825:
		{
			m.ip = 0x382b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x201)), uint16(1))
			m.cycles += 16
		}
	case 0x382b:
		{
			m.ip = 0x382d
			if !m.cf && !m.zf {
				m.ip = uint16(14474)
			}
			m.cycles += 8
		}
	case 0x382d:
		{
			m.ip = 0x382f
			if m.zf {
				m.ip = uint16(14441)
			}
			m.cycles += 8
		}
	case 0x382f:
		{
			m.ip = 0x3831
			m.r[2] = m.r[1]
			m.cycles += 2
		}
	case 0x3831:
		{
			m.ip = 0x3836
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)))
			m.cycles += 8
		}
	case 0x3836:
		{
			m.ip = 0x3838
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x3838:
		{
			m.ip = 0x383e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68)), uint16(10))
			m.cycles += 16
		}
	case 0x383e:
		{
			m.ip = 0x3840
			if !m.cf && !m.zf {
				m.ip = uint16(14423)
			}
			m.cycles += 8
		}
	case 0x3840:
		{
			m.ip = 0x3842
			m.set8(1, 0, m.unary("neg", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x3842:
		{
			m.ip = 0x3845
			m.set8(1, 0, m.alu("add", 8, ((m.r[1]>>0)&255), uint16(8)))
			m.cycles += 4
		}
	case 0x3845:
		{
			m.ip = 0x3848
			m.set8(2, 0, m.alu("sub", 8, ((m.r[2]>>0)&255), uint16(3)))
			m.cycles += 4
		}
	case 0x3848:
		{
			m.ip = 0x384c
			m.r[6] = m.alu("add", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x384c:
		{
			m.ip = 0x3850
			m.r[7] = m.alu("add", 16, m.r[7], uint16(240))
			m.cycles += 4
		}
	case 0x3850:
		{
			m.ip = 0x3852
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14405)
			}
			m.cycles += 17
		}
	case 0x3852:
		{
			m.ip = 0x3854
			m.r[1] = m.r[2]
			m.cycles += 2
		}
	case 0x3854:
		{
			m.ip = 0x3857
			m.ip = uint16(14551)
			m.cycles += 15
		}
	case 0x3857:
		{
			m.ip = 0x385a
			m.set8(2, 0, m.alu("sub", 8, ((m.r[2]>>0)&255), uint16(3)))
			m.cycles += 4
		}
	case 0x385a:
		{
			m.ip = 0x385e
			m.r[6] = m.alu("add", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x385e:
		{
			m.ip = 0x3862
			m.r[7] = m.alu("add", 16, m.r[7], uint16(240))
			m.cycles += 4
		}
	case 0x3862:
		{
			m.ip = 0x3864
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14423)
			}
			m.cycles += 17
		}
	case 0x3864:
		{
			m.ip = 0x3866
			m.r[1] = m.r[2]
			m.cycles += 2
		}
	case 0x3866:
		{
			m.ip = 0x3868
			m.ip = uint16(14551)
			m.cycles += 15
		}
	case 0x3868:
		{
			m.ip = 0x3869
			m.cycles += 3
		}
	case 0x3869:
		{
			m.ip = 0x386e
			m.set8(2, 0, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)))
			m.cycles += 8
		}
	case 0x386e:
		{
			m.ip = 0x3870
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), ((m.r[2]>>0)&255)))
			m.cycles += 4
		}
	case 0x3870:
		{
			m.ip = 0x3875
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), m.rd8(m.r[11], uint16(m.r[5]+0x1a3))))
			m.cycles += 16
		}
	case 0x3875:
		{
			m.ip = 0x3878
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), uint16(3)))
			m.cycles += 4
		}
	case 0x3878:
		{
			m.ip = 0x387e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68)), uint16(10))
			m.cycles += 16
		}
	case 0x387e:
		{
			m.ip = 0x3880
			if !m.cf && !m.zf {
				m.ip = uint16(14469)
			}
			m.cycles += 8
		}
	case 0x3880:
		{
			m.ip = 0x3882
			m.set8(2, 0, m.unary("neg", 8, ((m.r[2]>>0)&255)))
			m.cycles += 3
		}
	case 0x3882:
		{
			m.ip = 0x3885
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), uint16(24)))
			m.cycles += 4
		}
	case 0x3885:
		{
			m.ip = 0x3887
			m.set8(1, 0, m.alu("sub", 8, ((m.r[1]>>0)&255), ((m.r[2]>>0)&255)))
			m.cycles += 4
		}
	case 0x3887:
		{
			m.ip = 0x3889
			m.ip = uint16(14551)
			m.cycles += 15
		}
	case 0x3889:
		{
			m.ip = 0x388a
			m.cycles += 3
		}
	case 0x388a:
		{
			m.ip = 0x3890
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x201)), uint16(3))
			m.cycles += 16
		}
	case 0x3890:
		{
			m.ip = 0x3892
			if m.zf {
				m.ip = uint16(14526)
			}
			m.cycles += 8
		}
	case 0x3892:
		{
			m.ip = 0x3897
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)))
			m.cycles += 8
		}
	case 0x3897:
		{
			m.ip = 0x3899
			m.set8(3, 0, m.shift("shr", 8, ((m.r[3]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x3899:
		{
			m.ip = 0x389b
			m.set8(3, 0, m.unary("inc", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x389b:
		{
			m.ip = 0x389d
			m.r[2] = m.r[3]
			m.cycles += 2
		}
	case 0x389d:
		{
			m.ip = 0x38a3
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68)), uint16(10))
			m.cycles += 16
		}
	case 0x38a3:
		{
			m.ip = 0x38a5
			if m.cf {
				m.ip = uint16(14513)
			}
			m.cycles += 8
		}
	case 0x38a5:
		{
			m.ip = 0x38a7
			m.set8(3, 0, m.unary("neg", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x38a7:
		{
			m.ip = 0x38aa
			m.set8(3, 0, m.alu("add", 8, ((m.r[3]>>0)&255), uint16(4)))
			m.cycles += 4
		}
	case 0x38aa:
		{
			m.ip = 0x38ac
			m.r[6] = m.alu("add", 16, m.r[6], m.r[2])
			m.cycles += 4
		}
	case 0x38ac:
		{
			m.ip = 0x38ae
			m.r[7] = m.alu("add", 16, m.r[7], m.r[2])
			m.cycles += 4
		}
	case 0x38ae:
		{
			m.ip = 0x38b0
			m.ip = uint16(14551)
			m.cycles += 15
		}
	case 0x38b0:
		{
			m.ip = 0x38b1
			m.cycles += 3
		}
	case 0x38b1:
		{
			m.ip = 0x38b4
			m.r[6] = m.alu("add", 16, m.r[6], uint16(4))
			m.cycles += 4
		}
	case 0x38b4:
		{
			m.ip = 0x38b6
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x38b6:
		{
			m.ip = 0x38b9
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x38b9:
		{
			m.ip = 0x38bb
			m.r[7] = m.alu("sub", 16, m.r[7], m.r[3])
			m.cycles += 4
		}
	case 0x38bb:
		{
			m.ip = 0x38bd
			m.ip = uint16(14551)
			m.cycles += 15
		}
	case 0x38bd:
		{
			m.ip = 0x38be
			m.cycles += 3
		}
	case 0x38be:
		{
			m.ip = 0x38c3
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)))
			m.cycles += 8
		}
	case 0x38c3:
		{
			m.ip = 0x38c5
			m.set8(3, 0, m.shift("shr", 8, ((m.r[3]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x38c5:
		{
			m.ip = 0x38c7
			m.set8(3, 0, m.unary("inc", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x38c7:
		{
			m.ip = 0x38cd
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68)), uint16(10))
			m.cycles += 16
		}
	case 0x38cd:
		{
			m.ip = 0x38cf
			if m.cf {
				m.ip = uint16(14551)
			}
			m.cycles += 8
		}
	case 0x38cf:
		{
			m.ip = 0x38d1
			m.set8(3, 0, m.unary("neg", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x38d1:
		{
			m.ip = 0x38d4
			m.set8(3, 0, m.alu("add", 8, ((m.r[3]>>0)&255), uint16(4)))
			m.cycles += 4
		}
	case 0x38d4:
		{
			m.ip = 0x38d6
			m.ip = uint16(14551)
			m.cycles += 15
		}
	case 0x38d6:
		{
			m.ip = 0x38d7
			m.cycles += 3
		}
	case 0x38d7:
		{
			m.ip = 0x38d8
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x38d8:
		{
			m.ip = 0x38da
			m.r[1] = m.r[3]
			m.cycles += 2
		}
	case 0x38da:
		{
			m.ip = 0x38dd
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x38dd:
		{
			m.ip = 0x38e0
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x38e0:
		{
			m.ip = 0x38e1
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x38e1:
		{
			m.ip = 0x38e2
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x38e2:
		{
			m.ip = 0x38e4
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14554)
			}
			m.cycles += 17
		}
	case 0x38e4:
		{
			m.ip = 0x38e7
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x38e7:
		{
			m.ip = 0x38e9
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x38e9:
		{
			m.ip = 0x38ec
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x38ec:
		{
			m.ip = 0x38ee
			m.r[7] = m.alu("sub", 16, m.r[7], m.r[3])
			m.cycles += 4
		}
	case 0x38ee:
		{
			m.ip = 0x38ef
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x38ef:
		{
			m.ip = 0x38f1
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14551)
			}
			m.cycles += 17
		}
	case 0x38f1:
		{
			m.ip = 0x38f2
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x38f2:
		{
			m.ip = 0x38f3
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x38f3:
		{
			m.ip = 0x38f4
			m.push(m.r[5])
			m.cycles += 11
		}
	case 0x38f4:
		{
			m.ip = 0x38f8
			m.r[5] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x38f8:
		{
			m.ip = 0x38fa
			m.set8(3, 8, uint16(0))
			m.cycles += 2
		}
	case 0x38fa:
		{
			m.ip = 0x38ff
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[5]+0x19f)))
			m.cycles += 8
		}
	case 0x38ff:
		{
			m.ip = 0x3901
			m.r[0] = m.r[3]
			m.cycles += 2
		}
	case 0x3901:
		{
			m.ip = 0x3903
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x3903:
		{
			m.ip = 0x3905
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x3905:
		{
			m.ip = 0x3907
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x3907:
		{
			m.ip = 0x3909
			m.set8(2, 8, uint16(0))
			m.cycles += 2
		}
	case 0x3909:
		{
			m.ip = 0x390e
			m.set8(2, 0, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)))
			m.cycles += 8
		}
	case 0x390e:
		{
			m.ip = 0x3911
			m.set8(2, 0, m.alu("and", 8, ((m.r[2]>>0)&255), uint16(254)))
			m.cycles += 4
		}
	case 0x3911:
		{
			m.ip = 0x3913
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x3913:
		{
			m.ip = 0x3915
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x3915:
		{
			m.ip = 0x3917
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x3917:
		{
			m.ip = 0x3919
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x3919:
		{
			m.ip = 0x391b
			m.r[5] = m.shift("shl", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x391b:
		{
			m.ip = 0x3920
			m.r[6] = m.rd16(m.r[11], uint16(m.r[5]+0x1a5))
			m.cycles += 8
		}
	case 0x3920:
		{
			m.ip = 0x3925
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[6]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x3925:
		{
			m.ip = 0x3928
			m.r[6] = uint16(57952)
			m.cycles += 2
		}
	case 0x3928:
		{
			m.ip = 0x392a
			if m.zf {
				m.ip = uint16(14637)
			}
			m.cycles += 8
		}
	case 0x392a:
		{
			m.ip = 0x392d
			m.r[6] = uint16(51040)
			m.cycles += 2
		}
	case 0x392d:
		{
			m.ip = 0x392f
			m.r[6] = m.alu("add", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x392f:
		{
			m.ip = 0x3934
			m.wr16(m.r[11], uint16(m.r[5]+0x22f), m.r[6])
			m.cycles += 8
		}
	case 0x3934:
		{
			m.ip = 0x3936
			m.r[5] = m.shift("shr", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x3936:
		{
			m.ip = 0x393c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68)), uint16(0))
			m.cycles += 16
		}
	case 0x393c:
		{
			m.ip = 0x393e
			if !m.zf {
				m.ip = uint16(14705)
			}
			m.cycles += 8
		}
	case 0x393e:
		{
			m.ip = 0x3941
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x3941:
		{
			m.ip = 0x3944
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x3944:
		{
			m.ip = 0x3947
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3947:
		{
			m.ip = 0x394b
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x394b:
		{
			m.ip = 0x394f
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x394f:
		{
			m.ip = 0x3953
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x3953:
		{
			m.ip = 0x3957
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3957:
		{
			m.ip = 0x395b
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x395b:
		{
			m.ip = 0x395f
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x395f:
		{
			m.ip = 0x3963
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x4)))
			m.cycles += 8
		}
	case 0x3963:
		{
			m.ip = 0x3967
			m.wr8(m.r[8], uint16(m.r[7]+0x4), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3967:
		{
			m.ip = 0x396a
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x396a:
		{
			m.ip = 0x396d
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x396d:
		{
			m.ip = 0x396f
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14657)
			}
			m.cycles += 17
		}
	case 0x396f:
		{
			m.ip = 0x3970
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x3970:
		{
			m.ip = 0x3971
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3971:
		{
			m.ip = 0x3976
			m.wr8(m.r[11], uint16(m.r[5]+0x68), m.unary("dec", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68))))
			m.cycles += 15
		}
	case 0x3976:
		{
			m.ip = 0x397c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)), uint16(7))
			m.cycles += 16
		}
	case 0x397c:
		{
			m.ip = 0x397e
			if !m.zf {
				m.ip = uint16(14735)
			}
			m.cycles += 8
		}
	case 0x397e:
		{
			m.ip = 0x3982
			m.r[6] = uint16(0x70)
			m.cycles += 3
		}
	case 0x3982:
		{
			m.ip = 0x3984
			m.r[5] = m.shift("shl", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x3984:
		{
			m.ip = 0x3989
			m.r[6] = m.alu("add", 16, m.r[6], m.rd16(m.r[11], uint16(m.r[5]+0x6c)))
			m.cycles += 16
		}
	case 0x3989:
		{
			m.ip = 0x398b
			m.r[5] = m.shift("shr", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x398b:
		{
			m.ip = 0x398d
			target := m.rd16(m.r[11], uint16(m.r[6]))
			m.push(0x398d)
			m.ip = target
			m.cycles += 19
		}
	case 0x398d:
		{
			m.ip = 0x398e
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x398e:
		{
			m.ip = 0x398f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x398f:
		{
			m.ip = 0x3992
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x3992:
		{
			m.ip = 0x3995
			m.r[3] = uint16(5)
			m.cycles += 2
		}
	case 0x3995:
		{
			m.ip = 0x399b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x201)), uint16(3))
			m.cycles += 16
		}
	case 0x399b:
		{
			m.ip = 0x399d
			if m.zf {
				m.ip = uint16(14793)
			}
			m.cycles += 8
		}
	case 0x399d:
		{
			m.ip = 0x39a2
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)))
			m.cycles += 8
		}
	case 0x39a2:
		{
			m.ip = 0x39a4
			m.set8(3, 0, m.shift("shr", 8, ((m.r[3]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x39a4:
		{
			m.ip = 0x39a6
			m.set8(3, 0, m.unary("inc", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x39a6:
		{
			m.ip = 0x39a8
			m.r[2] = m.r[3]
			m.cycles += 2
		}
	case 0x39a8:
		{
			m.ip = 0x39ae
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68)), uint16(10))
			m.cycles += 16
		}
	case 0x39ae:
		{
			m.ip = 0x39b0
			if m.cf {
				m.ip = uint16(14780)
			}
			m.cycles += 8
		}
	case 0x39b0:
		{
			m.ip = 0x39b2
			m.set8(3, 0, m.unary("neg", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x39b2:
		{
			m.ip = 0x39b5
			m.set8(3, 0, m.alu("add", 8, ((m.r[3]>>0)&255), uint16(5)))
			m.cycles += 4
		}
	case 0x39b5:
		{
			m.ip = 0x39b7
			m.r[6] = m.alu("add", 16, m.r[6], m.r[2])
			m.cycles += 4
		}
	case 0x39b7:
		{
			m.ip = 0x39b9
			m.r[7] = m.alu("add", 16, m.r[7], m.r[2])
			m.cycles += 4
		}
	case 0x39b9:
		{
			m.ip = 0x39bb
			m.ip = uint16(14818)
			m.cycles += 15
		}
	case 0x39bb:
		{
			m.ip = 0x39bc
			m.cycles += 3
		}
	case 0x39bc:
		{
			m.ip = 0x39bf
			m.r[6] = m.alu("add", 16, m.r[6], uint16(5))
			m.cycles += 4
		}
	case 0x39bf:
		{
			m.ip = 0x39c1
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x39c1:
		{
			m.ip = 0x39c4
			m.r[7] = m.alu("add", 16, m.r[7], uint16(5))
			m.cycles += 4
		}
	case 0x39c4:
		{
			m.ip = 0x39c6
			m.r[7] = m.alu("sub", 16, m.r[7], m.r[3])
			m.cycles += 4
		}
	case 0x39c6:
		{
			m.ip = 0x39c8
			m.ip = uint16(14818)
			m.cycles += 15
		}
	case 0x39c8:
		{
			m.ip = 0x39c9
			m.cycles += 3
		}
	case 0x39c9:
		{
			m.ip = 0x39ce
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[5]+0x1a3)))
			m.cycles += 8
		}
	case 0x39ce:
		{
			m.ip = 0x39d0
			m.set8(3, 0, m.shift("shr", 8, ((m.r[3]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x39d0:
		{
			m.ip = 0x39d2
			m.set8(3, 0, m.unary("inc", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x39d2:
		{
			m.ip = 0x39d8
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[5]+0x68)), uint16(8))
			m.cycles += 16
		}
	case 0x39d8:
		{
			m.ip = 0x39da
			if m.cf || m.zf {
				m.ip = uint16(14818)
			}
			m.cycles += 8
		}
	case 0x39da:
		{
			m.ip = 0x39dc
			m.set8(3, 0, m.unary("neg", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x39dc:
		{
			m.ip = 0x39df
			m.set8(3, 0, m.alu("add", 8, ((m.r[3]>>0)&255), uint16(5)))
			m.cycles += 4
		}
	case 0x39df:
		{
			m.ip = 0x39e1
			m.ip = uint16(14818)
			m.cycles += 15
		}
	case 0x39e1:
		{
			m.ip = 0x39e2
			m.cycles += 3
		}
	case 0x39e2:
		{
			m.ip = 0x39e3
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x39e3:
		{
			m.ip = 0x39e5
			m.r[1] = m.r[3]
			m.cycles += 2
		}
	case 0x39e5:
		{
			m.ip = 0x39e8
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x39e8:
		{
			m.ip = 0x39eb
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x39eb:
		{
			m.ip = 0x39ec
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x39ec:
		{
			m.ip = 0x39ed
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x39ed:
		{
			m.ip = 0x39ef
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14821)
			}
			m.cycles += 17
		}
	case 0x39ef:
		{
			m.ip = 0x39f2
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x39f2:
		{
			m.ip = 0x39f4
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[3])
			m.cycles += 4
		}
	case 0x39f4:
		{
			m.ip = 0x39f7
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x39f7:
		{
			m.ip = 0x39f9
			m.r[7] = m.alu("sub", 16, m.r[7], m.r[3])
			m.cycles += 4
		}
	case 0x39f9:
		{
			m.ip = 0x39fa
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x39fa:
		{
			m.ip = 0x39fc
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14818)
			}
			m.cycles += 17
		}
	case 0x39fc:
		{
			m.ip = 0x39fd
			m.r[5] = m.pop()
			m.cycles += 8
		}
	case 0x39fd:
		{
			m.ip = 0x39fe
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x39fe:
		{
			m.ip = 0x3a02
			m.r[6] = m.rd16(m.r[11], uint16(0x21f))
			m.cycles += 8
		}
	case 0x3a02:
		{
			m.ip = 0x3a06
			m.r[6] = m.alu("add", 16, m.r[6], uint16(1920))
			m.cycles += 4
		}
	case 0x3a06:
		{
			m.ip = 0x3a08
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x3a08:
		{
			m.ip = 0x3a0a
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3a0a:
		{
			m.ip = 0x3a0e
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)))
			m.cycles += 8
		}
	case 0x3a0e:
		{
			m.ip = 0x3a10
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3a10:
		{
			m.ip = 0x3a12
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x3a12:
		{
			m.ip = 0x3a16
			m.r[6] = m.alu("sub", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x3a16:
		{
			m.ip = 0x3a18
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14866)
			}
			m.cycles += 17
		}
	case 0x3a18:
		{
			m.ip = 0x3a1b
			m.r[1] = uint16(3)
			m.cycles += 2
		}
	case 0x3a1b:
		{
			m.ip = 0x3a1e
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x3a1e:
		{
			m.ip = 0x3a21
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a21:
		{
			m.ip = 0x3a25
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x3a25:
		{
			m.ip = 0x3a29
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a29:
		{
			m.ip = 0x3a2d
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x3a2d:
		{
			m.ip = 0x3a31
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a31:
		{
			m.ip = 0x3a35
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x3a35:
		{
			m.ip = 0x3a39
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a39:
		{
			m.ip = 0x3a3c
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x3a3c:
		{
			m.ip = 0x3a3f
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x3a3f:
		{
			m.ip = 0x3a41
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14875)
			}
			m.cycles += 17
		}
	case 0x3a41:
		{
			m.ip = 0x3a42
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3a42:
		{
			m.ip = 0x3a46
			m.r[6] = m.rd16(m.r[11], uint16(0x21f))
			m.cycles += 8
		}
	case 0x3a46:
		{
			m.ip = 0x3a48
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x3a48:
		{
			m.ip = 0x3a4a
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3a4a:
		{
			m.ip = 0x3a4e
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)))
			m.cycles += 8
		}
	case 0x3a4e:
		{
			m.ip = 0x3a50
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3a50:
		{
			m.ip = 0x3a52
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x3a52:
		{
			m.ip = 0x3a56
			m.r[6] = m.alu("sub", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x3a56:
		{
			m.ip = 0x3a5a
			m.r[6] = m.alu("add", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x3a5a:
		{
			m.ip = 0x3a5c
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14934)
			}
			m.cycles += 17
		}
	case 0x3a5c:
		{
			m.ip = 0x3a5f
			m.r[1] = uint16(3)
			m.cycles += 2
		}
	case 0x3a5f:
		{
			m.ip = 0x3a62
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x3a62:
		{
			m.ip = 0x3a65
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a65:
		{
			m.ip = 0x3a69
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x3a69:
		{
			m.ip = 0x3a6d
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a6d:
		{
			m.ip = 0x3a71
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x3a71:
		{
			m.ip = 0x3a75
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a75:
		{
			m.ip = 0x3a79
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x3a79:
		{
			m.ip = 0x3a7d
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a7d:
		{
			m.ip = 0x3a80
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x3a80:
		{
			m.ip = 0x3a83
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x3a83:
		{
			m.ip = 0x3a85
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14943)
			}
			m.cycles += 17
		}
	case 0x3a85:
		{
			m.ip = 0x3a86
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3a86:
		{
			m.ip = 0x3a8a
			m.r[6] = m.rd16(m.r[11], uint16(0x21f))
			m.cycles += 8
		}
	case 0x3a8a:
		{
			m.ip = 0x3a8c
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x3a8c:
		{
			m.ip = 0x3a90
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)))
			m.cycles += 8
		}
	case 0x3a90:
		{
			m.ip = 0x3a92
			m.set8(1, 0, m.shift("shr", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x3a92:
		{
			m.ip = 0x3a94
			m.r[6] = m.alu("add", 16, m.r[6], m.r[1])
			m.cycles += 4
		}
	case 0x3a94:
		{
			m.ip = 0x3a97
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x3a97:
		{
			m.ip = 0x3a9a
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x3a9a:
		{
			m.ip = 0x3a9d
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3a9d:
		{
			m.ip = 0x3aa0
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x3aa0:
		{
			m.ip = 0x3aa3
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x3aa3:
		{
			m.ip = 0x3aa5
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(14999)
			}
			m.cycles += 17
		}
	case 0x3aa5:
		{
			m.ip = 0x3aa9
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x3aa9:
		{
			m.ip = 0x3aaa
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3aaa:
		{
			m.ip = 0x3aae
			m.r[6] = m.rd16(m.r[11], uint16(0x21f))
			m.cycles += 8
		}
	case 0x3aae:
		{
			m.ip = 0x3ab1
			m.r[6] = m.alu("add", 16, m.r[6], uint16(3))
			m.cycles += 4
		}
	case 0x3ab1:
		{
			m.ip = 0x3ab3
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x3ab3:
		{
			m.ip = 0x3ab7
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1a3)))
			m.cycles += 8
		}
	case 0x3ab7:
		{
			m.ip = 0x3ab9
			m.set8(1, 0, m.shift("shr", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x3ab9:
		{
			m.ip = 0x3abb
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[1])
			m.cycles += 4
		}
	case 0x3abb:
		{
			m.ip = 0x3abe
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x3abe:
		{
			m.ip = 0x3ac1
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x3ac1:
		{
			m.ip = 0x3ac5
			m.wr8(m.r[8], uint16(m.r[7]+0x4), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3ac5:
		{
			m.ip = 0x3ac8
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x3ac8:
		{
			m.ip = 0x3acb
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x3acb:
		{
			m.ip = 0x3acd
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(15038)
			}
			m.cycles += 17
		}
	case 0x3acd:
		{
			m.ip = 0x3ad1
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x3ad1:
		{
			m.ip = 0x3ad2
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3ad2:
		{
			m.ip = 0x3ad6
			m.r[3] = m.rd16(m.r[11], uint16(0x221))
			m.cycles += 8
		}
	case 0x3ad6:
		{
			m.ip = 0x3ad8
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3ad8:
		{
			m.ip = 0x3add
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5)), uint16(0))
			m.cycles += 16
		}
	case 0x3add:
		{
			m.ip = 0x3adf
			if m.zf {
				m.ip = uint16(15083)
			}
			m.cycles += 8
		}
	case 0x3adf:
		{
			m.ip = 0x3ae3
			m.wr16(m.r[11], uint16(m.r[3]+0x1d5), m.unary("dec", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5))))
			m.cycles += 15
		}
	case 0x3ae3:
		{
			m.ip = 0x3ae8
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5)), uint16(23))
			m.cycles += 16
		}
	case 0x3ae8:
		{
			m.ip = 0x3aea
			if m.cf || m.zf {
				m.ip = uint16(15083)
			}
			m.cycles += 8
		}
	case 0x3aea:
		{
			m.ip = 0x3aeb
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3aeb:
		{
			m.ip = 0x3aed
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3aed:
		{
			m.ip = 0x3af1
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b5)))
			m.cycles += 8
		}
	case 0x3af1:
		{
			m.ip = 0x3af3
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3af3:
		{
			m.ip = 0x3af5
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(255))
			m.cycles += 4
		}
	case 0x3af5:
		{
			m.ip = 0x3af7
			if !m.zf {
				m.ip = uint16(15098)
			}
			m.cycles += 8
		}
	case 0x3af7:
		{
			m.ip = 0x3afa
			m.ip = uint16(15404)
			m.cycles += 15
		}
	case 0x3afa:
		{
			m.ip = 0x3b00
			m.wr16(m.r[11], uint16(0x21f), uint16(28064))
			m.cycles += 8
		}
	case 0x3b00:
		{
			m.ip = 0x3b04
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x1c5))
			m.cycles += 8
		}
	case 0x3b04:
		{
			m.ip = 0x3b06
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3b06:
		{
			m.ip = 0x3b0b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)), uint16(7))
			m.cycles += 16
		}
	case 0x3b0b:
		{
			m.ip = 0x3b0d
			if !m.zf {
				m.ip = uint16(15125)
			}
			m.cycles += 8
		}
	case 0x3b0d:
		{
			m.ip = 0x3b0f
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3b0f:
		{
			m.ip = 0x3b13
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x1bd))
			m.cycles += 8
		}
	case 0x3b13:
		{
			m.ip = 0x3b15
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3b15:
		{
			m.ip = 0x3b1b
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(4))
			m.cycles += 16
		}
	case 0x3b1b:
		{
			m.ip = 0x3b1d
			if m.zf {
				m.ip = uint16(15139)
			}
			m.cycles += 8
		}
	case 0x3b1d:
		{
			m.ip = 0x3b20
			target := uint16(15549)
			m.push(0x3b20)
			m.ip = target
			m.cycles += 19
		}
	case 0x3b20:
		{
			m.ip = 0x3b22
			m.ip = uint16(15153)
			m.cycles += 15
		}
	case 0x3b22:
		{
			m.ip = 0x3b23
			m.cycles += 3
		}
	case 0x3b23:
		{
			m.ip = 0x3b29
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x3b29:
		{
			m.ip = 0x3b2b
			if m.zf {
				m.ip = uint16(15153)
			}
			m.cycles += 8
		}
	case 0x3b2b:
		{
			m.ip = 0x3b31
			m.wr16(m.r[11], uint16(0x21f), uint16(28068))
			m.cycles += 8
		}
	case 0x3b31:
		{
			m.ip = 0x3b33
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3b33:
		{
			m.ip = 0x3b37
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x1bd))
			m.cycles += 8
		}
	case 0x3b37:
		{
			m.ip = 0x3b39
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3b39:
		{
			m.ip = 0x3b3d
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x3b3d:
		{
			m.ip = 0x3b3f
			m.set8(0, 0, m.unary("inc", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x3b3f:
		{
			m.ip = 0x3b41
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(7)))
			m.cycles += 4
		}
	case 0x3b41:
		{
			m.ip = 0x3b45
			m.wr8(m.r[11], uint16(m.r[3]+0x1b9), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3b45:
		{
			m.ip = 0x3b49
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x3b49:
		{
			m.ip = 0x3b4b
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x3b4b:
		{
			m.ip = 0x3b4d
			if m.zf {
				m.ip = uint16(15188)
			}
			m.cycles += 8
		}
	case 0x3b4d:
		{
			m.ip = 0x3b51
			m.set8(2, 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b1)))
			m.cycles += 8
		}
	case 0x3b51:
		{
			m.ip = 0x3b53
			m.ip = uint16(15247)
			m.cycles += 15
		}
	case 0x3b53:
		{
			m.ip = 0x3b54
			m.cycles += 3
		}
	case 0x3b54:
		{
			m.ip = 0x3b58
			m.set8(2, 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b5)))
			m.cycles += 8
		}
	case 0x3b58:
		{
			m.ip = 0x3b5c
			m.wr8(m.r[11], uint16(m.r[3]+0x1b1), ((m.r[2] >> 8) & 255))
			m.cycles += 8
		}
	case 0x3b5c:
		{
			m.ip = 0x3b5e
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3b5e:
		{
			m.ip = 0x3b62
			m.wr16(m.r[11], uint16(m.r[3]+0x1c5), m.r[5])
			m.cycles += 8
		}
	case 0x3b62:
		{
			m.ip = 0x3b65
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x3b65:
		{
			m.ip = 0x3b67
			if !m.zf {
				m.ip = uint16(15214)
			}
			m.cycles += 8
		}
	case 0x3b67:
		{
			m.ip = 0x3b6b
			m.wr16(m.r[11], uint16(m.r[3]+0x1bd), m.unary("inc", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1bd))))
			m.cycles += 15
		}
	case 0x3b6b:
		{
			m.ip = 0x3b6d
			m.ip = uint16(15336)
			m.cycles += 15
		}
	case 0x3b6d:
		{
			m.ip = 0x3b6e
			m.cycles += 3
		}
	case 0x3b6e:
		{
			m.ip = 0x3b71
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(8))
			m.cycles += 4
		}
	case 0x3b71:
		{
			m.ip = 0x3b73
			if !m.zf {
				m.ip = uint16(15226)
			}
			m.cycles += 8
		}
	case 0x3b73:
		{
			m.ip = 0x3b77
			m.wr16(m.r[11], uint16(m.r[3]+0x1bd), m.unary("dec", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1bd))))
			m.cycles += 15
		}
	case 0x3b77:
		{
			m.ip = 0x3b7a
			m.ip = uint16(15370)
			m.cycles += 15
		}
	case 0x3b7a:
		{
			m.ip = 0x3b7d
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x3b7d:
		{
			m.ip = 0x3b7f
			if !m.zf {
				m.ip = uint16(15239)
			}
			m.cycles += 8
		}
	case 0x3b7f:
		{
			m.ip = 0x3b84
			m.wr16(m.r[11], uint16(m.r[3]+0x1bd), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1bd)), uint16(21)))
			m.cycles += 16
		}
	case 0x3b84:
		{
			m.ip = 0x3b86
			m.ip = uint16(15301)
			m.cycles += 15
		}
	case 0x3b86:
		{
			m.ip = 0x3b87
			m.cycles += 3
		}
	case 0x3b87:
		{
			m.ip = 0x3b8c
			m.wr16(m.r[11], uint16(m.r[3]+0x1bd), m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1bd)), uint16(21)))
			m.cycles += 16
		}
	case 0x3b8c:
		{
			m.ip = 0x3b8e
			m.ip = uint16(15273)
			m.cycles += 15
		}
	case 0x3b8e:
		{
			m.ip = 0x3b8f
			m.cycles += 3
		}
	case 0x3b8f:
		{
			m.ip = 0x3b91
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3b91:
		{
			m.ip = 0x3b94
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(0))
			m.cycles += 4
		}
	case 0x3b94:
		{
			m.ip = 0x3b96
			if !m.zf {
				m.ip = uint16(15257)
			}
			m.cycles += 8
		}
	case 0x3b96:
		{
			m.ip = 0x3b98
			m.ip = uint16(15336)
			m.cycles += 15
		}
	case 0x3b98:
		{
			m.ip = 0x3b99
			m.cycles += 3
		}
	case 0x3b99:
		{
			m.ip = 0x3b9c
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(8))
			m.cycles += 4
		}
	case 0x3b9c:
		{
			m.ip = 0x3b9e
			if !m.zf {
				m.ip = uint16(15265)
			}
			m.cycles += 8
		}
	case 0x3b9e:
		{
			m.ip = 0x3ba0
			m.ip = uint16(15370)
			m.cycles += 15
		}
	case 0x3ba0:
		{
			m.ip = 0x3ba1
			m.cycles += 3
		}
	case 0x3ba1:
		{
			m.ip = 0x3ba4
			m.alu("sub", 8, ((m.r[2] >> 8) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x3ba4:
		{
			m.ip = 0x3ba6
			if !m.zf {
				m.ip = uint16(15273)
			}
			m.cycles += 8
		}
	case 0x3ba6:
		{
			m.ip = 0x3ba8
			m.ip = uint16(15301)
			m.cycles += 15
		}
	case 0x3ba8:
		{
			m.ip = 0x3ba9
			m.cycles += 3
		}
	case 0x3ba9:
		{
			m.ip = 0x3bad
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1cd))
			m.cycles += 8
		}
	case 0x3bad:
		{
			m.ip = 0x3bb1
			m.r[7] = m.alu("add", 16, m.r[7], uint16(1680))
			m.cycles += 4
		}
	case 0x3bb1:
		{
			m.ip = 0x3bb3
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3bb3:
		{
			m.ip = 0x3bb6
			target := uint16(16569)
			m.push(0x3bb6)
			m.ip = target
			m.cycles += 19
		}
	case 0x3bb6:
		{
			m.ip = 0x3bba
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(2160))
			m.cycles += 4
		}
	case 0x3bba:
		{
			m.ip = 0x3bbc
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3bbc:
		{
			m.ip = 0x3bc0
			m.wr16(m.r[11], uint16(m.r[3]+0x1cd), m.r[7])
			m.cycles += 8
		}
	case 0x3bc0:
		{
			m.ip = 0x3bc2
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3bc2:
		{
			m.ip = 0x3bc4
			m.ip = uint16(15423)
			m.cycles += 15
		}
	case 0x3bc4:
		{
			m.ip = 0x3bc5
			m.cycles += 3
		}
	case 0x3bc5:
		{
			m.ip = 0x3bc9
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1cd))
			m.cycles += 8
		}
	case 0x3bc9:
		{
			m.ip = 0x3bcd
			m.r[7] = m.alu("add", 16, m.r[7], uint16(240))
			m.cycles += 4
		}
	case 0x3bcd:
		{
			m.ip = 0x3bd2
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x1d5)), uint16(15))
			m.cycles += 16
		}
	case 0x3bd2:
		{
			m.ip = 0x3bd4
			if !m.cf && !m.zf {
				m.ip = uint16(15327)
			}
			m.cycles += 8
		}
	case 0x3bd4:
		{
			m.ip = 0x3bd8
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(240))
			m.cycles += 4
		}
	case 0x3bd8:
		{
			m.ip = 0x3bda
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3bda:
		{
			m.ip = 0x3bdd
			target := uint16(16633)
			m.push(0x3bdd)
			m.ip = target
			m.cycles += 19
		}
	case 0x3bdd:
		{
			m.ip = 0x3bdf
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3bdf:
		{
			m.ip = 0x3be3
			m.wr16(m.r[11], uint16(m.r[3]+0x1cd), m.r[7])
			m.cycles += 8
		}
	case 0x3be3:
		{
			m.ip = 0x3be5
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3be5:
		{
			m.ip = 0x3be7
			m.ip = uint16(15423)
			m.cycles += 15
		}
	case 0x3be7:
		{
			m.ip = 0x3be8
			m.cycles += 3
		}
	case 0x3be8:
		{
			m.ip = 0x3bec
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1cd))
			m.cycles += 8
		}
	case 0x3bec:
		{
			m.ip = 0x3bee
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3bee:
		{
			m.ip = 0x3bf3
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)), uint16(1))
			m.cycles += 16
		}
	case 0x3bf3:
		{
			m.ip = 0x3bf5
			if !m.zf {
				m.ip = uint16(15355)
			}
			m.cycles += 8
		}
	case 0x3bf5:
		{
			m.ip = 0x3bf8
			target := uint16(16478)
			m.push(0x3bf8)
			m.ip = target
			m.cycles += 19
		}
	case 0x3bf8:
		{
			m.ip = 0x3bfa
			m.ip = uint16(15426)
			m.cycles += 15
		}
	case 0x3bfa:
		{
			m.ip = 0x3bfb
			m.cycles += 3
		}
	case 0x3bfb:
		{
			m.ip = 0x3bfe
			target := uint16(16697)
			m.push(0x3bfe)
			m.ip = target
			m.cycles += 19
		}
	case 0x3bfe:
		{
			m.ip = 0x3bff
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x3bff:
		{
			m.ip = 0x3c01
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3c01:
		{
			m.ip = 0x3c05
			m.wr16(m.r[11], uint16(m.r[3]+0x1cd), m.r[7])
			m.cycles += 8
		}
	case 0x3c05:
		{
			m.ip = 0x3c07
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3c07:
		{
			m.ip = 0x3c09
			m.ip = uint16(15423)
			m.cycles += 15
		}
	case 0x3c09:
		{
			m.ip = 0x3c0a
			m.cycles += 3
		}
	case 0x3c0a:
		{
			m.ip = 0x3c0e
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1cd))
			m.cycles += 8
		}
	case 0x3c0e:
		{
			m.ip = 0x3c10
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3c10:
		{
			m.ip = 0x3c15
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)), uint16(1))
			m.cycles += 16
		}
	case 0x3c15:
		{
			m.ip = 0x3c17
			if !m.zf {
				m.ip = uint16(15398)
			}
			m.cycles += 8
		}
	case 0x3c17:
		{
			m.ip = 0x3c18
			m.r[7] = m.unary("dec", 16, m.r[7])
			m.cycles += 3
		}
	case 0x3c18:
		{
			m.ip = 0x3c1a
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3c1a:
		{
			m.ip = 0x3c1e
			m.wr16(m.r[11], uint16(m.r[3]+0x1cd), m.r[7])
			m.cycles += 8
		}
	case 0x3c1e:
		{
			m.ip = 0x3c20
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3c20:
		{
			m.ip = 0x3c23
			target := uint16(16478)
			m.push(0x3c23)
			m.ip = target
			m.cycles += 19
		}
	case 0x3c23:
		{
			m.ip = 0x3c25
			m.ip = uint16(15426)
			m.cycles += 15
		}
	case 0x3c25:
		{
			m.ip = 0x3c26
			m.cycles += 3
		}
	case 0x3c26:
		{
			m.ip = 0x3c29
			target := uint16(16733)
			m.push(0x3c29)
			m.ip = target
			m.cycles += 19
		}
	case 0x3c29:
		{
			m.ip = 0x3c2b
			m.ip = uint16(15423)
			m.cycles += 15
		}
	case 0x3c2b:
		{
			m.ip = 0x3c2c
			m.cycles += 3
		}
	case 0x3c2c:
		{
			m.ip = 0x3c30
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1cd))
			m.cycles += 8
		}
	case 0x3c30:
		{
			m.ip = 0x3c32
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3c32:
		{
			m.ip = 0x3c37
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)), uint16(1))
			m.cycles += 16
		}
	case 0x3c37:
		{
			m.ip = 0x3c39
			if !m.zf {
				m.ip = uint16(15423)
			}
			m.cycles += 8
		}
	case 0x3c39:
		{
			m.ip = 0x3c3c
			target := uint16(16478)
			m.push(0x3c3c)
			m.ip = target
			m.cycles += 19
		}
	case 0x3c3c:
		{
			m.ip = 0x3c3e
			m.ip = uint16(15426)
			m.cycles += 15
		}
	case 0x3c3e:
		{
			m.ip = 0x3c3f
			m.cycles += 3
		}
	case 0x3c3f:
		{
			m.ip = 0x3c42
			target := uint16(16300)
			m.push(0x3c42)
			m.ip = target
			m.cycles += 19
		}
	case 0x3c42:
		{
			m.ip = 0x3c46
			m.r[1] = m.rd16(m.r[11], uint16(0x1ad))
			m.cycles += 8
		}
	case 0x3c46:
		{
			m.ip = 0x3c4c
			m.wr16(m.r[11], uint16(0x223), uint16(0))
			m.cycles += 8
		}
	case 0x3c4c:
		{
			m.ip = 0x3c52
			m.wr16(m.r[11], uint16(0x225), uint16(0))
			m.cycles += 8
		}
	case 0x3c52:
		{
			m.ip = 0x3c55
			target := uint16(15472)
			m.push(0x3c55)
			m.ip = target
			m.cycles += 19
		}
	case 0x3c55:
		{
			m.ip = 0x3c5a
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x3c5a:
		{
			m.ip = 0x3c5c
			if !m.zf {
				m.ip = uint16(15471)
			}
			m.cycles += 8
		}
	case 0x3c5c:
		{
			m.ip = 0x3c60
			m.r[1] = m.rd16(m.r[11], uint16(0x1af))
			m.cycles += 8
		}
	case 0x3c60:
		{
			m.ip = 0x3c66
			m.wr16(m.r[11], uint16(0x223), uint16(1))
			m.cycles += 8
		}
	case 0x3c66:
		{
			m.ip = 0x3c6c
			m.wr16(m.r[11], uint16(0x225), uint16(2))
			m.cycles += 8
		}
	case 0x3c6c:
		{
			m.ip = 0x3c6f
			target := uint16(15472)
			m.push(0x3c6f)
			m.ip = target
			m.cycles += 19
		}
	case 0x3c6f:
		{
			m.ip = 0x3c70
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3c70:
		{
			m.ip = 0x3c74
			m.r[3] = m.rd16(m.r[11], uint16(0x221))
			m.cycles += 8
		}
	case 0x3c74:
		{
			m.ip = 0x3c76
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3c76:
		{
			m.ip = 0x3c7a
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x1cd))
			m.cycles += 8
		}
	case 0x3c7a:
		{
			m.ip = 0x3c7d
			m.r[0] = m.alu("add", 16, m.r[0], uint16(1683))
			m.cycles += 4
		}
	case 0x3c7d:
		{
			m.ip = 0x3c7f
			m.r[0] = m.alu("sub", 16, m.r[0], m.r[1])
			m.cycles += 4
		}
	case 0x3c7f:
		{
			m.ip = 0x3c82
			m.alu("sub", 16, m.r[0], uint16(3360))
			m.cycles += 4
		}
	case 0x3c82:
		{
			m.ip = 0x3c84
			if m.cf || m.zf {
				m.ip = uint16(15493)
			}
			m.cycles += 8
		}
	case 0x3c84:
		{
			m.ip = 0x3c85
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3c85:
		{
			m.ip = 0x3c88
			m.alu("sub", 16, m.r[0], uint16(6))
			m.cycles += 4
		}
	case 0x3c88:
		{
			m.ip = 0x3c8a
			if m.cf || m.zf {
				m.ip = uint16(15532)
			}
			m.cycles += 8
		}
	case 0x3c8a:
		{
			m.ip = 0x3c8d
			m.alu("sub", 16, m.r[0], uint16(1680))
			m.cycles += 4
		}
	case 0x3c8d:
		{
			m.ip = 0x3c8f
			if m.zf {
				m.ip = uint16(15532)
			}
			m.cycles += 8
		}
	case 0x3c8f:
		{
			m.ip = 0x3c92
			m.alu("sub", 16, m.r[0], uint16(1686))
			m.cycles += 4
		}
	case 0x3c92:
		{
			m.ip = 0x3c94
			if m.zf {
				m.ip = uint16(15532)
			}
			m.cycles += 8
		}
	case 0x3c94:
		{
			m.ip = 0x3c97
			m.r[2] = uint16(241)
			m.cycles += 2
		}
	case 0x3c97:
		{
			m.ip = 0x3c9a
			m.r[1] = uint16(14)
			m.cycles += 2
		}
	case 0x3c9a:
		{
			m.ip = 0x3c9c
			m.alu("sub", 16, m.r[0], m.r[2])
			m.cycles += 4
		}
	case 0x3c9c:
		{
			m.ip = 0x3c9e
			if m.cf {
				m.ip = uint16(15531)
			}
			m.cycles += 8
		}
	case 0x3c9e:
		{
			m.ip = 0x3ca1
			m.r[2] = m.alu("add", 16, m.r[2], uint16(4))
			m.cycles += 4
		}
	case 0x3ca1:
		{
			m.ip = 0x3ca3
			m.alu("sub", 16, m.r[0], m.r[2])
			m.cycles += 4
		}
	case 0x3ca3:
		{
			m.ip = 0x3ca5
			if m.cf {
				m.ip = uint16(15532)
			}
			m.cycles += 8
		}
	case 0x3ca5:
		{
			m.ip = 0x3ca9
			m.r[2] = m.alu("add", 16, m.r[2], uint16(236))
			m.cycles += 4
		}
	case 0x3ca9:
		{
			m.ip = 0x3cab
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(15514)
			}
			m.cycles += 17
		}
	case 0x3cab:
		{
			m.ip = 0x3cac
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3cac:
		{
			m.ip = 0x3cae
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3cae:
		{
			m.ip = 0x3cb3
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x7a)), uint16(77))
			m.cycles += 16
		}
	case 0x3cb3:
		{
			m.ip = 0x3cb5
			if m.zf {
				m.ip = uint16(15545)
			}
			m.cycles += 8
		}
	case 0x3cb5:
		{
			m.ip = 0x3cb8
			target := uint16(15579)
			m.push(0x3cb8)
			m.ip = target
			m.cycles += 19
		}
	case 0x3cb8:
		{
			m.ip = 0x3cb9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3cb9:
		{
			m.ip = 0x3cbc
			target := uint16(15993)
			m.push(0x3cbc)
			m.ip = target
			m.cycles += 19
		}
	case 0x3cbc:
		{
			m.ip = 0x3cbd
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3cbd:
		{
			m.ip = 0x3cbe
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x3cbe:
		{
			m.ip = 0x3cc1
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x3cc1:
		{
			m.ip = 0x3cc5
			m.alu("sub", 16, m.r[5], m.rd16(m.r[11], uint16(m.r[3]+0x0)))
			m.cycles += 16
		}
	case 0x3cc5:
		{
			m.ip = 0x3cc7
			if m.zf {
				m.ip = uint16(15569)
			}
			m.cycles += 8
		}
	case 0x3cc7:
		{
			m.ip = 0x3cca
			m.r[3] = m.alu("add", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x3cca:
		{
			m.ip = 0x3ccd
			m.alu("sub", 16, m.r[3], uint16(24))
			m.cycles += 4
		}
	case 0x3ccd:
		{
			m.ip = 0x3ccf
			if !m.zf {
				m.ip = uint16(15553)
			}
			m.cycles += 8
		}
	case 0x3ccf:
		{
			m.ip = 0x3cd0
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x3cd0:
		{
			m.ip = 0x3cd1
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3cd1:
		{
			m.ip = 0x3cd5
			m.r[3] = m.rd16(m.r[11], uint16(m.r[3]+0x2))
			m.cycles += 8
		}
	case 0x3cd5:
		{
			m.ip = 0x3cd9
			m.wr16(m.r[11], uint16(0x21f), m.r[3])
			m.cycles += 8
		}
	case 0x3cd9:
		{
			m.ip = 0x3cda
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x3cda:
		{
			m.ip = 0x3cdb
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3cdb:
		{
			m.ip = 0x3cde
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x3cde:
		{
			m.ip = 0x3ce1
			target := uint16(9487)
			m.push(0x3ce1)
			m.ip = target
			m.cycles += 19
		}
	case 0x3ce1:
		{
			m.ip = 0x3ce5
			m.r[6] = m.rd16(m.r[11], uint16(0x466))
			m.cycles += 8
		}
	case 0x3ce5:
		{
			m.ip = 0x3ce9
			m.r[6] = m.rd16(m.r[11], uint16(m.r[6]+0xd93))
			m.cycles += 8
		}
	case 0x3ce9:
		{
			m.ip = 0x3ced
			m.wr16(m.r[11], uint16(0x21d), m.r[6])
			m.cycles += 8
		}
	case 0x3ced:
		{
			m.ip = 0x3cf0
			m.r[6] = m.alu("add", 16, m.r[6], uint16(18))
			m.cycles += 4
		}
	case 0x3cf0:
		{
			m.ip = 0x3cf4
			m.r[7] = uint16(0x1b1)
			m.cycles += 3
		}
	case 0x3cf4:
		{
			m.ip = 0x3cf7
			m.r[1] = uint16(44)
			m.cycles += 2
		}
	case 0x3cf7:
		{
			m.ip = 0x3cfa
			target := uint16(9392)
			m.push(0x3cfa)
			m.ip = target
			m.cycles += 19
		}
	case 0x3cfa:
		{
			m.ip = 0x3cfe
			m.r[6] = uint16(0xff5)
			m.cycles += 3
		}
	case 0x3cfe:
		{
			m.ip = 0x3d01
			m.r[3] = uint16(65535)
			m.cycles += 2
		}
	case 0x3d01:
		{
			m.ip = 0x3d04
			target := uint16(9346)
			m.push(0x3d04)
			m.ip = target
			m.cycles += 19
		}
	case 0x3d04:
		{
			m.ip = 0x3d0a
			m.wr16(m.r[11], uint16(0x55b), uint16(0))
			m.cycles += 8
		}
	case 0x3d0a:
		{
			m.ip = 0x3d0e
			m.r[0] = uint16(0x5ab)
			m.cycles += 3
		}
	case 0x3d0e:
		{
			m.ip = 0x3d11
			m.wr16(m.r[11], uint16(0x55d), m.r[0])
			m.cycles += 8
		}
	case 0x3d11:
		{
			m.ip = 0x3d14
			m.r[6] = uint16(45280)
			m.cycles += 2
		}
	case 0x3d14:
		{
			m.ip = 0x3d18
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3d18:
		{
			m.ip = 0x3d1c
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1ad))
			m.cycles += 8
		}
	case 0x3d1c:
		{
			m.ip = 0x3d1f
			m.r[1] = uint16(21)
			m.cycles += 2
		}
	case 0x3d1f:
		{
			m.ip = 0x3d21
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3d21:
		{
			m.ip = 0x3d26
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x6a)), uint16(0))
			m.cycles += 16
		}
	case 0x3d26:
		{
			m.ip = 0x3d28
			if m.zf {
				m.ip = uint16(15684)
			}
			m.cycles += 8
		}
	case 0x3d28:
		{
			m.ip = 0x3d2a
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3d2a:
		{
			m.ip = 0x3d2e
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x6c))
			m.cycles += 8
		}
	case 0x3d2e:
		{
			m.ip = 0x3d30
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3d30:
		{
			m.ip = 0x3d35
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0x68)), uint16(10))
			m.cycles += 16
		}
	case 0x3d35:
		{
			m.ip = 0x3d37
			if !m.cf && !m.zf {
				m.ip = uint16(15677)
			}
			m.cycles += 8
		}
	case 0x3d37:
		{
			m.ip = 0x3d3a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x3d3a:
		{
			m.ip = 0x3d3d
			m.r[7] = m.alu("and", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x3d3d:
		{
			m.ip = 0x3d40
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x3d40:
		{
			m.ip = 0x3d44
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x1f5))
			m.cycles += 8
		}
	case 0x3d44:
		{
			m.ip = 0x3d45
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x3d45:
		{
			m.ip = 0x3d46
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x3d46:
		{
			m.ip = 0x3d47
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x3d47:
		{
			m.ip = 0x3d4a
			target := uint16(10677)
			m.push(0x3d4a)
			m.ip = target
			m.cycles += 19
		}
	case 0x3d4a:
		{
			m.ip = 0x3d4d
			target := uint16(9749)
			m.push(0x3d4d)
			m.ip = target
			m.cycles += 19
		}
	case 0x3d4d:
		{
			m.ip = 0x3d50
			target := uint16(11325)
			m.push(0x3d50)
			m.ip = target
			m.cycles += 19
		}
	case 0x3d50:
		{
			m.ip = 0x3d56
			m.wr16(m.r[11], uint16(0x106a), uint16(20))
			m.cycles += 8
		}
	case 0x3d56:
		{
			m.ip = 0x3d59
			target := uint16(19182)
			m.push(0x3d59)
			m.ip = target
			m.cycles += 19
		}
	case 0x3d59:
		{
			m.ip = 0x3d5a
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x3d5a:
		{
			m.ip = 0x3d5b
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x3d5b:
		{
			m.ip = 0x3d5c
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x3d5c:
		{
			m.ip = 0x3d5f
			m.alu("sub", 16, m.r[1], uint16(16))
			m.cycles += 4
		}
	case 0x3d5f:
		{
			m.ip = 0x3d61
			if !m.cf && !m.zf {
				m.ip = uint16(15716)
			}
			m.cycles += 8
		}
	case 0x3d61:
		{
			m.ip = 0x3d64
			m.r[6] = m.alu("add", 16, m.r[6], uint16(4))
			m.cycles += 4
		}
	case 0x3d64:
		{
			m.ip = 0x3d66
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(15684)
			}
			m.cycles += 17
		}
	case 0x3d66:
		{
			m.ip = 0x3d6c
			m.wr16(m.r[11], uint16(0x106a), uint16(1500))
			m.cycles += 8
		}
	case 0x3d6c:
		{
			m.ip = 0x3d6f
			target := uint16(19311)
			m.push(0x3d6f)
			m.ip = target
			m.cycles += 19
		}
	case 0x3d6f:
		{
			m.ip = 0x3d72
			m.r[6] = uint16(28064)
			m.cycles += 2
		}
	case 0x3d72:
		{
			m.ip = 0x3d76
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3d76:
		{
			m.ip = 0x3d7a
			m.r[3] = m.rd16(m.r[11], uint16(m.r[3]+0x1a5))
			m.cycles += 8
		}
	case 0x3d7a:
		{
			m.ip = 0x3d7f
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x3d7f:
		{
			m.ip = 0x3d81
			if m.zf {
				m.ip = uint16(15748)
			}
			m.cycles += 8
		}
	case 0x3d81:
		{
			m.ip = 0x3d84
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x3d84:
		{
			m.ip = 0x3d87
			target := uint16(10677)
			m.push(0x3d87)
			m.ip = target
			m.cycles += 19
		}
	case 0x3d87:
		{
			m.ip = 0x3d8c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x37)), uint16(0))
			m.cycles += 16
		}
	case 0x3d8c:
		{
			m.ip = 0x3d8e
			if m.zf {
				m.ip = uint16(15769)
			}
			m.cycles += 8
		}
	case 0x3d8e:
		{
			m.ip = 0x3d92
			m.r[6] = m.rd16(m.r[11], uint16(0x12))
			m.cycles += 8
		}
	case 0x3d92:
		{
			m.ip = 0x3d96
			m.r[7] = m.rd16(m.r[11], uint16(0x3a))
			m.cycles += 8
		}
	case 0x3d96:
		{
			m.ip = 0x3d99
			target := uint16(10677)
			m.push(0x3d99)
			m.ip = target
			m.cycles += 19
		}
	case 0x3d99:
		{
			m.ip = 0x3d9e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(2))
			m.cycles += 16
		}
	case 0x3d9e:
		{
			m.ip = 0x3da0
			if !m.zf {
				m.ip = uint16(15782)
			}
			m.cycles += 8
		}
	case 0x3da0:
		{
			m.ip = 0x3da5
			m.wr8(m.r[11], uint16(0x76), uint16(1))
			m.cycles += 8
		}
	case 0x3da5:
		{
			m.ip = 0x3da6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3da6:
		{
			m.ip = 0x3dab
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x27a)), uint16(1))
			m.cycles += 16
		}
	case 0x3dab:
		{
			m.ip = 0x3daf
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3daf:
		{
			m.ip = 0x3db1
			if m.zf {
				m.ip = uint16(15903)
			}
			m.cycles += 8
		}
	case 0x3db1:
		{
			m.ip = 0x3db5
			m.wr16(m.r[11], uint16(m.r[3]+0x246), m.unary("dec", 16, m.rd16(m.r[11], uint16(m.r[3]+0x246))))
			m.cycles += 15
		}
	case 0x3db5:
		{
			m.ip = 0x3db7
			if !m.zf {
				m.ip = uint16(15903)
			}
			m.cycles += 8
		}
	case 0x3db7:
		{
			m.ip = 0x3dbc
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(0))
			m.cycles += 16
		}
	case 0x3dbc:
		{
			m.ip = 0x3dbe
			if m.zf {
				m.ip = uint16(15888)
			}
			m.cycles += 8
		}
	case 0x3dbe:
		{
			m.ip = 0x3dc3
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x61a)), uint16(2))
			m.cycles += 16
		}
	case 0x3dc3:
		{
			m.ip = 0x3dc5
			if !m.zf {
				m.ip = uint16(15888)
			}
			m.cycles += 8
		}
	case 0x3dc5:
		{
			m.ip = 0x3dcb
			m.wr16(m.r[11], uint16(m.r[3]+0x1ad), uint16(0))
			m.cycles += 8
		}
	case 0x3dcb:
		{
			m.ip = 0x3dd1
			m.wr16(m.r[11], uint16(m.r[3]+0x1a5), uint16(0))
			m.cycles += 8
		}
	case 0x3dd1:
		{
			m.ip = 0x3dd3
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3dd3:
		{
			m.ip = 0x3dd7
			m.wr8(m.r[11], uint16(0x61a), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3dd7:
		{
			m.ip = 0x3ddc
			m.wr8(m.r[11], uint16(m.r[3]+0x3e), uint16(1))
			m.cycles += 8
		}
	case 0x3ddc:
		{
			m.ip = 0x3ddf
			target := uint16(12292)
			m.push(0x3ddf)
			m.ip = target
			m.cycles += 19
		}
	case 0x3ddf:
		{
			m.ip = 0x3de2
			target := uint16(6912)
			m.push(0x3de2)
			m.ip = target
			m.cycles += 19
		}
	case 0x3de2:
		{
			m.ip = 0x3de7
			m.wr8(m.r[11], uint16(0x2077), uint16(1))
			m.cycles += 8
		}
	case 0x3de7:
		{
			m.ip = 0x3dec
			m.wr8(m.r[11], uint16(0x1984), uint16(1))
			m.cycles += 8
		}
	case 0x3dec:
		{
			m.ip = 0x3df1
			m.wr8(m.r[11], uint16(0x616), uint16(2))
			m.cycles += 8
		}
	case 0x3df1:
		{
			m.ip = 0x3df4
			m.r[0] = m.rd16(m.r[11], uint16(0x223))
			m.cycles += 8
		}
	case 0x3df4:
		{
			m.ip = 0x3df5
			m.r[0] = m.unary("inc", 16, m.r[0])
			m.cycles += 3
		}
	case 0x3df5:
		{
			m.ip = 0x3df8
			m.wr16(m.r[11], uint16(0x23b), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3df8:
		{
			m.ip = 0x3dfd
			m.wr8(m.r[11], uint16(0x2078), uint16(1))
			m.cycles += 8
		}
	case 0x3dfd:
		{
			m.ip = 0x3e00
			target := uint16(3091)
			m.push(0x3e00)
			m.ip = target
			m.cycles += 19
		}
	case 0x3e00:
		{
			m.ip = 0x3e05
			m.wr8(m.r[11], uint16(0x616), uint16(1))
			m.cycles += 8
		}
	case 0x3e05:
		{
			m.ip = 0x3e0a
			m.wr8(m.r[11], uint16(0x23b), uint16(1))
			m.cycles += 8
		}
	case 0x3e0a:
		{
			m.ip = 0x3e0d
			target := uint16(1919)
			m.push(0x3e0d)
			m.ip = target
			m.cycles += 19
		}
	case 0x3e0d:
		{
			m.ip = 0x3e0f
			m.ip = uint16(15903)
			m.cycles += 15
		}
	case 0x3e0f:
		{
			m.ip = 0x3e10
			m.cycles += 3
		}
	case 0x3e10:
		{
			m.ip = 0x3e16
			m.wr16(m.r[11], uint16(0x78), uint16(0))
			m.cycles += 8
		}
	case 0x3e16:
		{
			m.ip = 0x3e1b
			m.wr8(m.r[11], uint16(0x74), uint16(1))
			m.cycles += 8
		}
	case 0x3e1b:
		{
			m.ip = 0x3e1e
			target := uint16(6912)
			m.push(0x3e1e)
			m.ip = target
			m.cycles += 19
		}
	case 0x3e1e:
		{
			m.ip = 0x3e1f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3e1f:
		{
			m.ip = 0x3e24
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x3e24:
		{
			m.ip = 0x3e26
			if !m.zf {
				m.ip = uint16(15916)
			}
			m.cycles += 8
		}
	case 0x3e26:
		{
			m.ip = 0x3e29
			target := uint16(2340)
			m.push(0x3e29)
			m.ip = target
			m.cycles += 19
		}
	case 0x3e29:
		{
			m.ip = 0x3e2b
			m.ip = uint16(15919)
			m.cycles += 15
		}
	case 0x3e2b:
		{
			m.ip = 0x3e2c
			m.cycles += 3
		}
	case 0x3e2c:
		{
			m.ip = 0x3e2f
			target := uint16(2087)
			m.push(0x3e2f)
			m.ip = target
			m.cycles += 19
		}
	case 0x3e2f:
		{
			m.ip = 0x3e34
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x246)), uint16(0))
			m.cycles += 16
		}
	case 0x3e34:
		{
			m.ip = 0x3e36
			if m.zf {
				m.ip = uint16(15935)
			}
			m.cycles += 8
		}
	case 0x3e36:
		{
			m.ip = 0x3e3c
			m.wr16(m.r[11], uint16(0x106a), uint16(3000))
			m.cycles += 8
		}
	case 0x3e3c:
		{
			m.ip = 0x3e3f
			target := uint16(19311)
			m.push(0x3e3f)
			m.ip = target
			m.cycles += 19
		}
	case 0x3e3f:
		{
			m.ip = 0x3e40
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3e40:
		{
			m.ip = 0x3e43
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x3e43:
		{
			m.ip = 0x3e46
			m.r[7] = uint16(1120)
			m.cycles += 2
		}
	case 0x3e46:
		{
			m.ip = 0x3e49
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x616)))
			m.cycles += 8
		}
	case 0x3e49:
		{
			m.ip = 0x3e4d
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), m.rd8(m.r[11], uint16(0x23b))))
			m.cycles += 16
		}
	case 0x3e4d:
		{
			m.ip = 0x3e4f
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(3))
			m.cycles += 4
		}
	case 0x3e4f:
		{
			m.ip = 0x3e51
			if m.zf {
				m.ip = uint16(15972)
			}
			m.cycles += 8
		}
	case 0x3e51:
		{
			m.ip = 0x3e56
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x3e56:
		{
			m.ip = 0x3e58
			if !m.zf {
				m.ip = uint16(15969)
			}
			m.cycles += 8
		}
	case 0x3e58:
		{
			m.ip = 0x3e5c
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3e5c:
		{
			m.ip = 0x3e5f
			m.alu("sub", 16, m.r[3], uint16(0))
			m.cycles += 4
		}
	case 0x3e5f:
		{
			m.ip = 0x3e61
			if m.zf {
				m.ip = uint16(15972)
			}
			m.cycles += 8
		}
	case 0x3e61:
		{
			m.ip = 0x3e64
			m.r[7] = uint16(1180)
			m.cycles += 2
		}
	case 0x3e64:
		{
			m.ip = 0x3e68
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x246))
			m.cycles += 8
		}
	case 0x3e68:
		{
			m.ip = 0x3e69
			m.r[0] = m.unary("dec", 16, m.r[0])
			m.cycles += 3
		}
	case 0x3e69:
		{
			m.ip = 0x3e6b
			m.r[0] = m.shift("shl", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x3e6b:
		{
			m.ip = 0x3e6d
			m.r[0] = m.shift("shl", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x3e6d:
		{
			m.ip = 0x3e6f
			m.r[7] = m.alu("add", 16, m.r[7], m.r[0])
			m.cycles += 4
		}
	case 0x3e6f:
		{
			m.ip = 0x3e72
			m.r[6] = uint16(28076)
			m.cycles += 2
		}
	case 0x3e72:
		{
			m.ip = 0x3e75
			m.r[1] = uint16(1)
			m.cycles += 2
		}
	case 0x3e75:
		{
			m.ip = 0x3e78
			target := uint16(16988)
			m.push(0x3e78)
			m.ip = target
			m.cycles += 19
		}
	case 0x3e78:
		{
			m.ip = 0x3e79
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3e79:
		{
			m.ip = 0x3e7d
			m.r[7] = uint16(0x264)
			m.cycles += 3
		}
	case 0x3e7d:
		{
			m.ip = 0x3e82
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(2))
			m.cycles += 16
		}
	case 0x3e82:
		{
			m.ip = 0x3e84
			if m.zf {
				m.ip = uint16(16015)
			}
			m.cycles += 8
		}
	case 0x3e84:
		{
			m.ip = 0x3e89
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x223)), uint16(1))
			m.cycles += 16
		}
	case 0x3e89:
		{
			m.ip = 0x3e8b
			if m.zf {
				m.ip = uint16(16015)
			}
			m.cycles += 8
		}
	case 0x3e8b:
		{
			m.ip = 0x3e8f
			m.r[7] = uint16(0x25b)
			m.cycles += 3
		}
	case 0x3e8f:
		{
			m.ip = 0x3e91
			m.wr16(m.r[11], uint16(m.r[7]), m.unary("inc", 16, m.rd16(m.r[11], uint16(m.r[7]))))
			m.cycles += 15
		}
	case 0x3e91:
		{
			m.ip = 0x3e97
			m.wr16(m.r[11], uint16(0x55b), uint16(0))
			m.cycles += 8
		}
	case 0x3e97:
		{
			m.ip = 0x3e9b
			m.r[0] = uint16(0x5d7)
			m.cycles += 3
		}
	case 0x3e9b:
		{
			m.ip = 0x3e9e
			m.wr16(m.r[11], uint16(0x55d), m.r[0])
			m.cycles += 8
		}
	case 0x3e9e:
		{
			m.ip = 0x3ea2
			m.r[3] = m.rd16(m.r[11], uint16(0x221))
			m.cycles += 8
		}
	case 0x3ea2:
		{
			m.ip = 0x3ea7
			m.wr8(m.r[11], uint16(m.r[3]+0x7a), uint16(71))
			m.cycles += 8
		}
	case 0x3ea7:
		{
			m.ip = 0x3ea9
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3ea9:
		{
			m.ip = 0x3eaf
			m.wr16(m.r[11], uint16(m.r[3]+0x7e), uint16(0))
			m.cycles += 8
		}
	case 0x3eaf:
		{
			m.ip = 0x3eb3
			m.r[6] = uint16(0xff5)
			m.cycles += 3
		}
	case 0x3eb3:
		{
			m.ip = 0x3eb4
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x3eb4:
		{
			m.ip = 0x3eb7
			target := uint16(9346)
			m.push(0x3eb7)
			m.ip = target
			m.cycles += 19
		}
	case 0x3eb7:
		{
			m.ip = 0x3eb8
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x3eb8:
		{
			m.ip = 0x3ebb
			target := uint16(19095)
			m.push(0x3ebb)
			m.ip = target
			m.cycles += 19
		}
	case 0x3ebb:
		{
			m.ip = 0x3ebd
			m.set8(0, 0, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x3ebd:
		{
			m.ip = 0x3ec0
			m.r[0] = m.alu("and", 16, m.r[0], uint16(255))
			m.cycles += 4
		}
	case 0x3ec0:
		{
			m.ip = 0x3ec3
			m.r[0] = m.alu("add", 16, m.r[0], uint16(100))
			m.cycles += 4
		}
	case 0x3ec3:
		{
			m.ip = 0x3ec7
			m.wr16(m.r[11], uint16(m.r[3]+0x1d5), m.r[0])
			m.cycles += 8
		}
	case 0x3ec7:
		{
			m.ip = 0x3eca
			m.r[6] = uint16(28064)
			m.cycles += 2
		}
	case 0x3eca:
		{
			m.ip = 0x3ece
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1cd))
			m.cycles += 8
		}
	case 0x3ece:
		{
			m.ip = 0x3ed0
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3ed0:
		{
			m.ip = 0x3ed5
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)), uint16(1))
			m.cycles += 16
		}
	case 0x3ed5:
		{
			m.ip = 0x3ed7
			if !m.zf {
				m.ip = uint16(16100)
			}
			m.cycles += 8
		}
	case 0x3ed7:
		{
			m.ip = 0x3edc
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[3]+0x1b1)), uint16(16))
			m.cycles += 16
		}
	case 0x3edc:
		{
			m.ip = 0x3ede
			if !m.zf {
				m.ip = uint16(16100)
			}
			m.cycles += 8
		}
	case 0x3ede:
		{
			m.ip = 0x3ee1
			target := uint16(10719)
			m.push(0x3ee1)
			m.ip = target
			m.cycles += 19
		}
	case 0x3ee1:
		{
			m.ip = 0x3ee3
			m.ip = uint16(16103)
			m.cycles += 15
		}
	case 0x3ee3:
		{
			m.ip = 0x3ee4
			m.cycles += 3
		}
	case 0x3ee4:
		{
			m.ip = 0x3ee7
			target := uint16(10677)
			m.push(0x3ee7)
			m.ip = target
			m.cycles += 19
		}
	case 0x3ee7:
		{
			m.ip = 0x3ee9
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3ee9:
		{
			m.ip = 0x3eef
			m.wr16(m.r[11], uint16(m.r[3]+0x86), uint16(50))
			m.cycles += 8
		}
	case 0x3eef:
		{
			m.ip = 0x3ef3
			m.r[6] = m.rd16(m.r[11], uint16(0x66))
			m.cycles += 8
		}
	case 0x3ef3:
		{
			m.ip = 0x3ef7
			m.r[2] = m.rd16(m.r[11], uint16(m.r[6]+0x53a))
			m.cycles += 8
		}
	case 0x3ef7:
		{
			m.ip = 0x3efb
			m.r[6] = m.rd16(m.r[11], uint16(m.r[6]+0x532))
			m.cycles += 8
		}
	case 0x3efb:
		{
			m.ip = 0x3eff
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1bd))
			m.cycles += 8
		}
	case 0x3eff:
		{
			m.ip = 0x3f01
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3f01:
		{
			m.ip = 0x3f05
			m.r[5] = m.rd16(m.r[11], uint16(m.r[3]+0x0))
			m.cycles += 8
		}
	case 0x3f05:
		{
			m.ip = 0x3f0b
			m.wr8(m.r[11], uint16(m.r[5]+0x8e), m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[5]+0x8e)), uint16(251)))
			m.cycles += 16
		}
	case 0x3f0b:
		{
			m.ip = 0x3f0f
			m.wr16(m.r[11], uint16(m.r[3]+0x0), m.r[7])
			m.cycles += 8
		}
	case 0x3f0f:
		{
			m.ip = 0x3f13
			m.wr16(m.r[11], uint16(m.r[3]+0x2), m.r[6])
			m.cycles += 8
		}
	case 0x3f13:
		{
			m.ip = 0x3f15
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3f15:
		{
			m.ip = 0x3f1a
			m.wr8(m.r[11], uint16(m.r[7]+0x8e), m.alu("or", 8, m.rd8(m.r[11], uint16(m.r[7]+0x8e)), uint16(4)))
			m.cycles += 16
		}
	case 0x3f1a:
		{
			m.ip = 0x3f1c
			m.r[7] = m.shift("shl", 16, m.r[7], uint16(1))
			m.cycles += 8
		}
	case 0x3f1c:
		{
			m.ip = 0x3f20
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x61c))
			m.cycles += 8
		}
	case 0x3f20:
		{
			m.ip = 0x3f23
			target := uint16(10677)
			m.push(0x3f23)
			m.ip = target
			m.cycles += 19
		}
	case 0x3f23:
		{
			m.ip = 0x3f27
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x1c5))
			m.cycles += 8
		}
	case 0x3f27:
		{
			m.ip = 0x3f2c
			m.alu("and", 8, m.rd8(m.r[11], uint16(m.r[7]+0x8e)), uint16(8))
			m.cycles += 16
		}
	case 0x3f2c:
		{
			m.ip = 0x3f2e
			if m.zf {
				m.ip = uint16(16186)
			}
			m.cycles += 8
		}
	case 0x3f2e:
		{
			m.ip = 0x3f30
			m.r[7] = m.shift("shl", 16, m.r[7], uint16(1))
			m.cycles += 8
		}
	case 0x3f30:
		{
			m.ip = 0x3f34
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x61c))
			m.cycles += 8
		}
	case 0x3f34:
		{
			m.ip = 0x3f37
			m.r[6] = uint16(28068)
			m.cycles += 2
		}
	case 0x3f37:
		{
			m.ip = 0x3f3a
			target := uint16(10677)
			m.push(0x3f3a)
			m.ip = target
			m.cycles += 19
		}
	case 0x3f3a:
		{
			m.ip = 0x3f3c
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3f3c:
		{
			m.ip = 0x3f3f
			m.r[1] = uint16(3)
			m.cycles += 2
		}
	case 0x3f3f:
		{
			m.ip = 0x3f43
			m.r[6] = m.rd16(m.r[11], uint16(0x21d))
			m.cycles += 8
		}
	case 0x3f43:
		{
			m.ip = 0x3f46
			m.r[6] = m.alu("add", 16, m.r[6], uint16(18))
			m.cycles += 4
		}
	case 0x3f46:
		{
			m.ip = 0x3f4a
			m.r[7] = uint16(0x1b1)
			m.cycles += 3
		}
	case 0x3f4a:
		{
			m.ip = 0x3f4c
			m.r[7] = m.alu("add", 16, m.r[7], m.r[3])
			m.cycles += 4
		}
	case 0x3f4c:
		{
			m.ip = 0x3f4e
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x3f4e:
		{
			m.ip = 0x3f50
			m.wr8(m.r[11], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3f50:
		{
			m.ip = 0x3f53
			m.r[6] = m.alu("add", 16, m.r[6], uint16(4))
			m.cycles += 4
		}
	case 0x3f53:
		{
			m.ip = 0x3f56
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x3f56:
		{
			m.ip = 0x3f58
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16204)
			}
			m.cycles += 17
		}
	case 0x3f58:
		{
			m.ip = 0x3f5a
			m.r[7] = m.alu("add", 16, m.r[7], m.r[3])
			m.cycles += 4
		}
	case 0x3f5a:
		{
			m.ip = 0x3f5d
			m.r[1] = uint16(3)
			m.cycles += 2
		}
	case 0x3f5d:
		{
			m.ip = 0x3f5f
			m.r[0] = m.rd16(m.r[11], uint16(m.r[6]))
			m.cycles += 8
		}
	case 0x3f5f:
		{
			m.ip = 0x3f61
			m.wr16(m.r[11], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x3f61:
		{
			m.ip = 0x3f64
			m.r[6] = m.alu("add", 16, m.r[6], uint16(8))
			m.cycles += 4
		}
	case 0x3f64:
		{
			m.ip = 0x3f67
			m.r[7] = m.alu("add", 16, m.r[7], uint16(8))
			m.cycles += 4
		}
	case 0x3f67:
		{
			m.ip = 0x3f69
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16221)
			}
			m.cycles += 17
		}
	case 0x3f69:
		{
			m.ip = 0x3f6d
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x3f6d:
		{
			m.ip = 0x3f71
			m.wr16(m.r[11], uint16(m.r[3]+0x233), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x233)), m.r[2]))
			m.cycles += 16
		}
	case 0x3f71:
		{
			m.ip = 0x3f74
			target := uint16(17038)
			m.push(0x3f74)
			m.ip = target
			m.cycles += 19
		}
	case 0x3f74:
		{
			m.ip = 0x3f79
			m.wr16(m.r[11], uint16(0x66), m.alu("add", 16, m.rd16(m.r[11], uint16(0x66)), uint16(2)))
			m.cycles += 16
		}
	case 0x3f79:
		{
			m.ip = 0x3f7c
			m.r[5] = uint16(0)
			m.cycles += 2
		}
	case 0x3f7c:
		{
			m.ip = 0x3f81
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x66)), uint16(8))
			m.cycles += 16
		}
	case 0x3f81:
		{
			m.ip = 0x3f83
			if m.zf {
				m.ip = uint16(16260)
			}
			m.cycles += 8
		}
	case 0x3f83:
		{
			m.ip = 0x3f84
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3f84:
		{
			m.ip = 0x3f8a
			m.wr16(m.r[11], uint16(0x66), uint16(0))
			m.cycles += 8
		}
	case 0x3f8a:
		{
			m.ip = 0x3f90
			m.wr8(m.r[11], uint16(m.r[5]+0x7a), uint16(71))
			m.cycles += 8
		}
	case 0x3f90:
		{
			m.ip = 0x3f92
			m.r[5] = m.shift("shl", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x3f92:
		{
			m.ip = 0x3f99
			m.wr16(m.r[11], uint16(m.r[5]+0x7e), uint16(0))
			m.cycles += 8
		}
	case 0x3f99:
		{
			m.ip = 0x3f9b
			m.r[5] = m.shift("shr", 16, m.r[5], uint16(1))
			m.cycles += 8
		}
	case 0x3f9b:
		{
			m.ip = 0x3f9c
			m.r[5] = m.unary("inc", 16, m.r[5])
			m.cycles += 3
		}
	case 0x3f9c:
		{
			m.ip = 0x3f9f
			m.alu("sub", 16, m.r[5], uint16(4))
			m.cycles += 4
		}
	case 0x3f9f:
		{
			m.ip = 0x3fa1
			if !m.zf {
				m.ip = uint16(16260)
			}
			m.cycles += 8
		}
	case 0x3fa1:
		{
			m.ip = 0x3fa5
			m.r[6] = uint16(0xff5)
			m.cycles += 3
		}
	case 0x3fa5:
		{
			m.ip = 0x3fa8
			m.r[3] = uint16(65535)
			m.cycles += 2
		}
	case 0x3fa8:
		{
			m.ip = 0x3fab
			target := uint16(9346)
			m.push(0x3fab)
			m.ip = target
			m.cycles += 19
		}
	case 0x3fab:
		{
			m.ip = 0x3fac
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x3fac:
		{
			m.ip = 0x3fae
			m.set8(2, 8, uint16(0))
			m.cycles += 2
		}
	case 0x3fae:
		{
			m.ip = 0x3fb2
			m.set8(2, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x3fb2:
		{
			m.ip = 0x3fb4
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3fb4:
		{
			m.ip = 0x3fb8
			m.r[6] = m.rd16(m.r[11], uint16(m.r[3]+0x205))
			m.cycles += 8
		}
	case 0x3fb8:
		{
			m.ip = 0x3fba
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3fba:
		{
			m.ip = 0x3fbc
			m.set8(0, 8, ((m.r[2] >> 8) & 255))
			m.cycles += 2
		}
	case 0x3fbc:
		{
			m.ip = 0x3fc0
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b1)))
			m.cycles += 8
		}
	case 0x3fc0:
		{
			m.ip = 0x3fc2
			m.alu("and", 8, ((m.r[0] >> 0) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x3fc2:
		{
			m.ip = 0x3fc4
			if m.zf {
				m.ip = uint16(16340)
			}
			m.cycles += 8
		}
	case 0x3fc4:
		{
			m.ip = 0x3fc7
			m.r[6] = m.alu("add", 16, m.r[6], uint16(32))
			m.cycles += 4
		}
	case 0x3fc7:
		{
			m.ip = 0x3fca
			m.alu("and", 8, ((m.r[2] >> 0) & 255), uint16(1))
			m.cycles += 4
		}
	case 0x3fca:
		{
			m.ip = 0x3fcc
			if !m.zf {
				m.ip = uint16(16340)
			}
			m.cycles += 8
		}
	case 0x3fcc:
		{
			m.ip = 0x3fce
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3fce:
		{
			m.ip = 0x3fd2
			m.r[6] = m.rd16(m.r[11], uint16(m.r[3]+0x20d))
			m.cycles += 8
		}
	case 0x3fd2:
		{
			m.ip = 0x3fd4
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3fd4:
		{
			m.ip = 0x3fd7
			m.set8(2, 0, m.alu("and", 8, ((m.r[2]>>0)&255), uint16(254)))
			m.cycles += 4
		}
	case 0x3fd7:
		{
			m.ip = 0x3fd9
			m.r[2] = m.shift("shl", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x3fd9:
		{
			m.ip = 0x3fdb
			m.r[6] = m.alu("add", 16, m.r[6], m.r[2])
			m.cycles += 4
		}
	case 0x3fdb:
		{
			m.ip = 0x3fdd
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(8)))
			m.cycles += 4
		}
	case 0x3fdd:
		{
			m.ip = 0x3fdf
			m.r[0] = m.shift("shl", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x3fdf:
		{
			m.ip = 0x3fe1
			m.r[6] = m.alu("add", 16, m.r[6], m.r[0])
			m.cycles += 4
		}
	case 0x3fe1:
		{
			m.ip = 0x3fe3
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3fe3:
		{
			m.ip = 0x3fe7
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x1d5))
			m.cycles += 8
		}
	case 0x3fe7:
		{
			m.ip = 0x3fe9
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x3fe9:
		{
			m.ip = 0x3fec
			m.alu("sub", 16, m.r[0], uint16(16))
			m.cycles += 4
		}
	case 0x3fec:
		{
			m.ip = 0x3fee
			if !m.cf && !m.zf {
				m.ip = uint16(16408)
			}
			m.cycles += 8
		}
	case 0x3fee:
		{
			m.ip = 0x3ff1
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x3ff1:
		{
			m.ip = 0x3ff4
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x3ff4:
		{
			m.ip = 0x3ff7
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3ff7:
		{
			m.ip = 0x3ffb
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x3ffb:
		{
			m.ip = 0x3fff
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x3fff:
		{
			m.ip = 0x4003
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x4003:
		{
			m.ip = 0x4007
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4007:
		{
			m.ip = 0x400b
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x400b:
		{
			m.ip = 0x400f
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x400f:
		{
			m.ip = 0x4012
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x4012:
		{
			m.ip = 0x4015
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x4015:
		{
			m.ip = 0x4017
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16369)
			}
			m.cycles += 17
		}
	case 0x4017:
		{
			m.ip = 0x4018
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4018:
		{
			m.ip = 0x401b
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x401b:
		{
			m.ip = 0x401d
			m.r[2] = m.r[1]
			m.cycles += 2
		}
	case 0x401d:
		{
			m.ip = 0x4021
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x4021:
		{
			m.ip = 0x4023
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x4023:
		{
			m.ip = 0x4025
			m.set8(1, 0, m.unary("neg", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x4025:
		{
			m.ip = 0x4028
			m.set8(1, 0, m.alu("add", 8, ((m.r[1]>>0)&255), uint16(8)))
			m.cycles += 4
		}
	case 0x4028:
		{
			m.ip = 0x402b
			m.set8(2, 0, m.alu("sub", 8, ((m.r[2]>>0)&255), uint16(3)))
			m.cycles += 4
		}
	case 0x402b:
		{
			m.ip = 0x402f
			m.r[6] = m.alu("add", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x402f:
		{
			m.ip = 0x4033
			m.r[7] = m.alu("add", 16, m.r[7], uint16(240))
			m.cycles += 4
		}
	case 0x4033:
		{
			m.ip = 0x4035
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16424)
			}
			m.cycles += 17
		}
	case 0x4035:
		{
			m.ip = 0x4037
			m.r[1] = m.r[2]
			m.cycles += 2
		}
	case 0x4037:
		{
			m.ip = 0x403a
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x403a:
		{
			m.ip = 0x403d
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x403d:
		{
			m.ip = 0x4041
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x4041:
		{
			m.ip = 0x4045
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4045:
		{
			m.ip = 0x4049
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x4049:
		{
			m.ip = 0x404d
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x404d:
		{
			m.ip = 0x4051
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x4051:
		{
			m.ip = 0x4055
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4055:
		{
			m.ip = 0x4058
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x4058:
		{
			m.ip = 0x405b
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x405b:
		{
			m.ip = 0x405d
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16439)
			}
			m.cycles += 17
		}
	case 0x405d:
		{
			m.ip = 0x405e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x405e:
		{
			m.ip = 0x4060
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x4060:
		{
			m.ip = 0x4064
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b1)))
			m.cycles += 8
		}
	case 0x4064:
		{
			m.ip = 0x4066
			m.r[0] = m.r[1]
			m.cycles += 2
		}
	case 0x4066:
		{
			m.ip = 0x4068
			m.r[1] = m.alu("add", 16, m.r[1], m.r[1])
			m.cycles += 4
		}
	case 0x4068:
		{
			m.ip = 0x406a
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x406a:
		{
			m.ip = 0x406c
			m.r[1] = m.alu("add", 16, m.r[1], m.r[0])
			m.cycles += 4
		}
	case 0x406c:
		{
			m.ip = 0x406e
			m.set8(2, 8, uint16(0))
			m.cycles += 2
		}
	case 0x406e:
		{
			m.ip = 0x4072
			m.set8(2, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x4072:
		{
			m.ip = 0x4075
			m.set8(2, 0, m.alu("and", 8, ((m.r[2]>>0)&255), uint16(254)))
			m.cycles += 4
		}
	case 0x4075:
		{
			m.ip = 0x4077
			m.r[1] = m.alu("add", 16, m.r[1], m.r[2])
			m.cycles += 4
		}
	case 0x4077:
		{
			m.ip = 0x4079
			m.r[1] = m.alu("add", 16, m.r[1], m.r[2])
			m.cycles += 4
		}
	case 0x4079:
		{
			m.ip = 0x407b
			m.r[2] = m.shift("shr", 16, m.r[2], uint16(1))
			m.cycles += 8
		}
	case 0x407b:
		{
			m.ip = 0x407d
			m.r[1] = m.alu("add", 16, m.r[1], m.r[2])
			m.cycles += 4
		}
	case 0x407d:
		{
			m.ip = 0x407f
			m.r[3] = m.shift("shl", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x407f:
		{
			m.ip = 0x4083
			m.r[6] = m.rd16(m.r[11], uint16(m.r[3]+0x215))
			m.cycles += 8
		}
	case 0x4083:
		{
			m.ip = 0x4085
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4085:
		{
			m.ip = 0x4087
			m.r[6] = m.alu("add", 16, m.r[6], m.r[1])
			m.cycles += 4
		}
	case 0x4087:
		{
			m.ip = 0x408a
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x408a:
		{
			m.ip = 0x408d
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x408d:
		{
			m.ip = 0x4090
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4090:
		{
			m.ip = 0x4094
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x4094:
		{
			m.ip = 0x4098
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4098:
		{
			m.ip = 0x409c
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x409c:
		{
			m.ip = 0x40a0
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x40a0:
		{
			m.ip = 0x40a4
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x40a4:
		{
			m.ip = 0x40a8
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x40a8:
		{
			m.ip = 0x40ac
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x4)))
			m.cycles += 8
		}
	case 0x40ac:
		{
			m.ip = 0x40b0
			m.wr8(m.r[8], uint16(m.r[7]+0x4), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x40b0:
		{
			m.ip = 0x40b3
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x40b3:
		{
			m.ip = 0x40b6
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x40b6:
		{
			m.ip = 0x40b8
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16522)
			}
			m.cycles += 17
		}
	case 0x40b8:
		{
			m.ip = 0x40b9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x40b9:
		{
			m.ip = 0x40bd
			m.r[6] = m.rd16(m.r[11], uint16(0x21f))
			m.cycles += 8
		}
	case 0x40bd:
		{
			m.ip = 0x40c1
			m.r[6] = m.alu("add", 16, m.r[6], uint16(1920))
			m.cycles += 4
		}
	case 0x40c1:
		{
			m.ip = 0x40c3
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x40c3:
		{
			m.ip = 0x40c7
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x40c7:
		{
			m.ip = 0x40c9
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x40c9:
		{
			m.ip = 0x40cd
			m.r[6] = m.alu("sub", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x40cd:
		{
			m.ip = 0x40cf
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16585)
			}
			m.cycles += 17
		}
	case 0x40cf:
		{
			m.ip = 0x40d2
			m.r[1] = uint16(3)
			m.cycles += 2
		}
	case 0x40d2:
		{
			m.ip = 0x40d5
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x40d5:
		{
			m.ip = 0x40d8
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x40d8:
		{
			m.ip = 0x40dc
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x40dc:
		{
			m.ip = 0x40e0
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x40e0:
		{
			m.ip = 0x40e4
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x40e4:
		{
			m.ip = 0x40e8
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x40e8:
		{
			m.ip = 0x40ec
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x40ec:
		{
			m.ip = 0x40f0
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x40f0:
		{
			m.ip = 0x40f3
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x40f3:
		{
			m.ip = 0x40f6
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x40f6:
		{
			m.ip = 0x40f8
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16594)
			}
			m.cycles += 17
		}
	case 0x40f8:
		{
			m.ip = 0x40f9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x40f9:
		{
			m.ip = 0x40fd
			m.r[6] = m.rd16(m.r[11], uint16(0x21f))
			m.cycles += 8
		}
	case 0x40fd:
		{
			m.ip = 0x40ff
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x40ff:
		{
			m.ip = 0x4103
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x4103:
		{
			m.ip = 0x4105
			m.set8(1, 0, m.unary("inc", 8, ((m.r[1]>>0)&255)))
			m.cycles += 3
		}
	case 0x4105:
		{
			m.ip = 0x4109
			m.r[6] = m.alu("sub", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x4109:
		{
			m.ip = 0x410d
			m.r[6] = m.alu("add", 16, m.r[6], uint16(240))
			m.cycles += 4
		}
	case 0x410d:
		{
			m.ip = 0x410f
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16649)
			}
			m.cycles += 17
		}
	case 0x410f:
		{
			m.ip = 0x4112
			m.r[1] = uint16(3)
			m.cycles += 2
		}
	case 0x4112:
		{
			m.ip = 0x4115
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4115:
		{
			m.ip = 0x4118
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4118:
		{
			m.ip = 0x411c
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x411c:
		{
			m.ip = 0x4120
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4120:
		{
			m.ip = 0x4124
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x4124:
		{
			m.ip = 0x4128
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4128:
		{
			m.ip = 0x412c
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x412c:
		{
			m.ip = 0x4130
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4130:
		{
			m.ip = 0x4133
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x4133:
		{
			m.ip = 0x4136
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x4136:
		{
			m.ip = 0x4138
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16658)
			}
			m.cycles += 17
		}
	case 0x4138:
		{
			m.ip = 0x4139
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4139:
		{
			m.ip = 0x413d
			m.r[6] = m.rd16(m.r[11], uint16(0x21f))
			m.cycles += 8
		}
	case 0x413d:
		{
			m.ip = 0x413f
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x413f:
		{
			m.ip = 0x4143
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x4143:
		{
			m.ip = 0x4145
			m.set8(1, 0, m.shift("shr", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x4145:
		{
			m.ip = 0x4147
			m.r[6] = m.alu("add", 16, m.r[6], m.r[1])
			m.cycles += 4
		}
	case 0x4147:
		{
			m.ip = 0x414a
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x414a:
		{
			m.ip = 0x414d
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x414d:
		{
			m.ip = 0x4150
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4150:
		{
			m.ip = 0x4153
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x4153:
		{
			m.ip = 0x4156
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x4156:
		{
			m.ip = 0x4158
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16714)
			}
			m.cycles += 17
		}
	case 0x4158:
		{
			m.ip = 0x415c
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x415c:
		{
			m.ip = 0x415d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x415d:
		{
			m.ip = 0x4161
			m.r[6] = m.rd16(m.r[11], uint16(0x21f))
			m.cycles += 8
		}
	case 0x4161:
		{
			m.ip = 0x4164
			m.r[6] = m.alu("add", 16, m.r[6], uint16(3))
			m.cycles += 4
		}
	case 0x4164:
		{
			m.ip = 0x4166
			m.set8(1, 8, uint16(0))
			m.cycles += 2
		}
	case 0x4166:
		{
			m.ip = 0x416a
			m.set8(1, 0, m.rd8(m.r[11], uint16(m.r[3]+0x1b9)))
			m.cycles += 8
		}
	case 0x416a:
		{
			m.ip = 0x416c
			m.set8(1, 0, m.shift("shr", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x416c:
		{
			m.ip = 0x416e
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[1])
			m.cycles += 4
		}
	case 0x416e:
		{
			m.ip = 0x4171
			m.r[1] = uint16(24)
			m.cycles += 2
		}
	case 0x4171:
		{
			m.ip = 0x4174
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4174:
		{
			m.ip = 0x4178
			m.wr8(m.r[8], uint16(m.r[7]+0x4), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4178:
		{
			m.ip = 0x417b
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x417b:
		{
			m.ip = 0x417e
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x417e:
		{
			m.ip = 0x4180
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16753)
			}
			m.cycles += 17
		}
	case 0x4180:
		{
			m.ip = 0x4184
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(1920))
			m.cycles += 4
		}
	case 0x4184:
		{
			m.ip = 0x4185
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4185:
		{
			m.ip = 0x4188
			m.r[0] = uint16(5)
			m.cycles += 2
		}
	case 0x4188:
		{
			m.ip = 0x418b
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x418b:
		{
			m.ip = 0x418c
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x418c:
		{
			m.ip = 0x418f
			m.r[0] = uint16(0)
			m.cycles += 2
		}
	case 0x418f:
		{
			m.ip = 0x4190
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4190:
		{
			m.ip = 0x4193
			m.r[0] = uint16(1)
			m.cycles += 2
		}
	case 0x4193:
		{
			m.ip = 0x4194
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4194:
		{
			m.ip = 0x4197
			m.r[0] = uint16(2)
			m.cycles += 2
		}
	case 0x4197:
		{
			m.ip = 0x4198
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4198:
		{
			m.ip = 0x419b
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x419b:
		{
			m.ip = 0x419c
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x419c:
		{
			m.ip = 0x419f
			m.r[0] = uint16(4)
			m.cycles += 2
		}
	case 0x419f:
		{
			m.ip = 0x41a0
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x41a0:
		{
			m.ip = 0x41a3
			m.r[0] = uint16(3847)
			m.cycles += 2
		}
	case 0x41a3:
		{
			m.ip = 0x41a4
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x41a4:
		{
			m.ip = 0x41a7
			m.r[0] = uint16(65288)
			m.cycles += 2
		}
	case 0x41a7:
		{
			m.ip = 0x41a8
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x41a8:
		{
			m.ip = 0x41ac
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x41ac:
		{
			m.ip = 0x41ae
			m.r[7] = m.alu("xor", 16, m.r[7], m.r[7])
			m.cycles += 4
		}
	case 0x41ae:
		{
			m.ip = 0x41b1
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x41b1:
		{
			m.ip = 0x41b3
			m.set8(0, 8, uint16(15))
			m.cycles += 2
		}
	case 0x41b3:
		{
			m.ip = 0x41b5
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x41b5:
		{
			m.ip = 0x41b6
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x41b6:
		{
			m.ip = 0x41b9
			m.r[0] = uint16(3841)
			m.cycles += 2
		}
	case 0x41b9:
		{
			m.ip = 0x41ba
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x41ba:
		{
			m.ip = 0x41bd
			m.r[1] = uint16(14000)
			m.cycles += 2
		}
	case 0x41bd:
		{
			m.ip = 0x41be
			m.df = false
			m.cycles += 2
		}
	case 0x41be:
		{
			m.ip = 0x41c0
			m.stringOp("stos", 16)
			m.cycles += 2
		}
	case 0x41c0:
		{
			m.ip = 0x41c3
			m.r[0] = uint16(1)
			m.cycles += 2
		}
	case 0x41c3:
		{
			m.ip = 0x41c4
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x41c4:
		{
			m.ip = 0x41c7
			m.r[0] = uint16(261)
			m.cycles += 2
		}
	case 0x41c7:
		{
			m.ip = 0x41ca
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x41ca:
		{
			m.ip = 0x41cb
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x41cb:
		{
			m.ip = 0x41cc
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x41cc:
		{
			m.ip = 0x41cf
			m.r[6] = uint16(55440)
			m.cycles += 2
		}
	case 0x41cf:
		{
			m.ip = 0x41d2
			m.r[6] = m.alu("add", 16, m.r[6], uint16(52))
			m.cycles += 4
		}
	case 0x41d2:
		{
			m.ip = 0x41d5
			m.r[7] = uint16(1)
			m.cycles += 2
		}
	case 0x41d5:
		{
			m.ip = 0x41d8
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x41d8:
		{
			m.ip = 0x41d9
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x41d9:
		{
			m.ip = 0x41dc
			m.r[1] = uint16(18)
			m.cycles += 2
		}
	case 0x41dc:
		{
			m.ip = 0x41df
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x41df:
		{
			m.ip = 0x41e2
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x41e2:
		{
			m.ip = 0x41e3
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x41e3:
		{
			m.ip = 0x41e4
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x41e4:
		{
			m.ip = 0x41e6
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16860)
			}
			m.cycles += 17
		}
	case 0x41e6:
		{
			m.ip = 0x41e9
			m.r[6] = m.alu("add", 16, m.r[6], uint16(62))
			m.cycles += 4
		}
	case 0x41e9:
		{
			m.ip = 0x41ec
			m.r[7] = m.alu("add", 16, m.r[7], uint16(62))
			m.cycles += 4
		}
	case 0x41ec:
		{
			m.ip = 0x41ed
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x41ed:
		{
			m.ip = 0x41ef
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16856)
			}
			m.cycles += 17
		}
	case 0x41ef:
		{
			m.ip = 0x41f3
			m.push(m.rd16(m.r[11], uint16(0x233)))
			m.cycles += 11
		}
	case 0x41f3:
		{
			m.ip = 0x41f6
			m.r[0] = m.rd16(m.r[11], uint16(0x1985))
			m.cycles += 8
		}
	case 0x41f6:
		{
			m.ip = 0x41f9
			m.wr16(m.r[11], uint16(0x233), m.r[0])
			m.cycles += 8
		}
	case 0x41f9:
		{
			m.ip = 0x41ff
			m.wr16(m.r[11], uint16(0x102b), uint16(55440))
			m.cycles += 8
		}
	case 0x41ff:
		{
			m.ip = 0x4205
			m.wr16(m.r[11], uint16(0x102d), uint16(3128))
			m.cycles += 8
		}
	case 0x4205:
		{
			m.ip = 0x4208
			target := uint16(17038)
			m.push(0x4208)
			m.ip = target
			m.cycles += 19
		}
	case 0x4208:
		{
			m.ip = 0x420c
			m.wr16(m.r[11], uint16(0x233), m.pop())
			m.cycles += 8
		}
	case 0x420c:
		{
			m.ip = 0x420d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x420d:
		{
			m.ip = 0x4210
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x4210:
		{
			m.ip = 0x4211
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4211:
		{
			m.ip = 0x4214
			m.r[1] = uint16(16)
			m.cycles += 2
		}
	case 0x4214:
		{
			m.ip = 0x4217
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4217:
		{
			m.ip = 0x421a
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x421a:
		{
			m.ip = 0x421b
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x421b:
		{
			m.ip = 0x421c
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x421c:
		{
			m.ip = 0x421e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16916)
			}
			m.cycles += 17
		}
	case 0x421e:
		{
			m.ip = 0x4221
			m.r[6] = m.alu("add", 16, m.r[6], uint16(64))
			m.cycles += 4
		}
	case 0x4221:
		{
			m.ip = 0x4224
			m.r[7] = m.alu("add", 16, m.r[7], uint16(64))
			m.cycles += 4
		}
	case 0x4224:
		{
			m.ip = 0x4225
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4225:
		{
			m.ip = 0x4227
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16912)
			}
			m.cycles += 17
		}
	case 0x4227:
		{
			m.ip = 0x4228
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4228:
		{
			m.ip = 0x422c
			m.r[1] = m.rd16(m.r[11], uint16(0x246))
			m.cycles += 8
		}
	case 0x422c:
		{
			m.ip = 0x4231
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x616)), uint16(1))
			m.cycles += 16
		}
	case 0x4231:
		{
			m.ip = 0x4233
			if m.zf {
				m.ip = uint16(16962)
			}
			m.cycles += 8
		}
	case 0x4233:
		{
			m.ip = 0x4237
			m.r[1] = m.rd16(m.r[11], uint16(0x24a))
			m.cycles += 8
		}
	case 0x4237:
		{
			m.ip = 0x423c
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x23b)), uint16(1))
			m.cycles += 16
		}
	case 0x423c:
		{
			m.ip = 0x423e
			if m.zf {
				m.ip = uint16(16962)
			}
			m.cycles += 8
		}
	case 0x423e:
		{
			m.ip = 0x4242
			m.r[1] = m.rd16(m.r[11], uint16(0x24c))
			m.cycles += 8
		}
	case 0x4242:
		{
			m.ip = 0x4247
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x618)), uint16(1))
			m.cycles += 16
		}
	case 0x4247:
		{
			m.ip = 0x4249
			if !m.zf {
				m.ip = uint16(16980)
			}
			m.cycles += 8
		}
	case 0x4249:
		{
			m.ip = 0x424e
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x223)), uint16(0))
			m.cycles += 16
		}
	case 0x424e:
		{
			m.ip = 0x4250
			if m.zf {
				m.ip = uint16(16980)
			}
			m.cycles += 8
		}
	case 0x4250:
		{
			m.ip = 0x4254
			m.r[1] = m.rd16(m.r[11], uint16(0x248))
			m.cycles += 8
		}
	case 0x4254:
		{
			m.ip = 0x4257
			m.alu("sub", 16, m.r[1], uint16(0))
			m.cycles += 4
		}
	case 0x4257:
		{
			m.ip = 0x4259
			if m.zf {
				m.ip = uint16(17037)
			}
			m.cycles += 8
		}
	case 0x4259:
		{
			m.ip = 0x425c
			m.r[6] = uint16(28072)
			m.cycles += 2
		}
	case 0x425c:
		{
			m.ip = 0x425d
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x425d:
		{
			m.ip = 0x4260
			m.r[1] = uint16(18)
			m.cycles += 2
		}
	case 0x4260:
		{
			m.ip = 0x4263
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4263:
		{
			m.ip = 0x4266
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4266:
		{
			m.ip = 0x426a
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x426a:
		{
			m.ip = 0x426e
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x426e:
		{
			m.ip = 0x4272
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x2)))
			m.cycles += 8
		}
	case 0x4272:
		{
			m.ip = 0x4276
			m.wr8(m.r[8], uint16(m.r[7]+0x2), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4276:
		{
			m.ip = 0x427a
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x3)))
			m.cycles += 8
		}
	case 0x427a:
		{
			m.ip = 0x427e
			m.wr8(m.r[8], uint16(m.r[7]+0x3), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x427e:
		{
			m.ip = 0x4281
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x4281:
		{
			m.ip = 0x4284
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x4284:
		{
			m.ip = 0x4286
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16992)
			}
			m.cycles += 17
		}
	case 0x4286:
		{
			m.ip = 0x428a
			m.r[7] = m.alu("add", 16, m.r[7], uint16(64100))
			m.cycles += 4
		}
	case 0x428a:
		{
			m.ip = 0x428b
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x428b:
		{
			m.ip = 0x428d
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(16980)
			}
			m.cycles += 17
		}
	case 0x428d:
		{
			m.ip = 0x428e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x428e:
		{
			m.ip = 0x428f
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x428f:
		{
			m.ip = 0x4290
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4290:
		{
			m.ip = 0x4291
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x4291:
		{
			m.ip = 0x4295
			m.r[3] = m.rd16(m.r[11], uint16(0x225))
			m.cycles += 8
		}
	case 0x4295:
		{
			m.ip = 0x4299
			m.r[2] = m.rd16(m.r[11], uint16(m.r[3]+0x233))
			m.cycles += 8
		}
	case 0x4299:
		{
			m.ip = 0x429d
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x102d))
			m.cycles += 8
		}
	case 0x429d:
		{
			m.ip = 0x42a1
			m.r[3] = uint16(0x1023)
			m.cycles += 3
		}
	case 0x42a1:
		{
			m.ip = 0x42a3
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x42a3:
		{
			m.ip = 0x42a5
			m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3])))
			m.cycles += 16
		}
	case 0x42a5:
		{
			m.ip = 0x42a7
			if m.sf != m.of {
				m.ip = uint16(17069)
			}
			m.cycles += 8
		}
	case 0x42a7:
		{
			m.ip = 0x42a9
			m.set8(0, 0, m.unary("inc", 8, ((m.r[0]>>0)&255)))
			m.cycles += 3
		}
	case 0x42a9:
		{
			m.ip = 0x42ab
			m.r[2] = m.alu("sub", 16, m.r[2], m.rd16(m.r[11], uint16(m.r[3])))
			m.cycles += 16
		}
	case 0x42ab:
		{
			m.ip = 0x42ad
			m.ip = uint16(17059)
			m.cycles += 15
		}
	case 0x42ad:
		{
			m.ip = 0x42b0
			target := uint16(17102)
			m.push(0x42b0)
			m.ip = target
			m.cycles += 19
		}
	case 0x42b0:
		{
			m.ip = 0x42b3
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x42b3:
		{
			m.ip = 0x42b6
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x42b6:
		{
			m.ip = 0x42b8
			m.set8(0, 0, uint16(10))
			m.cycles += 2
		}
	case 0x42b8:
		{
			m.ip = 0x42bb
			m.alu("sub", 8, m.rd8(m.r[11], uint16(m.r[3]+0xfffe)), ((m.r[0] >> 0) & 255))
			m.cycles += 16
		}
	case 0x42bb:
		{
			m.ip = 0x42bd
			if !m.zf {
				m.ip = uint16(17057)
			}
			m.cycles += 8
		}
	case 0x42bd:
		{
			m.ip = 0x42bf
			m.set8(0, 0, ((m.r[2] >> 0) & 255))
			m.cycles += 2
		}
	case 0x42bf:
		{
			m.ip = 0x42c2
			target := uint16(17102)
			m.push(0x42c2)
			m.ip = target
			m.cycles += 19
		}
	case 0x42c2:
		{
			m.ip = 0x42c5
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x42c5:
		{
			m.ip = 0x42c7
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x42c7:
		{
			m.ip = 0x42ca
			target := uint16(17102)
			m.push(0x42ca)
			m.ip = target
			m.cycles += 19
		}
	case 0x42ca:
		{
			m.ip = 0x42cb
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x42cb:
		{
			m.ip = 0x42cc
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x42cc:
		{
			m.ip = 0x42cd
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x42cd:
		{
			m.ip = 0x42ce
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x42ce:
		{
			m.ip = 0x42cf
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x42cf:
		{
			m.ip = 0x42d1
			m.r[0] = m.r[8]
			m.cycles += 2
		}
	case 0x42d1:
		{
			m.ip = 0x42d5
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(0x276)))
			m.cycles += 16
		}
	case 0x42d5:
		{
			m.ip = 0x42d6
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x42d6:
		{
			m.ip = 0x42d8
			if m.zf {
				m.ip = uint16(17152)
			}
			m.cycles += 8
		}
	case 0x42d8:
		{
			m.ip = 0x42dc
			m.r[6] = m.rd16(m.r[11], uint16(0x102b))
			m.cycles += 8
		}
	case 0x42dc:
		{
			m.ip = 0x42de
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x42de:
		{
			m.ip = 0x42e0
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), ((m.r[0]>>0)&255)))
			m.cycles += 4
		}
	case 0x42e0:
		{
			m.ip = 0x42e2
			m.r[6] = m.alu("add", 16, m.r[6], m.r[0])
			m.cycles += 4
		}
	case 0x42e2:
		{
			m.ip = 0x42e5
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x42e5:
		{
			m.ip = 0x42e8
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x42e8:
		{
			m.ip = 0x42eb
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x42eb:
		{
			m.ip = 0x42ef
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6]+0x1)))
			m.cycles += 8
		}
	case 0x42ef:
		{
			m.ip = 0x42f3
			m.wr8(m.r[8], uint16(m.r[7]+0x1), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x42f3:
		{
			m.ip = 0x42f6
			m.r[6] = m.alu("add", 16, m.r[6], uint16(80))
			m.cycles += 4
		}
	case 0x42f6:
		{
			m.ip = 0x42f9
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x42f9:
		{
			m.ip = 0x42fb
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17125)
			}
			m.cycles += 17
		}
	case 0x42fb:
		{
			m.ip = 0x42ff
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(560))
			m.cycles += 4
		}
	case 0x42ff:
		{
			m.ip = 0x4300
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4300:
		{
			m.ip = 0x4302
			m.set8(0, 8, uint16(4))
			m.cycles += 2
		}
	case 0x4302:
		{
			m.ip = 0x4304
			m.set8(0, 0, m.alu("add", 8, ((m.r[0]>>0)&255), uint16(48)))
			m.cycles += 4
		}
	case 0x4304:
		{
			m.ip = 0x4307
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x4307:
		{
			m.ip = 0x4308
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4308:
		{
			m.ip = 0x430a
			m.set8(0, 0, m.input(uint16(97)))
			m.cycles += 8
		}
	case 0x430a:
		{
			m.ip = 0x430c
			m.set8(0, 0, m.alu("and", 8, ((m.r[0]>>0)&255), uint16(252)))
			m.cycles += 4
		}
	case 0x430c:
		{
			m.ip = 0x430e
			m.output(uint16(97), ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x430e:
		{
			m.ip = 0x4311
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x4311:
		{
			m.ip = 0x4313
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x4313:
		{
			m.ip = 0x4315
			m.set8(0, 8, uint16(76))
			m.cycles += 2
		}
	case 0x4315:
		{
			m.ip = 0x4317
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x4317:
		{
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4442:
		{
			m.ip = 0x4445
			m.r[0] = uint16(517)
			m.cycles += 2
		}
	case 0x4445:
		{
			m.ip = 0x4448
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4448:
		{
			m.ip = 0x4449
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4449:
		{
			m.ip = 0x444b
			m.set8(0, 0, uint16(3))
			m.cycles += 2
		}
	case 0x444b:
		{
			m.ip = 0x444d
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x444d:
		{
			m.ip = 0x444e
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x444e:
		{
			m.ip = 0x4454
			m.wr16(m.r[11], uint16(0x499), uint16(510))
			m.cycles += 8
		}
	case 0x4454:
		{
			m.ip = 0x445a
			m.wr16(m.r[11], uint16(0x49b), uint16(255))
			m.cycles += 8
		}
	case 0x445a:
		{
			m.ip = 0x445b
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x445b:
		{
			m.ip = 0x445c
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x445c:
		{
			m.ip = 0x445f
			target := uint16(17640)
			m.push(0x445f)
			m.ip = target
			m.cycles += 19
		}
	case 0x445f:
		{
			m.ip = 0x4462
			target := uint16(17738)
			m.push(0x4462)
			m.ip = target
			m.cycles += 19
		}
	case 0x4462:
		{
			m.ip = 0x4463
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x4463:
		{
			m.ip = 0x4464
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x4464:
		{
			m.ip = 0x4468
			m.set8(1, 0, m.rd8(m.r[11], uint16(0x49d)))
			m.cycles += 8
		}
	case 0x4468:
		{
			m.ip = 0x446b
			m.r[1] = m.alu("or", 16, m.r[1], uint16(16))
			m.cycles += 4
		}
	case 0x446b:
		{
			m.ip = 0x446d
			m.set8(1, 0, m.shift("shr", 8, ((m.r[1]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x446d:
		{
			m.ip = 0x446f
			if !m.cf {
				m.ip = uint16(17515)
			}
			m.cycles += 8
		}
	case 0x446f:
		{
			m.ip = 0x4470
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4470:
		{
			m.ip = 0x4474
			m.set8(1, 0, m.rd8(m.r[11], uint16(0x49d)))
			m.cycles += 8
		}
	case 0x4474:
		{
			m.ip = 0x4478
			m.wr16(m.r[11], uint16(0x499), m.shift("shr", 16, m.rd16(m.r[11], uint16(0x499)), ((m.r[1]>>0)&255)))
			m.cycles += 8
		}
	case 0x4478:
		{
			m.ip = 0x447d
			m.wr16(m.r[11], uint16(0x499), m.alu("and", 16, m.rd16(m.r[11], uint16(0x499)), uint16(65534)))
			m.cycles += 16
		}
	case 0x447d:
		{
			m.ip = 0x4481
			m.wr16(m.r[11], uint16(0x49b), m.shift("shr", 16, m.rd16(m.r[11], uint16(0x49b)), ((m.r[1]>>0)&255)))
			m.cycles += 8
		}
	case 0x4481:
		{
			m.ip = 0x4484
			m.r[5] = uint16(13920)
			m.cycles += 2
		}
	case 0x4484:
		{
			m.ip = 0x4487
			m.r[1] = uint16(1740)
			m.cycles += 2
		}
	case 0x4487:
		{
			m.ip = 0x448a
			target := uint16(18515)
			m.push(0x448a)
			m.ip = target
			m.cycles += 19
		}
	case 0x448a:
		{
			m.ip = 0x448d
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x448d:
		{
			m.ip = 0x448f
			if m.zf {
				m.ip = uint16(17553)
			}
			m.cycles += 8
		}
	case 0x448f:
		{
			m.ip = 0x4490
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4490:
		{
			m.ip = 0x4491
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4491:
		{
			m.ip = 0x4492
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4492:
		{
			m.ip = 0x4493
			m.push(m.r[7])
			m.cycles += 11
		}
	case 0x4493:
		{
			m.ip = 0x4497
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x4497:
		{
			m.ip = 0x449b
			m.r[7] = m.alu("add", 16, m.r[7], m.rd16(m.r[8], uint16(m.r[5]+0x0)))
			m.cycles += 16
		}
	case 0x449b:
		{
			m.ip = 0x449f
			m.r[6] = m.rd16(m.r[8], uint16(m.r[5]+0x2))
			m.cycles += 8
		}
	case 0x449f:
		{
			m.ip = 0x44a2
			m.r[5] = m.alu("add", 16, m.r[5], uint16(4))
			m.cycles += 4
		}
	case 0x44a2:
		{
			m.ip = 0x44a5
			m.r[1] = uint16(8)
			m.cycles += 2
		}
	case 0x44a5:
		{
			m.ip = 0x44aa
			m.wr8(m.r[11], uint16(0x1031), uint16(128))
			m.cycles += 8
		}
	case 0x44aa:
		{
			m.ip = 0x44ae
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x44ae:
		{
			m.ip = 0x44af
			m.push(m.r[6])
			m.cycles += 11
		}
	case 0x44af:
		{
			m.ip = 0x44b2
			target := uint16(17878)
			m.push(0x44b2)
			m.ip = target
			m.cycles += 19
		}
	case 0x44b2:
		{
			m.ip = 0x44b4
			m.set8(3, 0, uint16(15))
			m.cycles += 2
		}
	case 0x44b4:
		{
			m.ip = 0x44b8
			m.alu("sub", 16, m.r[6], uint16(13919))
			m.cycles += 4
		}
	case 0x44b8:
		{
			m.ip = 0x44ba
			if !m.cf && !m.zf {
				m.ip = uint16(17597)
			}
			m.cycles += 8
		}
	case 0x44ba:
		{
			m.ip = 0x44bd
			m.set8(3, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x44bd:
		{
			m.ip = 0x44be
			m.r[6] = m.pop()
			m.cycles += 8
		}
	case 0x44be:
		{
			m.ip = 0x44bf
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x44bf:
		{
			m.ip = 0x44c3
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x44c3:
		{
			m.ip = 0x44c5
			m.set8(0, 0, uint16(8))
			m.cycles += 2
		}
	case 0x44c5:
		{
			m.ip = 0x44c9
			m.set8(0, 8, m.rd8(m.r[11], uint16(0x1031)))
			m.cycles += 8
		}
	case 0x44c9:
		{
			m.ip = 0x44ca
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x44ca:
		{
			m.ip = 0x44ce
			m.wr8(m.r[11], uint16(0x1031), m.shift("shr", 8, m.rd8(m.r[11], uint16(0x1031)), uint16(1)))
			m.cycles += 8
		}
	case 0x44ce:
		{
			m.ip = 0x44d1
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[7])))
			m.cycles += 8
		}
	case 0x44d1:
		{
			m.ip = 0x44d4
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x44d4:
		{
			m.ip = 0x44d6
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17578)
			}
			m.cycles += 17
		}
	case 0x44d6:
		{
			m.ip = 0x44d7
			m.r[7] = m.pop()
			m.cycles += 8
		}
	case 0x44d7:
		{
			m.ip = 0x44d8
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x44d8:
		{
			m.ip = 0x44da
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17543)
			}
			m.cycles += 17
		}
	case 0x44da:
		{
			m.ip = 0x44db
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x44db:
		{
			m.ip = 0x44dd
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17519)
			}
			m.cycles += 17
		}
	case 0x44dd:
		{
			m.ip = 0x44e0
			m.r[0] = uint16(261)
			m.cycles += 2
		}
	case 0x44e0:
		{
			m.ip = 0x44e3
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x44e3:
		{
			m.ip = 0x44e4
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x44e4:
		{
			m.ip = 0x44e7
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x44e7:
		{
			m.ip = 0x44e8
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x44e8:
		{
			m.ip = 0x44eb
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x44eb:
		{
			m.ip = 0x44ee
			m.r[1] = uint16(87)
			m.cycles += 2
		}
	case 0x44ee:
		{
			m.ip = 0x44ef
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x44ef:
		{
			m.ip = 0x44f2
			m.r[1] = uint16(20)
			m.cycles += 2
		}
	case 0x44f2:
		{
			m.ip = 0x44f3
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x44f3:
		{
			m.ip = 0x44f7
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x44f7:
		{
			m.ip = 0x44fa
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x44fa:
		{
			m.ip = 0x44fd
			m.r[0] = uint16(772)
			m.cycles += 2
		}
	case 0x44fd:
		{
			m.ip = 0x44fe
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x44fe:
		{
			m.ip = 0x4501
			m.set8(3, 8, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4501:
		{
			m.ip = 0x4504
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4504:
		{
			m.ip = 0x4507
			m.r[0] = uint16(516)
			m.cycles += 2
		}
	case 0x4507:
		{
			m.ip = 0x4508
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4508:
		{
			m.ip = 0x450b
			m.set8(3, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x450b:
		{
			m.ip = 0x450e
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x450e:
		{
			m.ip = 0x4511
			m.r[0] = uint16(260)
			m.cycles += 2
		}
	case 0x4511:
		{
			m.ip = 0x4512
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4512:
		{
			m.ip = 0x4515
			m.set8(1, 8, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4515:
		{
			m.ip = 0x4518
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4518:
		{
			m.ip = 0x451b
			m.r[0] = uint16(4)
			m.cycles += 2
		}
	case 0x451b:
		{
			m.ip = 0x451c
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x451c:
		{
			m.ip = 0x451f
			m.set8(1, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x451f:
		{
			m.ip = 0x4521
			m.r[0] = m.r[1]
			m.cycles += 2
		}
	case 0x4521:
		{
			m.ip = 0x4525
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x4525:
		{
			m.ip = 0x4528
			m.r[1] = uint16(8)
			m.cycles += 2
		}
	case 0x4528:
		{
			m.ip = 0x452a
			m.set8(3, 8, m.shift("shl", 8, ((m.r[3]>>8)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x452a:
		{
			m.ip = 0x452d
			m.wr8(m.r[8], uint16(m.r[7]), m.shift("rcl", 8, m.rd8(m.r[8], uint16(m.r[7])), uint16(1)))
			m.cycles += 8
		}
	case 0x452d:
		{
			m.ip = 0x452f
			m.set8(3, 0, m.shift("shl", 8, ((m.r[3]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x452f:
		{
			m.ip = 0x4532
			m.wr8(m.r[8], uint16(m.r[7]), m.shift("rcl", 8, m.rd8(m.r[8], uint16(m.r[7])), uint16(1)))
			m.cycles += 8
		}
	case 0x4532:
		{
			m.ip = 0x4534
			m.set8(0, 8, m.shift("shl", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x4534:
		{
			m.ip = 0x4537
			m.wr8(m.r[8], uint16(m.r[7]), m.shift("rcl", 8, m.rd8(m.r[8], uint16(m.r[7])), uint16(1)))
			m.cycles += 8
		}
	case 0x4537:
		{
			m.ip = 0x4539
			m.set8(0, 0, m.shift("shl", 8, ((m.r[0]>>0)&255), uint16(1)))
			m.cycles += 8
		}
	case 0x4539:
		{
			m.ip = 0x453c
			m.wr8(m.r[8], uint16(m.r[7]), m.shift("rcl", 8, m.rd8(m.r[8], uint16(m.r[7])), uint16(1)))
			m.cycles += 8
		}
	case 0x453c:
		{
			m.ip = 0x453d
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x453d:
		{
			m.ip = 0x453f
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17704)
			}
			m.cycles += 17
		}
	case 0x453f:
		{
			m.ip = 0x4540
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x4540:
		{
			m.ip = 0x4541
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4541:
		{
			m.ip = 0x4543
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17650)
			}
			m.cycles += 17
		}
	case 0x4543:
		{
			m.ip = 0x4546
			m.r[6] = m.alu("add", 16, m.r[6], uint16(60))
			m.cycles += 4
		}
	case 0x4546:
		{
			m.ip = 0x4547
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4547:
		{
			m.ip = 0x4549
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17646)
			}
			m.cycles += 17
		}
	case 0x4549:
		{
			m.ip = 0x454a
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x454a:
		{
			m.ip = 0x454e
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x454e:
		{
			m.ip = 0x4551
			m.r[5] = uint16(0)
			m.cycles += 2
		}
	case 0x4551:
		{
			m.ip = 0x4554
			m.r[7] = uint16(0)
			m.cycles += 2
		}
	case 0x4554:
		{
			m.ip = 0x4557
			m.r[0] = uint16(0)
			m.cycles += 2
		}
	case 0x4557:
		{
			m.ip = 0x455a
			m.r[1] = uint16(87)
			m.cycles += 2
		}
	case 0x455a:
		{
			m.ip = 0x455b
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x455b:
		{
			m.ip = 0x455e
			m.r[1] = uint16(20)
			m.cycles += 2
		}
	case 0x455e:
		{
			m.ip = 0x4563
			m.wr16(m.r[8], uint16(m.r[7]+0x5190), m.r[5])
			m.cycles += 8
		}
	case 0x4563:
		{
			m.ip = 0x4566
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x4566:
		{
			m.ip = 0x456b
			m.wr16(m.r[8], uint16(m.r[7]+0x5190), m.r[0])
			m.cycles += 8
		}
	case 0x456b:
		{
			m.ip = 0x456e
			m.r[7] = m.alu("add", 16, m.r[7], uint16(2))
			m.cycles += 4
		}
	case 0x456e:
		{
			m.ip = 0x4571
			m.r[0] = m.alu("add", 16, m.r[0], uint16(8))
			m.cycles += 4
		}
	case 0x4571:
		{
			m.ip = 0x4572
			m.r[5] = m.unary("inc", 16, m.r[5])
			m.cycles += 3
		}
	case 0x4572:
		{
			m.ip = 0x4574
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17758)
			}
			m.cycles += 17
		}
	case 0x4574:
		{
			m.ip = 0x4577
			m.r[5] = m.alu("add", 16, m.r[5], uint16(60))
			m.cycles += 4
		}
	case 0x4577:
		{
			m.ip = 0x4578
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4578:
		{
			m.ip = 0x457a
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(17754)
			}
			m.cycles += 17
		}
	case 0x457a:
		{
			m.ip = 0x457d
			m.r[5] = uint16(27836)
			m.cycles += 2
		}
	case 0x457d:
		{
			m.ip = 0x4580
			m.r[7] = uint16(13920)
			m.cycles += 2
		}
	case 0x4580:
		{
			m.ip = 0x4583
			target := uint16(18980)
			m.push(0x4583)
			m.ip = target
			m.cycles += 19
		}
	case 0x4583:
		{
			m.ip = 0x4586
			m.set8(0, 8, m.alu("and", 8, ((m.r[0]>>8)&255), uint16(7)))
			m.cycles += 4
		}
	case 0x4586:
		{
			m.ip = 0x4589
			m.alu("sub", 16, m.r[0], uint16(1739))
			m.cycles += 4
		}
	case 0x4589:
		{
			m.ip = 0x458b
			if m.cf || m.zf {
				m.ip = uint16(17806)
			}
			m.cycles += 8
		}
	case 0x458b:
		{
			m.ip = 0x458e
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(308))
			m.cycles += 4
		}
	case 0x458e:
		{
			m.ip = 0x4590
			m.r[0] = m.shift("shl", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x4590:
		{
			m.ip = 0x4592
			m.r[0] = m.shift("shl", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x4592:
		{
			m.ip = 0x4594
			m.r[3] = m.r[0]
			m.cycles += 2
		}
	case 0x4594:
		{
			m.ip = 0x4599
			m.r[0] = m.rd16(m.r[8], uint16(m.r[3]+0x5190))
			m.cycles += 8
		}
	case 0x4599:
		{
			m.ip = 0x459c
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x459c:
		{
			m.ip = 0x459e
			if !m.zf {
				m.ip = uint16(17829)
			}
			m.cycles += 8
		}
	case 0x459e:
		{
			m.ip = 0x45a0
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x45a0:
		{
			m.ip = 0x45a3
			m.set8(3, 0, m.alu("and", 8, ((m.r[3]>>0)&255), uint16(252)))
			m.cycles += 4
		}
	case 0x45a3:
		{
			m.ip = 0x45a5
			m.ip = uint16(17812)
			m.cycles += 15
		}
	case 0x45a5:
		{
			m.ip = 0x45a8
			m.wr16(m.r[8], uint16(m.r[7]), m.r[0])
			m.cycles += 8
		}
	case 0x45a8:
		{
			m.ip = 0x45ad
			m.r[2] = m.rd16(m.r[8], uint16(m.r[3]+0x5192))
			m.cycles += 8
		}
	case 0x45ad:
		{
			m.ip = 0x45b1
			m.wr16(m.r[8], uint16(m.r[7]+0x2), m.r[2])
			m.cycles += 8
		}
	case 0x45b1:
		{
			m.ip = 0x45b4
			m.r[7] = m.alu("add", 16, m.r[7], uint16(4))
			m.cycles += 4
		}
	case 0x45b4:
		{
			m.ip = 0x45b8
			m.r[0] = m.rd16(m.r[8], uint16(m.r[5]+0x0))
			m.cycles += 8
		}
	case 0x45b8:
		{
			m.ip = 0x45bc
			m.r[2] = m.rd16(m.r[8], uint16(m.r[5]+0x2))
			m.cycles += 8
		}
	case 0x45bc:
		{
			m.ip = 0x45c1
			m.wr16(m.r[8], uint16(m.r[3]+0x5190), m.r[0])
			m.cycles += 8
		}
	case 0x45c1:
		{
			m.ip = 0x45c6
			m.wr16(m.r[8], uint16(m.r[3]+0x5192), m.r[2])
			m.cycles += 8
		}
	case 0x45c6:
		{
			m.ip = 0x45cc
			m.wr16(m.r[8], uint16(m.r[5]+0x0), uint16(65535))
			m.cycles += 8
		}
	case 0x45cc:
		{
			m.ip = 0x45cf
			m.r[5] = m.alu("sub", 16, m.r[5], uint16(4))
			m.cycles += 4
		}
	case 0x45cf:
		{
			m.ip = 0x45d3
			m.alu("sub", 16, m.r[7], uint16(20880))
			m.cycles += 4
		}
	case 0x45d3:
		{
			m.ip = 0x45d5
			if !m.zf {
				m.ip = uint16(17792)
			}
			m.cycles += 8
		}
	case 0x45d5:
		{
			m.ip = 0x45d6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x45d6:
		{
			m.ip = 0x45d7
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x45d7:
		{
			m.ip = 0x45db
			m.r[3] = m.rd16(m.r[11], uint16(0x499))
			m.cycles += 8
		}
	case 0x45db:
		{
			m.ip = 0x45df
			m.r[1] = m.rd16(m.r[11], uint16(0x49b))
			m.cycles += 8
		}
	case 0x45df:
		{
			m.ip = 0x45e2
			target := uint16(18980)
			m.push(0x45e2)
			m.ip = target
			m.cycles += 19
		}
	case 0x45e2:
		{
			m.ip = 0x45e3
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x45e3:
		{
			m.ip = 0x45e5
			m.r[0] = m.alu("and", 16, m.r[0], m.r[3])
			m.cycles += 4
		}
	case 0x45e5:
		{
			m.ip = 0x45e7
			m.r[6] = m.alu("add", 16, m.r[6], m.r[0])
			m.cycles += 4
		}
	case 0x45e7:
		{
			m.ip = 0x45e9
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[1])
			m.cycles += 4
		}
	case 0x45e9:
		{
			m.ip = 0x45ea
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x45ea:
		{
			m.ip = 0x45ee
			m.wr8(m.r[11], uint16(0x296), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x45ee:
		{
			m.ip = 0x45f0
			m.set8(3, 0, ((m.r[3] >> 8) & 255))
			m.cycles += 2
		}
	case 0x45f0:
		{
			m.ip = 0x45f4
			m.set8(3, 8, m.rd8(m.r[11], uint16(0x296)))
			m.cycles += 8
		}
	case 0x45f4:
		{
			m.ip = 0x45f6
			m.r[0] = m.alu("and", 16, m.r[0], m.r[3])
			m.cycles += 4
		}
	case 0x45f6:
		{
			m.ip = 0x45f8
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x45f8:
		{
			m.ip = 0x45fa
			m.r[6] = m.alu("add", 16, m.r[6], m.r[0])
			m.cycles += 4
		}
	case 0x45fa:
		{
			m.ip = 0x45fc
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x45fc:
		{
			m.ip = 0x45fe
			m.r[0] = m.shift("shr", 16, m.r[0], uint16(1))
			m.cycles += 8
		}
	case 0x45fe:
		{
			m.ip = 0x4600
			m.r[6] = m.alu("add", 16, m.r[6], m.r[0])
			m.cycles += 4
		}
	case 0x4600:
		{
			m.ip = 0x4604
			m.wr8(m.r[11], uint16(0x296), ((m.r[1] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4604:
		{
			m.ip = 0x4606
			m.set8(1, 0, ((m.r[1] >> 8) & 255))
			m.cycles += 2
		}
	case 0x4606:
		{
			m.ip = 0x460a
			m.set8(1, 8, m.rd8(m.r[11], uint16(0x296)))
			m.cycles += 8
		}
	case 0x460a:
		{
			m.ip = 0x460c
			m.r[1] = m.shift("shr", 16, m.r[1], uint16(1))
			m.cycles += 8
		}
	case 0x460c:
		{
			m.ip = 0x460e
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[1])
			m.cycles += 4
		}
	case 0x460e:
		{
			m.ip = 0x4610
			m.r[1] = m.shift("shr", 16, m.r[1], uint16(1))
			m.cycles += 8
		}
	case 0x4610:
		{
			m.ip = 0x4612
			m.r[1] = m.shift("shr", 16, m.r[1], uint16(1))
			m.cycles += 8
		}
	case 0x4612:
		{
			m.ip = 0x4614
			m.r[6] = m.alu("sub", 16, m.r[6], m.r[1])
			m.cycles += 4
		}
	case 0x4614:
		{
			m.ip = 0x4615
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4615:
		{
			m.ip = 0x4616
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4616:
		{
			m.ip = 0x4619
			target := uint16(17966)
			m.push(0x4619)
			m.ip = target
			m.cycles += 19
		}
	case 0x4619:
		{
			m.ip = 0x461e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x3b2)), uint16(1))
			m.cycles += 16
		}
	case 0x461e:
		{
			m.ip = 0x4620
			if m.zf {
				m.ip = uint16(17965)
			}
			m.cycles += 8
		}
	case 0x4620:
		{
			m.ip = 0x4623
			m.wr16(m.r[11], uint16(0x3aa), m.r[0])
			m.cycles += 8
		}
	case 0x4623:
		{
			m.ip = 0x4626
			target := uint16(18004)
			m.push(0x4626)
			m.ip = target
			m.cycles += 19
		}
	case 0x4626:
		{
			m.ip = 0x462a
			m.r[3] = m.rd16(m.r[11], uint16(0x3aa))
			m.cycles += 8
		}
	case 0x462a:
		{
			m.ip = 0x462d
			target := uint16(18132)
			m.push(0x462d)
			m.ip = target
			m.cycles += 19
		}
	case 0x462d:
		{
			m.ip = 0x462e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x462e:
		{
			m.ip = 0x4630
			m.set8(0, 8, uint16(61))
			m.cycles += 2
		}
	case 0x4630:
		{
			m.ip = 0x4632
			m.set8(0, 0, uint16(2))
			m.cycles += 2
		}
	case 0x4632:
		{
			m.ip = 0x4634
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x4634:
		{
			m.ip = 0x4636
			if m.cf {
				m.ip = uint16(17975)
			}
			m.cycles += 8
		}
	case 0x4636:
		{
			m.ip = 0x4637
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4637:
		{
			m.ip = 0x463b
			m.alu("sub", 16, m.r[2], uint16(886))
			m.cycles += 4
		}
	case 0x463b:
		{
			m.ip = 0x463d
			if !m.zf {
				m.ip = uint16(17987)
			}
			m.cycles += 8
		}
	case 0x463d:
		{
			m.ip = 0x4642
			m.wr8(m.r[11], uint16(0x3b2), uint16(1))
			m.cycles += 8
		}
	case 0x4642:
		{
			m.ip = 0x4643
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4643:
		{
			m.ip = 0x4646
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x4646:
		{
			m.ip = 0x4648
			m.interrupt(uint16(16))
			m.cycles += 50
		}
	case 0x4648:
		{
			m.ip = 0x464c
			m.r[2] = uint16(0x3b3)
			m.cycles += 3
		}
	case 0x464c:
		{
			m.ip = 0x464e
			m.set8(0, 8, uint16(9))
			m.cycles += 2
		}
	case 0x464e:
		{
			m.ip = 0x4650
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x4650:
		{
			m.ip = 0x4652
			m.set8(0, 8, uint16(76))
			m.cycles += 2
		}
	case 0x4652:
		{
			m.ip = 0x4654
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x4654:
		{
			m.ip = 0x4656
			m.r[5] = m.r[2]
			m.cycles += 2
		}
	case 0x4656:
		{
			m.ip = 0x4658
			m.set8(0, 8, uint16(66))
			m.cycles += 2
		}
	case 0x4658:
		{
			m.ip = 0x465a
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x465a:
		{
			m.ip = 0x465e
			m.r[3] = m.rd16(m.r[11], uint16(0x3aa))
			m.cycles += 8
		}
	case 0x465e:
		{
			m.ip = 0x4661
			m.r[1] = uint16(0)
			m.cycles += 2
		}
	case 0x4661:
		{
			m.ip = 0x4664
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x4664:
		{
			m.ip = 0x4666
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x4666:
		{
			m.ip = 0x4669
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x4669:
		{
			m.ip = 0x466d
			m.alu("sub", 16, m.r[5], uint16(886))
			m.cycles += 4
		}
	case 0x466d:
		{
			m.ip = 0x466f
			if m.zf {
				m.ip = uint16(18045)
			}
			m.cycles += 8
		}
	case 0x466f:
		{
			m.ip = 0x4671
			m.set8(0, 8, uint16(63))
			m.cycles += 2
		}
	case 0x4671:
		{
			m.ip = 0x4674
			m.r[1] = uint16(16)
			m.cycles += 2
		}
	case 0x4674:
		{
			m.ip = 0x4678
			m.r[2] = uint16(0x1045)
			m.cycles += 3
		}
	case 0x4678:
		{
			m.ip = 0x467a
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x467a:
		{
			m.ip = 0x467d
			m.r[2] = uint16(16)
			m.cycles += 2
		}
	case 0x467d:
		{
			m.ip = 0x467f
			m.set8(0, 8, uint16(66))
			m.cycles += 2
		}
	case 0x467f:
		{
			m.ip = 0x4681
			m.set8(0, 0, uint16(2))
			m.cycles += 2
		}
	case 0x4681:
		{
			m.ip = 0x4684
			m.r[1] = uint16(0)
			m.cycles += 2
		}
	case 0x4684:
		{
			m.ip = 0x4688
			m.r[3] = m.rd16(m.r[11], uint16(0x3aa))
			m.cycles += 8
		}
	case 0x4688:
		{
			m.ip = 0x468a
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x468a:
		{
			m.ip = 0x468d
			m.wr16(m.r[11], uint16(0x3ae), m.r[0])
			m.cycles += 8
		}
	case 0x468d:
		{
			m.ip = 0x4691
			m.wr16(m.r[11], uint16(0x3b0), m.r[2])
			m.cycles += 8
		}
	case 0x4691:
		{
			m.ip = 0x4694
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x4694:
		{
			m.ip = 0x4698
			m.alu("sub", 16, m.r[5], uint16(886))
			m.cycles += 4
		}
	case 0x4698:
		{
			m.ip = 0x469a
			if m.zf {
				m.ip = uint16(18077)
			}
			m.cycles += 8
		}
	case 0x469a:
		{
			m.ip = 0x469d
			m.r[2] = uint16(16)
			m.cycles += 2
		}
	case 0x469d:
		{
			m.ip = 0x469f
			m.set8(0, 8, uint16(66))
			m.cycles += 2
		}
	case 0x469f:
		{
			m.ip = 0x46a1
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x46a1:
		{
			m.ip = 0x46a3
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x46a3:
		{
			m.ip = 0x46a7
			m.r[8] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x46a7:
		{
			m.ip = 0x46ac
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x3b0)), uint16(0))
			m.cycles += 16
		}
	case 0x46ac:
		{
			m.ip = 0x46ae
			if m.zf {
				m.ip = uint16(18112)
			}
			m.cycles += 8
		}
	case 0x46ae:
		{
			m.ip = 0x46b1
			m.r[1] = uint16(65520)
			m.cycles += 2
		}
	case 0x46b1:
		{
			m.ip = 0x46b4
			target := uint16(18116)
			m.push(0x46b4)
			m.ip = target
			m.cycles += 19
		}
	case 0x46b4:
		{
			m.ip = 0x46b6
			m.r[0] = m.r[8]
			m.cycles += 2
		}
	case 0x46b6:
		{
			m.ip = 0x46b9
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4095))
			m.cycles += 4
		}
	case 0x46b9:
		{
			m.ip = 0x46bb
			m.r[8] = m.r[0]
			m.cycles += 2
		}
	case 0x46bb:
		{
			m.ip = 0x46c0
			m.wr16(m.r[11], uint16(0x3ae), m.alu("add", 16, m.rd16(m.r[11], uint16(0x3ae)), uint16(16)))
			m.cycles += 16
		}
	case 0x46c0:
		{
			m.ip = 0x46c4
			m.r[1] = m.rd16(m.r[11], uint16(0x3ae))
			m.cycles += 8
		}
	case 0x46c4:
		{
			m.ip = 0x46c8
			m.r[3] = m.rd16(m.r[11], uint16(0x3aa))
			m.cycles += 8
		}
	case 0x46c8:
		{
			m.ip = 0x46cb
			m.r[2] = uint16(0)
			m.cycles += 2
		}
	case 0x46cb:
		{
			m.ip = 0x46cd
			m.set8(0, 8, uint16(63))
			m.cycles += 2
		}
	case 0x46cd:
		{
			m.ip = 0x46ce
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x46ce:
		{
			m.ip = 0x46cf
			m.push(m.r[8])
			m.cycles += 11
		}
	case 0x46cf:
		{
			m.ip = 0x46d0
			m.r[11] = m.pop()
			m.cycles += 8
		}
	case 0x46d0:
		{
			m.ip = 0x46d2
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x46d2:
		{
			m.ip = 0x46d3
			m.r[11] = m.pop()
			m.cycles += 8
		}
	case 0x46d3:
		{
			m.ip = 0x46d4
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x46d4:
		{
			m.ip = 0x46d6
			m.set8(0, 8, uint16(62))
			m.cycles += 2
		}
	case 0x46d6:
		{
			m.ip = 0x46d8
			m.interrupt(uint16(33))
			m.cycles += 50
		}
	case 0x46d8:
		{
			m.ip = 0x46d9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x46d9:
		{
			m.ip = 0x46da
			m.push(m.r[11])
			m.cycles += 11
		}
	case 0x46da:
		{
			m.ip = 0x46dc
			m.r[8] = m.r[0]
			m.cycles += 2
		}
	case 0x46dc:
		{
			m.ip = 0x46e0
			m.r[11] = m.rd16(m.r[11], uint16(0x294))
			m.cycles += 8
		}
	case 0x46e0:
		{
			m.ip = 0x46e2
			m.r[6] = m.alu("xor", 16, m.r[6], m.r[6])
			m.cycles += 4
		}
	case 0x46e2:
		{
			m.ip = 0x46e4
			m.r[7] = m.alu("xor", 16, m.r[7], m.r[7])
			m.cycles += 4
		}
	case 0x46e4:
		{
			m.ip = 0x46e7
			m.r[0] = uint16(5)
			m.cycles += 2
		}
	case 0x46e7:
		{
			m.ip = 0x46ea
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x46ea:
		{
			m.ip = 0x46eb
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x46eb:
		{
			m.ip = 0x46ee
			m.r[0] = uint16(65288)
			m.cycles += 2
		}
	case 0x46ee:
		{
			m.ip = 0x46ef
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x46ef:
		{
			m.ip = 0x46f2
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x46f2:
		{
			m.ip = 0x46f3
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x46f3:
		{
			m.ip = 0x46f6
			m.r[0] = uint16(0)
			m.cycles += 2
		}
	case 0x46f6:
		{
			m.ip = 0x46f7
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x46f7:
		{
			m.ip = 0x46fa
			m.r[0] = uint16(3841)
			m.cycles += 2
		}
	case 0x46fa:
		{
			m.ip = 0x46fb
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x46fb:
		{
			m.ip = 0x46fe
			m.r[1] = uint16(28000)
			m.cycles += 2
		}
	case 0x46fe:
		{
			m.ip = 0x46ff
			m.df = false
			m.cycles += 2
		}
	case 0x46ff:
		{
			m.ip = 0x4701
			m.stringOp("stos", 8)
			m.cycles += 2
		}
	case 0x4701:
		{
			m.ip = 0x4704
			m.r[0] = uint16(0)
			m.cycles += 2
		}
	case 0x4704:
		{
			m.ip = 0x4705
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4705:
		{
			m.ip = 0x4707
			m.r[5] = m.alu("xor", 16, m.r[5], m.r[5])
			m.cycles += 4
		}
	case 0x4707:
		{
			m.ip = 0x470a
			m.r[1] = uint16(4)
			m.cycles += 2
		}
	case 0x470a:
		{
			m.ip = 0x470b
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x470b:
		{
			m.ip = 0x470d
			m.r[7] = m.alu("xor", 16, m.r[7], m.r[7])
			m.cycles += 4
		}
	case 0x470d:
		{
			m.ip = 0x470f
			m.set8(3, 8, uint16(1))
			m.cycles += 2
		}
	case 0x470f:
		{
			m.ip = 0x4712
			m.r[0] = uint16(4099)
			m.cycles += 2
		}
	case 0x4712:
		{
			m.ip = 0x4713
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4713:
		{
			m.ip = 0x4718
			m.r[0] = m.rd16(m.r[9], uint16(m.r[5]+0x477a))
			m.cycles += 8
		}
	case 0x4718:
		{
			m.ip = 0x4719
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4719:
		{
			m.ip = 0x471c
			m.r[1] = uint16(80)
			m.cycles += 2
		}
	case 0x471c:
		{
			m.ip = 0x471d
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x471d:
		{
			m.ip = 0x4720
			m.r[1] = uint16(350)
			m.cycles += 2
		}
	case 0x4720:
		{
			m.ip = 0x4722
			m.set8(3, 8, m.unary("dec", 8, ((m.r[3]>>8)&255)))
			m.cycles += 3
		}
	case 0x4722:
		{
			m.ip = 0x4725
			m.alu("and", 8, ((m.r[3] >> 8) & 255), uint16(127))
			m.cycles += 4
		}
	case 0x4725:
		{
			m.ip = 0x4727
			if m.zf {
				m.ip = uint16(18237)
			}
			m.cycles += 8
		}
	case 0x4727:
		{
			m.ip = 0x472a
			m.alu("and", 8, ((m.r[3] >> 8) & 255), uint16(128))
			m.cycles += 4
		}
	case 0x472a:
		{
			m.ip = 0x472c
			if !m.zf {
				m.ip = uint16(18265)
			}
			m.cycles += 8
		}
	case 0x472c:
		{
			m.ip = 0x472e
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x472e:
		{
			m.ip = 0x472f
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x472f:
		{
			m.ip = 0x4731
			if !m.zf {
				m.ip = uint16(18265)
			}
			m.cycles += 8
		}
	case 0x4731:
		{
			m.ip = 0x4732
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4732:
		{
			m.ip = 0x4734
			m.r[0] = m.r[11]
			m.cycles += 2
		}
	case 0x4734:
		{
			m.ip = 0x4737
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4096))
			m.cycles += 4
		}
	case 0x4737:
		{
			m.ip = 0x4739
			m.r[11] = m.r[0]
			m.cycles += 2
		}
	case 0x4739:
		{
			m.ip = 0x473a
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x473a:
		{
			m.ip = 0x473c
			m.ip = uint16(18265)
			m.cycles += 15
		}
	case 0x473c:
		{
			m.ip = 0x473d
			m.cycles += 3
		}
	case 0x473d:
		{
			m.ip = 0x473f
			m.set8(3, 8, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x473f:
		{
			m.ip = 0x4740
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x4740:
		{
			m.ip = 0x4742
			if !m.zf {
				m.ip = uint16(18251)
			}
			m.cycles += 8
		}
	case 0x4742:
		{
			m.ip = 0x4743
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4743:
		{
			m.ip = 0x4745
			m.r[0] = m.r[11]
			m.cycles += 2
		}
	case 0x4745:
		{
			m.ip = 0x4748
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4096))
			m.cycles += 4
		}
	case 0x4748:
		{
			m.ip = 0x474a
			m.r[11] = m.r[0]
			m.cycles += 2
		}
	case 0x474a:
		{
			m.ip = 0x474b
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x474b:
		{
			m.ip = 0x474d
			m.set8(3, 0, m.rd8(m.r[11], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x474d:
		{
			m.ip = 0x474e
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x474e:
		{
			m.ip = 0x4750
			if !m.zf {
				m.ip = uint16(18265)
			}
			m.cycles += 8
		}
	case 0x4750:
		{
			m.ip = 0x4751
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4751:
		{
			m.ip = 0x4753
			m.r[0] = m.r[11]
			m.cycles += 2
		}
	case 0x4753:
		{
			m.ip = 0x4756
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4096))
			m.cycles += 4
		}
	case 0x4756:
		{
			m.ip = 0x4758
			m.r[11] = m.r[0]
			m.cycles += 2
		}
	case 0x4758:
		{
			m.ip = 0x4759
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x4759:
		{
			m.ip = 0x475c
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[7])))
			m.cycles += 8
		}
	case 0x475c:
		{
			m.ip = 0x475f
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x475f:
		{
			m.ip = 0x4762
			m.r[7] = m.alu("add", 16, m.r[7], uint16(80))
			m.cycles += 4
		}
	case 0x4762:
		{
			m.ip = 0x4764
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(18208)
			}
			m.cycles += 17
		}
	case 0x4764:
		{
			m.ip = 0x4768
			m.r[7] = m.alu("sub", 16, m.r[7], uint16(27999))
			m.cycles += 4
		}
	case 0x4768:
		{
			m.ip = 0x4769
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4769:
		{
			m.ip = 0x476b
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(18204)
			}
			m.cycles += 17
		}
	case 0x476b:
		{
			m.ip = 0x476e
			m.r[5] = m.alu("add", 16, m.r[5], uint16(2))
			m.cycles += 4
		}
	case 0x476e:
		{
			m.ip = 0x476f
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x476f:
		{
			m.ip = 0x4771
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(18186)
			}
			m.cycles += 17
		}
	case 0x4771:
		{
			m.ip = 0x4774
			m.r[0] = uint16(261)
			m.cycles += 2
		}
	case 0x4774:
		{
			m.ip = 0x4777
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4777:
		{
			m.ip = 0x4778
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4778:
		{
			m.ip = 0x4779
			m.r[11] = m.pop()
			m.cycles += 8
		}
	case 0x4779:
		{
			m.ip = 0x477a
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4782:
		{
			m.ip = 0x4786
			m.alu("sub", 16, m.r[6], m.rd16(m.r[11], uint16(0x1032)))
			m.cycles += 16
		}
	case 0x4786:
		{
			m.ip = 0x4788
			if !m.zf {
				m.ip = uint16(18323)
			}
			m.cycles += 8
		}
	case 0x4788:
		{
			m.ip = 0x478c
			m.alu("sub", 16, m.r[6], uint16(4148))
			m.cycles += 4
		}
	case 0x478c:
		{
			m.ip = 0x478e
			if m.zf {
				m.ip = uint16(18323)
			}
			m.cycles += 8
		}
	case 0x478e:
		{
			m.ip = 0x4792
			m.wr16(m.r[11], uint16(0x1032), m.r[6])
			m.cycles += 8
		}
	case 0x4792:
		{
			m.ip = 0x4793
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4793:
		{
			m.ip = 0x4797
			m.wr16(m.r[11], uint16(0x1032), m.r[6])
			m.cycles += 8
		}
	case 0x4797:
		{
			m.ip = 0x479a
			m.r[0] = uint16(64)
			m.cycles += 2
		}
	case 0x479a:
		{
			m.ip = 0x479c
			m.r[8] = m.r[0]
			m.cycles += 2
		}
	case 0x479c:
		{
			m.ip = 0x47a1
			m.r[2] = m.rd16(m.r[8], uint16(0x63))
			m.cycles += 8
		}
	case 0x47a1:
		{
			m.ip = 0x47a4
			m.set8(2, 0, m.alu("add", 8, ((m.r[2]>>0)&255), uint16(6)))
			m.cycles += 4
		}
	case 0x47a4:
		{
			m.ip = 0x47a5
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x47a5:
		{
			m.ip = 0x47a6
			m.interrupts = false
			m.cycles += 2
		}
	case 0x47a6:
		{
			m.ip = 0x47a7
			m.set8(0, 0, m.input(m.r[2]))
			m.cycles += 8
		}
	case 0x47a7:
		{
			m.ip = 0x47a9
			m.set8(2, 0, uint16(192))
			m.cycles += 2
		}
	case 0x47a9:
		{
			m.ip = 0x47ac
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x47ac:
		{
			m.ip = 0x47ae
			m.set8(0, 0, ((m.r[3] >> 0) & 255))
			m.cycles += 2
		}
	case 0x47ae:
		{
			m.ip = 0x47af
			m.output(m.r[2], ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x47af:
		{
			m.ip = 0x47b1
			m.set8(0, 0, m.rd8(m.r[11], uint16(m.r[3]+m.r[6])))
			m.cycles += 8
		}
	case 0x47b1:
		{
			m.ip = 0x47b2
			m.output(m.r[2], ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x47b2:
		{
			m.ip = 0x47b4
			m.set8(3, 0, m.unary("inc", 8, ((m.r[3]>>0)&255)))
			m.cycles += 3
		}
	case 0x47b4:
		{
			m.ip = 0x47b7
			m.alu("sub", 8, ((m.r[3] >> 0) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x47b7:
		{
			m.ip = 0x47b9
			if !m.zf {
				m.ip = uint16(18348)
			}
			m.cycles += 8
		}
	case 0x47b9:
		{
			m.ip = 0x47ba
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x47ba:
		{
			m.ip = 0x47bb
			m.set8(0, 0, m.input(m.r[2]))
			m.cycles += 8
		}
	case 0x47bb:
		{
			m.ip = 0x47bd
			m.set8(2, 0, uint16(192))
			m.cycles += 2
		}
	case 0x47bd:
		{
			m.ip = 0x47bf
			m.set8(0, 0, uint16(32))
			m.cycles += 2
		}
	case 0x47bf:
		{
			m.ip = 0x47c0
			m.output(m.r[2], ((m.r[0] >> 0) & 255), 8)
			m.cycles += 8
		}
	case 0x47c0:
		{
			m.ip = 0x47c1
			m.interrupts = true
			m.cycles += 2
		}
	case 0x47c1:
		{
			m.ip = 0x47c2
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4845:
		{
			m.ip = 0x484a
			m.wr8(m.r[11], uint16(0x27f), uint16(0))
			m.cycles += 8
		}
	case 0x484a:
		{
			m.ip = 0x484c
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x484c:
		{
			m.ip = 0x484e
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x484e:
		{
			m.ip = 0x4850
			if m.zf {
				m.ip = uint16(18536)
			}
			m.cycles += 8
		}
	case 0x4850:
		{
			m.ip = 0x4852
			m.ip = uint16(18546)
			m.cycles += 15
		}
	case 0x4852:
		{
			m.ip = 0x4853
			m.cycles += 3
		}
	case 0x4853:
		{
			m.ip = 0x4855
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4855:
		{
			m.ip = 0x4857
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4857:
		{
			m.ip = 0x4859
			if m.zf {
				m.ip = uint16(18536)
			}
			m.cycles += 8
		}
	case 0x4859:
		{
			m.ip = 0x485b
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x485b:
		{
			m.ip = 0x485d
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x485d:
		{
			m.ip = 0x485e
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x485e:
		{
			m.ip = 0x4860
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4860:
		{
			m.ip = 0x4862
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4862:
		{
			m.ip = 0x4863
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x4863:
		{
			m.ip = 0x4865
			if !m.zf {
				m.ip = uint16(18521)
			}
			m.cycles += 8
		}
	case 0x4865:
		{
			m.ip = 0x4867
			m.ip = uint16(18546)
			m.cycles += 15
		}
	case 0x4867:
		{
			m.ip = 0x4868
			m.cycles += 3
		}
	case 0x4868:
		{
			m.ip = 0x486d
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x27d)), uint16(0))
			m.cycles += 16
		}
	case 0x486d:
		{
			m.ip = 0x486f
			if !m.zf {
				m.ip = uint16(18557)
			}
			m.cycles += 8
		}
	case 0x486f:
		{
			m.ip = 0x4872
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x4872:
		{
			m.ip = 0x4877
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(0))
			m.cycles += 16
		}
	case 0x4877:
		{
			m.ip = 0x4879
			if m.zf {
				m.ip = uint16(18556)
			}
			m.cycles += 8
		}
	case 0x4879:
		{
			m.ip = 0x487c
			target := uint16(5813)
			m.push(0x487c)
			m.ip = target
			m.cycles += 19
		}
	case 0x487c:
		{
			m.ip = 0x487d
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x487d:
		{
			m.ip = 0x487e
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x487e:
		{
			m.ip = 0x487f
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x487f:
		{
			m.ip = 0x4880
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x4880:
		{
			m.ip = 0x4885
			m.wr8(m.r[11], uint16(0x27f), uint16(1))
			m.cycles += 8
		}
	case 0x4885:
		{
			m.ip = 0x4888
			m.r[2] = uint16(513)
			m.cycles += 2
		}
	case 0x4888:
		{
			m.ip = 0x4889
			m.set8(0, 0, m.input(m.r[2]))
			m.cycles += 8
		}
	case 0x4889:
		{
			m.ip = 0x488b
			m.alu("and", 8, ((m.r[0] >> 0) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x488b:
		{
			m.ip = 0x488d
			if !m.zf {
				m.ip = uint16(18578)
			}
			m.cycles += 8
		}
	case 0x488d:
		{
			m.ip = 0x488f
			m.set8(0, 0, uint16(13))
			m.cycles += 2
		}
	case 0x488f:
		{
			m.ip = 0x4892
			m.ip = uint16(18706)
			m.cycles += 15
		}
	case 0x4892:
		{
			m.ip = 0x4894
			m.alu("and", 8, ((m.r[0] >> 0) & 255), uint16(32))
			m.cycles += 4
		}
	case 0x4894:
		{
			m.ip = 0x4896
			if !m.zf {
				m.ip = uint16(18587)
			}
			m.cycles += 8
		}
	case 0x4896:
		{
			m.ip = 0x4898
			m.set8(0, 0, uint16(27))
			m.cycles += 2
		}
	case 0x4898:
		{
			m.ip = 0x489a
			m.ip = uint16(18706)
			m.cycles += 15
		}
	case 0x489a:
		{
			m.ip = 0x489b
			m.cycles += 3
		}
	case 0x489b:
		{
			m.ip = 0x489e
			target := uint16(18605)
			m.push(0x489e)
			m.ip = target
			m.cycles += 19
		}
	case 0x489e:
		{
			m.ip = 0x489f
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x489f:
		{
			m.ip = 0x48a2
			target := uint16(18605)
			m.push(0x48a2)
			m.ip = target
			m.cycles += 19
		}
	case 0x48a2:
		{
			m.ip = 0x48a3
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x48a3:
		{
			m.ip = 0x48a5
			m.alu("sub", 16, m.r[0], m.r[3])
			m.cycles += 4
		}
	case 0x48a5:
		{
			m.ip = 0x48a7
			if m.zf {
				m.ip = uint16(18602)
			}
			m.cycles += 8
		}
	case 0x48a7:
		{
			m.ip = 0x48aa
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x48aa:
		{
			m.ip = 0x48ac
			m.ip = uint16(18706)
			m.cycles += 15
		}
	case 0x48ac:
		{
			m.ip = 0x48ad
			m.cycles += 3
		}
	case 0x48ad:
		{
			m.ip = 0x48af
			m.r[0] = m.alu("xor", 16, m.r[0], m.r[0])
			m.cycles += 4
		}
	case 0x48af:
		{
			m.ip = 0x48b1
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x48b1:
		{
			m.ip = 0x48b4
			target := uint16(6344)
			m.push(0x48b4)
			m.ip = target
			m.cycles += 19
		}
	case 0x48b4:
		{
			m.ip = 0x48b6
			m.r[3] = m.r[1]
			m.cycles += 2
		}
	case 0x48b6:
		{
			m.ip = 0x48b8
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x48b8:
		{
			m.ip = 0x48bb
			target := uint16(6344)
			m.push(0x48bb)
			m.ip = target
			m.cycles += 19
		}
	case 0x48bb:
		{
			m.ip = 0x48bf
			m.alu("sub", 16, m.r[3], m.rd16(m.r[11], uint16(0x280)))
			m.cycles += 16
		}
	case 0x48bf:
		{
			m.ip = 0x48c1
			if !m.cf {
				m.ip = uint16(18636)
			}
			m.cycles += 8
		}
	case 0x48c1:
		{
			m.ip = 0x48c3
			m.r[3] = m.unary("neg", 16, m.r[3])
			m.cycles += 3
		}
	case 0x48c3:
		{
			m.ip = 0x48c7
			m.r[3] = m.alu("add", 16, m.r[3], m.rd16(m.r[11], uint16(0x280)))
			m.cycles += 16
		}
	case 0x48c7:
		{
			m.ip = 0x48c9
			m.set8(2, 8, uint16(2))
			m.cycles += 2
		}
	case 0x48c9:
		{
			m.ip = 0x48cb
			m.ip = uint16(18642)
			m.cycles += 15
		}
	case 0x48cb:
		{
			m.ip = 0x48cc
			m.cycles += 3
		}
	case 0x48cc:
		{
			m.ip = 0x48d0
			m.r[3] = m.alu("sub", 16, m.r[3], m.rd16(m.r[11], uint16(0x280)))
			m.cycles += 16
		}
	case 0x48d0:
		{
			m.ip = 0x48d2
			m.set8(2, 8, uint16(0))
			m.cycles += 2
		}
	case 0x48d2:
		{
			m.ip = 0x48d5
			m.alu("sub", 16, m.r[3], uint16(100))
			m.cycles += 4
		}
	case 0x48d5:
		{
			m.ip = 0x48d7
			if !m.cf {
				m.ip = uint16(18650)
			}
			m.cycles += 8
		}
	case 0x48d7:
		{
			m.ip = 0x48da
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x48da:
		{
			m.ip = 0x48de
			m.alu("sub", 16, m.r[1], m.rd16(m.r[11], uint16(0x282)))
			m.cycles += 16
		}
	case 0x48de:
		{
			m.ip = 0x48e0
			if !m.cf {
				m.ip = uint16(18667)
			}
			m.cycles += 8
		}
	case 0x48e0:
		{
			m.ip = 0x48e2
			m.r[1] = m.unary("neg", 16, m.r[1])
			m.cycles += 3
		}
	case 0x48e2:
		{
			m.ip = 0x48e6
			m.r[1] = m.alu("add", 16, m.r[1], m.rd16(m.r[11], uint16(0x282)))
			m.cycles += 16
		}
	case 0x48e6:
		{
			m.ip = 0x48e8
			m.set8(2, 0, uint16(0))
			m.cycles += 2
		}
	case 0x48e8:
		{
			m.ip = 0x48ea
			m.ip = uint16(18673)
			m.cycles += 15
		}
	case 0x48ea:
		{
			m.ip = 0x48eb
			m.cycles += 3
		}
	case 0x48eb:
		{
			m.ip = 0x48ef
			m.r[1] = m.alu("sub", 16, m.r[1], m.rd16(m.r[11], uint16(0x282)))
			m.cycles += 16
		}
	case 0x48ef:
		{
			m.ip = 0x48f1
			m.set8(2, 0, uint16(8))
			m.cycles += 2
		}
	case 0x48f1:
		{
			m.ip = 0x48f4
			m.alu("sub", 16, m.r[1], uint16(100))
			m.cycles += 4
		}
	case 0x48f4:
		{
			m.ip = 0x48f6
			if !m.cf {
				m.ip = uint16(18686)
			}
			m.cycles += 8
		}
	case 0x48f6:
		{
			m.ip = 0x48f9
			m.alu("sub", 16, m.r[3], uint16(0))
			m.cycles += 4
		}
	case 0x48f9:
		{
			m.ip = 0x48fb
			if m.zf {
				m.ip = uint16(18702)
			}
			m.cycles += 8
		}
	case 0x48fb:
		{
			m.ip = 0x48fe
			m.r[1] = uint16(0)
			m.cycles += 2
		}
	case 0x48fe:
		{
			m.ip = 0x4900
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x4900:
		{
			m.ip = 0x4902
			m.alu("sub", 16, m.r[3], m.r[1])
			m.cycles += 4
		}
	case 0x4902:
		{
			m.ip = 0x4904
			if m.cf {
				m.ip = uint16(18697)
			}
			m.cycles += 8
		}
	case 0x4904:
		{
			m.ip = 0x4906
			m.set8(0, 8, uint16(77))
			m.cycles += 2
		}
	case 0x4906:
		{
			m.ip = 0x4908
			m.set8(0, 8, m.alu("sub", 8, ((m.r[0]>>8)&255), ((m.r[2]>>8)&255)))
			m.cycles += 4
		}
	case 0x4908:
		{
			m.ip = 0x4909
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4909:
		{
			m.ip = 0x490b
			m.set8(0, 8, uint16(72))
			m.cycles += 2
		}
	case 0x490b:
		{
			m.ip = 0x490d
			m.set8(0, 8, m.alu("add", 8, ((m.r[0]>>8)&255), ((m.r[2]>>0)&255)))
			m.cycles += 4
		}
	case 0x490d:
		{
			m.ip = 0x490e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x490e:
		{
			m.ip = 0x4911
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x4911:
		{
			m.ip = 0x4912
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4912:
		{
			m.ip = 0x4913
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x4913:
		{
			m.ip = 0x4914
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4914:
		{
			m.ip = 0x4915
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4915:
		{
			m.ip = 0x4916
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4916:
		{
			m.ip = 0x4918
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4918:
		{
			m.ip = 0x491a
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x491a:
		{
			m.ip = 0x491c
			if m.zf {
				m.ip = uint16(18729)
			}
			m.cycles += 8
		}
	case 0x491c:
		{
			m.ip = 0x491e
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x491e:
		{
			m.ip = 0x4920
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4920:
		{
			m.ip = 0x4921
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4921:
		{
			m.ip = 0x4923
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4923:
		{
			m.ip = 0x4925
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4925:
		{
			m.ip = 0x4926
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x4926:
		{
			m.ip = 0x4928
			if !m.zf {
				m.ip = uint16(18716)
			}
			m.cycles += 8
		}
	case 0x4928:
		{
			m.ip = 0x4929
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4929:
		{
			m.ip = 0x492e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x27d)), uint16(0))
			m.cycles += 16
		}
	case 0x492e:
		{
			m.ip = 0x4930
			if !m.zf {
				m.ip = uint16(18740)
			}
			m.cycles += 8
		}
	case 0x4930:
		{
			m.ip = 0x4933
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x4933:
		{
			m.ip = 0x4934
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4934:
		{
			m.ip = 0x4935
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4935:
		{
			m.ip = 0x4936
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4936:
		{
			m.ip = 0x4937
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x4937:
		{
			m.ip = 0x493c
			m.wr8(m.r[11], uint16(0x27f), uint16(1))
			m.cycles += 8
		}
	case 0x493c:
		{
			m.ip = 0x493f
			m.r[2] = uint16(513)
			m.cycles += 2
		}
	case 0x493f:
		{
			m.ip = 0x4940
			m.set8(0, 0, m.input(m.r[2]))
			m.cycles += 8
		}
	case 0x4940:
		{
			m.ip = 0x4942
			m.alu("and", 8, ((m.r[0] >> 0) & 255), uint16(16))
			m.cycles += 4
		}
	case 0x4942:
		{
			m.ip = 0x4944
			if !m.zf {
				m.ip = uint16(18761)
			}
			m.cycles += 8
		}
	case 0x4944:
		{
			m.ip = 0x4946
			m.set8(0, 0, uint16(13))
			m.cycles += 2
		}
	case 0x4946:
		{
			m.ip = 0x4949
			m.ip = uint16(18889)
			m.cycles += 15
		}
	case 0x4949:
		{
			m.ip = 0x494b
			m.alu("and", 8, ((m.r[0] >> 0) & 255), uint16(32))
			m.cycles += 4
		}
	case 0x494b:
		{
			m.ip = 0x494d
			if !m.zf {
				m.ip = uint16(18770)
			}
			m.cycles += 8
		}
	case 0x494d:
		{
			m.ip = 0x494f
			m.set8(0, 0, uint16(27))
			m.cycles += 2
		}
	case 0x494f:
		{
			m.ip = 0x4951
			m.ip = uint16(18889)
			m.cycles += 15
		}
	case 0x4951:
		{
			m.ip = 0x4952
			m.cycles += 3
		}
	case 0x4952:
		{
			m.ip = 0x4955
			target := uint16(18788)
			m.push(0x4955)
			m.ip = target
			m.cycles += 19
		}
	case 0x4955:
		{
			m.ip = 0x4956
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4956:
		{
			m.ip = 0x4959
			target := uint16(18788)
			m.push(0x4959)
			m.ip = target
			m.cycles += 19
		}
	case 0x4959:
		{
			m.ip = 0x495a
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x495a:
		{
			m.ip = 0x495c
			m.alu("sub", 16, m.r[0], m.r[3])
			m.cycles += 4
		}
	case 0x495c:
		{
			m.ip = 0x495e
			if m.zf {
				m.ip = uint16(18785)
			}
			m.cycles += 8
		}
	case 0x495e:
		{
			m.ip = 0x4961
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x4961:
		{
			m.ip = 0x4963
			m.ip = uint16(18889)
			m.cycles += 15
		}
	case 0x4963:
		{
			m.ip = 0x4964
			m.cycles += 3
		}
	case 0x4964:
		{
			m.ip = 0x4966
			m.r[0] = m.alu("xor", 16, m.r[0], m.r[0])
			m.cycles += 4
		}
	case 0x4966:
		{
			m.ip = 0x4968
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x4968:
		{
			m.ip = 0x496b
			target := uint16(6344)
			m.push(0x496b)
			m.ip = target
			m.cycles += 19
		}
	case 0x496b:
		{
			m.ip = 0x496d
			m.r[3] = m.r[1]
			m.cycles += 2
		}
	case 0x496d:
		{
			m.ip = 0x496f
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x496f:
		{
			m.ip = 0x4972
			target := uint16(6344)
			m.push(0x4972)
			m.ip = target
			m.cycles += 19
		}
	case 0x4972:
		{
			m.ip = 0x4976
			m.alu("sub", 16, m.r[3], m.rd16(m.r[11], uint16(0x280)))
			m.cycles += 16
		}
	case 0x4976:
		{
			m.ip = 0x4978
			if !m.cf {
				m.ip = uint16(18819)
			}
			m.cycles += 8
		}
	case 0x4978:
		{
			m.ip = 0x497a
			m.r[3] = m.unary("neg", 16, m.r[3])
			m.cycles += 3
		}
	case 0x497a:
		{
			m.ip = 0x497e
			m.r[3] = m.alu("add", 16, m.r[3], m.rd16(m.r[11], uint16(0x280)))
			m.cycles += 16
		}
	case 0x497e:
		{
			m.ip = 0x4980
			m.set8(2, 8, uint16(2))
			m.cycles += 2
		}
	case 0x4980:
		{
			m.ip = 0x4982
			m.ip = uint16(18825)
			m.cycles += 15
		}
	case 0x4982:
		{
			m.ip = 0x4983
			m.cycles += 3
		}
	case 0x4983:
		{
			m.ip = 0x4987
			m.r[3] = m.alu("sub", 16, m.r[3], m.rd16(m.r[11], uint16(0x280)))
			m.cycles += 16
		}
	case 0x4987:
		{
			m.ip = 0x4989
			m.set8(2, 8, uint16(0))
			m.cycles += 2
		}
	case 0x4989:
		{
			m.ip = 0x498c
			m.alu("sub", 16, m.r[3], uint16(100))
			m.cycles += 4
		}
	case 0x498c:
		{
			m.ip = 0x498e
			if !m.cf {
				m.ip = uint16(18833)
			}
			m.cycles += 8
		}
	case 0x498e:
		{
			m.ip = 0x4991
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x4991:
		{
			m.ip = 0x4995
			m.alu("sub", 16, m.r[1], m.rd16(m.r[11], uint16(0x282)))
			m.cycles += 16
		}
	case 0x4995:
		{
			m.ip = 0x4997
			if !m.cf {
				m.ip = uint16(18850)
			}
			m.cycles += 8
		}
	case 0x4997:
		{
			m.ip = 0x4999
			m.r[1] = m.unary("neg", 16, m.r[1])
			m.cycles += 3
		}
	case 0x4999:
		{
			m.ip = 0x499d
			m.r[1] = m.alu("add", 16, m.r[1], m.rd16(m.r[11], uint16(0x282)))
			m.cycles += 16
		}
	case 0x499d:
		{
			m.ip = 0x499f
			m.set8(2, 0, uint16(0))
			m.cycles += 2
		}
	case 0x499f:
		{
			m.ip = 0x49a1
			m.ip = uint16(18856)
			m.cycles += 15
		}
	case 0x49a1:
		{
			m.ip = 0x49a2
			m.cycles += 3
		}
	case 0x49a2:
		{
			m.ip = 0x49a6
			m.r[1] = m.alu("sub", 16, m.r[1], m.rd16(m.r[11], uint16(0x282)))
			m.cycles += 16
		}
	case 0x49a6:
		{
			m.ip = 0x49a8
			m.set8(2, 0, uint16(8))
			m.cycles += 2
		}
	case 0x49a8:
		{
			m.ip = 0x49ab
			m.alu("sub", 16, m.r[1], uint16(100))
			m.cycles += 4
		}
	case 0x49ab:
		{
			m.ip = 0x49ad
			if !m.cf {
				m.ip = uint16(18869)
			}
			m.cycles += 8
		}
	case 0x49ad:
		{
			m.ip = 0x49b0
			m.alu("sub", 16, m.r[3], uint16(0))
			m.cycles += 4
		}
	case 0x49b0:
		{
			m.ip = 0x49b2
			if m.zf {
				m.ip = uint16(18885)
			}
			m.cycles += 8
		}
	case 0x49b2:
		{
			m.ip = 0x49b5
			m.r[1] = uint16(0)
			m.cycles += 2
		}
	case 0x49b5:
		{
			m.ip = 0x49b7
			m.set8(0, 0, uint16(0))
			m.cycles += 2
		}
	case 0x49b7:
		{
			m.ip = 0x49b9
			m.alu("sub", 16, m.r[3], m.r[1])
			m.cycles += 4
		}
	case 0x49b9:
		{
			m.ip = 0x49bb
			if m.cf {
				m.ip = uint16(18880)
			}
			m.cycles += 8
		}
	case 0x49bb:
		{
			m.ip = 0x49bd
			m.set8(0, 8, uint16(77))
			m.cycles += 2
		}
	case 0x49bd:
		{
			m.ip = 0x49bf
			m.set8(0, 8, m.alu("sub", 8, ((m.r[0]>>8)&255), ((m.r[2]>>8)&255)))
			m.cycles += 4
		}
	case 0x49bf:
		{
			m.ip = 0x49c0
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x49c0:
		{
			m.ip = 0x49c2
			m.set8(0, 8, uint16(72))
			m.cycles += 2
		}
	case 0x49c2:
		{
			m.ip = 0x49c4
			m.set8(0, 8, m.alu("add", 8, ((m.r[0]>>8)&255), ((m.r[2]>>0)&255)))
			m.cycles += 4
		}
	case 0x49c4:
		{
			m.ip = 0x49c5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x49c5:
		{
			m.ip = 0x49c8
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x49c8:
		{
			m.ip = 0x49c9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x49c9:
		{
			m.ip = 0x49ca
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x49ca:
		{
			m.ip = 0x49cb
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x49cb:
		{
			m.ip = 0x49cc
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x49cc:
		{
			m.ip = 0x49cd
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x49cd:
		{
			m.ip = 0x49d0
			target := uint16(18515)
			m.push(0x49d0)
			m.ip = target
			m.cycles += 19
		}
	case 0x49d0:
		{
			m.ip = 0x49d3
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x49d3:
		{
			m.ip = 0x49d5
			if !m.zf {
				m.ip = uint16(18906)
			}
			m.cycles += 8
		}
	case 0x49d5:
		{
			m.ip = 0x49d8
			m.wr16(m.r[11], uint16(0x288), m.r[0])
			m.cycles += 8
		}
	case 0x49d8:
		{
			m.ip = 0x49da
			m.ip = uint16(18893)
			m.cycles += 15
		}
	case 0x49da:
		{
			m.ip = 0x49df
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x27f)), uint16(0))
			m.cycles += 16
		}
	case 0x49df:
		{
			m.ip = 0x49e1
			if m.zf {
				m.ip = uint16(18922)
			}
			m.cycles += 8
		}
	case 0x49e1:
		{
			m.ip = 0x49e5
			m.alu("sub", 16, m.r[0], m.rd16(m.r[11], uint16(0x288)))
			m.cycles += 16
		}
	case 0x49e5:
		{
			m.ip = 0x49e7
			if m.zf {
				m.ip = uint16(18893)
			}
			m.cycles += 8
		}
	case 0x49e7:
		{
			m.ip = 0x49ea
			m.wr16(m.r[11], uint16(0x288), m.r[0])
			m.cycles += 8
		}
	case 0x49ea:
		{
			m.ip = 0x49eb
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x49eb:
		{
			m.ip = 0x49ed
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x49ed:
		{
			m.ip = 0x49ef
			m.interrupt(uint16(26))
			m.cycles += 50
		}
	case 0x49ef:
		{
			m.ip = 0x49f2
			m.r[2] = m.alu("or", 16, m.r[2], uint16(1))
			m.cycles += 4
		}
	case 0x49f2:
		{
			m.ip = 0x49f6
			m.wr16(m.r[11], uint16(0x1067), m.r[2])
			m.cycles += 8
		}
	case 0x49f6:
		{
			m.ip = 0x49fb
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(1))
			m.cycles += 16
		}
	case 0x49fb:
		{
			m.ip = 0x49fd
			if !m.zf {
				m.ip = uint16(18947)
			}
			m.cycles += 8
		}
	case 0x49fd:
		{
			m.ip = 0x4a03
			m.wr16(m.r[11], uint16(0x1067), uint16(12345))
			m.cycles += 8
		}
	case 0x4a03:
		{
			m.ip = 0x4a04
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4a04:
		{
			m.ip = 0x4a07
			target := uint16(19029)
			m.push(0x4a07)
			m.ip = target
			m.cycles += 19
		}
	case 0x4a07:
		{
			m.ip = 0x4a0a
			m.set8(0, 8, m.alu("and", 8, ((m.r[0]>>8)&255), uint16(3)))
			m.cycles += 4
		}
	case 0x4a0a:
		{
			m.ip = 0x4a0d
			m.alu("sub", 16, m.r[0], uint16(639))
			m.cycles += 4
		}
	case 0x4a0d:
		{
			m.ip = 0x4a0f
			if !m.zf && m.sf == m.of {
				m.ip = uint16(18960)
			}
			m.cycles += 8
		}
	case 0x4a0f:
		{
			m.ip = 0x4a10
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4a10:
		{
			m.ip = 0x4a13
			m.set8(0, 8, m.alu("and", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 4
		}
	case 0x4a13:
		{
			m.ip = 0x4a14
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4a14:
		{
			m.ip = 0x4a17
			target := uint16(19029)
			m.push(0x4a17)
			m.ip = target
			m.cycles += 19
		}
	case 0x4a17:
		{
			m.ip = 0x4a1a
			m.set8(0, 8, m.alu("and", 8, ((m.r[0]>>8)&255), uint16(1)))
			m.cycles += 4
		}
	case 0x4a1a:
		{
			m.ip = 0x4a1d
			m.alu("sub", 16, m.r[0], uint16(349))
			m.cycles += 4
		}
	case 0x4a1d:
		{
			m.ip = 0x4a1f
			if !m.zf && m.sf == m.of {
				m.ip = uint16(18976)
			}
			m.cycles += 8
		}
	case 0x4a1f:
		{
			m.ip = 0x4a20
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4a20:
		{
			m.ip = 0x4a23
			m.set8(0, 8, m.alu("and", 8, ((m.r[0]>>8)&255), uint16(0)))
			m.cycles += 4
		}
	case 0x4a23:
		{
			m.ip = 0x4a24
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4a24:
		{
			m.ip = 0x4a25
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4a25:
		{
			m.ip = 0x4a26
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x4a26:
		{
			m.ip = 0x4a2a
			m.r[3] = m.rd16(m.r[11], uint16(0x1067))
			m.cycles += 8
		}
	case 0x4a2a:
		{
			m.ip = 0x4a2c
			m.r[2] = m.r[3]
			m.cycles += 2
		}
	case 0x4a2c:
		{
			m.ip = 0x4a2e
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a2e:
		{
			m.ip = 0x4a30
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a30:
		{
			m.ip = 0x4a32
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x4a32:
		{
			m.ip = 0x4a34
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a34:
		{
			m.ip = 0x4a36
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a36:
		{
			m.ip = 0x4a38
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a38:
		{
			m.ip = 0x4a3a
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x4a3a:
		{
			m.ip = 0x4a3c
			m.set8(0, 0, ((m.r[3] >> 8) & 255))
			m.cycles += 2
		}
	case 0x4a3c:
		{
			m.ip = 0x4a3e
			m.r[2] = m.r[3]
			m.cycles += 2
		}
	case 0x4a3e:
		{
			m.ip = 0x4a40
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a40:
		{
			m.ip = 0x4a42
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a42:
		{
			m.ip = 0x4a44
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x4a44:
		{
			m.ip = 0x4a46
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a46:
		{
			m.ip = 0x4a48
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a48:
		{
			m.ip = 0x4a4a
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a4a:
		{
			m.ip = 0x4a4c
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x4a4c:
		{
			m.ip = 0x4a50
			m.wr16(m.r[11], uint16(0x1067), m.r[3])
			m.cycles += 8
		}
	case 0x4a50:
		{
			m.ip = 0x4a52
			m.set8(0, 8, ((m.r[3] >> 8) & 255))
			m.cycles += 2
		}
	case 0x4a52:
		{
			m.ip = 0x4a53
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x4a53:
		{
			m.ip = 0x4a54
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4a54:
		{
			m.ip = 0x4a55
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4a55:
		{
			m.ip = 0x4a59
			m.set8(0, 8, m.rd8(m.r[11], uint16(0x1069)))
			m.cycles += 8
		}
	case 0x4a59:
		{
			m.ip = 0x4a5c
			m.set8(0, 8, m.alu("add", 8, ((m.r[0]>>8)&255), uint16(2)))
			m.cycles += 4
		}
	case 0x4a5c:
		{
			m.ip = 0x4a5f
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(100))
			m.cycles += 4
		}
	case 0x4a5f:
		{
			m.ip = 0x4a61
			if m.sf != m.of {
				m.ip = uint16(19046)
			}
			m.cycles += 8
		}
	case 0x4a61:
		{
			m.ip = 0x4a63
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x4a63:
		{
			m.ip = 0x4a66
			target := uint16(18923)
			m.push(0x4a66)
			m.ip = target
			m.cycles += 19
		}
	case 0x4a66:
		{
			m.ip = 0x4a6a
			m.wr8(m.r[11], uint16(0x1069), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x4a6a:
		{
			m.ip = 0x4a6e
			m.r[3] = m.rd16(m.r[11], uint16(0x1067))
			m.cycles += 8
		}
	case 0x4a6e:
		{
			m.ip = 0x4a70
			m.r[2] = m.r[3]
			m.cycles += 2
		}
	case 0x4a70:
		{
			m.ip = 0x4a72
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a72:
		{
			m.ip = 0x4a74
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a74:
		{
			m.ip = 0x4a76
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x4a76:
		{
			m.ip = 0x4a78
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a78:
		{
			m.ip = 0x4a7a
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a7a:
		{
			m.ip = 0x4a7c
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a7c:
		{
			m.ip = 0x4a7e
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x4a7e:
		{
			m.ip = 0x4a80
			m.set8(0, 0, ((m.r[3] >> 8) & 255))
			m.cycles += 2
		}
	case 0x4a80:
		{
			m.ip = 0x4a82
			m.r[2] = m.r[3]
			m.cycles += 2
		}
	case 0x4a82:
		{
			m.ip = 0x4a84
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a84:
		{
			m.ip = 0x4a86
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a86:
		{
			m.ip = 0x4a88
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x4a88:
		{
			m.ip = 0x4a8a
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a8a:
		{
			m.ip = 0x4a8c
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a8c:
		{
			m.ip = 0x4a8e
			m.r[3] = m.alu("add", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4a8e:
		{
			m.ip = 0x4a90
			m.r[3] = m.alu("add", 16, m.r[3], m.r[2])
			m.cycles += 4
		}
	case 0x4a90:
		{
			m.ip = 0x4a94
			m.wr16(m.r[11], uint16(0x1067), m.r[3])
			m.cycles += 8
		}
	case 0x4a94:
		{
			m.ip = 0x4a96
			m.set8(0, 8, ((m.r[3] >> 8) & 255))
			m.cycles += 2
		}
	case 0x4a96:
		{
			m.ip = 0x4a97
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4a97:
		{
			m.ip = 0x4a98
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4a98:
		{
			m.ip = 0x4a99
			m.push(m.r[2])
			m.cycles += 11
		}
	case 0x4a99:
		{
			m.ip = 0x4a9d
			m.set8(0, 8, m.rd8(m.r[11], uint16(0x1069)))
			m.cycles += 8
		}
	case 0x4a9d:
		{
			m.ip = 0x4a9f
			m.set8(0, 8, m.unary("inc", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x4a9f:
		{
			m.ip = 0x4aa2
			m.alu("sub", 8, ((m.r[0] >> 8) & 255), uint16(100))
			m.cycles += 4
		}
	case 0x4aa2:
		{
			m.ip = 0x4aa4
			if m.sf != m.of {
				m.ip = uint16(19113)
			}
			m.cycles += 8
		}
	case 0x4aa4:
		{
			m.ip = 0x4aa6
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x4aa6:
		{
			m.ip = 0x4aa9
			target := uint16(18923)
			m.push(0x4aa9)
			m.ip = target
			m.cycles += 19
		}
	case 0x4aa9:
		{
			m.ip = 0x4aad
			m.wr8(m.r[11], uint16(0x1069), ((m.r[0] >> 8) & 255))
			m.cycles += 8
		}
	case 0x4aad:
		{
			m.ip = 0x4ab0
			m.r[0] = m.rd16(m.r[11], uint16(0x1067))
			m.cycles += 8
		}
	case 0x4ab0:
		{
			m.ip = 0x4ab2
			m.r[3] = m.r[0]
			m.cycles += 2
		}
	case 0x4ab2:
		{
			m.ip = 0x4ab4
			m.r[0] = m.alu("add", 16, m.r[0], m.r[0])
			m.cycles += 4
		}
	case 0x4ab4:
		{
			m.ip = 0x4ab6
			m.r[0] = m.alu("add", 16, m.r[0], m.r[0])
			m.cycles += 4
		}
	case 0x4ab6:
		{
			m.ip = 0x4ab8
			m.r[0] = m.alu("add", 16, m.r[0], m.r[3])
			m.cycles += 4
		}
	case 0x4ab8:
		{
			m.ip = 0x4aba
			m.r[0] = m.alu("add", 16, m.r[0], m.r[0])
			m.cycles += 4
		}
	case 0x4aba:
		{
			m.ip = 0x4abc
			m.r[0] = m.alu("add", 16, m.r[0], m.r[0])
			m.cycles += 4
		}
	case 0x4abc:
		{
			m.ip = 0x4abe
			m.r[0] = m.alu("add", 16, m.r[0], m.r[0])
			m.cycles += 4
		}
	case 0x4abe:
		{
			m.ip = 0x4ac0
			m.r[0] = m.alu("add", 16, m.r[0], m.r[3])
			m.cycles += 4
		}
	case 0x4ac0:
		{
			m.ip = 0x4ac3
			m.wr16(m.r[11], uint16(0x1067), m.r[0])
			m.cycles += 8
		}
	case 0x4ac3:
		{
			m.ip = 0x4ac4
			m.r[2] = m.pop()
			m.cycles += 8
		}
	case 0x4ac4:
		{
			m.ip = 0x4ac5
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4ac5:
		{
			m.ip = 0x4ac6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4ac6:
		{
			m.ip = 0x4acb
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(7))
			m.cycles += 16
		}
	case 0x4acb:
		{
			m.ip = 0x4acd
			if !m.zf {
				m.ip = uint16(19157)
			}
			m.cycles += 8
		}
	case 0x4acd:
		{
			m.ip = 0x4ad2
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f6)), uint16(14))
			m.cycles += 16
		}
	case 0x4ad2:
		{
			m.ip = 0x4ad4
			if !m.zf {
				m.ip = uint16(19157)
			}
			m.cycles += 8
		}
	case 0x4ad4:
		{
			m.ip = 0x4ad5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4ad5:
		{
			m.ip = 0x4ad6
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4ad6:
		{
			m.ip = 0x4ad7
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4ad7:
		{
			m.ip = 0x4ad8
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4ad8:
		{
			m.ip = 0x4adc
			m.r[3] = m.rd16(m.r[11], uint16(0x106c))
			m.cycles += 8
		}
	case 0x4adc:
		{
			m.ip = 0x4ae0
			m.r[1] = m.rd16(m.r[11], uint16(0x106a))
			m.cycles += 8
		}
	case 0x4ae0:
		{
			m.ip = 0x4ae3
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4ae3:
		{
			m.ip = 0x4ae5
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19168)
			}
			m.cycles += 17
		}
	case 0x4ae5:
		{
			m.ip = 0x4ae7
			m.r[1] = m.r[3]
			m.cycles += 2
		}
	case 0x4ae7:
		{
			m.ip = 0x4ae8
			m.r[3] = m.unary("dec", 16, m.r[3])
			m.cycles += 3
		}
	case 0x4ae8:
		{
			m.ip = 0x4aea
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19164)
			}
			m.cycles += 17
		}
	case 0x4aea:
		{
			m.ip = 0x4aeb
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4aeb:
		{
			m.ip = 0x4aec
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4aec:
		{
			m.ip = 0x4aed
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x4aed:
		{
			m.ip = 0x4aee
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4aee:
		{
			m.ip = 0x4af3
			m.wr8(m.r[11], uint16(0x1070), uint16(0))
			m.cycles += 8
		}
	case 0x4af3:
		{
			m.ip = 0x4af8
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(7))
			m.cycles += 16
		}
	case 0x4af8:
		{
			m.ip = 0x4afa
			if !m.zf {
				m.ip = uint16(19221)
			}
			m.cycles += 8
		}
	case 0x4afa:
		{
			m.ip = 0x4aff
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f6)), uint16(14))
			m.cycles += 16
		}
	case 0x4aff:
		{
			m.ip = 0x4b01
			if !m.zf {
				m.ip = uint16(19221)
			}
			m.cycles += 8
		}
	case 0x4b01:
		{
			m.ip = 0x4b03
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4b03:
		{
			m.ip = 0x4b05
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4b05:
		{
			m.ip = 0x4b07
			if m.zf {
				m.ip = uint16(19217)
			}
			m.cycles += 8
		}
	case 0x4b07:
		{
			m.ip = 0x4b0c
			m.wr8(m.r[11], uint16(0x1070), uint16(1))
			m.cycles += 8
		}
	case 0x4b0c:
		{
			m.ip = 0x4b0e
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x4b0e:
		{
			m.ip = 0x4b10
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4b10:
		{
			m.ip = 0x4b11
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4b11:
		{
			m.ip = 0x4b14
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x4b14:
		{
			m.ip = 0x4b15
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4b15:
		{
			m.ip = 0x4b16
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4b16:
		{
			m.ip = 0x4b17
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4b17:
		{
			m.ip = 0x4b1b
			m.r[3] = m.rd16(m.r[11], uint16(0x106c))
			m.cycles += 8
		}
	case 0x4b1b:
		{
			m.ip = 0x4b20
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x27d)), uint16(0))
			m.cycles += 16
		}
	case 0x4b20:
		{
			m.ip = 0x4b22
			if m.zf {
				m.ip = uint16(19245)
			}
			m.cycles += 8
		}
	case 0x4b22:
		{
			m.ip = 0x4b24
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4b24:
		{
			m.ip = 0x4b26
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4b26:
		{
			m.ip = 0x4b28
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4b28:
		{
			m.ip = 0x4b2a
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4b2a:
		{
			m.ip = 0x4b2c
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4b2c:
		{
			m.ip = 0x4b2d
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x4b2d:
		{
			m.ip = 0x4b31
			m.r[1] = m.rd16(m.r[11], uint16(0x106a))
			m.cycles += 8
		}
	case 0x4b31:
		{
			m.ip = 0x4b34
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4b34:
		{
			m.ip = 0x4b36
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19249)
			}
			m.cycles += 17
		}
	case 0x4b36:
		{
			m.ip = 0x4b3b
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(0))
			m.cycles += 16
		}
	case 0x4b3b:
		{
			m.ip = 0x4b3d
			if m.zf {
				m.ip = uint16(19270)
			}
			m.cycles += 8
		}
	case 0x4b3d:
		{
			m.ip = 0x4b3f
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4b3f:
		{
			m.ip = 0x4b41
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4b41:
		{
			m.ip = 0x4b43
			if !m.zf {
				m.ip = uint16(19303)
			}
			m.cycles += 8
		}
	case 0x4b43:
		{
			m.ip = 0x4b45
			m.ip = uint16(19278)
			m.cycles += 15
		}
	case 0x4b45:
		{
			m.ip = 0x4b46
			m.cycles += 3
		}
	case 0x4b46:
		{
			m.ip = 0x4b49
			target := uint16(18515)
			m.push(0x4b49)
			m.ip = target
			m.cycles += 19
		}
	case 0x4b49:
		{
			m.ip = 0x4b4c
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x4b4c:
		{
			m.ip = 0x4b4e
			if !m.zf {
				m.ip = uint16(19286)
			}
			m.cycles += 8
		}
	case 0x4b4e:
		{
			m.ip = 0x4b50
			m.r[1] = m.r[3]
			m.cycles += 2
		}
	case 0x4b50:
		{
			m.ip = 0x4b51
			m.r[3] = m.unary("dec", 16, m.r[3])
			m.cycles += 3
		}
	case 0x4b51:
		{
			m.ip = 0x4b53
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19245)
			}
			m.cycles += 17
		}
	case 0x4b53:
		{
			m.ip = 0x4b55
			m.ip = uint16(19308)
			m.cycles += 15
		}
	case 0x4b55:
		{
			m.ip = 0x4b56
			m.cycles += 3
		}
	case 0x4b56:
		{
			m.ip = 0x4b58
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x4b58:
		{
			m.ip = 0x4b5a
			if !m.zf {
				m.ip = uint16(19278)
			}
			m.cycles += 8
		}
	case 0x4b5a:
		{
			m.ip = 0x4b5f
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2074)), uint16(0))
			m.cycles += 16
		}
	case 0x4b5f:
		{
			m.ip = 0x4b61
			if !m.zf {
				m.ip = uint16(19303)
			}
			m.cycles += 8
		}
	case 0x4b61:
		{
			m.ip = 0x4b64
			target := uint16(3091)
			m.push(0x4b64)
			m.ip = target
			m.cycles += 19
		}
	case 0x4b64:
		{
			m.ip = 0x4b66
			m.ip = uint16(19308)
			m.cycles += 15
		}
	case 0x4b66:
		{
			m.ip = 0x4b67
			m.cycles += 3
		}
	case 0x4b67:
		{
			m.ip = 0x4b6c
			m.wr8(m.r[11], uint16(0x1070), uint16(1))
			m.cycles += 8
		}
	case 0x4b6c:
		{
			m.ip = 0x4b6d
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4b6d:
		{
			m.ip = 0x4b6e
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4b6e:
		{
			m.ip = 0x4b6f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4b6f:
		{
			m.ip = 0x4b74
			m.wr8(m.r[11], uint16(0x1070), uint16(0))
			m.cycles += 8
		}
	case 0x4b74:
		{
			m.ip = 0x4b79
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x24f2)), uint16(7))
			m.cycles += 16
		}
	case 0x4b79:
		{
			m.ip = 0x4b7b
			if !m.zf {
				m.ip = uint16(19350)
			}
			m.cycles += 8
		}
	case 0x4b7b:
		{
			m.ip = 0x4b80
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x24f6)), uint16(14))
			m.cycles += 16
		}
	case 0x4b80:
		{
			m.ip = 0x4b82
			if !m.zf {
				m.ip = uint16(19350)
			}
			m.cycles += 8
		}
	case 0x4b82:
		{
			m.ip = 0x4b84
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4b84:
		{
			m.ip = 0x4b86
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4b86:
		{
			m.ip = 0x4b88
			if m.zf {
				m.ip = uint16(19346)
			}
			m.cycles += 8
		}
	case 0x4b88:
		{
			m.ip = 0x4b8d
			m.wr8(m.r[11], uint16(0x1070), uint16(1))
			m.cycles += 8
		}
	case 0x4b8d:
		{
			m.ip = 0x4b8f
			m.set8(0, 8, uint16(0))
			m.cycles += 2
		}
	case 0x4b8f:
		{
			m.ip = 0x4b91
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4b91:
		{
			m.ip = 0x4b92
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4b92:
		{
			m.ip = 0x4b95
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x4b95:
		{
			m.ip = 0x4b96
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4b96:
		{
			m.ip = 0x4b97
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4b97:
		{
			m.ip = 0x4b98
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4b98:
		{
			m.ip = 0x4b9c
			m.r[3] = m.rd16(m.r[11], uint16(0x106c))
			m.cycles += 8
		}
	case 0x4b9c:
		{
			m.ip = 0x4ba1
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x27d)), uint16(0))
			m.cycles += 16
		}
	case 0x4ba1:
		{
			m.ip = 0x4ba3
			if m.zf {
				m.ip = uint16(19374)
			}
			m.cycles += 8
		}
	case 0x4ba3:
		{
			m.ip = 0x4ba5
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4ba5:
		{
			m.ip = 0x4ba7
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4ba7:
		{
			m.ip = 0x4ba9
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4ba9:
		{
			m.ip = 0x4bab
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4bab:
		{
			m.ip = 0x4bad
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4bad:
		{
			m.ip = 0x4bae
			m.r[3] = m.unary("inc", 16, m.r[3])
			m.cycles += 3
		}
	case 0x4bae:
		{
			m.ip = 0x4bb2
			m.r[1] = m.rd16(m.r[11], uint16(0x106a))
			m.cycles += 8
		}
	case 0x4bb2:
		{
			m.ip = 0x4bb5
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4bb5:
		{
			m.ip = 0x4bb7
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19378)
			}
			m.cycles += 17
		}
	case 0x4bb7:
		{
			m.ip = 0x4bbc
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1396)), uint16(0))
			m.cycles += 16
		}
	case 0x4bbc:
		{
			m.ip = 0x4bbe
			if m.zf {
				m.ip = uint16(19399)
			}
			m.cycles += 8
		}
	case 0x4bbe:
		{
			m.ip = 0x4bc0
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4bc0:
		{
			m.ip = 0x4bc2
			m.interrupt(uint16(22))
			m.cycles += 50
		}
	case 0x4bc2:
		{
			m.ip = 0x4bc4
			if !m.zf {
				m.ip = uint16(19432)
			}
			m.cycles += 8
		}
	case 0x4bc4:
		{
			m.ip = 0x4bc6
			m.ip = uint16(19407)
			m.cycles += 15
		}
	case 0x4bc6:
		{
			m.ip = 0x4bc7
			m.cycles += 3
		}
	case 0x4bc7:
		{
			m.ip = 0x4bca
			target := uint16(18501)
			m.push(0x4bca)
			m.ip = target
			m.cycles += 19
		}
	case 0x4bca:
		{
			m.ip = 0x4bcd
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x4bcd:
		{
			m.ip = 0x4bcf
			if !m.zf {
				m.ip = uint16(19415)
			}
			m.cycles += 8
		}
	case 0x4bcf:
		{
			m.ip = 0x4bd1
			m.r[1] = m.r[3]
			m.cycles += 2
		}
	case 0x4bd1:
		{
			m.ip = 0x4bd2
			m.r[3] = m.unary("dec", 16, m.r[3])
			m.cycles += 3
		}
	case 0x4bd2:
		{
			m.ip = 0x4bd4
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19374)
			}
			m.cycles += 17
		}
	case 0x4bd4:
		{
			m.ip = 0x4bd6
			m.ip = uint16(19437)
			m.cycles += 15
		}
	case 0x4bd6:
		{
			m.ip = 0x4bd7
			m.cycles += 3
		}
	case 0x4bd7:
		{
			m.ip = 0x4bd9
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), uint16(27))
			m.cycles += 4
		}
	case 0x4bd9:
		{
			m.ip = 0x4bdb
			if !m.zf {
				m.ip = uint16(19432)
			}
			m.cycles += 8
		}
	case 0x4bdb:
		{
			m.ip = 0x4be0
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x2074)), uint16(0))
			m.cycles += 16
		}
	case 0x4be0:
		{
			m.ip = 0x4be2
			if !m.zf {
				m.ip = uint16(19432)
			}
			m.cycles += 8
		}
	case 0x4be2:
		{
			m.ip = 0x4be5
			target := uint16(3091)
			m.push(0x4be5)
			m.ip = target
			m.cycles += 19
		}
	case 0x4be5:
		{
			m.ip = 0x4be7
			m.ip = uint16(19437)
			m.cycles += 15
		}
	case 0x4be7:
		{
			m.ip = 0x4be8
			m.cycles += 3
		}
	case 0x4be8:
		{
			m.ip = 0x4bed
			m.wr8(m.r[11], uint16(0x1070), uint16(1))
			m.cycles += 8
		}
	case 0x4bed:
		{
			m.ip = 0x4bee
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4bee:
		{
			m.ip = 0x4bef
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4bef:
		{
			m.ip = 0x4bf0
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4bf0:
		{
			m.ip = 0x4bf3
			target := uint16(18515)
			m.push(0x4bf3)
			m.ip = target
			m.cycles += 19
		}
	case 0x4bf3:
		{
			m.ip = 0x4bf6
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x4bf6:
		{
			m.ip = 0x4bf8
			if m.zf {
				m.ip = uint16(19449)
			}
			m.cycles += 8
		}
	case 0x4bf8:
		{
			m.ip = 0x4bf9
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4bf9:
		{
			m.ip = 0x4bfc
			m.r[0] = m.rd16(m.r[11], uint16(0x3457))
			m.cycles += 8
		}
	case 0x4bfc:
		{
			m.ip = 0x4bff
			m.wr16(m.r[11], uint16(0x106a), m.r[0])
			m.cycles += 8
		}
	case 0x4bff:
		{
			m.ip = 0x4c02
			m.r[0] = m.rd16(m.r[11], uint16(0x497))
			m.cycles += 8
		}
	case 0x4c02:
		{
			m.ip = 0x4c06
			m.r[3] = m.rd16(m.r[11], uint16(0x106c))
			m.cycles += 8
		}
	case 0x4c06:
		{
			m.ip = 0x4c0a
			m.wr16(m.r[11], uint16(0x497), m.r[3])
			m.cycles += 8
		}
	case 0x4c0a:
		{
			m.ip = 0x4c0d
			m.wr16(m.r[11], uint16(0x106c), m.r[0])
			m.cycles += 8
		}
	case 0x4c0d:
		{
			m.ip = 0x4c11
			m.r[2] = uint16(0x299)
			m.cycles += 3
		}
	case 0x4c11:
		{
			m.ip = 0x4c14
			target := uint16(17942)
			m.push(0x4c14)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c14:
		{
			m.ip = 0x4c17
			m.r[0] = m.rd16(m.r[11], uint16(0x274))
			m.cycles += 8
		}
	case 0x4c17:
		{
			m.ip = 0x4c1a
			target := uint16(18137)
			m.push(0x4c1a)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c1a:
		{
			m.ip = 0x4c1e
			m.r[6] = uint16(0x1056)
			m.cycles += 3
		}
	case 0x4c1e:
		{
			m.ip = 0x4c21
			target := uint16(18306)
			m.push(0x4c21)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c21:
		{
			m.ip = 0x4c24
			target := uint16(16773)
			m.push(0x4c24)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c24:
		{
			m.ip = 0x4c27
			target := uint16(19632)
			m.push(0x4c27)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c27:
		{
			m.ip = 0x4c2b
			m.r[6] = uint16(0x1045)
			m.cycles += 3
		}
	case 0x4c2b:
		{
			m.ip = 0x4c2e
			target := uint16(18306)
			m.push(0x4c2e)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c2e:
		{
			m.ip = 0x4c32
			m.r[8] = m.rd16(m.r[11], uint16(0x272))
			m.cycles += 8
		}
	case 0x4c32:
		{
			m.ip = 0x4c38
			m.wr16(m.r[11], uint16(0x345b), uint16(28000))
			m.cycles += 8
		}
	case 0x4c38:
		{
			m.ip = 0x4c3e
			m.wr16(m.r[11], uint16(0x345d), uint16(37440))
			m.cycles += 8
		}
	case 0x4c3e:
		{
			m.ip = 0x4c44
			m.wr16(m.r[11], uint16(0x345f), uint16(11613))
			m.cycles += 8
		}
	case 0x4c44:
		{
			m.ip = 0x4c4a
			m.wr16(m.r[11], uint16(0x3461), uint16(11657))
			m.cycles += 8
		}
	case 0x4c4a:
		{
			m.ip = 0x4c50
			m.wr16(m.r[11], uint16(0x3467), uint16(145))
			m.cycles += 8
		}
	case 0x4c50:
		{
			m.ip = 0x4c56
			m.wr16(m.r[11], uint16(0x3469), uint16(65535))
			m.cycles += 8
		}
	case 0x4c56:
		{
			m.ip = 0x4c5c
			m.wr16(m.r[11], uint16(0x3463), uint16(80))
			m.cycles += 8
		}
	case 0x4c5c:
		{
			m.ip = 0x4c62
			m.wr16(m.r[11], uint16(0x3465), uint16(65456))
			m.cycles += 8
		}
	case 0x4c62:
		{
			m.ip = 0x4c65
			target := uint16(20389)
			m.push(0x4c65)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c65:
		{
			m.ip = 0x4c68
			m.r[6] = uint16(46880)
			m.cycles += 2
		}
	case 0x4c68:
		{
			m.ip = 0x4c6b
			m.r[7] = uint16(10910)
			m.cycles += 2
		}
	case 0x4c6b:
		{
			m.ip = 0x4c6e
			target := uint16(17474)
			m.push(0x4c6e)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c6e:
		{
			m.ip = 0x4c71
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x4c71:
		{
			m.ip = 0x4c73
			if !m.zf {
				m.ip = uint16(19603)
			}
			m.cycles += 8
		}
	case 0x4c73:
		{
			m.ip = 0x4c76
			target := uint16(19864)
			m.push(0x4c76)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c76:
		{
			m.ip = 0x4c79
			target := uint16(20389)
			m.push(0x4c79)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c79:
		{
			m.ip = 0x4c7c
			target := uint16(18515)
			m.push(0x4c7c)
			m.ip = target
			m.cycles += 19
		}
	case 0x4c7c:
		{
			m.ip = 0x4c7f
			m.alu("sub", 16, m.r[0], uint16(65535))
			m.cycles += 4
		}
	case 0x4c7f:
		{
			m.ip = 0x4c81
			if !m.zf {
				m.ip = uint16(19603)
			}
			m.cycles += 8
		}
	case 0x4c81:
		{
			m.ip = 0x4c86
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x1397)), uint16(1))
			m.cycles += 16
		}
	case 0x4c86:
		{
			m.ip = 0x4c88
			if !m.zf {
				m.ip = uint16(19571)
			}
			m.cycles += 8
		}
	case 0x4c88:
		{
			m.ip = 0x4c8e
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x3467)), uint16(145))
			m.cycles += 16
		}
	case 0x4c8e:
		{
			m.ip = 0x4c90
			if !m.zf {
				m.ip = uint16(19571)
			}
			m.cycles += 8
		}
	case 0x4c90:
		{
			m.ip = 0x4c93
			m.r[0] = uint16(65535)
			m.cycles += 2
		}
	case 0x4c93:
		{
			m.ip = 0x4c97
			m.r[3] = m.rd16(m.r[11], uint16(0x106a))
			m.cycles += 8
		}
	case 0x4c97:
		{
			m.ip = 0x4c9b
			m.wr16(m.r[11], uint16(0x3457), m.r[3])
			m.cycles += 8
		}
	case 0x4c9b:
		{
			m.ip = 0x4c9f
			m.r[3] = m.rd16(m.r[11], uint16(0x497))
			m.cycles += 8
		}
	case 0x4c9f:
		{
			m.ip = 0x4ca3
			m.r[1] = m.rd16(m.r[11], uint16(0x106c))
			m.cycles += 8
		}
	case 0x4ca3:
		{
			m.ip = 0x4ca7
			m.wr16(m.r[11], uint16(0x497), m.r[1])
			m.cycles += 8
		}
	case 0x4ca7:
		{
			m.ip = 0x4cab
			m.wr16(m.r[11], uint16(0x106c), m.r[3])
			m.cycles += 8
		}
	case 0x4cab:
		{
			m.ip = 0x4caf
			m.r[2] = uint16(0x299)
			m.cycles += 3
		}
	case 0x4caf:
		{
			m.ip = 0x4cb0
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4cb0:
		{
			m.ip = 0x4cb3
			target := uint16(19799)
			m.push(0x4cb3)
			m.ip = target
			m.cycles += 19
		}
	case 0x4cb3:
		{
			m.ip = 0x4cb6
			m.r[6] = uint16(47780)
			m.cycles += 2
		}
	case 0x4cb6:
		{
			m.ip = 0x4cb9
			m.r[7] = uint16(6583)
			m.cycles += 2
		}
	case 0x4cb9:
		{
			m.ip = 0x4cbc
			m.r[1] = uint16(51)
			m.cycles += 2
		}
	case 0x4cbc:
		{
			m.ip = 0x4cbd
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4cbd:
		{
			m.ip = 0x4cc0
			m.r[1] = uint16(34)
			m.cycles += 2
		}
	case 0x4cc0:
		{
			m.ip = 0x4cc3
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4cc3:
		{
			m.ip = 0x4cc6
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4cc6:
		{
			m.ip = 0x4cc7
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x4cc7:
		{
			m.ip = 0x4cc8
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x4cc8:
		{
			m.ip = 0x4cca
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19648)
			}
			m.cycles += 17
		}
	case 0x4cca:
		{
			m.ip = 0x4ccd
			m.r[6] = m.alu("add", 16, m.r[6], uint16(46))
			m.cycles += 4
		}
	case 0x4ccd:
		{
			m.ip = 0x4cd0
			m.r[7] = m.alu("add", 16, m.r[7], uint16(46))
			m.cycles += 4
		}
	case 0x4cd0:
		{
			m.ip = 0x4cd1
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4cd1:
		{
			m.ip = 0x4cd3
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19644)
			}
			m.cycles += 17
		}
	case 0x4cd3:
		{
			m.ip = 0x4cd5
			m.r[3] = m.alu("xor", 16, m.r[3], m.r[3])
			m.cycles += 4
		}
	case 0x4cd5:
		{
			m.ip = 0x4cd8
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x4cd8:
		{
			m.ip = 0x4cd9
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4cd9:
		{
			m.ip = 0x4cdc
			m.r[6] = uint16(47824)
			m.cycles += 2
		}
	case 0x4cdc:
		{
			m.ip = 0x4ce0
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x346b))
			m.cycles += 8
		}
	case 0x4ce0:
		{
			m.ip = 0x4ce3
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x4ce3:
		{
			m.ip = 0x4ce6
			m.r[1] = uint16(66)
			m.cycles += 2
		}
	case 0x4ce6:
		{
			m.ip = 0x4ce7
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4ce7:
		{
			m.ip = 0x4cea
			m.r[1] = uint16(12)
			m.cycles += 2
		}
	case 0x4cea:
		{
			m.ip = 0x4ced
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4ced:
		{
			m.ip = 0x4cf0
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4cf0:
		{
			m.ip = 0x4cf1
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x4cf1:
		{
			m.ip = 0x4cf2
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x4cf2:
		{
			m.ip = 0x4cf4
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19690)
			}
			m.cycles += 17
		}
	case 0x4cf4:
		{
			m.ip = 0x4cf7
			m.r[6] = m.alu("add", 16, m.r[6], uint16(68))
			m.cycles += 4
		}
	case 0x4cf7:
		{
			m.ip = 0x4cfa
			m.r[7] = m.alu("add", 16, m.r[7], uint16(68))
			m.cycles += 4
		}
	case 0x4cfa:
		{
			m.ip = 0x4cfb
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4cfb:
		{
			m.ip = 0x4cfd
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19686)
			}
			m.cycles += 17
		}
	case 0x4cfd:
		{
			m.ip = 0x4cfe
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4cfe:
		{
			m.ip = 0x4d00
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19672)
			}
			m.cycles += 17
		}
	case 0x4d00:
		{
			m.ip = 0x4d03
			m.r[6] = uint16(53760)
			m.cycles += 2
		}
	case 0x4d03:
		{
			m.ip = 0x4d06
			m.r[7] = uint16(18104)
			m.cycles += 2
		}
	case 0x4d06:
		{
			m.ip = 0x4d09
			m.r[1] = uint16(3)
			m.cycles += 2
		}
	case 0x4d09:
		{
			m.ip = 0x4d0a
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4d0a:
		{
			m.ip = 0x4d0d
			m.r[1] = uint16(9)
			m.cycles += 2
		}
	case 0x4d0d:
		{
			m.ip = 0x4d0e
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4d0e:
		{
			m.ip = 0x4d11
			m.r[1] = uint16(32)
			m.cycles += 2
		}
	case 0x4d11:
		{
			m.ip = 0x4d14
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4d14:
		{
			m.ip = 0x4d17
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4d17:
		{
			m.ip = 0x4d18
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x4d18:
		{
			m.ip = 0x4d19
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x4d19:
		{
			m.ip = 0x4d1b
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19729)
			}
			m.cycles += 17
		}
	case 0x4d1b:
		{
			m.ip = 0x4d1e
			m.r[6] = m.alu("add", 16, m.r[6], uint16(48))
			m.cycles += 4
		}
	case 0x4d1e:
		{
			m.ip = 0x4d21
			m.r[7] = m.alu("add", 16, m.r[7], uint16(48))
			m.cycles += 4
		}
	case 0x4d21:
		{
			m.ip = 0x4d22
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4d22:
		{
			m.ip = 0x4d24
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19725)
			}
			m.cycles += 17
		}
	case 0x4d24:
		{
			m.ip = 0x4d28
			m.r[7] = m.alu("add", 16, m.r[7], uint16(3840))
			m.cycles += 4
		}
	case 0x4d28:
		{
			m.ip = 0x4d29
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4d29:
		{
			m.ip = 0x4d2b
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19721)
			}
			m.cycles += 17
		}
	case 0x4d2b:
		{
			m.ip = 0x4d2e
			m.r[7] = uint16(18824)
			m.cycles += 2
		}
	case 0x4d2e:
		{
			m.ip = 0x4d31
			m.r[1] = uint16(2)
			m.cycles += 2
		}
	case 0x4d31:
		{
			m.ip = 0x4d32
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4d32:
		{
			m.ip = 0x4d35
			m.r[6] = uint16(52192)
			m.cycles += 2
		}
	case 0x4d35:
		{
			m.ip = 0x4d38
			m.r[1] = uint16(48)
			m.cycles += 2
		}
	case 0x4d38:
		{
			m.ip = 0x4d39
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4d39:
		{
			m.ip = 0x4d3c
			m.r[1] = uint16(32)
			m.cycles += 2
		}
	case 0x4d3c:
		{
			m.ip = 0x4d3f
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4d3f:
		{
			m.ip = 0x4d42
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4d42:
		{
			m.ip = 0x4d43
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x4d43:
		{
			m.ip = 0x4d44
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x4d44:
		{
			m.ip = 0x4d46
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19772)
			}
			m.cycles += 17
		}
	case 0x4d46:
		{
			m.ip = 0x4d49
			m.r[6] = m.alu("add", 16, m.r[6], uint16(48))
			m.cycles += 4
		}
	case 0x4d49:
		{
			m.ip = 0x4d4c
			m.r[7] = m.alu("add", 16, m.r[7], uint16(48))
			m.cycles += 4
		}
	case 0x4d4c:
		{
			m.ip = 0x4d4d
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4d4d:
		{
			m.ip = 0x4d4f
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19768)
			}
			m.cycles += 17
		}
	case 0x4d4f:
		{
			m.ip = 0x4d53
			m.r[7] = m.alu("add", 16, m.r[7], uint16(720))
			m.cycles += 4
		}
	case 0x4d53:
		{
			m.ip = 0x4d54
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4d54:
		{
			m.ip = 0x4d56
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19761)
			}
			m.cycles += 17
		}
	case 0x4d56:
		{
			m.ip = 0x4d57
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4d57:
		{
			m.ip = 0x4d5a
			m.r[6] = uint16(53346)
			m.cycles += 2
		}
	case 0x4d5a:
		{
			m.ip = 0x4d5d
			m.r[7] = uint16(26881)
			m.cycles += 2
		}
	case 0x4d5d:
		{
			m.ip = 0x4d60
			m.r[1] = uint16(13)
			m.cycles += 2
		}
	case 0x4d60:
		{
			m.ip = 0x4d61
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4d61:
		{
			m.ip = 0x4d64
			m.r[1] = uint16(11)
			m.cycles += 2
		}
	case 0x4d64:
		{
			m.ip = 0x4d67
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4d67:
		{
			m.ip = 0x4d6a
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4d6a:
		{
			m.ip = 0x4d6b
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x4d6b:
		{
			m.ip = 0x4d6c
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x4d6c:
		{
			m.ip = 0x4d6e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19812)
			}
			m.cycles += 17
		}
	case 0x4d6e:
		{
			m.ip = 0x4d71
			m.r[6] = m.alu("add", 16, m.r[6], uint16(69))
			m.cycles += 4
		}
	case 0x4d71:
		{
			m.ip = 0x4d74
			m.r[7] = m.alu("add", 16, m.r[7], uint16(69))
			m.cycles += 4
		}
	case 0x4d74:
		{
			m.ip = 0x4d75
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4d75:
		{
			m.ip = 0x4d77
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19808)
			}
			m.cycles += 17
		}
	case 0x4d77:
		{
			m.ip = 0x4d7a
			m.r[6] = uint16(54786)
			m.cycles += 2
		}
	case 0x4d7a:
		{
			m.ip = 0x4d7d
			m.r[7] = uint16(26952)
			m.cycles += 2
		}
	case 0x4d7d:
		{
			m.ip = 0x4d80
			m.r[1] = uint16(13)
			m.cycles += 2
		}
	case 0x4d80:
		{
			m.ip = 0x4d81
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4d81:
		{
			m.ip = 0x4d84
			m.r[1] = uint16(7)
			m.cycles += 2
		}
	case 0x4d84:
		{
			m.ip = 0x4d87
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4d87:
		{
			m.ip = 0x4d8a
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4d8a:
		{
			m.ip = 0x4d8b
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x4d8b:
		{
			m.ip = 0x4d8c
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x4d8c:
		{
			m.ip = 0x4d8e
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19844)
			}
			m.cycles += 17
		}
	case 0x4d8e:
		{
			m.ip = 0x4d91
			m.r[6] = m.alu("add", 16, m.r[6], uint16(73))
			m.cycles += 4
		}
	case 0x4d91:
		{
			m.ip = 0x4d94
			m.r[7] = m.alu("add", 16, m.r[7], uint16(73))
			m.cycles += 4
		}
	case 0x4d94:
		{
			m.ip = 0x4d95
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4d95:
		{
			m.ip = 0x4d97
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(19840)
			}
			m.cycles += 17
		}
	case 0x4d97:
		{
			m.ip = 0x4d98
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4d98:
		{
			m.ip = 0x4d9b
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4d9b:
		{
			m.ip = 0x4d9e
			m.r[0] = uint16(517)
			m.cycles += 2
		}
	case 0x4d9e:
		{
			m.ip = 0x4d9f
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4d9f:
		{
			m.ip = 0x4da2
			m.r[3] = uint16(0)
			m.cycles += 2
		}
	case 0x4da2:
		{
			m.ip = 0x4da6
			m.wr16(m.r[11], uint16(0x347f), m.r[3])
			m.cycles += 8
		}
	case 0x4da6:
		{
			m.ip = 0x4daa
			m.r[7] = m.rd16(m.r[11], uint16(m.r[3]+0x34e9))
			m.cycles += 8
		}
	case 0x4daa:
		{
			m.ip = 0x4dae
			m.r[7] = m.rd16(m.r[11], uint16(m.r[7]+0x351d))
			m.cycles += 8
		}
	case 0x4dae:
		{
			m.ip = 0x4daf
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4daf:
		{
			m.ip = 0x4db3
			m.r[0] = m.rd16(m.r[11], uint16(m.r[3]+0x34b5))
			m.cycles += 8
		}
	case 0x4db3:
		{
			m.ip = 0x4db7
			m.r[3] = m.rd16(m.r[11], uint16(m.r[3]+0x3481))
			m.cycles += 8
		}
	case 0x4db7:
		{
			m.ip = 0x4db9
			target := m.r[7]
			m.push(0x4db9)
			m.ip = target
			m.cycles += 19
		}
	case 0x4db9:
		{
			m.ip = 0x4dba
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4dba:
		{
			m.ip = 0x4dbf
			m.wr16(m.r[11], uint16(m.r[3]+0x34e9), m.alu("add", 16, m.rd16(m.r[11], uint16(m.r[3]+0x34e9)), uint16(2)))
			m.cycles += 16
		}
	case 0x4dbf:
		{
			m.ip = 0x4dc4
			m.alu("sub", 16, m.rd16(m.r[11], uint16(m.r[3]+0x34e9)), uint16(64))
			m.cycles += 16
		}
	case 0x4dc4:
		{
			m.ip = 0x4dc6
			if m.cf || m.zf {
				m.ip = uint16(19916)
			}
			m.cycles += 8
		}
	case 0x4dc6:
		{
			m.ip = 0x4dcc
			m.wr16(m.r[11], uint16(m.r[3]+0x34e9), uint16(0))
			m.cycles += 8
		}
	case 0x4dcc:
		{
			m.ip = 0x4dcf
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x4dcf:
		{
			m.ip = 0x4dd2
			m.alu("sub", 16, m.r[3], uint16(50))
			m.cycles += 4
		}
	case 0x4dd2:
		{
			m.ip = 0x4dd4
			if m.cf || m.zf {
				m.ip = uint16(19874)
			}
			m.cycles += 8
		}
	case 0x4dd4:
		{
			m.ip = 0x4dd7
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4dd7:
		{
			m.ip = 0x4dda
			m.r[0] = uint16(261)
			m.cycles += 2
		}
	case 0x4dda:
		{
			m.ip = 0x4ddb
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4ddb:
		{
			m.ip = 0x4ddc
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4ddc:
		{
			m.ip = 0x4de1
			m.wr8(m.r[11], uint16(0x355f), uint16(0))
			m.cycles += 8
		}
	case 0x4de1:
		{
			m.ip = 0x4de4
			target := uint16(20238)
			m.push(0x4de4)
			m.ip = target
			m.cycles += 19
		}
	case 0x4de4:
		{
			m.ip = 0x4de5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4de5:
		{
			m.ip = 0x4dea
			m.wr8(m.r[11], uint16(0x355f), uint16(0))
			m.cycles += 8
		}
	case 0x4dea:
		{
			m.ip = 0x4deb
			m.r[0] = m.unary("dec", 16, m.r[0])
			m.cycles += 3
		}
	case 0x4deb:
		{
			m.ip = 0x4dee
			target := uint16(20238)
			m.push(0x4dee)
			m.ip = target
			m.cycles += 19
		}
	case 0x4dee:
		{
			m.ip = 0x4df1
			m.r[0] = m.alu("add", 16, m.r[0], uint16(2))
			m.cycles += 4
		}
	case 0x4df1:
		{
			m.ip = 0x4df4
			target := uint16(20238)
			m.push(0x4df4)
			m.ip = target
			m.cycles += 19
		}
	case 0x4df4:
		{
			m.ip = 0x4df5
			m.r[0] = m.unary("dec", 16, m.r[0])
			m.cycles += 3
		}
	case 0x4df5:
		{
			m.ip = 0x4df6
			m.r[3] = m.unary("dec", 16, m.r[3])
			m.cycles += 3
		}
	case 0x4df6:
		{
			m.ip = 0x4df9
			target := uint16(20238)
			m.push(0x4df9)
			m.ip = target
			m.cycles += 19
		}
	case 0x4df9:
		{
			m.ip = 0x4dfc
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x4dfc:
		{
			m.ip = 0x4dff
			target := uint16(20238)
			m.push(0x4dff)
			m.ip = target
			m.cycles += 19
		}
	case 0x4dff:
		{
			m.ip = 0x4e00
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4e00:
		{
			m.ip = 0x4e05
			m.wr8(m.r[11], uint16(0x355f), uint16(9))
			m.cycles += 8
		}
	case 0x4e05:
		{
			m.ip = 0x4e08
			m.r[0] = m.alu("add", 16, m.r[0], uint16(2))
			m.cycles += 4
		}
	case 0x4e08:
		{
			m.ip = 0x4e0b
			target := uint16(20238)
			m.push(0x4e0b)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e0b:
		{
			m.ip = 0x4e0e
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(4))
			m.cycles += 4
		}
	case 0x4e0e:
		{
			m.ip = 0x4e11
			target := uint16(20238)
			m.push(0x4e11)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e11:
		{
			m.ip = 0x4e14
			m.r[0] = m.alu("add", 16, m.r[0], uint16(2))
			m.cycles += 4
		}
	case 0x4e14:
		{
			m.ip = 0x4e17
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x4e17:
		{
			m.ip = 0x4e1a
			target := uint16(20238)
			m.push(0x4e1a)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e1a:
		{
			m.ip = 0x4e1d
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x4e1d:
		{
			m.ip = 0x4e20
			target := uint16(20238)
			m.push(0x4e20)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e20:
		{
			m.ip = 0x4e21
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4e21:
		{
			m.ip = 0x4e26
			m.wr8(m.r[11], uint16(0x355f), uint16(9))
			m.cycles += 8
		}
	case 0x4e26:
		{
			m.ip = 0x4e29
			m.r[0] = m.alu("add", 16, m.r[0], uint16(3))
			m.cycles += 4
		}
	case 0x4e29:
		{
			m.ip = 0x4e2c
			target := uint16(20238)
			m.push(0x4e2c)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e2c:
		{
			m.ip = 0x4e2f
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(6))
			m.cycles += 4
		}
	case 0x4e2f:
		{
			m.ip = 0x4e32
			target := uint16(20238)
			m.push(0x4e32)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e32:
		{
			m.ip = 0x4e35
			m.r[0] = m.alu("add", 16, m.r[0], uint16(3))
			m.cycles += 4
		}
	case 0x4e35:
		{
			m.ip = 0x4e38
			m.r[3] = m.alu("add", 16, m.r[3], uint16(3))
			m.cycles += 4
		}
	case 0x4e38:
		{
			m.ip = 0x4e3b
			target := uint16(20238)
			m.push(0x4e3b)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e3b:
		{
			m.ip = 0x4e3e
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(6))
			m.cycles += 4
		}
	case 0x4e3e:
		{
			m.ip = 0x4e41
			target := uint16(20238)
			m.push(0x4e41)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e41:
		{
			m.ip = 0x4e42
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4e42:
		{
			m.ip = 0x4e47
			m.wr8(m.r[11], uint16(0x355f), uint16(6))
			m.cycles += 8
		}
	case 0x4e47:
		{
			m.ip = 0x4e4a
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4))
			m.cycles += 4
		}
	case 0x4e4a:
		{
			m.ip = 0x4e4d
			target := uint16(20238)
			m.push(0x4e4d)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e4d:
		{
			m.ip = 0x4e50
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(8))
			m.cycles += 4
		}
	case 0x4e50:
		{
			m.ip = 0x4e53
			target := uint16(20238)
			m.push(0x4e53)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e53:
		{
			m.ip = 0x4e56
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4))
			m.cycles += 4
		}
	case 0x4e56:
		{
			m.ip = 0x4e59
			m.r[3] = m.alu("add", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x4e59:
		{
			m.ip = 0x4e5c
			target := uint16(20238)
			m.push(0x4e5c)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e5c:
		{
			m.ip = 0x4e5f
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x4e5f:
		{
			m.ip = 0x4e62
			target := uint16(20238)
			m.push(0x4e62)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e62:
		{
			m.ip = 0x4e63
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4e63:
		{
			m.ip = 0x4e68
			m.wr8(m.r[11], uint16(0x355f), uint16(6))
			m.cycles += 8
		}
	case 0x4e68:
		{
			m.ip = 0x4e6b
			m.r[0] = m.alu("add", 16, m.r[0], uint16(5))
			m.cycles += 4
		}
	case 0x4e6b:
		{
			m.ip = 0x4e6e
			target := uint16(20238)
			m.push(0x4e6e)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e6e:
		{
			m.ip = 0x4e71
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(10))
			m.cycles += 4
		}
	case 0x4e71:
		{
			m.ip = 0x4e74
			target := uint16(20238)
			m.push(0x4e74)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e74:
		{
			m.ip = 0x4e77
			m.r[0] = m.alu("add", 16, m.r[0], uint16(5))
			m.cycles += 4
		}
	case 0x4e77:
		{
			m.ip = 0x4e7a
			m.r[3] = m.alu("add", 16, m.r[3], uint16(5))
			m.cycles += 4
		}
	case 0x4e7a:
		{
			m.ip = 0x4e7d
			target := uint16(20238)
			m.push(0x4e7d)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e7d:
		{
			m.ip = 0x4e80
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(10))
			m.cycles += 4
		}
	case 0x4e80:
		{
			m.ip = 0x4e83
			target := uint16(20238)
			m.push(0x4e83)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e83:
		{
			m.ip = 0x4e84
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4e84:
		{
			m.ip = 0x4e87
			target := uint16(20287)
			m.push(0x4e87)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e87:
		{
			m.ip = 0x4e88
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4e88:
		{
			m.ip = 0x4e89
			m.r[0] = m.unary("dec", 16, m.r[0])
			m.cycles += 3
		}
	case 0x4e89:
		{
			m.ip = 0x4e8c
			target := uint16(20287)
			m.push(0x4e8c)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e8c:
		{
			m.ip = 0x4e8f
			m.r[0] = m.alu("add", 16, m.r[0], uint16(2))
			m.cycles += 4
		}
	case 0x4e8f:
		{
			m.ip = 0x4e92
			target := uint16(20287)
			m.push(0x4e92)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e92:
		{
			m.ip = 0x4e93
			m.r[0] = m.unary("dec", 16, m.r[0])
			m.cycles += 3
		}
	case 0x4e93:
		{
			m.ip = 0x4e94
			m.r[3] = m.unary("dec", 16, m.r[3])
			m.cycles += 3
		}
	case 0x4e94:
		{
			m.ip = 0x4e97
			target := uint16(20287)
			m.push(0x4e97)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e97:
		{
			m.ip = 0x4e9a
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x4e9a:
		{
			m.ip = 0x4e9d
			target := uint16(20287)
			m.push(0x4e9d)
			m.ip = target
			m.cycles += 19
		}
	case 0x4e9d:
		{
			m.ip = 0x4e9e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4e9e:
		{
			m.ip = 0x4ea1
			m.r[0] = m.alu("add", 16, m.r[0], uint16(2))
			m.cycles += 4
		}
	case 0x4ea1:
		{
			m.ip = 0x4ea4
			target := uint16(20287)
			m.push(0x4ea4)
			m.ip = target
			m.cycles += 19
		}
	case 0x4ea4:
		{
			m.ip = 0x4ea7
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(4))
			m.cycles += 4
		}
	case 0x4ea7:
		{
			m.ip = 0x4eaa
			target := uint16(20287)
			m.push(0x4eaa)
			m.ip = target
			m.cycles += 19
		}
	case 0x4eaa:
		{
			m.ip = 0x4ead
			m.r[0] = m.alu("add", 16, m.r[0], uint16(2))
			m.cycles += 4
		}
	case 0x4ead:
		{
			m.ip = 0x4eb0
			m.r[3] = m.alu("add", 16, m.r[3], uint16(2))
			m.cycles += 4
		}
	case 0x4eb0:
		{
			m.ip = 0x4eb3
			target := uint16(20287)
			m.push(0x4eb3)
			m.ip = target
			m.cycles += 19
		}
	case 0x4eb3:
		{
			m.ip = 0x4eb6
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x4eb6:
		{
			m.ip = 0x4eb9
			target := uint16(20287)
			m.push(0x4eb9)
			m.ip = target
			m.cycles += 19
		}
	case 0x4eb9:
		{
			m.ip = 0x4eba
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4eba:
		{
			m.ip = 0x4ebd
			m.r[0] = m.alu("add", 16, m.r[0], uint16(3))
			m.cycles += 4
		}
	case 0x4ebd:
		{
			m.ip = 0x4ec0
			target := uint16(20287)
			m.push(0x4ec0)
			m.ip = target
			m.cycles += 19
		}
	case 0x4ec0:
		{
			m.ip = 0x4ec3
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(6))
			m.cycles += 4
		}
	case 0x4ec3:
		{
			m.ip = 0x4ec6
			target := uint16(20287)
			m.push(0x4ec6)
			m.ip = target
			m.cycles += 19
		}
	case 0x4ec6:
		{
			m.ip = 0x4ec9
			m.r[0] = m.alu("add", 16, m.r[0], uint16(3))
			m.cycles += 4
		}
	case 0x4ec9:
		{
			m.ip = 0x4ecc
			m.r[3] = m.alu("add", 16, m.r[3], uint16(3))
			m.cycles += 4
		}
	case 0x4ecc:
		{
			m.ip = 0x4ecf
			target := uint16(20287)
			m.push(0x4ecf)
			m.ip = target
			m.cycles += 19
		}
	case 0x4ecf:
		{
			m.ip = 0x4ed2
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(6))
			m.cycles += 4
		}
	case 0x4ed2:
		{
			m.ip = 0x4ed5
			target := uint16(20287)
			m.push(0x4ed5)
			m.ip = target
			m.cycles += 19
		}
	case 0x4ed5:
		{
			m.ip = 0x4ed6
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4ed6:
		{
			m.ip = 0x4ed9
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4))
			m.cycles += 4
		}
	case 0x4ed9:
		{
			m.ip = 0x4edc
			target := uint16(20287)
			m.push(0x4edc)
			m.ip = target
			m.cycles += 19
		}
	case 0x4edc:
		{
			m.ip = 0x4edf
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(8))
			m.cycles += 4
		}
	case 0x4edf:
		{
			m.ip = 0x4ee2
			target := uint16(20287)
			m.push(0x4ee2)
			m.ip = target
			m.cycles += 19
		}
	case 0x4ee2:
		{
			m.ip = 0x4ee5
			m.r[0] = m.alu("add", 16, m.r[0], uint16(4))
			m.cycles += 4
		}
	case 0x4ee5:
		{
			m.ip = 0x4ee8
			m.r[3] = m.alu("add", 16, m.r[3], uint16(4))
			m.cycles += 4
		}
	case 0x4ee8:
		{
			m.ip = 0x4eeb
			target := uint16(20287)
			m.push(0x4eeb)
			m.ip = target
			m.cycles += 19
		}
	case 0x4eeb:
		{
			m.ip = 0x4eee
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(8))
			m.cycles += 4
		}
	case 0x4eee:
		{
			m.ip = 0x4ef1
			target := uint16(20287)
			m.push(0x4ef1)
			m.ip = target
			m.cycles += 19
		}
	case 0x4ef1:
		{
			m.ip = 0x4ef2
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4ef2:
		{
			m.ip = 0x4ef5
			m.r[0] = m.alu("add", 16, m.r[0], uint16(5))
			m.cycles += 4
		}
	case 0x4ef5:
		{
			m.ip = 0x4ef8
			target := uint16(20287)
			m.push(0x4ef8)
			m.ip = target
			m.cycles += 19
		}
	case 0x4ef8:
		{
			m.ip = 0x4efb
			m.r[0] = m.alu("sub", 16, m.r[0], uint16(10))
			m.cycles += 4
		}
	case 0x4efb:
		{
			m.ip = 0x4efe
			target := uint16(20287)
			m.push(0x4efe)
			m.ip = target
			m.cycles += 19
		}
	case 0x4efe:
		{
			m.ip = 0x4f01
			m.r[0] = m.alu("add", 16, m.r[0], uint16(5))
			m.cycles += 4
		}
	case 0x4f01:
		{
			m.ip = 0x4f04
			m.r[3] = m.alu("add", 16, m.r[3], uint16(5))
			m.cycles += 4
		}
	case 0x4f04:
		{
			m.ip = 0x4f07
			target := uint16(20287)
			m.push(0x4f07)
			m.ip = target
			m.cycles += 19
		}
	case 0x4f07:
		{
			m.ip = 0x4f0a
			m.r[3] = m.alu("sub", 16, m.r[3], uint16(10))
			m.cycles += 4
		}
	case 0x4f0a:
		{
			m.ip = 0x4f0d
			target := uint16(20287)
			m.push(0x4f0d)
			m.ip = target
			m.cycles += 19
		}
	case 0x4f0d:
		{
			m.ip = 0x4f0e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4f0e:
		{
			m.ip = 0x4f0f
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4f0f:
		{
			m.ip = 0x4f10
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4f10:
		{
			m.ip = 0x4f12
			m.set8(1, 0, ((m.r[3] >> 0) & 255))
			m.cycles += 2
		}
	case 0x4f12:
		{
			m.ip = 0x4f15
			m.r[2] = uint16(80)
			m.cycles += 2
		}
	case 0x4f15:
		{
			m.ip = 0x4f17
			m.multiplyDivide("mul", m.r[2])
			m.cycles += 100
		}
	case 0x4f17:
		{
			m.ip = 0x4f19
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4f19:
		{
			m.ip = 0x4f1b
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4f1b:
		{
			m.ip = 0x4f1d
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4f1d:
		{
			m.ip = 0x4f1f
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x4f1f:
		{
			m.ip = 0x4f22
			m.set8(1, 0, m.alu("and", 8, ((m.r[1]>>0)&255), uint16(7)))
			m.cycles += 4
		}
	case 0x4f22:
		{
			m.ip = 0x4f25
			m.set8(1, 0, m.alu("xor", 8, ((m.r[1]>>0)&255), uint16(7)))
			m.cycles += 4
		}
	case 0x4f25:
		{
			m.ip = 0x4f27
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4f27:
		{
			m.ip = 0x4f29
			m.set8(0, 8, m.shift("shl", 8, ((m.r[0]>>8)&255), ((m.r[1]>>0)&255)))
			m.cycles += 8
		}
	case 0x4f29:
		{
			m.ip = 0x4f2c
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4f2c:
		{
			m.ip = 0x4f2e
			m.set8(0, 0, uint16(8))
			m.cycles += 2
		}
	case 0x4f2e:
		{
			m.ip = 0x4f2f
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4f2f:
		{
			m.ip = 0x4f32
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x4f32:
		{
			m.ip = 0x4f33
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4f33:
		{
			m.ip = 0x4f36
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[3])))
			m.cycles += 8
		}
	case 0x4f36:
		{
			m.ip = 0x4f39
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x355f)))
			m.cycles += 8
		}
	case 0x4f39:
		{
			m.ip = 0x4f3c
			m.wr8(m.r[8], uint16(m.r[3]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4f3c:
		{
			m.ip = 0x4f3d
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4f3d:
		{
			m.ip = 0x4f3e
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x4f3e:
		{
			m.ip = 0x4f3f
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4f3f:
		{
			m.ip = 0x4f40
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4f40:
		{
			m.ip = 0x4f41
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4f41:
		{
			m.ip = 0x4f43
			m.set8(1, 0, ((m.r[3] >> 0) & 255))
			m.cycles += 2
		}
	case 0x4f43:
		{
			m.ip = 0x4f46
			m.r[2] = uint16(80)
			m.cycles += 2
		}
	case 0x4f46:
		{
			m.ip = 0x4f48
			m.multiplyDivide("mul", m.r[2])
			m.cycles += 100
		}
	case 0x4f48:
		{
			m.ip = 0x4f4a
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4f4a:
		{
			m.ip = 0x4f4c
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4f4c:
		{
			m.ip = 0x4f4e
			m.r[3] = m.shift("shr", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4f4e:
		{
			m.ip = 0x4f50
			m.r[3] = m.alu("add", 16, m.r[3], m.r[0])
			m.cycles += 4
		}
	case 0x4f50:
		{
			m.ip = 0x4f53
			m.set8(1, 0, m.alu("and", 8, ((m.r[1]>>0)&255), uint16(7)))
			m.cycles += 4
		}
	case 0x4f53:
		{
			m.ip = 0x4f56
			m.set8(1, 0, m.alu("xor", 8, ((m.r[1]>>0)&255), uint16(7)))
			m.cycles += 4
		}
	case 0x4f56:
		{
			m.ip = 0x4f58
			m.set8(0, 8, uint16(1))
			m.cycles += 2
		}
	case 0x4f58:
		{
			m.ip = 0x4f59
			m.push(m.r[0])
			m.cycles += 11
		}
	case 0x4f59:
		{
			m.ip = 0x4f5a
			m.push(m.r[3])
			m.cycles += 11
		}
	case 0x4f5a:
		{
			m.ip = 0x4f5b
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x4f5b:
		{
			m.ip = 0x4f5f
			m.r[3] = m.alu("add", 16, m.r[3], uint16(41197))
			m.cycles += 4
		}
	case 0x4f5f:
		{
			m.ip = 0x4f64
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x347f)), uint16(36))
			m.cycles += 16
		}
	case 0x4f64:
		{
			m.ip = 0x4f66
			if m.cf || m.zf {
				m.ip = uint16(20330)
			}
			m.cycles += 8
		}
	case 0x4f66:
		{
			m.ip = 0x4f6a
			m.r[3] = m.alu("add", 16, m.r[3], uint16(60309))
			m.cycles += 4
		}
	case 0x4f6a:
		{
			m.ip = 0x4f6c
			m.set8(1, 8, ((m.r[0] >> 8) & 255))
			m.cycles += 2
		}
	case 0x4f6c:
		{
			m.ip = 0x4f6e
			m.set8(1, 8, m.shift("shl", 8, ((m.r[1]>>8)&255), ((m.r[1]>>0)&255)))
			m.cycles += 8
		}
	case 0x4f6e:
		{
			m.ip = 0x4f70
			m.r[6] = m.r[3]
			m.cycles += 2
		}
	case 0x4f70:
		{
			m.ip = 0x4f72
			m.set8(3, 0, m.alu("xor", 8, ((m.r[3]>>0)&255), ((m.r[3]>>0)&255)))
			m.cycles += 4
		}
	case 0x4f72:
		{
			m.ip = 0x4f75
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4f75:
		{
			m.ip = 0x4f78
			m.r[0] = uint16(772)
			m.cycles += 2
		}
	case 0x4f78:
		{
			m.ip = 0x4f79
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4f79:
		{
			m.ip = 0x4f7c
			m.set8(3, 8, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x4f7c:
		{
			m.ip = 0x4f7e
			m.set8(3, 8, m.alu("and", 8, ((m.r[3]>>8)&255), ((m.r[1]>>8)&255)))
			m.cycles += 4
		}
	case 0x4f7e:
		{
			m.ip = 0x4f80
			m.set8(3, 8, m.unary("neg", 8, ((m.r[3]>>8)&255)))
			m.cycles += 3
		}
	case 0x4f80:
		{
			m.ip = 0x4f82
			m.r[3] = m.shift("rol", 16, m.r[3], uint16(1))
			m.cycles += 8
		}
	case 0x4f82:
		{
			m.ip = 0x4f84
			m.set8(0, 8, m.unary("dec", 8, ((m.r[0]>>8)&255)))
			m.cycles += 3
		}
	case 0x4f84:
		{
			m.ip = 0x4f86
			if m.sf == m.of {
				m.ip = uint16(20344)
			}
			m.cycles += 8
		}
	case 0x4f86:
		{
			m.ip = 0x4f8a
			m.wr8(m.r[11], uint16(0x355f), ((m.r[3] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4f8a:
		{
			m.ip = 0x4f8b
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x4f8b:
		{
			m.ip = 0x4f8c
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4f8c:
		{
			m.ip = 0x4f8d
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x4f8d:
		{
			m.ip = 0x4f8f
			m.set8(0, 8, m.shift("shl", 8, ((m.r[0]>>8)&255), ((m.r[1]>>0)&255)))
			m.cycles += 8
		}
	case 0x4f8f:
		{
			m.ip = 0x4f92
			m.r[2] = uint16(974)
			m.cycles += 2
		}
	case 0x4f92:
		{
			m.ip = 0x4f94
			m.set8(0, 0, uint16(8))
			m.cycles += 2
		}
	case 0x4f94:
		{
			m.ip = 0x4f95
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4f95:
		{
			m.ip = 0x4f98
			m.r[0] = uint16(3)
			m.cycles += 2
		}
	case 0x4f98:
		{
			m.ip = 0x4f99
			m.output(m.r[2], m.r[0], 16)
			m.cycles += 8
		}
	case 0x4f99:
		{
			m.ip = 0x4f9c
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[3])))
			m.cycles += 8
		}
	case 0x4f9c:
		{
			m.ip = 0x4f9f
			m.set8(0, 0, m.rd16(m.r[11], uint16(0x355f)))
			m.cycles += 8
		}
	case 0x4f9f:
		{
			m.ip = 0x4fa2
			m.wr8(m.r[8], uint16(m.r[3]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x4fa2:
		{
			m.ip = 0x4fa3
			m.r[3] = m.pop()
			m.cycles += 8
		}
	case 0x4fa3:
		{
			m.ip = 0x4fa4
			m.r[0] = m.pop()
			m.cycles += 8
		}
	case 0x4fa4:
		{
			m.ip = 0x4fa5
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x4fa5:
		{
			m.ip = 0x4fa8
			target := uint16(19142)
			m.push(0x4fa8)
			m.ip = target
			m.cycles += 19
		}
	case 0x4fa8:
		{
			m.ip = 0x4fac
			m.r[6] = m.rd16(m.r[11], uint16(0x345b))
			m.cycles += 8
		}
	case 0x4fac:
		{
			m.ip = 0x4fb0
			m.r[7] = m.rd16(m.r[11], uint16(0x345f))
			m.cycles += 8
		}
	case 0x4fb0:
		{
			m.ip = 0x4fb3
			target := uint16(20622)
			m.push(0x4fb3)
			m.ip = target
			m.cycles += 19
		}
	case 0x4fb3:
		{
			m.ip = 0x4fb7
			m.r[6] = m.rd16(m.r[11], uint16(0x345d))
			m.cycles += 8
		}
	case 0x4fb7:
		{
			m.ip = 0x4fbb
			m.r[7] = m.rd16(m.r[11], uint16(0x3461))
			m.cycles += 8
		}
	case 0x4fbb:
		{
			m.ip = 0x4fbe
			target := uint16(20622)
			m.push(0x4fbe)
			m.ip = target
			m.cycles += 19
		}
	case 0x4fbe:
		{
			m.ip = 0x4fc2
			m.wr16(m.r[11], uint16(0x3467), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x3467))))
			m.cycles += 15
		}
	case 0x4fc2:
		{
			m.ip = 0x4fc6
			m.wr16(m.r[11], uint16(0x3469), m.unary("inc", 16, m.rd16(m.r[11], uint16(0x3469))))
			m.cycles += 15
		}
	case 0x4fc6:
		{
			m.ip = 0x4fcc
			m.alu("sub", 16, m.rd16(m.r[11], uint16(0x3467)), uint16(291))
			m.cycles += 16
		}
	case 0x4fcc:
		{
			m.ip = 0x4fce
			if !m.zf {
				m.ip = uint16(20444)
			}
			m.cycles += 8
		}
	case 0x4fce:
		{
			m.ip = 0x4fd4
			m.wr16(m.r[11], uint16(0x3467), uint16(0))
			m.cycles += 8
		}
	case 0x4fd4:
		{
			m.ip = 0x4fd8
			m.wr16(m.r[11], uint16(0x3463), m.unary("neg", 16, m.rd16(m.r[11], uint16(0x3463))))
			m.cycles += 15
		}
	case 0x4fd8:
		{
			m.ip = 0x4fdc
			m.wr16(m.r[11], uint16(0x3465), m.unary("neg", 16, m.rd16(m.r[11], uint16(0x3465))))
			m.cycles += 15
		}
	case 0x4fdc:
		{
			m.ip = 0x4fe0
			m.r[6] = m.rd16(m.r[11], uint16(0x3463))
			m.cycles += 8
		}
	case 0x4fe0:
		{
			m.ip = 0x4fe4
			m.r[7] = m.rd16(m.r[11], uint16(0x3465))
			m.cycles += 8
		}
	case 0x4fe4:
		{
			m.ip = 0x4fe8
			m.wr16(m.r[11], uint16(0x345f), m.alu("add", 16, m.rd16(m.r[11], uint16(0x345f)), m.r[6]))
			m.cycles += 16
		}
	case 0x4fe8:
		{
			m.ip = 0x4fec
			m.wr16(m.r[11], uint16(0x3461), m.alu("add", 16, m.rd16(m.r[11], uint16(0x3461)), m.r[7]))
			m.cycles += 16
		}
	case 0x4fec:
		{
			m.ip = 0x4fef
			m.r[0] = m.rd16(m.r[11], uint16(0x3469))
			m.cycles += 8
		}
	case 0x4fef:
		{
			m.ip = 0x4ff2
			m.alu("sub", 16, m.r[0], uint16(31))
			m.cycles += 4
		}
	case 0x4ff2:
		{
			m.ip = 0x4ff4
			if !m.zf {
				m.ip = uint16(20489)
			}
			m.cycles += 8
		}
	case 0x4ff4:
		{
			m.ip = 0x4ffa
			m.wr16(m.r[11], uint16(0x3469), uint16(65535))
			m.cycles += 8
		}
	case 0x4ffa:
		{
			m.ip = 0x5000
			m.wr16(m.r[11], uint16(0x345b), uint16(28000))
			m.cycles += 8
		}
	case 0x5000:
		{
			m.ip = 0x5006
			m.wr16(m.r[11], uint16(0x345d), m.alu("add", 16, m.rd16(m.r[11], uint16(0x345d)), uint16(4650)))
			m.cycles += 16
		}
	case 0x5006:
		{
			m.ip = 0x5008
			m.ip = uint16(20507)
			m.cycles += 15
		}
	case 0x5008:
		{
			m.ip = 0x5009
			m.cycles += 3
		}
	case 0x5009:
		{
			m.ip = 0x500c
			m.alu("sub", 16, m.r[0], uint16(15))
			m.cycles += 4
		}
	case 0x500c:
		{
			m.ip = 0x500e
			if !m.zf {
				m.ip = uint16(20590)
			}
			m.cycles += 8
		}
	case 0x500e:
		{
			m.ip = 0x5014
			m.wr16(m.r[11], uint16(0x345b), m.alu("add", 16, m.rd16(m.r[11], uint16(0x345b)), uint16(4650)))
			m.cycles += 16
		}
	case 0x5014:
		{
			m.ip = 0x501a
			m.wr16(m.r[11], uint16(0x345d), uint16(28000))
			m.cycles += 8
		}
	case 0x501a:
		{
			m.ip = 0x501b
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x501b:
		{
			m.ip = 0x5020
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x345a)), uint16(0))
			m.cycles += 16
		}
	case 0x5020:
		{
			m.ip = 0x5022
			if m.zf {
				m.ip = uint16(20589)
			}
			m.cycles += 8
		}
	case 0x5022:
		{
			m.ip = 0x5027
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x345a)), uint16(2))
			m.cycles += 16
		}
	case 0x5027:
		{
			m.ip = 0x5029
			if m.zf {
				m.ip = uint16(20560)
			}
			m.cycles += 8
		}
	case 0x5029:
		{
			m.ip = 0x502e
			m.alu("sub", 8, m.rd8(m.r[11], uint16(0x345a)), uint16(3))
			m.cycles += 16
		}
	case 0x502e:
		{
			m.ip = 0x5030
			if m.zf {
				m.ip = uint16(20578)
			}
			m.cycles += 8
		}
	case 0x5030:
		{
			m.ip = 0x5033
			target := uint16(557)
			m.push(0x5033)
			m.ip = target
			m.cycles += 19
		}
	case 0x5033:
		{
			m.ip = 0x5035
			m.set8(0, 0, uint16(65))
			m.cycles += 2
		}
	case 0x5035:
		{
			m.ip = 0x5038
			m.alu("sub", 16, m.r[3], uint16(120))
			m.cycles += 4
		}
	case 0x5038:
		{
			m.ip = 0x503a
			if !m.cf {
				m.ip = uint16(20540)
			}
			m.cycles += 8
		}
	case 0x503a:
		{
			m.ip = 0x503c
			m.set8(0, 0, uint16(66))
			m.cycles += 2
		}
	case 0x503c:
		{
			m.ip = 0x5040
			m.alu("sub", 8, ((m.r[0] >> 0) & 255), m.rd8(m.r[11], uint16(0x3459)))
			m.cycles += 16
		}
	case 0x5040:
		{
			m.ip = 0x5042
			if !m.zf {
				m.ip = uint16(20554)
			}
			m.cycles += 8
		}
	case 0x5042:
		{
			m.ip = 0x5045
			target := uint16(9161)
			m.push(0x5045)
			m.ip = target
			m.cycles += 19
		}
	case 0x5045:
		{
			m.ip = 0x5049
			m.wr16(m.r[11], uint16(0x24f7), m.r[2])
			m.cycles += 8
		}
	case 0x5049:
		{
			m.ip = 0x504a
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x504a:
		{
			m.ip = 0x504f
			m.wr8(m.r[11], uint16(0x345a), uint16(0))
			m.cycles += 8
		}
	case 0x504f:
		{
			m.ip = 0x5050
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x5050:
		{
			m.ip = 0x5053
			target := uint16(557)
			m.push(0x5053)
			m.ip = target
			m.cycles += 19
		}
	case 0x5053:
		{
			m.ip = 0x5058
			m.wr8(m.r[11], uint16(0x3459), uint16(65))
			m.cycles += 8
		}
	case 0x5058:
		{
			m.ip = 0x505b
			m.alu("sub", 16, m.r[3], uint16(120))
			m.cycles += 4
		}
	case 0x505b:
		{
			m.ip = 0x505d
			if !m.cf {
				m.ip = uint16(20578)
			}
			m.cycles += 8
		}
	case 0x505d:
		{
			m.ip = 0x5062
			m.wr8(m.r[11], uint16(0x3459), uint16(66))
			m.cycles += 8
		}
	case 0x5062:
		{
			m.ip = 0x5066
			m.wr8(m.r[11], uint16(0x345a), m.unary("dec", 8, m.rd8(m.r[11], uint16(0x345a))))
			m.cycles += 15
		}
	case 0x5066:
		{
			m.ip = 0x5069
			target := uint16(9161)
			m.push(0x5069)
			m.ip = target
			m.cycles += 19
		}
	case 0x5069:
		{
			m.ip = 0x506d
			m.wr16(m.r[11], uint16(0x24f7), m.r[2])
			m.cycles += 8
		}
	case 0x506d:
		{
			m.ip = 0x506e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x506e:
		{
			m.ip = 0x5071
			m.r[0] = m.alu("and", 16, m.r[0], uint16(7))
			m.cycles += 4
		}
	case 0x5071:
		{
			m.ip = 0x5074
			m.alu("sub", 16, m.r[0], uint16(7))
			m.cycles += 4
		}
	case 0x5074:
		{
			m.ip = 0x5076
			if !m.zf {
				m.ip = uint16(20611)
			}
			m.cycles += 8
		}
	case 0x5076:
		{
			m.ip = 0x507c
			m.wr16(m.r[11], uint16(0x345b), m.alu("add", 16, m.rd16(m.r[11], uint16(0x345b)), uint16(4650)))
			m.cycles += 16
		}
	case 0x507c:
		{
			m.ip = 0x5082
			m.wr16(m.r[11], uint16(0x345d), m.alu("add", 16, m.rd16(m.r[11], uint16(0x345d)), uint16(4650)))
			m.cycles += 16
		}
	case 0x5082:
		{
			m.ip = 0x5083
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x5083:
		{
			m.ip = 0x5088
			m.wr16(m.r[11], uint16(0x345b), m.alu("add", 16, m.rd16(m.r[11], uint16(0x345b)), uint16(10)))
			m.cycles += 16
		}
	case 0x5088:
		{
			m.ip = 0x508d
			m.wr16(m.r[11], uint16(0x345d), m.alu("add", 16, m.rd16(m.r[11], uint16(0x345d)), uint16(10)))
			m.cycles += 16
		}
	case 0x508d:
		{
			m.ip = 0x508e
			m.ip = m.pop()
			m.cycles += 8
		}
	case 0x508e:
		{
			m.ip = 0x5091
			m.r[1] = uint16(59)
			m.cycles += 2
		}
	case 0x5091:
		{
			m.ip = 0x5092
			m.push(m.r[1])
			m.cycles += 11
		}
	case 0x5092:
		{
			m.ip = 0x5095
			m.r[1] = uint16(10)
			m.cycles += 2
		}
	case 0x5095:
		{
			m.ip = 0x5098
			m.set8(0, 0, m.rd8(m.r[8], uint16(m.r[6])))
			m.cycles += 8
		}
	case 0x5098:
		{
			m.ip = 0x509b
			m.wr8(m.r[8], uint16(m.r[7]), ((m.r[0] >> 0) & 255))
			m.cycles += 8
		}
	case 0x509b:
		{
			m.ip = 0x509c
			m.r[6] = m.unary("inc", 16, m.r[6])
			m.cycles += 3
		}
	case 0x509c:
		{
			m.ip = 0x509d
			m.r[7] = m.unary("inc", 16, m.r[7])
			m.cycles += 3
		}
	case 0x509d:
		{
			m.ip = 0x509f
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(20629)
			}
			m.cycles += 17
		}
	case 0x509f:
		{
			m.ip = 0x50a2
			m.r[6] = m.alu("add", 16, m.r[6], uint16(70))
			m.cycles += 4
		}
	case 0x50a2:
		{
			m.ip = 0x50a5
			m.r[7] = m.alu("add", 16, m.r[7], uint16(70))
			m.cycles += 4
		}
	case 0x50a5:
		{
			m.ip = 0x50a6
			m.r[1] = m.pop()
			m.cycles += 8
		}
	case 0x50a6:
		{
			m.ip = 0x50a8
			m.r[1]--
			if m.r[1] != 0 {
				m.ip = uint16(20625)
			}
			m.cycles += 17
		}
	case 0x50a8:
		{
			m.ip = 0x50a9
			m.ip = m.pop()
			m.cycles += 8
		}
	default:
		m.fail("invalid game execution state")
	}
}
