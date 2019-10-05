FROM golang as go

WORKDIR /go/test
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' main.go

FROM alpine

WORKDIR /app
COPY --from=go /go/test/ .
EXPOSE 8123
CMD ["/app/main"]
