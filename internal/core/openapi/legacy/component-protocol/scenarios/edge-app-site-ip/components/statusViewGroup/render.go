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

package statusviewgroup

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ping-cloudnative/moonlight/apistructs"
	protocol "github.com/ping-cloudnative/moonlight/internal/core/openapi/legacy/component-protocol"
	"github.com/ping-cloudnative/moonlight/internal/core/openapi/legacy/component-protocol/scenarios/edge-app-site-ip/i18n"
	i18r "github.com/ping-cloudnative/moonlight/pkg/i18n"
)

func (c ComponentViewGroup) Render(ctx context.Context, component *apistructs.Component, scenario apistructs.ComponentProtocolScenario, event apistructs.ComponentEvent, gs *apistructs.GlobalStateData) error {

	bdl := ctx.Value(protocol.GlobalInnerKeyCtxBundle.String()).(protocol.ContextBundle)

	if err := c.SetBundle(bdl); err != nil {
		return err
	}
	if err := c.SetComponent(component); err != nil {
		return err
	}

	orgID, err := strconv.ParseInt(c.ctxBundle.Identity.OrgID, 10, 64)
	if err != nil {
		return fmt.Errorf("component %s parse org id error: %v", component.Name, err)
	}

	if c.component.State == nil {
		c.component.State = map[string]interface{}{}
	}

	identity := c.ctxBundle.Identity

	if event.Operation == apistructs.EdgeOperationChangeRadio || event.Operation == apistructs.InitializeOperation {
		err = c.OperationChangeViewGroup()
		if err != nil {
			return err
		}

		err = c.Operation(orgID, identity)
		if err != nil {
			return err
		}

		c.component.Operations = getOperations()
	}

	return nil
}

func getProps(total, success, error int, lr *i18r.LocaleResource) apistructs.EdgeRadioProps {
	return apistructs.EdgeRadioProps{
		RadioType:   "button",
		ButtonStyle: "outline",
		Size:        "small",
		Options: []apistructs.EdgeButtonOption{
			{Text: fmt.Sprintf("%s(%d)", lr.Get(i18n.I18nKeyAll), total), Key: "total"},
			{Text: fmt.Sprintf("%s(%d)", lr.Get(i18n.I18nKeyRunning), success), Key: "success"},
			{Text: fmt.Sprintf("%s(%d)", lr.Get(i18n.I18nKeyStopped), error), Key: "error"},
		},
	}
}

func getOperations() apistructs.EdgeOperations {
	return apistructs.EdgeOperations{
		"onChange": apistructs.EdgeOperation{
			Key:    apistructs.EdgeOperationChangeRadio,
			Reload: true,
		},
	}
}
