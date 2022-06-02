package main

import (
	"context" // replace greeter-service if you didn't name greeter-service when go mod init
	greeter "greeter-service/proto"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/registry"
)

// GreeterService ...
type GreeterService struct{}

// Greet ... Implement interface left in proto/greeter.pb.micro.go server part
func (g *GreeterService) Greet(ctx context.Context, req *greeter.Request, res *greeter.Response) error {
	log.Println("Greeter service handle Greet", req.Name)
	res.Greeting = "Hello, " + req.Name
	return nil
}

func init() {
	cmd.DefaultRegistries["etcd"] = func(opts ...registry.Option) registry.Registry {
		opts = append(opts, etcd.Auth("user", "pass"))
		return etcd.NewRegistry(opts...)
	}
}

func main() {
	service := micro.NewService(
		micro.Name("micro.service.greeter"), // The service name to register in the registry
	)

	service.Init()

	// The 'RegisterGreeterHandler' is implement in the proto/greeter.pb.micro.go file
	greeter.RegisterGreeterHandler(service.Server(), &GreeterService{})

	if err := service.Run(); err != nil {
		log.Print(err.Error())
		return
	}
}
