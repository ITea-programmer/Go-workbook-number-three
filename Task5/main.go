package main

import "fmt"

func main() {
	fmt.Println("Введите число")
	var ch uint
	fmt.Scan(&ch)
	fmt.Println(fn(ch))
}

func fn(ch uint) uint {
	result := 0
	multiplier := 1
	for ch > 0 {
		digit := ch % 10
		if digit%2 == 0 && digit != 0 {
			result = int(digit)*multiplier + result
			multiplier *= 10
		}
		ch /= 10
	}
	
	if result > 0 {
		return uint(result)
	}
	return 100
}
