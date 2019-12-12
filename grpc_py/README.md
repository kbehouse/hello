
# gRPC


## python gRPC
 
refer: https://grpc.io/docs/quickstart/python/

Install
```
python -m pip install grpcio-tools
```

-I: proto directory

--python_out: for protobuf

--grpc_python_out: for gRPC

```
cd pb
python -m grpc_tools.protoc -I./ --python_out=. --grpc_python_out=. helloworld.proto
```

## Test

Run Server
```
py greeter_server.py
```

Test Client
```
py greeter_client.py
```

Test grpcurl
```
echo '{"name": "kartik"}' | grpcurl -plaintext  -proto ./pb/helloworld.proto -d @ localhost:50051 pb.Greeter.SayHello
```


## Docker

Build for grpc-python-alpine

```
docker build -t kbehouse/grpc-python-alpine -f Dockerfile_grpc_alpine .
```

Build Server app
```
docker build -t kbehouse/grpc_greet_py  .
```

Run Server
```
docker run -it -p 50051:50051  kbehouse/grpc_greet_py
```

Test (other terminal)

(Intstall grpcurl from https://github.com/fullstorydev/grpcurl)

```
echo '{"name": "kartik"}' | grpcurl -plaintext  -proto ./pb/helloworld.proto -d @ localhost:50051 pb.Greeter.SayHello
```