FROM golang:1.18.0-alpine

RUN apk add --update && apk add git

RUN apk add curl vim

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz

RUN mv ./migrate.linux-amd64 /usr/bin/migrate

# appディレクトリの作成
RUN mkdir /app

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . /app
