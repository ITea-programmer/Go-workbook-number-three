package main

import (
	"fmt"
	"io"
)

func main() {
	var num, sum int
	
	for {
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		sum += num
	}
	
	fmt.Println(sum)
}
