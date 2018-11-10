package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ordersProcessed := 0
	cashier := func(orderNum int) {
		if ordersProcessed < 10 {
			fmt.Println("Processing order", orderNum)
			ordersProcessed++
		} else {
			fmt.Println("I am tired! I want to take rest!", orderNum)
		}
		wg.Done()
	}

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(orderNum int) {
			cashier(orderNum)
		}(i)
	}
	wg.Wait()

}
