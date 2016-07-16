package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/context"

	"github.com/boltdb/bolt"
	proto "github.com/golang/protobuf/proto"

	pb "github.com/dan-compton/gofigure/gofigure"
	_ "github.com/lib/pq"
	log "github.com/opsee/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	port = ":9114"
)

// server is used to implement grpc methods defined in /gofigure
type server struct {
	store *bolt.DB
}

// Registers configuration for service and sets configuration timestamp
func (s *server) NewConfig(ctx context.Context, in *pb.NewConfigRequest) (*pb.NewConfigResponse, error) {
	log.Debugf("received newconfig request for version: %s", in.Configuration.Version.Id)
	response := &pb.NewConfigResponse{
		Status: pb.NewConfigResponse_ERRUNSPECIFIED,
	}

	// validation
	if in.Configuration == nil || in.Configuration.Version == nil {
		return response, fmt.Errorf("missing configuration or configuration version")
	}

	err := s.store.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(in.ServiceName))
		if err != nil {
			return err
		}

		// ignore client config timestamp
		t := time.Now().UTC().Unix()
		in.Configuration.Version.Timestamp = t

		// store config as proto bytes
		d, err := proto.Marshal(in.Configuration)
		if err != nil {
			return err
		}

		// store the configuration
		err = b.Put([]byte(in.Configuration.Version.Id), d)
		if err != nil {
			return err
		}

		// set current service configuration version
		err = b.Put([]byte("current_version"), []byte(in.Configuration.Version.Id))
		if err != nil {
			return err
		}

		response.Status = pb.NewConfigResponse_SUCCESS
		return nil
	})
	if err != nil {
		log.WithError(err).Error("couldn't create new service config")
	}

	return response, err
}

// Returns configuration for service under latest version bucket
func (s *server) GetConfig(ctx context.Context, in *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	log.Debugf("received getconfig request")
	response := &pb.GetConfigResponse{
		Status: pb.GetConfigResponse_ERRUNSPECIFIED,
	}

	// validation
	if in.Version == nil {
		return response, fmt.Errorf("configuration version")
	}

	err := s.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(in.ServiceName))

		var version string
		if in.Version != nil {
			// get non current version of config
			version = in.Version.Id
		} else {
			// get current config version
			v := b.Get([]byte("current_version"))
			if v == nil {
				response.Status = pb.GetConfigResponse_ERRINVALID
				return nil
			}
			log.Debugf("Got current config version number: %d", v)
			version = string(v)
		}

		// get config
		d := b.Get([]byte(version))
		if d == nil {
			response.Status = pb.GetConfigResponse_ERRINVALID
			return nil
		}

		// unmarshal config
		config := &pb.Config{}
		err := proto.Unmarshal(d, config)
		if err != nil {
			return err
		}

		response.Status = pb.GetConfigResponse_SUCCESS
		response.Configuration = config
		return nil
	})
	if err != nil {
		log.WithError(err).Error("couldn't create new service config")
	}

	return response, err
}

// Increments version bucket number for service and adds new configuration under that bucket
// if Config.Current == true, updates config_version to reflect Config.Version
func (s *server) UpdateConfig(ctx context.Context, in *pb.UpdateConfigRequest) (*pb.UpdateConfigResponse, error) {
	log.Debugf("received updateconfig request")
	response := &pb.UpdateConfigResponse{
		Status: pb.UpdateConfigResponse_ERRUNSPECIFIED,
	}

	// validation
	if in.Configuration == nil || in.Configuration.Version == nil {
		return response, fmt.Errorf("missing configuration or configuration version")
	}

	err := s.store.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(in.ServiceName))

		// store config as proto bytes
		d, err := proto.Marshal(in.Configuration)
		if err != nil {
			return err
		}

		// ignore client config timestamp
		t := time.Now().UTC().Unix()
		in.Configuration.Version.Timestamp = t

		err = b.Put([]byte(in.Configuration.Version.Id), d)
		if err != nil {
			return err
		}

		// store version number
		err = b.Put([]byte("current_version"), []byte(in.Configuration.Version.Id))
		if err != nil {
			return err
		}
		log.Debugf("updated current_version of %s to %s", in.ServiceName, in.Configuration.Version.Id)

		response.Status = pb.UpdateConfigResponse_SUCCESS
		return nil
	})
	if err != nil {
		log.WithError(err).Error("couldn't create new service config")
	}

	return response, err
}

// Returns a status code indicating if a service's configuration is up to date, outdated, from the future, etc
func (s *server) ConfigCheck(ctx context.Context, in *pb.ConfigCheckRequest) (*pb.ConfigCheckResponse, error) {
	log.Debugf("received checkconfig request")
	response := &pb.ConfigCheckResponse{
		Status: pb.ConfigCheckResponse_ERRUNSPECIFIED,
	}

	err := s.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(in.ServiceName))

		// get current config version
		// TODO store the entire version proto here
		v := b.Get([]byte("current_version"))
		if v == nil {
			response.Status = pb.ConfigCheckResponse_ERRDOESNOTEXIST
			return nil
		}
		log.Debugf("got config current_version: %s", v)

		// get config
		d := b.Get([]byte(v))
		if v == nil {
			response.Status = pb.ConfigCheckResponse_ERRDOESNOTEXIST
			return nil
		}

		// unmarshal config
		config := &pb.Config{}
		err := proto.Unmarshal(d, config)
		if err != nil {
			return err
		}

		response.Status = pb.ConfigCheckResponse_SUCCESS
		response.Version = config.Version
		return nil
	})
	if err != nil {
		log.WithError(err).Error("couldn't create new service config")
	}

	return response, err
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
	log.Info("connected to boltdb")

	log.Info("staring gofigure server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGoFiguratorServer(s, &server{store: db})
	s.Serve(lis)
}
