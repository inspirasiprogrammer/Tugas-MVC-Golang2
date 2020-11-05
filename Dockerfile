FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["go","run","main.go"]