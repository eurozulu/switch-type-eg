package main

import (
	"fmt"
	"reflect"
)

func main() {
	cars := []Car{
		&DieselCar{litresPer100km: 6, fuelLitres: 3},
		&PetrolCar{litresPer100km: 14, fuelLitres: 3},
		&ElectricCar{KWHPer100km: 2, batteryLevelKWH: 1},
	}
	fmt.Println("prefuel")
	showRange(cars)
	refuel(cars, 10)
	fmt.Println()
	fmt.Println("postfuel")
	showRange(cars)
}

func showRange(cars []Car) {
	for _, c := range cars {
		fmt.Printf("%s has %.2f km range\n", reflect.ValueOf(c).Elem().Type().Name(), c.Range())
	}
}

// refuel applies the 'level' to all the cars, casting each type into the concrete struct pointer as 'v'
// Each car has an implementation specific method, which is called within the case of each car
// The case statements should have the type value the switch casting is expected to result in
func refuel(cars []Car, level float64) {
	for _, c := range cars {
		switch v := c.(type) {
		case *DieselCar:
			v.FillWithDiesel(level)

		case *PetrolCar:
			v.FillWithPetrol(level)

		case *ElectricCar:
			v.Charge(level)
		}
	}
}
