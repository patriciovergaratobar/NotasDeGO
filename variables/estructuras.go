package main

import "fmt"

type User struct {
	edad int
	nombre, apellido string
}

//De esta forma se crean los metodos para las estructura, 
//siempre y cuando estas se encuentren en el mismo package.
func(u User) nombreCompleto() string {
	return u.nombre + " " + u.apellido
}

func(u User) set_nombre(n string) {
	u.nombre = n;
}

func(u *User) set_nombreConPuntero(n string) {
	u.nombre = n;
}


func main() {

	patricio := User{edad:29, nombre:"Patricio", apellido:"Vergara"}

	user2 := new(User)

	user2.edad = 27
	user2.nombre = "zen"

	fmt.Println(patricio)
	fmt.Println(patricio.nombreCompleto())

	patricio.set_nombre("Pato")
	fmt.Println(patricio.nombre) // Aqui sigue mostrando el mismo nombre porque cuando se ejecuta la funcion esta genera una copia del objeto con el nuevo nombre


	patricio.set_nombreConPuntero("Pato")
	fmt.Println(patricio.nombre) // Aqui se muestra el nuevo nombre por que el cambio de valor se hace en la misma posicion de memoria

	fmt.Println(user2)
	fmt.Println((*user2))
	fmt.Println(user2.nombre)
	fmt.Println((*user2).nombre)
}
