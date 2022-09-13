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

package dingtalktest

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/proto-go/pkg/dingtalktest/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg                 *config
	Log                 logs.Logger
	dingTalkTestService *dingTalkTestService
	Register            transport.Register `autowired:"service-register" optional:"true"`
}

// Run this is optional
func (p *provider) Init(ctx servicehub.Context) error {
	p.dingTalkTestService = &dingTalkTestService{
		Log: p.Log,
		bdl: bundle.New(bundle.WithErdaServer()),
	}
	if p.Register != nil {
		pb.RegisterDingTalkTestServiceImp(p.Register, p.dingTalkTestService, apis.Options())
	}
	return nil
}

func init() {
	servicehub.Register("erda.pkg.dingtalktest", &servicehub.Spec{
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
