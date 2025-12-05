FROM golang:1.24:web

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/web

EXPOSE 8085

CMD ["./server"]
