package hashtable

import "testing"

func TestHashTable(t *testing.T) {
	h := NewHashTable[int, string]().SetHashFunc(func(v int) int {
		return int(v)
	})

	h.Set(2, "no.2")
	h.Set(7, "no.7")
	h.Set(5, "no.5")

	dat, ok := h.Get(2)
	if !ok {
		t.Error("errorerror")
	}
	t.Logf("value(%s) len(%d)", dat, h.Len())

	h.Set(2, "changed no.2")
	dat2, ok := h.Get(2)
	if !ok {
		t.Error("errorerror")
	}
	t.Logf("value(%s) len(%d)", dat2, h.Len())

	t.Error("temp")
}
