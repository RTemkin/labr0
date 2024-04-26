package main

import (
	"sync"

	"labbb4/funclab4"
)

func main() {
	wg := sync.WaitGroup{}
	size := 5
	msg := 12
	rb := funclab4.NewBuf(size)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			for w := 1; w <= msg; w++ {
				rb.Writer(w+100, i)
			}

			wg.Done()
		}(i)

	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {

			defer wg.Done()

			for i := 1; i <= msg; i++ {
				rb.Reader(i)
			}
		}(i)
	}

	wg.Wait()

}
