package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

	"caminho/para/seu/pacote/controllers" // Substituir em baixo
	"caminho/para/seu/pacote/repositories" // Substituir em baixo

    // Modelos dos dados
    "./models"

    // Repositórios
    "./repositories"

    // Serviços
    "./services"

    // Controladores
    "./controllers"
)

// Função para conectar ao banco de dados MySQL
func conectarBancoDeDados() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:Rootroot+@tcp(localhost:3306)/banco_pedidos")
    if err != nil {
        return nil, fmt.Errorf("erro ao conectar no banco dedados: %w", err)
    }

    return db, nil
}

func main() {
    // Conexão com o banco de dados
    db, err := conectarBancoDeDados()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() // Fecha a conexão ao final da execução

    // Repositórios
    pedidoRepository := repositories.NovoPedidoRepository(db)
    produtoRepository := repositories.NovoProdutoRepository(db)
    clienteRepository := repositories.NovoClienteRepository(db)

    // Serviços
    emailService := services.NovoEmailService()
    estoqueService := services.NovoEstoqueService(produtoRepository)

    // Controladores
    pedidoController := controllers.NovoPedidoController(pedidoRepository, emailService, estoqueService)

    // Registrar rotas
    router := http.NewServeMux()
    router.HandleFunc("/pedidos", pedidoController.CriarPedido)
    router.HandleFunc("/pedidos/{pedidoID}", pedidoController.AtualizarPedido)
    // Registrar outras rotas para outras funcionalidades

    // Iniciar servidor web
    log.Println("Servidor escutando na porta 8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Erro ao iniciar servidor: %v", err)
    }
}
