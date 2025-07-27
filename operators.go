package main

import "fmt"
// arithmetic + - * / %
//logical and or not
//relational < > == != <= >=
//bitwise & | ^
// assignment = += := -= /= *=

func main(){
	x := true
	y := false
	// logical:
	fmt.Println(x && y)//false
	fmt.Println(x || y)//true
	fmt.Println(!x)//false
	a:=4// 0100  0-> false
	b:=5// 0101  1-> true
	// relational
	fmt.Println(a>b)
	fmt.Println(a==b)
	fmt.Println(a!=b)
	// bitwise
	fmt.Println(a&b)
	fmt.Println(a|b)
	fmt.Println(a^b)
	//assignment
	a+=4
	b-=3
	fmt.Println(a)
	fmt.Println(b)

}