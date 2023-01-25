package main

import . "fmt"

var g float32 = 9.8


type energy struct {
	m,v,h,ke,pe,me float32
}

func newEnergy() *energy {
	new_e := energy{}
	return &new_e
}

func (e *energy) k_energy(){
	(*e).ke = (*e).m*(*e).v*(*e).v/2
	Print((*e).ke)
}

func (e *energy) p_energy(){
	(*e).pe = (*e).m*g*(*e).h
	Print((*e).pe)
}


func main() {
	var n int
	var m,v,h float32
	Scanln(&n)	

	e := make([]energy,n)

	for i:=0; i<n; i++ {	
		Scanln(&m,&v,&h)
		e[i].m = m
		e[i].v = v
		e[i].h = h
	}
	
	for i:=0; i<n; i++ {
		e[i].k_energy()
		Print(" ")
		e[i].p_energy()
		Print(" ")
		Println(e[i].ke+e[i].pe)
	}
	
}

