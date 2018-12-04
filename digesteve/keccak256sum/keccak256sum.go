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
	keccak "github.com/gxed/hashland/keccakpg"

	"github.com/ipfn/go-digest/digest"
)

// Bytes - Sums Keccak256 secure hash.
func Bytes(data ...[]byte) []byte {
	return digest.SumBytes(keccak.New256(), data...)
}

// Digest - Sums Keccak256 secure hash.
func Digest(data ...[]byte) digest.Digest {
	return digest.Sum(keccak.New256(), data...)
}
