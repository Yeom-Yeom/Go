package main
 
import (
	. "fmt"
	"sync"
	"time"
)
 
func add(a *int, b *int, result *int, wait *sync.WaitGroup) {
	*result = *a + *b
	wait.Wait()
}


func main() {
	var num1, num2 int = 10, 5
	var result int
	var wait sync.WaitGroup
	wait.Add(100)
		
	go add(&num1, &num2, &result, &wait)

	time.Sleep(time.Second)
	
	Println(result)
	
	wait.Done()

}