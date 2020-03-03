package casbin_enforcers

import "testing"

func testEnforceCache(t *testing.T, e *CachedEnforcer, sub string, obj interface{}, act string, res bool) {
	t.Helper()
	if myRes, _ := e.Enforce(sub, obj, act); myRes != res {
		t.Errorf("%s, %v, %s: %t, supposed to be %t", sub, obj, act, myRes, res)
	}
}

func TestCache(t *testing.T) {
	e, _ := NewCachedEnforcer("examples/basic_model.conf", "examples/basic_policy.csv")
	e.autoClear = false
	// The cache is enabled by default for NewCachedEnforcer.

	testEnforceCache(t, e, "alice", "data1", "read", true)
	testEnforceCache(t, e, "alice", "data1", "write", false)
	testEnforceCache(t, e, "alice", "data2", "read", false)
	testEnforceCache(t, e, "alice", "data2", "write", false)

	// The cache is enabled, so even if we remove a policy rule, the decision
	// for ("alice", "data1", "read") will still be true, as it uses the cached result.
	e.RemovePolicy("alice", "data1", "read")

	testEnforceCache(t, e, "alice", "data1", "read", true)
	testEnforceCache(t, e, "alice", "data1", "write", false)
	testEnforceCache(t, e, "alice", "data2", "read", false)
	testEnforceCache(t, e, "alice", "data2", "write", false)

	// Now we invalidate the cache, then all first-coming Enforce() has to be evaluated in real-time.
	// The decision for ("alice", "data1", "read") will be false now.
	e.InvalidateCache()

	testEnforceCache(t, e, "alice", "data1", "read", false)
	testEnforceCache(t, e, "alice", "data1", "write", false)
	testEnforceCache(t, e, "alice", "data2", "read", false)
	testEnforceCache(t, e, "alice", "data2", "write", false)
}
