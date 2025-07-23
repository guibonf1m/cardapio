package database

import (
	"fmt"
	"github.com/guibonf1m/cardapio/internal/domain/cliente"
	"gorm.io/gorm"
	"time"
)

type ClienteModel struct {
	ID        uint   `gorm:"primaryKey"`
	Nome      string `gorm:"not null"`
	Telefone  string `gorm:"uniqueIndex;not null"`
	Email     string
	Endereco  string
	Status    string         `gorm:"default:ativo"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedBy *int
	UpdatedBy *int
}

// Conversão de domínio para model
func (m *ClienteModel) FromDomain(c *cliente.Cliente) {
	m.Nome = c.Nome
	m.Telefone = c.Telefone
	m.Email = c.Email
	m.Endereco = c.Endereco
	m.Status = string(c.Status)
	// ID, timestamps são gerenciados pelo GORM
}

// Conversão de model para domínio
func (m *ClienteModel) ToDomain() *cliente.Cliente {
	return &cliente.Cliente{
		ID:        fmt.Sprintf("%d", m.ID),
		Nome:      m.Nome,
		Telefone:  m.Telefone,
		Email:     m.Email,
		Endereco:  m.Endereco,
		Status:    cliente.Status(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
	}
}
