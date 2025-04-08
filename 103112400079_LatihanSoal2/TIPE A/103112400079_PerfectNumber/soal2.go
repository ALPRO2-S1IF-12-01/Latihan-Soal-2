package main

import "fmt"

func jumlahFaktor(n int) int {
	jumlah := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			jumlah += i
			if i*i != n {
				jumlah += n / i
			}
		}
	}
	return jumlah
}

func main() {
	var a, b_103112400079 int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b_103112400079)

	fmt.Printf("Bilangan sempurna antara %d dan %d: ", a, b_103112400079)
	ada := false
	for i := a; i <= b_103112400079; i++ {
		if jumlahFaktor(i) == i && i > 1 {
			if ada {
				fmt.Print(", ")
			}
			fmt.Print(i)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Tidak ada")
	} else {
		fmt.Println()
	}
}
