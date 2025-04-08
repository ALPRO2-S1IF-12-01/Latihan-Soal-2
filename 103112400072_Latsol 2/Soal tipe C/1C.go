package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan int
	
	
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&bilangan)
	
	
	panjang := hitungPanjangDigit(bilangan)
	
	
	var titikPotong int
	if panjang%2 == 0 {
		
		titikPotong = panjang / 2
	} else {
		
		titikPotong = (panjang + 1) / 2
	}
	
	
	pembagi := int(math.Pow10(panjang - titikPotong))
	
	
	bilangan1 := bilangan / pembagi
	bilangan2 := bilangan % pembagi
	
	
	hasilPenjumlahan := bilangan1 + bilangan2
	
	
	fmt.Printf("Bilangan 1: %d\n", bilangan1)
	fmt.Printf("Bilangan 2: %d\n", bilangan2)
	fmt.Printf("Hasil penjumlahan: %d\n", hasilPenjumlahan)
}


func hitungPanjangDigit(n int) int {
	if n == 0 {
		return 1
	}
	
	panjang := 0
	for n > 0 {
		panjang++
		n /= 10
	}
	
	return panjang
}