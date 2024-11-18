package main

import (
	"fmt"
	"time"
)

// aksk:LTAI5tKyzGUuNmjhkGHkRvWY:R3bhBjOGiZFmcjm2MS3Pm1GEa6d55W

func main1(){
	for i := 0; i < 20; i++ {
		go func() {
			// fmt.Println(xx)
			time.Sleep(3 * time.Second)
			defer fmt.Println(i)
		}()

	}

	time.Sleep(6 * time.Second)
}