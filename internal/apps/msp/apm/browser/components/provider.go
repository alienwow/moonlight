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

package components

import (
	"embed"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	componentprotocol "github.com/ping-cloudnative/moonlight-utils/providers/component-protocol"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/protocol"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"

	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/browser/components/browser-overview"
)

//go:embed scenarios
var scenarioFS embed.FS

type config struct {
}

type provider struct {
	Cfg *config
	Log logs.Logger

	Protocol componentprotocol.Interface
	CPTran   i18n.I18n `autowired:"i18n"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.Protocol.SetI18nTran(p.CPTran)
	protocol.MustRegisterProtocolsFromFS(scenarioFS)
	return nil
}

func init() {
	servicehub.Register("browser-components", &servicehub.Spec{
		Services:    []string{"browser-components"},
		Description: "browser-components",
		ConfigFunc:  func() interface{} { return &config{} },
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
