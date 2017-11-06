package repro_test

import (
	"crypto/rand"
	"os"
	"testing"

	"github.com/dgraph-io/badger"
)

func randBytes(n int) []byte {
	recv := make([]byte, n)
	in, _ := rand.Reader.Read(recv)
	return recv[:in]
}

var benchmarkData = []struct {
	key, value []byte
}{
	{randBytes(100), nil},
	{randBytes(1000), []byte("foo")},
	{[]byte("foo"), randBytes(1000)},
	{[]byte(""), randBytes(1000)},
	{nil, randBytes(1000000)},
	{randBytes(100000), nil},
	{randBytes(1000000), nil},
}

func BenchmarkIt(b *testing.B) {
	dir := "testbadger_db"
	os.RemoveAll(dir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll(dir)

	opts := new(badger.Options)
	*opts = badger.DefaultOptions
	opts.ValueLogFileSize = 1024 * 1024 * 1024
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(*opts)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		tx := db.NewTransaction(true)
		for _, kv := range benchmarkData {
			tx.Set(kv.key, kv.value)
		}
		if err := tx.Commit(nil); err != nil {
			b.Fatalf("#%d: batchSet err: %v", i, err)
		}
	}
}
