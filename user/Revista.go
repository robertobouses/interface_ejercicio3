package user

type Revista struct {
	Titulo string
	Autor  string
	Codigo int
}

func (R Revista) ImprimirTituloAutor() (string, string) {
	return R.Titulo, R.Autor
}
