FROM golang:1.9
COPY ./ /app/
WORKDIR /app
RUN go get github.com/DataDog/datadog-go/statsd
RUN go build -o /app/canary-tester /app/cmd/canary-tester.go
CMD ["/app/canary-tester"]
