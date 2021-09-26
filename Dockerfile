FROM golang:1.16 AS builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . ./

RUN go build . 

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/docker_builds .

CMD ["/app/docker_builds"]