FROM golang:1.17-buster

COPY ./src/go.mod /app/
COPY ./src/go.sum /app/
WORKDIR /app

RUN go mod download

COPY ./src /app
RUN go build -o main

CMD [ "./main" ]
