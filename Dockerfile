FROM golang:1.20

WORKDIR /app

ADD . /app

CMD ["go","run","main.go"]