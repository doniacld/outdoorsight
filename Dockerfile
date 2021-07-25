FROM golang:1.14

RUN mkdir /app
WORKDIR /app
COPY ./bin/cmd/main /app/main
CMD ["/app/main"]
