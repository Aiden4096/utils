package main

import "fmt"

func main() {
	//fmt.Println(1 + 2 + 4 + 8)
	//fmt.Println(16 + 32 + 64 + 1 + 2 + 4 + 8)
	worker := NewSnowFlake(12, 12)
	for i := 0; i < 10000; i++ {
		fmt.Println(worker.Acquire())
	}
}
