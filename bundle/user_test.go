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

package bundle

//import (
//	"os"
//	"testing"
//
//	"github.com/davecgh/go-spew/spew"
//	"github.com/stretchr/testify/assert"
//
//	"github.com/ping-cloudnative/moonlight/apistructs"
//)
//
//func TestGetCurrentUser(t *testing.T) {
//	os.Setenv("CMDB_ADDR", "cmdb.default.svc.cluster.local:9093")
//	b := New(WithCMDB())
//	userInfo, err := b.GetCurrentUser("2")
//	assert.NoError(t, err)
//	spew.Dump(userInfo)
//}
//
//func TestListUsers(t *testing.T) {
//	os.Setenv("CMDB_ADDR", "cmdb.default.svc.cluster.local:9093")
//	b := New(WithCMDB())
//	userInfo, err := b.ListUsers(apistructs.UserListRequest{
//		Query:   "",
//		UserIDs: []string{"1", "2"},
//	})
//	assert.NoError(t, err)
//	spew.Dump(userInfo)
//}
