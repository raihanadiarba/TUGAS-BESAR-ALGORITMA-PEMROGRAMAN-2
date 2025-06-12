package main

import (
	"fmt"
	"time"
)

func TampilkanPertanyaan() {
	pertanyaan := []string{
		"Seberapa sering Anda merasa cemas/khawatir berlebihan (1-5)?",
		"Seberapa sering mengalami kesulitan tidur (1-5)?",
		"Seberapa sering kehilangan minat pada aktivitas (1-5)?",
		"Seberapa sering kesulitan berkonsentrasi (1-5)?",
		"Seberapa sering merasa sedih/putus asa (1-5)?\n",
	}

	fmt.Println("\n=== ASSESSMENT KESEHATAN MENTAL ===")
	fmt.Println("Skala: 1(Tidak pernah) - 5(Selalu)")
	fmt.Println("================================\n")

	for i, p := range pertanyaan {
		fmt.Printf("%d. %s\n", i+1, p)
	}
}

func InputJawabanKuesioner(jumlahPertanyaan int) []int {
	jawaban := make([]int, jumlahPertanyaan)
	for i := 0; i < jumlahPertanyaan; i++ {
		inputValid := false
		for !inputValid {
			fmt.Printf("Pertanyaan %d: ", i+1)
			var input int
			fmt.Scan(&input)
			if input >= 1 && input <= 5 {
				jawaban[i] = input
				inputValid = true
			} else {
				fmt.Println("Error: Masukkan angka 1-5")
			}
		}
	}
	return jawaban
}

func HitungTotalSkor(jawaban []int) int {
	total := 0
	for _, nilai := range jawaban {
		total += nilai
	}
	return total
}

func FormatTanggal(tanggalString string) time.Time {
	tanggal, _ := time.Parse("01-01-2025", tanggalString)
	return tanggal
}
