FROM golang:latest AS builder
LABEL maintainer="Konstantin Kriviakin <krivakin@gmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o service1 cmd/service1/*.go


FROM busybox
COPY --from=builder /app/service1 service1
COPY --from=builder /app/cmd/service1/config config

EXPOSE 8080

CMD ["./service1", "--config_dir", "config"]