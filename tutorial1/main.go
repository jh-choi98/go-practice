package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice))
}
func sumIntSlice(slice []int) int{
	var sum int
	for _, v := range slice{
		sum += v
	}
	return sum
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}
