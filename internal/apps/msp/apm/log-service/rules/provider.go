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

package rules

import (
	"time"

	"github.com/ping-cloudnative/moonlight/internal/core/org"
	metricpb "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/metric/pb"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver/interceptors"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight-utils/providers/mysql"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/log-service/rules/db"
	"github.com/ping-cloudnative/moonlight/pkg/http/httpclient"
)

type config struct{}

type provider struct {
	C          *config
	L          logs.Logger
	db         *db.DB
	bdl        *bundle.Bundle
	MetricMeta metricpb.MetricMetaServiceServer `autowired:"erda.core.monitor.metric.MetricMetaService"`
	t          i18n.Translator
	Org        org.ClientInterface
}

func (p *provider) Init(ctx servicehub.Context) error {
	hc := httpclient.New(httpclient.WithTimeout(time.Second, time.Second*60))
	p.bdl = bundle.New(
		bundle.WithHTTPClient(hc),
		bundle.WithErdaServer(),
	)
	p.t = ctx.Service("i18n").(i18n.I18n).Translator("log-metrics")
	p.db = db.New(ctx.Service("mysql").(mysql.Interface).DB())
	routes := ctx.Service("http-server", interceptors.Recover(p.L)).(httpserver.Router)
	return p.intRoutes(routes)
}

func init() {
	servicehub.Register("log-metric-rules", &servicehub.Spec{
		Services:     []string{"log-metric-rules"},
		Dependencies: []string{"http-server", "mysql", "i18n"},
		Description:  "logs metric rules",
		ConfigFunc:   func() interface{} { return &config{} },
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
