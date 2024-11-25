FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o server server2.go

EXPOSE 8082

CMD ["./server"]
