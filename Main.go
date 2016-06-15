package main

import (
	"fmt"
	"io/ioutil"

	"zvm/zvm"
)

func main() {
	dat, _ := ioutil.ReadFile("minizork.z3")
	story := zvm.StoryFromBytes(dat)

	var n zvm.DictionaryNumber
	for count := zvm.GetEntryCount(story); count > 0; count-- {
		e := zvm.GetDictionaryEntry(story, n)
		fmt.Println(e)

		n++
	}
}
