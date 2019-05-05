package main

import (
	"fmt"
)

type Humano struct {

	nombre string
}

func(h Humano) hablar() string {

	return "bla bla bla"
}

func(h Humano) imprimirNombre() string {

	return h.nombre
}

type Tutor struct {

	Humano
}

func(t Tutor) imprimirNombre() string {

	return "Oculto =0"
}


func main() {

	fmt.Println("Inicio")

	tutor := Tutor{Humano{"Patricio"}}

	
	fmt.Println(tutor.nombre)

	//Llamando metodo heredado
	fmt.Println(tutor.hablar())

	//Llamando a metodo sobre escrito
	fmt.Println(tutor.imprimirNombre())
}