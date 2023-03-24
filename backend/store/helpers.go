package store

import "strings"

// This file contains some helper functions that are used by the store.
// These are functions that are used to build queries or parse parameters.
// Some of these functions will probably be replaced by the standard library once go 1.19 is released.

// === query helpers ===

// whereQuery builds a where query based on the given conditions
func whereQuery(conditions []string) string {
	if len(conditions) > 0 {
		return "WHERE " + strings.Join(conditions, " AND ")
	}
	return ""
}

// orderByQuery builds an order by query based on the given conditions
func orderByQuery(conditions []string) string {
	if len(conditions) > 0 {
		return "ORDER BY " + strings.Join(conditions, ", ")
	}
	return ""
}

// inQuery repeats an x amount of times a placeholder
func inQuery(times int) string {
	if times > 0 {
		return strings.Repeat("?, ", times-1) + "?"
	}
	return ""
}

// === generic helper functions ===

// mapSlice takes a slice and calls `mapTo` on each element, returning a new slice with the results
func mapSlice[T any, M any](in []T, mapTo func(T) M) []M {
	slice := make([]M, len(in))
	for idx, val := range in {
		slice[idx] = mapTo(val)
	}
	return slice
}

// mapToIdx takes a slice and calls `mapTo` on each element, and returns the results mapped by index
func mapToIdx[T any, K comparable](in []T, mapTo func(T) K) map[K]int {
	dict := make(map[K]int, len(in))
	for idx, val := range in {
		dict[mapTo(val)] = idx
	}
	return dict
}

// mapSliceToAny maps a slice of type `T` to any
func mapSliceToAny[T any](in []T) []any {
	slice := make([]any, len(in))
	for idx, val := range in {
		slice[idx] = val
	}
	return slice
}

// filter calls `filter` on each element of `in` and returns a new slice with the elements that returned true
func filter[T any](in []T, filter func(T) bool) []T {
	slice := make([]T, 0, len(in))
	for _, val := range in {
		if filter(val) {
			slice = append(slice, val)
		}
	}
	return slice
}

// difference returns the elements in `a` that aren't in `b`
func difference[T comparable](a, b []T) (diff []T) {
	set := make(map[T]any, len(b))
	for _, val := range b {
		set[val] = struct{}{} // dummy value that doesn't allocate space
	}

	for _, val := range a {
		if _, ok := set[val]; !ok {
			diff = append(diff, val)
		}
	}
	return diff
}
