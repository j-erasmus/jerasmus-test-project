package main 

import "fmt"
import "time"

func main() { 
	i:=2
	fmt.Println("write ",i,"as")

	switch i{ 
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){ 
	case time.Saturday,time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println(time.Now().Weekday(),"weekday")
	}

	t:= time.Now()
	switch{ 
	case t.Hour()<12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	whatAmI := func(i interface{}){ 
		switch t := i.(type){ 
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("an int")
		default:
			fmt.Printf("unknown type %T \n",t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("tushar")
}