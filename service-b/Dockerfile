FROM golang:1.23

WORKDIR /app
COPY . .

RUN go mod download -x
RUN go build -o svcb ./cmd/main.go

CMD ["./svcb"]
