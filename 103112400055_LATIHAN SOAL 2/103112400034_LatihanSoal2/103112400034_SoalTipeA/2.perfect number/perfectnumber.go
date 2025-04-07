package main

import "fmt"

func adalahPerfect(n int) bool {
	jumlah := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			jumlah += i
		}
	}
	return jumlah == n
}

func main() {
	var a, b int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	ada := false
	for i := a; i <= b; i++ {
		if adalahPerfect(i) {
			fmt.Print(i, " ")
			ada = true
		}
	}

	if !ada {
		fmt.Print("Tidak ada")
	}

	fmt.Println()
}
