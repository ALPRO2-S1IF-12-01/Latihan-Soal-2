//Anastasia Adida N.I
package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Perfect numbers antara %d dan %d:", a, b)
	for i := a; i <= b; i++ {
		if perfectNumber_103112400085(i) {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println()
}

func perfectNumber_103112400085(x int) bool {
	jumlah := 0
	for i := 1; i < x; i++ {
		if x%i == 0 {
			jumlah += i
		}
	}
	return jumlah == x
}
