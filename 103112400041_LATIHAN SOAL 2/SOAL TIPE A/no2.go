package main

import "fmt"

func Perfect(a, b_103112400041 int) {
	var perfect int
	for i := a; i <= b_103112400041; i++ {
		perfect = 0
		for j := 1; j < i; j++ {
			if i % j == 0 {
				perfect=perfect+j
			}
		}
		if perfect == i {
			fmt.Println(i)
		}
	}
}
func main() {
	var a,b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	Perfect(a, b)
}