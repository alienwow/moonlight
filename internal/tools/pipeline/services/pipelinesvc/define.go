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

package pipelinesvc

import (
	"github.com/ping-cloudnative/moonlight-utils/providers/mysqlxorm"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/pkg/websocket"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/dbclient"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/actionagent"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/actionmgr"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/app"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/cache"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/clusterinfo"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/cron/daemon"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/edgepipeline_register"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/edgereporter"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/engine"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/permission"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/queuemanager"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/resource"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/run"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/secret"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/user"
	"github.com/ping-cloudnative/moonlight/pkg/jsonstore"
	"github.com/ping-cloudnative/moonlight/pkg/jsonstore/etcd"
	"github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cms/pb"
	cronpb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cron/pb"
)

type PipelineSvc struct {
	appSvc          app.Interface
	crondSvc        daemon.Interface
	actionAgentSvc  actionagent.Interface
	pipelineCronSvc cronpb.CronServiceServer
	permissionSvc   permission.Interface
	queueManage     queuemanager.Interface
	cache           cache.Interface

	dbClient  *dbclient.Client
	bdl       *bundle.Bundle
	publisher *websocket.Publisher

	engine engine.Interface

	js      jsonstore.JsonStore
	etcdctl *etcd.Store

	// providers
	cmsService   pb.CmsServiceServer
	clusterInfo  clusterinfo.Interface
	edgeRegister edgepipeline_register.Interface
	edgeReporter edgereporter.Interface
	secret       secret.Interface
	user         user.Interface
	run          run.Interface
	actionMgr    actionmgr.Interface
	mysql        mysqlxorm.Interface
	resource     resource.Interface
}

func New(appSvc app.Interface, crondSvc daemon.Interface,
	actionAgentSvc actionagent.Interface,
	pipelineCronSvc cronpb.CronServiceServer, permissionSvc permission.Interface,
	queueManage queuemanager.Interface,
	dbClient *dbclient.Client, bdl *bundle.Bundle, publisher *websocket.Publisher,
	engine engine.Interface, js jsonstore.JsonStore, etcd *etcd.Store, clusterInfo clusterinfo.Interface,
	edgeRegister edgepipeline_register.Interface, cache cache.Interface, resource resource.Interface) *PipelineSvc {

	s := PipelineSvc{}
	s.appSvc = appSvc
	s.crondSvc = crondSvc
	s.actionAgentSvc = actionAgentSvc
	s.pipelineCronSvc = pipelineCronSvc
	s.permissionSvc = permissionSvc
	s.queueManage = queueManage
	s.dbClient = dbClient
	s.bdl = bdl
	s.publisher = publisher
	s.engine = engine
	s.js = js
	s.etcdctl = etcd
	s.clusterInfo = clusterInfo
	s.edgeRegister = edgeRegister
	s.cache = cache
	s.resource = resource
	return &s
}

func (s *PipelineSvc) WithCmsService(cmsService pb.CmsServiceServer) {
	s.cmsService = cmsService
}

func (s *PipelineSvc) WithSecret(secret secret.Interface) {
	s.secret = secret
}

func (s *PipelineSvc) WithUser(user user.Interface) {
	s.user = user
}

func (s *PipelineSvc) WithRun(run run.Interface) {
	s.run = run
}

func (s *PipelineSvc) WithActionMgr(actionMgr actionmgr.Interface) {
	s.actionMgr = actionMgr
}

func (s *PipelineSvc) WithMySQL(mysql mysqlxorm.Interface) {
	s.mysql = mysql
}

func (s *PipelineSvc) WithEdgeReporter(r edgereporter.Interface) {
	s.edgeReporter = r
}
