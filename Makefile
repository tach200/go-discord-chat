protogen:
	protoc -I/usr/local/include -I. \
      --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=:. --go-grpc_opt=paths=source_relative \
      --grpc-gateway_out . \
      --grpc-gateway_opt logtostderr=true \
      --grpc-gateway_opt paths=source_relative \
      --grpc-gateway_opt generate_unbound_methods=true \
    proto/message.proto