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

type Pairs struct {
	val  string
	time int
}

type TimeMap struct {
	mp map[string][]Pairs
}

func Constructor() TimeMap {
	return TimeMap{mp: make(map[string][]Pairs)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	pair := Pairs{
		val:  value,
		time: timestamp,
	}
	this.mp[key] = append(this.mp[key], pair)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.mp[key]; !ok {
		return ""
	}
	arr := this.mp[key]
	l, r := 0, len(arr)-1

	res := ""

	for l <= r {
		mid := l + (r-l)/2
		value, time := arr[mid].val, arr[mid].time
		if time <= timestamp {
			res = value
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
