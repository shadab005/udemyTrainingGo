package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

//Defining Square
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) fun() {
	fmt.Println("Hello from rectangle fun method")
}

//Defining Circle
type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println("Area = ", s.area())
}

func main() {
	fmt.Println("----Welcome to interfaces chapter-----:")

	sq := Square{10.1}
	c1 := Circle{2.5}
	info(sq)
	info(&c1)

	fmt.Println(c1.area())
	fmt.Println((&c1).area())

}

type MyCar struct {
	name string
}

func (c MyCar) move() {
	fmt.Println("Car is moving")
}

func (c MyCar) radius() float64 {
	return 5.4
}

// play

type Car interface {
	Wheel
	move()
}

type Wheel interface {
	radius() float64
}

func carInfo(c Car) {
	fmt.Println(c)

}
