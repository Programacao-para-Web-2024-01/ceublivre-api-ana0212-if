package repositories

import (
    "errors"
    "pedido-gerenciamento/models"
)

var pedidos = make(map[string]models.Pedido)

func SalvarPedido(pedido models.Pedido) error {
    pedidos[pedido.ID] = pedido
    return nil
}

func BuscarPedido(id string) (models.Pedido, error) {
    pedido, exists := pedidos[id]
    if !exists {
        return models.Pedido{}, errors.New("pedido não encontrado")
    }
    return pedido, nil
}

func AtualizarPedido(pedido models.Pedido) error {
    pedidos[pedido.ID] = pedido
    return nil
}

func ListarPedidos() []models.Pedido {
    var lista []models.Pedido
    for _, pedido := range pedidos {
        lista = append(lista, pedido)
    }
    return lista
}

func AtualizarStatusPedido(id string, status string) error {
    pedido, exists := pedidos[id]
    if !exists {
        return errors.New("pedido não encontrado")
    }
    pedido.Status = status
    pedidos[id] = pedido
    return nil
}

func AtualizarEnderecoPedido(id string, endereco string) error {
    pedido, exists := pedidos[id]
    if !exists {
        return errors.New("pedido não encontrado")
    }
    pedido.Endereco = endereco
    pedidos[id] = pedido
    return nil
}
