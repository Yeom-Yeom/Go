package main
import "fmt"

func main(){
	a := [2][2]int{{7,3},{5,2}}
	var result bool = false
	d := a[0][0]*a[1][1] - a[0][1]*a[1][0]
	if d !=0{
		result = true
		fmt.Println(result)
		fmt.Printf("-%d %d\n",a[1][1],a[0][1])
		fmt.Printf("%d -%d\n",a[1][0],a[0][0])
	}else{
		fmt.Print(result)
	}
}