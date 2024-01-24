FROM golang:1.17

WORKDIR /app

COPY . .

RUN go mod init main.go

RUN go build -o concurrent .

ENTRYPOINT ["./concurrent", "-file", "testcur.csv", "-concurrent", "false"]


