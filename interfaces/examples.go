package main

import "math"

type Rectangle struct {
	Weight float64
	Height float64
}

type Cirle struct {
	Radius float64
}

func (c Cirle) FindSquare() float64 {
	return c.Radius * math.Pi
}

func (c Rectangle) FindSquare() float64 {
	return c.Height * c.Weight
}

type FigureSquareFinder interface {
	FindSquare() float64
}
