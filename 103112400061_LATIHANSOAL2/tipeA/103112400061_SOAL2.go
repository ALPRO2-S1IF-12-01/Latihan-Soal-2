package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	perfectNumber(a, b)
}

func perfectNumber(a, b int) {
	var perfectNumber int
	fmt.Printf("Perfect numbers antara %v dan %v: ", a, b)
	for i := b; i > a; i-- {
		var j int
		var factors_103112400061 int = 0
		// AUTHOR: KEISHIN NAUFA ALFARIDZHI
		for j = 1; j < i; j++ {
			if i%j == 0 {
				// fmt.Print(j, " ")
				factors_103112400061 += j
			}
		}
		if factors_103112400061 == i { // NIM: 103112400061
			perfectNumber = factors_103112400061
			fmt.Printf("%v\n", perfectNumber)
		}
	}
}