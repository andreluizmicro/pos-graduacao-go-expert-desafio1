# Desafio Client Server API

Este é um desafio proposto no curso expert go

### Estruturação do projeto

O diretório `client` contém a implementação do serviço client que faz uma requisição para api (server) criar um nova cotação.

e projeto server está na raiz do desafio, a ideia de manter essa estrutura, foi para manter todo o código do desafio em um único repositório


### Banco de dados

Criar a tabela `cotation`

```SQL
CREATE TABLE cotation (
	id varchar(255), 
	code varchar(80), 
	code_in varchar(80), 
	name varchar(255), 
	high decimal(8,4), 
	low decimal(8,4), 
	var_bid decimal(8,4), 
	pct_change decimal(8,4), 
	bid decimal(8,4), 
	ask decimal(8,4), 
	created_at timestamp, 
	primary key (id)
);
```

### Rodando o server

Para rodar o server basta rodar o comando abaixo ainda na raiz do diretório principal com isso temos um server rodando no http://localhost:8080

    go run cmd/main.go

### Rodando o cliente

Para o rodar o cliente basta acessar o diretório cliente e executar o seguinte comando

    go run cmd/main.go

### Funcionamento do projeto

Ao rodar o comando no cliente, será realizada um chamada ao server, e após a consulta o dados dever ser armazenado no banco dad dados na tabela cotation.

Também deve ser criado um arquivo chamado `cotacao.txt` no diretório tmp do projeto client, e o arquivo deve conter:

```txt
Dólar:{5.0445}
```