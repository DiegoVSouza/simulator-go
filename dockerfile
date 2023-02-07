FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bind:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

CMD [ "tail", "-f", "/dev/null" ]