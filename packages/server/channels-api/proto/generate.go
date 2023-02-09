package proto

//go:generate mkdir -p ../proto/generated
//go:generate protoc -I ../protobuf-import -I . --go_out=../proto/generated --go_opt=paths=source_relative --go-grpc_out=../proto/generated --go-grpc_opt=paths=source_relative  ./common.proto ./messageevent.proto
