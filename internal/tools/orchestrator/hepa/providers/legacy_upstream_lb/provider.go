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

package legacy_upstream_lb

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/common"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/services/legacy_upstream_lb/impl"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/proto-go/core/hepa/legacy_upstream_lb/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg               *config
	Log               logs.Logger
	Register          transport.Register
	upstreamLbService *upstreamLbService
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.upstreamLbService = &upstreamLbService{p}
	err := impl.NewGatewayUpstreamLbServiceImpl()
	if err != nil {
		return err
	}
	if p.Register != nil {
		pb.RegisterUpstreamLbServiceImp(p.Register, p.upstreamLbService, apis.Options(), common.AccessLogWrap(common.AccessLog))
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.hepa.legacy_upstream_lb.UpstreamLbService" || ctx.Type() == pb.UpstreamLbServiceServerType() || ctx.Type() == pb.UpstreamLbServiceHandlerType():
		return p.upstreamLbService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.legacy_upstream_lb", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Description:          "",
		Dependencies: []string{
			"hepa",
		},
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
