package main
import "fmt"

func main(){
	names := make([]string,0,100)
	var name string
	for {
		fmt.Scanln(&name)
		if name == "1"{
			break;
		}else{
			names = append(names,name)		
		} 
	}
	var result string = names[0]
	for i:=0; i<len(names); i++{
		if(len(names[i])>len(result)){
			result = names[i]
		}
	}
	fmt.Printf("%s %d",result,len(result))
}