package adder

import (
	"calendar/api/proto"
	"context"
)

type GRPCServer struct {
	proto.UnimplementedAdderServer
}

// ADD
func (s *GRPCServer) Add(c context.Context, r *proto.AddRequest) (*proto.AddResponse, error) {

	return &proto.AddResponse{Result: r.GetX() + r.GetY()}, nil
}
