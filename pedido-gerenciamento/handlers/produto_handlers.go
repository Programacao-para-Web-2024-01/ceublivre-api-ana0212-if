package handlers

import (
    "encoding/json"
    "net/http"
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/repositories"
)

func AdicionarProduto(w http.ResponseWriter, r *http.Request) {
    var produto models.Produto
    err := json.NewDecoder(r.Body).Decode(&produto)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = repositories.SalvarProduto(produto)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(produto)
}

func ListarProdutos(w http.ResponseWriter, r *http.Request) {
    produtos := repositories.ListarProdutos()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(produtos)
}
