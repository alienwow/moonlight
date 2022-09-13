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

package micro_api

import (
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/common"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/services/micro_api/impl"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	perm "github.com/ping-cloudnative/moonlight/pkg/common/permission"
	"github.com/ping-cloudnative/moonlight/proto-go/core/hepa/api/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg        *config
	Register   transport.Register
	Perm       perm.Interface `autowired:"permission"`
	apiService *apiService
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.apiService = &apiService{
		p: p,
	}
	err := impl.NewGatewayApiServiceImpl()
	if err != nil {
		return err
	}
	if p.Register != nil {
		type apiService = pb.ApiServiceServer
		pb.RegisterApiServiceImp(p.Register, p.apiService, apis.Options(), p.Perm.Check(
			perm.Method(apiService.CreateApi, perm.ScopeProject, "project", perm.ActionGet, perm.FieldValue("ApiRequest.ProjectId")),
			perm.Method(apiService.UpdateApi, perm.ScopeProject, "project", perm.ActionGet, perm.FieldValue("ApiRequest.ProjectId")),
			perm.Method(apiService.GetApis, perm.ScopeProject, "project", perm.ActionGet, perm.FieldValue("ProjectId")),
			perm.Method(apiService.DeleteApi, perm.ScopeOrg, "org", perm.ActionGet, perm.OrgIDValue()),
		), common.AccessLogWrap(common.AccessLog))
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.hepa.api.ApiService" || ctx.Type() == pb.ApiServiceServerType() || ctx.Type() == pb.ApiServiceHandlerType():
		return p.apiService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.api", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Dependencies: []string{
			"hepa",
			"erda.core.hepa.global.GlobalService",
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
