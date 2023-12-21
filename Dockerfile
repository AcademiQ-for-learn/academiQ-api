FROM golang:1.22rc1-alpine3.19

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
