package user

type Novela struct {
	Titulo   string
	Autor    string
	Tematica string
}

func (N Novela) ImprimirTituloAutor() (string, string) {
	return N.Titulo, N.Autor
}
