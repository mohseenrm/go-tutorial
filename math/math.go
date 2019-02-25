package main

import "fmt"

func add(x, y float64) float64 { return x + y }

func main() {
	var num1, num2 float64 = 1.9, 2.1
	fmt.Println(add(num1, num2))

	num3, num4 := 2.1, 6.66666
	fmt.Println(add(num3, num4))
}
