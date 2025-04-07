package main

import "fmt"

func jumlahPertemuan(a, b int) int {
	jumlahHari := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%a == 0 && hari%b != 0 {
			jumlahHari++
		}
	}
	return jumlahHari
}

func main() {
	var x, y int

	fmt.Print("masukkan nilai x : ")
	fmt.Scan(&x)

	fmt.Print("masukkan nilai y : ")
	fmt.Scan(&y)

	totalPertemuan := jumlahPertemuan(x, y)

	fmt.Printf("jumlah pertemuan dalam setahun: %d hari\n", totalPertemuan)
	fmt.Print("Sheila Stephanie A [103112400086]")
}
