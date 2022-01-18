# Reast API em GOLANG
<p>🚀 Projeto de estudo sobre a liguagem Go com uma API Rest e banco de dados MySQL</p>

## Para rodar o projeto 
- Ter o golang instalado <https://go.dev/doc/install>
- Ter o mysql instalado <https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/>
- Criar uma tabela chamada `users` para que possa testar as operações do CRUD
- Possuir um API Client. Ex: Postman, Insomnia



### 🏁  Criando a tabela do Projeto

```mysql
CREATE TABLE Persons (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(50) NOT NULL,
    email varchar(100),
    PRIMARY KEY (id)
);

```


### 🎲  Rodando o Back End (servidor)

```bash
# Clone este repositório
$ git clone <https://github.com/bahcasac/CRUD-GOLANG.git>

# Acesse a pasta raiz do projeto
$ cd CRUD-GOLANG

# Configurando o seu ambiente
$ cp .env.sample .env

# Rode o comando de execução no arquivo `main.go`
$ go run main.go

# O servidor inciará na porta:5001 - acesse <http://localhost:5001>
```

### 🎉  Consumindo a API 

```bash
# Criar um usuário
curl --location --request POST 'http://localhost:5001/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"Maria",
    "email":"maria@gmail.com"
}'

# Listagem de usuários
curl --location --request GET 'http://localhost:5001/users'

#  Listar usuário por id
curl --location --request GET 'http://localhost:5001/users/1'

# Atualizar um usuário
curl --location --request PUT 'http://localhost:5001/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"jose",
    "email":"jose@gmail.com"
}'

# Apagar um usuário
curl --location --request DELETE 'http://localhost:5001/users/1'
```


