package repositories

import (
	"context"
	"github.com/guibonf1m/cardapio/internal/domain/cliente"
)

type ClienteRepository interface {
	Create(ctx context.Context, c *cliente.Cliente) error
	FindByID(ctx context.Context, id string) (*cliente.Cliente, error)
	FindByCPF(ctx context.Context, cpf string) (*cliente.Cliente, error)
	FindByTelefone(ctx context.Context, telefone string) (*cliente.Cliente, error)
	Update(ctx context.Context, c *cliente.Cliente) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filters map[string]interface{}) ([]*cliente.Cliente, error)
}
