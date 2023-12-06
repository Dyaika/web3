package main

import (
	"bytes"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/myuser/myrepo/internal/images"
	"github.com/myuser/myrepo/internal/quad"
	"github.com/nfnt/resize"
	"google.golang.org/grpc"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"sync"

	imgg "github.com/myuser/myrepo/proto/images"
	qa "github.com/myuser/myrepo/proto/quadratic"
)

type ImageProcessingServer struct {
	imgg.ImageProcessingServer
}

func (s *ImageProcessingServer) ResizeImage(ctx context.Context, req *imgg.ResizeImageRequest) (*imgg.ResizeImageResponse, error) {
	img, _, err := image.Decode(bytes.NewReader(req.GetImageData()))
	if err != nil {
		return nil, err
	}

	resizedImg := resize.Resize(uint(req.GetWidth()), uint(req.GetHeight()), img, resize.Lanczos3)

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, resizedImg, nil); err != nil {
		return nil, err
	}

	return &imgg.ResizeImageResponse{ResizedImageData: buf.Bytes()}, nil
}

func (s *ImageProcessingServer) ConvertToGrayscale(ctx context.Context, req *imgg.ConvertToGrayscaleRequest) (*imgg.ConvertToGrayscaleResponse, error) {
	img, _, err := image.Decode(bytes.NewReader(req.GetImageData()))
	if err != nil {
		return nil, err
	}

	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, grayImg, nil); err != nil {
		return nil, err
	}

	return &imgg.ConvertToGrayscaleResponse{GrayscaleImageData: buf.Bytes()}, nil
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go quad.Start(&wg)

	wg.Add(1)
	go images.Start(&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}

		// Register QuadraticEquationSolver gRPC service
		if err := qa.RegisterQuadraticEquationSolverHandlerFromEndpoint(ctx, mux, "localhost:"+os.Getenv("QUAD_PORT"), opts); err != nil {
			log.Fatalf("failed to register QuadraticEquationSolver gRPC Gateway endpoint: %v", err)
		}

		// Register ImageProcessing gRPC service
		if err := imgg.RegisterImageProcessingHandlerFromEndpoint(ctx, mux, "localhost:"+os.Getenv("IMAGE_PORT"), opts); err != nil {
			log.Fatalf("failed to register ImageProcessing gRPC Gateway endpoint: %v", err)
		}

		// Start HTTP server for gRPC Gateway
		log.Println("gRPC Gateway is running on port " + os.Getenv("GATEWAY_PORT"))
		if err := http.ListenAndServe(":"+os.Getenv("GATEWAY_PORT"), mux); err != nil {
			log.Fatalf("failed to serve gRPC Gateway: %v", err)
		}
	}()

	wg.Wait()
}
