package main

import (
	"fmt"
	"time"
	_ "embed"
)

//go:embed test8
var aksk string

func main7(){
	for i := 0; i < 20; i++ {
		go func() {
			// fmt.Println(xx)
			time.Sleep(3 * time.Second)
			defer fmt.Println(i)
		}()

	}

	fmt.Println(aksk)

	time.Sleep(6 * time.Second)
}