package main

import (
	"fmt"
	"strconv"
)

func main() {

	edad := 28
	//String a entero
	//masEdad, err := strconv.Atoi("1") 
	masEdad, _ := strconv.Atoi("1")  //Esta variable retorna 2 valores: valor transformado y error si es que hay
	//Con el _ podemos omitir uno de los valores.
	edad = edad + masEdad
	edad_str := strconv.Itoa(edad) //Entero a string
	fmt.Println("Mi edad es "+edad_str)
}
