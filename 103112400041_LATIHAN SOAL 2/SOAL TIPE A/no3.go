package main

import "fmt"

func secretMeeting(x, y_103112400041 int) int {
	jumlahHari := 0
	for i := 1; i <= 365; i++ {
		if i % x == 0 && i % y_103112400041 != 0 {
			jumlahHari=jumlahHari+1
		}
	}
	return jumlahHari
}

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)
	fmt.Print("Jumlah pertemuan dalam setahun: ", secretMeeting(x, y))
}