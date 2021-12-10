package main

type Car interface {
	Range() float64
}

type PetrolCar struct {
	litresPer100km float64
	fuelLitres     float64
}

func (c *PetrolCar) FillWithPetrol(litres float64) {
	c.fuelLitres += litres
}

func (c PetrolCar) Range() float64 {
	if c.litresPer100km == 0 {
		return 0
	}
	return (c.fuelLitres / c.litresPer100km) * 100
}

type ElectricCar struct {
	KWHPer100km     float64
	batteryLevelKWH float64
}

func (c *ElectricCar) Charge(kwh float64) {
	c.batteryLevelKWH += kwh
}

func (c ElectricCar) Range() float64 {
	if c.KWHPer100km == 0 {
		return 0
	}
	return (c.batteryLevelKWH / c.KWHPer100km) * 100
}

type DieselCar struct {
	litresPer100km float64
	fuelLitres     float64
}

func (c *DieselCar) FillWithDiesel(litres float64) {
	c.fuelLitres += litres
}

func (c DieselCar) Range() float64 {
	if c.litresPer100km == 0 {
		return 0
	}
	return (c.fuelLitres / c.litresPer100km) * 100
}
