// DWI OKTA SURYANINGRUM / 103112400066

package main

import "fmt"

func main() {
    var bilangan int
    
    fmt.Print("Masukkan bilangan bulat positif (>10): ")
    fmt.Scan(&bilangan)
    
    // Validasi input
    if bilangan <= 10 {
        fmt.Println("Error: Bilangan harus lebih besar dari 10")
        return
    }
    
    // Hitung jumlah digit
    temp := bilangan
    digitCount := 0
    for temp > 0 {
        temp /= 10
        digitCount++
    }
    
    // Tentukan titik potong
    splitPos := digitCount / 2
    if digitCount%2 != 0 {
        splitPos++
    }
    
    // Hitung pembagi untuk memisahkan bilangan
    divisor := 1
    for i := 0; i < digitCount-splitPos; i++ {
        divisor *= 10
    }
    
    // Pisahkan bilangan
    bagian1 := bilangan / divisor
    bagian2 := bilangan % divisor
    
    // Hitung dan tampilkan hasil
    fmt.Printf("Bilangan 1: %d\n", bagian1)
    fmt.Printf("Bilangan 2: %d\n", bagian2)
    fmt.Printf("Hasil penjumlahan: %d\n", bagian1 + bagian2)
}