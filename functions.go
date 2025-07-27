package main

import "fmt"

func main()  {
	hello()
	output:= greet("john")
	fmt.Println(output)
	quotient,_:=divide(7,3) // _ blank identifier used to ignore 
	fmt.Println(quotient)

} 
// func <func name> ([args],[args]....)[return value][return value, return value]
func hello ()  {
	fmt.Println("hello")
	return
}

func greet(name string) string{
	return "hi " + name
}

func divide(dividend int , divisor int)(quotient int,remainder int){
	quotient=dividend/divisor
	remainder=dividend % divisor
	return //quotient,remainder
}
/*
string ""
int 0
*/