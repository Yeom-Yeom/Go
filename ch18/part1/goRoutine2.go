package main
 
import . "fmt"
 
func add(num1 int, num2 int, c chan int){
	c <- num1+num2
}

func main() {
	var num1, num2 int

	Scanln(&num1, &num2)
	c := make(chan int)
	go func(){
		add(num1,num2,c)
	}()
	Print(<-c)
}

// main() 문에서 바로 add 함수를 호출하면 error
// go func()를 사용하여 error solved