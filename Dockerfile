FROM golang:1.16-alpine

RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./build/app ./src/cmd

CMD ["./build/app"]