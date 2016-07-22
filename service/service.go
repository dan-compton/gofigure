package service

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/boltdb/bolt"
	pb "github.com/dan-compton/gofigure/gofigure"
	proto "github.com/golang/protobuf/proto"
	log "github.com/opsee/logrus"
	"golang.org/x/net/context"
)

// server is used to implement grpc methods defined in /gofigure
type service struct {
	store *bolt.DB
}

// Stores a new service configuration at service_name->version_id->timestamp
func (s *service) NewConfig(ctx context.Context, in *pb.NewConfigRequest) (*pb.NewConfigResponse, error) {
	log.Debugf("received newconfig request for version: %s", in.Configuration.Version.Id)
	response := &pb.NewConfigResponse{
		Status: pb.Status_ERR_UNSPECIFIED,
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
		timestamp := fmt.Sprintf("%d", t)
		configId := in.Configuration.Version.Id

		// create version bucket
		bb, err := b.CreateBucketIfNotExists([]byte(configId))
		if err != nil {
			return err
		}
		log.Debugf("created bucket %s in %s", configId, in.ServiceName)

		// store config as proto bytes
		d, err := proto.Marshal(in.Configuration)
		if err != nil {
			return err
		}

		// store the configuration at service_name->Id->timestamp=protobuf
		err = bb.Put([]byte(timestamp), d)
		if err != nil {
			return err
		}

		// store the current version (timestamp) of this configuration version
		err = bb.Put([]byte("current_timestamp"), []byte(timestamp))
		if err != nil {
			return err
		}
		log.Debugf("new current_timestamp %s for %s", timestamp, configId)

		// set current version of this service to configuration version
		err = b.Put([]byte("current_version"), []byte(configId))
		if err != nil {
			return err
		}
		log.Debugf("new current_version %s for %s", configId, in.ServiceName)

		response.Status = pb.Status_SUCCESS
		return nil
	})
	if err != nil {
		log.WithError(err).Error("couldn't create new service config")
	}

	return response, err
}

// Returns configuration for service under latest version bucket
func (s *service) GetConfig(ctx context.Context, in *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	log.Debugf("received getconfig request")
	response := &pb.GetConfigResponse{
		Status: pb.Status_ERR_UNSPECIFIED,
	}

	// validation
	if in.Version == nil {
		return response, fmt.Errorf("configuration version")
	}

	err := s.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(in.ServiceName))
		if b == nil {
			response.Status = pb.Status_ERR_INVALID_SERVICE
			return nil
		}

		var version string
		if in.Version != nil {
			// get non current version of config
			version = in.Version.Id
		} else {
			// get current config version
			v := b.Get([]byte("current_version"))
			if v == nil {
				response.Status = pb.Status_ERR_INVALID_CONFIG
				return nil
			}
			version = string(v)
		}
		log.Debugf("fetching config version: %s", version)

		// get config current version bucket
		bb := b.Bucket([]byte(version))
		if bb == nil {
			log.Debugf("couldn't fetch bucket %s", in.Version.Id)
			response.Status = pb.Status_ERR_INVALID_CONFIG
			return nil
		}

		timestamp := fmt.Sprintf("%d", in.Version.Timestamp)
		log.Debugf("initial timestamp: %s", timestamp)
		if timestamp == "0" {
			// get current config version
			t := bb.Get([]byte("current_timestamp"))
			if t == nil {
				response.Status = pb.Status_ERR_INVALID_CONFIG
				return nil
			}
			timestamp = string(t)
		}
		log.Debugf("fetching config %s timestamp: %s", version, timestamp)

		// get config
		d := bb.Get([]byte(timestamp))
		if d == nil {
			response.Status = pb.Status_ERR_INVALID_CONFIG
			return nil
		}

		// unmarshal config
		config := &pb.Config{}
		err := proto.Unmarshal(d, config)
		if err != nil {
			return err
		}

		response.Status = pb.Status_SUCCESS
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
func (s *service) UpdateConfig(ctx context.Context, in *pb.UpdateConfigRequest) (*pb.UpdateConfigResponse, error) {
	log.Debugf("received updateconfig request")
	response := &pb.UpdateConfigResponse{
		Status: pb.Status_ERR_UNSPECIFIED,
	}

	// validation
	if in.Configuration == nil || in.Configuration.Version == nil {
		return response, fmt.Errorf("missing configuration or configuration version")
	}

	err := s.store.Update(func(tx *bolt.Tx) error {
		// get service bucket
		b := tx.Bucket([]byte(in.ServiceName))
		if b == nil {
			response.Status = pb.Status_ERR_INVALID_SERVICE
			return nil
		}

		var version string
		if in.Configuration.Version != nil {
			// get non current version of config
			version = in.Configuration.Version.Id
		} else {
			// get current config version
			v := b.Get([]byte("current_version"))
			if v == nil {
				response.Status = pb.Status_ERR_INVALID_CONFIG
				return nil
			}
			log.Debugf("Got current config version number: %d", v)
			version = string(v)
		}

		// get specified or current version bucket
		bb := b.Bucket([]byte(version))
		if bb == nil {
			response.Status = pb.Status_ERR_INVALID_CONFIG
			return nil
		}

		// ignore client config timestamp, we don't overwrite old versions ever
		t := time.Now().UTC().Unix()
		in.Configuration.Version.Timestamp = t

		// marshal config to bytes
		d, err := proto.Marshal(in.Configuration)
		if err != nil {
			return err
		}

		// store the protobuf in the current version bucket with timestamp as key
		err = bb.Put([]byte(fmt.Sprintf("%d", t)), d)
		if err != nil {
			return err
		}

		// store current timestamp as current
		err = bb.Put([]byte("current_timestamp"), []byte(fmt.Sprintf("%d", t)))
		if err != nil {
			return err
		}

		// store version number
		err = b.Put([]byte("current_version"), []byte(in.Configuration.Version.Id))
		if err != nil {
			return err
		}
		log.Debugf("updated current_version of %s to Id %s and timestamp %d", in.ServiceName, in.Configuration.Version.Id, t)

		response.Status = pb.Status_SUCCESS
		return nil
	})
	if err != nil {
		log.WithError(err).Error("couldn't create new service config")
	}

	return response, err
}

// Returns a status code indicating if a service's configuration is up to date, outdated, from the future, etc
func (s *service) CheckConfig(ctx context.Context, in *pb.CheckConfigRequest) (*pb.CheckConfigResponse, error) {
	log.Debugf("received checkconfig request")
	response := &pb.CheckConfigResponse{
		Status: pb.Status_ERR_UNSPECIFIED,
	}

	err := s.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(in.ServiceName))

		// get current config version
		v := b.Get([]byte("current_version"))
		if v == nil {
			response.Status = pb.Status_ERR_INVALID_SERVICE
			return nil
		}

		// get config current version bucket
		bb := b.Bucket([]byte(v))
		if bb == nil {
			response.Status = pb.Status_ERR_INVALID_CONFIG
			return nil
		}

		t := bb.Get([]byte("current_timestamp"))
		if t == nil {
			response.Status = pb.Status_ERR_INVALID_CONFIG
			return nil
		}
		log.Debugf("got config current_timestamp: %s", t)

		// get config
		d := bb.Get([]byte(t))
		if v == nil {
			response.Status = pb.Status_ERR_INVALID_CONFIG
			return nil
		}

		// unmarshal config
		config := &pb.Config{}
		err := proto.Unmarshal(d, config)
		if err != nil {
			return err
		}

		response.Status = pb.Status_SUCCESS
		response.Version = config.Version
		return nil
	})
	if err != nil {
		log.WithError(err).Error("couldn't create new service config")
	}

	return response, err
}

func New(db *bolt.DB) *service {
	return &service{
		store: db,
	}
}

func (s *service) Start(addr string) {
	log.Info("staring gofigure server")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs := grpc.NewServer()
	pb.RegisterGoFiguratorServer(gs, s)
	gs.Serve(lis)
}
