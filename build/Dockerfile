FROM golang:1.18.0-alpine

RUN apk add --update && apk add git

# appディレクトリの作成
RUN mkdir /app

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . /app
