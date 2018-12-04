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

package keccak256sum

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ipfn/go-digest/digest"
)

var testDigest = digest.FromHex("9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658")

func TestBytes(t *testing.T) {
	bytes := Bytes([]byte("test"))
	assert.Equal(t, testDigest.Bytes(), bytes)
}

func TestDigest(t *testing.T) {
	hashed := Digest([]byte("test"))
	assert.Equal(t, testDigest, hashed)
}
