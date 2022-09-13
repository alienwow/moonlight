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

package tenant

import (
	"github.com/jinzhu/gorm"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/internal/apps/msp/instance/db/monitor"
	"github.com/ping-cloudnative/moonlight/internal/apps/msp/tenant/db"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	perm "github.com/ping-cloudnative/moonlight/pkg/common/permission"
	"github.com/ping-cloudnative/moonlight/proto-go/msp/tenant/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg           *config
	Log           logs.Logger
	Register      transport.Register
	tenantService *tenantService
	DB            *gorm.DB       `autowired:"mysql-client"`
	Perm          perm.Interface `autowired:"permission"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.tenantService = &tenantService{
		p:            p,
		MSPTenantDB:  &db.MSPTenantDB{DB: p.DB},
		MSPProjectDB: &db.MSPProjectDB{DB: p.DB},
		MonitorDB:    &monitor.MonitorDB{DB: p.DB},
	}
	if p.Register != nil {
		type TenantService pb.TenantServiceServer
		pb.RegisterTenantServiceImp(p.Register, p.tenantService, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.msp.tenant.TenantService" || ctx.Type() == pb.TenantServiceServerType() || ctx.Type() == pb.TenantServiceHandlerType():
		return p.tenantService
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.tenant", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
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
