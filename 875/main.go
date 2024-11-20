package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}

func minEatingSpeed(piles []int, h int) int {
	max := 0
	// Находим самую большую кучу
	for _, i := range piles {
		if i > max {
			max = i
		}
	}

	// указыаем минимальную скорость 1 и максимальную, которая равна велечине самой большой кучи
	l, r := 1, max
	// используем бинарный поиск, для нахождения минимально возможной скорости
	for l <= r {
		// Находим скорость
		speed := l + (r-l)/2
		// Проходимся по срезу куч и высчитываем, сколько у нас часов это займет
		totalTime := 0
		for _, j := range piles {
			totalTime += int(math.Ceil(float64(j) / float64(speed)))
		}
		// если заняло меньше или сколько нужно, то записываем результат
		if totalTime <= h {
			max = speed
			r = speed - 1
			// Иначе двигаем левый указатель на новое место
		} else {
			l = speed + 1
		}
	}

	return max
}
