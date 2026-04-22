package main

import "slices"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums = slices.Insert(nums, 2, 99999)                            //отдаст новый слайс с 99999 на индексе 2
	slices.Contains(nums, 40)                                       //проверяент если 40 в намс
	slices.ContainsFunc(nums, func(v int) bool { return v%2 == 0 }) //проверяет есть ли число которое делится на два без остатка
	//slices.Compare(s1, s2)
	//slices.CompareFunc(s1, s2)
	//slices.Concat() передаются слайсы и потом объхединяются в один
	//slices.Index(slice, to_find )
	//slices.Delete(slice, от(вкл), до(невкл) )
	//к каждой функции есть доп Func передается булева функция хотнелоб бы побольше примеров таких функций какие они бывают котрые передаются в фанк
	//slices.Reverse(numbers) //перевернуть

}
