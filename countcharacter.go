package main
import "fmt"

func CountChar(str string, c rune) int {
    sum := 0
    //slic := []rune(str)
    for _, i:= range str{
        if i == c{
            sum+=1
        }
    }
    return sum
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}