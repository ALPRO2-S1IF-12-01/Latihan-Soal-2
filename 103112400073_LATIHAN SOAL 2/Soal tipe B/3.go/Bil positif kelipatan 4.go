package main

import "fmt"


func jumlahKelipatan4(total int) int {
    var bilangan_103112400073 int
    fmt.Print("Masukkan bilangan (negatif untuk berhenti): ")
    fmt.Scan(&bilangan_103112400073)
    
    if bilangan_103112400073 < 0 {
        return total
    }
   
    if bilangan_103112400073 > 0 && bilangan_103112400073%4 == 0 {
        total += bilangan_103112400073
    }

    return jumlahKelipatan4(total)
}

func main() {
    total := jumlahKelipatan4(0) 
    fmt.Printf("Jumlah bilangan kelipatan 4 : %d\n", total)
}