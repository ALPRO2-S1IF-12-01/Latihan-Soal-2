//Pratama Bintang Daniswara 103112400051
package main

import "fmt"

func main() {
	var a, b int
	var jml int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	jml = 0
	i := a
	for i <= b {
		if i%2 == 1 {
			jml = jml + 1
		}
		i++
	}

	fmt.Println("Banyaknya angka ganjil:", jml)
}
