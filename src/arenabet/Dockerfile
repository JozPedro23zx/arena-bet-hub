FROM golang:1.17

WORKDIR /go/src

RUN apt-get update && apt-get install build-essential librdkafka-dev -y

COPY . .

CMD ["tail", "-f", "/dev/null"]