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

package builder

import (
	"fmt"
	"strings"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/clickhouse"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/oap/collector/core/model/odata"
)

type BuilderConfig struct {
	DataType string `file:"data_type"`
	// could be org_name, cluster_name, terminus_key
	TenantIdKey string `file:"tenant_id_key"`
}

func GetClickHouseInf(ctx servicehub.Context, dt odata.DataType) (clickhouse.Interface, error) {
	svc := ctx.Service("clickhouse@" + strings.ToLower(string(dt)))
	if svc == nil {
		svc = ctx.Service("clickhouse")
	}
	if svc == nil {
		return nil, fmt.Errorf("service clickhouse is required")
	}
	ch, ok := svc.(clickhouse.Interface)
	if !ok {
		return nil, fmt.Errorf("convert svc<%T> failed", svc)
	}
	return ch, nil
}
