FROM golang:latest
LABEL maintainer="Konstantin Kriviakin <krivakin@gmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o service1 cmd/service1/*.go

EXPOSE 8080

CMD ["./service1"]