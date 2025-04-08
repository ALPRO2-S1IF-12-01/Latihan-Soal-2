package main

import "fmt"

func bilanganSempurna(n int) bool {
	jumlah := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			jumlah += i
		}
	}
	return jumlah == n
}

func tampilkanBilanganSempurna(a, b int) {
	fmt.Printf("Bilangan sempurna antara %d dan %d: ", a, b)
	ada := false
	for i := a; i <= b; i++ {
		if bilanganSempurna(i) {
			fmt.Print(i, " ")
			ada = true
		}
	}
	if !ada {
		fmt.Print("Tidak ada")
	}
	fmt.Println()
}

func main() {
	fmt.Println("Nama: Raja Muhammad Lufhti\nNim : 103112400027")
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Nilai a harus lebih kecil atau sama dengan b.")
		return
	}

	tampilkanBilanganSempurna(a, b)
}
