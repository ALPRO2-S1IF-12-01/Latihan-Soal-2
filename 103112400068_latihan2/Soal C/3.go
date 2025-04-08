package main

import "fmt"

func multiply(n, m int) int {
	if m == 0 {
		return 0
	}
	return n + multiply(n, m-1)
}

func main() {
	var n, m int

	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)

	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	result := multiply(n, m)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, result)
}
