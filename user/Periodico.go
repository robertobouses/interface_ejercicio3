package user

type Periodico struct {
	Titulo    string
	Autor     string
	Idiologia string
}

func (P Periodico) ImprimirTituloAutor() (string, string) {
	return P.Titulo, P.Autor
}
