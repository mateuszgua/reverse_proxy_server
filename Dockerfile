FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /reverse_proxy_server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /reverse_proxy_server /reverse_proxy_server

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT [ "/reverse_proxy_server" ]