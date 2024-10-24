package main

import (
	"fmt"
	"strings"
)

const maxTime = 301

func hitungSkor(waktu [8]int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, waktuSoal := range waktu {
		if waktuSoal < maxTime {
			*soal++
			*skor += waktuSoal
		}
	}
}

func main() {
	var pemenang string
	var soalPemenang, skorPemenang int
	soalPemenang = 0
	skorPemenang = maxTime * 8

	for {
		var nama string
		var waktu [8]int

		fmt.Print("Masukkan nama peserta (atau 'Selesai' untuk mengakhiri) : ")
		fmt.Scan(&nama)
		nama = strings.TrimSpace(nama)
		if nama == "Selesai" {
			break
		}
		fmt.Print("Masukkan waktu penyelesaian 8 soal (dalam menit) : ")
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}
		var soal, skor int
		hitungSkor(waktu, &soal, &skor)
		if soal > soalPemenang || (soal == soalPemenang && skor < skorPemenang) {
			pemenang = nama
			soalPemenang = soal
			skorPemenang = skor
		}
	}
	fmt.Printf("Pemenang : %s, Soal yang diselesaikan :%d, Total waktu : %d menit\n", pemenang, soalPemenang, skorPemenang)
}
