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

package log_service

import (
	"io/ioutil"
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/internal/apps/msp/instance/db"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub/index/query"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	monitorpb "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/log/query/pb"
	"github.com/ping-cloudnative/moonlight/proto-go/msp/apm/log-service/pb"
)

type config struct {
	QueryLogESEnabled  bool `file:"query_log_elasticsearch_enabled"`
	IndexFieldSettings struct {
		File            string               `file:"file"`
		DefaultSettings defaultFieldSettings `file:"default_settings"`
	} `file:"index_field_settings"`
}

type defaultFieldSettings struct {
	Fields []logField `file:"fields" yaml:"fields"`
}

type logField struct {
	FieldName          string `file:"field_name" yaml:"field_name"`
	SupportAggregation bool   `file:"support_aggregation" yaml:"support_aggregation"`
	Display            bool   `file:"display"`
	AllowEdit          bool   `file:"allow_edit" default:"true" yaml:"allow_edit"`
	Group              int32  `file:"group" yaml:"group"`
}

// +provider
type provider struct {
	Cfg                 *config
	Log                 logs.Logger
	Register            transport.Register
	logService          *logService
	MonitorLogService   monitorpb.LogQueryServiceServer `autowired:"erda.core.monitor.log.query.LogQueryService" optional:"true"`
	MonitorLogSvcClient monitorpb.LogQueryServiceClient `autowired:"erda.core.monitor.log.query.LogQueryService.client" optional:"true"`
	LoghubQuery         query.LoghubService             `autowired:"logs-index-query"`
	DB                  *gorm.DB                        `autowired:"mysql-client"`
	I18n                i18n.Translator                 `autowired:"i18n" translator:"msp-i18n"`
	Router              httpserver.Router               `autowired:"http-router"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	if len(p.Cfg.IndexFieldSettings.File) > 0 {
		f, err := ioutil.ReadFile(p.Cfg.IndexFieldSettings.File)
		if err != nil {
			return err
		}
		var defaultSettings defaultFieldSettings
		err = yaml.Unmarshal(f, &defaultSettings)
		if err != nil {
			return err
		}
		p.Cfg.IndexFieldSettings.DefaultSettings = defaultSettings
	}

	p.logService = &logService{p, &db.LogDeploymentDB{DB: p.DB}, &db.LogInstanceDB{DB: p.DB}, time.Now().UnixNano()}
	if p.Register != nil {
		pb.RegisterLogServiceImp(p.Register, p.logService, apis.Options())
	}

	p.initRoutes(p.Router)
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.msp.apm.log_service.LogService" || ctx.Type() == pb.LogServiceServerType() || ctx.Type() == pb.LogServiceHandlerType():
		return p.logService
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.apm.log_service", &servicehub.Spec{
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
