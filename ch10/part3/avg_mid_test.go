package main
import . "fmt"

func main(){
	var a map[string]int
	a = make(map[string]int)
	var class string
	var score int

	var avg float32
	for {
		Scanf("%s %d",&class, &score)
		if class == "0"{
			break
		}
		a[class] = score
	}
	for _,j := range a{
		avg+=float32(j)
	}
	avg /= float32(len(a))
	
	for x, y := range a{
		Printf("%s %d\n",x,y)
	}
	Printf("%.2f",avg)
}