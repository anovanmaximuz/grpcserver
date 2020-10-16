package service

import (
	"context"
	"math"
	"log"

	"github.com/anovanmaximuz/grpcserver/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MathService is used for call grpc math service server
type MathService struct {
}

// Ordinal function for math process ordinal with grpc
func (ms *MathService) Ordinal(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Printf("have a request with number %d and ordinal %d", request.Number, request.Ordinal)
	total := math.Pow(float64(request.Number), float64(request.Ordinal))
	return &proto.Response{Result: int64(total)}, nil
}
