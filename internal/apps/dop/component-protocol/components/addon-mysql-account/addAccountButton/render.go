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

package addAccountButton

import (
	"context"

	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cpregister/base"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/utils/cputil"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/component-protocol/components/addon-mysql-account/common"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/component-protocol/types"
	"github.com/ping-cloudnative/moonlight/internal/tools/monitor/utils"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	addonmysqlpb "github.com/ping-cloudnative/moonlight/proto-go/orchestrator/addon/mysql/pb"
)

type comp struct {
}

func init() {
	base.InitProviderWithCreator("addon-mysql-account", "addAccountButton",
		func() servicehub.Provider { return &comp{} })
}

func (f *comp) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	userID := apis.GetUserID(ctx)
	ac, err := common.LoadAccountData(ctx)
	if err != nil {
		return err
	}
	props := map[string]interface{}{
		"text":          cputil.I18n(ctx, "one_click_create"),
		"type":          "primary",
		"requestIgnore": []string{"props", "data", "operations"},
		"tipProps": map[string]interface{}{
			"placement": "left",
		},
	}
	if !ac.EditPerm {
		props["disabled"] = true
		props["disabledTip"] = cputil.I18n(ctx, "create_no_perm_tip")
	}
	c.Props = props
	c.Operations = map[string]interface{}{
		"click": cptype.LegacyOperation{
			Key:    "addAccount",
			Reload: true,
		},
	}
	addonMySQLSvc := ctx.Value(types.AddonMySQLService).(addonmysqlpb.AddonMySQLServiceServer)
	pg := common.LoadPageDataAccount(ctx)
	switch event.Operation {
	case "addAccount":
		_, err := addonMySQLSvc.GenerateMySQLAccount(utils.NewContextWithHeader(ctx), &addonmysqlpb.GenerateMySQLAccountRequest{
			InstanceId: pg.InstanceID,
			UserID:     userID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
