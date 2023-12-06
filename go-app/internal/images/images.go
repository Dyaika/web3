package images

import (
	"bytes"
	"context"
	imgg "github.com/myuser/myrepo/proto/images"
	"github.com/nfnt/resize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net"
	"os"
	"sync"
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

func Start(wg *sync.WaitGroup) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":"+os.Getenv("IMAGE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	imgg.RegisterImageProcessingServer(s, &ImageProcessingServer{})
	reflection.Register(s)

	log.Println("ImageProcessing gRPC Server is running on port " + os.Getenv("IMAGE_PORT"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
