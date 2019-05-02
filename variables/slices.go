package main

import "fmt"

func main() {

	arregloSlice := []int{1,3,6,2,3,5}

	arregloNormal := [3]int{1,3,6}
	slice := arregloNormal[0:] // pasar un arreglo a un slice

	fmt.Println(slice)

	fmt.Println(arregloSlice)

	arregloSlice = append(arregloSlice, 100, 90, 87)

	fmt.Println(arregloSlice)

}