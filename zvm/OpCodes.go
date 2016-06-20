package zvm

type opCodeT byte

const (
	illegal = opCodeT(0)

	op2_1  = opCodeT(1)
	op2_2  = opCodeT(2)
	op2_3  = opCodeT(3)
	op2_4  = opCodeT(4)
	op2_5  = opCodeT(5)
	op2_6  = opCodeT(6)
	op2_7  = opCodeT(7)
	op2_8  = opCodeT(8)
	op2_9  = opCodeT(9)
	op2_10 = opCodeT(10)
	op2_11 = opCodeT(11)
	op2_12 = opCodeT(12)
	op2_13 = opCodeT(13)
	op2_14 = opCodeT(14)
	op2_15 = opCodeT(15)
	op2_16 = opCodeT(16)
	op2_17 = opCodeT(17)
	op2_18 = opCodeT(18)
	op2_19 = opCodeT(19)
	op2_20 = opCodeT(20)
	op2_21 = opCodeT(21)
	op2_22 = opCodeT(22)
	op2_23 = opCodeT(23)
	op2_24 = opCodeT(24)
	op2_25 = opCodeT(25)
	op2_26 = opCodeT(26)
	op2_27 = opCodeT(27)
	op2_28 = opCodeT(28)

	op1_128 = opCodeT(29)
	op1_129 = opCodeT(30)
	op1_130 = opCodeT(31)
	op1_131 = opCodeT(32)
	op1_132 = opCodeT(33)
	op1_133 = opCodeT(34)
	op1_134 = opCodeT(35)
	op1_135 = opCodeT(36)
	op1_136 = opCodeT(37)
	op1_137 = opCodeT(38)
	op1_138 = opCodeT(39)
	op1_139 = opCodeT(40)
	op1_140 = opCodeT(41)
	op1_141 = opCodeT(42)
	op1_142 = opCodeT(43)
	op1_143 = opCodeT(44)

	op0_176 = opCodeT(45)
	op0_177 = opCodeT(46)
	op0_178 = opCodeT(47)
	op0_179 = opCodeT(48)
	op0_180 = opCodeT(49)
	op0_181 = opCodeT(50)
	op0_182 = opCodeT(51)
	op0_183 = opCodeT(52)
	op0_184 = opCodeT(53)
	op0_185 = opCodeT(54)
	op0_186 = opCodeT(55)
	op0_187 = opCodeT(56)
	op0_188 = opCodeT(57)
	op0_189 = opCodeT(58)
	op0_190 = opCodeT(59)
	op0_191 = opCodeT(60)

	var224 = opCodeT(61)
	var225 = opCodeT(62)
	var226 = opCodeT(63)
	var227 = opCodeT(64)
	var228 = opCodeT(65)
	var229 = opCodeT(66)
	var230 = opCodeT(67)
	var231 = opCodeT(68)
	var232 = opCodeT(69)
	var233 = opCodeT(70)
	var234 = opCodeT(71)
	var235 = opCodeT(72)
	var236 = opCodeT(73)
	var237 = opCodeT(74)
	var238 = opCodeT(75)
	var239 = opCodeT(76)
	var240 = opCodeT(77)
	var241 = opCodeT(78)
	var242 = opCodeT(79)
	var243 = opCodeT(80)
	var244 = opCodeT(81)
	var245 = opCodeT(82)
	var246 = opCodeT(83)
	var247 = opCodeT(84)
	var248 = opCodeT(85)
	var249 = opCodeT(86)
	var250 = opCodeT(87)
	var251 = opCodeT(88)
	var252 = opCodeT(89)
	var253 = opCodeT(90)
	var254 = opCodeT(91)
	var255 = opCodeT(92)

	ext0  = opCodeT(93)
	ext1  = opCodeT(94)
	ext2  = opCodeT(95)
	ext3  = opCodeT(96)
	ext4  = opCodeT(97)
	ext5  = opCodeT(98)
	ext6  = opCodeT(99)
	ext7  = opCodeT(100)
	ext8  = opCodeT(101)
	ext9  = opCodeT(102)
	ext10 = opCodeT(103)
	ext11 = opCodeT(104)
	ext12 = opCodeT(105)
	ext13 = opCodeT(106)
	ext14 = opCodeT(107)
	ext16 = opCodeT(108)
	ext17 = opCodeT(109)
	ext18 = opCodeT(110)
	ext19 = opCodeT(111)
	ext20 = opCodeT(112)
	ext21 = opCodeT(113)
	ext22 = opCodeT(114)
	ext23 = opCodeT(115)
	ext24 = opCodeT(116)
	ext25 = opCodeT(117)
	ext26 = opCodeT(118)
	ext27 = opCodeT(119)
	ext28 = opCodeT(120)
	ext29 = opCodeT(121)
)

var (
	extByteCodes = []opCodeT{
		ext0, ext1, ext2, ext3, ext4, ext5, ext6, ext7, ext8, ext9, ext10, ext11, ext12, ext13, ext14, illegal,
		ext16, ext17, ext18, ext19, ext20, ext21, ext22, ext23, ext24, ext25, ext26, ext27, ext28, ext29, illegal, illegal}

	op0ByteCodes = []opCodeT{
		op0_176, op0_177, op0_178, op0_179, op0_180, op0_181, op0_182, op0_183,
		op0_184, op0_185, op0_186, op0_187, op0_188, op0_189, op0_190, op0_191}

	op1ByteCodes = []opCodeT{
		op1_128, op1_129, op1_130, op1_131, op1_132, op1_133, op1_134, op1_135,
		op1_136, op1_137, op1_138, op1_139, op1_140, op1_141, op1_142, op1_143}

	op2ByteCodes = []opCodeT{
		illegal, op2_1, op2_2, op2_3, op2_4, op2_5, op2_6, op2_7, op2_8, op2_9, op2_10,
		op2_11, op2_12, op2_13, op2_14, op2_15, op2_16, op2_17, op2_18, op2_19, op2_20,
		op2_21, op2_22, op2_23, op2_24, op2_25, op2_26, op2_27, op2_28, illegal, illegal, illegal}

	varByteCodes = []opCodeT{
		var224, var225, var226, var227, var228, var229, var230, var231, var232, var233, var234,
		var235, var236, var237, var238, var239, var240, var241, var242, var243, var244, var245,
		var246, var247, var248, var249, var250, var251, var252, var253, var254, var255}
)
