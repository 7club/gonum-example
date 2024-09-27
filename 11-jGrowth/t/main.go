package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 10; i < 200; i += 10 {
		fmt.Println(i, 5*math.Pow(0.99, float64(i)))
	}
}
