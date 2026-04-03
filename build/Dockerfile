FROM docker.io/library/golang:latest

WORKDIR /app

COPY . .

RUN go build -o basic_api .

EXPOSE 8080

CMD ["./basic_api"]