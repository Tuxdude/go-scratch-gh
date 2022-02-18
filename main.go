package main

import (
	"fmt"
	"os"
)

func run() int {
	fmt.Println("This command is a no-op, does literally nothing except print this line")
	fmt.Println("Another no-op output")
	return 0
}

func main() {
	os.Exit(run())
}
