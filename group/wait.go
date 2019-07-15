package group

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// WaitGroup tbd
func WaitGroup() {
	fmt.Println("\nWaitGroup")

	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	fmt.Println()
	bar()

	fmt.Println("\nCPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
