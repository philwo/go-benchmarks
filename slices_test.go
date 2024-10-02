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
	testdata = "0123456789012345678901234567890123456789abc"
)

var (
	// Lowercase letters
	lowercase charmap

	// Digits
	digits charmap
)

func init() {
	for ch := byte('a'); ch <= byte('z'); ch++ {
		lowercase.set(ch)
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

// skipBytesAny returns offset in buf where the byte is not in charmap.
// It is the same as indexBytesAny, but with the condition inverted.
func skipBytesAny(buf []byte, cm charmap) int {
	for i, ch := range buf {
		if !cm.contains(ch) {
			return i
		}
	}
	return len(buf)
}

// indexBytesAnyIndexFunc returns offset in buf where the byte is in charmap.
// It directly passes the charmap.contains method to slices.IndexFunc.
func indexBytesAnyIndexFuncDirect(buf []byte, cm charmap) int {
	return slices.IndexFunc(buf, cm.contains)
}

// indexBytesAnyIndexFunc returns offset in buf where the byte is in charmap.
// It wraps the charmap.contains method in a closure and passes it to slices.IndexFunc.
func indexBytesAnyIndexFuncWrapped(buf []byte, cm charmap) int {
	return slices.IndexFunc(buf, func(ch byte) bool { return cm.contains(ch) })
}

// skipBytesAnyIndexFunc returns offset in buf where the byte is not in charmap.
// It wraps the charmap.contains method in a closure and passes it to slices.IndexFunc.
func skipBytesAnyIndexFunc(buf []byte, cm charmap) int {
	return slices.IndexFunc(buf, func(ch byte) bool { return !cm.contains(ch) })
}

// TestIndexBytesAny tests indexBytesAny.
func TestIndexBytesAny(t *testing.T) {
	buf := []byte(testdata)
	if i := indexBytesAny(buf, lowercase); i != 40 {
		t.Errorf("indexBytesAny: expected 40, got %d", i)
	}
}

// TestSkipBytesAny tests skipBytesAny.
func TestSkipBytesAny(t *testing.T) {
	buf := []byte(testdata)
	if i := skipBytesAny(buf, digits); i != 40 {
		t.Errorf("skipBytesAny: expected 40, got %d", i)
	}
}

// TestIndexBytesAnyIndexFuncDirect tests indexBytesAnyIndexFuncDirect.
func TestIndexBytesAnyIndexFuncDirect(t *testing.T) {
	buf := []byte(testdata)
	if i := indexBytesAnyIndexFuncDirect(buf, lowercase); i != 40 {
		t.Errorf("indexBytesAnyIndexFunc: expected 40, got %d", i)
	}
}

// TestIndexBytesAnyIndexFuncWrapped tests indexBytesAnyIndexFuncWrapped.
func TestIndexBytesAnyIndexFuncWrapped(t *testing.T) {
	buf := []byte(testdata)
	if i := indexBytesAnyIndexFuncWrapped(buf, lowercase); i != 40 {
		t.Errorf("indexBytesAnyIndexFunc: expected 40, got %d", i)
	}
}

// TestSkipBytesAnyIndexFunc tests skipBytesAnyIndexFunc.
func TestSkipBytesAnyIndexFunc(t *testing.T) {
	buf := []byte(testdata)
	if i := skipBytesAnyIndexFunc(buf, digits); i != 40 {
		t.Errorf("skipBytesAnyIndexFunc: expected 40, got %d", i)
	}
}

// BenchmarkIndexBytesAny benchmarks indexBytesAny.
func BenchmarkIndexBytesAny(b *testing.B) {
	buf := []byte(testdata)
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

// BenchmarkIndexBytesAnyIndexFuncWrapped benchmarks indexBytesAnyIndexFuncWrapped.
func BenchmarkIndexBytesAnyIndexFuncWrapped(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		indexBytesAnyIndexFuncWrapped(buf, lowercase)
	}
}

// BenchmarkSkipBytesAny benchmarks skipBytesAny.
func BenchmarkSkipBytesAny(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		skipBytesAny(buf, digits)
	}
}

// BenchmarkSkipBytesAnyIndexFunc benchmarks skipBytesAnyIndexFunc.
func BenchmarkSkipBytesAnyIndexFunc(b *testing.B) {
	buf := []byte(testdata)
	for i := 0; i < b.N; i++ {
		skipBytesAnyIndexFunc(buf, digits)
	}
}
