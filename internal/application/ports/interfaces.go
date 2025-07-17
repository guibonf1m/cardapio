package ports

import (
	c "github.com/guibonf1m/cardapio/internal/domain/cliente"
	pe "github.com/guibonf1m/cardapio/internal/domain/pedido"
	pr "github.com/guibonf1m/cardapio/internal/domain/produto"
)

// ============ TIPOS DO CORE (sem JSON tags) ============
type ItemPedidoData struct {
	ProdutoID   int
	Quantidade  int
	Observacoes string
}

type ProdutoService interface {
	CriarProduto(nome, descricao string, preco float64, categoria pr.Categoria) (*pr.Produto, error)
	ListarProdutos() ([]pr.Produto, error)
	ListarProdutosDisponiveis() ([]pr.Produto, error)
	ListarPorCategoria(categoria pr.Categoria) ([]pr.Produto, error)
	BuscarProduto(id int) (*pr.Produto, error)
	AtualizarProduto(produto *pr.Categoria) (*pr.Produto, error)
	AlterarStatusProduto(id int, status pr.Status) error
	DeletarProduto(id int) error
	VerificarDisponibilidade(id int) (bool, error)
}

type ClienteService interface {
	CriarCliente(nome, telefone, email, endereco string) (*c.Cliente, error)
	ListarClientes() ([]c.Cliente, error)
	ListarClientesAtivos() ([]c.Cliente, error)
	BuscarCliente(id int) (*c.Cliente, error)
	BuscarClientePorTelefone(telefone string) (*c.Cliente, error)
	BuscarClientePorEmail(email string) (*c.Cliente, error)
	AtualizarCliente(cliente c.Cliente) error
	AlterarStatusCliente(id int, status c.Status) error
	DeletarCliente(id int) error
	ValidarCliente(telefone, email string) error
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
