FROM golang:1.24.6

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

ARG bin_to_build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/${bin_to_build}/*.go

CMD [ "./svr" ]
