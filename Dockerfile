FROM golang:1.2

WORKDIR /app

ADD . /app

CMD ["go","run","main.go"]