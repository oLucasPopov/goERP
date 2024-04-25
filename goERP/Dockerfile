FROM golang:1.21

# Diretório de trabalho na VM linux
WORKDIR /app/public/goERP

# Copiar a pasta atual na máquina física para o diretório de trabalho do linux
COPY . .

# Baixa as dependências da aplicação
RUN go mod download

# Compila a aplicação go
RUN go build -o /app/public/goERP

# Dá permissão de execução e escrita (de logs) para a aplicação
RUN chmod +x+w /app/public/goERP

# Roda a aplicação
CMD ["/app/public/goERP/goERP"]

# Expõe a porta rodada no docker na porta 8080
EXPOSE 8080