package zvm

type dictionaryTableBase ByteAddress
type dictionaryAddress WordZstringAddress

func wordSeparatorBase(base DictionaryBase) ByteAddress {
	return ByteAddress(base)
}

func countWordSeparators(story Story) uint8 {
	dictionaryBase := dictionaryBase(story)
	wordSeparatorBase := wordSeparatorBase(dictionaryBase)
	count := story.readByte(wordSeparatorBase)

	return count
}

func getDictionaryTableBase(story Story) dictionaryTableBase {
	dictionaryBase := dictionaryBase(story)
	count := countWordSeparators(story)
	offset := uint16(count) + 1
	address := uint16(dictionaryBase) + offset
	return dictionaryTableBase(address)
}

func getEntryLength(story Story, base dictionaryTableBase) uint8 {
	return story.readByte(ByteAddress(base))
}

func toDictionaryAddress(base dictionaryTableBase, n DictionaryNumber, entryLength uint8) dictionaryAddress {
	entryBase := uint16(base) + 3
	offset := uint16(n) * uint16(entryLength)
	return dictionaryAddress(entryBase + offset)
}

/* GetEntryCount gets the number of dictionary entries from the story */
func GetEntryCount(story Story) uint16 {
	base := getDictionaryTableBase(story)
	address := WordAddress(base + 1)
	return story.readWord(address)
}

/*  GetDictionaryEntry returns the dictionary entry at position n from story */
func GetDictionaryEntry(story Story, n DictionaryNumber) string {
	base := getDictionaryTableBase(story)
	entryLength := getEntryLength(story, base)
	address := toDictionaryAddress(base, n, entryLength)
	zstringAddress := Zstring(address)

	return ReadZstring(story, zstringAddress)
}
