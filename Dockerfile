FROM golang:1.14.2 as builder

WORKDIR /go/src/
COPY . .
RUN go test -v ./handles

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main basic-rest

FROM scratch
ENV PORT 80
COPY --from=builder /go/src/main ./basic-rest
CMD ["./basic-rest"]