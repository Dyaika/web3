package quad

import (
	"context"
	qa "github.com/myuser/myrepo/proto/quadratic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math"
	"net"
	"os"
	"sync"
)

type QuadraticEquationSolverServer struct {
	qa.QuadraticEquationSolverServer
}

// SolveQuadraticEquation решает квадратное уравнение
func (s *QuadraticEquationSolverServer) SolveQuadraticEquation(ctx context.Context, req *qa.QuadraticEquationRequest) (*qa.QuadraticEquationResponse, error) {
	a := req.GetA()
	b := req.GetB()
	c := req.GetC()

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		// Уравнение не имеет действительных корней
		return &qa.QuadraticEquationResponse{}, nil
	}

	sqrtDiscriminant := math.Sqrt(discriminant)

	root1 := (-b + sqrtDiscriminant) / (2 * a)
	root2 := (-b - sqrtDiscriminant) / (2 * a)

	return &qa.QuadraticEquationResponse{
		Roots: []float64{root1, root2},
	}, nil
}

func Start(wg *sync.WaitGroup) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":"+os.Getenv("QUAD_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	qa.RegisterQuadraticEquationSolverServer(s, &QuadraticEquationSolverServer{})
	reflection.Register(s)

	log.Println("QuadraticEquationSolver gRPC Server is running on port " + os.Getenv("QUAD_PORT"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
