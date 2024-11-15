package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(4421))
}

// Решение 1
func subtractProductAndSum(n int) int {
	multi := 1
	sum := 0

	for n != 0 {
		multi *= n % 10
		sum += n % 10
		n = n / 10
	}
	return multi - sum
}
