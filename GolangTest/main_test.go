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

func TestOdd(t *testing.T) {
    testFunction := OddOrEven(5)
    expected := []int{1,3,5}

    for i, v := range testFunction {
        if v != expected[i] {
            t.Errorf("Test OddOrEven Failed")
        }
    }
}

func TestEven(t *testing.T) {
    testFunction := OddOrEven(10)
    expected := []int{2,4,6,8,10}

    for i, v := range testFunction {
        if v != expected[i] {
            t.Errorf("Test OddOrEven Failed")
        }
    }
}
