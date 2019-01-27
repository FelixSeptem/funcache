// Package funcache inspired by https://docs.python.org/3/library/functools.html powered LRU cache, LFU cache ARC cache and decorator pattern
package funcache

import (
	"reflect"

	"github.com/FelixSeptem/collections/arc"
	"github.com/FelixSeptem/collections/lfu"
	"github.com/FelixSeptem/collections/lru"
)

// wrapped function input and out put both are slice of interface{}
// all elements in it shall be comparable type due to https://golang.org/ref/spec#Comparison_operators generally not pointer,map,slice
type WrappedFun func([]interface{}) ([]interface{}, error)

// cache interface to get and set data use cache
type Cache interface {
	Set(key, value interface{}) bool
	Get(key interface{}) (interface{}, bool)
}

// input an array or a slice convert to a slice
func ToSlice(in interface{}) interface{} {
	a := reflect.ValueOf(in)
	if a.Kind() == reflect.Slice {
		return in
	}
	if a.Kind() != reflect.Array {
		panic("shall input a slice or an array")
	}
	t := reflect.SliceOf(a.Type().Elem())
	s := reflect.New(t).Elem()
	reflect.Copy(s, a)
	return s.Interface()
}

// input a slice or an array convert to an array
func ToArray(in interface{}) interface{} {
	s := reflect.ValueOf(in)
	if s.Kind() == reflect.Array {
		return in
	}
	if s.Kind() != reflect.Slice {
		panic("shall input a slice or an array")
	}
	t := reflect.ArrayOf(s.Len(), s.Type().Elem())
	a := reflect.New(t).Elem()
	reflect.Copy(a, s)
	return a.Interface()
}

// use given cache to cache expensive operate results
func CachedFun(wrapped WrappedFun, retries int, cache Cache) WrappedFun {
	return func(in []interface{}) ([]interface{}, error) {
		key := ToArray(in)
		if out, ok := cache.Get(key); ok {
			result := ToSlice(out)
			return result.([]interface{}), nil
		}
		var (
			res []interface{}
			err error
		)
		for i := 0; i < retries; i++ {
			res, err = wrapped(in)
			if err == nil {
				cache.Set(key, res)
				return res, nil
			}
		}
		return res, err
	}
}

// use LRU with CacheFun
func CacheFunWithLru(cacheSize int, wrapped WrappedFun, retries int) WrappedFun {
	return CachedFun(wrapped, retries, lru.NewLRUCache(cacheSize))
}

// use LFU with CacheFun
func CacheFunWithLfu(cacheSize int, wrapped WrappedFun, retries int) WrappedFun {
	return CachedFun(wrapped, retries, lfu.NewLFUCache(cacheSize))
}

// use ARC with CacheFun
func CacheFunWithArc(cacheSize int, wrapped WrappedFun, retries int) WrappedFun {
	return CachedFun(wrapped, retries, arc.NewARCCache(cacheSize))
}
