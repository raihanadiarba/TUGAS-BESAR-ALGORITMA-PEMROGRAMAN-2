package main

import (
	"fmt"
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
