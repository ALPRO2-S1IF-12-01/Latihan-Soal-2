package main

import "fmt"

func Hadiah(d1, d2, d3 int) string {
    if d1 == d2 && d2 == d3 {
        return "Hadiah A"
    } else if d1 != d2 && d2 != d3 && d3 != d1 {
        return "Hadiah B"
    } else {
		return "Hadiah C"
	}
}

func main() {
    var peserta_103112400041, jumlahPeserta, n, d1, d2, d3 int
    var jumlahA, jumlahB, jumlahC int = 0, 0, 0
    fmt.Print("Masukkan jumlah peserta: ")
    fmt.Scanln(&peserta_103112400041)
    jumlahPeserta = 0
    for i := 1; i <= peserta_103112400041; i++ {
        jumlahPeserta++
        fmt.Printf("Masukkan nomor kartu peserta ke-%d: ", jumlahPeserta)
        fmt.Scanln(&n)
        d1 = (n / 100)
        d2 = (n / 10) % 10
        d3 = n % 10
        fmt.Println(Hadiah(d1, d2, d3))
        switch Hadiah(d1, d2, d3) {
        case "Hadiah A":
            jumlahA=jumlahA+1
        case "Hadiah B":
            jumlahB=jumlahB+1
        case "Hadiah C":
            jumlahC=jumlahC+1
        }
    }
    fmt.Println()
    fmt.Println("Jumlah yang memperoleh Hadiah A: ", jumlahA)
    fmt.Println("Jumlah yang memperoleh Hadiah B: ", jumlahB)
    fmt.Println("Jumlah yang memperoleh Hadiah C: ", jumlahC)
}
