FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o assignment-2 .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/assignment-2 .

CMD ["./assignment-2"]