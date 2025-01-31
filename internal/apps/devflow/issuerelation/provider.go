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

package issuerelation

import (
	"github.com/jinzhu/gorm"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/devflow/issuerelation/db"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/pkg/database/dbengine"
	"github.com/ping-cloudnative/moonlight/proto-go/apps/devflow/issuerelation/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg                  *config
	Log                  logs.Logger
	Register             transport.Register
	issueRelationService *issueRelationService
	DB                   *gorm.DB `autowired:"mysql-client"`
	Org                  org.ClientInterface

	bdl *bundle.Bundle
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.bdl = bundle.New(bundle.WithAllAvailableClients())
	p.issueRelationService = &issueRelationService{p, &db.Client{DBEngine: &dbengine.DBEngine{
		DB: p.DB}}, p.Org}
	if p.Register != nil {
		pb.RegisterIssueRelationServiceImp(p.Register, p.issueRelationService)
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.apps.devflow.issuerelation.IssueRelationService" || ctx.Type() == pb.IssueRelationServiceServerType() || ctx.Type() == pb.IssueRelationServiceHandlerType():
		return p.issueRelationService
	}
	return p
}

func init() {
	servicehub.Register("erda.apps.devflow.issuerelation", &servicehub.Spec{
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
