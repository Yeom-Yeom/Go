package main
import "fmt"

func main(){
	for i:=2; i<10; i++{
		for j:=1; j<10 ;j++{
			if i%2==1{
				if i<j{
					fmt.Printf("\n")
					break
				}
				fmt.Printf("%d x %d = %d\n",i,j,i*j)
			}else{
				continue
			}
		}
	}
	//테스트는 같은 형태로 나오지만 구름IDE 상에서 제출을 하면 오답으로 뜬다.
	
}