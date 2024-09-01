# Fiber API com GoLang, GORM, SQLite e Swagger

## Descrição

Esta API foi desenvolvida em GoLang utilizando o framework Fiber e o ORM GORM, com o banco de dados SQLite. Ela implementa dois endpoints principais para gerenciar usuários:

1. **Cadastro de Usuário**: Permite o registro de um novo usuário com um nome de usuário, login e senha.
2. **Autenticação de Usuário**: Permite a autenticação de um usuário com base no login e senha, retornando um token JWT para acesso seguro.

A API também inclui documentação automatizada utilizando Swagger, o que facilita o entendimento e a utilização dos endpoints.

## Funcionalidades

- **Registrar um usuário**: Cadastra um novo usuário com um nome de usuário, login e senha.
- **Autenticar um usuário**: Autentica um usuário existente com um login e senha e gera um token JWT.
- **Gerar documentação Swagger**: A documentação da API é gerada automaticamente utilizando Swagger.

## Tecnologias Utilizadas

- **GoLang**: Linguagem de programação utilizada para desenvolver a API.
- **Fiber**: Um framework web leve e rápido para Go, inspirado no Express.js.
- **GORM**: Um ORM para Go que facilita a manipulação de banco de dados.
- **SQLite**: Um banco de dados leve utilizado para armazenar os dados dos usuários.
- **Swagger**: Utilizado para gerar documentação automática para a API.
- **JWT (JSON Web Token)**: Utilizado para a autenticação segura dos usuários.

## Endpoints

### 1. Registro de Usuário

- **Rota**: `POST /register`
- **Descrição**: Registra um novo usuário.
- **Corpo da Requisição**:
  ```json
  {
    "usuarioNome": "string",
    "usuarioLogin": "string",
    "senha": "string"
  }

- **Resposta**:
  ```json
  {
    "usuarioNome": "string",
    "usuarioLogin": "string",
    "dataHoraExpiracao": "2024-09-01T13:36:22.941Z",
    "token": "string",
    "autenticado": true
  }

- **Resposta Status**:
  ```json
  
    201 Created: Usuário registrado com sucesso.
    400 Bad Request: Erro na requisição.
    500 Internal Server Error: Erro ao salvar o usuário no banco de dados.
  

### 2. Autenticação de Usuário

- **Rota**: `POST /login`
- **Descrição**: Registra um novo usuário.
- **Corpo da Requisição**:
  ```json
  {
    "login": "string",
    "senha": "string"
  }


# Instalação e Execução
### 1. Clone o repositório:
  ```json
    git clone https://github.com/thiiagofernando/fiber-api.git
    cd fiber-api
 ```
### 2. Instale as dependências:
  ```json
    go mod tidy
  ``` 

### 3. Gere a documentação Swagger:
  ```json
    swag init
  ``` 

### 4. Execute a API:
  ```json
    go run main.go
  ```


### 5. Acesse a documentação Swagger em:
  ```json
    http://localhost:3000/swagger/index.html
  ```  


# Contribuição
#### Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.