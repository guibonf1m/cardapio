package dtos

type CriarPedidoRequest struct {
	ClienteID   int                 `json:"cliente_id" validate:"required"`
	Itens       []ItemPedidoRequest `json:"itens" validate:"required,min=1"`
	Observacoes string              `json:"observacoes,omitempty" validate:"max=300"`
}
type ItemPedidoRequest struct {
	ProdutoID   int    `json:"produto_id" validate:"required"`
	Quantidade  int    `json:"quantidade" validate:"required,min=1,max=20"`
	Observacoes string `json:"observacoes,omitempty" validate:"max=100"`
}
type AtualizarStatusPedidoRequest struct {
	Status string `json:"status" validate:"required,oneof=pendente confirmado preparo pronto entregue cancelado"`
}
type AdicionarObservacaoRequest struct {
	Observacao string `json:"observacao" validate:"required,max=300"`
}

type PedidoResponse struct {
	ID          int                  `json:"id"`
	ClienteID   int                  `json:"cliente_id"`
	Cliente     ClienteResponse      `json:"cliente,omitempty"`
	Status      string               `json:"status"`
	Total       float64              `json:"total"`
	Observacoes string               `json:"observacoes"`
	Itens       []ItemPedidoResponse `json:"itens,omitempty"`
	CreatedAt   string               `json:"created_at"`
	UpdatedAt   string               `json:"updated_at"`
}

type ItemPedidoResponse struct {
	ID            int             `json:"id"`
	ProdutoID     int             `json:"produto_id"`
	Produto       ProdutoResponse `json:"produto,omitempty"`
	Quantidade    int             `json:"quantidade"`
	PrecoUnitario float64         `json:"preco_unitario"`
	Subtotal      float64         `json:"subtotal"`
	Observacoes   string          `json:"observacoes"`
}
