package zvm

type ImmutableBytes struct {
	bytes []uint8
	edits map[ByteAddress]uint8
}

func (bytes ImmutableBytes) Size() uint16 {
	len := len(bytes.bytes)
	return uint16(len)
}

func IsInRange(address ByteAddress, size uint16) bool {
	return uint16(address) < size
}

func IsOutOfRange(address ByteAddress, size uint16) bool {
	return !IsInRange(address, size)
}

func Make(bytes []uint8) ImmutableBytes {
	return ImmutableBytes{bytes: bytes}
}

func (bytes ImmutableBytes) ReadByte(address ByteAddress) uint8 {
	if IsOutOfRange(address, bytes.Size()) {
		panic("address is out of range")
	}

	byte, edited := bytes.edits[address]
	if edited {
		return byte

	}
	return bytes.bytes[address]
}

func (bytes ImmutableBytes) WriteByte(address ByteAddress, value uint8) ImmutableBytes {
	if IsOutOfRange(address, bytes.Size()) {
		panic("address is out of range")
	}

	bytes.edits[address] = value
	return bytes
}
