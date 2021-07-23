package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

func add(sum *int, rating int){

	someRandomTime := int(rand.Intn(1000))
	for i:=0;i<someRandomTime ;i++{
		time.Sleep(time.Microsecond)
	}
	//time.Sleep(someRandomTime*time.Millisecond)// I am waiting for some random time duration

	mu.Lock()
	*sum+=rating
	mu.Unlock()

	wg.Done()
}

func main(){
	rand.Seed(time.Now().UnixNano()) //to get different pseudo random numbers

	var ratings []int

	//generating random ratings
	for i:=0;i<200;i++{
		someRandomNumber := rand.Intn(5)
		ratings=append(ratings,someRandomNumber)
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
