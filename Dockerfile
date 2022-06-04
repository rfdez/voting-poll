# syntax=docker/dockerfile:1

FROM golang:1.18-alpine as build

WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

COPY go.mod ./
COPY go.sum ./

RUN go mod download \
    && go mod verify

COPY cmd ./cmd
COPY internal ./internal
COPY kit ./kit

RUN go build -o /poll -a ./cmd/api/main.go


FROM gcr.io/distroless/base-debian11

WORKDIR /

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT [ "/poll" ]

COPY --from=build /poll /poll
