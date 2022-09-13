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

package main

import (
	_ "embed"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cpregister"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub"
	"github.com/ping-cloudnative/moonlight/pkg/common"

	// modules and providers
	_ "github.com/ping-cloudnative/moonlight-utils/providers"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/cassandra"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/messenger/notify/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/messenger/notifygroup/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/alert/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/event/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/log/query/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/metric/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/org/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/oap/entity/client"

	_ "github.com/ping-cloudnative/moonlight-utils/providers/grpcclient"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/diagnotor/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/settings/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/token/client"

	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/adapter"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/alert"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/apis"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/plugins/certificate"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/plugins/dns"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/plugins/http"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/plugins/page"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/plugins/tcp"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/storage/sync-cache"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/task"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/task/fetcher/fixed"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/checker/task/fetcher/scheduled"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/diagnotor"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/exception/query"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/log-service/query"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/log-service/rules"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/metric"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/notifygroup"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/service"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/trace/query"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/trace/storage/elasticsearch"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/configcenter"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/credential"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/instance/permission"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/member"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/menu"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/registercenter"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/coordinator"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/apigateway"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/configcenter"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/etcd"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/jvmprofiler"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/loganalytics"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/loges"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/logexporter"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/logservice"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/monitor"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/monitorcollector"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/monitorkafka"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/monitorzk"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/mysql"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/nacos"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/postgresql"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/registercenter"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/servicemesh"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/zkproxy"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/resource/deploy/handlers/zookeeper"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/tenant"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/tenant/project"
	_ "github.com/ping-cloudnative/moonlight/internal/pkg/audit"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/settings/retention-strategy"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/storekit/elasticsearch/index/cleaner"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/storekit/elasticsearch/index/loader"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/storekit/elasticsearch/index/retention-strategy"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub/index/query"
	_ "github.com/ping-cloudnative/moonlight/pkg/common/permission"

	_ "github.com/ping-cloudnative/moonlight-utils/providers/component-protocol"
	// components
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/alert/components"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/browser/components"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/service/components"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/service/datasources"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/trace/query/components"
)

//go:embed bootstrap.yaml
var bootstrapCfg string

func main() {
	common.RegisterInitializer(loghub.Init)
	common.RegisterHubListener(cpregister.NewHubListener())
	common.Run(&servicehub.RunOptions{
		Content: bootstrapCfg,
	})
}
