package main

func SelectionSortBySkor(data *[]Assessment, ascending bool) {
	n := len(*data)
	for i := 0; i < n; i++ {
		extreme := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if (*data)[j].SkorTotal < (*data)[extreme].SkorTotal {
					extreme = j
				}
			} else {
				if (*data)[j].SkorTotal > (*data)[extreme].SkorTotal {
					extreme = j
				}
			}
		}
		(*data)[i], (*data)[extreme] = (*data)[extreme], (*data)[i]
	}
}

func InsertionSortByTanggal(data *[]Assessment, ascending bool) {
	for i := 1; i < len(*data); i++ {
		key := (*data)[i]
		j := i - 1

		if ascending {
			for j >= 0 && (*data)[j].Tanggal.After(key.Tanggal) {
				(*data)[j+1] = (*data)[j]
				j--
			}
		} else {
			for j >= 0 && (*data)[j].Tanggal.Before(key.Tanggal) {
				(*data)[j+1] = (*data)[j]
				j--
			}
		}
		(*data)[j+1] = key
	}
}
