# 安装  protoc
--  protoc --version

# 安装 protoc-gen-go 
-- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# go get google.golang.org/protobuf

# protoc --go_out=. --go_opt=paths=source_relative protocbuf/std.proto  当前目录

# protoc --go_out=.  --go-grpc_out=.  rpc/rpc.proto 当前目录