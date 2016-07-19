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

	"google.golang.org/grpc"

	pb "github.com/dan-compton/gofigure/gofigure"
	"github.com/dan-compton/gofigure/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	configFile    string
	configVersion string
	serviceName   string
	address       string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Stores a new configuration in the configuration store",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Could not connect: %s", err.Error())
			return
		}
		defer conn.Close()
		c := pb.NewGoFiguratorClient(conn)

		yamlBytes, err := util.ReadYAML(configFile)
		if err != nil {
			fmt.Printf("Could not create new configuration: %s", err.Error)
			return
		}

		fmt.Printf("%s", string(yamlBytes))

		ncr := &pb.NewConfigRequest{
			ServiceName: serviceName,
			Configuration: &pb.Config{
				Version: &pb.ConfigVersion{
					Id: configVersion,
				},
				AConfig: &pb.Config_Yaml{
					Yaml: &pb.YamlConfig{
						RawData: yamlBytes,
					},
				},
			},
		}

		r, err := c.NewConfig(context.Background(), ncr)
		if err != nil {
			fmt.Printf("Could not create new config: %s", err.Error())
			return
		} else {
			if r.Status == pb.NewConfigResponse_SUCCESS {
				fmt.Printf("Created new config.")
			} else {
				fmt.Printf("Could not create new config.  Code: %v", r.Status)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&configFile, "config_file", "c", "config.yml", "yaml file representing service configuration")
	newCmd.Flags().StringVarP(&configVersion, "config_version", "v", "default_version", "identifier for given configuration")
	newCmd.Flags().StringVarP(&serviceName, "service_name", "s", "default_service", "identifier for given service")
	newCmd.Flags().StringVarP(&address, "address", "a", "localhost:9113", "address of target gofigure server")
}
