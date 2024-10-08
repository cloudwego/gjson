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
	"strconv"
	"testing"
)

func TestGetFast(t *testing.T) {
	type args struct {
		s    string
		p    int
		path []interface{}
	}
	tests := []struct {
		name  string
		args  args
		start int
		end   int
	}{
		{"", args{`{"a":1}`, 0, []interface{}{"a"}}, 5, 6},
		{"", args{`{"b":[],"c":{},"a":[0,1,2,3,4,5,6,7,8,9]}`, 0, []interface{}{"a", 9}}, 38, 39},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, e, _, err := Get(tt.args.s, tt.args.path...)
			if err != nil {
				t.Fatal(err)
			}
			if s != tt.start {
				t.Fatal(s)
			}
			if e != tt.end {
				t.Fatal(e)
			}
		})
	}
}

func BenchmarkGet(b *testing.B) {
	var js = `{"b":[],"c":{},"a":[0,1,2,3,4,5,6,7,8,9]}`
	b.Run("Wrapper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _, _, _ = Get(js, "a", 9)
		}
	})
}

func TestString(t *testing.T) {
	type args struct {
		json string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		hasEsc  bool
		wantErr bool
	}{
		{"", args{`"abc"`}, "abc", false, false},
		{"", args{`\u263a`}, ``, false, true},
		{"", args{`"\u263a"`}, `\u263a`, true, false},
		{"", args{`"中文"`}, `中文`, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got, hasEsc, err := String(tt.args.json, 0)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if hasEsc != tt.hasEsc {
				t.Fatal()
			}
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkString(b *testing.B) {
	var js = `"abc中文"`
	b.Run("Wrapper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _, _, _ = String(js, 0)
		}
	})
	b.Run("std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.Unquote(js)
		}
	})
}
