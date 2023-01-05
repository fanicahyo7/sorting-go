package main

import (
	"fmt"
)

func main() {
	data := []int{1, 4, 5, 6, 8, 2}

	fmt.Println("========Soal 1========")
	fmt.Println("Input ", data)
	diagram(data)
	fmt.Println()

	fmt.Println("========Soal 2========")
	fmt.Println("Input ", data)
	fmt.Println("Sortir ", insertionSortir(data, false))
	fmt.Println("Steps visualization")
	_ = insertionSortir(data, true)
	fmt.Println()

	fmt.Println("========Soal 3========")
	fmt.Println("Input ", data)
	fmt.Println("Sortir ", reverseSorting(data, false))
	fmt.Println("Steps visualization")
	_ = reverseSorting(data, true)
}

func insertionSortir(data []int, withDiagram bool) []int {
	var dataArray []int
	dataArray = append(dataArray, data...)
	for i := 1; i < len(dataArray); i++ {
		value := dataArray[i]
		j := i - 1
		for j >= 0 && dataArray[j] > value {
			titip := dataArray[j+1]
			dataArray[j+1] = dataArray[j]
			dataArray[j] = titip
			j--

			if withDiagram {
				diagram(dataArray)
			}
		}
		dataArray[j+1] = value
	}
	return dataArray
}

func reverseSorting(data []int, withDiagram bool) []int {
	var dataArray []int
	dataArray = append(dataArray, data...)

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if dataArray[j] < dataArray[j+1] {
				titip := dataArray[j]
				dataArray[j] = dataArray[j+1]
				dataArray[j+1] = titip
				if withDiagram {
					diagram(dataArray)
				}
			}
		}
	}
	return dataArray
}

func diagram(data []int) {
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
