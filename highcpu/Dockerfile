
FROM golang:1.16-alpine

WORKDIR /app

COPY main.go .

RUN go mod init main.go


RUN go build -o highcpu


ENTRYPOINT ["./highcpu"]