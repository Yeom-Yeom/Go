package main
import . "fmt"

var g float32 = 9.8
type calc func(float32, float32) float32

func calMechEnergy(calc calc, ke float32, pe float32) float32 {
	result := calc(ke,pe)
	 
  return result
}  

func main() {
	var m, v, h float32
	Scanln(&m, &v, &h)
	

	kinEnergy := func(m float32, v float32) float32 {
		return ((m*v*v)/2)
	}
	potEnergy := func(m float32, h float32) float32{
		return (m*g*h)
	}
	
	ke := calMechEnergy(kinEnergy,m,v)
	pe := calMechEnergy(potEnergy,m,h)
	Println(ke,pe,ke+pe)
}