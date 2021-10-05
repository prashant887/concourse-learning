package main
import "fmt"

func add(a int,b int) int {
return a+b
}

func main() {
fmt.Println("Hello ")
res:=add(4,8)
fmt.Println("res : ",res)
}