FROM golang:stretch

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8090
EXPOSE 50051

CMD ["go", "run", "main.go"]
