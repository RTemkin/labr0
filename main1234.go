package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup

func produser(num int, ch chan int, boo chan bool) {

	num1 := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		defer wg.Done()
		mutex.Lock()
		num1 += num
		mutex.Unlock()

		ch <- num1

	}

	boo <- true
}

func cosumer(ch chan int) {
	for i := 0; i < 5; i++ {

		a := <-ch

		fmt.Println(i, a)
	}

}

func main() {

	ch := make(chan int)
	boo := make(chan bool)

	for i := 0; i < 3; i++ {
		go produser(10, ch, boo)
	}
	for i := 0; i < 3; i++ {
		go cosumer(ch)
	}
	for i := 0; i < 3; i++ {
		<-boo
	}

	wg.Wait()
	fmt.Println("The End")
}
