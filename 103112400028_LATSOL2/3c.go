// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

func main() {
	bilangan1, bilangan2 := bacaInput()
	hasil := hitungPerkalian(bilangan1, bilangan2)
	tampilkanHasil(bilangan1, bilangan2, hasil)
}

func bacaInput() (int, int) {
	var a, b int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&b)
	return a, b
}

func hitungPerkalian(a, b int) int {
	switch {
	case a == 0 || b == 0:
		return 0
	case a == 1:
		return b
	case b == 1:
		return a
	case a < 0:
		return -hitungPerkalian(-a, b)
	default:
		return b + hitungPerkalian(a-1, b)
	}
}

func tampilkanHasil(a, b, hasil int) {
	fmt.Printf("Hasil dari %d x %d = %d\n", a, b, hasil)
}
