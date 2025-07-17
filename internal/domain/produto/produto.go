package produto

import "time"

type Status string

const (
	StatusDisponivel   Status = "disponivel"
	StatusIndisponivel Status = "indisponivel"
)

type Categoria string

const (
	CategoriaMarmita    Categoria = "marmita"
	CategoriaBebida     Categoria = "bebida"
	CategoriaPorcao     Categoria = "porcao"
	CategoriaHamburguer Categoria = "hamburguer"
	CategoriaSalgado    Categoria = "salgado"
	CategoriaBatata     Categoria = "batata-frita"
)

type Produto struct {
	ID        int
	Nome      string
	Descricao string
	Preco     float64
	Categoria Categoria
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *int
	UpdatedBy *int
}
