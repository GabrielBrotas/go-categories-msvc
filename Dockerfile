FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o api ./cmd/api/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/api .

ENV ENVIRONMENT=prod

EXPOSE 8080

CMD [ "/app/api" ]