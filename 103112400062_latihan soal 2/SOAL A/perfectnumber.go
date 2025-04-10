// M. DAVI ILYAS RENALDO/103112400062 
package main

import "fmt"

func isPerfectNumber(n int) bool {
	if n <= 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i*i != n {
				sum += n / i
			}
		}
	}
	return sum == n
}

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	found := false
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada")
	} else {
		fmt.Println()
	}
}