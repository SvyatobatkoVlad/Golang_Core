package main

import "fmt"

type Point struct{
	X, Y int
}

func (p *Point) movePoint(x, y int) {
	p.X += x
	p.Y += y
}

func movePointer(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func movePointPtr(p *Point, x, y int){
	p.X += x
	p.Y += y
}


func main(){
	p := Point{1, 2}
	fmt.Println(p)
	//fmt.Println(movePointer(p, 10, 10))
	movePointPtr(&p, 10, 10)
	fmt.Println(p)

	p.movePoint(10, 10)
	fmt.Println(p)
}
