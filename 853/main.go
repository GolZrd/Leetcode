package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
}

func carFleet(target int, position []int, speed []int) int {
	// пары позиция и скорость
	cars := [][2]int{}
	for i := 0; i < len(position); i++ {
		cars = append(cars, [2]int{position[i], speed[i]})
	}
	// сортируем по позиции
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	stack := []float64{}

	for _, car := range cars {
		timeToDest := float64(target-car[0]) / float64(car[1])
		stack = append(stack, timeToDest)
		// Проверяем что в стаке больше 2 записей и если предпоследнее значение больше
		// или равно только что добавленному, то мы последнее удаляем из стэка, и у нас по итогу в стэке будут
		// только уникальные значения
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
