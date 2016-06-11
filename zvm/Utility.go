package zvm

func FetchBits(word int, high BitNumber, length BitSize) int {
	mask := ^(-1 << length)
	shr := (uint8(high) - uint8(length)) + 1
	return (word >> shr) & mask
}
