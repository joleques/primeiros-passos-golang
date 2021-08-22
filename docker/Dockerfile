FROM ubuntu:20.04

RUN apt-get update && \
apt-get install -y \
build-essential \
curl \
gcc \
git \
gnupg \
libc6-dev \
wget \
vim

ENV GO_INSTALL_DIR="/usr/local/"
ENV PATH=$PATH:$GO_INSTALL_DIR/go/bin
ENV GOPATH="/go"
ENV GOBIN="/usr/local/go/bin"

RUN wget https://dl.google.com/go/go1.17.linux-amd64.tar.gz
RUN tar zxvf go1.17.linux-amd64.tar.gz -C $GO_INSTALL_DIR

WORKDIR /usr/local/go

EXPOSE 8080

CMD [ "bash" ]