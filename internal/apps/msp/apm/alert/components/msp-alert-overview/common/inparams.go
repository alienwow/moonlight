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

package common

import (
	"github.com/mitchellh/mapstructure"

	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
)

type InParams struct {
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Scope     string `json:"scope"`
	ScopeId   string `json:"scopeId"`
}

func ParseFromCpSdk(sdk *cptype.SDK) (*InParams, error) {
	var param InParams
	err := mapstructure.Decode(sdk.InParams, &param)
	if err != nil {
		return nil, err
	}
	return &param, nil
}
