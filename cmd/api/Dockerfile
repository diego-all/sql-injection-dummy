FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod . 
COPY go.sum . 
RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN go build -o main .

EXPOSE 9090

CMD ["./main"]
