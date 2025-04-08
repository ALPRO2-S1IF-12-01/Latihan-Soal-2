// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Nilai b: ")
	fmt.Scan(&b)

	jmlganjil := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jmlganjil++
		}
	}
	fmt.Printf("Banyaknya angka ganjil: %d\n", jmlganjil)
}