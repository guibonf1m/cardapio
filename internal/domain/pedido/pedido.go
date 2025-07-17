package pedido

import "time"

type Status string

const (
	StatusPendente   Status = "pendente"
	StatusConfirmado Status = "confirmado"
	StatusPreparo    Status = "preparo"
	StatusPronto     Status = "pronto"
	StatusEntregue   Status = "entregue"
	StatusCancelado  Status = "cancelado"
)

type Pedido struct {
	ID          int
	ClienteID   int
	Status      Status
	Total       float64
	Observacoes string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   *int
	UpdatedBy   *int
	Itens       []ItemPedido `json:"itens,omitempty"`
}
type ItemPedido struct {
	ID            int
	PedidoID      int
	ProdutoID     int
	Quantidade    int
	PrecoUnitario float64
	Subtotal      float64
	Observacoes   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
