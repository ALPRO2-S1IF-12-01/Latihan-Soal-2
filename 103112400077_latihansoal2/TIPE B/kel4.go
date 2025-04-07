package main

import "fmt"

func jumlahKelipatan4(total int) int {		
	var angka_103112400077 int
	fmt.Scan(&angka_103112400077)

	if angka_103112400077 < 0 {
		return total 
	}

	if angka_103112400077%4 == 0 {
		return jumlahKelipatan4(total + angka_103112400077) 
	}

	return jumlahKelipatan4(total) 
}

func main() {
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	hasil := jumlahKelipatan4(0) 
	fmt.Println("Jumlah bilangan kelipatan 4:", hasil)
}
