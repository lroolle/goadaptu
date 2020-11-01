FROM golang:alpine
MAINTAINER Eric Wang <wrqatw@gmail.com>

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
