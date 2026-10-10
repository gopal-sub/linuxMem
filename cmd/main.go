
package main

import (
	"fmt"
	"time"
	"os"
)

var globalData [1024 * 1024]byte

func main() {
	localData := [1024]byte{}
	heapData := make([]byte, 10*1024*1024) //10 Mb

	globalData[0] = 1
	localData[0] = 2
	heapData[0] = 3
	processID := os.Getpid()
	fmt.Printf("PID: %d\n", processID) // Replace 0 with os.Getpid()

	time.Sleep(5 * time.Minute)
}
