FROM node:lts-alpine

WORKDIR /server

#ファイルのコピー
COPY ./index.js /server/
COPY ./package.json /server/
COPY ./yarn.lock /server/
COPY ./public/ /server/public/

#依存関係のインストール
RUN apk add --no-cache bind-tools && \
    yarn install && \
    yarn cache clean && \
    rm -rf /root/.npm

CMD /usr/local/bin/yarn start
