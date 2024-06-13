package handlers

import (
    "encoding/json"
    "net/http"
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/services"
    "github.com/gorilla/mux"
)

func CriarDevolucao(w http.ResponseWriter, r *http.Request) {
    var devolucao models.Devolucao
    _ = json.NewDecoder(r.Body).Decode(&devolucao)
    devolucao, err := services.CriarDevolucao(devolucao)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(devolucao)
}

func ListarDevolucoes(w http.ResponseWriter, r *http.Request) {
    devolucoes := services.ListarDevolucoes()
    json.NewEncoder(w).Encode(devolucoes)
}

func AtualizarDevolucao(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var devolucao models.Devolucao
    _ = json.NewDecoder(r.Body).Decode(&devolucao)
    devolucao, err := services.AtualizarDevolucao(id, devolucao)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(devolucao)
}
