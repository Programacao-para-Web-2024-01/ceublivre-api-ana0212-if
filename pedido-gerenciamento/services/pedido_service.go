package services

import (
    "errors"
    "log"
    "net/url"
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/repositories"
    "pedido-gerenciamento/utils"
    "strconv"
    "time"
)

func CriarPedido(pedido models.Pedido) (models.Pedido, error) {
    pedido.ID = "pedido_" + strconv.FormatInt(time.Now().Unix(), 10)
    pedido.DataCriacao = time.Now().Format("2006-01-02 15:04:05")
    pedido.Status = "Pendente"
    
    log.Printf("Criando pedido: %+v", pedido)
    
    for _, produtoNome := range pedido.Produtos {
        produto, err := repositories.BuscarProduto(produtoNome)
        if err != nil {
            log.Printf("Erro ao buscar produto '%s': %v", produtoNome, err)
            return pedido, err
        }
        if produto.Quantidade <= 0 {
            log.Printf("Produto esgotado: %s", produtoNome)
            return pedido, errors.New("produto esgotado: " + produtoNome)
        }
        produto.Quantidade -= 1
        err = repositories.AtualizarEstoque(produtoNome, -1) // Reduz uma unidade do estoque
        if err != nil {
            log.Printf("Erro ao atualizar estoque do produto '%s': %v", produtoNome, err)
            return pedido, err
        }
    }

    err := repositories.SalvarPedido(pedido)
    if err != nil {
        log.Printf("Erro ao salvar pedido: %v", err)
        return pedido, err
    }

    // Notificar cliente via email
    utils.Notify(pedido.Cliente, "Pedido criado com sucesso")
    log.Printf("Pedido criado com sucesso: %+v", pedido)
    return pedido, nil
}

func AtualizarPedido(id string, pedidoAtualizado models.Pedido) (models.Pedido, error) {
    pedido, err := repositories.BuscarPedido(id)
    if err != nil {
        log.Printf("Erro ao buscar pedido '%s': %v", id, err)
        return pedido, err
    }
    if pedido.ID == "" {
        log.Printf("Pedido não encontrado: %s", id)
        return pedido, errors.New("pedido não encontrado")
    }

    pedido.Endereco = pedidoAtualizado.Endereco
    pedido.Status = pedidoAtualizado.Status

    err = repositories.AtualizarPedido(pedido)
    if err != nil {
        log.Printf("Erro ao atualizar pedido '%s': %v", id, err)
        return pedido, err
    }

    // Notificar cliente via email
    utils.Notify(pedido.Cliente, "Pedido atualizado para status: " + pedido.Status)
    log.Printf("Pedido atualizado com sucesso: %+v", pedido)
    return pedido, nil
}

func AtualizarStatusPedido(id string, status string) error {
    pedido, err := repositories.BuscarPedido(id)
    if err != nil {
        log.Printf("Erro ao buscar pedido '%s': %v", id, err)
        return err
    }
    if pedido.ID == "" {
        log.Printf("Pedido não encontrado: %s", id)
        return errors.New("pedido não encontrado")
    }

    pedido.Status = status

    err = repositories.AtualizarPedido(pedido)
    if err != nil {
        log.Printf("Erro ao atualizar status do pedido '%s': %v", id, err)
        return err
    }

    // Notificar cliente via email
    utils.Notify(pedido.Cliente, "Status do pedido atualizado para: " + pedido.Status)
    log.Printf("Status do pedido atualizado com sucesso: %+v", pedido)
    return nil
}

func AtualizarEnderecoPedido(id string, endereco string) error {
    pedido, err := repositories.BuscarPedido(id)
    if err != nil {
        log.Printf("Erro ao buscar pedido '%s': %v", id, err)
        return err
    }
    if pedido.ID == "" {
        log.Printf("Pedido não encontrado: %s", id)
        return errors.New("pedido não encontrado")
    }

    pedido.Endereco = endereco

    err = repositories.AtualizarPedido(pedido)
    if err != nil {
        log.Printf("Erro ao atualizar endereço do pedido '%s': %v", id, err)
        return err
    }

    // Notificar cliente via email
    utils.Notify(pedido.Cliente, "Endereço do pedido atualizado para: " + pedido.Endereco)
    log.Printf("Endereço do pedido atualizado com sucesso: %+v", pedido)
    return nil
}

func ObterPedido(id string) (models.Pedido, error) {
    pedido, err := repositories.BuscarPedido(id)
    if err != nil {
        log.Printf("Erro ao buscar pedido '%s': %v", id, err)
        return models.Pedido{}, err
    }
    if pedido.ID == "" {
        log.Printf("Pedido não encontrado: %s", id)
        return models.Pedido{}, errors.New("pedido não encontrado")
    }
    return pedido, nil
}

func ListarPedidos() []models.Pedido {
    pedidos := repositories.ListarPedidos()
    log.Printf("Pedidos listados: %+v", pedidos)
    return pedidos
}

func GerarRelatorio(filtros url.Values) []models.Pedido {
    pedidos := repositories.ListarPedidos()
    var pedidosFiltrados []models.Pedido
    for _, pedido := range pedidos {
        // Adicione lógica de filtragem aqui
        pedidosFiltrados = append(pedidosFiltrados, pedido)
    }
    log.Printf("Relatório gerado com filtros %+v: %+v", filtros, pedidosFiltrados)
    return pedidosFiltrados
}
