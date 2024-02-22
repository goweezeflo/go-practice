package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range 200 {
		go printStuff(rune(i))
	}
	time.Sleep(3000 * time.Millisecond)
}

func printStuff(stuff rune) {
	fmt.Println(stuff)
}
