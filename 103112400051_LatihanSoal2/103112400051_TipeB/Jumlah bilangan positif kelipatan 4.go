//Pratama Bintang Daniswara 103112400051
package main

import "fmt"

func main() {
	var x, t int

	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")
	for {
		fmt.Scan(&x)
		switch {
		case x < 0:
			break
		case x > 0 && x%4 == 0:
			t = t + x
		}
		if x < 0 {
			break
		}
	}
	fmt.Println("Jumlah bilangan kelipatan 4:", t)
}
