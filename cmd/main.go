
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
	fmt.Printf("globalData pt: %p\n", &globalData[0]) // Replace 0 with os.Getpid()
	fmt.Printf("localDate pt: %p\n", &localData[0]) // Replace 0 with os.Getpid()
	fmt.Printf("heapData pt: %p\n", &heapData[0]) // Replace 0 with os.Getpid()
	
	for i :=0; i < len(heapData); i += 4096{
		heapData[i] = 12

		time.Sleep(2 *time.Second)
		fmt.Printf("touched byte: %d\n", i)
	}



	time.Sleep(5 * time.Minute)
}
