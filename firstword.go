package main
import "fmt"
func FirstWord(s string) string {
    sum := ""
    for _,i := range s{
        if i != ' '{
        sum+=string(i)
        }else {
            break
        }
    }
    return sum + "\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}
