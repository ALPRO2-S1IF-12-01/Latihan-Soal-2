package main

import "fmt"

func main() {
	var n, r int
	fmt.Print("Masukkan Nilai A: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan Nilai B: ")
	fmt.Scan(&r)

	fmt.Printf("Perfect numbers antara %d dan %d: ", n, r)
	found := false

	for i := n; i <= r; i++ {
		if isPerfect(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}

	if !found {
		fmt.Print("tidak ada")
	}
}

func isPerfect(angka int) bool {
	if angka <= 1 {
		return false
	}

	total := 0
	for i := 1; i < angka; i++ {
		if angka%i == 0 {
			total += i
		}
	}
	return total == angka
}
