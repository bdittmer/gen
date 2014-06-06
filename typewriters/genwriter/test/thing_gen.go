// This file was generated by C:\Users\msherman\AppData\Local\Temp\go-build095643687\command-line-arguments\_obj\exe\setup.exe, using the gen typewriter

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

package main

import (
	"errors"
)

// Things is a slice of type Thing, for use with gen methods below. Use this type where you would use []Thing. (This is required because slices cannot be method receivers.)
type Things []Thing

// All verifies that all elements of Things return true for the passed func. See: http://clipperhouse.github.io/gen/#All
func (rcv Things) All(fn func(Thing) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Any verifies that one or more elements of Things return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv Things) Any(fn func(Thing) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}

// Count gives the number elements of Things that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv Things) Count(fn func(Thing) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// Distinct returns a new Things slice whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv Things) Distinct() (result Things) {
	appended := make(map[Thing]bool)
	for _, v := range rcv {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// DistinctBy returns a new Things slice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv Things) DistinctBy(equal func(Thing, Thing) bool) (result Things) {
	for _, v := range rcv {
		eq := func(_app Thing) bool {
			return equal(v, _app)
		}
		if !result.Any(eq) {
			result = append(result, v)
		}
	}
	return result
}

// Each iterates over Things and executes the passed func against each element. See: http://clipperhouse.github.io/gen/#Each
func (rcv Things) Each(fn func(Thing)) {
	for _, v := range rcv {
		fn(v)
	}
}

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv Things) First(fn func(Thing) bool) (result Thing, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no Things elements return true for passed func")
	return
}

// IsSortedBy reports whether an instance of Things is sorted, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv Things) IsSortedBy(less func(Thing, Thing) bool) bool {
	n := len(rcv)
	for i := n - 1; i > 0; i-- {
		if less(rcv[i], rcv[i-1]) {
			return false
		}
	}
	return true
}

// IsSortedDesc reports whether an instance of Things is sorted in descending order, using the pass func to define ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv Things) IsSortedByDesc(less func(Thing, Thing) bool) bool {
	greater := func(a, b Thing) bool {
		return a != b && !less(a, b)
	}
	return rcv.IsSortedBy(greater)
}

// MaxBy returns an element of Things containing the maximum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally maximal, the last such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MaxBy
func (rcv Things) MaxBy(less func(Thing, Thing) bool) (result Thing, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the MaxBy of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if rcv[i] != rcv[m] && !less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// MinBy returns an element of Things containing the minimum value, when compared to other elements using a passed func defining ‘less’. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#MinBy
func (rcv Things) MinBy(less func(Thing, Thing) bool) (result Thing, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Min of an empty slice")
		return
	}
	m := 0
	for i := 1; i < l; i++ {
		if less(rcv[i], rcv[m]) {
			m = i
		}
	}
	result = rcv[m]
	return
}

// Single returns exactly one element of Things that returns true for the passed func. Returns error if no or multiple elements return true. See: http://clipperhouse.github.io/gen/#Single
func (rcv Things) Single(fn func(Thing) bool) (result Thing, err error) {
	var candidate Thing
	found := false
	for _, v := range rcv {
		if fn(v) {
			if found {
				err = errors.New("multiple Things elements return true for passed func")
				return
			}
			candidate = v
			found = true
		}
	}
	if found {
		result = candidate
	} else {
		err = errors.New("no Things elements return true for passed func")
	}
	return
}

// SortBy returns a new ordered Things slice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv Things) SortBy(less func(Thing, Thing) bool) Things {
	result := make(Things, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortThings(result, less, 0, n, maxDepth)
	return result
}

// SortByDesc returns a new, descending-ordered Things slice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv Things) SortByDesc(less func(Thing, Thing) bool) Things {
	greater := func(a, b Thing) bool {
		return a != b && !less(a, b)
	}
	return rcv.SortBy(greater)
}

// Where returns a new Things slice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv Things) Where(fn func(Thing) bool) (result Things) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// AggregateOther iterates over Things, operating on each element while maintaining ‘state’. See: http://clipperhouse.github.io/gen/#Aggregate
func (rcv Things) AggregateOther(fn func(Other, Thing) Other) (result Other) {
	for _, v := range rcv {
		result = fn(result, v)
	}
	return
}

// AverageOther sums Other over all elements and divides by len(Things). See: http://clipperhouse.github.io/gen/#Average
func (rcv Things) AverageOther(fn func(Thing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine AverageOther of zero-length Things")
		return
	}
	for _, v := range rcv {
		result += fn(v)
	}
	result = result / Other(l)
	return
}

// GroupByOther groups elements into a map keyed by Other. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv Things) GroupByOther(fn func(Thing) Other) map[Other]Things {
	result := make(map[Other]Things)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// MaxOther selects the largest value of Other in Things. Returns error on Things with no elements. See: http://clipperhouse.github.io/gen/#MaxCustom
func (rcv Things) MaxOther(fn func(Thing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine MaxOther of zero-length Things")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f > result {
				result = f
			}
		}
	}
	return
}

// MinOther selects the least value of Other in Things. Returns error on Things with no elements. See: http://clipperhouse.github.io/gen/#MinCustom
func (rcv Things) MinOther(fn func(Thing) Other) (result Other, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine MinOther of zero-length Things")
		return
	}
	result = fn(rcv[0])
	if l > 1 {
		for _, v := range rcv[1:] {
			f := fn(v)
			if f < result {
				result = f
			}
		}
	}
	return
}

// SelectOther returns a slice of Other in Things, projected by passed func. See: http://clipperhouse.github.io/gen/#Select
func (rcv Things) SelectOther(fn func(Thing) Other) (result []Other) {
	for _, v := range rcv {
		result = append(result, fn(v))
	}
	return
}

// SumOther sums Other over elements in Things. See: http://clipperhouse.github.io/gen/#Sum
func (rcv Things) SumOther(fn func(Thing) Other) (result Other) {
	for _, v := range rcv {
		result += fn(v)
	}
	return
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapThings(rcv Things, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortThings(rcv Things, less func(Thing, Thing) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapThings(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownThings(rcv Things, less func(Thing, Thing) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapThings(rcv, first+root, first+child)
		root = child
	}
}

func heapSortThings(rcv Things, less func(Thing, Thing) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownThings(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapThings(rcv, first, first+i)
		siftDownThings(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeThings(rcv Things, less func(Thing, Thing) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapThings(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapThings(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapThings(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeThings(rcv Things, a, b, n int) {
	for i := 0; i < n; i++ {
		swapThings(rcv, a+i, b+i)
	}
}

func doPivotThings(rcv Things, less func(Thing, Thing) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeThings(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeThings(rcv, less, m, m-s, m+s)
		medianOfThreeThings(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeThings(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapThings(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapThings(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapThings(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeThings(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeThings(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortThings(rcv Things, less func(Thing, Thing) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortThings(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotThings(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortThings(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortThings(rcv, mhi, b)
		} else {
			quickSortThings(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortThings(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortThings(rcv, less, a, b)
	}
}
