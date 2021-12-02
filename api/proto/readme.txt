use this command
.\protoc.exe --go_out=. --go_opt=paths=source_relative   --go-grpc_out=. --go-grpc_opt=paths=source_relative    added.proto


add this  -> 	proto.UnimplementedAdderServer to the struct 

type GRPCServer struct {
	proto.UnimplementedAdderServer
}
