package main

import (
    "fmt"
)

func Fibonacci(askNumber int) []int {

    listArray := []int{1}

    if askNumber <= 1 {
        return listArray
    }

    for i := range askNumber - 1 {

        indexArray := 0;

        if i >= 1 {
            indexArray = i - 1
        }

        itemArray := listArray[indexArray] + i

        listArray = append(listArray, itemArray)
    }

    return listArray
}


func OddOrEven(askNumber int) []int {

    var listArray []int

    if askNumber <= 0 {
        return listArray
    }

    checkOddOrEven := askNumber % 2

    for i:= 1; i <= askNumber; i++ {
        
        if i % 2 == checkOddOrEven {
            listArray = append(listArray, i)
        }
    }

    return listArray
}

func main() {
    fmt.Printf("%d", Fibonacci(5))
    fmt.Printf("%d", OddOrEven(10))
}
