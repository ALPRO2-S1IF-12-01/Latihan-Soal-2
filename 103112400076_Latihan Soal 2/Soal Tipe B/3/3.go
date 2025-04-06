package main

import (
	"fmt"
)
func inputDanHitung(total int) int {
	var angka int
	fmt.Scan(&angka)
	if angka < 0 {
		return total
	}
	fmt.Printf("%d ", angka)
	if angka%4 == 0 {
		total += angka
	}
	return inputDanHitung(total)
}
func main() {
	NIM := "103112400076" 
	fmt.Println("NIM:", NIM) 
	fmt.Print("Masukkan bilangan (negatif untuk berhenti):\n: ")
	total := inputDanHitung(0)
	fmt.Println() 
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}
