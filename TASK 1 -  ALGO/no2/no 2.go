package main

import (
	"fmt"
	"math"
)

var denominations = []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

func calculateChange(totalBelanja, uangDibayar int) (map[int]int, bool) {
	kembalian := uangDibayar - totalBelanja

	if kembalian < 0 {
		return nil, false
	}

	result := make(map[int]int)
	for _, denom := range denominations {
		if kembalian >= denom {
			count := int(math.Floor(float64(kembalian / denom)))
			result[denom] = count
			kembalian -= count * denom
		}
	}

	return result, true
}

func main() {
	totalBelanja := 700649
	uangDibayar := 800000

	kembalian, success := calculateChange(totalBelanja, uangDibayar)
	if success {
		totalKembalian := 0
		for denom, count := range kembalian {
			totalKembalian += denom * count
		}
		kembalianRounded := int(math.Floor(float64(totalKembalian)/100)) * 100
		fmt.Printf("Kembalian yang harus diberikan kasir: Rp %d,\ndibulatkan menjadi Rp %d\n", totalKembalian, kembalianRounded)
		fmt.Println("Pecahan uang:")
		for denom, count := range kembalian {
			if denom >= 100 {
				fmt.Printf("%d lembar Rp%d\n", count, denom)
			} else {
				fmt.Printf("%d koin %d\n", count, denom)
			}
		}
	} else {
		fmt.Println("Jumlah uang yang dibayarkan kurang dari total belanja.")
	}
}
