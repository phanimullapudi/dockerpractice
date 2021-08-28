FROM golang:latest as builder
WORKDIR /go-testrun-files/
ADD https://raw.githubusercontent.com/phanimullapudi/golearning/master/samples/arrays.go ./arrays.go
RUN go env -w GO111MODULE=auto
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM scratch 
COPY --from=builder /go-testrun-files/app .
CMD ["./app"]