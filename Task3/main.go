package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("Введите строку с числами, которые нужно сложить.")
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(adding(s1, s2))
}

func adding(par_1, par_2 string) int64 {
	number := func(s string) string {
		var result string
		for _, i := range s {
			if unicode.IsDigit(i) {
				result += string(i)
			}
		}
		return result
	}
	d_1 := number(par_1)
	d_2 := number(par_2)
	num1, err1 := strconv.ParseInt(d_1, 10, 64)
	num2, err2 := strconv.ParseInt(d_2, 10, 64)
	if err1 != nil {
		num1 = 0
	}
	if err2 != nil {
		num2 = 0
	}
	return num1 + num2
}
