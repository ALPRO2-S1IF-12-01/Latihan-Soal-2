//Pratama Bintang Daniswara 103112400051
package main

import "fmt"

func main() {
    var bilangan int
    
    fmt.Print("Masukkan bilangan bulat positif (>10): ")
    fmt.Scanln(&bilangan)
    
    if bilangan <= 10 {
        fmt.Println("Error: Bilangan harus lebih besar dari 10")
        return
    }
    
    temp := bilangan
    jumlahDigit := 0
    for temp != 0 {
        temp = temp / 10
        jumlahDigit++
    }
    
    potong := jumlahDigit / 2
    if jumlahDigit % 2 != 0 {
        potong++
    }
    
    pembagi := 1
    for i := 0; i < jumlahDigit - potong; i++ {
        pembagi = pembagi * 10
    }
    
    bilangan1 := bilangan / pembagi
    bilangan2 := bilangan % pembagi
    total := bilangan1 + bilangan2

    fmt.Printf("Bilangan 1: %d\n", bilangan1)
    fmt.Printf("Bilangan 2: %d\n", bilangan2)
    fmt.Printf("Hasil penjumlahan: %d\n", total)
}