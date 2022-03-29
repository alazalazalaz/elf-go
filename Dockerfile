FROM golang:1.14.10 AS compile
ADD ../.. /build
WORKDIR /build
ENV GO111MODULE auto
RUN CGO_ENABLED=0 go build -o application example/app/main.go

FROM alpine
COPY --from=compile /build/application /application
#chmod +x /build/app
ENTRYPOINT ["/application"]


