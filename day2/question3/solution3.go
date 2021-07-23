package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
type person struct{
	mu sync.Mutex
	balance int
	transactions []int
}

func update(p *person,add int){
	p.mu.Lock()
	p.balance+=add
	p.mu.Unlock()
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
		{balance: 500,transactions:[]int{20,-30}},
		{balance:500,transactions:[]int{10}},
		{balance:500,transactions:[]int{-10}}}

	for key,_ := range persons{
		wg.Add(1)
		go processAllTransaction(&persons[key])
	}
	wg.Wait()
	for _,p := range persons{
		fmt.Println(p.balance)
	}
}
