FROM golang:1.17-alpine3.15

ARG MOCKGEN_VERSION=v1.6.0

WORKDIR /go/src

RUN apk update && apk upgrade
RUN apk add --no-cache git bash
RUN apk add gcc g++ make cmake

# librdkafka
RUN apk add librdkafka-dev
RUN apk add librdkafka-static
RUN apk add librdkafka
# RUN git clone https://github.com/edenhill/librdkafka.git && \
#   cd librdkafka && \
#   ./configure && \
#   make && \
#   make install \
#   make clean

RUN go install github.com/golang/mock/mockgen@${MOCKGEN_VERSION}

COPY . .

# RUN chmod +x .docker/entrypoint.sh

CMD ["tail", "-f", "/dev/null"]
# ENTRYPOINT [".docker/entrypoint.sh"]
