FROM golang:1.27

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o runlog-server .

EXPOSE 8080

CMD ["./runlog-server"]
