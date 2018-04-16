FROM golang:1.9
COPY ./ /app/
WORKDIR /app
RUN go build -o /app/canary-tester /app/cmd/canary-tester.go
CMD ["/app/canary-tester"]
