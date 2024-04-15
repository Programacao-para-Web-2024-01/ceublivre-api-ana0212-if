package controllers

import (
	"encoding/json"
	"net/http"

	"caminho/para/seu/pacote/services" // Substitua pelo caminho real
	"caminho/para/seu/pacote/models"    // Substitua pelo caminho real
)

type PedidoController interface {
	CreatePedido(w http.ResponseWriter, r *http.Request)
}

type pedidoControllerImpl struct {
	service services.PedidoService
}

func NovoPedidoController(service services.PedidoService) PedidoController {
	return &pedidoControllerImpl{service: service}
}

func (c *pedidoControllerImpl) CreatePedido(w http.ResponseWriter, r *http.Request) {
	// Analise o corpo da requisição para desempacotar os dados do Pedido
	var pedido models.Pedido
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		// Trate o erro (por exemplo, escreva uma resposta de erro)
		return
	}

	// Chame o serviço para criar o pedido
	err = c.service.CriarPedido(&pedido)
	if err != != nil {
		// Trate o erro (por exemplo, escreva uma resposta de erro)
		return
	}

	// Escreva uma resposta de sucesso com os dados do Pedido criado (opcional)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pedido)
}
