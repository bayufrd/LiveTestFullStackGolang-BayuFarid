package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

func (p Product) GetPriceWithTax(taxRate float64) float64 {
	return p.Price * (1 + taxRate)
}

func CalculateAverage(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("cannot calculate average of an empty slice")
	}

	total := 0
	for _, number := range numbers {
		total += number
	}

	return float64(total) / float64(len(numbers)), nil
}

func main() {
	product := Product{
		Name:  "Laptop",
		Price: 1200.0,
	}

	fmt.Printf("%s price with tax: %.2f\n", product.Name, product.GetPriceWithTax(0.07))

	numbers := []int{10, 20, 30, 40}
	average, err := CalculateAverage(numbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Average: %.2f\n", average)
	}

	emptyNumbers := []int{}
	emptyAverage, err := CalculateAverage(emptyNumbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Average: %.2f\n", emptyAverage)
	}

	inventory := map[string]int{
		"apple":  10,
		"banana": 5,
		"orange": 0,
	}

	for productName, quantity := range inventory {
		if quantity == 0 {
			fmt.Printf("%s is out of stock.\n", productName)
			continue
		}

		fmt.Printf("%s has %d items.\n", productName, quantity)
	}
}
