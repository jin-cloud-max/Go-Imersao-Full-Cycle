FROM golang:1.17

WORKDIR /go/src

RUN apt-get update && apt-get install build-essential librdkafka-dev -y

RUN apt-get sqlite3

CMD ["tail", "-f", "/dev/null"]
