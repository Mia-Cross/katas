package cakes

import (
	"strings"
)

type CakeMaker interface {
	Name() string
	Price() float64
}

type Bundle struct {
	cakesToPack []CakeMaker
}

// func CalculatePrice(bundle []CakeMaker) float64 {
// 	price := 0.0
// 	for i := range bundle {
// 		price += bundle[i].Price()
// 	}
// 	return price * 0.9
// }

type Topping struct {
	cake CakeMaker
}

type Cupcake struct {
	cake CakeMaker
}

func (cake Cupcake) Name() string {
	return "🧁"
}
func (cake Cupcake) Price() float64 {
	return 1.0
}

type Cookie struct {
	cake CakeMaker
}

func (cookie Cookie) Name() string {
	return "🍪"
}
func (cookie Cookie) Price() float64 {
	return 2.0
}

type Chocolate struct {
	cake CakeMaker
}

func addToppingName(emot string, name string) string {
	if strings.Contains(name, "with") {
		name += " and " + emot
	} else {
		name += " with " + emot
	}
	return name
}

func (choco Chocolate) Name() string {
	return addToppingName("🍫", choco.cake.Name())
}
func (choco Chocolate) Price() float64 {
	return choco.cake.Price() + 0.1
}

type Nuts struct {
	cake CakeMaker
}

func (nut Nuts) Name() string {
	return addToppingName("🥜", nut.cake.Name())
}

func (nut Nuts) Price() float64 {
	return nut.cake.Price() + 0.2
}
