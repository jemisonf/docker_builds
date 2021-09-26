FROM golang:1.16

WORKDIR /app

COPY . ./

RUN go build .

CMD ./server