package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
}

type myInterface interface {
	printCoba(wg *sync.WaitGroup)
	printBisa(wg *sync.WaitGroup)
}

func (t task) printCoba(wg *sync.WaitGroup) {

	for i := 1; i <= 4; i++ {

		fmt.Println("coba1 coba2 coba3 ", i)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}

func (t task) printBisa(wg *sync.WaitGroup) {

	for i := 1; i <= 4; i++ {

		fmt.Println("bisa1 bisa2 bisa3 ", i)
		time.Sleep(time.Millisecond * 200)
	}
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	var myTask1 myInterface = task{}

	wg.Add(2)
	go myTask1.printCoba(&wg)
	go myTask1.printBisa(&wg)

	wg.Wait()
}
