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

package notify

import (
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight-utils/providers/mysql"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/core/messenger/notify/db"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/proto-go/core/messenger/notify/pb"
)

type config struct {
}

type provider struct {
	C             *config
	Register      transport.Register `autowired:"service-register" optional:"true"`
	notifyService *notifyService
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.notifyService = &notifyService{}
	p.notifyService.DB = db.New(ctx.Service("mysql").(mysql.Interface).DB())
	p.notifyService.bdl = bundle.New(bundle.WithScheduler(), bundle.WithErdaServer())
	if p.Register != nil {
		type NotifyService = pb.NotifyServiceServer
		pb.RegisterNotifyServiceImp(p.Register, p.notifyService, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.messenger.notify.NotifyService" || ctx.Type() == pb.NotifyServiceServerType() || ctx.Type() == pb.NotifyServiceHandlerType():
		return p.notifyService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.messenger.notify", &servicehub.Spec{
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
