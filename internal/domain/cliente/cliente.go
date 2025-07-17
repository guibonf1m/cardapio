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
	Telefone  string // Campo principal (WhatsApp)
	Email     string // Para segurança e comunicação
	Endereco  string // Para delivery
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *int
	UpdatedBy *int
}
