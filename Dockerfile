FROM golang:1.15-alpine

RUN mkdir /app
ADD ./ /app

WORKDIR /app
RUN go build -o chatcli

ENTRYPOINT ["/app/chatcli", "-type", "server"]