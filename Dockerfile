#
# Build go project
#
FROM golang:1.15-alpine as go-builder

WORKDIR /go/src/github.com/tanalam2411/mutatingwebhook

COPY pkg pkg
COPY go.mod go.mod
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o mutatingwebhook main.go

#
# Runtime container
#
FROM alpine:latest  

RUN mkdir -p /app && \
    addgroup -S app && adduser -S app -G app && \
    chown app:app /app

WORKDIR /app

COPY --from=go-builder /go/src/github.com/tanalam2411/mutatingwebhook .

USER app

CMD ["./mutatingwebhook"]  
