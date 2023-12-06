package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	length_service "go-app/internal/length-service"
	prime_service "go-app/internal/prime-service"
	"go-app/maintenance"
	lng "go-app/proto/length"
	prm "go-app/proto/prime"
	"google.golang.org/grpc"
	"net/http"
	"sync"
)

func main() {
	maintenance.LoadEnv()
	var wg sync.WaitGroup

	wg.Add(1)
	go length_service.Start(&wg)

	wg.Add(1)
	go prime_service.Start(&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}

		len_port := maintenance.GetEnv("LEN_PORT", "50054")
		if err := lng.RegisterCalcLengthHandlerFromEndpoint(ctx, mux, "localhost:"+len_port, opts); err != nil {
			maintenance.LogData("ERROR: " + err.Error())
		}

		pri_port := maintenance.GetEnv("PRI_PORT", "50051")
		if err := prm.RegisterCheckPrimeHandlerFromEndpoint(ctx, mux, "localhost:"+pri_port, opts); err != nil {
			maintenance.LogData("ERROR: " + err.Error())
		}

		port := maintenance.GetEnv("GATEWAY_PORT", "8080")
		maintenance.LogData("gRPC Gateway is running on port " + port)
		if err := http.ListenAndServe(":"+port, mux); err != nil {
			maintenance.LogData("ERROR: " + err.Error())
		}
	}()

	wg.Wait()
}
