package zvm

type instructionT struct {
}

type instruction uint16

type opCodeFormT byte

const (
	longForm     = opCodeFormT(1)
	shortForm    = opCodeFormT(2)
	variableForm = opCodeFormT(3)
	extendedForm = opCodeFormT(4)
)

type opCountT byte

const (
	op0     = opCountT(0)
	op1     = opCountT(1)
	op2     = opCountT(2)
	variant = opCountT(3)
)

func getOpcodeForm(story Story, address ByteAddress) opCodeFormT {
	byte := uint16(story.readByte(address))
	switch fetchBits(byte, bit7, twoBits) {
	case 3:
		return variableForm

	case 2:
		if byte == 190 {
			return extendedForm
		}
		return shortForm
	}

	return longForm
}

func getOpCount(story Story, address ByteAddress, form opCodeFormT) opCountT {
	byte := uint16(story.readByte(address))
	switch form {
	case shortForm:
		if fetchBits(byte, bit5, twoBits) == 3 {
			return op0
		}
		return op1

	case longForm:
		return op2

	case variableForm:
		if fetchBit(byte, bit5) {
			return variant
		}
		return op2

	default:
		return variant
	}
}

func decodeOpCode(story Story, address ByteAddress, form opCodeFormT, opCount opCountT) opCodeT {
	byte := uint16(story.readByte(address))

	if form == extendedForm {
		const maximumExtended = 29
		ext := story.readByte(address + 1)
		if ext > maximumExtended {
			return illegal
		}

		return extByteCodes[ext]
	}

	switch opCount {
	case op0:
		return op0ByteCodes[fetchBits(byte, bit3, fourBits)]
	case op1:
		return op1ByteCodes[fetchBits(byte, bit3, fourBits)]
	case op2:
		return op2ByteCodes[fetchBits(byte, bit4, fiveBits)]
	case variant:
		return varByteCodes[fetchBits(byte, bit4, fiveBits)]
	}

	return opCodeT(0)
}

func decode(story Story) instructionT {
	return instructionT{}
}
