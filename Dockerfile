FROM golang:1.14 AS builder

WORKDIR /go/src/app

COPY . .

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o main .

FROM scratch

WORKDIR /dist

COPY --from=builder /go/src/app/main .

EXPOSE 8080

CMD ["./main"]