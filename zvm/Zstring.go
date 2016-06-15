package zvm

import (
	"bytes"
	"fmt"
)

var alphabetTable = [][]string{
	{" ", "?", "?", "?", "?", "?", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	{" ", "?", "?", "?", "?", "?", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
	{" ", "?", "?", "?", "?", "?", "?", "\n", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", ",", "!", "?", "_", "#", "'", "\"", "/", "\\", "-", ":", "(", ")"}}

// DisplayBytes
func DisplayBytes(story Story, zstring Zstring) {
	var word uint16
	var zchar1, zchar2, zchar3 uint

	address := WordAddress(uint16(zstring))
	for {
		word = story.readWord(address)
		isEnd := fetchBit(word, bit15)
		zchar1 = fetchBits(word, bit14, fiveBits)
		zchar2 = fetchBits(word, bit9, fiveBits)
		zchar3 = fetchBits(word, bit4, fiveBits)

		fmt.Printf("%02x %s %02x %s %02x %s ", zchar1, alphabetTable[1][zchar1], zchar2, alphabetTable[1][zchar2], zchar3, alphabetTable[1][zchar3])
		if isEnd {
			fmt.Println()
			return
		}

		address = nextWordAddress(address)
	}
}

type zcharT byte
type alphabetT uint

const (
	abbrev   = iota
	alphabet = iota
	leading  = iota
	trailing = iota
)

type stringState struct {
	mode         int
	abbreviation AbbreviationNumber
	alphabet     alphabetT
	leading      zcharT
}

var (
	abbrev0  = stringState{abbreviation: AbbreviationNumber(0), mode: abbrev}
	abbrev32 = stringState{abbreviation: AbbreviationNumber(32), mode: abbrev}
	abbrev64 = stringState{abbreviation: AbbreviationNumber(64), mode: abbrev}

	alphabet0 = stringState{alphabet: alphabetT(0), mode: alphabet}
	alphabet1 = stringState{alphabet: alphabetT(1), mode: alphabet}
	alphabet2 = stringState{alphabet: alphabetT(2), mode: alphabet}
)

func ReadZstring(story Story, zstring Zstring) string {
	var (
		buffer                 bytes.Buffer
		word                   uint16
		zchar1, zchar2, zchar3 zcharT
	)

	state := alphabet0
	address := WordAddress(uint16(zstring))

	for {
		word = story.readWord(address)

		isEnd := fetchBit(word, bit15)

		zchar1 = zcharT(fetchBits(word, bit14, fiveBits))
		zchar2 = zcharT(fetchBits(word, bit9, fiveBits))
		zchar3 = zcharT(fetchBits(word, bit4, fiveBits))

		text1, state1 := processZchar(story, zchar1, state)
		text2, state2 := processZchar(story, zchar2, state1)
		text3, state3 := processZchar(story, zchar3, state2)

		acc := fmt.Sprint(text1, text2, text3)
		buffer.WriteString(acc)

		if isEnd {
			str := buffer.String()
			return str
		}

		address = nextWordAddress(address)
		state = state3
	}
}

func processZchar(story Story, zchar zcharT, state stringState) (string, stringState) {
	if state.mode == alphabet {
		switch zchar {
		case 1:
			return "", abbrev0
		case 2:
			return "", abbrev32
		case 3:
			return "", abbrev64
		case 4:
			return "", alphabet1
		case 5:
			return "", alphabet2
		}

		if zchar == zcharT(6) && state == alphabet2 {
			return "", stringState{mode: leading}
		}

		alphabet := alphabetTable[state.alphabet]
		return alphabet[zchar], alphabet0
	}

	switch {
	case state.mode == abbrev:
		{
			abbrv := uint8(state.abbreviation) + uint8(zchar)
			addr := zstringFromAbbreviation(story, AbbreviationNumber(abbrv))
			abbreviation := ReadZstring(story, addr)
			return abbreviation, alphabet0
		}

	case state.mode == leading:
		{
			return "", stringState{mode: trailing, leading: zchar}
		}

	case state.mode == trailing:
		{
			high := state.leading * 32
			return string(high + zchar), alphabet0
		}
	}

	panic(fmt.Sprint("unknown state:", zchar, state))
}
