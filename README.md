# API-GO

API-GO é um projeto em Golang que se conecta a um banco de dados PostgreSQL e consome mensagens de uma fila RabbitMQ para processar e armazenar dados de usuários. O projeto demonstra a integração entre Golang, RabbitMQ e PostgreSQL, utilizando goroutines para processamento paralelo das mensagens.

## Funcionalidades

- **Conexão ao PostgreSQL**: Conecta-se a um banco de dados PostgreSQL e realiza operações de CRUD (Create, Read, Update, Delete) em uma tabela de usuários.
- **Consumo de Mensagens RabbitMQ**: Consome mensagens de uma fila RabbitMQ e processa-as em paralelo utilizando goroutines.
- **Inserção de Usuários**: As mensagens recebidas contendo dados de usuários são deserializadas e inseridas na tabela de usuários no PostgreSQL.
- **Criação Dinâmica de Tabelas**: Usa Makefile para gerenciar a criação e migração de tabelas no banco de dados de forma dinâmica.

## Estrutura do Projeto

- `config/`: Contém a configuração do banco de dados e RabbitMQ.
- `domain/`: Contém a definição da estrutura de usuário e funções de interação com o banco de dados.
- `mock/`: Gera usuários mock para testes.
- `message/`: Contém a lógica para consumir e processar mensagens do RabbitMQ.
- `main.go`: Arquivo principal que inicializa a conexão com o banco de dados, insere usuários mock e inicia o listener de mensagens.
- `Makefile`: Contém comandos para gerenciar migrações de banco de dados.
- `user-generator.py`: Script Python para gerar usuários fictícios e enviar mensagens para a fila RabbitMQ para testes.

## Bibliotecas Usadas

- [pq](https://github.com/lib/pq): Driver PostgreSQL para Golang.
- [amqp](https://github.com/streadway/amqp): Biblioteca para interagir com RabbitMQ.
- [uuid](https://github.com/google/uuid): Biblioteca para geração de UUIDs.
- [pika](https://pika.readthedocs.io/): Biblioteca Python para interagir com RabbitMQ.
- [pgx](https://github.com/jackc/pgx): Driver PostgreSQL para Golang com suporte a pool de conexões.

## Pré-requisitos

- Docker
- Docker Compose
- Go 1.22 ou superior
- Python 3.8 ou superior
- PostgreSQL
- RabbitMQ

## Configuração

### Gerando Usuários Fictícios e Enviando para RabbitMQ
O script user-generator.py é usado para gerar usuários fictícios e enviar mensagens para a fila RabbitMQ para testar a aplicação. Para vizualização das mensagens na fila pode estar acessando localmente o endereço: http://localhost:15672/#/queues

### Instale as dependências Python:

```bash
  pip install faker pika
```

### Execute o script user-generator.py:
```bash
  python3 user-generator.py
```

### Variáveis de Ambiente

Defina as seguintes variáveis de ambiente para configuração do RabbitMQ e PostgreSQL:

```bash
export POSTGRES_USER=your_postgres_user
export POSTGRES_PASSWORD=your_postgres_password
export POSTGRES_HOST=your_postgres_host
export POSTGRES_DB=your_postgres_db
export POSTGRES_PORT=5432

export RABBITMQ_USER=guest
export RABBITMQ_PASSWORD=guest
export RABBITMQ_URL=localhost
export RABBITMQ_PORT=5672
export RABBITMQ_VHOST=/

```
### Example de Makefile

```bash
.PHONY: migrate-up migrate-down

migrate-up:
	@echo "Running migrations up..."
	@go run cmd/migrate/main.go up

migrate-down:
	@echo "Running migrations down..."
	@go run cmd/migrate/main.go down

```

### Clone o repositório:
```bash
  git clone https://github.com/seu-usuario/API-GO.git
  cd API-GO
```

### Inicie os contêineres do Docker:
```bash
   docker-compose up -d ||  docker-compose up 
```

### Execute as migrações do banco de dados usando o Makefile:
```bash
  make migrate-up
```

### Execute o projeto:
```bash
  go run main.go

```
## Visualizando Dados no Banco de Dados

Para visualizar os dados inseridos na tabela do banco de dados, você pode usar uma ferramenta de gerenciamento de banco de dados como o DBeaver.