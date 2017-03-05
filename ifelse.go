package main

import "fmt"


func main() {

    if true {
        fmt.Println("Lord of the Rings")
    } else {
        fmt.Println("Harry Potter")
    }

    if 8%4 == 123 {
        fmt.Println("What if this is true ?")
    }

    if num := 6; num > 0 {
        fmt.Println("larger")
    } else if num < 10 {
        fmt.Println("smaller")
    } else {
        fmt.Println("What is this ?")
    }
}
