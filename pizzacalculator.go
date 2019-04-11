package main

import (
	"errors"
)

var toppingsAvailable = []string{"Anchovies","Chicken","Bacon","Mushrooms","Olives","Jalapenos","Red Onions","Mixed Peppers","Tomato","Green Chillies"}

func IsToppingAvailable(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func PizzaCalculator(base string, toppings []string) (float64, error) {
	var totalPrice = 0.0
	var flag = 0

	if base == "" {
		return totalPrice,errors.New("You have to choose a base for pizza")
	}

	switch base {
	case "Classic":
		totalPrice = 8.99
		break
	case "Stuffed":
		totalPrice = 11.99
		break
	case "Pan":
		totalPrice = 9.99
		break
	case "Italian":
		totalPrice = 10.99
		break
	default:
		return totalPrice,errors.New("The base you chose is not available")
	}

	if len(toppings) > 0 {
		for _, v := range toppings {
			if !IsToppingAvailable(v, toppingsAvailable) {
				flag = 1
			}
		}

		if flag == 1 {
			return totalPrice, errors.New("Toppings you chose is not available")
		}

		toppingsPrice := float64(len(toppings)) * 1.4
		totalPrice += toppingsPrice
	}
	return totalPrice,nil
}