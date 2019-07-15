package condition

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Atom tbd
func Atom() {
	fmt.Println("\nAtom")

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:", atomic.LoadInt64(&counter))

			runtime.Gosched()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
