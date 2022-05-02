package sort

import (
	"github.com/dmitry-bakeev/flight-schedule/pkg/models"
)

func merge(ldata []*models.Flight, rdata []*models.Flight) []*models.Flight {
	result := make([]*models.Flight, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < len(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx].NumberFlight < rdata[ridx].NumberFlight:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}
	return result
}

func MergeSort(data []*models.Flight, r chan<- []*models.Flight) {
	if len(data) == 1 {
		r <- data
		return
	}

	leftChan := make(chan []*models.Flight)
	rightChan := make(chan []*models.Flight)
	middle := len(data) / 2

	go MergeSort(data[:middle], leftChan)
	go MergeSort(data[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	r <- merge(ldata, rdata)
}
