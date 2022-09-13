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

package executeHistoryButton

import (
	"context"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cpregister/base"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/utils/cputil"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/component-protocol/components/auto-test-scenes/common/gshelper"
)

type ComponentAction struct {
}

func getProps(ctx context.Context, visible bool) map[string]interface{} {
	return map[string]interface{}{
		"text":    cputil.I18n(ctx, "execHistory"),
		"visible": visible,
	}
}

func (ca *ComponentAction) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	gh := gshelper.NewGSHelper(gs)
	c.Props = getProps(ctx, gh.GetExecuteTaskBreadcrumbVisible())
	return nil
}

func init() {
	base.InitProviderWithCreator("auto-test-scenes", "executeHistoryButton",
		func() servicehub.Provider { return &ComponentAction{} })
}
