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

package resource_test

import (
	"context"
	"testing"

	"github.com/rancher/apiserver/pkg/types"

	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/cmp/resource"
	"github.com/ping-cloudnative/moonlight/proto-go/cmp/dashboard/pb"
)

type fakeCmp struct {
}

func (f fakeCmp) ListSteveResource(ctx context.Context, req *apistructs.SteveRequest) ([]types.APIObject, error) {
	return nil, nil
}

func (f fakeCmp) GetNamespacesResources(ctx context.Context, nReq *pb.GetNamespacesResourcesRequest) (*pb.GetNamespacesResourcesResponse, error) {
	return nil, nil
}

func TestNewReportTable(t *testing.T) {
	var bdl bundle.Bundle
	var cmp fakeCmp
	var trans i18n.Translator
	resource.NewReportTable(
		resource.ReportTableWithBundle(&bdl),
		resource.ReportTableWithCMP(cmp),
		resource.ReportTableWithTrans(trans),
	)
}
