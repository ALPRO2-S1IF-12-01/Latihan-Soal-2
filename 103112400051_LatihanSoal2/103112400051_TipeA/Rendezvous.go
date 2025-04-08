//Pratama Bintang Daniswara 103112400051
package main

import "fmt"


func main() {
	var x, y int
	fmt.Print("Masukan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukan nilai y: ")
	fmt.Scan(&y)

	switch {
	case x <= 0 || y <= 0:
		return
	}
	h := 0
	for i := 1; i <= 365; i++ {
		if i%x == 0 && i%y != 0 {
			h++
		}
	}
	fmt.Println("Jumlah pertemuan dalam setahun:", h)
}
