//Anastasia Adinda N.I
package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan_103112400085 int

	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&bilangan_103112400085)

	jumlahDigit := int(math.Log10(float64(bilangan_103112400085))) + 1
	tengah := jumlahDigit / 2
	if jumlahDigit%2 != 0 {
		tengah++
	}

	pangkat := int(math.Pow(10, float64(jumlahDigit-tengah)))

	bilangan1 := bilangan_103112400085 / pangkat
	bilangan2 := bilangan_103112400085 % pangkat

	fmt.Println("Bilangan 1:", bilangan1)
	fmt.Println("Bilangan 2:", bilangan2)
	fmt.Println("Hasil penjumlahan:", bilangan1+bilangan2)
}
