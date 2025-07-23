package services

import pe "github.com/guibonf1m/cardapio/internal/domain/pedido"

type ItemPedidoData struct {
	ProdutoID   int
	Quantidade  int
	Observacoes string
}

type PedidoService interface {
	CriarPedido(clienteID int, itens []ItemPedidoData, observacoes string) (*pe.Pedido, error)
	ListarPedidos() ([]pe.Pedido, error)
	ListarPedidosPorCliente(clienteID int) ([]pe.Pedido, error)
	ListarPedidosPorStatus(status pe.Status) ([]pe.Pedido, error)
	BuscarPedido(id int) (*pe.Pedido, error)
	BuscarPedidoCompleto(id int) (*pe.Pedido, error)
	AtualizarStatusPedido(id int, status pe.Status) error
	CancelarPedido(id int, motivo string) error
	CalcularTotalPedido(itens []ItemPedidoData) (float64, error)
	AdicionarObservacao(id int, observacao string) error
}
