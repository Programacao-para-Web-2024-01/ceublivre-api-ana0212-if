package repositories

import (
    "pedido-gerenciamento/models"
    "errors"
    "strconv"
)

var produtos = []models.Produto{}

func SalvarProduto(produto models.Produto) error {
    for _, p := range produtos {
        if p.Nome == produto.Nome {
            return errors.New("produto já existe")
        }
    }
    produto.ID = "prod_" + strconv.Itoa(len(produtos)+1)
    produtos = append(produtos, produto)
    return nil
}

func ListarProdutos() []models.Produto {
    return produtos
}

func BuscarProduto(nome string) (models.Produto, error) {
    for _, produto := range produtos {
        if produto.Nome == nome {
            return produto, nil
        }
    }
    return models.Produto{}, errors.New("produto não encontrado")
}

func AtualizarEstoque(produtoNome string, quantidade int) error {
    for i, produto := range produtos {
        if produto.Nome == produtoNome {
            produtos[i].Quantidade += quantidade
            return nil
        }
    }
    return errors.New("produto não encontrado")
}
