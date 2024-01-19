FROM golang:1.21.6-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN  go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /little-heu-api -ldflags="-s -w"

FROM gcr.io/distroless/static-debian11
COPY --from=builder /little-heu-api /little-heu-api
COPY ./openapi.json /openapi.json
COPY /data /data
CMD ["/little-heu-api"]
