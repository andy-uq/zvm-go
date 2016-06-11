package zvm

// Story struct
type Story struct {
	dynamic memory_t
	static  memory_t
}

// StoryFromBytes returns a Story from a byte array
func StoryFromBytes(data []byte) Story {
	const headerSize = 64
	const staticMemoryBaseOffset = WordAddress(14)

	len := uint16(len(data))
	if len < headerSize {
		panic("not a valid story file")
	}

	high := staticMemoryBaseOffset.High()
	low := staticMemoryBaseOffset.Low()

	dynamicLength := uint16(data[high])*256 + uint16(data[low])

	if dynamicLength > len {
		panic("not a valid story file")
	}

	dynamic := memoryFromBytes(data[0:dynamicLength])
	static := memoryFromBytes(data[dynamicLength:])

	return Story{dynamic: dynamic, static: static}
}

func (story Story) readByte(address ByteAddress) byte {
	len := story.dynamic.size()
	if isInRange(address, len) {
		return story.dynamic.read(address)
	}

	staticAddress := offsetByteAddress(address, -int(len))
	return story.static.read(staticAddress)
}

func (story Story) writeByte(address ByteAddress, value byte) {
	story.dynamic.writeByte(address, value)
}

func (story Story) readWord(address WordAddress) uint16 {
	high := story.readByte(address.High())
	low := story.readByte(address.Low())

	return uint16(high)*256 + uint16(low)
}

func (story Story) writeWord(address WordAddress, value uint16) {
	high := uint8(value >> 8)
	low := uint8(value & 0xff)

	story.writeByte(address.High(), high)
	story.writeByte(address.Low(), low)
}

// Version of the story
func (story Story) Version() byte {	
	const version = ByteAddress(0)
	return story.readByte(version)
}

func abbreviationTableBase(story Story) AbbreviationTableBase {
	const abbreviationTableBaseOffset = WordAddress(24)
	return AbbreviationTableBase(story.readWord(abbreviationTableBaseOffset))
}

func firstAbbreviationAddress(base AbbreviationTableBase) WordAddress {
	return WordAddress(base)
}

func zstringFromAbbreviation(story Story, n AbbreviationNumber) Zstring {
	const abbreviationTableLength = 96
	if n > abbreviationTableLength {
		panic ("bad offset into abbreviation table")
	}

	base := abbreviationTableBase(story)
	baseAddress := firstAbbreviationAddress(base)
	address := offsetWordAddress(baseAddress, int(n))
	pointer := WordZstringAddress(story.readWord(address))

	return decodeWordAddress(pointer)
}

// DisplayAbbreviation
func DisplayAbbreviation(story Story, n AbbreviationNumber) {
	zstring := zstringFromAbbreviation(story, n)
	DisplayBytes(story, zstring)
}