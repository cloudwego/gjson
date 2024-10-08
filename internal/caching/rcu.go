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
	"sync"
	"sync/atomic"
	"unsafe"
)

// IncRCU wraps normal map as
// an incremental RCU (Read-Copy-Update) cache
//
// WARN: It's only used for **Fixed** key-val relation,
// changing a value mapped to same key is **NOT** allowed
type IncRCU struct {
	// original pointer of the map
	//
	// TODO: it only accepts map[int64]interface{}
	// and map[string]interface{} now
	p unsafe.Pointer

	// maxCAS means the max Catch-And-Swap times when Copy-Update.
	maxCAS int

	// If maxCAS gets exceeded when updating,
	// we will use Read-Write-Lock way to store key-val
	mux sync.RWMutex
	bp  unsafe.Pointer
}

// NewRCUI64 accepts map[int]interface{} as RCU
//
// WARN: this API can be only used on the IncRCU returned by `NewRCUI64()`
func NewRCUI64(p *map[int64]interface{}, maxCAS int) IncRCU {
	bp := make(map[int64]interface{})
	return IncRCU{
		maxCAS: maxCAS,
		p:      unsafe.Pointer(p),
		bp:     unsafe.Pointer(&bp),
	}
}

// DumpStr dumps all the key-values in the cache
func (self *IncRCU) DumpI64() (ret map[int64]interface{}) {
	m := (*(*map[int64]interface{})(atomic.LoadPointer(&self.p)))
	ret = make(map[int64]interface{}, len(m))

	for k, v := range m {
		ret[k] = v
	}

	self.mux.Lock()
	defer self.mux.Unlock()

	n := (*(*map[int64]interface{})(unsafe.Pointer(self.bp)))
	for k, v := range n {
		ret[k] = v
	}

	return
}

// NewRCUI64 accepts map[int]interface{} as RCU
//
// WARN: this API can be only used on the IncRCU returned by `NewRCUStr()`
func NewRCUStr(p *map[string]interface{}, maxCAS int) IncRCU {
	bp := make(map[string]interface{})
	return IncRCU{
		p:      unsafe.Pointer(p),
		maxCAS: maxCAS,
		bp:     unsafe.Pointer(&bp),
	}
}

// DumpStr dumps all the key-values in the cache
func (self *IncRCU) DumpStr() (ret map[string]interface{}) {
	m := (*(*map[string]interface{})(atomic.LoadPointer(&self.p)))
	ret = make(map[string]interface{}, len(m))

	for k, v := range m {
		ret[k] = v
	}

	self.mux.Lock()
	defer self.mux.Unlock()

	n := (*(*map[string]interface{})(unsafe.Pointer(self.bp)))
	for k, v := range n {
		ret[k] = v
	}

	return
}

// GetByI64 returns corresponding val of int64 id
//
// WARN: this API can be only used on the IncRCU returned by `NewRCUI64()`
func (self *IncRCU) GetByI64(id int64) interface{} {
	// try RCU first
	if v := (*(*map[int64]interface{})(atomic.LoadPointer(&self.p)))[id]; v != nil {
		return v
	} else {

		// not found int RCU, use RW-Lock map
		var ok bool
		self.mux.RLock()
		v, ok = (*(*map[int64]interface{})(unsafe.Pointer(self.bp)))[id]
		self.mux.RUnlock()

		// try store back to RCU
		if ok {
			self.SetByI64(id, v)
		}
		return v
	}
}

// GetByStr returns corresponding val of string id to
//
// WARN: this API can be only used on the IncRCU returned by `NewRCUStr()`
func (self *IncRCU) GetByStr(id string) interface{} {
	// try RCU first
	if v := (*(*map[string]interface{})(atomic.LoadPointer(&self.p)))[id]; v != nil {
		return v

	} else {

		// not found int RCU, use RW-Lock map
		var ok bool
		self.mux.RLock()
		v, ok = (*(*map[string]interface{})(unsafe.Pointer(self.bp)))[id]
		self.mux.RUnlock()

		// try store back to RCU
		if ok {
			self.SetByStr(id, v)
		}
		return v
	}
}

// SetByI64 stores int64 id and any val into the cache,
// and tells if the id already set
//
// WARN: this API can be only used on the IncRCU returned by `NewRCUI64()`
func (self *IncRCU) SetByI64(id int64, val interface{}) (exist bool) {
	for c := 0; c < self.maxCAS; c++ {
		// check if key exist
		m := (*map[int64]interface{})(atomic.LoadPointer(&self.p))
		if _, exist = (*m)[id]; exist {
			return
		}

		// copy old kvs
		n := make(map[int64]interface{}, len(*m))

		for k, v := range *m {
			n[k] = v
			if k == id {
				exist = true
				// duplicated key is not allowed
				// since two map may have the same key at the same time
				return
			}
		}

		// set new value
		n[id] = val

		// try update cache (CAS)
		if atomic.CompareAndSwapPointer(&self.p, unsafe.Pointer(m), unsafe.Pointer(&n)) {
			return
		}
	}

	// still not succeed, try use backup map using RW-Lock
	self.mux.Lock()
	defer self.mux.Unlock()
	bm := *(*map[int64]interface{})(unsafe.Pointer(self.bp))

	// check and set
	_, exist = bm[id]
	if !exist {
		bm[id] = val
	}
	return
}

// SetByStr stores string id and any val into the cache,
// and tells if the id already set
//
// WARN: this API can be only used on the IncRCU returned by `NewRCUStr()`
func (self *IncRCU) SetByStr(id string, val interface{}) (exist bool) {
	for c := 0; c < self.maxCAS; c++ {
		// check if key exist
		m := (*map[string]interface{})(atomic.LoadPointer(&self.p))
		if _, exist = (*m)[id]; exist {
			return
		}
		// copy old kvs
		n := make(map[string]interface{}, len(*m))
		for k, v := range *m {
			n[k] = v
			if k == id {
				exist = true
				// duplicated key is not allowed
				// since two map may have the same key at the same time
				return
			}
		}

		// set new value
		n[id] = val

		// try update cache (CAS)
		if atomic.CompareAndSwapPointer(&self.p, unsafe.Pointer(m), unsafe.Pointer(&n)) {
			return
		}
	}

	// still not succeed, try use backup map using RW-Lock
	self.mux.Lock()
	defer self.mux.Unlock()
	bm := *(*map[string]interface{})(unsafe.Pointer(self.bp))

	// check and set
	_, exist = bm[id]
	if !exist {
		bm[id] = val
	}
	return
}
