// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func cekPerfectNumber(x int) bool {
	jumlah := 0
	for i := 1; i < x; i++ {
		if x % i == 0 {
			jumlah += i
		}
	}
	return jumlah == x
}

func cetakPerfectNumber(a, b int) {
	fmt.Printf("Perfect numbers antara %v dan %v: ", a, b)
	for i := a; i <= b; i++ {
		if cekPerfectNumber(i) {
			fmt.Printf("%v ", i)
		}
	}
}

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	cetakPerfectNumber(a, b)
}