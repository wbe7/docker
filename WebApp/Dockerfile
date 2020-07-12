FROM golang:1.14
WORKDIR /app
COPY main.go .
RUN go build -o WebApp
CMD ["/app/WebApp"]