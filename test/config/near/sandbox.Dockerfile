FROM debian:bookworm as builder
WORKDIR /usr/src/app
RUN apt-get update && apt-get install --assume-yes curl
RUN curl -LJO https://s3-us-west-1.amazonaws.com/build.nearprotocol.com/nearcore/Linux-x86_64/1.38.0/aac5e42fe8975e27faca53e31f53f9c67a5b4e35/near-sandbox.tar.gz
RUN tar -xf near-sandbox.tar.gz

FROM debian:bookworm-slim as runtime
WORKDIR /usr/local/bin
COPY --from=builder /usr/src/app/Linux-x86_64/near-sandbox /usr/local/bin/near-sandbox
RUN apt-get update && apt-get install --assume-yes curl jq
RUN near-sandbox --home /root/.near init

COPY * /root/.near

RUN cat /root/.near/validator_key.json

ENTRYPOINT [ "near-sandbox", "--home", "/root/.near", "run" ]
