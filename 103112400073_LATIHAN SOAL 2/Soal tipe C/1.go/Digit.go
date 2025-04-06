package main

import (
    "fmt"
    "strconv"
)

func main() {
    var bilangan_103112400073 int

    
    fmt.Print("Masukkan bilangan bulat positif (>10): ")
    fmt.Scan(&bilangan_103112400073)

    
    bilanganStr := strconv.Itoa(bilangan_103112400073)
    panjang := len(bilanganStr)

    
    tengah := panjang / 2

    
    var kiriStr, kananStr string
    if panjang%2 == 0 {
       
        kiriStr = bilanganStr[:tengah]
        kananStr = bilanganStr[tengah:]
    } else {
        
        kiriStr = bilanganStr[:tengah+1]
        kananStr = bilanganStr[tengah+1:]
    }

   
    kiri, _ := strconv.Atoi(kiriStr)
    kanan, _ := strconv.Atoi(kananStr)

    
    hasilPenjumlahan := kiri + kanan

   
    fmt.Printf("Bilangan 1: %d\n", kiri)
    fmt.Printf("Bilangan 2: %d\n", kanan)
    fmt.Printf("Hasil penjumlahan: %d\n", hasilPenjumlahan)
}