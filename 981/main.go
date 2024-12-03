package main

import "fmt"

func main() {
	timeMap := Constructor()

	timeMap.Set("foo", "bar", 1)
	timeMap.Get("foo", 1)
	timeMap.Get("foo", 3)
	timeMap.Set("foo", "bar2", 4)
	timeMap.Get("foo", 4)
	fmt.Println(timeMap.Get("foo", 5))
}

// Так как у нас задача создать time-based key-value структуру данных, то сначала мы определим
// структуру Pairs, которая будет хранить в себе значение (строку) и время.
// Такая структура будет хранить информацию о времени, когда эти значения были установлены.
type Pairs struct {
	val  string
	time int
}

// Основная структура TimeMap - это map, в котором ключом будет строка, а значением массив Pairs.
type TimeMap struct {
	mp map[string][]Pairs
}

// Данная функция нужна для инициализации структуры.
func Constructor() TimeMap {
	return TimeMap{mp: make(map[string][]Pairs)}
}

// Данная функция нужна для установки значения в структуру.
func (this *TimeMap) Set(key string, value string, timestamp int) {
	// Создаем структуру Pairs и передаем в нее значение и время.
	pair := Pairs{
		val:  value,
		time: timestamp,
	}

	// Теперь добавляем эту структуру в массив, который находится под ключом key.
	this.mp[key] = append(this.mp[key], pair)
}

// Данная функция нужна для получения значения из структуры.
func (this *TimeMap) Get(key string, timestamp int) string {
	// Проверяем существует ли массив под ключом key.
	if _, ok := this.mp[key]; !ok {
		return ""
	}

	// Получаем массив под ключом key.
	arr := this.mp[key]

	// Создаем два указателя на начало и конец массива.
	l, r := 0, len(arr)-1

	// Также создадим результирующую строку
	res := ""

	// Запускаем цикл пока указатель l меньше или равен указателю r.
	for l <= r {
		// Получаем среднее значение
		mid := l + (r-l)/2

		// Получаем значение и время
		value, time := arr[mid].val, arr[mid].time

		// Теперь нам нужно сравнить время, которое передали в функцию с тем,
		// что у нас были добавлены в структуру.

		// Если полученное по индексу mid время меньше или равно timestamp,
		// то мы это значение запишем в res, но все еще будем продолжать проходить по циклу.
		if time <= timestamp {
			res = value
			// Так как по условию при добавлении значения в структуру время будет постоянно увеличиваться
			// Поэтому мы знаем, что все значения будут отсортированы по времени. И справа будет время больше
			// И поэтому мы сдвинем левый указатель на mid+1
			l = mid + 1
			// иначе мы двигаем правый указатель
		} else {
			r = mid - 1
		}
	}

	return res
}
