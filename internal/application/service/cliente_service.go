package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/guibonf1m/cardapio/internal/application/ports/services"
	c "github.com/guibonf1m/cardapio/internal/domain/cliente"

	"github.com/guibonf1m/cardapio/internal/application/ports/repositories"
)

var (
	ErrClienteNotFound = errors.New("cliente não encontrado")
	ErrClienteExists   = errors.New("cliente já existe")
	ErrInvalidCPF      = errors.New("CPF inválido")
	ErrInvalidPhone    = errors.New("telefone inválido")
)

type ClienteService struct {
	repo repositories.ClienteRepository
}

func NewClienteService(repo repositories.ClienteRepository) services.ClienteService {
	return &ClienteService{
		repo: repo,
	}
}

func (s *ClienteService) CriarCliente(nome string, telefone string, email string, endereco string) (*c.Cliente, error) {
	if err := s.validarDadosCliente(nome, telefone, email); err != nil {
		return nil, err
	}

	if clienteExistente, _ := s.repo.FindByTelefone(context.Background(), telefone); clienteExistente != nil {
		return nil, fmt.Errorf("cliente com telefone %s já existe", telefone)
	}

	novoCliente := c.NewCliente(nome, telefone, email, endereco)

	if err := s.repo.Create(context.Background(), novoCliente); err != nil {
		return nil, fmt.Errorf("erro ao salvar cliente: %w", err)
	}

	return novoCliente, nil
}

func (s *ClienteService) validarDadosCliente(nome, telefone, email string) error {
	if nome == "" {
		return errors.New("nome é obrigatório")
	}

	if telefone == "" {
		return errors.New("telefone é obrigatório")
	}

	if err := s.validarTelefone(telefone); err != nil {
		return err
	}

	if email != "" {
		if err := s.validarEmail(email); err != nil {
			return err
		}
	}

	return nil
}

func (s *ClienteService) validarTelefone(telefone string) error {
	// Validação básica de telefone brasileiro
	if len(telefone) < 10 || len(telefone) > 11 {
		return ErrInvalidPhone
	}
	return nil
}

func (s *ClienteService) validarEmail(email string) error {
	// Validação básica de email
	if email != "" && len(email) < 5 {
		return errors.New("email inválido")
	}
	return nil
}
