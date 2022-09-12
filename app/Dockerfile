# hadolint ignore=DL3006
FROM golang:1.18 as build

LABEL maintainer="alex@ideaup.cl"

WORKDIR /app

COPY cli/go.mod cli/go.sum ./

RUN go mod download

COPY cli ./

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -v -o glonner-cli main.go

FROM gcr.io/distroless/base
WORKDIR /app

COPY --from=build app/ ./
ENTRYPOINT ["./glonner-cli"]
