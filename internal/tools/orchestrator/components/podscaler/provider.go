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

package podscaler

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/internal/pkg/audit"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/events"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/scheduler/impl/servicegroup"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/proto-go/orchestrator/podscaler/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg             *config
	Log             logs.Logger
	Register        transport.Register
	hpscalerService pb.PodScalerServiceServer //`autowired:"erda.orchestrator.podscaler.PodScalerService"`
	audit           audit.Auditor

	DB           *gorm.DB             `autowired:"mysql-client"`
	EventManager *events.EventManager `autowired:"erda.orchestrator.events.event-manager"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.audit = audit.GetAuditor(ctx)
	p.hpscalerService = NewRuntimeHPScalerService(
		WithBundleService(NewBundleService()),
		WithDBService(NewDBService(p.DB)),
		WithServiceGroupImpl(servicegroup.NewServiceGroupImplInit()),
	)

	if p.Register != nil {
		type HPScalerService = pb.PodScalerServiceServer
		pb.RegisterPodScalerServiceImp(p.Register, p.hpscalerService, apis.Options(), p.audit.Audit(
			audit.Method(HPScalerService.CreateRuntimeHPARules, audit.AppScope, string(apistructs.CreateAndApplyHPARule),
				func(ctx context.Context, req, resp interface{}, err error) (interface{}, map[string]interface{}, error) {
					r := req.(*pb.HPARuleCreateRequest)
					services := make([]string, 0)
					for _, svc := range r.Services {
						services = append(services, svc.ServiceName)
					}
					return services, map[string]interface{}{
						"runtimeId": r.RuntimeID,
					}, nil
				}),
			audit.Method(HPScalerService.UpdateRuntimeHPARules, audit.AppScope, string(apistructs.UpdateHPARule),
				func(ctx context.Context, req, resp interface{}, err error) (interface{}, map[string]interface{}, error) {
					r := req.(*pb.ErdaRuntimeHPARules)
					services := make([]string, 0)
					serviceToRule := make(map[string]interface{})
					serviceToRule["runtimeId"] = r.RuntimeID
					for _, rule := range r.Rules {
						services = append(services, rule.ServiceName)
						serviceToRule[rule.ServiceName] = rule.RuleID
					}
					return services, serviceToRule, nil
					//return nil, map[string]interface{}{}, nil
				}),
			audit.Method(HPScalerService.DeleteHPARulesByIds, audit.AppScope, string(apistructs.DeleteHPARule),
				func(ctx context.Context, req, resp interface{}, err error) (interface{}, map[string]interface{}, error) {
					r := req.(*pb.DeleteRuntimePARulesRequest)
					rules := make([]string, 0)
					for _, ruleId := range r.Rules {
						rules = append(rules, ruleId)
					}
					return rules, map[string]interface{}{
						"runtimeId": r.RuntimeID,
					}, nil
				}),
			audit.Method(HPScalerService.ApplyOrCancelHPARulesByIds, audit.AppScope, string(apistructs.ApplyOrCancelHPARule),
				func(ctx context.Context, req, resp interface{}, err error) (interface{}, map[string]interface{}, error) {
					r := req.(*pb.ApplyOrCancelPARulesRequest)
					rules := make([]string, 0)
					serviceToRule := make(map[string]interface{})
					serviceToRule["runtimeId"] = r.RuntimeID
					for _, ruleAction := range r.RuleAction {
						rules = append(rules, ruleAction.RuleId)
						serviceToRule[ruleAction.RuleId] = ruleAction.Action
					}
					return rules, serviceToRule, nil
					//return nil, map[string]interface{}{}, nil
				}),
		))
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.orchestrator.podscaler.HPScalerService" || ctx.Type() == pb.PodScalerServiceServerType() || ctx.Type() == pb.PodScalerServiceHandlerType():
		return p.hpscalerService
	}
	return p
}

func init() {
	servicehub.Register("erda.orchestrator.podscaler", &servicehub.Spec{
		Services: pb.ServiceNames(),
		Types:    pb.Types(),
		OptionalDependencies: []string{
			"erda.orchestrator.events",
			"service-register",
			"mysql",
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
