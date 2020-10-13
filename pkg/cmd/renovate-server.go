/*
Copyright 2020 The arhat.dev Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"

	"arhat.dev/pkg/log"
	"github.com/spf13/cobra"

	"arhat.dev/renovate-server/pkg/conf"
	"arhat.dev/renovate-server/pkg/constant"
)

func NewRenovateServerCmd() *cobra.Command {
	var (
		appCtx       context.Context
		configFile   string
		config       = new(conf.Config)
		cliLogConfig = new(log.Config)
	)

	renovateServerCmd := &cobra.Command{
		Use:           "renovate-server",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Use == "version" {
				return nil
			}

			var err error
			appCtx, err = conf.ReadConfig(cmd, &configFile, cliLogConfig, config)
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(appCtx, config)
		},
	}

	flags := renovateServerCmd.PersistentFlags()

	flags.StringVarP(&configFile, "config", "c",
		constant.DefaultRenovateServerConfigFile, "path to the config file")
	flags.AddFlagSet(conf.FlagsForServer("", &config.Server))

	return renovateServerCmd
}

func run(appCtx context.Context, config *conf.Config) error {
	logger := log.Log.WithName("server")

	_ = logger
	return nil
}
