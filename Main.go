package main

import (
	"fmt"
	"zvm/zvm"	
	"io/ioutil"
)

func main() {
	dat, _ := ioutil.ReadFile("minizork.z3")
	story := zvm.StoryFromBytes(dat)
	 	
	addr := zvm.Zstring(0xb106)
	zstring := zvm.ReadZstring(story, addr)

	fmt.Println(zstring)
}
