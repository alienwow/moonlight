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

package devflowrule

import (
	"context"
	"reflect"

	"gorm.io/gorm"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/devflowrule/db"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/proto-go/dop/devflowrule/pb"
)

type config struct{}

type provider struct {
	Cfg      *config
	Log      logs.Logger
	bundle   *bundle.Bundle
	DB       *gorm.DB           `autowired:"mysql-gorm.v2-client"`
	Register transport.Register `autowired:"service-register" required:"true"`
	Trans    i18n.Translator    `translator:"project-pipeline" required:"true"`
	Org      org.Interface

	dbClient *db.Client
}

func (p *provider) Run(ctx context.Context) error {
	return nil
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.bundle = bundle.New(bundle.WithGittar())
	p.dbClient = &db.Client{DB: p.DB}
	if p.Register != nil {
		pb.RegisterDevFlowRuleServiceImp(p.Register, p, apis.Options())
	}
	return nil
}

func init() {
	servicehub.Register("erda.dop.devFlowRule", &servicehub.Spec{
		Services:     []string{"erda.dop.devFlowRule"},
		Types:        []reflect.Type{reflect.TypeOf((*Interface)(nil)).Elem()},
		Dependencies: nil,
		Description:  "devFlowRule",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
