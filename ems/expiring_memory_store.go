// Copyright 2014 Brighcove Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package ems provides an expiring memory store
package ems

import "errors"
import "github.com/streamrail/concurrent-map"

type ExpiringMemoryStore struct {
	elementMap *cmap.ConcurrentMap
}

var expiredElementError = errors.New("Element expired")
var elementNotFoundError = errors.New("Element not found")

func NewExpiringMemoryStore() *ExpiringMemoryStore {
	return &ExpiringMemoryStore{elementMap: cmap.NewConcurrentMap()}
}

func (store *ExpiringMemoryStore) Read(name string) (string, error) {
	map_entry, ok := store.elementMap.Get(name)

	// Checks if item exists
	if ok == true {
		// Map stores items as interface{}, hence we'll have to cast.
		element := map_entry.(*Element)
		if element.IsExpired() {
			return "", expiredElementError
		} else {
			return element.value, nil
		}
	}

	return "", elementNotFoundError
}

func (store *ExpiringMemoryStore) Write(name string, value string) {
	element := NewElement(value)
	store.elementMap.Add(name, element)
}

func (store *ExpiringMemoryStore) WriteWithExpiration(name string, value string, expiresIn int64) {
	element := NewElementWithExpiration(value, expiresIn)
	store.elementMap.Add(name, element)
}

func (store *ExpiringMemoryStore) Remove(name string) {
	store.elementMap.Remove(name)
}

func (store *ExpiringMemoryStore) Exists(name string) bool {
	return store.elementMap.Has(name)
}

func (store *ExpiringMemoryStore) Clear() {
	store.elementMap.Clear()
}

func (store *ExpiringMemoryStore) CountActive() int {
	activeCount := 0
	for t := range store.elementMap.Iter() {
		element := t.Val.(*Element)
		if !element.IsExpired() {
			activeCount = activeCount + 1
		}
	}
	return activeCount
}

func (store *ExpiringMemoryStore) CountAll() int {
	return store.elementMap.Count()
}
