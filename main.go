package main

import (
	"fmt"
)

func main() {
	var dataAssessment []Assessment
	jumlahPertanyaan := 5

	for {
		fmt.Println("\n=== SISTEM ASSESSMENT KESEHATAN MENTAL ===")
		fmt.Println("1. Tambah Assessment Baru")
		fmt.Println("2. Ubah Assessment")
		fmt.Println("3. Hapus Assessment")
		fmt.Println("4. Tampilkan 5 Assessment Terakhir")
		fmt.Println("5. Rata-rata Skor 30 Hari Terakhir")
		fmt.Println("6. Urutkan Assessment")
		fmt.Println("7. Cari Assessment")
		fmt.Println("8. Tampilkan Rekomendasi Terbaru")
		fmt.Println("9. Tampilkan Semua Assessment")
		fmt.Println("0. Keluar")
		fmt.Print("[*] Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var idAssessment, idUser, tanggalStr string
			fmt.Print("\n=> ")
			fmt.Print("Masukkan ID Assessment: ")
			fmt.Scan(&idAssessment)
			fmt.Print("\n=> ")
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			fmt.Print("\n=> ")
			fmt.Print("Masukkan Tanggal (dd-mm-yyyy): ")
			fmt.Scan(&tanggalStr)
			tanggal := FormatTanggal(tanggalStr)

			TampilkanPertanyaan()
			jawaban := InputJawabanKuesioner(jumlahPertanyaan)
			skorTotal := HitungTotalSkor(jawaban)

			assessmentBaru := Assessment{
				IDAssessment: idAssessment,
				IDUser:       idUser,
				Tanggal:      tanggal,
				Jawaban:      jawaban,
				SkorTotal:    skorTotal,
			}

			TambahAssessment(&dataAssessment, assessmentBaru)
			fmt.Printf("Skor Total: %d\n", skorTotal)
			fmt.Println("Rekomendasi:", Rekomendasi(skorTotal))

		case 2:
			var idAssessment, idUser string
			fmt.Print("\n=> ")
			fmt.Print("Masukkan ID Assessment yang ingin diubah: ")
			fmt.Scan(&idAssessment)
			fmt.Print("\n=> ")
			fmt.Print("Masukkan ID User: ")
			fmt.Scan(&idUser)

			assessmentPtr, found := UbahAssessment(&dataAssessment, idAssessment, idUser)
			if !found {
				break
			}

			var tanggalStr string
			fmt.Print("\n=> ")
			fmt.Print("Masukkan Tanggal Baru (dd-mm-yyyy): ")
			fmt.Scan(&tanggalStr)
			tanggal := FormatTanggal(tanggalStr)

			TampilkanPertanyaan()
			jawaban := InputJawabanKuesioner(jumlahPertanyaan)
			skorTotal := HitungTotalSkor(jawaban)

			assessmentPtr.Tanggal = tanggal
			assessmentPtr.Jawaban = jawaban
			assessmentPtr.SkorTotal = skorTotal

			fmt.Printf("Skor Total: %d\n", skorTotal)
			fmt.Println("Rekomendasi:", Rekomendasi(skorTotal))

		case 3:
			var idAssessment, idUser string
			fmt.Print("\n=> Masukkan ID Assessment yang ingin dihapus: ")
			fmt.Scan(&idAssessment)
			fmt.Print("\n=> Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)

			if HapusAssessment(&dataAssessment, idAssessment, idUser) {
				fmt.Println("Assessment dengan ID", idAssessment, "milik user", idUser, "telah dihapus")
			} else {
				fmt.Println("Gagal menghapus - assessment tidak ditemukan atau tidak sesuai dengan user")
			}
		case 4:
			var idUser string
			fmt.Print("\n=> ")
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			TampilkanLimaTerakhir(dataAssessment, idUser)

		case 5:
			var idUser string
			fmt.Print("\n=> ")
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			rata := HitungRataRataSebulan(dataAssessment, idUser)
			if rata == 0 {
				fmt.Println("Tidak ada data dalam 30 hari terakhir.")
			} else {
				fmt.Printf("Rata-rata skor: %.2f\n", rata)
				fmt.Println("Rekomendasi:", Rekomendasi(int(rata)))
			}

		case 6:
			fmt.Println("\n1. Urutkan berdasarkan skor total (Ascending)")
			fmt.Println("2. Urutkan berdasarkan skor total (Descending)")
			fmt.Println("3. Urutkan berdasarkan tanggal (Ascending)")
			fmt.Println("4. Urutkan berdasarkan tanggal (Descending)")
			fmt.Print("Pilih metode pengurutan: ")
			var opsi int
			fmt.Scan(&opsi)

			switch opsi {
			case 1:
				SelectionSortBySkor(&dataAssessment, true)
				fmt.Println("\nData telah diurutkan berdasarkan skor total (Ascending):")
				TampilkanSemuaAssessment(dataAssessment)
			case 2:
				SelectionSortBySkor(&dataAssessment, false)
				fmt.Println("\nData telah diurutkan berdasarkan skor total (Descending):")
				TampilkanSemuaAssessment(dataAssessment)
			case 3:
				InsertionSortByTanggal(&dataAssessment, true)
				fmt.Println("\nData telah diurutkan berdasarkan tanggal (Ascending):")
				TampilkanSemuaAssessment(dataAssessment)
			case 4:
				InsertionSortByTanggal(&dataAssessment, false)
				fmt.Println("\nData telah diurutkan berdasarkan tanggal (Descending):")
				TampilkanSemuaAssessment(dataAssessment)
			default:
				fmt.Println("Pilihan tidak valid.")
			}

		case 7:
			var idUser string
			fmt.Print("\n=> ")
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			hasil := SequentialSearch(dataAssessment, idUser)
			if len(hasil) == 0 {
				fmt.Println("Data tidak ditemukan.")
			} else {
				fmt.Println("\nHasil pencarian:")
				for _, a := range hasil {
					fmt.Printf("ID: %s | Tanggal: %s | Skor: %d\n",
						a.IDAssessment,
						a.Tanggal.Format("02-01-2006"),
						a.SkorTotal)
					fmt.Println("Rekomendasi:", Rekomendasi(a.SkorTotal))
					fmt.Println("------------------------------")
				}
			}

		case 8:
			var idUser string
			fmt.Print("\n=> ")
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scan(&idUser)
			hasil := SequentialSearch(dataAssessment, idUser)
			if len(hasil) == 0 {
				fmt.Println("Data tidak ditemukan.")
			} else {
				skor := hasil[len(hasil)-1].SkorTotal
				fmt.Println("\nAssessment Terakhir:")
				fmt.Printf("Tanggal: %s\n", hasil[len(hasil)-1].Tanggal.Format("02-01-2006"))
				fmt.Printf("Skor: %d\n", skor)
				fmt.Println("Rekomendasi:", Rekomendasi(skor))
			}

		case 9:
			if len(dataAssessment) == 0 {
				fmt.Println("\nBelum ada data assessment.")
			} else {
				fmt.Println("\nSemua Assessment:")
				TampilkanSemuaAssessment(dataAssessment)
			}

		case 0:
			fmt.Println("\nTerima kasih telah menggunakan sistem assessment kesehatan mental!")
			return

		default:
			fmt.Println("\nPilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func TampilkanSemuaAssessment(data []Assessment) {
	for i, a := range data {
		fmt.Printf("%d. ID: %s | User: %s | Tanggal: %s | Skor: %d\n",
			i+1,
			a.IDAssessment,
			a.IDUser,
			a.Tanggal.Format("02-01-2006"),
			a.SkorTotal)
		fmt.Println("   Rekomendasi:", Rekomendasi(a.SkorTotal))
		fmt.Println("------------------------------")
	}
}
