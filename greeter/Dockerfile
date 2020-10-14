FROM golang:1.12
WORKDIR /go/src/greeter
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

FROM scratch
COPY --from=0 /go/src/greeter/server .
EXPOSE 80
ENTRYPOINT ["/server"]
