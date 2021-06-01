FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build
COPY . .

RUN go mod download

#compilar e criar binário
RUN go build -o bin ./cmd/demo/main.go

FROM alpine
COPY --from=0 /build/bin /hello

ENTRYPOINT ["/hello"]
EXPOSE 1323