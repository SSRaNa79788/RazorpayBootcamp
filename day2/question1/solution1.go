package main

import (
	"fmt"
	"sync"
)

var fre map[string]int
var wg sync.WaitGroup

func countFrequency(s string) {
	l := len(s)
	for i := 0; i < l; i++ {
		fre[string(s[i])]++
	}
	wg.Done()
}
func main() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}

	fre = make(map[string]int)
	for _, word := range words {
		wg.Add(1)
		go countFrequency(word)
	}
	wg.Wait()
	//fmt.Println(string(byte("a"[0])+1))
	s := "a"
	for {
		if f, ok := fre[s]; ok {
			fmt.Println(s, " : ", f)
		}
		if s == "z" {
			break
		}
		s = string(s[0] + 1)
	}

}
