package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code
	pb "github.com/cabotmoose/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/cabotmoose/shippy/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
	// "golang.org/x/net/context"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Setup micro instance
	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}