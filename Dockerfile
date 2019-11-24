FROM golang:1.12.1-alpine3.9 AS stage-build

RUN apk add --no-cache ca-certificates cmake make g++ openssl-dev git curl pkgconfig

RUN git clone -b 1.7 https://github.com/neo4j-drivers/seabolt.git /seabolt

WORKDIR /seabolt/build

RUN cmake -D CMAKE_BUILD_TYPE=Release \
    -D CMAKE_INSTALL_LIBDIR=lib .. \
    && cmake --build . --target install

WORKDIR /go/src/github.com/piotrpersona/gg

RUN apk update && apk add git dep

COPY Gopkg.* ./

RUN dep ensure --vendor-only

COPY . .

RUN OOS=linux GOARCH=amd64 go build \
    -tags seabolt_static \
    -ldflags="-w -s" \
    -o /go/bin/gg \
    main.go

ENTRYPOINT [ "/go/bin/gcomv" ]

FROM alpine:3.9

RUN apk update && apk add ca-certificates

COPY --from=stage-build \
    /go/bin/gg /usr/local/bin/gg

# Microbadger: https://microbadger.com/labels
ARG VCS_REF
LABEL org.label-schema.vcs-ref=$VCS_REF \
        org.label-schema.vcs-url="https://github.com/piotrpersona/gg"

ENTRYPOINT [ "/usr/local/bin/gg" ]
