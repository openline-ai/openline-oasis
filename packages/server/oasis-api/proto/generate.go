package proto

//go:generate mkdir -p ../proto/generated
//go:generate protoc --go_out=../proto/generated --go_opt=paths=source_relative --go-grpc_out=../proto/generated --go-grpc_opt=paths=source_relative  ./oasisapi.proto
