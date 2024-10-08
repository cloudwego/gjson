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

package caching

import (
	"math/rand"
	"reflect"
	"sync"
	"testing"
)

func TestIncRCUI64(t *testing.T) {
	N := 100000
	M := 100
	D := 100

	m := map[int64]interface{}{}
	rcu := NewRCUI64(&m, M)
	exp := map[int64]interface{}{}

	type pair struct {
		k int64
		v int
	}
	kchan := make(chan pair, M)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < N; i++ {
			k := rand.Int63n(int64(D))
			if _, ok := exp[k]; !ok {
				exp[k] = i
			}
			kchan <- pair{k, i}
		}
	}()

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pair := <-kchan
			rcu.SetByI64(pair.k, i)
		}(i)
	}

	wg.Wait()

	reflect.DeepEqual(exp, rcu.DumpI64())
}

func TestIncRCUStr(t *testing.T) {
	N := 10000
	M := 100
	Dbuf := "0123456789"

	m := map[string]interface{}{}
	rcu := NewRCUStr(&m, M)
	exp := map[string]interface{}{}

	type pair struct {
		k string
		v int
	}
	kchan := make(chan pair, M)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < N; i++ {
			k := Dbuf[:rand.Intn(len(Dbuf))]
			if _, ok := exp[k]; !ok {
				exp[k] = i
			}
			kchan <- pair{k, i}
		}
	}()

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pair := <-kchan
			if !rcu.SetByStr(pair.k, pair.v) {
				println(pair.k, pair.v)
			}
		}(i)
	}

	wg.Wait()
	reflect.DeepEqual(exp, rcu.DumpStr())
}
