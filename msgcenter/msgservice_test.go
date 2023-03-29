package msgcenter

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var m sync.Map

	m.Store("123", 13)
	m.Delete("321")
	t.Logf("%+v", m)

}
