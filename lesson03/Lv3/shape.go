package main

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r *Rectangle) Area() float64 {
	return r.length * r.width
}
