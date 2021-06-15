FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build
COPY . .

RUN go mod download

#compilar e criar binário
RUN go build -o bin ./cmd/dev/main.go

FROM alpine
COPY --from=0 /build/bin /podchef

ENTRYPOINT ["/podchef"]
EXPOSE 1323