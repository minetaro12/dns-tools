# dns-tools

```
#ビルド
$ go build -o main

#実行
$ ./main

$ PORT=8080 ./main
```

## 使い方

```
$ curl localhost:8000/whois -X POST --form "domain=0sn.net"

$ curl localhost:8000/dig -X POST --form "fqdn=example.com" --form "type=a" --form "dns=1.1.1.1"
```
