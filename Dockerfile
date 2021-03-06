FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build
COPY . .

# RUN go mod download -x

#compilar e criar binário
RUN go build -mod vendor -o bin ./cmd/prod/main.go

FROM alpine
WORKDIR /app
COPY --from=0 /build/bin bin

ENTRYPOINT ["./bin"]
EXPOSE 1323