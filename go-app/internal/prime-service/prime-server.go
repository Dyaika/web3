package prime_service

import (
	"context"
	"go-app/maintenance"
	prm "go-app/proto/prime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math"
	"net"
	"strconv"
	"sync"
)

type PrimeServer struct {
	prm.CheckPrimeServer
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	maxDivisor := int(math.Sqrt(float64(n)))
	for i := 2; i <= maxDivisor; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func (s *PrimeServer) CheckPrime(ctx context.Context, req *prm.IsPrimeRequest) (*prm.IsPrimeResponse, error) {
	number := int(req.Message)
	answ := ""
	if isPrime(number) {
		answ = strconv.Itoa(number) + " - простое число"
	} else {
		answ = strconv.Itoa(number) + " - не простое число"
	}
	return &prm.IsPrimeResponse{Answer: answ}, nil
}

func Start(wg *sync.WaitGroup) {
	defer wg.Done()
	var port = maintenance.GetEnv("PRI_PORT", "50051")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	prm.RegisterCheckPrimeServer(s, &PrimeServer{})
	reflection.Register(s)

	maintenance.LogData("Prime server started on port " + port)
	if err := s.Serve(lis); err != nil {
		maintenance.LogData("ERROR: " + err.Error())
	}
}
