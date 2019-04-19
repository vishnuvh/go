package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

	var x = 7

	fmt.Println("Hello World")
	fmt.Println("Time is ", time.Now())
	fmt.Println("Random number is", rand.Intn(100))
	fmt.Println("Square root of", x, "is", math.Sqrt(float64(x)))
	fmt.Println("pi value is", math.Pi)
	fmt.Println("add(42, 13)", add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Printf("split 17 into ")
	fmt.Println(split(17))
}
