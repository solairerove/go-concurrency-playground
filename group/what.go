package group

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

// What tbd
func What() {
	fmt.Println("\nWhat")

	ch := make(chan int)

	go func() {
		ch <- doSomething(5)
	}()

	fmt.Println(<-ch)
}
