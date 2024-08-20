# Build stage
FROM --platform=linux/amd64 golang:1.18-alpine AS builder

WORKDIR /go/src/app

COPY ./service/go.mod ./service/go.sum ./

RUN go mod download

COPY ./service/ .

RUN go build -o app .

# Runtime stage
FROM --platform=linux/amd64 alpine AS app

COPY --from=builder /go/src/app/app /bin/app

COPY --from=builder /go/src/app/banks-response.json /banks-response.json

EXPOSE 8080

CMD ["/bin/app"]
