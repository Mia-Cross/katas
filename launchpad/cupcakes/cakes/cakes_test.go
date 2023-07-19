package cakes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameCupcakeShouldReturnName(t *testing.T) {
	my_cake := Cupcake{}
	assert.Equal(t, "🧁", my_cake.Name())
}
func TestAddingChocolateShouldReturnCupcakeWithChocolate(t *testing.T) {
	my_cake := Chocolate{Cupcake{}}
	assert.Equal(t, "🧁 with 🍫", my_cake.Name())
}

func TestAddingChocolateAndNutsShouldReturnCupcakeWithChocoAndNuts(t *testing.T) {
	my_cake := Nuts{Chocolate{Cupcake{}}}
	assert.Equal(t, "🧁 with 🍫 and 🥜", my_cake.Name())
}

func TestNameCookieShouldReturnName(t *testing.T) {
	my_cook := Cookie{}
	assert.Equal(t, "🍪", my_cook.Name())
}

func TestAddingNutsandChocoShouldReturnCupcakeWithNutsAndChoco(t *testing.T) {
	my_cake := Chocolate{Nuts{Cupcake{}}}
	assert.Equal(t, "🧁 with 🥜 and 🍫", my_cake.Name())
}

func TestCupcakePrice(t *testing.T) {
	my_cake := Cupcake{}
	assert.Equal(t, 1.0, my_cake.Price())
}
func TestCupcakeWithChocolatePriceIsGood(t *testing.T) {
	my_cake := Chocolate{Cupcake{}}
	assert.Equal(t, 1.1, my_cake.Price())
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-9
}

func TestCookieChocolateNutsPriceIsGood(t *testing.T) {
	my_cake := Nuts{Chocolate{Cookie{}}}
	assert.Equal(t, true, almostEqual(2.3, my_cake.Price()))
}

func TestBundleCupcakeCookieHasGoodPrice(t *testing.T) {
	my_bundle := Bundle{[]CakeMaker{Cupcake{}, Cookie{}}}
	assert.Equal(t, 3.0, my_bundle.Price())
}
