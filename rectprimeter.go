package main

import (
	"fmt"
)
func RectPerimeter(w, h int) int {
    mm := (w + h) * 2
    if w < 0 || h < 0{
        return -1
    }
    return mm

}
func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}