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

// Package cmp Core components of multi-cloud management platform
package cmp

import (
	"context"
	"embed"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	componentprotocol "github.com/ping-cloudnative/moonlight-utils/providers/component-protocol"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/protocol"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/cmp/component-protocol/types"
	"github.com/ping-cloudnative/moonlight/internal/apps/cmp/metrics"
	"github.com/ping-cloudnative/moonlight/internal/apps/cmp/resource"
	"github.com/ping-cloudnative/moonlight/internal/apps/cmp/steve"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/pkg/http/httpclient"
	alertpb "github.com/ping-cloudnative/moonlight/proto-go/cmp/alert/pb"
	pb2 "github.com/ping-cloudnative/moonlight/proto-go/cmp/dashboard/pb"
	clusterpb "github.com/ping-cloudnative/moonlight/proto-go/core/clustermanager/cluster/pb"
	monitor "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/alert/pb"
	"github.com/ping-cloudnative/moonlight/proto-go/core/monitor/metric/pb"
	cronpb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cron/pb"
	tokenpb "github.com/ping-cloudnative/moonlight/proto-go/core/token/pb"
)

//go:embed component-protocol/scenarios
var scenarioFS embed.FS

type provider struct {
	Server      pb.MetricServiceServer         `autowired:"erda.core.monitor.metric.MetricService"`
	Credential  tokenpb.TokenServiceServer     `autowired:"erda.core.token.TokenService" optional:"true"`
	Register    transport.Register             `autowired:"service-register" optional:"true"`
	CronService cronpb.CronServiceServer       `autowired:"erda.core.pipeline.cron.CronService" required:"true"`
	ClusterSvc  clusterpb.ClusterServiceServer `autowired:"erda.core.clustermanager.cluster.ClusterService"`
	Router      httpserver.Router              `autowired:"http-router"`

	Metrics         *metrics.Metric
	Monitor         monitor.AlertServiceServer `autowired:"erda.core.monitor.alert.AlertService" optional:"true"`
	Protocol        componentprotocol.Interface
	Resource        *resource.Resource
	CPTran          i18n.I18n       `autowired:"i18n"`
	Tran            i18n.Translator `translator:"common"`
	SteveAggregator *steve.Aggregator
	Org             org.Interface
}

// Run Run the provider
func (p *provider) Run(ctx context.Context) error {
	runtime.GOMAXPROCS(2)
	p.Metrics = metrics.New(p.Server, ctx)
	logrus.Info("cmp provider is running...")
	p.Resource = resource.New(ctx, p.Tran, p, p.ClusterSvc)
	ctxNew := context.WithValue(ctx, "metrics", p.Metrics)
	ctxNew = context.WithValue(ctxNew, "resource", p.Resource)
	return p.initialize(ctxNew)
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.Protocol.SetI18nTran(p.CPTran)
	p.Protocol.WithContextValue(types.GlobalCtxKeyBundle, bundle.New(
		bundle.WithAllAvailableClients(),
		bundle.WithHTTPClient(
			httpclient.New(
				httpclient.WithTimeout(time.Second*30, time.Second*90),
				httpclient.WithEnableAutoRetry(false),
			)),
	))
	p.Protocol.WithContextValue(types.ClusterSvc, p.ClusterSvc)
	protocol.MustRegisterProtocolsFromFS(scenarioFS)
	pb2.RegisterClusterResourceImp(p.Register, p, apis.Options())
	alertpb.RegisterAlertServiceImp(p.Register, p, apis.Options())

	return nil
}

func init() {
	servicehub.Register("cmp", &servicehub.Spec{
		Services:    append([]string{"cmp"}, pb2.ServiceNames()...),
		Description: "Core components of multi-cloud management platform.",
		Creator:     func() servicehub.Provider { return &provider{} },
	})
}
