package main

import (
	"fmt"
	"time"
)

func TambahAssessment(dataAssessment *[]Assessment, assessmentBaru Assessment) {
	*dataAssessment = append(*dataAssessment, assessmentBaru)
	fmt.Println("Assessment berhasil ditambahkan!")
}

func UbahAssessment(dataAssessment *[]Assessment, idAssessment string, idUser string) (*Assessment, bool) {
	for i := range *dataAssessment {
		if (*dataAssessment)[i].IDAssessment == idAssessment && (*dataAssessment)[i].IDUser == idUser {
			return &(*dataAssessment)[i], true
		}
	}
	fmt.Println("Assessment tidak ditemukan atau tidak sesuai dengan user.")
	return nil, false
}
func HapusAssessment(dataAssessment *[]Assessment, idAssessment string, idUser string) bool {
	for i := range *dataAssessment {
		if (*dataAssessment)[i].IDAssessment == idAssessment && (*dataAssessment)[i].IDUser == idUser {
			*dataAssessment = append((*dataAssessment)[:i], (*dataAssessment)[i+1:]...)
			fmt.Println("Assessment berhasil dihapus!")
			return true
		}
	}
	fmt.Println("Assessment tidak ditemukan atau tidak sesuai dengan user.")
	return false
}

func TampilkanSemuaAssessment(data []Assessment) {
	for i, a := range data {
		fmt.Printf("%d. ID: %s | User: %s | Tanggal: %s | Skor: %d\n",
			i+1,
			a.IDAssessment,
			a.IDUser,
			a.Tanggal.Format("01-01-2025"),
			a.SkorTotal)
		fmt.Println("   Rekomendasi:", Rekomendasi(a.SkorTotal))
		fmt.Println("------------------------------")
	}
}

func TampilkanLimaTerakhir(data []Assessment, idUser string) {
	count := 0
	fmt.Println("\n5 Assessment Terakhir:")
	for i := len(data) - 1; i >= 0 && count < 5; i-- {
		if data[i].IDUser == idUser {
			fmt.Printf("%d. Tanggal: %s, Skor: %d\n", count+1,
				data[i].Tanggal.Format("01-01-2025"),
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
