FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7

ENV REDIS_ADDR="127.0.0.1:6379"
ENV REDIS_PW="root"
ENV REDIS_DB="2"
ENV MysqlDSN="root:root@tcp(127.0.0.1:3306)/to_do_list?charset=utf8&parseTime=true"
ENV GIN_MODE="release"
ENV JWT_SECRET = "Todo-List"
ENV PORT=3000


WORKDIR /app

COPY --from = build /app /usr/bin/api_server
ADD ./conf /app/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]