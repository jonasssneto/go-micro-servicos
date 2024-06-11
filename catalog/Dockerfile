FROM golang:1.22

WORKDIR /go/src/
copy . .
RUN GOOS=linux go build -ldflags="-s -w" 
CMD ["./catalog"]