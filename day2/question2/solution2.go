package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func add(sum *int, rating int){
	*sum+=rating
	wg.Done()
}

func main(){
	ratings := []int{2,5,3,4,2,3,2,3,4,5,5,5,3,1,4}
	sum:=0
	l:=len(ratings)
	for _,rating:=range ratings{
		wg.Add(1)
		go add(&sum,rating)
	}

	wg.Wait()
	fmt.Println(float64(sum)/float64(l))
}
