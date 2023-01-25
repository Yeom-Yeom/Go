package main

import (
	"fmt"
	"math"
)

type inter interface{
	area() float32
	volume() float32
}

type cuboid struct{
	a,b,c float32
}

type cylinder struct{
	r,h float32
}

func (c cuboid) area() float32 {
	return 2*c.a*c.b+2*c.b*c.c+2*c.a*c.c
}

func (c cuboid) volume() float32{
	return c.a*c.b*c.c
}

func (cy cylinder) area() float32 {
	return (math.Pi*cy.r*cy.r)*2+(2*math.Pi*cy.r)*cy.h
}

func (cy cylinder) volume() float32 {
	return (math.Pi*cy.r*cy.r*cy.h)
}


func main() {	

	cy1 := cylinder{10,10}
	cy2 := cylinder{4.2,15.6}
	cu1 := cuboid{10.5,20.2,20}
	cu2 := cuboid {4,10,23}
	
	printMeasure(cy1, cy2, cu1, cu2)	
}

func printMeasure(m ...inter) {
	for _,val := range m{
		fmt.Printf("%.2f, %.2f\n",val.area(), val.volume())
	}
}
