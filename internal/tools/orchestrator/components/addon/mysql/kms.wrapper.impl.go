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

package mysql

import (
	"encoding/base64"

	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/pkg/kms/kmstypes"
)

type kmsWrapperImpl struct {
	bdl *bundle.Bundle
}

func NewKMSWrapper(bdl *bundle.Bundle) KMSWrapper {
	return &kmsWrapperImpl{bdl: bdl}
}

func (k *kmsWrapperImpl) CreateKey() (*kmstypes.CreateKeyResponse, error) {
	return k.bdl.KMSCreateKey(apistructs.KMSCreateKeyRequest{
		CreateKeyRequest: kmstypes.CreateKeyRequest{
			PluginKind: kmstypes.PluginKind_DICE_KMS,
		},
	})
}

func (k *kmsWrapperImpl) Encrypt(plaintext, keyID string) (*kmstypes.EncryptResponse, error) {
	return k.bdl.KMSEncrypt(apistructs.KMSEncryptRequest{
		EncryptRequest: kmstypes.EncryptRequest{
			KeyID:           keyID,
			PlaintextBase64: base64.StdEncoding.EncodeToString([]byte(plaintext)),
		},
	})
}

func (k *kmsWrapperImpl) Decrypt(ciphertext, keyID string) (*kmstypes.DecryptResponse, error) {
	return k.bdl.KMSDecrypt(apistructs.KMSDecryptRequest{
		DecryptRequest: kmstypes.DecryptRequest{
			KeyID:            keyID,
			CiphertextBase64: ciphertext,
		},
	})
}
