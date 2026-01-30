package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите строку с числами, которые нужно поделить.")
	var s string
	fmt.Scanf("%[^\n]", &s)
	
	result, err := delenie(s)
	if err != nil {
		fmt.Printf("%.4f\n", 0.0)
		return
	}
	
	fmt.Printf("%.4f\n", result)
}

func delenie(s string) (float64, error) {
	s = strings.ReplaceAll(s, " ", "")
	
	if s == "" {
		return 0.0, fmt.Errorf("Пустой ввод")
	}
	
	nums := strings.Split(s, ";")
	
	if len(nums) != 2 {
		return 0.0, fmt.Errorf("Неверный формат ввода, ожидается: число1;число2")
	}
	
	num1 := strings.ReplaceAll(nums[0], ",", ".")
	num2 := strings.ReplaceAll(nums[1], ",", ".")
	
	Num1, err1 := strconv.ParseFloat(num1, 64)
	if err1 != nil {
		return 0.0, fmt.Errorf("Первое число '%s' имеет неверный формат", nums[0])
	}
	
	Num2, err2 := strconv.ParseFloat(num2, 64)
	if err2 != nil {
		return 0.0, fmt.Errorf("Второе число '%s' имеет неверный формат", nums[1])
	}
	
	if Num2 == 0.0 {
		return 0.0, fmt.Errorf("деление на ноль")
	}
	
	return Num1 / Num2, nil
}
