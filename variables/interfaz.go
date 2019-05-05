package main

import (
	"fmt"
)

type User interface {
	permisos() int
	getNombre() string
}

type Admin struct {

	nombre string
}

func (a Admin) permisos() int {
	return 5;
}

func (a Admin) getNombre() string {
	return a.nombre;
}

type Visita struct {

	nombre string
}

func (v Visita) permisos() int {
	return 1;
}

func (v Visita) getNombre() string {
	return v.nombre;
}

func login(user User) string {

	if user.permisos() > 4 {
		return user.getNombre()
	} else {
		return "No tiene permiso"
	}
}


func main() {

	fmt.Println("Inicio")

	admin := Admin{"Administrador"}

	visita := Visita{"Administrador"}

	fmt.Println(login(admin))

	fmt.Println(login(visita))
}