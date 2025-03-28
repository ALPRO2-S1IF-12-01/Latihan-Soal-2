package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	if x <= 0 || y <= 0 {
		fmt.Println("Error: Nilai x dan y harus bilangan bulat positif")
		return
	}

	// menghitung jumlah pertemuan langsung dari kelipatan x
	jumlahPertemuan := 0
	for hari := x; hari <= 365; hari += x {
		if hari%y != 0 {
			jumlahPertemuan++
		}
	}

	fmt.Println("Jumlah pertemuan dalam setahun:", jumlahPertemuan)
}
