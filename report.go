package main

import (
	"fmt"
	"time"
)

func TampilkanLimaTerakhir(data []Assessment, idUser string) {
	count := 0
	fmt.Println("\n5 Assessment Terakhir:")
	for i := len(data) - 1; i >= 0 && count < 5; i-- {
		if data[i].IDUser == idUser {
			fmt.Printf("%d. Tanggal: %s, Skor: %d\n", count+1,
				data[i].Tanggal.Format("02-01-2006"),
				data[i].SkorTotal)
			count++
		}
	}
	if count == 0 {
		fmt.Println("Data tidak ditemukan.")
	}
}

func HitungRataRataSebulan(data []Assessment, idUser string) float64 {
	now := time.Now()
	var total, count int
	for _, a := range data {
		if a.IDUser == idUser && now.Sub(a.Tanggal).Hours() <= 24*30 {
			total += a.SkorTotal
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(total) / float64(count)
}

func Rekomendasi(skor int) string {
	switch {
	case skor <= 10:
		return "Kondisi stabil. Terus jaga keseimbangan mentalmu!"
	case skor <= 15:
		return "Perlu perhatian. Coba lakukan: \n- Relaksasi\n- Curhat ke teman\n- Aktivitas menyenangkan"
	case skor <= 20:
		return "Risiko sedang. Pertimbangkan:\n- Konsultasi ke psikolog\n- Teknik manajemen stres\n- Rutin olahraga"
	default:
		return "Risiko tinggi. Segera cari bantuan profesional:\n- Psikolog/Psikiater\n- Layanan konseling\n- Dukungan sosial"
	}
}
