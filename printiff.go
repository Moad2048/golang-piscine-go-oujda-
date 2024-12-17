package main

import (
	"fmt"
)

func PrintIf(str string) string {
    if len(str) > 0 && len(str) < 3{
        return "Invalid Input\n"
    }else if len(str) >= 3{
        return "G\n"
    }else if len(str) <= 0{
        return "G\n"
    }
    return "G\n"
    
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
