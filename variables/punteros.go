package main

import "fmt"

func main() {

	/*
	1- Es una direccion en memoria
	2- En lugar del valor guardamos la direccion de memoria que esta el valor
	3- *T => *int *string *Estructura
	4- El valor zero es nil
	*/
	var x,y *int

	entero:= 5

	x = &entero // el & entrega la direccion en memoria
	y = &entero

	*x = 6 //como Y apunta a la misma posicion que X, sera 6 igual y para asignar un valor hay que colocar *

	fmt.Println(x) //Imprime la posicion en memoria
	fmt.Println(*x) //EL * da acceso al valor que tiene la memoria
	fmt.Println(y)
	fmt.Println(*x)



}
