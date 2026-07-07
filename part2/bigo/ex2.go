package main

import (
	"fmt"
)

// MyBuffer буфер фиксированного размера
type MyBuffer struct {
	data     []int
	capacity int
}

func NewMyBuffer(capacity int) *MyBuffer {
	return &MyBuffer{
		data:     make([]int, 0, capacity),
		capacity: capacity,
	}
}

// Push добавляет элемент в конец.
func (mb *MyBuffer) Push(val int) {
	if len(mb.data) == mb.capacity {
		// Мы вынуждены сдвинуть весь массив на 1 ячейку влево,
		// чтобы освободить место в конце для нового элемента.
		copy(mb.data, mb.data[1:])
		mb.data[mb.capacity-1] = val
	} else {
		// Если место есть свободное, то просто добавляем значение.
		mb.data = append(mb.data, val)
	}
}

// Pop забирает самый старый элемент (с начала).
func (mb *MyBuffer) Pop() (int, bool) {
	if len(mb.data) == 0 {
		return 0, false
	}
	val := mb.data[0]
	// Сдвигаем оставшиеся элементы на место удаленного.
	copy(mb.data, mb.data[1:])
	mb.data = mb.data[:len(mb.data)-1]
	return val, true
}

// Iterate проходит по элементам от старых к новым
func (mb *MyBuffer) Iterate(callback func(int, int)) {
	for i := 0; i < len(mb.data); i++ {
		callback(i, mb.data[i])
	}
}

func main() {
	mb := NewMyBuffer(3)

	fmt.Println("Записываем 3 значения: 10, 20 и 30")
	mb.Push(10)
	mb.Push(20)
	mb.Push(30)

	fmt.Print("Текущие значения: ")
	mb.Iterate(func(i, v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Printf("\n\n")

	fmt.Println("Достаем один самый старый элемент")
	val, _ := mb.Pop()
	fmt.Printf("Самый старый элемент: %d", val)
	fmt.Printf("\n\n")

	fmt.Println("Записываем 40 и 50")
	mb.Push(40) // это значение просто добавится, так как .Pop() выше одно место освободил
	mb.Push(50) // Буфер заполнен, поэтому 50 перезапишет самый старый элемент (которым сейчас является 20).

	fmt.Print("Текущие значения: ")
	mb.Iterate(func(i, v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()
}
