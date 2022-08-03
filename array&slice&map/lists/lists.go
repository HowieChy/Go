package lists

import "fmt"

func main() {

	//数组 arrays
	// var productNames [4]string = [4]string{"A Book"}
	// prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	//切片
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
