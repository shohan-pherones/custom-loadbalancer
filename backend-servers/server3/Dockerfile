FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o server server3.go

EXPOSE 8083

CMD ["./server"]
