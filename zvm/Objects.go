package zvm

type propertyDefaultsTableT uint16
type objectTreeBaseT uint16
type objectAddressT uint16
type propertyHeaderAddressT uint16

const (
	defaultPropertyTableEntrySize = 2
)

func defaultPropertyTableSize(story Story) uint8 {
	if story.v3OrLower() {
		return 31
	}

	return 63
}

func getEntrySize(story Story) uint8 {
	if story.v3OrLower() {
		return 9
	}

	return 14
}

func treeBase(story Story) objectTreeBaseT {
	base := objectBase(story)
	tableSize := uint16(defaultPropertyTableSize(story) * defaultPropertyTableEntrySize)
	return objectTreeBaseT(uint16(base) + tableSize)
}

func toObjectAddress(story Story, n ObjectNumber) objectAddressT {
	base := treeBase(story)
	entrySize := getEntrySize(story)
	offset := uint16(n-1) * uint16(entrySize)
	return objectAddressT(uint16(base) + offset)
}

func getPropertyHeader(story Story, n ObjectNumber) propertyHeaderAddressT {
	var objectPropertyOffset uint16
	if story.v3OrLower() {
		objectPropertyOffset = 7
	} else {
		objectPropertyOffset = 12
	}

	objectAddress := toObjectAddress(story, n)
	offset := uint16(objectAddress) + objectPropertyOffset
	propertyHeaderAddress := story.readWord(WordAddress(offset))
	return propertyHeaderAddressT(propertyHeaderAddress)
}

func GetParent(story Story, n ObjectNumber) ObjectNumber {
	address := toObjectAddress(story, n)
	if story.v3OrLower() {
		parent := story.readByte(ByteAddress(address) + 4)
		return ObjectNumber(parent)
	}

	parent := story.readWord(WordAddress(address) + 6)
	return ObjectNumber(parent)
}

func GetSibling(story Story, n ObjectNumber) ObjectNumber {
	address := toObjectAddress(story, n)
	if story.v3OrLower() {
		parent := story.readByte(ByteAddress(address) + 5)
		return ObjectNumber(parent)
	}

	parent := story.readWord(WordAddress(address) + 8)
	return ObjectNumber(parent)
}

func GetChild(story Story, n ObjectNumber) ObjectNumber {
	address := toObjectAddress(story, n)
	if story.v3OrLower() {
		parent := story.readByte(ByteAddress(address) + 6)
		return ObjectNumber(parent)
	}

	parent := story.readWord(WordAddress(address) + 10)
	return ObjectNumber(parent)
}

func GetObjectName(story Story, n ObjectNumber) string {
	address := getPropertyHeader(story, n)

	if length := story.readByte(ByteAddress(address)); length == 0 {
		return "<unnamed>"
	}

	zstring := Zstring(address + 1)
	return ReadZstring(story, zstring)
}
