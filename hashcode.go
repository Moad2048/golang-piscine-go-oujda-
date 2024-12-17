package main

import (
	"fmt"
)
func HashCode(dec string) string {
    mk := ""
    for _, c:= range dec{
        m := (int(c) + len(dec)) % 127
        if m < 32{
            m+=32
        }
        mk+=string(m)
    }
    return mk
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}