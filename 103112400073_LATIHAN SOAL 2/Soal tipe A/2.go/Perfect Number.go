package main

import "fmt"


func PerfectNumber(n int) bool {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func main() {
	var a_103112400073, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a_103112400073)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)


	fmt.Printf("Perfect numbers antara %d dan %d: ", a_103112400073, b)
	for i := a_103112400073; i <= b; i++ {
		if PerfectNumber(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
