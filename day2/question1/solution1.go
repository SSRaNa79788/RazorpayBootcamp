package main

import (
	"fmt"
	"sync"
)
type characterData struct{
	char string
	frequency int
	mu sync.Mutex
}

var wg sync.WaitGroup
var frequencyOfCharacters map[string]int
var mapCharacterToCharacterMutex map[string]*characterData

func incrementByOne(character string){
	characterDataNodeForCurrentCharacter := mapCharacterToCharacterMutex[character]
	if characterDataNodeForCurrentCharacter==nil{
		wg.Done()
		return
	}
	characterDataNodeForCurrentCharacter.mu.Lock()
	frequencyOfCharacters[character]++
	characterDataNodeForCurrentCharacter.mu.Unlock()
	wg.Done()
}

func countFrequency(s string) {
	l := len(s)
	for i := 0; i < l; i++ {
		wg.Add(1)
		go incrementByOne(string(s[i]))
	}
	wg.Done()
}
func main() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}

	frequencyOfCharacters = make(map[string]int)
	mapCharacterToCharacterMutex=make(map[string]*characterData)

	initialCharacter :="a"
	for i:=0;i<26;i++{
		currentCharacter := string(initialCharacter[0] +byte(i))
		currentCharacterDataNode := characterData{char:currentCharacter, frequency:0}
		mapCharacterToCharacterMutex[currentCharacter] = &currentCharacterDataNode
	}

	//for k,v := range mapCharacterToCharacterMutex{
	//	fmt.Println(k,v)
	//}

	for _, word := range words {
		wg.Add(1)
		go countFrequency(word)
	}
	wg.Wait()
	//fmt.Println(string(byte("a"[0])+1))

	s := "a"
	for {
		if f, ok := frequencyOfCharacters[s]; ok {
			fmt.Println(s, " : ", f)
		}
		if s == "z" {
			break
		}
		s = string(s[0] + 1)
	}

}
