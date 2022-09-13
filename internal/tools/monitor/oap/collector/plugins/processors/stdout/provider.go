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

package stdout

import (
	"encoding/json"
	"fmt"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/trace"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/log"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/metric"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/core/model"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/core/model/odata"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/plugins"
)

var providerName = plugins.WithPrefixProcessor("stdout")

type config struct {
	Keypass map[string][]string `file:"keypass"`
}

var _ model.Processor = (*provider)(nil)

// +provider
type provider struct {
	Cfg *config
	Log logs.Logger
}

func (p *provider) ComponentClose() error {
	return nil
}

func (p *provider) ComponentConfig() interface{} {
	return p.Cfg
}

func (p *provider) ProcessMetric(item *metric.Metric) (*metric.Metric, error) {
	printJSON(item)
	return item, nil
}
func (p *provider) ProcessLog(item *log.Log) (*log.Log, error) {
	printJSON(item)
	return item, nil
}
func (p *provider) ProcessSpan(item *trace.Span) (*trace.Span, error) {
	printJSON(item)
	return item, nil
}
func (p *provider) ProcessRaw(item *odata.Raw) (*odata.Raw, error) {
	printJSON(item)
	return item, nil
}

func printJSON(data interface{}) {
	buf, _ := json.Marshal(data)
	fmt.Printf("%s\n", string(buf))
}

// Run this is optional
func (p *provider) Init(ctx servicehub.Context) error {
	return nil
}

func init() {
	servicehub.Register(providerName, &servicehub.Spec{
		Services: []string{
			providerName,
		},
		Description: "here is description of erda.oap.collector.processor.stdout",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
