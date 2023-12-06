package length_service

import (
	"context"
	"go-app/maintenance"
	lng "go-app/proto/length"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
	"sync"
)

type LengthServer struct {
	lng.CalcLengthServer
}

func (s *LengthServer) CalcLength(ctx context.Context, req *lng.LengthRequest) (*lng.LengthResponse, error) {
	number := strconv.Itoa(len(req.Message))
	return &lng.LengthResponse{Answer: "Длина: " + number}, nil
}

func Start(wg *sync.WaitGroup) {
	defer wg.Done()
	var port = maintenance.GetEnv("LEN_PORT", "50054")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	lng.RegisterCalcLengthServer(s, &LengthServer{})
	reflection.Register(s)

	maintenance.LogData("Length server started on port " + port)
	if err := s.Serve(lis); err != nil {
		maintenance.LogData("ERROR: " + err.Error())
	}
}
