package main

import (
	"errors"
	"fmt"
	"slices"
)

var (
	ErrInvalidRange = errors.New("invalid range")
	ErrSeatOccupied = errors.New("seat already occupied")
)

type Range struct {
	Start int
	End   int // включительно
}

type Cinema struct {
	seats []bool
}

func NewCinema(seats []bool) *Cinema {
	return &Cinema{seats: slices.Clone(seats)}
}

// FindAvailable ищет свободные диапазоны длиной >= k.
// Метод должен вернуть все найденные диапазоны.
// Если передается отрицательный аргумент или ноль - вернется пустой слайс.
func (c *Cinema) FindAvailable(k int) []Range {
	res := make([]Range, 0)

	if k <= 0 {
		return res
	}

	counter := 0
	start := 0

	for current, el := range c.seats {
		if el {
			if counter >= k {
				res = append(res, Range{
					Start: start,
					End:   current - 1,
				})
			}
			counter = 0
		} else {
			if counter == 0 {
				start = current
			}
			counter++
		}
	}

	if counter >= k {
		res = append(res, Range{
			Start: start,
			End:   len(c.seats) - 1,
		})
	}

	return res
}

// Book бронирует диапазон мест [start, end] включительно.
func (c *Cinema) Book(start, end int) error {

	if start > end {
		return ErrInvalidRange
	}
	wantedSeats := c.seats[start : end+1]
	for _, el := range wantedSeats {
		if el == true {
			return ErrSeatOccupied
		}
	}

	for i := start; i <= end; i++ {
		c.seats[i] = true
	}

	return nil
}

func main() {

	seats := []bool{true, false, false, false, true, false, false, false, false, true}

	Cinema := NewCinema(
		seats,
	)

	fmt.Println(Cinema.FindAvailable(3))

}
