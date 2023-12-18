package model

import (
	"fmt"
	"github.com/google/uuid"
	"math"
)

type Shape interface {
	Area() float64
	Print()
}

type Ellipse struct {
	Id      uuid.UUID
	RadiusA float64
	RadiusB float64
}

func (e Ellipse) Area() float64 {
	return math.Pi * e.RadiusA * e.RadiusB
}

func (e Ellipse) Print() {
	fmt.Printf("Ellipse:\n\tId: %s\n\tRadius A: %f\n\tRadius B: %f\n\tArea: %f\n", e.Id, e.RadiusA, e.RadiusB, e.Area())
}

func CreateEllipse(radiusA, radiusB float64) *Ellipse {
	return &Ellipse{
		Id:      uuid.New(),
		RadiusA: radiusA,
		RadiusB: radiusB,
	}
}

type Rectangle struct {
	Id     uuid.UUID
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Print() {
	fmt.Printf("Rectangle:\n\tId: %s\n\tWidth: %f\n\tHeight: %f\n\tArea: %f\n", r.Id, r.Width, r.Height, r.Area())
	//fmt.Printf("Rectangle: %+v\n", r)
}

func CreateRectangle(height, width float64) *Rectangle {
	return &Rectangle{
		Id:     uuid.New(),
		Height: height,
		Width:  width,
	}
}

type Triangle struct {
	Id     uuid.UUID
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) Print() {
	fmt.Printf("Triangle:\n\tId: %s\n\tBase: %f\n\tHeight: %f\n\tArea: %f\n", t.Id, t.Base, t.Height, t.Area())
}

func CreateTriangle(base, height float64) *Triangle {
	return &Triangle{
		Id:     uuid.New(),
		Base:   base,
		Height: height,
	}
}
