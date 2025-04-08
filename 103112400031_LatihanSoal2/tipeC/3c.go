// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func perkalianRekursif(n, m, hasilPerkalian int) int {
	if m == 0 {
		return hasilPerkalian
	}
	return perkalianRekursif(n, m-1, hasilPerkalian+n)
}

func main() {
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)

	hasilPerkalian := perkalianRekursif(n, m, 0)

	fmt.Printf("hasilPerkalian dari %v x %v = %v", n, m, hasilPerkalian)
}
