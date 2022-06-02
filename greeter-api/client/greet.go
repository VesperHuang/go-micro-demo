package client

import (
	"context"
	greeter "greeter-api/proto"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/registry"
)

var client greeter.GreeterService

// Init ... In the main function, you should call Init() first,
// so the 'client' defined above can be initialized.
func Init() {

	cmd.DefaultRegistries["etcd"] = func(opts ...registry.Option) registry.Registry {
		opts = append(opts, etcd.Auth("user", "pass"))
		return etcd.NewRegistry(opts...)
	}

	service := micro.NewService(
		micro.Name("micro.client.greeter"),
	)

	service.Init()

	// NewGreeterService is defined at proto/greeter.pb.micro.go file,
	// "micro.service.greeter" should match the name you defined in the server part.
	client = greeter.NewGreeterService("micro.service.greeter", service.Client())
}

// Greet ...
func Greet(ctx *gin.Context) {
	name := ctx.Query("name") // ctx.Query will return the GET request query string
	log.Println("Client handle Greet, name =", name)

	// Client request the RPC server for response
	res, err := client.Greet(context.TODO(), &greeter.Request{Name: name})
	if err != nil {
		log.Print(err.Error())
		// return with 400 error code and error message
		ctx.JSON(400, gin.H{"message": "server error"})
		return
	}

	// return 200 success code and the response from server
	ctx.JSON(200, gin.H{"message": res.Greeting})
}
