FROM golang:1.24.6

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/httpserver/*.go

EXPOSE 8080
CMD [ "./svr" ]
