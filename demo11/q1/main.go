package main

import (
	"fmt"
	"math/rand"
)

type Notifier interface {
	SendInt(ch chan<- int)
}

func main() {
	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}
}

func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}
