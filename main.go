package main

import (
	"fmt"

	race "github.com/solairerove/go-concurrency-playground/condition"
	wait "github.com/solairerove/go-concurrency-playground/group"
)

func main() {
	fmt.Println("go-concurrency-playground")

	wait.WaitGroup()
	wait.What()

	race.Race()
}
