FROM golang

WORKDIR /go/src/app

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]