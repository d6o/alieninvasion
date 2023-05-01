// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps defines various functions useful with maps of any type.
package maps

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
//
// Hard-forked "Keys" func from the "maps" package of Go 1.21
//
// This function has been hard-forked from the unreleased Go 1.21 version,
// where a new standard library "maps" package has been introduced.
// The "Keys" func can be found at the following URL:
// https://github.com/golang/go/blob/master/src/maps/maps.go#L8-L16
//
// The decision to hard-fork this function was made in order to take advantage
// of the new "Keys" func before the official release of Go 1.21.
//
// This hard-forked implementation will be removed and replaced with the
// official "maps" package once Go 1.21 is released and this project is
// updated to use the new version.
//
// To follow the progress of Go 1.21, visit the following link:
// https://github.com/golang/go/milestone/279
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
