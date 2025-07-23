package cliente

import "time"

type Status string

const (
	StatusAtivo   Status = "ativo"
	StatusInativo Status = "inativo"
)

type Cliente struct {
	ID        int
	Nome      string
	Telefone  string
	Email     string
	Endereco  string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *int
	UpdatedBy *int
}

func NewCliente(nome, telefone, email, endereco string) *Cliente {
	return &Cliente{
		Nome:      nome,
		Telefone:  telefone,
		Email:     email,
		Endereco:  endereco,
		Status:    StatusAtivo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Métodos de domínio
func (c *Cliente) Ativar() {
	c.Status = StatusAtivo
	c.UpdatedAt = time.Now()
}

func (c *Cliente) Desativar() {
	c.Status = StatusInativo
	c.UpdatedAt = time.Now()
}

func (c *Cliente) IsAtivo() bool {
	return c.Status == StatusAtivo
}

func (c *Cliente) AtualizarDados(nome, telefone, email, endereco string) {
	c.Nome = nome
	c.Telefone = telefone
	c.Email = email
	c.Endereco = endereco
	c.UpdatedAt = time.Now()
}
