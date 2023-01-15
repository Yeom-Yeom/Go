package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)

	a := float32(num1)
	b := uint(num2)
	c := int64(num3)

	fmt.Printf("%T, %f\n", a, a)
	fmt.Printf("%T, %d\n", b, b)
	fmt.Printf("%T, %d\n", c, c)

}
