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

package testdata

import (
	"strconv"
	"testing"

	. "github.com/cloudwego/gjson"
	"github.com/cloudwego/gjson/internal/fast"
)

func BenchmarkGetComplexPath(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(basicJSON, `loggy.programmers.#[tag="good"]#.firstName`)
		}
	})
	b.Run("medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(TwitterJsonMedium, `statuses.#[friends_count>100]#.id`)
		}
	})
	b.Run("large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(twitterLarge, `statuses.#[friends_count>100]#.id`)
		}
	})
}

func BenchmarkGetSimplePath(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(basicJSON, `loggy.programmers.0.firstName`)
		}
	})
	b.Run("medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(TwitterJsonMedium, `statuses.3.id`)
		}
	})
	b.Run("Large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x := Get(twitterLarge, `statuses.50.id`)
			if !x.Exists() {
				b.Fatal()
			}
		}
	})
}

func BenchmarkFastPath(b *testing.B) {
	opt := fast.FastPathEnable
	b.Run("normal", func(b *testing.B) {
		fast.FastPathEnable = false
		b.Run("small", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Get(basicJSON, `loggy.programmers.0.firstName`)
			}
		})
		b.Run("medium", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Get(TwitterJsonMedium, `statuses.3.id`)
			}
		})
		b.Run("Large", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := Get(twitterLarge, `statuses.50.id`)
				if !x.Exists() {
					b.Fatal()
				}
			}
		})
	})
	b.Run("fast-path", func(b *testing.B) {
		fast.FastPathEnable = true
		b.Run("small", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Get(basicJSON, `loggy.programmers.0.firstName`)
			}
		})
		b.Run("medium", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Get(TwitterJsonMedium, `statuses.3.id`)
			}
		})
		b.Run("Large", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := Get(twitterLarge, `statuses.50.id`)
				if !x.Exists() {
					b.Fatal()
				}
			}
		})
	})
	fast.FastPathEnable = opt
}

func TestEscapeT(t *testing.T) {
	_ = Get(TwitterJsonMedium, `statuses.#[friends_count>100]#.id`)
	// _ = gjson.Get(TwitterJsonMedium, `statuses.#[friends_count>100]#.id`)
}

func BenchmarkParseString(b *testing.B) {
	opt := fast.FastStringEnable
	opt2 := fast.ValidStringEnable
	b.Run("normal", func(b *testing.B) {
		fast.FastStringEnable = false
		fast.ValidStringEnable = false
		b.Run("small", func(b *testing.B) {
			var str = `"<a href=\"//twitter.com/download/iphone%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for iPhone</a>"`
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
		b.Run("medium", func(b *testing.B) {
			var str = strconv.Quote(complicatedJSON)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)

			}
		})
		b.Run("large", func(b *testing.B) {
			var str = strconv.Quote(TwitterJsonMedium)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
	})
	b.Run("fast-string", func(b *testing.B) {
		fast.FastStringEnable = true
		b.Run("small", func(b *testing.B) {
			var str = `"<a href=\"//twitter.com/download/iphone%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for iPhone</a>"`
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
		b.Run("medium", func(b *testing.B) {
			var str = strconv.Quote(complicatedJSON)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)

			}
		})
		b.Run("large", func(b *testing.B) {
			var str = strconv.Quote(TwitterJsonMedium)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
	})
	b.Run("validate-string", func(b *testing.B) {
		fast.FastStringEnable = true
		fast.ValidStringEnable = true
		b.Run("small", func(b *testing.B) {
			var str = `"<a href=\"//twitter.com/download/iphone%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for iPhone</a>"`
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
		b.Run("medium", func(b *testing.B) {
			var str = strconv.Quote(complicatedJSON)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)

			}
		})
		b.Run("large", func(b *testing.B) {
			var str = strconv.Quote(TwitterJsonMedium)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
	})
	fast.FastStringEnable = opt
	fast.ValidStringEnable = opt2
}
