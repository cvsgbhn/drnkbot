package main

import (
	"fmt"
	"time"
)

func whatthedrink(num int) string {
	if num%2 == 0 {
		return "чй"
	} else {
		return "кф"
	}
}

func main() {
	hours, minutes, seconds := time.Now().Clock()
	fmt.Println(whatthedrink(hours + minutes + seconds - 5))
}
