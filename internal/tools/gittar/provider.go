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

package gittar

import (
	"context"

	"github.com/coreos/etcd/clientv3"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/etcd"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	tokenpb "github.com/ping-cloudnative/moonlight/proto-go/core/token/pb"
	userpb "github.com/ping-cloudnative/moonlight/proto-go/core/user/pb"
)

type provider struct {
	ETCD         etcd.Interface             // autowired
	EtcdClient   *clientv3.Client           // autowired
	TokenService tokenpb.TokenServiceServer `autowired:"erda.core.token.TokenService"`
	Identity     userpb.UserServiceServer
	Org          org.ClientInterface
}

func (p *provider) Run(ctx context.Context) error { return p.Initialize() }

func init() {
	servicehub.Register("gittar", &servicehub.Spec{
		Services:     []string{"gittar"},
		Dependencies: []string{"etcd"},
		Creator:      func() servicehub.Provider { return &provider{} },
	})
}
