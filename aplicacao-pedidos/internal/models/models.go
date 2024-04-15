package models

import (
    "time"
)

type Pedido struct {
    ID              int            `json:"id"`
    ClienteID        int            `json:"clienteID"`
    EntregaRua       string         `json:"entregaRua"`
    EntregaNumero   int            `json:"entregaNumero"`
    EntregaBairro    string         `json:"entregaBairro"`
    EntregaCidade    string         `json:"entregaCidade"`
    EntregaEstado    string         `json:"entregaEstado"`
    EntregaCep      string         `json:"entregaCep"`
    Produtos        string         `json:"produtos"` // Assuming Produtos is a JSON string (update if it's a custom struct)
    ValorTotal       float64        `json:"valorTotal"`
    ValorFrete      float64        `json:"valorFrete"`
    Status          string         `json:"status"`
    DataCriacao     time.Time      `json:"dataCriacao"`
}

type Produto struct {
    ID              int            `json:"id"`
    Nome            string         `json:"nome"`
    Descricao       string         `json:"descricao"`
    Preco           float64        `json:"preco"`
    Estoque          int            `json:"estoque"`
}

type Cliente struct {
    ID              int            `json:"id"`
    Nome            string         `json:"nome"`
    Email           string         `json:"email"`
    EnderecoRua     string         `json:"enderecoRua"`
    EnderecoNumero   int            `json:"enderecoNumero"`
    EnderecoBairro    string         `json:"enderecoBairro"`
    EnderecoCidade    string         `json:"enderecoCidade"`
    EnderecoEstado    string         `json:"enderecoEstado"`
    EnderecoCep      string         `json:"enderecoCep"`
}
