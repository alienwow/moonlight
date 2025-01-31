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

package edgepipeline_register

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ping-cloudnative/moonlight/apistructs"
)

func TestRegisterEventHandler(t *testing.T) {
	p := &provider{}
	p.eventHandlers = make([]EventHandler, 0)
	p.edgeClients = make(map[string]apistructs.ClusterManagerClientDetail)

	p.RegisterEventHandler(func(ctx context.Context, eventDetail apistructs.ClusterManagerClientDetail) {
		t.Log("event handler called")
	})
	p.RegisterEventHandler(func(ctx context.Context, eventDetail apistructs.ClusterManagerClientDetail) {
		t.Log("event handler called")
	})
	p.emitClientEvent(context.Background(), apistructs.ClusterManagerClientDetail{})
}

func Test_updateClientByEvent(t *testing.T) {
	p := &provider{}
	p.eventHandlers = make([]EventHandler, 0)
	p.edgeClients = make(map[string]apistructs.ClusterManagerClientDetail)

	p.updateClientByEvent(&apistructs.ClusterManagerClientEvent{
		Content: apistructs.ClusterManagerClientDetail{
			apistructs.ClusterManagerDataKeyClusterKey: "erda",
		},
	})

	assert.Equal(t, apistructs.ClusterManagerClientDetail{
		apistructs.ClusterManagerDataKeyClusterKey: "erda",
	}, p.edgeClients["erda"])
}
