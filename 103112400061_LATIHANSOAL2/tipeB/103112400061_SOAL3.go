package main

import "fmt"

func main() {
	var jumlah_103112400061 int
	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	jumlah_103112400061 = kelipatanEmpat(0)
	fmt.Println("jumlah bilangan kelipatan 4:", jumlah_103112400061) // NIM: 103112400061
}

func kelipatanEmpat(jumlah_103112400061 int) int {
	var x int
	fmt.Scan(&x)
	// AUTHOR: KEISHIN NAUFA ALFARIDZHI
	if x < 0{
		return jumlah_103112400061
	}

	if x%4 == 0 {
		jumlah_103112400061 += x
	}

	return kelipatanEmpat(jumlah_103112400061)
}