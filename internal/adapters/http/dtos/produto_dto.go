package dtos

type CriarProdutoRequest struct {
	Nome      string  `json:"nome" validate:"required,min=3,max=100"`
	Descricao string  `json:"descricao,omitempty" validate:"max=500"`
	Preco     float64 `json:"preco" validate:"required,gt=0"`
	Categoria string  `json:"categoria" validate:"required"`
}
type AtualizarProdutoRequest struct {
	Nome      *string  `json:"nome,omitempty" validate:"omitempty,min=3,max=100"`
	Descricao *string  `json:"descricao,omitempty" validate:"omitempty,max=500"`
	Preco     *float64 `json:"preco,omitempty" validate:"omitempty,gt=0"`
	Categoria *string  `json:"categoria,omitempty" validate:"omitempty"`
}
type AlterarStatusProdutoRequest struct {
	Status string `json:"status" validate:"required,oneof=disponivel indisponivel"`
}

type ProdutoResponse struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Categoria string  `json:"categoria"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
