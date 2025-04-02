FROM golang:1.24.1-bookworm as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN  go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./little-heu-api -ldflags="-s -w"

FROM gcr.io/distroless/static-debian12 as runtime

LABEL maintainer="Jan Novák <jan.novak@silen.org>"
LABEL version="1.1.1"
LABEL description="Little Heureka API"

EXPOSE 8080

COPY --from=builder /build/little-heu-api /little-heu-api
COPY ./openapi.json /openapi.json
COPY ./data /data
CMD ["/little-heu-api"]
