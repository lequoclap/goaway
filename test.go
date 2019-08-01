package main

import (
	"fmt"
	"math"
	"math/rand"
)

func wrap(x string, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Sum(x int, y int) int {
	return x + y
}

func main() {

	Sum(5, 5)

	fmt.Println("Hello world")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Pi is:", math.Pi)
	fmt.Println(wrap("hello", "world"))

	fmt.Println(split(17))
}
