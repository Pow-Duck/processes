package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i += 1
		time.Sleep(time.Second)
		fmt.Println("hello world: ", i)
	}
}
