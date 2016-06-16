package main

import (
	"fmt"
	"io/ioutil"
	"zvm/zvm"
)

func main() {
	dat, _ := ioutil.ReadFile("minizork.z3")
	story := zvm.StoryFromBytes(dat)

	n1 := zvm.GetObjectName(story, zvm.ObjectNumber(1))
	fmt.Println(n1)
}
