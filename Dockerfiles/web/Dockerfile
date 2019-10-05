FROM node:lts-alpine

# 静的コンテンツを配信するシンプルな http サーバをインストールする
#RUN npm install -g http-server

# カレントワーキングディレクトリとして 'app' フォルダを指定する
WORKDIR /app

# `package.json` と `package-lock.json` （あれば）を両方コピーする
COPY ./web/package*.json ./

# プロジェクトの依存ライブラリをインストールする
RUN npm install

# カレントワーキングディレクトリ(つまり 'app' フォルダ)にプロジェクトのファイルやフォルダをコピーする
COPY ./web/ .

# 本番向けに圧縮しながらアプリケーションをビルドする
#RUN npm run build
#RUN nohup npm run dev &

EXPOSE 3000
#CMD [ "http-server", "dist" ]
CMD npm run dev
