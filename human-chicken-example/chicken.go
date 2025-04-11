package main

import "fmt"

type Chicken struct{}

func (c *Chicken) TakeABath() error {
	fmt.Println("Chicken: Taking a bath")
	return nil
}

func (c *Chicken) PluckFeathers() error {
	fmt.Println("Chicken: Plucking feathers")
	return nil
}
