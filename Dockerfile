FROM golang:1.19

WORKDIR /usr/src/app

COPY . .

CMD ["go", "run", "main.go"]