package sort

import "github.com/dmitry-bakeev/flight-schedule/pkg/models"

func Reverse(data []*models.Flight) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
