FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.* /

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o /main .

FROM alpine:latest

WORKDIR /app

# Install necessary tools (optional)
RUN apk add --no-cache bash
# Install necessary tools (optional)

COPY --from=builder /main .
COPY wait-for-it.sh .
COPY .env .

RUN chmod +x main wait-for-it.sh

EXPOSE 8080

CMD ["./wait-for-it.sh", "db:5432", "--", "./main"]