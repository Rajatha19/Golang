package main

import "fmt"

func main(){
	//array
	//marks := [5]int{50,60,70,80,90}
	//sum:= 0
	/*for i:= 0 ;i<len(marks); i++{
		sum = sum + marks[i]
	}
	fmt.Println(sum/len(marks))*/
/*
	for _,value := range marks{
	//	fmt.Println(index,marks[index])
	    fmt.Println(value)
	}*/


	//slice
	//marksList := [5]int{50,60,70,80,90}
    /*for _,value := range marksList{
	//	fmt.Println(index,marks[index])
	    fmt.Println(value)
	}*/
	//subarray
	/*fmt.Println(marksList[1:4])
	fmt.Println(marksList[2:])
	fmt.Println(marksList[:2])
*/
// capacity = cap
    /*fmt.Println(len(marksList),cap(marksList))
    values := make([]int,3,5) // diff len and cap
	fmt.Println(len(values),cap(values),values)
	//add value use append
	values=append(values,100)
	fmt.Println(len(values),cap(values),values)
	
	values=append(values,200,300)
	fmt.Println(len(values),cap(values),values)*/
	// map= key value pair of list
    // keys are custom defined

    /*students:= map[int]int{
		0:50,
		1:100,
		100:1000,
	}
	/*
	for key,value:= range students {
		fmt.Println(key,value)
	}

	students[2] = 200
	fmt.Println("after insert")
	for key,value:= range students {
		fmt.Println(key,value)
	}

	delete(students,100)
	fmt.Println("after delete")
	for key,value:= range students {
		fmt.Println(key,value)
	}*/
	/*delete(students,2)
	fmt.Println(students)
	*/
	/*students := map[string]int{
		"john":90,
		"hel":98,
	}
	students["john"]=96
	value , ok :=students["john"]
	fmt.Println(value, ok)
	fmt.Println(len(students))*/

	students := make(map[string]int)
	students["j"]=80
	students["p"]=90
	fmt.Println(students)



}