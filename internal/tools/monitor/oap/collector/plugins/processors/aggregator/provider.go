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

package aggregator

import (
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

var providerName = plugins.WithPrefixProcessor("aggregator")

type config struct {
	Keypass    map[string][]string `file:"keypass"`
	Keydrop    map[string][]string `file:"keydrop"`
	Keyinclude []string            `file:"keyinclude"`
	Keyexclude []string            `file:"keyexclude"`

	Rules []RuleConfig `file:"rules"`
}

var _ model.Processor = (*provider)(nil)

// +provider
// TODO. Watch out: only work with metric's Fields now, so specify field key without `fields.` prefix
type provider struct {
	Cfg *config
	Log logs.Logger

	cache  map[uint64]aggregate
	rulers []*ruler
}

func (p *provider) ComponentClose() error {
	return nil
}

type aggregate struct {
	data *metric.Metric
}

func (p *provider) ComponentConfig() interface{} {
	return p.Cfg
}

func (p *provider) ProcessMetric(item *metric.Metric) (*metric.Metric, error) {
	return p.add(item), nil
}

func (p *provider) ProcessLog(item *log.Log) (*log.Log, error)        { return item, nil }
func (p *provider) ProcessSpan(item *trace.Span) (*trace.Span, error) { return item, nil }
func (p *provider) ProcessRaw(item *odata.Raw) (*odata.Raw, error)    { return item, nil }

func (p *provider) add(item *metric.Metric) *metric.Metric {
	id := item.Hash()
	_, ok := p.cache[id]
	if !ok {
		agg := aggregate{
			data: item,
		}
		for _, rule := range p.rulers {
			agg.data = rule.Fn(nil, agg.data)
		}
		p.cache[id] = agg
		return item
	}

	pre := p.cache[id]
	for _, rule := range p.rulers {
		pre.data = rule.Fn(pre.data, item)
	}
	p.cache[id] = pre

	return pre.data
}

// Run this is optional
func (p *provider) Init(ctx servicehub.Context) error {
	p.cache = make(map[uint64]aggregate)
	p.rulers = make([]*ruler, len(p.Cfg.Rules))
	for idx, r := range p.Cfg.Rules {
		rr, err := newRuler(r)
		if err != nil {
			return fmt.Errorf("newRuler err: %w", err)
		}
		p.rulers[idx] = rr
	}
	return nil
}

func init() {
	servicehub.Register(providerName, &servicehub.Spec{
		Services: []string{
			providerName,
		},
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Description: "Only work with Metric.Fields",
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
