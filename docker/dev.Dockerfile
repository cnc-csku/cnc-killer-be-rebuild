FROM golang:1.24-alpine
RUN apk add make
RUN mkdir app

ADD . /app/

WORKDIR /app

RUN go install github.com/air-verse/air@latest

CMD ["air"]