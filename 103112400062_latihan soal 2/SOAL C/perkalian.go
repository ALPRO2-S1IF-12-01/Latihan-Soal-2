// M. DAVI ILYAS RENALDO/103112400062 
package main

import "fmt"

func perkalian(a, b int) int {
	if b == 0 {
		return 0
	}
	return a + perkalian(a, b-1)
}

func main() {
	var n, m int
	fmt.Print("masukkan bilangan(n): ")
	fmt.Scan(&n)
	fmt.Print("masukkan bilangan(m): ")
	fmt.Scan(&m)

	hasil := perkalian(n, m)
	fmt.Printf("hasil dari %d x %d = %d\n", n, m, hasil)
}