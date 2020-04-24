FROM chain_base:latest AS build-env

FROM alpine:3.10
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/SaKu2110/chain_dev/miner_api /usr/local/bin/miner_api
ENV LISTEN_ADDR 50051
ENV LISTEN_IP localhost

EXPOSE 50051
CMD [ "/usr/local/bin/miner_api" ]