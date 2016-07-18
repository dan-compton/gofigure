package main

import (
	"time"

	pb "github.com/dan-compton/gofigure/gofigure"
	"github.com/golang/protobuf/ptypes/struct"
	log "github.com/opsee/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9114"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGoFiguratorClient(conn)

	serviceName := "blah"
	serviceVersion := "someversion"

	// Get
	gcr := &pb.GetConfigRequest{
		ServiceName: serviceName,
		Version: &pb.ConfigVersion{
			Id: serviceVersion,
		},
	}

	gr, err := c.GetConfig(context.Background(), gcr)
	log.WithFields(log.Fields{"response": gr, "error": err}).Info("get config response")

	// Update
	ucr := &pb.UpdateConfigRequest{
		ServiceName: serviceName,
		Configuration: &pb.ProtoConfig{
			Version: &pb.ConfigVersion{
				Id:        serviceVersion,
				Timestamp: time.Now().Unix(),
			},
			Configuration: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"somekey": &structpb.Value{&structpb.Value_StringValue{"some"}},
				},
			},
		},
	}

	ur, err := c.UpdateConfig(context.Background(), ucr)
	log.WithFields(log.Fields{"response": ur, "error": err}).Info("update config response")

	// Get same key
	gr, err = c.GetConfig(context.Background(), gcr)
	log.WithFields(log.Fields{"response": gr, "error": err}).Info("get updated config response")

	// Check Config
	ccr := &pb.CheckConfigRequest{
		ServiceName: serviceName,
		Version: &pb.ConfigVersion{
			Id:        serviceVersion,
			Timestamp: 1468699,
		},
	}

	cr, err := c.CheckConfig(context.Background(), ccr)
	log.WithFields(log.Fields{"response": cr, "error": err}).Info("check config response")
}
