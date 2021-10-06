package main
import "fmt" 

func add(a int,b int) int {
	return a+b 
}

func main(){
	fmt.Println("Hello World")
	fmt.Printf(" Sum of %d and %d is %d ",8,4,add(8,4))
}