FROM golang:latest AS builder

WORKDIR /app
COPY . .

# Gera binário Linux compatível com Ubuntu (sem linkagem estática)
RUN go build -o loadtester main.go

# Imagem final baseada no Ubuntu
FROM ubuntu:latest

# Instala dependências mínimas para rodar binários Go
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/loadtester .

ENTRYPOINT ["./loadtester"]
