package services

import c "github.com/guibonf1m/cardapio/internal/domain/cliente"

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
