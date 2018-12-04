// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sha256sum

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ipfn/go-digest/digest"
)

var testDigest = digest.FromHex("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08")

func TestBytes(t *testing.T) {
	bytes := Bytes([]byte("test"))
	assert.Equal(t, testDigest.Bytes(), bytes)
}

func TestDigest(t *testing.T) {
	hashed := Digest([]byte("test"))
	assert.Equal(t, testDigest, hashed)
}
