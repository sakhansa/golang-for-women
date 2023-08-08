package main

import (
	"fmt"
	"sync"
)

type task struct {
}

type myInterface interface {
	printCoba(i int)
	printBisa(i int)
}

func (t task) printCoba(i int) {

	fmt.Println("coba1 coba2 coba3 ", i)

}

func (t task) printBisa(i int) {

	fmt.Println("bisa1 bisa2 bisa3 ", i)

}

func loop(wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	var myTask1 myInterface = task{}

	for i := 1; i <= 4; i++ {
		mtx.Lock()
		myTask1.printCoba(i)
		myTask1.printBisa(i)
		mtx.Unlock()
	}
}

func main() {

	var wg sync.WaitGroup
	var mtx sync.Mutex

	wg.Add(1)
	go loop(&wg, &mtx)
	wg.Wait()
}
