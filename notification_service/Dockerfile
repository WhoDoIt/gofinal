FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN ls -a
RUN go mod download

COPY . .
RUN go build -o ./app cmd/server/*.go
CMD ["./app"]