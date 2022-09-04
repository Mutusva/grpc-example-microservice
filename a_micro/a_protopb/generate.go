package a_protopb

//go:generate protoc --go_opt=paths=source_relative --go_out=plugins=grpc:. ./a_proto.proto
