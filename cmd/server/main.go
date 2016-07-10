package main

import (
	"net"

	"golang.org/x/net/context"

	"github.com/boltdb/bolt"
	pb "github.com/dan-compton/gofigure/gofigure"
	_ "github.com/lib/pq"
	log "github.com/opsee/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement grpc methods defined in /gofigure
type server struct {
	store *bolt.DB
}

// Registers new service and configuration under version bucket 0
func (s *server) NewConfig(ctx context.Context, in *pb.NewConfigRequest) (*pb.NewConfigResponse, error) {
	// TODO
	return nil, nil
}

// Returns configuration for service under latest version bucket
func (s *server) GetConfig(ctx context.Context, in *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	// TODO
	return nil, nil
}

// Increments version bucket number for service and adds new configuration under that bucket
func (s *server) UpdateConfig(ctx context.Context, in *pb.UpdateConfigRequest) (*pb.UpdateConfigResponse, error) {
	// TODO
	return nil, nil
}

// Returns a status code indicating if a service's configuration is up to date, outdated, from the future, etc
func (s *server) CheckConfig(ctx context.Context, in *pb.UpdateConfigRequest) (*pb.UpdateConfigResponse, error) {
	// TODO
	return nil, nil
}

func main() {
	viper.SetEnvPrefix("gofigure")
	viper.AutomaticEnv()

	viper.SetDefault("log_level", "debug")
	logLevelStr := viper.GetString("log_level")
	logLevel, err := log.ParseLevel(logLevelStr)
	if err != nil {
		log.WithError(err).Error("Could not parse log level, using default.")
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	db, err := bolt.Open("gofigure.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGoFiguratorServer(s, &server{store: db})
	s.Serve(lis)
}
