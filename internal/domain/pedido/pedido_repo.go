package pedido

type PedidoRepository interface {
	GetPedido(id int) (*Pedido, error)
	GetPedidos() ([]Pedido, error)
	GetPedidosPorCliente(clienteID int) ([]Pedido, error)
	GetPedidosPorStatus(status Status) ([]Pedido, error)
	AddPedido(pedido Pedido) (Pedido, error)
	UpdateStatusPedido(pedidoID int, status Status) error
	DeletePedido(id int) error
}
