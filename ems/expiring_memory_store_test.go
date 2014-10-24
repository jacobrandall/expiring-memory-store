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

import (
	// "sort"
	"strconv"
	"testing"
	"time"
)

func TestMapCreation(t *testing.T) {
	m := NewExpiringMemoryStore()
	if m == nil {
		t.Error("map is null.")
	}

	if m.CountAll() != 0 {
		t.Error("new map should be empty.")
	}
}

func TestWrite(t *testing.T) {
	m := NewExpiringMemoryStore()
	m.Write("animal1", "elephant")
	m.Write("animal2", "monkey")

	if m.CountAll() != 2 {
		t.Error("map should contain exactly two elements.")
	}
}

func TestWriteWithExpiration(t *testing.T) {
	m := NewExpiringMemoryStore()
	m.WriteWithExpiration("animal1", "elephant", 1)

	if m.CountAll() != 1 {
		t.Error("map should contain exactly one element.")
	}

	time.Sleep(time.Duration(2) * time.Second)

	val, err := m.Read("animal1")
	if err != expiredElementError {
		t.Error("Expired error should be returned.")
	}

	if val != "" {
		t.Error("Expected empty string to be returned when element is expired.")
	}

	if m.CountActive() != 0 {
		t.Error("map should be empty.")
	}

	if m.CountAll() != 1 {
		t.Error("map should contain one inactive entry.")
	}
}

func TestRead(t *testing.T) {
	m := NewExpiringMemoryStore()

	// Get a missing element.
	val, err := m.Read("Money")

	if err != elementNotFoundError {
		t.Error("Expected elementNotFoundError when item is missing from map.")
	}

	if val != "" {
		t.Error("Missing values should return as an empty string.")
	}

	m.Write("animal1", "elephant")

	// Retrieve inserted element.

	tmp, err := m.Read("animal1")

	if err == elementNotFoundError {
		t.Error("Error should be nil for item stored within the map.")
	}

	if &tmp == nil {
		t.Error("expecting an element, not null.")
	}

	if tmp != "elephant" {
		t.Error("item was modified.")
	}
}

func TestExists(t *testing.T) {
	m := NewExpiringMemoryStore()

	// Get a missing element.
	if m.Exists("Money") == true {
		t.Error("element shouldn't exist")
	}

	m.Write("foo", "bar")

	if m.Exists("foo") == false {
		t.Error("element exists, expecting Exists to return True.")
	}
}

func TestRemove(t *testing.T) {
	m := NewExpiringMemoryStore()

	m.Write("foo", "bar")

	m.Remove("foo")

	if m.CountAll() != 0 {
		t.Error("Expecting count to be zero once item was removed.")
	}

	temp, err := m.Read("monkey")

	if err != elementNotFoundError {
		t.Error("Expecting elementNotFoundError for missing items.")
	}

	if temp != "" {
		t.Error("Expecting item to be nil after its removal.")
	}

	// Remove a none existing element.
	m.Remove("noone")
}

func TestCount(t *testing.T) {
	m := NewExpiringMemoryStore()
	for i := 0; i < 100; i++ {
		m.Write(strconv.Itoa(i), strconv.Itoa(i))
	}

	if m.CountAll() != 100 {
		t.Error("Expecting 100 element within map.")
	}
}

func TestClear(t *testing.T) {
	m := NewExpiringMemoryStore()

	m.Clear()
	if m.CountAll() != 0 {
		t.Error("Expecting an empty map")
	}

	m.Write("foo", "bar")

	m.Clear()
	if m.CountAll() != 0 {
		t.Error("Expecting an empty map")
	}
}
