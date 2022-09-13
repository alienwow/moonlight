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

package filter

import (
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/components/filter/impl"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight/bundle"
)

type ComponentFilter struct {
	impl.DefaultFilter
	sdk              *cptype.SDK
	bdl              *bundle.Bundle
	State            State
	FrontendUrlQuery string
	projectID        uint64

	fixedIterationID uint64
}

type State struct {
	Base64UrlQueryParams string             `json:"filter__urlQuery,omitempty"`
	Values               FrontendConditions `json:"values,omitempty"`
}

type FrontendConditions struct {
	IterationIDs []int64  `json:"iterationIDs,omitempty"`
	AssigneeIDs  []string `json:"assignee,omitempty"`
	LabelIDs     []uint64 `json:"label,omitempty"`
}
