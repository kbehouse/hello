# GATEWAY_FLAGS := -I. -I/usr/local/include -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I/usr/local/include
# GATEWAY_FLAGS := -I. -I/usr/local/include -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.1/third_party/googleapis -I/usr/local/include
GATEWAY_PATH = $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway)
GATEWAY_FLAGS := -I. -I/usr/local/include -I$(GATEWAY_PATH)/third_party/googleapis -I/usr/local/include
GRPC_FLAGS := --python_out=. --grpc_python_out=.

code:
	python -m grpc_tools.protoc $(GRPC_FLAGS) $(GATEWAY_FLAGS) *.proto

gw:
	protoc $(GATEWAY_FLAGS) \
        --go_out=plugins=grpc:. \
        --grpc-gateway_out=logtostderr=true:. \
        *.proto

deps:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
