FROM golang:1.25.1

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN mkdir ./temp

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build  -o /usr/local/bin/app ./avenuectl

CMD ["app"]