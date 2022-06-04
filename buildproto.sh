# 用于生成pb文件 
# protoc --go_out=plugins=grpc:. hello.proto

# 用于生成pb文件，pb.go文件， gateway文件
protoc --go_out . --go-grpc_out . --grpc-gateway_out . hello.proto
# protoc --go_out=plugins=grpc:. --grpc-gateway_out . hello.proto
