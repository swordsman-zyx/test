FROM golang:1.18

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=8800

WORKDIR /app

COPY . .
RUN go build .

CMD ["./test.exe"]