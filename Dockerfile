FROM golang:1.16 as builder

WORKDIR /work

COPY . /work/

RUN CGO_ENABLED=0 go build -o server .

FROM scratch

LABEL org.opencontainers.image.source https://github.com/hk220/echo

WORKDIR /work

COPY --from=builder /work/server /work/server

EXPOSE 9001

ENTRYPOINT [ "/work/server" ]
