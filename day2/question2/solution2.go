package main

import (
	"fmt"
	"math/rand"
	"sync"
)
var wg sync.WaitGroup
func add(sum *int, rating int){
	*sum+=rating
	wg.Done()
}

func main(){
	var ratings []int

	//generating random ratings
	for i:=0;i<200;i++{
		someRandomNumber := rand.Intn(5)
		ratings[i]=someRandomNumber
	}

	sum:=0
	l:=len(ratings)
	for _,rating:=range ratings{
		wg.Add(1)
		go add(&sum,rating)
	}

	wg.Wait()
	fmt.Println(float64(sum)/float64(l))
}
