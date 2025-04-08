//Pratama Bintang Daniswara 103112400051
package main

import "fmt"


func isSempurna(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func main() {
	var a, b int

	fmt.Print("Masukan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukan nilai b: ")
	fmt.Scan(&b)

	switch {
	case a <= 0 || b <= 0:
		return
	case a > b:
		return
	default:
		fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
		for i := a; i <= b; i++ {
			if isSempurna(i) {
				fmt.Print(i, " ")
			}
		}
		fmt.Println()
	}
}

