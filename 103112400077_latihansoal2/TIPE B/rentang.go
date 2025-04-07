package main

import "fmt"

func ganjilRentang(a, b int) int {
	bilB := (b + 1) / 2 
	bilA := (a - 1) / 2 
	return bilB - bilA
}
func main() {
	var a_103112400077, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a_103112400077)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Banyaknya angka ganjil: ", ganjilRentang(a_103112400077, b))
}