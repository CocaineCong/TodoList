FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /TodoList

WORKDIR /TodoList

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7

ENV REDIS_ADDR="127.0.0.1:6379"
ENV REDIS_PW="root"
ENV REDIS_DB="2"
ENV MysqlDSN="root:root@tcp(127.0.0.1:3306)/to_do_list?charset=utf8&parseTime=true"
ENV GIN_MODE="release"
ENV JWT_SECRET = "Todo-List"
ENV PORT=3000


#RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
#    apk update && \
#    apk add ca-certificates && \
#    echo "hosts: files dns" > /etc/nsswitch.conf && \
#    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /TodoList/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]