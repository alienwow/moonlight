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

package org_client

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/common"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/services/org_client/impl"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	perm "github.com/ping-cloudnative/moonlight/pkg/common/permission"
	"github.com/ping-cloudnative/moonlight/proto-go/core/hepa/org_client/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg              *config
	Log              logs.Logger
	Register         transport.Register
	orgClientService *orgClientService
	Perm             perm.Interface `autowired:"permission"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.orgClientService = &orgClientService{p}
	err := impl.NewGatewayOrgClientServiceImpl()
	if err != nil {
		return err
	}
	if p.Register != nil {
		type clientService = pb.OrgClientServiceServer
		pb.RegisterOrgClientServiceImp(p.Register, p.orgClientService, apis.Options(), common.AccessLogWrap(common.AccessLog))
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.hepa.org_client.OrgClientService" || ctx.Type() == pb.OrgClientServiceServerType() || ctx.Type() == pb.OrgClientServiceHandlerType():
		return p.orgClientService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.org_client", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Dependencies: []string{
			"hepa",
			"erda.core.hepa.openapi_rule.OpenapiRuleService",
			"erda.core.hepa.openapi_consumer.OpenapiConsumerService",
		},
		Description: "",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
