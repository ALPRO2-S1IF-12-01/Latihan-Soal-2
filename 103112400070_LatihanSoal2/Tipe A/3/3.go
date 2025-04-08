// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("nilai x: ")
	fmt.Scan(&x)
	fmt.Print("nilai y: ")
	fmt.Scan(&y)
	pertemuan := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x != 0 && hari%y != 0 {
			pertemuan++
		}
	}
	fmt.Printf("Pertemuan rahasia dalam setahun: %d\n", pertemuan)
}