FROM golang:1.18 AS build

WORKDIR /usr/src/service

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOGC=off go build -a -installsuffix cgo -ldflags="-w -s" -v -o ./service ./cmd/app

EXPOSE 8080

CMD ./service