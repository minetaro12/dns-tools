# dns-tools

```
#ビルド
$ go build

#実行
$ ./dns-tools

$ PORT=8080 ./dns-tools
```

## 使い方

```
$ curl localhost:8000/whois/ -X POST --form "domain=0sn.net"

$ curl localhost:8000/dig/ -X POST --form "domain=example.com" --form "type=a" --form "dns=1.1.1.1"
```
