/**
 * Copyright 2024 Cloudwego Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fast

import (
	"errors"

	"github.com/bytedance/sonic/ast"
)

// Get search paths from p, and returns the start and end position of that value
func Get(s string, path ...interface{}) (start int, end int, typ int, err error) {
	return getByPath(s, path...)
}

func errJSON(json string) error {
	return errors.New("invalid json " + json)
}

func Skip(src string, i int) (end int, remian string) {
	s, e, err := skipFast(src, i)
	if err != nil {
		return i + 1, src[i : i+1]
	}
	return e, src[s:e]
}

func String(json string, i int) (end int, str string, hasEsc bool, error error) {
	v, r, hasEsc := decodeString(json, i, false, ValidStringEnable)
	if r < 0 {
		return i + 1, "", false, errJSON(json)
	}
	return r, v, hasEsc, nil
}

func Unquote(str string) (string, error) {
	out, err := unquote(str, false)
	if err != nil {
		return "", err
	}
	return out, nil
}

func JSONType(sonic int) int {
	switch sonic {
	case ast.V_NULL:
		return 0
	case ast.V_FALSE:
		return 1
	case ast.V_TRUE:
		return 4
	case ast.V_NUMBER:
		return 2
	case ast.V_STRING:
		return 3
	case ast.V_ARRAY:
		return 5
	case ast.V_OBJECT:
		return 5
	default:
		return 0
	}
}

func Valid(json string) bool {
	return validSyntax(json)
}
