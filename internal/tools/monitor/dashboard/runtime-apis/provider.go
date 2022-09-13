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

package runtimeapis

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver/interceptors"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/metric/query/metricq"
)

type provider struct {
	L       logs.Logger
	metricq metricq.Queryer
	Org     org.ClientInterface
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.metricq = ctx.Service("metrics-query").(metricq.Queryer)
	routes := ctx.Service("http-server", interceptors.Recover(p.L)).(httpserver.Router)
	return p.intRoutes(routes)
}

func init() {
	servicehub.Register("runtime-apis", &servicehub.Spec{
		Services:     []string{"runtime-apis"},
		Dependencies: []string{"http-server", "metrics-query"},
		Description:  "runtime apis",
		Creator:      func() servicehub.Provider { return &provider{} },
	})
}
