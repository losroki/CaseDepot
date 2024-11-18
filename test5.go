package main

import (
	"fmt"
	"time"
)

func main5(){
	for i := 0; i < 20; i++ {
		go func() {
			// fmt.Println(xx)
			time.Sleep(3 * time.Second)
			defer fmt.Println(i)
		}()

	}

	// ak:LTAI5tKyzGUuNmjhkGHkRvWY
	// sk:R3bhBjOGiZFmcjm2MS3Pm1GEa6d55W

	time.Sleep(6 * time.Second)
}