package main

import "fmt"

func perfek(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n && n != 1
}

func main() {
	var a_103112400077, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a_103112400077)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Perfect numbers antara %d dan %d:", a_103112400077, b)
	first := true

	for i := a_103112400077; i <= b; i++ {
		if perfek(i) {
			if !first {
				fmt.Print(",") 
			}
			fmt.Printf(" %d", i)
			first = false
		}
	}

	if first {
		fmt.Print(" Tidak ada")
	}
	fmt.Println()
}
