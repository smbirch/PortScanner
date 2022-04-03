package main

import (
	"fmt"
	"portscanner/port"
)

func main() {
	fmt.Println("Port Scanner")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
