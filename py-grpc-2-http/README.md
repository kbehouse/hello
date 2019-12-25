## Install

```
pip install grpcio
pip install grpcio-tools
pip install google-api-core
```

## Run

```
cd cmd/
go run main.go
```

```
py server.py
```

Make sure protoc install
```
http://google.github.io/proto-lens/installing-protoc.html
```

## Test

```
curl -X "POST" "http://localhost:8081/v1/hello" -H 'Content-Type: application/json; charset=utf-8' -d '{"name": "fff"}'
```


## Refer

1. https://gist.github.com/doubleyou/41f0828e4b9b50a38f7db815feed0a6c
2. https://www.cnblogs.com/lienhua34/p/6285829.html
3. thanks to [hykuan](https://github.com/hykuan)