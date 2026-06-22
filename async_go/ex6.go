package main

import "fmt"

type Statistic struct {
	UserId int
	steps  int
}

type Result struct {
	UserIds []int
	steps   int
}

// достает из мапа копию если там масив - поэтому если хотим менять элемент масива нужно хранить в мапе указатель или сначала считать и потом записать
// но с слайс - ссылочный тип поэтому для него указатель не надо
func GetChampion(Statistics [][]Statistic) Result {
	ResMap := make(map[int][]int)
	for _, statistic := range Statistics {
		for _, v := range statistic {

			ResMap[v.UserId][0] += 1
			ResMap[v.UserId][0] 
			fmt.Println(v.UserId, v.steps)
		}
	}

	return Result{}
}

func main() {

	input := [][]Statistic{
		{
			Statistic{
				UserId: 1,
				steps:  1000,
			},
			Statistic{
				UserId: 2,
				steps:  600000,
			},
		},

		{
			Statistic{
				UserId: 1,
				steps:  1000,
			},
		},
	}

	GetChampion(input)
}
