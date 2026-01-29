package main

import (
	"fmt"
	"strings"
)
type Battery struct {
	zaryad string
}

func (bat Battery) String() string {
	zaryazheno := strings.Count(bat.zaryad, "1")
	pusto := len(bat.zaryad) - zaryazheno
	
	result := "[" + strings.Repeat(" ", pusto) + strings.Repeat("X", zaryazheno) + "]"
	return result
}

func main() {
	var s string
	fmt.Scan(&s)
	if len(s) != 10 {
		fmt.Println("Ошибка: требуется строка из 10 символов")
		return
	}
	battery := Battery{zaryad: s}
	fmt.Println(battery)
}
