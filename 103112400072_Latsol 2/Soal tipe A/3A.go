package main

import "fmt"

func main() {
	
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	
	meetings := 0
	for day := 1; day <= 365; day++ {
		
		if day%x == 0 && day%y != 0 {
			meetings++
		}
	}

	
	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", meetings)
}