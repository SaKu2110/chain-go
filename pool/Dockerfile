FROM chain_base:latest AS build-env

FROM alpine:3.10
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/SaKu2110/chain_dev/pool_api /usr/local/bin/pool_api
ENV LISTEN_ADDR 50051

EXPOSE 50051
CMD [ "/usr/local/bin/pool_api" ]