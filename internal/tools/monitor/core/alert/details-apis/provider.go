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

package details_apis

import (
	"time"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight-utils/providers/clickhouse"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver/interceptors"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/storekit/clickhouse/table/loader"

	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/internal/pkg/bundle-ex/cmdb"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/metric/query/metricq"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/metric/storage/elasticsearch"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	perm "github.com/ping-cloudnative/moonlight/pkg/common/permission"
	"github.com/ping-cloudnative/moonlight/pkg/http/httpclient"
	"github.com/ping-cloudnative/moonlight/proto-go/core/monitor/alertdetail/pb"
)

type config struct {
	QueryMetricFromCk bool `file:"query_metric_from_clickhouse"`
	DebugSQL          bool `file:"debug_sql"`
}

type provider struct {
	C           *config
	L           logs.Logger
	metricq     metricq.Queryer
	EsSearchRaw elasticsearch.Interface `autowired:"metric-storage" optional:"true"`
	//metricq  metricpb.MetricServiceServer `autowired:"erda.core.monitor.metric.MetricService"`
	cmdb *cmdb.Cmdb

	Register           transport.Register `autowired:"service-register"`
	Perm               perm.Interface     `autowired:"permission"`
	alertDetailService *alertDetailService
	Org                org.ClientInterface
	Clickhouse         clickhouse.Interface `autowired:"clickhouse" optional:"true"`
	Loader             loader.Interface     `autowired:"clickhouse.table.loader@metric" optional:"true"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	hc := httpclient.New(httpclient.WithTimeout(time.Second, time.Second*60))
	p.cmdb = cmdb.New(cmdb.WithHTTPClient(hc), cmdb.WithOrgSvc(p.Org))
	p.metricq = ctx.Service("metrics-query").(metricq.Queryer)
	p.alertDetailService = &alertDetailService{
		p: p,
		ClickhouseSource: ClickhouseSource{
			Clickhouse: p.Clickhouse,
			Log:        p.L,
			DebugSQL:   p.C.DebugSQL,
			Loader:     p.Loader,
			Org:        p.Org,
		},
	}

	if p.Register != nil {
		type AlertDetailService = pb.AlertDetailServiceServer
		pb.RegisterAlertDetailServiceImp(p.Register, p.alertDetailService, apis.Options(), p.Perm.Check(
			perm.Method(AlertDetailService.QuerySystemPodMetrics, perm.ScopeOrg, "monitor_org_center", perm.ActionGet, p.OrgIDByCluster("clusterName")),
		))
	}
	routes := ctx.Service("http-server",
		//telemetry.HttpMetric(),
		interceptors.Recover(p.L)).(httpserver.Router)
	return p.intRoutes(routes)
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.monitor.alertdetail" || ctx.Type() == pb.AlertDetailServiceServerType() || ctx.Type() == pb.AlertDetailServiceHandlerType():
		return p.alertDetailService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.monitor.alertdetail", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		Dependencies:         []string{"metrics-query"},
		OptionalDependencies: []string{"service-register"},
		Description:          "",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
