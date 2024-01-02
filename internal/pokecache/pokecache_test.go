package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Minute)
	if cache.cache == nil {
		t.Error("cache is nil")
	}

}

func TestCacheAddGet(t *testing.T) {
	cache := NewCache(time.Minute)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		entry, ok := cache.Get(c.inputKey)

		if !ok {
			t.Errorf("%s not found", c.inputKey)
			continue
		}
		if string(entry) != string(c.inputVal) {
			t.Errorf("%s does not match %s", string(entry), c.inputVal)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key1 := "key1"
	val1 := []byte("val1")
	cache.Add(key1, val1)

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key1)

	if ok {
		t.Error("Entry not deleted/reaped after interval")
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key1 := "key1"
	val1 := []byte("val1")
	cache.Add(key1, val1)

	time.Sleep(interval / 2)

	_, ok := cache.Get(key1)

	if !ok {
		t.Error("Entry should not be deleted/reaped before interval")
	}
}
