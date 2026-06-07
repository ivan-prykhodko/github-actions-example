ARG GO_VERSION=1.25


FROM golang:${GO_VERSION}-alpine AS builder
ENV CGO_ENABLED=0
ADD . /build
WORKDIR /build
RUN cd cmd/app && go build -ldflags "-s -w" -o /build/app


FROM scratch AS app_prod
COPY --from=builder /build/app /srv/app
WORKDIR /srv
CMD ["/srv/app"]


FROM golang:${GO_VERSION}-alpine AS app_dev
WORKDIR /srv
CMD ["go", "run", "cmd/app/main.go"]
