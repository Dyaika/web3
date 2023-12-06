FROM golang:latest

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["go-cookie-service"]
EXPOSE 8080
