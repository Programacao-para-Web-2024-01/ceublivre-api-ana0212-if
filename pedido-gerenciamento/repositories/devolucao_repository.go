package repositories

import (
    "pedido-gerenciamento/models"
    "errors"
)

var devolucoes = []models.Devolucao{}

func SalvarDevolucao(devolucao models.Devolucao) error {
    devolucoes = append(devolucoes, devolucao)
    return nil
}

func ListarDevolucoes() []models.Devolucao {
    return devolucoes
}

func BuscarDevolucao(id string) (models.Devolucao, error) {
    for _, devolucao := range devolucoes {
        if devolucao.ID == id {
            return devolucao, nil
        }
    }
    return models.Devolucao{}, errors.New("devolução não encontrada")
}

func AtualizarDevolucao(devolucaoAtualizada models.Devolucao) error {
    for i, devolucao := range devolucoes {
        if devolucao.ID == devolucaoAtualizada.ID {
            devolucoes[i] = devolucaoAtualizada
            return nil
        }
    }
    return errors.New("devolução não encontrada")
}
