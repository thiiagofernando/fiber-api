# Etapa 1: Build
FROM golang:1.20 AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar go mod e sum arquivos
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código fonte para o container
COPY . .

# Construir o binário da aplicação
RUN go build -o main .

# Etapa 2: Runtime
FROM alpine:latest

# Instalar as dependências do runtime
RUN apk --no-cache add ca-certificates

# Definir o diretório de trabalho
WORKDIR /root/

# Copiar o binário da aplicação da etapa de build
COPY --from=builder /app/main .

# Definir o comando para iniciar a aplicação
CMD ["./main"]

# Expor a porta que a aplicação utiliza
EXPOSE 3000
