package main

import (
	"fmt"
	"github.com/iron-io/iron_go/worker"
)

func main() {
	worker.ParseFlags() // just to use the dependency
	fmt.Println("Hello Go!")
}
