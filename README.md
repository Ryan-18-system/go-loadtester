# Stress Tester CLI (Go)

Ferramenta de linha de comando em Go para executar testes de estresse em endpoints HTTP, com suporte a concorrência configurável. Pode ser executada nativamente ou via Docker.

---

## 🔧 Funcionalidades

* Envia múltiplas requisições HTTP concorrentes para uma URL alvo.
* Permite configurar:

  * URL alvo (`--url`)
  * Número total de requisições (`--requests`)
  * Nível de concorrência (`--concurrency`)
* Exibe relatório final com:

  * Tempo total do teste
  * Total de requisições bem-sucedidas (200 OK)
  * Distribuição de códigos HTTP retornados

---

## 📂 Estrutura do Projeto

```
stress-tester/
├── main.go
├── Dockerfile
└── README.md
```

---

## 🚀 Como Executar o Stress Tester

### 1. Clonar o repositório

```bash
git clone https://github.com/seu-usuario/stress-tester.git
cd stress-tester
```

### 2. Rodar localmente com Go

```bash
go run main.go --url http://localhost:8080 --requests 1000 --concurrency 50
```

### 3. Rodar via Docker

#### a) Build da imagem

```bash
docker build -t stress-tester .
```

#### b) Executar o container

```bash
docker run --rm stress-tester --url http://host.docker.internal:8080 --requests 1000 --concurrency 50
```

> Obs: use `host.docker.internal` no Windows/Mac para acessar serviços rodando no host. No Linux, use o IP da rede bridge (ex: `172.17.0.1`).

---

## ✅ Exemplo de Saída Esperada

```
========== Relatório ==========
Tempo total: 2.345s
Total de requisições: 1000
Sucesso (200 OK): 978
Distribuição de status HTTP:
  200: 978
  500: 22
================================
```

---

## 🧪 Testes Automatizados

Atualmente não há testes automatizados incluídos. Para adicionar testes, crie arquivos no diretório `tests/` e utilize `go test ./...`.

---
