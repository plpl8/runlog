FROM golang:1.27

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go build -o runlog-server .

EXPOSE 8080

CMD ["./runlog-server"]
