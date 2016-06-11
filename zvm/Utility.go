package zvm

func fetchBits(word uint16, high BitNumber, length BitSize) uint {
	mask := ^(-1 << length)
	shr := (uint8(high) - uint8(length)) + 1
	return uint(word >> shr) & uint(mask)
}

func fetchBit(word uint16, n BitNumber) bool {
	return fetchBits(word, n, BitSize(1)) == 1
}

func (address WordAddress) High() ByteAddress {
	return ByteAddress(uint16(address))
}

func (address WordAddress) Low() ByteAddress {
	return ByteAddress(uint16(address) + 1)
}

func offsetByteAddress(address ByteAddress, offset int) ByteAddress {
	return ByteAddress(int(address) + offset)
}

func offsetWordAddress(address WordAddress, offset int) WordAddress {
	const wordSize = 2
	offset *= wordSize
	return WordAddress(int(address) + offset)
}

func nextWordAddress(address WordAddress) WordAddress {
	return offsetWordAddress(address, 1)
}

func decodeWordAddress(address WordZstringAddress) Zstring {
	value := int(address) * 2
	return Zstring(value)
}