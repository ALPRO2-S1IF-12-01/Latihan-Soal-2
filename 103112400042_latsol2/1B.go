// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	found := false
	for num := a; num <= b; num++ {
		if isPerfectNumber(num) {
			fmt.Printf("%d ", num)
			found = true
		}
	}
	if !found {
		fmt.Print("Tidak ada")
	}
	fmt.Println()
}
func isPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if otherFactor := num / i; otherFactor != i {
				sum += otherFactor
			}
		}
	}
	return sum == num
}
