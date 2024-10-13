package main

import "fmt"

func change(total_cart int, payment int) ([]string, bool) {
	balance := payment - total_cart
	var result []string
	var keyword string

	if total_cart > payment {
		return nil, false
	}

	if balance >= 100000 {
		helper := balance / 100000

		balance = balance - (helper * 100000)

		keyword = fmt.Sprintf("%d lembar 100000", helper)
		result = append(result, keyword)
	}

	if balance >= 50000 {
		helper := balance / 50000

		balance = balance - (helper * 50000)

		keyword = fmt.Sprintf("%d lembar 50000", helper)
		result = append(result, keyword)
	}

	if balance >= 20000 {
		helper := balance / 20000

		balance = balance - (helper * 20000)

		keyword = fmt.Sprintf("%d lembar 20000", helper)
		result = append(result, keyword)
	}

	if balance >= 10000 {
		helper := balance / 10000

		balance = balance - (helper * 10000)

		keyword = fmt.Sprintf("%d lembar 10000", helper)
		result = append(result, keyword)
	}

	if balance >= 5000 {
		helper := balance / 5000

		balance = balance - (helper * 5000)

		keyword = fmt.Sprintf("%d lembar 5000", helper)
		result = append(result, keyword)
	}

	if balance >= 2000 {
		helper := balance / 2000

		balance = balance - (helper * 2000)

		keyword = fmt.Sprintf("%d lembar 2000", helper)
		result = append(result, keyword)
	}

	if balance >= 1000 {
		helper := balance / 1000

		balance = balance - (helper * 1000)

		keyword = fmt.Sprintf("%d lembar 1000", helper)
		result = append(result, keyword)
	}

	if balance >= 500 {
		helper := balance / 500

		balance = balance - (helper * 500)

		keyword = fmt.Sprintf("%d koin 500", helper)
		result = append(result, keyword)
	}

	if balance >= 200 {
		helper := balance / 200

		balance = balance - (helper * 200)

		keyword = fmt.Sprintf("%d koin 200", helper)
		result = append(result, keyword)
	}

	if balance >= 100 {
		helper := (balance / 100)

		keyword = fmt.Sprintf("%d koin 100", helper)
		result = append(result, keyword)
	}

	return result, true
}

func main() {
	total_cart := 657650
	payment := 600000

	res, status := change(total_cart, payment)

	if status {
		fmt.Printf("%v", res)
	} else {
		fmt.Print("kurang bayar")
	}

	total_cart1 := 700649
	payment1 := 800000

	res1, status1 := change(total_cart1, payment1)

	if status1 {
		fmt.Printf("%v", res1)
	} else {
		fmt.Print("kurang bayar")
	}
}
