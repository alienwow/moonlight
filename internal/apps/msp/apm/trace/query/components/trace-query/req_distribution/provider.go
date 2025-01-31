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

package req_distribution

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/components/bubblegraph"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/components/bubblegraph/impl"
	structure "github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/components/commodel/data-structure"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cpregister"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/trace/query"
	"github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/trace/query/commom/custom"
	metricpb "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/metric/pb"
)

type provider struct {
	impl.DefaultBubbleGraph
	custom.TraceInParams
	Log          logs.Logger
	I18n         i18n.Translator              `autowired:"i18n"`
	TraceService *query.TraceService          `autowired:"erda.msp.apm.trace.TraceService"`
	Metric       metricpb.MetricServiceServer `autowired:"erda.core.monitor.metric.MetricService"`
}

// RegisterInitializeOp .
func (p *provider) RegisterInitializeOp() (opFunc cptype.OperationFunc) {
	return func(sdk *cptype.SDK) cptype.IStdStructuredPtr {
		params := p.TraceInParams.InParamsPtr

		if params.TenantId == "" {
			return nil
		}

		// lang := sdk.Lang
		response, err := p.TraceService.GetTraceReqDistribution(sdk.Ctx, *p.TraceInParams.InParamsPtr)

		if err != nil {
			p.Log.Error(err)
		}
		dataBuilder := bubblegraph.NewDataBuilder().WithTitle(p.I18n.Text(sdk.Lang, "traceDistribution")).
			WithYOptions(bubblegraph.NewOptionsBuilder().WithType(structure.Time).WithPrecision(structure.Nanosecond).WithEnable(true).Build())
		if response == nil {
			p.StdDataPtr = dataBuilder.Build()
			return nil
		}

		for _, row := range response {
			x := row.Date
			y := row.AvgDuration
			size := row.Count

			dataBuilder.WithBubble(bubblegraph.NewBubbleBuilder().
				WithValueX(x).
				WithValueY(y).
				WithValueSize(float64(size)).
				WithDimension("Req Distribution").
				Build())
		}
		p.StdDataPtr = dataBuilder.Build()
		return nil
	}
}

// RegisterRenderingOp .
func (p *provider) RegisterRenderingOp() (opFunc cptype.OperationFunc) {
	return p.RegisterInitializeOp()
}

// Provide .
func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	return p
}

func init() {
	cpregister.RegisterProviderComponent("trace-query", "reqDistribution", &provider{})
}
