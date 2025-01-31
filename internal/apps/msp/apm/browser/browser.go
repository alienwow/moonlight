// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package browser

import (
	"fmt"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	writer "github.com/ping-cloudnative/moonlight-utils/pkg/parallel-writer"
	"github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/browser/ipdb"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/lib/kafka"
)

type config struct {
	Ipdb   string               `file:"ipdb"`
	Output kafka.ProducerConfig `file:"output"`
	Input  kafka.ConsumerConfig `file:"input"`
}

type provider struct {
	Cfg    *config
	Log    logs.Logger
	ipdb   *ipdb.Locator
	output writer.Writer
	Kafka  kafka.Interface `autowired:"kafkago"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	ipdb, err := ipdb.NewLocator(p.Cfg.Ipdb)
	if err != nil {
		return fmt.Errorf("fail to load ipdb: %s", err)
	}
	p.ipdb = ipdb
	p.Log.Infof("load ipdb from %s", p.Cfg.Ipdb)

	w, err := p.Kafka.NewProducer(&p.Cfg.Output)
	if err != nil {
		return fmt.Errorf("fail to create kafka producer: %s", err)
	}
	p.output = w
	return nil
}

// Start .
func (p *provider) Start() error {
	return p.Kafka.NewConsumer(&p.Cfg.Input, p.invoke)
}

// Close .
func (p *provider) Close() error {
	p.Log.Debug("not support close kafka consumer")
	return nil
}

func init() {
	servicehub.Register("browser-analytics", &servicehub.Spec{
		Services:    []string{"browser-analytics"},
		Description: "browser-analytics",
		ConfigFunc:  func() interface{} { return &config{} },
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
