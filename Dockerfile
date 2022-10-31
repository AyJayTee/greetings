FROM golang:1.19-alpine AS builder

WORKDIR /go/src/app
COPY . .
RUN go build cmd/http/main.go


FROM alpine

WORKDIR /usr/home
COPY --from=builder /go/src/app/main /usr/home
ENTRYPOINT [ "./main" ]