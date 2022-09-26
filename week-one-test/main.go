package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name string
	Price float64
}

func printProduct(product Product) {
	fmt.Println(product.Name,"-", product.Price)
}

func filterProduct(in []Product, price float64) []Product {
	var out []Product
	for _, each := range in {
		 if each.Price == price {
				out = append(out, each)
		 }
	}
	return out
}

func main() {
	products := []Product {
		{"Benih Lele", 50000 },
		{"Pakan Lele Cap Menara", 25000 },
		{"Probiotik A", 75000},
		{"Probiotik Nila B", 10000},
		{"Pakan Nila", 20000},
		{"Benih Nila", 20000},
		{"Cupang", 5000},
		{"Benih Nila", 30000},
		{"Bening Cupang", 10000},
		{"Probiotik B", 10000},
	}

	unsortedProducts := append([]Product{}, products...)

	sort.Slice(products[:], func(i, j int) bool {
		return products[i].Price < products[j].Price
	})

	fmt.Println("Total produk dengan harga dibawah Rp 100.000 :")
	var priceTemp float64 = 0
	for i := 0; i < len(products); i++ {
		if priceTemp >= 100000 {
			break
		}

		priceTemp += products[i].Price
		printProduct(products[i])
	}
	fmt.Println("================================================")
	var cheapestPrice float64 = 1000000
	var mostExpensivePrice float64 = 0

	for _, p := range unsortedProducts {
		if p.Price < cheapestPrice {
			cheapestPrice = p.Price
		}

		if p.Price >= mostExpensivePrice {
			mostExpensivePrice = p.Price
		}
	}

	fmt.Println("Daftar produk Termurah: ", cheapestPrice)
	fmt.Println("Daftar produk Termahal: ", mostExpensivePrice)
	fmt.Println("================================================")
	filteredProducts := filterProduct(unsortedProducts, 10000)
	fmt.Println("Daftar dengan harga Rp 10.000: ")
	for _, p := range filteredProducts {
		printProduct(p)
	}
 }