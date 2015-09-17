package TestArea

import (
	"math"
	"testing"
)

type circleTestPair struct {
	ci Circle
	co float64
}

type rectangleTestPair struct {
	ri Rectangle
	ro float64
}

var testCircleAreaCases = []circleTestPair{
	{Circle{x: 0, y: 0, r: 5}, 31.410000},
	{Circle{x: 0, y: 0, r: 7.8}, 49.000000},
	{Circle{x: 0, y: 0, r: 10.3}, 64.710000},
	{Circle{x: 0, y: 0, r: 0.7}, 4.390000},
}

var testCirclePerimeterCases = []circleTestPair{
	{Circle{x: 0, y: 0, r: 5}, 78.530000},
	{Circle{x: 0, y: 0, r: 7.8}, 191.130000},
	{Circle{x: 0, y: 0, r: 10.3}, 333.290000},
	{Circle{x: 0, y: 0, r: 0.7}, 1.530000},
}

var testRectangleAreaCases = []rectangleTestPair{
	{Rectangle{x1: 0, y1: 0, x2: 5, y2: 5}, 25.000000},
	{Rectangle{x1: 0, y1: 0, x2: 5, y2: 10}, 50.000000},
	{Rectangle{x1: 0, y1: 0, x2: 10, y2: 20}, 200.000000},
	{Rectangle{x1: 0, y1: 0, x2: 20, y2: 30}, 600.000000},
}

var testRectanglePerimeterCases = []rectangleTestPair{
	{Rectangle{0, 0, 5, 5}, 20.000000},
	{Rectangle{0, 0, 5, 10}, 30.000000},
	{Rectangle{0, 0, 10, 20}, 60.000000},
	{Rectangle{0, 0, 20, 30}, 100.000000},
}

func TestCirclePerimeter(t *testing.T) {
	for _, val := range testCirclePerimeterCases {
		v := val.ci.perimeter()
		if v != val.co {
			t.Errorf("Circle Perimeter returned %f, expected %f", v, val.co)
		}
	}
}

func TestCircleArea(t *testing.T) {
	for _, val := range testCircleAreaCases {
		v := val.ci.area()
		if v != val.co {
			t.Errorf("Circle Area returned %f, expected %f", v, val.co)
		}
	}
}

func TestRectanglePerimeter(t *testing.T) {
	//r := Rectangle{0, 0, 10}
	for _, val := range testRectanglePerimeterCases {
		v := val.ri.perimeter()
		if v != val.ro {
			t.Errorf("Rectangle Perimeter returned %f, expected %f", v, val.ro)
		}
	}
}

func TestRectangleArea(t *testing.T) {
	for _, val := range testRectangleAreaCases {
		v := val.ri.area()
		if v != val.ro {
			t.Errorf("Rectangle Area returned %f, expected %f", v, val.ro)
		}
	}
}

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

func (c *Circle) area() float64 {
	value := 2 * math.Pi * c.r
	value = float64(int(value*100)) / 100
	return value
}

func (c *Circle) perimeter() float64 {
	value := math.Pi * c.r * c.r
	value = float64(int(value*100)) / 100
	return value
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	value := l * w
	value = float64(int(value*100)) / 100
	return value
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	value := 2 * (l + w)
	value = float64(int(value*100)) / 100
	return value
}
