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

package podsTableTitle

import (
	"golang.org/x/net/context"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cpregister/base"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/utils/cputil"
)

func init() {
	base.InitProviderWithCreator("cmp-dashboard-workload-detail", "podsTableTitle", func() servicehub.Provider {
		return &ComponentPodsTableTitle{}
	})
}

func (p *ComponentPodsTableTitle) Render(ctx context.Context, component *cptype.Component, _ cptype.Scenario,
	event cptype.ComponentEvent, _ *cptype.GlobalStateData) error {
	component.Props = cputil.MustConvertProps(Props{
		Title: cputil.I18n(ctx, "workloadPods"),
		Size:  "small",
	})
	return nil
}
