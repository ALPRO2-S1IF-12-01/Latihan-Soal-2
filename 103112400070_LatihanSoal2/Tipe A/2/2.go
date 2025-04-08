// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func perfectnumber(n int) bool {
	if n <= 1 {
		return false
	}
	s := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			s += i
			if i*i != n {
				s += n / i
			}
		}
	}
	return s == n
}

func main() {
	var a, b int
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Printf("Bilangan sempurna antara %d dan %d: ", a, b)
	bil := false
	for i := a; i <= b; i++ {
		if perfectnumber(i) {
			if bil {
				fmt.Print(", ")
			}
			fmt.Print(i)
			bil = true
		}
	}
}