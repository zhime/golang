package main

import "fmt"

// Animal 需实现全部接口
type Animal interface {
	sing()
	jump()
	rap()
}

type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Printf("%s sing", c.Name)
}

func (c Chicken) jump() {
	fmt.Printf("%s jump", c.Name)
}

func (c Chicken) rap() {
	fmt.Printf("%s rap", c.Name)
}

func main() {
	var c = Chicken{Name: "dd"}
	c.sing()
}
