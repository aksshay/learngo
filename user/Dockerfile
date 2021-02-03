FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go get -u github.com/gorilla/mux
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o userservice

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/userservice .

ENTRYPOINT ["./userservice"]