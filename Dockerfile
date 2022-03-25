FROM golang:1.14.10 AS compile
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 go build -o app example/app/main.go

FROM alpine
COPY --from=compile /build/app /app

RUN chmod +x /app
ENTRYPOINT ["/app"]


