FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod download -x
RUN go build -o svca ./cmd/main.go

CMD ["./svca"]
