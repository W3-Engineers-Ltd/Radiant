// Copyright 2014 beego Author. All Rights Reserved.
//

package utils

import (
	"github.com/W3-Engineers-Ltd/Radiant/core/utils"
)

type reducetype func(interface{}) interface{}

type filtertype func(interface{}) bool

// InSlice checks given string in string slice or not.
func InSlice(v string, sl []string) bool {
	return utils.InSlice(v, sl)
}

// InSliceIface checks given interface in interface slice.
func InSliceIface(v interface{}, sl []interface{}) bool {
	return utils.InSliceIface(v, sl)
}

// SliceRandList generate an int slice from min to max.
func SliceRandList(min, max int) []int {
	return utils.SliceRandList(min, max)
}

// SliceMerge merges interface slices to one slice.
func SliceMerge(slice1, slice2 []interface{}) (c []interface{}) {
	return utils.SliceMerge(slice1, slice2)
}

// SliceReduce generates a new slice after parsing every value by reduce function
func SliceReduce(slice []interface{}, a reducetype) (dslice []interface{}) {
	return utils.SliceReduce(slice, func(i interface{}) interface{} {
		return a(i)
	})
}

// SliceRand returns random one from slice.
func SliceRand(a []interface{}) (b interface{}) {
	return utils.SliceRand(a)
}

// SliceSum sums all values in int64 slice.
func SliceSum(intslice []int64) (sum int64) {
	return utils.SliceSum(intslice)
}

// SliceFilter generates a new slice after filter function.
func SliceFilter(slice []interface{}, a filtertype) (ftslice []interface{}) {
	return utils.SliceFilter(slice, func(i interface{}) bool {
		return a(i)
	})
}

// SliceDiff returns diff slice of slice1 - slice2.
func SliceDiff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	return utils.SliceDiff(slice1, slice2)
}

// SliceIntersect returns slice that are present in all the slice1 and slice2.
func SliceIntersect(slice1, slice2 []interface{}) (diffslice []interface{}) {
	return utils.SliceIntersect(slice1, slice2)
}

// SliceChunk separates one slice to some sized slice.
func SliceChunk(slice []interface{}, size int) (chunkslice [][]interface{}) {
	return utils.SliceChunk(slice, size)
}

// SliceRange generates a new slice from begin to end with step duration of int64 number.
func SliceRange(start, end, step int64) (intslice []int64) {
	return utils.SliceRange(start, end, step)
}

// SlicePad prepends size number of val into slice.
func SlicePad(slice []interface{}, size int, val interface{}) []interface{} {
	return utils.SlicePad(slice, size, val)
}

// SliceUnique cleans repeated values in slice.
func SliceUnique(slice []interface{}) (uniqueslice []interface{}) {
	return utils.SliceUnique(slice)
}

// SliceShuffle shuffles a slice.
func SliceShuffle(slice []interface{}) []interface{} {
	return utils.SliceShuffle(slice)
}
