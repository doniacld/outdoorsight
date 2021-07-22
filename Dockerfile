FROM golang:1.14

RUN mkdir /app
WORKDIR /app
COPY . /app/
CMD ["/app/bin/cmd/main"]
