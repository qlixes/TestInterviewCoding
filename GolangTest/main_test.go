package main

import (
    "testing"
)

func TestFibonacci(t *testing.T) {
    testFunction := Fibonacci(5)
    expected := []int{1,1,2,3,5}

    for i, v := range testFunction {
        if v != expected[i] {
            t.Errorf("Test Fibonacci Failed")
        }
    }
}

func TestOddOrEven(t *testing.T) {
    testFunction := OddOrEven(5)
    expected := []int{1,3,5}

    for i, v := range testFunction {
        if v != expected[i] {
            t.Errorf("Test OddOrEven Failed")
        }
    }
}
