package main

import "fmt"

func perkalian(n, m, hasil int) int {
	if m == 0 {
		return hasil
	}
	return perkalian(n, m-1, hasil+n)
}

func main() {
	var n, m_103112400079 int

	fmt.Print("Masukkan bilangan n: ")
	fmt.Scanln(&n)

	fmt.Print("Masukkan bilangan m: ")
	fmt.Scanln(&m_103112400079)

	hasil := perkalian(n, m_103112400079, 0)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m_103112400079, hasil)
}
