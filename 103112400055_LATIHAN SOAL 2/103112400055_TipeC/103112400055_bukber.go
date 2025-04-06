package main

import "fmt"

func main() {
	var n_103112400055 int
	fmt.Print("Masukkan jumlah peserta: ")
	fmt.Scan(&n_103112400055)

	countA, countB, countC := 0, 0, 0

	for i := 1; i <= n_103112400055; i++ {
		var number int
		fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", i)
		fmt.Scan(&number)

		numStr := fmt.Sprintf("%d", number)
		if allSame(numStr) {
			fmt.Println("Hadiah A")
			countA++
		} else if allDifferent(numStr) {
			fmt.Println("Hadiah B")
			countB++
		} else {
			fmt.Println("Hadiah C")
			countC++
		}
	}

	fmt.Println("Jumlah yang memperoleh Hadiah A:", countA)
	fmt.Println("Jumlah yang memperoleh Hadiah B:", countB)
	fmt.Println("Jumlah yang memperoleh Hadiah C:", countC)
}

func allSame(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func allDifferent(s string) bool {
	seen := make(map[rune]bool)
	for _, ch := range s {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}
