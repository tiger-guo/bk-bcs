/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"testing"
)

func TestCustomHeaderMatcher(t *testing.T) {
	// 自定义头字段
	ret, ok := CustomHeaderMatcher("X-Request-Id")
	if !ok || ret != "X-Request-Id" {
		t.Errorf("Excepted: ('X-Request-Id', true), Result: (%s, %v)", ret, ok)
	}

	// 标准头字段
	ret, ok = CustomHeaderMatcher("Content-Type")
	if !ok || ret != "grpcgateway-Content-Type" {
		t.Errorf("Excepted: ('grpcgateway-Content-Type', true), Result: (%s, %v)", ret, ok)
	}
}
