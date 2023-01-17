package gomap

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertTest[K comparable, V comparable](t *testing.T, key []K, val []V) {
	mm := Make[K, V]()
	var defaultKey K
	var defaultVal V
	var wg sync.WaitGroup
	for i, v := range val {
		wg.Add(1)
		i, v := i, v
		go func() {
			defer wg.Done()
			if i > len(key)-1 {
				return
			}
			assert.Equal(t, mm.Get(defaultKey), defaultVal)
			mm.Put(key[i], v)
			assert.Equal(t, mm.Get(key[i]), v)
		}()
	}
	wg.Wait()
}

func TestSMap(t *testing.T) {
	t.Parallel()

	t.Run("safe for smap(int, int)", func(t *testing.T) {
		t.Parallel()

		key := []int{
			2233, 2234, 2235, 2236, 2237, 2233, 2234, 2235, 2236, 2237, 2233,
		}
		val := []uint{
			2233, 2234, 2235, 2236, 2237, 2233, 2234, 2235, 2236, 2237, 2233,
		}
		assertTest(t, key, val)
	})

	t.Run("safe for smap(int64, string)", func(t *testing.T) {
		t.Parallel()

		key := []uint64{
			2233, 2234, 2235, 2236, 2237, 2233, 2234, 2235, 2236, 2237, 2233,
		}
		val := []string{
			"2233", "2234", "2235", "2236", "2237", "2233", "2234", "2235", "2236", "2237", "2233",
		}
		assertTest(t, key, val)
	})

	t.Run("safe for smap(string, uint32)", func(t *testing.T) {
		t.Parallel()
		key := []string{
			"2233", "2234", "2235", "2236", "2237", "2233", "2234", "2235", "2236", "2237", "2233",
		}
		val := []int32{
			2233, 2234, 2235, 2236, 2237, 2233, 2234, 2235, 2236, 2237, 2233,
		}

		assertTest(t, key, val)
	})

	t.Run("safe for smap range", func(t *testing.T) {
		t.Parallel()
		key := []int32{
			2233, 2234, 2235, 2236, 2237, 2233, 2234, 2235, 2236, 2237, 2233,
		}
		val := []uint64{
			2233, 2234, 2235, 2236, 2237, 2233, 2234, 2235, 2236, 2237, 2233,
		}

		mm := Make[int32, uint64]()
		var wg sync.WaitGroup
		for i, v := range val {
			wg.Add(1)
			i, v := i, v
			go func() {
				defer wg.Done()
				mm.Put(key[i], v)
				assert.Equal(t, mm.Get(key[i]), v)
			}()
		}
		wg.Wait()
		mm.Range(func(key int32, val uint64) bool {
			assert.Equal(t, mm.Get(key), val)
			return true
		})
	})

	t.Run("safe for smap(int16, func())", func(t *testing.T) {
		t.Parallel()

		key := []int16{
			2233, 2234, 2235, 2236, 2237, 2233, 2234, 2235, 2236, 2237, 2233,
		}
		val := []func(){
			func() {}, func() { _ = 1 }, nil, func() {}, func() {}, nil, nil, func() { _ = 2 },
		}

		mm := Make[int16, func()]()
		var defaultKey int16
		var defaultVal func()
		var wg sync.WaitGroup
		for i, v := range val {
			wg.Add(1)
			i, v := i, v
			go func() {
				defer wg.Done()
				assert.IsType(t, mm.Get(defaultKey), defaultVal)
				mm.Put(key[i], v)
				assert.IsType(t, mm.Get(key[i]), v)
			}()
		}
		wg.Wait()
	})
}

func TestMap(t *testing.T) {
	t.Parallel()

	mm := MakeMap[int, string](10)
	t.Run("safe for map", func(t *testing.T) {
		key := []int{
			2233, 2234, 2235, 2236, 2237, 2233, 2234, 2235, 2236, 2237, 2233,
		}
		val := []string{
			"2233", "2234", "2235", "2236", "2237", "2233", "2234", "2235", "2236", "2237", "2233",
		}

		var wg sync.WaitGroup
		for i, v := range val {
			wg.Add(1)
			i, v := i, v
			go func() {
				defer wg.Done()
				assert.Equal(t, mm.Get(0), "")
				mm.Put(key[i], v)
				assert.Equal(t, mm.Get(key[i]), v)
			}()
		}
		wg.Wait()
	})

	t.Run("safe for map range", func(t *testing.T) {
		mm.Range(func(key int, val string) bool {
			assert.Equal(t, mm.Get(key), val)
			return true
		})
	})
}
