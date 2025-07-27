package main

import "fmt"
// for init cond; test cond; inc/dec 

func main(){
	n:=10
	/*for i:=1;i<=n;i++{
		fmt.Println(i)
	}*/
/*
    j:=1
	for j<=n{
		fmt.Println(j)
		j=j+2
	}*/

	/*j:= 1
	for {
		fmt.Println(j)
		j=j+1
		if j>10{
			break
		}
	}*/

	for i:=1;i<=n;i++{
		if i%2 !=0{
			continue
		}
		fmt.Println(i)
	}


}