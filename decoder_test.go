// Copyright 2022 CloudWeGo Authors.
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

package hessian2

import (
	"testing"
)

func TestDecoder(t *testing.T) {
	javaSimpleClassEncodeByteArray := []byte{
		67, 10, 97, 46, 98, 46, 99, 46, 84, 101, 115, 116, 153, 4, 110, 97, 109, 101, 3, 97, 103, 101, 4, 97, 103, 101,
		50, 5, 119, 105, 103, 104, 116, 6, 119, 105, 103, 104, 116, 50, 4, 97, 103, 101, 51, 5, 98, 121, 116, 101, 115,
		4, 108, 105, 115, 116, 3, 109, 97, 112, 96, 9, 67, 104, 97, 110, 103, 101, 100, 101, 110, 154, 95, 0, 0, 44,
		236, 236, 84, 74, 0, 0, 1, 136, 35, 86, 112, 163, 35, 1, 0, 6, 113, 6, 77, 97, 105, 110, 36, 49, 96, 1, 116,
		73, 100, 99, 40, 143, 68, 65, 164, 19, 212, 224, 0, 0, 0, 76, 0, 0, 1, 136, 35, 86, 112, 163, 70, 74, 0, 0, 1,
		136, 35, 86, 112, 163, 35, 1, 0, 6, 112, 12, 97, 46, 98, 46, 99, 46, 84, 101, 115, 116, 36, 49, 77, 12, 97, 46,
		98, 46, 99, 46, 84, 101, 115, 116, 36, 50, 90, 77, 6, 77, 97, 105, 110, 36, 50, 1, 84, 96, 1, 116, 73, 100, 99,
		40, 143, 68, 65, 164, 19, 212, 224, 0, 0, 0, 76, 0, 0, 1, 136, 35, 86, 112, 163, 70, 74, 0, 0, 1, 136, 35, 86,
		112, 163, 35, 1, 0, 6, 112, 145, 77, 146, 90, 90,
	}

	decoder := NewDecoderWithByteArray(javaSimpleClassEncodeByteArray)
	obj := decoder.ReadObject()
	switch vc := obj.(type) {
	case *VirtualClass:
		if vc.JavaClassPackage() != "a.b.c" {
			t.Errorf("包名解析错误, 期望: %s", "a.b.c")
		}
		if vc.JavaClassName() != "Test" {
			t.Errorf("类名解析错误, 期望: %s", "Test")
		}

	default:
		t.Error("解析错误")
	}
}
