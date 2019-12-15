FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8090

CMD ["go", "run", "main.go"]
