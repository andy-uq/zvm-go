package zvm

type memory_t struct {
	bytes []byte
}

func (bytes memory_t) size() uint16 {
	len := len(bytes.bytes)
	return uint16(len)
}

func isInRange(address ByteAddress, size uint16) bool {
	return uint16(address) < size
}

func isOutOfRange(address ByteAddress, size uint16) bool {
	return !isInRange(address, size)
}

func memoryFromBytes(bytes []byte) memory_t {
	return memory_t{bytes: bytes}
}

func (memory memory_t) read(address ByteAddress) byte {
	if isOutOfRange(address, memory.size()) {
		panic("address is out of range")
	}

	return memory.bytes[address]
}

func (memory memory_t) writeByte(address ByteAddress, value byte) memory_t {
	if isOutOfRange(address, memory.size()) {
		panic("address is out of range")
	}

	memory.bytes[address] = value
	return memory
}
