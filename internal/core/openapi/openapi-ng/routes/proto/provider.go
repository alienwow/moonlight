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

package proto

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	transhttp "github.com/ping-cloudnative/moonlight-utils/pkg/transport/http"
	"github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng"
	"github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/auth"
	"github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/proxy"
	discover "github.com/ping-cloudnative/moonlight/internal/pkg/service-discover"
	common "github.com/ping-cloudnative/moonlight/proto-go/common/pb"
)

// +provider
type provider struct {
	Log      logs.Logger
	Discover discover.Interface `autowired:"discover"`
	Auth     auth.Interface     `autowired:"openapi-auth"`
	Router   openapi.Interface  `autowired:"openapi-router"`

	proxy proxy.Proxy
}

func (p *provider) Init(ctx servicehub.Context) (err error) {
	p.proxy = proxy.Proxy{
		Log:      p.Log,
		Discover: p.Discover,
	}
	return p.RegisterTo(p.Router)
}

func (p *provider) RegisterTo(router transhttp.Router) (err error) {
	var oneOpenAPIProxyHandler OneOpenAPIProxyHandler = func(method, publishPath, backendPath, serviceName string, opt *common.OpenAPIOption) error {
		handler, err := p.proxy.Wrap(method, publishPath, backendPath, serviceName)
		if err != nil {
			return err
		}
		handler = p.Auth.Interceptor(handler, GetAuthOption(opt.Auth, func(opts map[string]interface{}) map[string]interface{} {
			opts["path"] = publishPath
			opts["method"] = method
			return opts
		}))
		router.Add(method, publishPath, transhttp.HandlerFunc(handler))
		return nil
	}
	return RangeOpenAPIsProxy(oneOpenAPIProxyHandler)
}

func init() {
	servicehub.Register("openapi-protobuf-routes", &servicehub.Spec{
		Creator: func() servicehub.Provider { return &provider{} },
	})
}
