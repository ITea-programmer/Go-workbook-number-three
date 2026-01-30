package main

import "fmt"
import "strconv"

func main() {
    fmt.Println("Введите число")
    var ch uint
    fmt.Scan(&ch)
   fmt.Println(fn(ch))
}
func fn(ch uint) uint{
    ch_1 = strconv.Itoa(ch)
    answer := ""
    for i in ch_1{
        if i in "2468"{
            s += i
        }
    }
    if len(answer) != 0{
        answer := srconv.Atoi(answer)
        return answer
    } else{
        return 100
    }
}
