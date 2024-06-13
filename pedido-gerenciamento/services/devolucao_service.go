package services

import (
    "pedido-gerenciamento/models"
    "pedido-gerenciamento/repositories"
    "strconv"
    "time"
    "errors"
)

func CriarDevolucao(devolucao models.Devolucao) (models.Devolucao, error) {
    devolucao.ID = "devolucao_" + strconv.FormatInt(time.Now().Unix(), 10)
    devolucao.DataCriacao = time.Now().Format("2006-01-02 15:04:05")
    devolucao.Status = "Pendente"
    err := repositories.SalvarDevolucao(devolucao)
    if err != nil {
        return devolucao, err
    }
    return devolucao, nil
}

func ListarDevolucoes() []models.Devolucao {
    return repositories.ListarDevolucoes()
}

func AtualizarDevolucao(id string, devolucaoAtualizada models.Devolucao) (models.Devolucao, error) {
    devolucao, err := repositories.BuscarDevolucao(id)
    if err != nil {
        return devolucao, err
    }
    if devolucao.ID == "" {
        return devolucao, errors.New("devolução não encontrada")
    }
    devolucao.Status = devolucaoAtualizada.Status
    err = repositories.AtualizarDevolucao(devolucao)
    if err != nil {
        return devolucao, err
    }
    return devolucao, nil
}
