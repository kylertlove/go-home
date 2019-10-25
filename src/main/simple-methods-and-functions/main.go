package simplemethodsandfunctions

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	classElem string = "asdf"
)

func main() {
	println("\n\n---------")
	println("Methods and simple funcs")
	fmt.Println(timeFunction())
	goLooping()

	println("\n\n---------")
	println("Methods with error handling")
	mathFunction, err := simpleMath(3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(mathFunction)

	println("\n\n---------")
	println("Pointers")
	pointerTesting()
}

/**
* Function with both type inference and manual type declaration and error handling
 */
func simpleMath(first int) (int, error) {
	if first < 0 {
		err := errors.New("Cant handle negative numbers")
		return 0, err
	}
	newFirst := first
	var second int = 32
	return newFirst + second, nil
}

func timeFunction() string {
	now := time.Now().String()
	return now
}

/**
Go has one for loop but is used like for/while/switch
*/
func goLooping() {
	fmt.Println("Standard ForLoop")
	for index := 0; index < 5; index++ {
		fmt.Print(index)
	}
	fmt.Println("")
	fmt.Println("ForLoop that acts like a while loop")
	canContinue := true
	i := 0
	for canContinue {
		if i == 4 {
			canContinue = false
		}
		fmt.Print(i)
		i++
	}

	fmt.Println("")
	fmt.Println("For break on condition")
	j := 0
	for {
		if j == 5 {
			break
		}
		fmt.Print(j)
		j++
	}
}

// Testing pass by value vs reference
func pointerTesting() {
	//go can pass by reference or by value.
	//to change a var by value equate it to the return obj
	passByValue := "value"
	passByValue = getPassByValue(passByValue)
	println("Pass by value: " + passByValue)

	//better memory management - pass by ref
	// & prepended to variable is the mem location
	// * prepended gets the data from the passed in reference
	passByReference := "Reference"
	getPassByRef(&passByReference)
	println("Pass By Ref: " + passByReference)
	println("\n\n---------")
	println("Structs and Interfaces")
}

func getPassByValue(v string) string {
	return "changed"
}

func getPassByRef(v *string) {
	*v = "changed"
}
