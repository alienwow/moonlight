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

package endpoints

import (
	"context"
	"net/http"

	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/services/apierrors"
	"github.com/ping-cloudnative/moonlight/internal/pkg/user"
	"github.com/ping-cloudnative/moonlight/pkg/http/httpserver"
)

func (e *Endpoints) ListAPIGateways(ctx context.Context, r *http.Request, vars map[string]string) (httpserver.Responser, error) {
	identity, err := user.GetIdentityInfo(r)
	if err != nil {
		return apierrors.CreateContract.NotLogin().ToResp(), nil
	}
	orgID, err := user.GetOrgID(r)
	if err != nil {
		return apierrors.CreateContract.MissingParameter(apierrors.MissingOrgID).ToResp(), nil
	}

	var req = apistructs.ListAPIGatewaysReq{
		OrgID:     orgID,
		Identity:  &identity,
		URIParams: &apistructs.ListAPIGatewaysURIParams{AssetID: vars[urlPathAssetID]},
	}

	list, apiError := e.assetSvc.ListAPIGateways(&req)
	if apiError != nil {
		return apiError.ToResp(), nil
	}
	return httpserver.OkResp(apistructs.ListAPIGatewayRespData{Total: len(list), List: list})
}

func (e *Endpoints) ListProjectAPIGateways(ctx context.Context, r *http.Request, vars map[string]string) (httpserver.Responser, error) {
	identity, err := user.GetIdentityInfo(r)
	if err != nil {
		return apierrors.ListAPIGateways.NotLogin().ToResp(), nil
	}
	orgID, err := user.GetOrgID(r)
	if err != nil {
		return apierrors.ListAPIGateways.MissingParameter(apierrors.MissingOrgID).ToResp(), nil
	}

	var req = apistructs.ListProjectAPIGatewaysReq{
		OrgID:     orgID,
		Identity:  &identity,
		URIParams: &apistructs.ListProjectAPIGatewaysURIParams{ProjectID: vars[urlPathProjectID]},
	}

	list, apiError := e.assetSvc.ListProjectAPIGateways(&req)
	if apiError != nil {
		return apiError.ToResp(), nil
	}
	return httpserver.OkResp(map[string]interface{}{"total": len(list), "list": list})
}
