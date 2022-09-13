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

package errors

import (
	transhttp "github.com/ping-cloudnative/moonlight-utils/pkg/transport/http"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
)

type Error interface {
	error
	transhttp.Error
	i18n.Internationalizable
}

func suffixIfNotEmpty(sep, err string) string {
	if len(err) > 0 {
		return sep + " " + err
	}
	return ""
}
