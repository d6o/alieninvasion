// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maps

import (
	"golang.org/x/exp/slices"
	"sort"
	"testing"
)

// reference: https://github.com/golang/go/blob/master/src/maps/maps_test.go#LL15C1-L32C2
var m1 = map[int]int{1: 2, 2: 4, 4: 8, 8: 16}
var m2 = map[int]string{1: "2", 2: "4", 4: "8", 8: "16"}

func TestKeys(t *testing.T) {
	want := []int{1, 2, 4, 8}

	got1 := Keys(m1)
	sort.Ints(got1)
	if !slices.Equal(got1, want) {
		t.Errorf("Keys(%v) = %v, want %v", m1, got1, want)
	}

	got2 := Keys(m2)
	sort.Ints(got2)
	if !slices.Equal(got2, want) {
		t.Errorf("Keys(%v) = %v, want %v", m2, got2, want)
	}
}
