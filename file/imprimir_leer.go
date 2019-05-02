package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var edad int
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d\n",&edad)
	fmt.Printf("Tu edad es %d\n", edad)

	//Otra forma de leer con budio y os
	leer := bufio.NewReader(os.Stdin) // Indicamos de donde leeremos con os.Stdin (Teclado)
	fmt.Print("Ingresa tu nombre: ")
	// Con el \n estamos diciendo que se leera hasta que haya un salto de linea
	text, err := leer.ReadString('\n')
	if  err != nil {
		fmt.Println("Error al leer. ", err)
	} else {

		println("El nombre es " + text)
	}


}