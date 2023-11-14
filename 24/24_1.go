package main

import (
	"fmt"
	"math"
)

type PointInterface interface {
	DistanceTo(Point) float64
}

type Point struct {
	x float64
	y float64
}

var _ PointInterface = (*Point)(nil)

func NewPoint(x float64, y float64) Point {
	return Point{x: x, y: y}
}

func (p *Point) DistanceTo(p2 Point) float64 {
	dx := p.x - p2.x
	dy := p.y - p2.y

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	distance := point1.DistanceTo(point2)

	fmt.Printf("distance between points: %.2f\n", distance)
}
