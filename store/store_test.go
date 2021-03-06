package store

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/modb-io/modb"
	"github.com/modb-io/modb/store/badger"
	"github.com/modb-io/modb/store/bbolt"
)

func tmpName() string {
	return os.TempDir() + "/" + strconv.Itoa(rand.Int())
}

func TestStore(t *testing.T) {
	// setup each store

	// create ServerService
	bboltPath := tmpName()
	defer os.Remove(bboltPath)

	bboltStore, err := bbolt.Open(bboltPath)
	if err != nil {
		t.Fatal(err)
	}
	defer bboltStore.Close()

	// create ServerService
	badgerPath := tmpName()
	defer os.RemoveAll(badgerPath)

	badgerStore, err := badger.Open(badgerPath)
	if err != nil {
		t.Fatal(err)
	}
	defer badgerStore.Close()

	t.Run("Store Tests", func(t *testing.T) {
		t.Run("bbolt", func(t *testing.T) {
			ServerServiceTests(t, bboltStore)
		})
		t.Run("badger", func(t *testing.T) {
			ServerServiceTests(t, badgerStore)
		})
	})
}

func ServerServiceTests(t *testing.T, store modb.ServerService) {

	t.Run("Set", func(t *testing.T) {
		err := store.Set("account/chilts", "{}")
		if err != nil {
			t.Fatal(err)
			return
		}
	})

	t.Run("Inc", func(t *testing.T) {
		err := store.Inc("account/chilts", "logins")
		if err != nil {
			t.Fatal(err)
			return
		}
	})

}
