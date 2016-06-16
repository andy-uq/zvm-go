package zvm

type BitNumber uint8
type BitSize uint8

const bit0 = BitNumber(0)
const bit1 = BitNumber(1)
const bit2 = BitNumber(2)
const bit3 = BitNumber(3)
const bit4 = BitNumber(4)
const bit5 = BitNumber(5)
const bit6 = BitNumber(6)
const bit7 = BitNumber(7)
const bit8 = BitNumber(8)
const bit9 = BitNumber(9)
const bit10 = BitNumber(10)
const bit11 = BitNumber(11)
const bit12 = BitNumber(12)
const bit13 = BitNumber(13)
const bit14 = BitNumber(14)
const bit15 = BitNumber(15)

const fiveBits = BitSize(5)
const fourBits = BitSize(4)
const threeBits = BitSize(3)
const twoBits = BitSize(2)

type ByteAddress uint16
type WordAddress uint16

type AbbreviationNumber uint8
type AbbreviationTableBase uint16

type WordZstringAddress uint16
type Zstring uint32

type DictionaryBase uint16
type DictionaryNumber uint8

type ObjectBase uint16
type ObjectNumber uint8
