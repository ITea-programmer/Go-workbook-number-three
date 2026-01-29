package main

import "fmt"

func work(x int) int {
    if x%2 == 0 {
        return x
    }
    return x % 2
}

func main() {
    var a, b, c, d, e, f, g, h, i, j int
    fmt.Println("Введите 10 чисел через пробел")
    fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i, &j)
    
    numbers := []int{a, b, c, d, e, f, g, h, i, j}
    
    for i, num := range numbers {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(work(num))
    }
    fmt.Println()
}
