package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

func GetPriceWithTax(product Product, taxRate float64) Product {

	tax := float64((taxRate * product.Price) / 100)

	product.Price = tax
	
	return product
}

func CalculateAverage(num []int) (float64, error) {

	total := len(num)
	sum := 0

	if total == 0 {
		return 0, errors.New("Empty list")
	}

	for _, j := range num {
		sum = sum + j
	}

	return float64(sum / total), nil
}

func main() {
	products := Product{
		Name:  "Laptop",
		Price: 12000,
	}

	fmt.Println(GetPriceWithTax(products, 7))

	nums := []int{10, 20, 30, 40}
	empties := []int{}

	fmt.Println(CalculateAverage(nums))
	fmt.Println(CalculateAverage(empties))

	inventory := map[string]int{
		"apple":  10,
		"banana": 5,
		"orange": 0,
	}

	for i, j := range inventory {
		if j == 0 {
			fmt.Printf("%s is out of stock \n", i)
		} else {
			fmt.Printf("%s has %d items. \n", i, j)
		}
	}
}
