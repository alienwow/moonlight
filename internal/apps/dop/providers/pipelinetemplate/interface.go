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

package pipelinetemplate

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	dbclient "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/pipelinetemplate/db"
	"github.com/ping-cloudnative/moonlight/proto-go/dop/pipelinetemplate/pb"
)

type Interface interface {
	pb.TemplateServiceServer
}

type ServiceImpl struct {
	log logs.Logger
	db  *dbclient.DBClient
}
