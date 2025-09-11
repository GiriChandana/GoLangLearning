package main

func discountedPrice(product string, price float64) float64 {
	var discount float64
	switch product {
	case "apples":
		discount = 0.1
	case "bananas":
		discount = 0.2
	default:
		discount = 0
	}
	return price * (1 - discount)
}
