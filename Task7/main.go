package main

import (
	"fmt"
	"io"
)

func main() {
	var ch, sum int
	
	for {
		_, err := fmt.Scan(&ch)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
		sum += ch
	}
	
	fmt.Println(sum)
}
