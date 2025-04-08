// DWI OKTA SURYANINGRUM / 103112400066

package main

import "fmt"

// Fungsi utama buat menghitung total harga sewa
func hitungTotal(jamSewa, menitSewa int, isMember bool, voucherCode string) float64 {
	durasi := konversiJam(jamSewa, menitSewa) // Ubah jam + menit jadi durasi dalam jam (pakai float)
	var tarifPerJam, totalHarga float64       // Siapin variabel buat tarif dan total biaya

	// Cek apakah user adalah member
	if isMember {
		tarifPerJam = 3500.0 // Member dapet tarif lebih murah
	} else {
		tarifPerJam = 5000.0 // Non-member pakai tarif normal
	}

	totalHarga = durasi * tarifPerJam // Hitung total biaya berdasarkan durasi

	// Kasih diskon 10% kalau durasi lebih dari 3 jam dan voucher valid
	if durasi > 3 && validVoucher(voucherCode) {
		totalHarga = totalHarga * 0.9
	}

	return totalHarga // Balikin total harga ke pemanggil fungsi
}

// Fungsi buat ngubah jam + menit jadi durasi total dalam jam
func konversiJam(jam, menit int) float64 {
	var totalDurasi float64
	if jam < 1 {
		totalDurasi = 1 // Minimal sewa tetap dihitung 1 jam
	} else if jam >= 1 && menit <= 10 {
		totalDurasi = float64(jam) // Kalau cuma nambah menit dikit, nggak dihitung tambahan
	} else if jam >= 1 && menit > 10 {
		totalDurasi = float64(jam) + float64(menit)/60 // Kalau lewat 10 menit, dihitung proporsional
	}
	return totalDurasi
}

// Fungsi buat cek apakah kode voucher valid
func validVoucher(kode string) bool {
	jumlahDigit := len(kode)
	return jumlahDigit == 5 || jumlahDigit == 6 // Voucher dianggap valid kalau panjangnya 5 atau 6 digit
}

// Fungsi utama yang dijalankan pertama kali
func main() {
	var durasiJam, durasiMenit int
	var kodeVoucher string
	var statusMember bool

	// Ambil input dari user
	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&durasiJam)

	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&durasiMenit)

	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&statusMember)

	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&kodeVoucher)

	// Tampilkan hasil akhir biaya sewa
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", hitungTotal(durasiJam, durasiMenit, statusMember, kodeVoucher))
}