package main

import (
	"math"
)

type Range struct {
	From int `json:"from"` // начальный индекс
	To   int `json:"to"`   // конечный индекс (включительно)
}

//input{"temps":[1.5, 2.3, 3.1, 2.8, 1.9, 3.5, 3.5, 2.0], "k": 3}
//output[{"from":4,"to":6},{"from":5,"to":7}]

// FindMaxAvgWindows находит окна с максимальной средней температурой.
func FindMaxAvgWindows(temps []float64, k int) []Range {
	if k <= 0 || len(temps) < k {
		return nil // ✅ именно nil, не []Range{}
	}
	sum := 0.0
	avg := 0.0
	max := -1.0
	max_range := []Range{}

	for i := 0; i <= len(temps)-k; i++ {

		if i == 0 {
			//считаем среднее для первого окна
			for j := 0; j < k; j++ {
				sum += temps[j]
			}
			avg = sum / float64(k)
			avg = math.Round(avg*10) / 10
			max = avg
			max_range = append(max_range, Range{From: i, To: i + k - 1})
		} else {
			//считаем среднее для последующих окон, используя предыдущую сумму
			sum = sum - temps[i-1] + temps[i+k-1]
			avg = sum / float64(k)
			avg = math.Round(avg*10) / 10
			if avg > max {
				max = avg
				max_range = []Range{Range{From: i, To: i + k - 1}}
			} else if avg == max {
				max_range = append(max_range, Range{From: i, To: i + k - 1})
			}
		}

	}

	return max_range
}
