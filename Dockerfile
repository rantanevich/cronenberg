FROM golang:1.17.8-alpine AS go

WORKDIR /app

ARG GOOS=linux \
    CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o cronenberg ./cmd/cronenberg


FROM scratch
COPY --from=go /app/cronenberg /

ENTRYPOINT ["/cronenberg"]
