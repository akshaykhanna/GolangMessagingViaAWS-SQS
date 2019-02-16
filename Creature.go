package main

import "fmt"

type Creature struct {
	Name   string
	IsReal bool
}

func (c Creature) Dump() {
	fmt.Printf("Name : %s and IsReal: %s", c.Name, c.IsReal)
}

type FlyingCreature struct {
	WingSpan int
	Creature
}

func (fc FlyingCreature) Dump() {
	fmt.Printf("Name: %s, IsReal: %s, & WingSpan: %d", fc.Name, fc.IsReal, fc.WingSpan)
}

type MyInterface interface {
	Meth1()
	Meth2()
	Meth3()
}

type implementClass struct {
}

func (ic implementClass) Meth1() {
	fmt.Println("Meth1...")
}
func (ic implementClass) Meth2() {
	fmt.Println("Meth2...")
}
func (ic implementClass) Meth3() {
	fmt.Println("Meth3...")
}

func NewMyInterface() MyInterface {
	return &implementClass{}
}