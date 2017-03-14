package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println(a)
    fmt.Println(a[4])

    a[4] = 100

    fmt.Println(a[4])


    b := [10]int{123,12351,124351,45231,5423,1235,124,354,1235}
    fmt.Println(b)

    const s int = 5
    var twoD [s][s]int

    for o := 0; o<s; o++ {
        for p := 0; p<s; p++ {
            twoD[o][p] = o + p
        }
    }
    fmt.Println(twoD)

}
