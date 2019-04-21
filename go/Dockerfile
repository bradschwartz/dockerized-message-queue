FROM golang:1.8

WORKDIR /app

COPY . .

ENV GOBIN /go/bin
RUN go get

CMD ["go", "run", "main.go"]