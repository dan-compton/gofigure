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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	address       string
	service       string
	version       string
	serviceConfig string
)

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "client",
	Short: "Creates or updates service configurations from Gofigure store",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.PersistentFlags().StringVarP(&address, "address", "a", "localhost:9114", "address of target gofigure server")
	RootCmd.PersistentFlags().StringVarP(&service, "service", "s", "default_service", "identifier for given service")
	RootCmd.PersistentFlags().StringVarP(&version, "version", "v", "default_version", "identifier for given configuration")
	RootCmd.PersistentFlags().StringVarP(&serviceConfig, "config_file", "c", "config.yml", "yaml file representing a service configuration")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("gofigure")
	viper.AutomaticEnv() // read in environment variables that match
}
