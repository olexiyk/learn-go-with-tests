package main

import "math"

type Rectangle struct {
	w, h float64
}

type Circle struct {
	r float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return (r.w + r.h) * 2
}

func (self Rectangle) Area() float64 {
	return self.w * self.h
}

func (self Circle) Area() float64 {
	return self.r * self.r * math.Pi
}
