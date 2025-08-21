package main

import (
	"errors"
	"fmt"
)

func division(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("invalid operation....can't divide by zero")
	}
	return (x / y), nil
}

func main() {
	var dividend float64
	var divisor float64
	var choice string
	for {
		fmt.Print("Enter the dividend: ")
		_, err := fmt.Scan(&dividend)
		if err != nil {
			fmt.Println("oops! couldn't carryout that operation..")
			return
		}
		fmt.Print("Enter thr divisor: ")
		_, errr := fmt.Scan(&divisor)
		if errr != nil {
			fmt.Println("oops! couldn't carryout that operation..")
			return
		}
		result, err := division(dividend, divisor)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("The quotient of %.2f and %.2f is.. %.2f\n", dividend, divisor, result)

		fmt.Println("#############")

		fmt.Print("Do you want to carry-out another operation (y/n): ")
		fmt.Scan(&choice)
		if choice == "y" || choice == "Y" {
			continue

		} else if choice == "n" || choice == "N" {
			fmt.Println("Exiting the program")
			break
		} else {
			fmt.Println("invalid command....")
			fmt.Println("Exiting program..")
			break

		}

	}
}
