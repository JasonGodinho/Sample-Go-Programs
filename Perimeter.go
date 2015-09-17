package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func main() {

	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Area of rectangle:", r.area())           //Prints 100 as rectangle area
	fmt.Println("Perimeter of rectangle:", r.perimeter()) //Prints 40 as rectangle perimeter

	c := Circle{0, 0, 7}
	fmt.Println("Area of circle:", c.area())           //Prints 153.93 as circle area
	fmt.Println("Perimeter of circle:", c.perimeter()) //Prints 43.98 as circle perimeter
}
