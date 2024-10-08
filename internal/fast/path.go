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
	"os"
	"strconv"

	"github.com/cloudwego/gjson/internal/caching"
)

var psCache caching.IncRCU

var (
	FastPathEnable    = os.Getenv("GJSON_FAST_PATH") != ""
	FastStringEnable  = os.Getenv("GJSON_FAST_STRING") == "1" || os.Getenv("GJSON_FAST_STRING") == "2"
	ValidStringEnable = os.Getenv("GJSON_FAST_STRING") == "2"
)

func init() {
	m := map[string]interface{}{}
	psCache = caching.NewRCUStr(&m, 1000)
}

type parsedPaths struct {
	simples []interface{}
}

// check if it is a simple path and return parsed values.
// Now only a path contains only alphabets, '_' and numbers can be regarded as simple
// TODO: support more delicated simpliable cases
func FastPaths(path string) []interface{} {
	if !FastPathEnable {
		return nil
	}
	// read cache first
	ps := psCache.GetByStr(path)
	if ps != nil {
		pp, ok := ps.(parsedPaths)
		if !ok {
			return nil
		}
		return pp.simples
	}

	// parse paths
	var paths = []interface{}{}
	var s = 0
	var isNum = true
	for i := 0; i < len(path); i++ {
		c := path[i]
		if c == '.' {
			if err := add_sub(&paths, path[s:i], isNum); err != nil {
				paths = nil
				goto set_paths
			}
			if i+1 >= len(path) {
				paths = nil
				goto set_paths
			}
			s = i + 1
			isNum = true
			continue
		}
		is := isNameChar(c)
		if !is {
			paths = nil
			goto set_paths
		}
		isNum = isNum && isNumCahr(c)
		// escaped chars
		// if c == '\\' && i+1 < len(path) && escapableChars[path[i+1]] {
		// 	i++
		// 	continue
		// }
		// if complexChars[c] {
		// 	paths = nil
		// 	break
		// }
	}
	if err := add_sub(&paths, path[s:], isNum); err != nil {
		paths = nil
	}

	// set to cache
set_paths:
	_ = psCache.SetByStr(path, parsedPaths{paths})
	return paths
}

func add_sub(paths *[]interface{}, sub string, isNum bool) error {
	if isNum {
		index, err := strconv.ParseInt(sub, 10, 64)
		if err != nil {
			return err
		} else if index < 0 {
			return errors.New("negative index")
		}
		*paths = append(*paths, int(index))
	} else {
		*paths = append(*paths, sub)
	}
	return nil
}

func isNameChar(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '_'
}

func isNumCahr(c byte) bool {
	return ('0' <= c && c <= '9')
}

// var escapableChars = [256]bool{
// 	'\\': true,
// 	'.':  true,
// 	'*':  true,
// 	'?':  true,
// }

// var complexChars = [256]bool{
// 	'*': true,
// 	'"': true,
// 	'%': true,
// 	'?': true,
// 	'#': true,
// 	'|': true,
// 	'@': true,
// 	'[': true,
// 	']': true,
// 	'{': true,
// 	'}': true,
// 	'(': true,
// 	')': true,
// 	',': true,
// 	'~': true,
// 	'!': true,
// }
