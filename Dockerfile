FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o loadbalancer main.go

EXPOSE 8080

CMD ["./loadbalancer"]
