package main

import (
    "fmt"
    "strconv"
)

func Dipisah(n int) (int, int) {
    jadiString := strconv.Itoa(n)
    jumlahDigit := len(jadiString)
    tengah := jumlahDigit / 2
    if jumlahDigit%2 != 0 {
        tengah += 1
    }
    bilangan1, _ := strconv.Atoi(jadiString[:tengah])
    bilangan2, _ := strconv.Atoi(jadiString[tengah:])

    return bilangan1, bilangan2
}

func main() {
    var n_103112400041 int
    fmt.Print("Masukkan bilangan bulat positif (>10): ")
    fmt.Scanln(&n_103112400041)
    bil1, bil2 := Dipisah(n_103112400041)
    fmt.Printf("Bilangan 1: %d\n", bil1)
    fmt.Printf("Bilangan 2: %d\n", bil2)
    fmt.Print("Hasil Penjumlahan: ",bil1 + bil2)
}