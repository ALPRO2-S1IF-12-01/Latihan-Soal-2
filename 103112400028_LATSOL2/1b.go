// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import "fmt"

func main() {
	var start, end int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&start)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&end)

	fmt.Printf("Perfect numbers antara %d dan %d: ", start, end)

	hasPerfect := printPerfectNumbers(start, end)

	if !hasPerfect {
		fmt.Print("Tidak ada")
	}
	fmt.Println()
}

func printPerfectNumbers(start, end int) bool {
	found := false
	for n := start; n <= end; n++ {
		if checkPerfect(n) {
			fmt.Printf("%d ", n)
			found = true
		}
	}
	return found
}
func checkPerfect(n int) bool {
	if n < 2 {
		return false
	}
	total := 0
	for d := 1; d <= n/2; d++ {
		if n%d == 0 {
			total += d
		}
	}
	return total == n
}
