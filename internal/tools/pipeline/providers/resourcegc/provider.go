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

package resourcegc

import (
	"context"
	"reflect"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/mysqlxorm"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/dbclient"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/dbgc"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/leaderworker"
	"github.com/ping-cloudnative/moonlight/pkg/jsonstore"
	"github.com/ping-cloudnative/moonlight/pkg/jsonstore/etcd"
)

type config struct{}

type provider struct {
	js       jsonstore.JsonStore
	etcd     *etcd.Store
	dbClient *dbclient.Client

	Cfg   *config
	Log   logs.Logger
	MySQL mysqlxorm.Interface
	LW    leaderworker.Interface
	DBGC  dbgc.Interface
}

func (r *provider) Init(ctx servicehub.Context) error {
	// dbclient
	r.dbClient = &dbclient.Client{Engine: r.MySQL.DB()}
	js, err := jsonstore.New()
	if err != nil {
		return err
	}
	etcdClient, err := etcd.New()
	if err != nil {
		return err
	}
	r.js = js
	r.etcd = etcdClient
	return nil
}

func (r *provider) Run(ctx context.Context) error {
	// gc
	r.LW.OnLeader(r.listenGC)
	r.LW.OnLeader(r.compensateGCNamespaces)
	return nil
}

func init() {
	servicehub.Register("resourcegc", &servicehub.Spec{
		Services:     []string{"resourcegc"},
		Types:        []reflect.Type{reflect.TypeOf((*Interface)(nil)).Elem()},
		Dependencies: nil,
		Description:  "pipeline resourcegc",
		ConfigFunc:   func() interface{} { return &config{} },
		Creator:      func() servicehub.Provider { return &provider{} },
	})
}
