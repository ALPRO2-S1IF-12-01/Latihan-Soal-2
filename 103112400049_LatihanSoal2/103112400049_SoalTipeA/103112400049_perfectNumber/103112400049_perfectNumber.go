package main

//HISYAM NURDIATMOKO - 103112400049
import "fmt"

func main() {
	var a_103112400049, b int
	fmt.Print("masukkan nilai a: ")
	fmt.Scan(&a_103112400049)
	fmt.Print("masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("perfect numbers antara ", a_103112400049, " dan ", b, ": ")
	pertama := true
	for n := a_103112400049; n <= b; n++ {
		if perfect(n) {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(n)
			pertama = false
		}
	}
	fmt.Println()
}

func perfect(n int) bool {
	if n <= 1 {
		return false
	}
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}
