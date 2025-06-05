package main

func SequentialSearch(data []Assessment, idUser string) []Assessment {
	var hasil []Assessment
	for _, a := range data {
		if a.IDUser == idUser {
			hasil = append(hasil, a)
		}
	}
	return hasil
}
