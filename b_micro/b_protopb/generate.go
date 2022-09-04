package b_protopb

//go:generate protoc --go_opt=paths=source_relative --go_out=plugins=grpc:. ./b_proto.proto
