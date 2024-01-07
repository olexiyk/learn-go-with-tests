package main

import "math"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}

func (self Rectangle) Area() float64 {
	return self.Width * self.Height
}

func (self Circle) Area() float64 {
	return self.Radius * self.Radius * math.Pi
}

func (self Triangle) Area() float64 {
	return self.Base * self.Height / 2
}
