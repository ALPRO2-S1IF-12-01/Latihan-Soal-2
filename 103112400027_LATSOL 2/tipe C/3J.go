package main

import "fmt"

func hasilPerkalian(n, m int) int {
	if m == 0 {
		return 0
	}
	return n + hasilPerkalian(n, m-1) 
}

func main() {
	fmt.Println("Nama: Raja Muhammad Lufhti\nNim : 103112400027")
	var n, m int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan nilai m: ")
	fmt.Scan(&m)

	fmt.Printf("Hasil perkalian %d * %d = %d", n, m, hasilPerkalian(n, m))
}
