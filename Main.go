package main

import (
	"fmt"
	"zvm/zvm"	
	"io/ioutil"
)

func main() {
	dat, _ := ioutil.ReadFile("minizork.z3")
	story := zvm.StoryFromBytes(dat)
	 	
	zvm.DisplayAbbreviation(story, zvm.AbbreviationNumber(0))
	zvm.DisplayAbbreviation(story, zvm.AbbreviationNumber(4))

	fmt.Println(story.Version())
}
