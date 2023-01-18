package main
import "fmt"

func main(){
	var result int
	for i:=0; i<10; i++{
		for j:=0; j<10; j++{
			if i!=j{				
				result = (i*10+j) + (j*10+i)
				if result == 99{
					if i == 0{
						fmt.Printf("0%d + %d = %d\n",  i*10+j, j*10+i, result)
					}else if j==0{
						fmt.Printf("%d + 0%d = %d\n",  i*10+j, j*10+i, result)
					}else{
						fmt.Printf("%d + %d = %d\n", i*10+j, j*10+i, result)
					}					
				}else{
					continue
				}
			}			
		}
	}
}