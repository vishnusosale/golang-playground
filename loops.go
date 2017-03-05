package main

import "fmt"


func main() {
    var i int = 3;

    for i <= 9 {
        fmt.Println(i)
        i = i + 1
    }

    for {
        fmt.Println("Hi")
        break
    }

    for j := 0; j<=9; j++ {
        fmt.Println("j = ", j)
    }

    for n := 0; n <=5; n++ {
        if n%2 == 0{
            continue
        }
        fmt.Println(n)
    }
}
