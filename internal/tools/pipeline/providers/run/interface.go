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

package run

import (
	"context"

	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/spec"
)

type Interface interface {
	RunOnePipeline(ctx context.Context, req *apistructs.PipelineRunRequest) (*spec.Pipeline, error)
	CanManualRun(ctx context.Context, p *spec.Pipeline) (reason string, can bool)
}
