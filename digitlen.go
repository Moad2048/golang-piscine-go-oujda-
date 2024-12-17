package main

import (
	"fmt"
)
func DigitLen(n, base int) int {
    sum := 0
    if base < 2 || base > 36{
        return -1
    }
    if n < 0 {
        n = -n
    }
    for n > 0 {
        n = n / base
        sum++
    }
    if sum == 0{
        return 1
    }
    return sum
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}