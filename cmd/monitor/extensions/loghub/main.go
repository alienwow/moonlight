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
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub"
	"github.com/ping-cloudnative/moonlight/pkg/common"
	"github.com/ping-cloudnative/moonlight/pkg/common/addon"

	// providers and modules
	_ "github.com/ping-cloudnative/moonlight/internal/apps/msp/apm/log-service/analysis"

	// // log export outputs
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub/exporter"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub/exporter/output/elasticsearch"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub/exporter/output/elasticsearch-proxy"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub/exporter/output/stdout"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub/exporter/output/udp"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/extensions/loghub/index/manager"

	// infra
	_ "github.com/ping-cloudnative/moonlight-utils/providers/health"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/pprof"
)

func main() {
	common.RegisterInitializer(addon.OverrideEnvs)
	common.RegisterInitializer(loghub.Init)
	common.Run(&servicehub.RunOptions{})
}
