package hs_test

import (
	"net/http"
	"sort"
	"testing"
)

var header http.Header

func init() {
	header = http.Header{
		"Content-Type":     {"1024"},
		"Content-Encoding": {"gzip"},
		"X-Trailer":        {"X-Expires-At", "Delta"},
		"Expires":          {"-1"},
		"a1":               {"a3", "z32"},
		"k1z":              {"a7", "z40"},
		"b3a":              {"xa9", "q1"},
		"b2a":              {"ba9", "u8"},
		"abcdefghi..z":     {"a", "z"},
		"Date":             {"Sun  5 Feb 2017 15:05:42 MST"},
		"foo":              {"bar"},
		"_abcdef":          {"bar"},
		"Aabcdef":          {"bar"},
		"FOOZ":             {"bar"},
		"f00Z":             {"Bar"},
		"User-Agent":       {""},
	}
}

func headerToKVs(h http.Header) (kvs []keyValues) {
	for key, values := range h {
		kvs = append(kvs, keyValues{key, values})
	}
	return kvs
}

type keyValues struct {
	key   string
	value []string
}

type headerSorter struct {
	kvs []keyValues
}

func (s *headerSorter) Len() int           { return len(s.kvs) }
func (s *headerSorter) Swap(i, j int)      { s.kvs[i], s.kvs[j] = s.kvs[j], s.kvs[i] }
func (s *headerSorter) Less(i, j int) bool { return s.kvs[i].key < s.kvs[j].key }

func sortHeaderSlice(kvs []keyValues) []keyValues {
	sort.Slice(kvs, func(i, j int) bool { return kvs[i].key < kvs[j].key })
	return kvs
}

func sortHeaderCustomSorter(kvs []keyValues) []keyValues {
	hs := &headerSorter{kvs: kvs}
	sort.Sort(hs)
	return kvs
}

func BenchmarkHeaderSortSlice(b *testing.B) {
	kvs := headerToKVs(header)
	var sink interface{}
	for i := 0; i < b.N; i++ {
		sink = sortHeaderSlice(kvs[:])
	}
	if sink != nil {
	}
	b.ReportAllocs()
}

func BenchmarkHeaderSortSorter(b *testing.B) {
	kvs := headerToKVs(header)
	var sink interface{}
	for i := 0; i < b.N; i++ {
		sink = sortHeaderCustomSorter(kvs[:])
	}
	if sink != nil {
	}
	b.ReportAllocs()
}
