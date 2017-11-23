package main

import (
	"errors"
	"fmt"
)

type Car struct {
	Model string
}
type Cars []Car

func (cars *Cars) Add(car Car) {
	myCars = append(myCars, car)
}

func (cars *Cars) Find(model string) (*Car, error) {
	for _, car := range *cars {
		if car.Model == model {
			return &car, nil
		}
	}
	return nil, errors.New("car not found")
}

var myCars Cars

func main() {
	myCars.Add(Car{"IS250"})
	myCars.Add(Car{"Blazer"})
	myCars.Add(Car{"Highlander"})

	car, err := myCars.Find("Highlander")
	if err != nil {
		fmt.Printf("ERROR: %v", car)
	} else {
		fmt.Printf("Found %v", car)
	}
}

