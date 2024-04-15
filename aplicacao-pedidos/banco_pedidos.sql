CREATE TABLE pedidos (
  id INT PRIMARY KEY AUTO_INCREMENT,
  cliente_id INT NOT NULL,
  entrega_rua VARCHAR(255) NOT NULL,
  entrega_numero INT NOT NULL,
  entrega_bairro VARCHAR(255) NOT NULL,
  entrega_cidade VARCHAR(255) NOT NULL,
  entrega_estado VARCHAR(255) NOT NULL,
  entrega_cep VARCHAR(255) NOT NULL,
  produtos JSON NOT NULL,
  valor_total DECIMAL(10,2) NOT NULL,
  valor_frete DECIMAL(10,2) NOT NULL,
  status VARCHAR(255) NOT NULL DEFAULT 'pendente',
  data_criacao DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE produtos (
  id INT PRIMARY KEY AUTO_INCREMENT,
  nome VARCHAR(255) NOT NULL,
  descricao TEXT NOT NULL,
  preco DECIMAL(10,2) NOT NULL,
  estoque INT NOT NULL
);

CREATE TABLE clientes (
  id INT PRIMARY KEY AUTO_INCREMENT,
  nome VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  endereco_rua VARCHAR(255) NOT NULL,
  endereco_numero INT NOT NULL,
  endereco_bairro VARCHAR(255) NOT NULL,
  endereco_cidade VARCHAR(255) NOT NULL,
  endereco_estado VARCHAR(255) NOT NULL,
  endereco_cep VARCHAR(255) NOT NULL
);
