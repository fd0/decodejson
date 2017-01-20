package decoder

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

const testFileName = "index.json"

func TestIndexDecodeJSON(t *testing.T) {
	f, err := os.Open(testFileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	dec := json.NewDecoder(f)
	var idx Index
	err = dec.Decode(&idx)
	if err != nil {
		panic(err)
	}

	blobs := 0
	for _, p := range idx.Packs {
		blobs += len(p.Blobs)
	}

	t.Logf("index has %v packs and %v blobs", len(idx.Packs), blobs)
}

func BenchmarkIndexDecodeJSON(b *testing.B) {
	f, err := os.Open(testFileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	dec := json.NewDecoder(f)

	var idx Index

	fi, err := f.Stat()
	if err != nil {
		b.Fatal(err)
	}

	b.SetBytes(fi.Size())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = f.Seek(0, 0)
		if err != nil {
			b.Fatal(err)
		}

		err = dec.Decode(&idx)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIndexUnmarshalJSON(b *testing.B) {
	f, err := os.Open(testFileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var idx Index

	fi, err := f.Stat()
	if err != nil {
		b.Fatal(err)
	}

	b.SetBytes(fi.Size())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(buf, &idx)
		if err != nil {
			b.Fatal(err)
		}
	}
}
