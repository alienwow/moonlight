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

package project

import (
	"time"

	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/dbclient"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/services/namespace"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/pkg/cache"
	dashboardPb "github.com/ping-cloudnative/moonlight/proto-go/cmp/dashboard/pb"
	clusterpb "github.com/ping-cloudnative/moonlight/proto-go/core/clustermanager/cluster/pb"
	tokenpb "github.com/ping-cloudnative/moonlight/proto-go/core/token/pb"
)

type Project struct {
	db         *dbclient.DBClient
	bdl        *bundle.Bundle
	trans      i18n.Translator
	cmp        dashboardPb.ClusterResourceServer
	namespace  *namespace.Namespace
	clusterSvc clusterpb.ClusterServiceServer

	appOwnerCache    *cache.Cache
	CreateFileRecord func(req apistructs.TestFileRecordRequest) (uint64, error)
	UpdateFileRecord func(req apistructs.TestFileRecordRequest) error
	tokenService     tokenpb.TokenServiceServer
	org              org.Interface
}

func New(options ...Option) *Project {
	p := new(Project)
	for _, f := range options {
		f(p)
	}
	p.appOwnerCache = cache.New("ApplicationOwnerCache", time.Minute, p.updateMemberCache)
	return p
}
