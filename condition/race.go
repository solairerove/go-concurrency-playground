package condition

import (
	"fmt"
	"runtime"
	"sync"
)

// Race tbd
func Race() {
	fmt.Println("\nRace")

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
