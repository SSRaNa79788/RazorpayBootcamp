package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
type person struct{
	balance int
	transactions []int
}

func update(p *person,add int){
	p.balance+=add
	wg.Done()
}

func processAllTransaction(p *person){
	for _,add := range p.transactions{
		wg.Add(1)
		go update(p,add)
	}
	wg.Done()
}

func main(){
	persons := []person{
		{500,[]int{20,-30}},
		{500,[]int{10}},
		{500,[]int{-10}}}

	for key,_ := range persons{
		wg.Add(1)
		go processAllTransaction(&persons[key])
	}
	wg.Wait()
	for _,p := range persons{
		fmt.Println(p.balance)
	}
}
