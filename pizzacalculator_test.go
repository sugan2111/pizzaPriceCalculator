package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClassicBaseAndOneToppingReturnsPrice(t *testing.T) {
	toppings := []string{"Anchovies"}
	actual, err := PizzaCalculator("Classic",toppings)
	expected := 10.39

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestClassicBaseNoToppingsReturnsPrice(t *testing.T) {
	actual, err := PizzaCalculator("Classic",[]string{})
	expected := 8.99

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestPanBaseAndThreeToppingsReturnsPrice(t *testing.T) {
	toppings := []string{"Mixed Peppers","Tomato","Green Chillies"}
	actual, err := PizzaCalculator("Pan",toppings)
	expected := 14.19

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestStuffedBaseAndTwoToppingReturnsPrice(t *testing.T) {
	toppings := []string{"Jalapenos","Red Onions"}
	actual, err := PizzaCalculator("Stuffed",toppings)
	expected := 14.79

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestItalianBaseAndOneToppingReturnsPrice(t *testing.T) {
	toppings := []string{"Chicken"}
	actual, err := PizzaCalculator("Italian",toppings)
	expected := 12.39

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestBaseAndThreeToppingsReturnsPrice(t *testing.T) {
	toppings := []string{"Mushrooms","Olives","Jalapenos"}
	actual, err := PizzaCalculator("Pan",toppings)
	expected := 14.19

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestIncorrectBaseReturnsError(t *testing.T) {
	_,err := PizzaCalculator("Random",[]string{})
	assert.Error(t,err,"The base you chose is not available")
}

func TestNoBaseAndOnlyToppingsReturnsError(t *testing.T) {
	_,err := PizzaCalculator("",[]string{"Olives"})
	assert.Error(t,err,"You have to choose a base for pizza")
}

func TestBaseAndIncorrectToppingsReturnsError(t *testing.T) {
	toppings := []string{"Beef,Pork"}
	_,err := PizzaCalculator("Classic",toppings)
	assert.Error(t,err,"The toppings you chose is not available")
}

