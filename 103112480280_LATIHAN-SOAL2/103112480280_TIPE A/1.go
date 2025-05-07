// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import (
    "fmt"
    "strings"
)

func durasi(jam_103112480280, menit int) int {
    if jam_103112480280 < 1 {
        return 1
    }
    if menit >= 10 {
        return jam_103112480280 + 1
    }
    return jam_103112480280
}

func biayaSewa(durasi int, member bool) float64 {
    if member {
        return float64(durasi) * 3500
    }
    return float64(durasi) * 5000
}

func diskon(biaya float64, voucher string, durasi int) float64 {
    panjang := len(strings.TrimSpace(voucher))
    if (panjang == 5 || panjang == 6) && durasi > 3 {
        return biaya * 0.9
    }
    return biaya
}

func main() {
    var jam_103112480280, menit int
    var member bool
    var voucher string

    fmt.Print("Masukkan durasi (jam): ")
    fmt.Scan(&jam_103112480280)
    fmt.Print("Masukkan durasi (menit): ")
    fmt.Scan(&menit)
    fmt.Print("Apakah member? (true/false): ")
    fmt.Scan(&member)
    fmt.Print("Masukkan nomor voucher: ")
    fmt.Scan(&voucher)

    totalDurasi := durasi(jam_103112480280, menit)
    totalBiaya := biayaSewa(totalDurasi, member)
    totalSetelahDiskon := diskon(totalBiaya, voucher, totalDurasi)

    fmt.Printf("Biaya sewa setelah diskon: Rp %.2f\n", totalSetelahDiskon)
}