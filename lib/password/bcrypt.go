// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package password

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/thinkgos/go-core-package/internal/bytesconv"
)

var _ Verify = BCrypt{}

// BCrypt bcrypt password encryption
type BCrypt struct{}

// Hash password hash encryption
func (sf BCrypt) Hash(password, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(salt+password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return bytesconv.Bytes2Str(bytes), err
}

// Compare password hash verification
func (sf BCrypt) Compare(password, salt, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(salt+password))
}
