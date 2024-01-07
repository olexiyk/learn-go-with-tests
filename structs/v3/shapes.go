package main

type Rectangle struct {
	w, h float64
}

func Perimeter(r Rectangle) float64 {
	return (r.w + r.h) * 2
}

func Area(r Rectangle) float64 {
	return r.w * r.h
}
