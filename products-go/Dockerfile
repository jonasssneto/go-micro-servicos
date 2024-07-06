FROM golang:1.22.1

WORKDIR /go/src/app
COPY . . 

EXPOSE 8000

RUN go build -o main cmd/main.go

ENV GIN_MODE=release

CMD ["./main"]