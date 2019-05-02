package main

import "fmt"

func main() {

	x:= 10
	y := 101

	if x > y {

		fmt.Printf("El valor %d es mayor que el valor %d \n", x, y)

	} else if y > x {
		fmt.Printf("El valor %d es mayor que el valor %d \n", y, x)
	} else {
		fmt.Printf("El valor %d es igual a el valor %d \n", x, y)
	}
}