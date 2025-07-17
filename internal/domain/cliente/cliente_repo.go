package cliente

type ClienteRepository interface {
	GetCliente(id int) (*Cliente, error)
	GetClientes() ([]Cliente, error)
	GetClientesFiltrados(status *Status) ([]Cliente, error)
	GetClientePorTelefone(telefone string) (*Cliente, error)
	GetClientePorEmail(email string) (*Cliente, error)
	TelefoneExists(telefone string) (bool, error)
	EmailExists(email string) (bool, error)
	AddCliente(cliente Cliente) (Cliente, error)
	UpdateCliente(cliente Cliente) error
	DeleteCliente(id int) error
}
