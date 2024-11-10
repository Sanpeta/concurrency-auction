# Concurrency Auction

Este projeto é uma aplicação de leilão concorrente construída com Go e MongoDB. A aplicação permite criar leilões, fazer lances e gerenciar usuários.

## Estrutura do Projeto

- **cmd/auction**: Contém o ponto de entrada da aplicação.
  - `main.go`: Carrega variáveis de ambiente, conecta ao MongoDB, inicializa controladores e define rotas da API.
- **config**: Configurações da aplicação.
  - `database/mongodb/connection.go`: Configura a conexão com o MongoDB.
  - `logger/logger.go`: Configura o logger da aplicação.
  - `rest_err/rest_err.go`: Define a estrutura de erros REST.
- **internal**: Contém a lógica interna do projeto, dividida em várias subpastas.
  - `entity/`: Definições das entidades do domínio.
    - Subpastas: `auction_entity/`, `bid_entity/`, `user_entity/`.
  - `infra/`: Infraestrutura do projeto.
    - Subpastas:
      - `api/web/`: Controladores da API web.
        - Exemplo: `controller/auction_controller/`, `controller/bid_controller/`, `controller/user_controller/`.
      - `database/`: Implementações de acesso ao banco de dados.
  - `internal_error/`: Definições de erros internos.
    - Exemplo: `internal_error.go`.
  - `usecase/`: Casos de uso da aplicação.
    - Subpastas: `auction_usecase/`, `bid_usecase/`, `user_usecase/`.

## Pré-requisitos

- Go 1.16+
- Docker
- Docker Compose

## Configuração

1. Clone o repositório:

   ```sh
   git clone https://github.com/Sanpeta/concurrency-auction.git
   cd concurrency-auction
   ```

2. Crie um arquivo `app.env` em `./` com as seguintes variáveis:

   ```env
   MONGODB_HOST=mongodb://localhost:27017
   MONGODB_DATABASE=auctiondb
   ```

## Executando Localmente

1. Instale as dependências:

   ```sh
   go mod tidy
   ```

2. Inicie o MongoDB:

   ```sh
   docker run --name mongodb -p 27017:27017 -d mongo:latest
   ```

3. Execute a aplicação:

   ```sh
   go run cmd/auction/main.go
   ```

A aplicação estará disponível em `http://localhost:8080`.

## Executando com Docker

1. Construa e inicie os serviços com Docker Compose:

   ```sh
   docker-compose up --build
   ```

A aplicação estará disponível em `http://localhost:8080`.

## Endpoints da API

- `GET /auction`: Lista todos os leilões.
- `GET /auction/:auctionId`: Obtém detalhes de um leilão específico.
- `POST /auction`: Cria um novo leilão.
- `GET /auction/winner/:auctionId`: Obtém o lance vencedor de um leilão.
- `POST /bid`: Cria um novo lance.
- `GET /bid/:auctionId`: Lista todos os lances de um leilão.
- `GET /user/:userId`: Obtém detalhes de um usuário específico.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
