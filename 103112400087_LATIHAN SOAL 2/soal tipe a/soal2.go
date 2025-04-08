package main

import "fmt"

func isperfect(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func main() {
	var a, b_10312400087 int

	fmt.Print("masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b: ")
	fmt.Scan(&b_10312400087)

	fmt.Printf("perfect numbers antara %d dan %d: ", a, b_10312400087)
	for i := a; i <= b_10312400087; i++ {
		if isperfect(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
