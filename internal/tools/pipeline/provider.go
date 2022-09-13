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

package pipeline

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/etcd"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	"github.com/ping-cloudnative/moonlight-utils/providers/mysqlxorm"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/internal/pkg/metrics/report"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/pipeline/aop/plugins"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/actionagent"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/actionmgr"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/app"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/cache"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/cancel"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/clusterinfo"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/cron/compensator"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/cron/daemon"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/dbgc"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/dispatcher"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/edgepipeline"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/edgepipeline_register"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/edgereporter"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/engine"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/leaderworker"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/permission"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/queuemanager"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/reconciler"
	reportsvc "github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/report"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/resource"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/resourcegc"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/run"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/secret"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/user"
	"github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cms/pb"
	cronpb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cron/pb"
)

type provider struct {
	CmsService     pb.CmsServiceServer      `autowired:"erda.core.pipeline.cms.CmsService"`
	MetricReport   report.MetricReport      `autowired:"metric-report-client" optional:"true"`
	Router         httpserver.Router        `autowired:"http-router"`
	CronService    cronpb.CronServiceServer `autowired:"erda.core.pipeline.cron.CronService" required:"true"`
	ReportSvc      reportsvc.Interface
	CronDaemon     daemon.Interface
	CronCompensate compensator.Interface
	MySQL          mysqlxorm.Interface `autowired:"mysql-xorm"`

	Engine       engine.Interface
	QueueManager queuemanager.Interface
	Reconciler   reconciler.Interface
	EdgePipeline edgepipeline.Interface
	EdgeRegister edgepipeline_register.Interface
	EdgeReporter edgereporter.Interface
	LeaderWorker leaderworker.Interface
	ClusterInfo  clusterinfo.Interface
	DBGC         dbgc.Interface
	ResourceGC   resourcegc.Interface
	Cache        cache.Interface
	PipelineRun  run.Interface
	Cancel       cancel.Interface
	User         user.Interface
	App          app.Interface
	Secret       secret.Interface
	ActionMgr    actionmgr.Interface
	Resource     resource.Interface
	Org          org.ClientInterface
	ActionAgent  actionagent.Interface
	Permission   permission.Interface
}

func (p *provider) Run(ctx context.Context) error {
	logrus.Infof("[alert] starting pipeline instance")
	var err error

	select {
	case <-ctx.Done():
	}
	return err
}

func init() {
	servicehub.Register("pipeline", &servicehub.Spec{
		Services:     []string{"pipeline"},
		Dependencies: []string{"etcd"},
		Creator:      func() servicehub.Provider { return &provider{} },
	})
}
