FROM golang:1.14.10 AS compile
ADD . /build
WORKDIR /build
ENV GO111MODULE auto
#注意build出来的可执行文件名不能和同层的文件夹名重合，比如go build -o app example/app/main.go就会出现报错。
RUN CGO_ENABLED=0 go build -o application example/app/main.go

FROM alpine
COPY --from=compile /build/application /application
#chmod +x /build/app
ENTRYPOINT ["/application"]


