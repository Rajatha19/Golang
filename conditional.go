package main
import "fmt"

//>90 A
//>80<=90 B
//>70<=80 C
//<70 D
func main(){
	/*
	marks := 65
	if marks > 90{
		fmt.Println("A")
	} else if marks > 80 && marks <=90{
		fmt.Println("B")
	} else if marks >70 && marks <=80{
		fmt.Println("C")
	} else{
		fmt.Println("D")
	}
		*/
	day:="saturday"
	/*
	if day == "monday" || day=="tuesday"	|| day == "wednesday"{
		fmt.Println("weekday")
	} else if day== "saturday" || day =="sunday"{
		fmt.Println("weekend")
	}*/


	switch day {
	case "monday","tuesday","wednesday","thursday","friday":
		fmt.Println("weekday")
	case "saturday":
		fmt.Println("weekend")
		fallthrough
	case "sunday":
		fmt.Println("weekend bonus")
    default:
		fmt.Println("invalid")
	}
}