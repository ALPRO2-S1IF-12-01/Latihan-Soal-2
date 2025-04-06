package main

import (
	"fmt"
)
func kali(n, m, hasil int) int {
	if m == 0 {
		return hasil
	}
	return kali(n, m-1, hasil+n)
}
func main() {
	NIM := "103112400076"
	fmt.Println("NIM:", NIM)
	var n, m int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan m: ")
	fmt.Scan(&m)
	hasil := kali(n, m, 0)
	fmt.Printf("Hasil dari %d x %d = %d\n", n, m, hasil)
}
