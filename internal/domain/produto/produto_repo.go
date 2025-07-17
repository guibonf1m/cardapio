package produto

type ProdutoRepository interface {
	GetProduto(id int) (*Produto, error)
	GetProdutos() ([]Produto, error)
	GetProdutosFiltrados(status *Status, categoria *Categoria) ([]Produto, error)
	GetProdutosPorCategoria(categoria Categoria) ([]Produto, error)
	NomeExists(nome string) (bool, error)
	AddProduto(produto Produto) (Produto, error)
	UpdateProduto(produto Produto) error
	DeleteProduto(id int) error
}
