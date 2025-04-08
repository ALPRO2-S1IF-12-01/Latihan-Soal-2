package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	fmt.Println("Jumlah pertemuan dalam setahun:", rendezvous(x, y))
}

func rendezvous(x, y int) int {
	var total_103112400061 int
	for i := 1; i <= 365; i++ {
		if i % x == 0 && i % y != 0 {
			total_103112400061++
		}
	}
	return total_103112400061
}