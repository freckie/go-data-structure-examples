package hashtable

const (
	bucketCap       uint    = 8
	bucketCapFactor uint    = 2
	loadThreshold   float32 = 0.75
)

type hashFuncType[K comparable] func(K) int

type HashTable[K comparable, V any] struct {
	buckets  []*HashTableEntry[K, V]
	hash     hashFuncType[K]
	len      uint
	capacity uint
}

func NewHashTable[K comparable, V any]() *HashTable[K, V] {
	buckets := make([]*HashTableEntry[K, V], bucketCap)
	table := &HashTable[K, V]{
		buckets:  buckets,
		hash:     nil,
		len:      0,
		capacity: bucketCap,
	}
	return table
}

func (t *HashTable[K, V]) SetHashFunc(f hashFuncType[K]) *HashTable[K, V] {
	t.hash = f
	return t
}

func (t HashTable[K, V]) IsEmpty() bool {
	return t.len == 0
}

func (t HashTable[K, V]) Len() uint {
	return t.len
}

func (t HashTable[K, V]) Get(key K) (V, bool) {
	h := t.hash(key)
	idx := h % int(t.capacity)

	it := t.buckets[idx]
	if it == nil {
		return *new(V), false
	}

	for it != nil {
		if it.key == key {
			break
		}
		it = it.next
	}
	return it.value, true
}

func (t *HashTable[K, V]) Set(key K, value V) {
	if (float32(t.len) / float32(t.capacity)) >= loadThreshold {
		t.resize(t.capacity * bucketCapFactor)
	}

	t.set(key, value)
}

func (t *HashTable[K, V]) set(key K, value V) {
	h := t.hash(key)
	idx := h % int(t.capacity)

	entry := &HashTableEntry[K, V]{
		key:   key,
		value: value,
		next:  nil,
	}

	if bucket := t.buckets[idx]; bucket != nil {
		var prev *HashTableEntry[K, V]
		for it := bucket; it != nil; it = it.next {
			prev = it
			if it.key == key {
				it.value = value
				return
			}
		}
		prev.next = entry
	} else {
		t.buckets[idx] = entry
	}
	t.len++
}

func (t *HashTable[K, V]) resize(newCap uint) {
	oldBuckets := t.buckets
	newBuckets := make([]*HashTableEntry[K, V], newCap)
	t.capacity = newCap

	for i := 0; i < len(oldBuckets); i++ {
		for oldBuckets[i] != nil {
			entry := oldBuckets[i]
			newHash := t.hash(entry.key)
			newIdx := newHash % int(newCap)

			entry.next = newBuckets[newIdx]
			newBuckets[newIdx] = entry
			oldBuckets[i] = entry.next
		}
	}
	t.buckets = newBuckets
}

type HashTableEntry[K comparable, V any] struct {
	key   K
	value V
	next  *HashTableEntry[K, V]
}
