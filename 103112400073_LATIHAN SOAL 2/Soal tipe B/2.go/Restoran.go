package main

import "fmt"


func hitungBiaya(jumlahMenu, jumlahOrang int, adaSisa bool) int {
    var biayaPerMenu int
    var totalBiaya_103112400073 int

    
    if jumlahMenu <= 3 {
        biayaPerMenu = 10000
    } else if jumlahMenu > 50 {
        biayaPerMenu = 100000
    } else {
        biayaPerMenu = 10000 + (2500 * (jumlahMenu - 3))
    }

    
    totalBiaya_103112400073 = biayaPerMenu

   
    if adaSisa {
        totalBiaya_103112400073 *= jumlahOrang
    }

    return totalBiaya_103112400073
}

func main() {
    var jumlahRombongan int

    
    fmt.Print("Masukkan jumlah rombongan: ")
    fmt.Scan(&jumlahRombongan)

    
    for i := 1; i <= jumlahRombongan; i++ {
        var jumlahMenu, jumlahOrang int
        var sisaMakanan int 

        
        fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
        fmt.Scan(&jumlahMenu, &jumlahOrang, &sisaMakanan)

        
        adaSisa := sisaMakanan == 1

        
        totalBiaya := hitungBiaya(jumlahMenu, jumlahOrang, adaSisa)

        
        fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, totalBiaya)
    }
}