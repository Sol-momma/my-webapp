# ベースイメージとしてGoのAlpine版を利用
FROM golang:1.21-alpine

# 必要なパッケージをインストール
RUN apk add --no-cache git

# 作業ディレクトリの作成
WORKDIR /app

# Goモジュールファイルをコピーして依存関係をダウンロード
# COPY go.mod go.sum ./
# RUN go mod download

# ソースコード全体をコピー
COPY . .

# アプリケーションのビルド
# RUN go build -o main .

# コンテナ起動時のコマンド
CMD ["/app/main"]
