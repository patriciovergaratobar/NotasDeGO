package main

import "fmt"

func main() {

	arreglo1 := [3]int{16,20,3}
	
	fmt.Println(arreglo1)

	for i:=0; i < len(arreglo1); i++ {
		fmt.Println(arreglo1[i])
	}
}