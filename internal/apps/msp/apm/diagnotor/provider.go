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

package diagnotor

import (
	"github.com/jinzhu/gorm"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/bundle"
	monitordb "github.com/ping-cloudnative/moonlight/internal/apps/msp/instance/db/monitor"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	perm "github.com/ping-cloudnative/moonlight/pkg/common/permission"
	basepb "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/diagnotor/pb"
	"github.com/ping-cloudnative/moonlight/proto-go/msp/apm/diagnotor/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg                  *config
	Log                  logs.Logger
	Register             transport.Register            `autowired:"service-register" optional:"true"`
	BaseDiagnotorService basepb.DiagnotorServiceServer `autowired:"erda.core.monitor.diagnotor.DiagnotorService"`
	Perm                 perm.Interface                `autowired:"permission"`
	DB                   *gorm.DB                      `autowired:"mysql-client"`

	monitor          *monitordb.MonitorDB
	bdl              *bundle.Bundle
	diagnotorService *diagnotorService
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.bdl = bundle.New(bundle.WithScheduler(), bundle.WithErdaServer())
	p.monitor = &monitordb.MonitorDB{DB: p.DB}

	p.diagnotorService = &diagnotorService{p: p}
	if p.Register != nil {
		type DiagnotorService = pb.DiagnotorServiceServer
		pb.RegisterDiagnotorServiceImp(p.Register, p.diagnotorService, apis.Options(), p.Perm.Check(
			perm.Method(DiagnotorService.ListServices, perm.ScopeProject, "monitor_status", perm.ActionList, p.getScopeID),
			perm.Method(DiagnotorService.StartDiagnosis, perm.ScopeProject, "monitor_status", perm.ActionCreate, p.checkScopeID),
			perm.Method(DiagnotorService.StopDiagnosis, perm.ScopeProject, "monitor_status", perm.ActionDelete, p.checkScopeID),
			perm.Method(DiagnotorService.QueryDiagnosisStatus, perm.ScopeProject, "monitor_status", perm.ActionGet, p.checkScopeID),
			perm.Method(DiagnotorService.ListProcesses, perm.ScopeProject, "monitor_status", perm.ActionGet, p.checkScopeID),
		))
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.msp.apm.diagnotor.DiagnotorService" || ctx.Type() == pb.DiagnotorServiceServerType() || ctx.Type() == pb.DiagnotorServiceHandlerType():
		return p.diagnotorService
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.apm.diagnotor", &servicehub.Spec{
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
