# Client-Server API Challenge

## 📌 Visão Geral

Este projeto implementa uma arquitetura **client-server** em Go para consumo de uma API externa de cotação do dólar (USD/BRL), persistência do dado em banco SQLite e geração de um arquivo texto com o valor obtido.

O objetivo principal é demonstrar:

* Uso de **contexts com timeout**
* Comunicação HTTP client/server
* Consumo de API externa
* Persistência com **GORM + SQLite**
* Organização em pacotes

---

## 🚀 Entrypoint

O **entrypoint da aplicação** está localizado em:

```
cmd/main.go
```

É a partir desse arquivo que:

1. O banco de dados é inicializado
2. O servidor HTTP é iniciado
3. O client é executado para consumir a API local


---

## ▶️ Como Executar

```bash
go run cmd/main.go
```

---

## 🧱 Estrutura do Projeto

```
.
├── cmd/
│   └── main.go          # Entrypoint da aplicação
│
├── client/
│   └── client.go        # Client HTTP que consome o servidor
│
├── server/
│   ├── server.go        # Servidor HTTP e handler /cotacao
│   ├── db.go            # Inicialização do banco SQLite
│   └── model.go         # Model Cotacao (GORM)
│
├── cotacao.db           # Banco SQLite (gerado em runtime)
├── cotacao.txt          # Arquivo texto com a cotação (gerado em runtime)
└── go.mod / go.sum
```

---

## ⚙️ Funcionamento da Aplicação

### 1️⃣ Inicialização (`main.go`)

* Inicializa o banco de dados SQLite
* Sobe o servidor HTTP na porta `8080`
* Aguarda brevemente para garantir que o servidor esteja ativo
* Executa o client

Fluxo:

```
main → InitDB → StartServer → client.Run
```

---

### 2️⃣ Servidor (`server`)

#### Endpoint

```
GET /cotacao
```

#### Responsabilidades

* Criar contexto com timeout de **200ms** para a API externa

* Consumir a API:

  ```
  https://economia.awesomeapi.com.br/json/last/USD-BRL
  ```

* Extrair o campo `bid`

* Criar contexto com timeout de **10ms** para o banco

* Persistir a cotação no SQLite

* Retornar JSON no formato:

```json
{
  "bid": "5.12"
}
```

#### Banco de Dados

* SQLite
* Inicializado automaticamente via `AutoMigrate`
* Tabela `cotacaos`

---

### 3️⃣ Client (`client`)

#### Comportamento

* Cria contexto com timeout de **300ms**
* Faz request para:

```
http://localhost:8080/cotacao
```

* Decodifica a resposta JSON
* Extrai o valor do `bid`
* Gera o arquivo `cotacao.txt`

Conteúdo do arquivo:

```
Dólar: 5.12
```

---

## 📦 Outputs Gerados

Ao final da execução, a aplicação gera:

* 📁 **Banco de dados SQLite**

  ```
  cotacao.db
  ```

* 📄 **Arquivo texto com a cotação**

  ```
  cotacao.txt
  ```

---

## 🧪 Timeouts Implementados

| Camada       | Timeout |
| ------------ | ------- |
| API Externa  | 200ms   |
| Banco SQLite | 10ms    |
| Client HTTP  | 300ms   |

Esses timeouts garantem controle de recursos e resiliência da aplicação.
