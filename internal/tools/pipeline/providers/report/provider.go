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

package report

import (
	"context"
	"reflect"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight-utils/providers/mysqlxorm"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/dbclient"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/report/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg      *config
	Log      logs.Logger
	Register transport.Register
	*reportService

	MySQL mysqlxorm.Interface
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.reportService = &reportService{p, &dbclient.Client{p.MySQL.DB()}}
	if p.Register != nil {
		pb.RegisterReportServiceImp(p.Register, p.reportService, apis.Options())
	}
	return nil
}

func (p *provider) Run(ctx context.Context) error {
	return nil
}

func init() {
	interfaceType := reflect.TypeOf((*Interface)(nil)).Elem()
	servicehub.Register("erda.core.pipeline.report", &servicehub.Spec{
		Types:                []reflect.Type{interfaceType},
		OptionalDependencies: []string{"service-register"},
		Description:          "pipeline report",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
