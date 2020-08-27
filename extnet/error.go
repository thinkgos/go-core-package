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

package extnet

import (
	"net"
	"strings"
)

// IsErrClosed is error closed
func IsErrClosed(err error) bool {
	return err != nil && strings.Contains(err.Error(), "use of closed network connection")
}

// IsErrTimeout is net error timeout
func IsErrTimeout(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(net.Error)
	return ok && e.Timeout()
}

// IsErrTemporary is net error timeout
func IsErrTemporary(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(net.Error)
	return ok && e.Temporary()
}

// IsErrRefused is error connection refused
func IsErrRefused(err error) bool {
	return err != nil && strings.Contains(err.Error(), "connection refused")
}

// IsErrDeadline is error i/o deadline reached
func IsErrDeadline(err error) bool {
	return err != nil && strings.Contains(err.Error(), "i/o deadline reached")
}

// IsErrSocketNotConnected is error socket is not connected
func IsErrSocketNotConnected(err error) bool {
	return err != nil && strings.Contains(err.Error(), "socket is not connected")
}
