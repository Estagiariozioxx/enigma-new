
# Enigma-New: Sistema de Gerenciamento de Criptografia

Este projeto simula um sistema inspirado na máquina de criptografia ENIGMA, usado para decifrar mensagens. Com esta aplicação, é possível gerenciar chaves criptográficas, realizar autenticação de usuários e descriptografar documentos com uma chave específica. A aplicação foi desenvolvida em Golang 

## Índice

- [Descrição do Projeto](#descrição-do-projeto)
- [Pré-requisitos](#pré-requisitos)
- [Configuração](#configuração)
- [Inicialização e Execução da Aplicação](#inicialização-e-execução-da-aplicação)
- [Uso dos Endpoints](#uso-dos-endpoints)
- [Considerações sobre Segurança e Regras de Negócio](#considerações-sobre-segurança-e-regras-de-negócio)
- [Documentação Swagger](#documentação-swagger)

## Descrição do Projeto

O sistema implementa:
- Autenticação com JWT.
- Gerenciamento de chaves criptográficas com persistência no banco de dados.
- Descriptografia de documentos com a cifra de César.
- Controle de acesso para o usuário MESTRE, que possui privilégios exclusivos.

## Pré-requisitos

- **Golang** (versão 1.16 ou superior)
- **Banco de Dados**: Pode-se utilizar PostgreSQL, MySQL ou SQLite. Configure conforme o arquivo `.env`.
- **Docker** (Opcional): Para facilitar a configuração e execução em um ambiente isolado.

## Configuração

1. Clone o repositório:
   ```bash
   git clone <URL-do-repositório>
   cd enigma-new
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Configure o arquivo `.env`:
   - Configure as variáveis de ambiente no arquivo `.env`, incluindo as credenciais do banco de dados e outras configurações necessárias.

4. Configure o banco de dados:
   - Certifique-se de que o banco de dados está rodando.
   - Rode as migrações iniciais se necessário.

## Inicialização e Execução da Aplicação

Inicie o servidor com o comando:

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080` (ou conforme configurado).

## Uso dos Endpoints

### Autenticação e Autorização

1. **Login do usuário**:
   - **Rota**: `POST /api/login`
   - **Descrição**: Realiza o login e retorna um token JWT.
   - **Corpo**:
     ```json
     {
       "username": "nome_do_usuario",
       "password": "senha_do_usuario"
     }
     ```

2. **Endpoints principais**:

   - **GET /current-key**: Retorna a chave de criptografia atual.
   - **GET /keys?page={n}**: Retorna uma lista de todas as chaves paginada.
   - **GET /users**: Lista todos os usuários.
   - **POST /users**: Cria um novo usuário.
   - **PUT /users/{id}**: Atualiza o nome de um usuário específico.
   - **DELETE /users/{id}**: Deleta um usuário.

3. **Descriptografia**:
   - **Rota**: `POST /decrypt`
   - **Descrição**: Recebe um arquivo criptografado e a primeira palavra descriptografada.
   - **Corpo**:
     ```json
     {
       "file": "<arquivo>",
       "first_word": "primeira_palavra"
     }
     ```


## Considerações sobre Segurança e Regras de Negócio

1. **Usuário MESTRE**:
   - Possui privilégios exclusivos e não pode ser modificado ou deletado.

2. **Segurança**:
   - As senhas são armazenadas com hashing usando salt.
   - A autenticação utiliza JWT com expiração de 1 dia.

3. **Controle de Acesso**:
   - Endpoints principais requerem autenticação JWT.
   - Aplicação protege contra ataques de força bruta com limitador de tentativas de login.

## Documentação Swagger

Ao iniciar a aplicação a documentação da API fica disponível em http://localhost:8080/swagger/index.html.

## Docker

O uso de Docker facilita o setup. Para rodar a aplicação com Docker, utilize o arquivo `Dockerfile` e siga os passos:

1. Construa a imagem:
   ```bash
   docker build -t enigma-new .
   ```

2. Rode o container:
   ```bash
   docker run -p 8080:8080 --env-file .env enigma-new
   ```
