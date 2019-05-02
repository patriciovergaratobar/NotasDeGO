package main

import "fmt"

func main() {

	slice := make([]int, 2, 5) // Se contruye un arreglo con 2 posiciones, pero tiene una capacidad de 5

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(append(slice, 3,4,6)) // Aqui se genera un slice apuntando a la posicion de memoria del slice y respetando su capacidad
	fmt.Println(append(slice, 3,4,6, 33,22,44)) // Aqui ya no se respeta la capacidad y todo se pone menos eficiente ya que el append tiene que generar otro slice con mas capacidad en memoria.

	origen := []int {2,3,4}
	destino := make([]int, 2)

	copy(destino, origen) 
	fmt.Println(destino) // Se copia el minimo ejemplo se copian la minima cantidad definida en destino

	destino2 := make([]int, len(origen))
	copy(destino2, origen) 
	fmt.Println(destino2)

}