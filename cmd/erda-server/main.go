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

	// apps
	_ "github.com/ping-cloudnative/moonlight/internal/apps/admin"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/admin/personal-workbench"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/gallery"

	// pkg
	_ "github.com/ping-cloudnative/moonlight/internal/pkg/audit"
	_ "github.com/ping-cloudnative/moonlight/internal/pkg/dingtalktest"
	_ "github.com/ping-cloudnative/moonlight/internal/pkg/service-discover/erda-discover"
	_ "github.com/ping-cloudnative/moonlight/internal/pkg/service-discover/fixed-discover"
	"github.com/ping-cloudnative/moonlight/pkg/common"

	// core
	_ "github.com/ping-cloudnative/moonlight-utils/providers/pprof"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/redis"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/dicehub"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/dicehub/extension"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/dicehub/image"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/dicehub/release"
	_ "github.com/ping-cloudnative/moonlight/internal/core/file"
	_ "github.com/ping-cloudnative/moonlight/internal/core/legacy"
	_ "github.com/ping-cloudnative/moonlight/internal/core/legacy/providers/token"
	_ "github.com/ping-cloudnative/moonlight/internal/core/legacy/services/dingtalk/api"
	_ "github.com/ping-cloudnative/moonlight/internal/core/messenger/eventbox"
	_ "github.com/ping-cloudnative/moonlight/internal/core/messenger/notify"
	_ "github.com/ping-cloudnative/moonlight/internal/core/messenger/notify-channel"
	_ "github.com/ping-cloudnative/moonlight/internal/core/messenger/notifygroup"
	_ "github.com/ping-cloudnative/moonlight/internal/core/project"
	_ "github.com/ping-cloudnative/moonlight/internal/core/user"

	// infra
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/component-protocol"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cpregister"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/etcd"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/etcd-election"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/grpcserver"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/health"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/mysql"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/mysql/v2"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/serviceregister"

	// grpc
	_ "github.com/ping-cloudnative/moonlight-utils/providers/grpcclient"
	_ "github.com/ping-cloudnative/moonlight/proto-go/apps/gallery/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/cmp/dashboard/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/clustermanager/cluster/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/dicehub/release/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/alert/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/monitor/metric/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/org/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cms/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cron/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/definition/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/graph/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/queue/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/source/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/services/errorbox/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/token/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/user/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/msp/menu/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/msp/tenant/project/client"
	_ "github.com/ping-cloudnative/moonlight/proto-go/orchestrator/addon/mysql/client"

	// openapi
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/legacy"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/auth"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/auth/compatibility"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/auth/ory-kratos"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/auth/password"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/auth/token"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/auth/uc-session"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/example/backend"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/example/custom-register"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/example/custom-route-source"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/example/publish"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/interceptors/common"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/interceptors/csrf"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/interceptors/dump"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/interceptors/user-info"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/routes/custom"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/routes/dynamic"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/routes/dynamic/temporary"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/routes/openapi-v1"
	_ "github.com/ping-cloudnative/moonlight/internal/core/openapi/openapi-ng/routes/proto"

	// uc-adaptor
	_ "github.com/ping-cloudnative/moonlight/internal/core/user/impl/uc/uc-adaptor"

	// dop
	_ "github.com/ping-cloudnative/moonlight/internal/apps/devflow/flow"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/devflow/issuerelation"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/component-protocol/components"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/api-management"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/autotest/testplan"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/cms"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/contribution"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/devflowrule"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/guide"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/core"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/core/query"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/stream"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/stream/core"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/sync"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/pipelinetemplate"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/project/home"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/projectpipeline"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/publishitem"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/queue"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/rule"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/search"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/taskerror"

	// cmp
	_ "github.com/ping-cloudnative/moonlight/internal/apps/cmp"
	_ "github.com/ping-cloudnative/moonlight/internal/apps/cmp/component-protocol/components"
)

//go:embed bootstrap.yaml
var bootstrapCfg string

func main() {
	common.RegisterHubListener(cpregister.NewHubListener())
	common.Run(&servicehub.RunOptions{
		Content: bootstrapCfg,
	})
}
