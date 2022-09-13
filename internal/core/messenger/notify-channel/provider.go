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

package notify_channel

import (
	"github.com/jinzhu/gorm"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/core/messenger/notify-channel/db"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/proto-go/core/messenger/notifychannel/pb"
	userpb "github.com/ping-cloudnative/moonlight/proto-go/core/user/pb"
)

type config struct{}

type provider struct {
	Cfg                 *config
	Log                 logs.Logger
	Register            transport.Register
	Identity            userpb.UserServiceServer
	notifyChanelService *notifyChannelService
	bdl                 *bundle.Bundle
	DB                  *gorm.DB        `autowired:"mysql-client"`
	I18n                i18n.Translator `autowired:"i18n" translator:"cs-i18n"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.bdl = bundle.New(bundle.WithKMS(), bundle.WithErdaServer())
	p.notifyChanelService = &notifyChannelService{
		p:               p,
		NotifyChannelDB: &db.NotifyChannelDB{DB: p.DB},
	}
	if p.Register != nil {
		pb.RegisterNotifyChannelServiceImp(p.Register, p.notifyChanelService, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.messenger.notifychannel.NotifyChannelService" || ctx.Type() == pb.NotifyChannelServiceServerType() || ctx.Type() == pb.NotifyChannelServiceHandlerType():
		return p.notifyChanelService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.messenger.notifychannel", &servicehub.Spec{
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
