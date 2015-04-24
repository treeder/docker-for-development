package main

import (
	"fmt"
	"github.com/iron-io/iron_go/worker"
)

type Person struct {
	Name string
}

func main() {
	worker.ParseFlags() // just ot use the dependency
	fmt.Println("Hello Go!")
}
