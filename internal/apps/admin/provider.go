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

package admin

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	hs "github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/internal/apps/admin/dao"
	"github.com/ping-cloudnative/moonlight/internal/apps/admin/manager"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/pkg/database/dbengine"
	"github.com/ping-cloudnative/moonlight/pkg/http/httpserver"
	clusterpb "github.com/ping-cloudnative/moonlight/proto-go/core/clustermanager/cluster/pb"
)

type Config struct {
	Debug bool `default:"false" env:"DEBUG" desc:"enable debug logging"`
}

type provider struct {
	Config Config

	Log    logs.Logger
	Router hs.Router       `autowired:"http-router"`
	Tran   i18n.Translator `translator:"common"`

	DB         *gorm.DB                       `autowired:"mysql-client"`
	ClusterSvc clusterpb.ClusterServiceServer `autowired:"erda.core.clustermanager.cluster.ClusterService"`
	Org        org.Interface
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.Log.Info("init admin")
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000000000",
	})
	logrus.SetOutput(os.Stdout)

	if p.Config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	db := &dao.DBClient{
		DBEngine: &dbengine.DBEngine{
			DB: p.DB,
		},
	}
	admin := manager.NewAdminManager(
		manager.WithDB(db),
		manager.WithBundle(manager.NewBundle()),
		manager.WithClusterSvc(p.ClusterSvc),
		manager.WithOrg(p.Org),
	)
	server := httpserver.New("")
	server.Router().UseEncodedPath()
	server.RegisterEndpoint(admin.Routers())
	return server.RegisterToNewHttpServerRouter(p.Router)
}

func init() {
	servicehub.Register("service.admin", &servicehub.Spec{
		Services:     []string{"admin"},
		Dependencies: []string{"http-server"},
		Description:  "erda platform admin",
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
