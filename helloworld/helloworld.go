package main
import "fmt"
func main(){//go的左大括号必须和函数名在同一行，否则会报错
	fmt.Println(Hello())
}
func Hello() string{
	return "HelloGo"
}
