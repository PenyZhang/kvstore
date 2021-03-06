package etcdv3

import (
	"testing"
	"time"

	"github.com/YuleiXiao/kvstore"
	"github.com/YuleiXiao/kvstore/store"
	"github.com/YuleiXiao/kvstore/testutils"
	"github.com/stretchr/testify/assert"
)

var (
	client = "localhost:2379"
)

func makeEtcdClient(t *testing.T) store.Store {
	kv, err := New(
		[]string{client},
		&store.Config{
			ConnectionTimeout: 3 * time.Second,
			Username:          "test",
			Password:          "very-secure",
		},
	)

	if err != nil {
		t.Fatalf("cannot create store: %v", err)
	}

	return kv
}

func TestRegister(t *testing.T) {
	Register()

	kv, err := kvstore.NewStore(store.ETCDV3, []string{client}, nil)
	assert.NoError(t, err)
	assert.NotNil(t, kv)

	if _, ok := kv.(*Etcd); !ok {
		t.Fatal("Error registering and initializing etcd")
	}
}

func TestEtcdStore(t *testing.T) {
	kv := makeEtcdClient(t)
	lockKV := makeEtcdClient(t)
	ttlKV := makeEtcdClient(t)

	testutils.RunCleanup(t, kv)
	testutils.RunTestCommon(t, kv)
	testutils.RunTestAtomic(t, kv)
	testutils.RunTestWatch(t, kv)
	testutils.RunTestLockV3(t, kv)
	testutils.RunTestLockTTLV3(t, kv, lockKV)
	testutils.RunTestTTL(t, kv, ttlKV)
}
