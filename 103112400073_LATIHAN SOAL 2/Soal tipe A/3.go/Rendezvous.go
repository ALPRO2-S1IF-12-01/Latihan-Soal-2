package main

import "fmt"


func HitungKetemu(x, y int) int {
    hitung := 0
    for day := 1; day <= 365; day++ {
        
        if day%x == 0 && day%y != 0 {
            hitung++
        }
    }
    return hitung
}

func main() {
    var x_103112400073, y int
  
    fmt.Print("Masukkan nilai x: ")
    fmt.Scan(&x_103112400073)
    fmt.Print("Masukkan nilai y: ")
    fmt.Scan(&y)
 
    PertemuanRahasia := HitungKetemu(x_103112400073, y)

    fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", PertemuanRahasia)
}