package main

import "fmt"

type ChickenInterface interface {
	TakeABath() error
	PluckFeathers() error
}

type Human struct{}

func (h Human) Eat(chicken ChickenInterface) error {
	err := chicken.TakeABath()
	if err != nil {
		return err
	}
	err = chicken.PluckFeathers()
	if err != nil {
		return err
	}
	fmt.Println("Human: Eating")
	return nil
}
