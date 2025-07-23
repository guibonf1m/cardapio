package services

import pr "github.com/guibonf1m/cardapio/internal/domain/produto"

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
