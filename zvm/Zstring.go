package zvm

import (
	"fmt"
)

var alphabetTable = []string{"_","?","?","?","?","?","a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

// DisplayBytes
func DisplayBytes(story Story, zstring Zstring) {
    address := WordAddress(uint16(zstring))
    for {
        word := story.readWord(address)
        isEnd := fetchBit(word, bit15)
        zchar1 := fetchBits(word, bit14, fiveBits)
        zchar2 := fetchBits(word, bit9, fiveBits)
        zchar3 := fetchBits(word, bit4, fiveBits)

        fmt.Printf("%02x %s %02x %s %02x %s ", zchar1, alphabetTable[zchar1], zchar2, alphabetTable[zchar2], zchar3, alphabetTable[zchar3])
        if isEnd {
            fmt.Println()
            return
        }

        address = nextWordAddress(address)
    }
}