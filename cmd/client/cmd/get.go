// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/dan-compton/gofigure/gofigure"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

var timestamp int64

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Could not connect: %s", err.Error())
			return
		}
		defer conn.Close()
		c := pb.NewGoFiguratorClient(conn)

		ncr := &pb.GetConfigRequest{
			ServiceName: serviceName,
			Version: &pb.ConfigVersion{
				Id:        configVersion,
				Timestamp: timestamp,
			},
		}

		r, err := c.GetConfig(context.Background(), ncr)
		if err != nil {
			fmt.Printf("Could not get config: %s", err.Error())
			return
		} else {
			if r.Status == pb.GetConfigResponse_SUCCESS {
				if configFile != "" {
					switch t := r.Configuration.AConfig.(type) {
					case *pb.Config_Yaml:
						z := t.Yaml
						fmt.Printf("%s, %d", z.RawData, len(z.RawData))
						//ioutil.WriteFile(configFile, , 0644)
					case *pb.Config_Proto:
						z := t.Proto
						fmt.Printf("%v, %d", z.Data)
					default:
						fmt.Printf("not that %s", t)
					}
				}
			} else {
				fmt.Printf("Could not get config.  Code: %v", r.Status)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&configFile, "output_file", "o", "config.yml", "yaml file to write the output config")
	getCmd.Flags().StringVarP(&configVersion, "config_version", "v", "default_version", "identifier for given configuration")
	getCmd.Flags().StringVarP(&serviceName, "service_name", "s", "default_service", "identifier for given service")
	getCmd.Flags().StringVarP(&address, "address", "a", "localhost:9113", "address of target gofigure server")
	getCmd.Flags().Int64VarP(&timestamp, "timestamp", "t", 0, "Unix timestamp of configuration commit")
}
