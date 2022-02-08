# dns-tools

`yarn install` 依存関係のインストール

`yarn start` ポート8000で起動

`PORT=9000 yarn start` ポート9000で起動

## 使い方

```
curl localhost:8000/whois?domain=example.com

curl localhost:8000/dig?domain=example.com&ytype=ns
```