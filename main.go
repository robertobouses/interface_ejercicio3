package main

import "fmt"

type Organizar interface {
	ImprimirTituloAutor() (string, string)
}

type Novela struct {
	titulo   string
	autor    string
	tematica string
}

type Periodico struct {
	titulo    string
	autor     string
	idiologia string
}

type Revista struct {
	titulo string
	autor  string
	codigo int
}

func (N Novela) ImprimirTituloAutor() (string, string) {
	return N.titulo, N.autor
}

func (P Periodico) ImprimirTituloAutor() (string, string) {
	return P.titulo, P.autor
}
func (R Revista) ImprimirTituloAutor() (string, string) {
	return R.titulo, R.autor
}

func Datos(O Organizar) {
	titulo, autor := O.ImprimirTituloAutor()
	fmt.Printf("Título: %s\nAutor: %s\n\n", titulo, autor)
}

func main() {
	fmt.Println("hola les presento la biblioteca chachi")
	Novela1 := Novela{"El Quijote", "Cervantes", "picaresca"}
	Novela2 := Novela{"Tokio Blues", "Murakami", "romántica"}

	Periodico1 := Periodico{"La Razón", "Luís Mari Anson", "anormal"}
	Revista1 := Revista{"AR", "Doña Ama Rosa Quintana", 69}
	fmt.Print("mensaje subliminal: Anna te quiero")
	Datos(Novela1)
	Datos(Novela2)
	Datos(Periodico1)
	Datos(Revista1)

}
