FROM golang:1.12.1-alpine3.9 AS stage-build

RUN apk add --no-cache ca-certificates cmake make g++ openssl-dev git curl pkgconfig

RUN git clone -b 1.7 https://github.com/neo4j-drivers/seabolt.git /seabolt

WORKDIR /seabolt/build

RUN cmake -D CMAKE_BUILD_TYPE=Release \
    -D CMAKE_INSTALL_LIBDIR=lib .. \
    && cmake --build . --target install

WORKDIR /go/src/github.com/piotrpersona/gcomv

RUN apk update && apk add git dep

COPY Gopkg.* ./

RUN dep ensure --vendor-only

COPY main.go .
COPY cmd cmd
COPY neo neo
COPY app app

RUN OOS=linux GOARCH=amd64 go build \
    -tags seabolt_static \
    -ldflags="-w -s" \
    -o /go/bin/gcomv \
    main.go

ENTRYPOINT [ "/go/bin/gcomv" ]

FROM alpine:3.9

COPY --from=stage-build \
    /go/bin/gcomv /usr/local/bin/gcomv

ENTRYPOINT [ "/usr/local/bin/gcomv" ]
