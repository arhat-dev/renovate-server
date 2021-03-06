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

package conf

import (
	"time"

	"arhat.dev/pkg/kubehelper"
	"arhat.dev/pkg/log"
	"arhat.dev/pkg/tlshelper"
	"github.com/spf13/pflag"

	"arhat.dev/renovate-server/pkg/constant"
)

type Config struct {
	Server ServerConfig `json:"server" yaml:"server"`

	GitHub []PlatformConfig `json:"github" yaml:"github"`
	GitLab []PlatformConfig `json:"gitlab" yaml:"gitlab"`
}

type ServerConfig struct {
	Log log.ConfigSet `json:"log" yaml:"log"`

	Webhook struct {
		Listen string              `json:"listen" yaml:"listen"`
		TLS    tlshelper.TLSConfig `json:"tls" yaml:"tls"`
	} `json:"webhook" yaml:"webhook"`

	Scheduling struct {
		// Delay period for webhook event
		Delay time.Duration `json:"delay" yaml:"delay"`

		// CronTabs is the crontab string array to schedule renovate periodically
		CronTabs []string `json:"cronTabs" yaml:"cronTabs"`
		// Timezone
		Timezone string `json:"timezone" yaml:"timezone"`
	} `json:"scheduling" yaml:"scheduling"`

	Executor struct {
		Kubernetes *KubernetesExecutorConfig `json:"kubernetes" yaml:"kubernetes"`
	} `json:"executor" yaml:"executor"`
}

type KubernetesExecutorConfig struct {
	KubeClient kubehelper.KubeClientConfig `json:"kubeClient" yaml:"kubeClient"`

	// JobTTL delete after specified time period
	JobTTL time.Duration `json:"jobTTL" yaml:"jobTTL"`

	RenovateImage           string `json:"renovateImage" yaml:"renovateImage"`
	RenovateImagePullPolicy string `json:"renovateImagePullPolicy" yaml:"renovateImagePullPolicy"`
}

func FlagsForServer(prefix string, config *ServerConfig) *pflag.FlagSet {
	fs := pflag.NewFlagSet("app", pflag.ExitOnError)

	fs.StringVar(&config.Webhook.Listen, prefix+"webhook.listen",
		constant.DefaultWebhookListenAddress, "set webhook listener address",
	)
	fs.AddFlagSet(tlshelper.FlagsForTLSConfig(prefix+"webhook.tls", &config.Webhook.TLS))
	fs.DurationVar(&config.Scheduling.Delay, prefix+"scheduling.delay",
		constant.DefaultSchedulingDelay, "set delay time before actually invoke executor",
	)

	return fs
}
