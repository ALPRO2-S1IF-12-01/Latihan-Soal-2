package main

import "fmt"

// rekursif
func jumlahKelipatan(total int) int {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		return total
	}
	if bilangan > 0 && bilangan%4 == 0 {
		total = total + bilangan
	}

	return jumlahKelipatan(total)
}

func main() {
	fmt.Println("masukkan bilangan (negatif untuk berhenti) : ")

	hasil := jumlahKelipatan(0)
	fmt.Println("jumlah bilangan kelipatan 4 : ", hasil)
	fmt.Print("Sheila Stephanie A [103112400086]")
}
