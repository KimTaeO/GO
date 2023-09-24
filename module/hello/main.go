package main

import (
	"fmt"
	"rack/repository/Go/module/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Printf(message)
}
