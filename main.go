package main

import "fmt"

func main() {
	data := []int{1, 4, 5, 6, 8, 2}

	// cari nilai tertinggi yang nantinya sebagai jumlah baris
	var nilaiTertinggi = data[0]
	for _, value := range data {
		if value > nilaiTertinggi {
			nilaiTertinggi = value
		}
	}

	// cetak baris
	for a := nilaiTertinggi; a >= 0; a-- {
		// bila dibaris terakhir akan mencetak nilai tiap index
		if a == 0 {
			for _, value := range data {
				fmt.Printf("%d ", value)
			}
		} else {
			// cetak kolom
			for _, value := range data {
				if value >= a {
					fmt.Print("| ")
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}

}
