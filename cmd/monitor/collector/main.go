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
	"os"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight/pkg/common"

	// modules
	_ "github.com/ping-cloudnative/moonlight-utils/providers/health"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/kafka"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/kubernetes"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/pprof"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/prometheus"
	_ "github.com/ping-cloudnative/moonlight-utils/providers/serviceregister"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/lib/kafka"

	// providers
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/collector"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/authentication"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/interceptor"

	// grpc
	_ "github.com/ping-cloudnative/moonlight-utils/providers/grpcclient"
	_ "github.com/ping-cloudnative/moonlight/proto-go/core/token/client"

	// data pipeline
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/core"
	_ "github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/plugins/all"
)

//go:embed bootstrap.yaml
var centralBootstrapCfg string

//go:embed bootstrap-agent.yaml
var edgeBootstrapCfg string

//go:generate sh -c "cd ${PROJ_PATH} && go generate -v -x github.com/ping-cloudnative/moonlight/internal/tools/monitor/core/collector"
func main() {
	cfg := centralBootstrapCfg
	if os.Getenv("DICE_IS_EDGE") == "true" {
		cfg = edgeBootstrapCfg
	}
	common.Run(&servicehub.RunOptions{
		Content: cfg,
	})
}
