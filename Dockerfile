FROM golang:1.22-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o beego-mailexam .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/beego-mailexam .

ENV HTTP_HOST=0.0.0.0
ENV HTTP_PORT=8080

EXPOSE 8080

CMD ["./beego-mailexam"]
