// Copyright 2024 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package slicesbench

import (
	"slices"
	"testing"
)

type charmap [8]uint32

func (m *charmap) set(ch byte) {
	(*m)[ch>>5] |= 1 << uint(ch&31)
}

func (m *charmap) contains(ch byte) bool {
	return (*m)[ch>>5]&(1<<uint(ch&31)) != 0
}

const (
	testdata = "0123456789012345678901234567890123456789abc" // 40 digits, then lowercase
)

var (
	// Lowercase letters
	lowercase charmap

	// Uppercase letters
	uppercase charmap

	// Digits
	digits charmap
)

func init() {
	for ch := byte('a'); ch <= byte('z'); ch++ {
		lowercase.set(ch)
	}
	for ch := byte('A'); ch <= byte('Z'); ch++ {
		uppercase.set(ch)
	}
	for ch := byte('0'); ch <= byte('9'); ch++ {
		digits.set(ch)
	}
}

// indexBytesAny returns offset in buf where the byte is in charmap.
// It doesn't use slices.IndexFunc, but instead iterates over the buffer.
func indexBytesAny(buf []byte, cm charmap) int {
	for i, ch := range buf {
		if cm.contains(ch) {
			return i
		}
	}
	return len(buf)
}

// indexBytesAnyIndexFunc returns offset in buf where the byte is in charmap.
// It directly passes the charmap.contains method to slices.IndexFunc.
func indexBytesAnyIndexFuncDirect(buf []byte, cm charmap) int {
	r := slices.IndexFunc(buf, cm.contains)
	if r == -1 {
		return len(buf)
	}
	return r
}

// indexBytesAnyIndexFunc returns offset in buf where the byte is in charmap.
// It wraps the charmap.contains method in a closure and passes it to slices.IndexFunc.
func indexBytesAnyIndexFuncWrapped(buf []byte, cm charmap) int {
	r := slices.IndexFunc(buf, func(ch byte) bool { return cm.contains(ch) })
	if r == -1 {
		return len(buf)
	}
	return r
}

// TestIndexBytesAny tests indexBytesAny.
func TestIndexBytesAny(t *testing.T) {
	buf := []byte(testdata)
	if i := indexBytesAny(buf, lowercase); i != 40 {
		t.Errorf("indexBytesAny: expected 40, got %d", i)
	}
	if i := indexBytesAny(buf, uppercase); i != len(buf) {
		t.Errorf("indexBytesAny: expected %d, got %d", len(buf), i)
	}
}

// TestIndexBytesAnyIndexFuncDirect tests indexBytesAnyIndexFuncDirect.
func TestIndexBytesAnyIndexFuncDirect(t *testing.T) {
	buf := []byte(testdata)
	if i := indexBytesAnyIndexFuncDirect(buf, lowercase); i != 40 {
		t.Errorf("indexBytesAnyIndexFunc: expected 40, got %d", i)
	}
	if i := indexBytesAnyIndexFuncDirect(buf, uppercase); i != len(buf) {
		t.Errorf("indexBytesAnyIndexFunc: expected %d, got %d", len(buf), i)
	}
}

// TestIndexBytesAnyIndexFuncWrapped tests indexBytesAnyIndexFuncWrapped.
func TestIndexBytesAnyIndexFuncWrapped(t *testing.T) {
	buf := []byte(testdata)
	if i := indexBytesAnyIndexFuncWrapped(buf, lowercase); i != 40 {
		t.Errorf("indexBytesAnyIndexFunc: expected 40, got %d", i)
	}
	if i := indexBytesAnyIndexFuncWrapped(buf, uppercase); i != len(buf) {
		t.Errorf("indexBytesAnyIndexFunc: expected %d, got %d", len(buf), i)
	}
}

// TestIndexBytesSameResult tests that indexBytesAny and indexBytesAnyIndexFunc
// return the same result.
func TestIndexBytesSameResult(t *testing.T) {
	buf := []byte(testdata)

	if i, j := indexBytesAny(buf, lowercase), indexBytesAnyIndexFuncDirect(buf, lowercase); i != j {
		t.Errorf("indexBytesAny and indexBytesAnyIndexFuncDirect: expected %d, got %d", i, j)
	}
	if i, j := indexBytesAny(buf, lowercase), indexBytesAnyIndexFuncWrapped(buf, lowercase); i != j {
		t.Errorf("indexBytesAny and indexBytesAnyIndexFuncWrapped: expected %d, got %d", i, j)
	}

	if i, j := indexBytesAny(buf, uppercase), indexBytesAnyIndexFuncDirect(buf, uppercase); i != j {
		t.Errorf("indexBytesAny and indexBytesAnyIndexFuncDirect: expected %d, got %d", i, j)
	}
	if i, j := indexBytesAny(buf, uppercase), indexBytesAnyIndexFuncWrapped(buf, uppercase); i != j {
		t.Errorf("indexBytesAny and indexBytesAnyIndexFuncWrapped: expected %d, got %d", i, j)
	}
}

// BenchmarkIndexBytesAny benchmarks indexBytesAny.
func BenchmarkIndexBytesAny(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAny(buf, lowercase)
	}
}

// BenchmarkIndexBytesAnyDoesntContain benchmarks indexBytesAny with a buffer that doesn't contain any of the characters.
func BenchmarkIndexBytesAnyDoesntContain(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAny(buf, uppercase)
	}
}

// BenchmarkIndexBytesAnyImmediateHit benchmarks indexBytesAny with a buffer that contains the character at the beginning.
func BenchmarkIndexBytesAnyImmediateHit(b *testing.B) {
	buf := []byte("a" + testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAny(buf, lowercase)
	}
}

// BenchmarkIndexBytesAnyIndexFuncDirect benchmarks indexBytesAnyIndexFuncDirect.
func BenchmarkIndexBytesAnyIndexFuncDirect(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAnyIndexFuncDirect(buf, lowercase)
	}
}

// BenchmarkIndexBytesAnyIndexFuncDirectDoesntContain benchmarks indexBytesAnyIndexFuncDirect with a buffer that doesn't contain any of the characters.
func BenchmarkIndexBytesAnyIndexFuncDirectDoesntContain(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAnyIndexFuncDirect(buf, uppercase)
	}
}

// BenchmarkIndexBytesAnyIndexFuncDirectImmediateHit benchmarks indexBytesAnyIndexFuncDirect with a buffer that contains the character at the beginning.
func BenchmarkIndexBytesAnyIndexFuncDirectImmediateHit(b *testing.B) {
	buf := []byte("a" + testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAnyIndexFuncDirect(buf, lowercase)
	}
}

// BenchmarkIndexBytesAnyIndexFuncWrapped benchmarks indexBytesAnyIndexFuncWrapped.
func BenchmarkIndexBytesAnyIndexFuncWrapped(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAnyIndexFuncWrapped(buf, lowercase)
	}
}

// BenchmarkIndexBytesAnyIndexFuncWrappedDoesntContain benchmarks indexBytesAnyIndexFuncWrapped with a buffer that doesn't contain any of the characters.
func BenchmarkIndexBytesAnyIndexFuncWrappedDoesntContain(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAnyIndexFuncWrapped(buf, uppercase)
	}
}

// BenchmarkIndexBytesAnyIndexFuncWrappedImmediateHit benchmarks indexBytesAnyIndexFuncWrapped with a buffer that contains the character at the beginning.
func BenchmarkIndexBytesAnyIndexFuncWrappedImmediateHit(b *testing.B) {
	buf := []byte("a" + testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAnyIndexFuncWrapped(buf, lowercase)
	}
}
