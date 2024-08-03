FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o receipt-processor ./cmd/api

EXPOSE 8080

CMD [ "./receipt-processor" ]
