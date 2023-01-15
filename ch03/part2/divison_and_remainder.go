package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Scanln(&num1, &num2)
	fmt.Printf("몫 : %d 나머지 :  %d", num1/num2, num1%num2)
}
