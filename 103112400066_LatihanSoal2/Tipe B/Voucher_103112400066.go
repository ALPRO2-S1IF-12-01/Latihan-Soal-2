// DWI OKTA SURYANINGRUM / 103112400066

package main

import "fmt"

func main() {
	var a, b int

	// Ambil input nilai batas awal dan akhir
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	// Cek validasi input
	if a < 0 || b < 0 {
		fmt.Println("Error: Masukan harus bilangan bulat positif")
		return
	}
	if a > b {
		fmt.Println("Error: Nilai a harus lebih kecil atau sama dengan b")
		return
	}

	// Hitung jumlah bilangan ganjil di antara a dan b
	jumlahGanjil := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlahGanjil++
		}
	}

	fmt.Printf("Banyaknya angka ganjil: %d\n", jumlahGanjil)
}
