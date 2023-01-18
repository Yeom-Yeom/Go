package main
import "fmt"

func main(){
	var n int = 0
	var i int = 0
	var j int = 0
	fmt.Scanln(&n)
	for i=0; i<n; i++{
		for j=0; j<i+1; j++{
			if i == j{
				fmt.Print("* ")
			}else{
				fmt.Print("o ")
			}
		}
		fmt.Print("\n")
	}
}