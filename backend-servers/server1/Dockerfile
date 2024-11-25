FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o server server1.go

EXPOSE 8081

CMD ["./server"]
