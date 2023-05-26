package main

import (
	"fmt"

	"github.com/robertobouses/interface_ejercicio3/user"
)

type Organizar interface {
	ImprimirTituloAutor() (string, string)
}

func Datos(O Organizar) {
	titulo, autor := O.ImprimirTituloAutor()
	fmt.Printf("Título: %s\nAutor: %s\n\n", titulo, autor)
}

func main() {
	fmt.Println("hola les presento la biblioteca chachi")
	Novela1 := user.Novela{"El Quijote", "Cervantes", "picaresca"}
	Novela2 := user.Novela{"Tokio Blues", "Murakami", "romántica"}

	Periodico1 := user.Periodico{"Marca", "Paco Goles", "Deportes"}
	Revista1 := user.Revista{"AR", "Doña Ama Rosa Quintana", 69}
	fmt.Print("mensaje subliminal:")
	Datos(Novela1)
	Datos(Novela2)
	Datos(Periodico1)
	Datos(Revista1)

}
