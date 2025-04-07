package main

import "fmt"

func perfectNumber(angka int) bool {
	jumlah := 0
	for i := 1; i <= angka/2; i++ {
		if angka%i == 0 {
			jumlah += i
		}
	}
	return jumlah == angka
}

func main() {
	var a, b int

	fmt.Print("nilai a : ")
	fmt.Scan(&a)
	fmt.Print("nilai b : ")
	fmt.Scan(&b)

	fmt.Printf("perfect number dari %d sampai %d: ", a, b)
	for i := a; i <= b; i++ {
		if perfectNumber(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	fmt.Print("Sheila Stephanie A [103112400086]")
}
