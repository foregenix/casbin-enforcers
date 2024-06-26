package casbin_enforcers

import (
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/effector"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/casbin/casbin/v2/rbac"
	"github.com/casbin/govaluate"
)

var _ FullEnforcer = &CachedEnforcer{}

// CachedEnforcer wraps Enforcer and provides decision cache.
type CachedEnforcer struct {
	base        BasicEnforcer
	api         APIEnforcer
	m           map[string]bool
	enableCache bool
	autoClear   bool
	locker      *sync.RWMutex
}

// NewCachedEnforcer creates a cached enforcer from an existing enforcer or via file or DB.
func NewCachedEnforcer(params ...interface{}) (*CachedEnforcer, error) {
	e := &CachedEnforcer{}
	if len(params) == 1 {
		if parent, ok := params[0].(FullEnforcer); ok {
			e.base = parent
			e.api = parent
		} else if parent, ok := params[0].(BasicEnforcer); ok {
			e.base = parent
			e.api = &DummyEnforcer{}
		}
	}
	if e.base == nil {
		ef, err := casbin.NewEnforcer(params...)
		if err != nil {
			return nil, err
		}
		e.base = ef
		e.api = ef
	}

	e.enableCache = true
	e.m = make(map[string]bool)
	e.locker = new(sync.RWMutex)
	return e, nil
}

// EnableCache determines whether to enable cache on Enforce(). When enableCache is enabled, cached result (true | false) will be returned for previous decisions.
func (e *CachedEnforcer) EnableCache(enableCache bool) {
	e.enableCache = enableCache
}

// EnableAutoCLear determines whether to automatically invalidate the cache when the policy is changed.
func (e *CachedEnforcer) EnableAutoClear(enableAuto bool) {
	e.autoClear = enableAuto
}

// Enforce decides whether a "subject" can access a "object" with the operation "action", input parameters are usually: (sub, obj, act).
// if rvals is not string , ignore the cache.
func (e *CachedEnforcer) Enforce(rvals ...interface{}) (bool, error) {
	if !e.enableCache {
		return e.base.Enforce(rvals...)
	}

	var key strings.Builder
	for _, rval := range rvals {
		if val, ok := rval.(string); ok {
			key.WriteString(val)
			key.WriteString("$$")
		} else {
			return e.base.Enforce(rvals...)
		}
	}

	if res, ok := e.getCachedResult(key.String()); ok {
		return res, nil
	}
	res, err := e.base.Enforce(rvals...)
	if err != nil {
		return false, err
	}

	e.setCachedResult(key.String(), res)
	return res, nil
}

func (e *CachedEnforcer) getCachedResult(key string) (res bool, ok bool) {
	e.locker.RLock()
	defer e.locker.RUnlock()
	res, ok = e.m[key]
	return res, ok
}

func (e *CachedEnforcer) setCachedResult(key string, res bool) {
	e.locker.Lock()
	defer e.locker.Unlock()
	e.m[key] = res
}

// InvalidateCache deletes all the existing cached decisions.
func (e *CachedEnforcer) InvalidateCache() {
	e.locker.Lock()
	defer e.locker.Unlock()
	if len(e.m) > 0 {
		e.m = make(map[string]bool)
	}
}

// GetParentEnforcer returns the parent enforcer wrapped by this instance.
func (e *CachedEnforcer) GetParentEnforcer() BasicEnforcer {
	return e.base
}

// InitWithFile initializes an enforcer with a model file and a policy file.
func (e *CachedEnforcer) InitWithFile(modelPath string, policyPath string) error {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.base.InitWithFile(modelPath, policyPath)
}

// InitWithAdapter initializes an enforcer with a database adapter.
func (e *CachedEnforcer) InitWithAdapter(modelPath string, adapter persist.Adapter) error {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.base.InitWithAdapter(modelPath, adapter)
}

// InitWithModelAndAdapter initializes an enforcer with a model and a database adapter.
func (e *CachedEnforcer) InitWithModelAndAdapter(m model.Model, adapter persist.Adapter) error {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.base.InitWithModelAndAdapter(m, adapter)
}

// LoadModel reloads the model from the model CONF file.
// Because the policy is attached to a model, so the policy is invalidated and needs to be reloaded by calling LoadPolicy().
func (e *CachedEnforcer) LoadModel() error {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.base.LoadModel()
}

// GetModel gets the current model.
func (e *CachedEnforcer) GetModel() model.Model {
	return e.base.GetModel()
}

// SetModel sets the current model.
func (e *CachedEnforcer) SetModel(m model.Model) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	e.base.SetModel(m)
}

// GetAdapter gets the current adapter.
func (e *CachedEnforcer) GetAdapter() persist.Adapter {
	return e.base.GetAdapter()
}

// SetAdapter sets the current adapter.
func (e *CachedEnforcer) SetAdapter(adapter persist.Adapter) {
	e.base.SetAdapter(adapter)
}

// SetWatcher sets the current watcher.
func (e *CachedEnforcer) SetWatcher(watcher persist.Watcher) error {
	return e.base.SetWatcher(watcher)
}

// GetRoleManager gets the current role manager.
func (e *CachedEnforcer) GetRoleManager() rbac.RoleManager {
	return e.base.GetRoleManager()
}

// SetRoleManager sets the current role manager.
func (e *CachedEnforcer) SetRoleManager(rm rbac.RoleManager) {
	e.base.SetRoleManager(rm)
	if e.autoClear {
		e.InvalidateCache()
	}
}

// SetEffector sets the current effector.
func (e *CachedEnforcer) SetEffector(eft effector.Effector) {
	e.base.SetEffector(eft)
	if e.autoClear {
		e.InvalidateCache()
	}
}

// ClearPolicy clears all policy.
func (e *CachedEnforcer) ClearPolicy() {
	e.base.ClearPolicy()
	if e.autoClear {
		e.InvalidateCache()
	}
}

// LoadPolicy reloads the policy from file/database.
func (e *CachedEnforcer) LoadPolicy() error {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.base.LoadPolicy()
}

// LoadFilteredPolicy reloads a filtered policy from file/database.
func (e *CachedEnforcer) LoadFilteredPolicy(filter interface{}) error {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.base.LoadFilteredPolicy(filter)
}

// IsFiltered returns true if the loaded policy has been filtered.
func (e *CachedEnforcer) IsFiltered() bool {
	return e.base.IsFiltered()
}

// SavePolicy saves the current policy (usually after changed with Casbin API) back to file/database.
func (e *CachedEnforcer) SavePolicy() error {
	return e.base.SavePolicy()
}

// EnableEnforce changes the enforcing state of Casbin, when Casbin is disabled, all access will be allowed by the Enforce() function.
func (e *CachedEnforcer) EnableEnforce(enable bool) {
	e.base.EnableEnforce(enable)
}

// EnableLog changes whether Casbin will log messages to the Logger.
func (e *CachedEnforcer) EnableLog(enable bool) {
	e.base.EnableLog(enable)
}

// EnableAutoSave controls whether to save a policy rule automatically to the adapter when it is added or removed.
func (e *CachedEnforcer) EnableAutoSave(autoSave bool) {
	e.base.EnableAutoSave(autoSave)
}

// EnableAutoBuildRoleLinks controls whether to rebuild the role inheritance relations when a role is added or deleted.
func (e *CachedEnforcer) EnableAutoBuildRoleLinks(autoBuildRoleLinks bool) {
	e.base.EnableAutoBuildRoleLinks(autoBuildRoleLinks)
}

// BuildRoleLinks manually rebuild the role inheritance relations.
func (e *CachedEnforcer) BuildRoleLinks() error {
	return e.base.BuildRoleLinks()
}

// EnforceWithMatcher use a custom matcher to decides whether a "subject" can access a "object" with the operation "action", input parameters are usually: (matcher, sub, obj, act), use model matcher by default when matcher is "".
func (e *CachedEnforcer) EnforceWithMatcher(matcher string, rvals ...interface{}) (bool, error) {
	return e.base.EnforceWithMatcher(matcher, rvals)
}

// GetAllSubjects gets the list of subjects that show up in the current policy.
func (e *CachedEnforcer) GetAllSubjects() ([]string, error) {
	return e.api.GetAllSubjects()
}

// GetAllNamedSubjects gets the list of subjects that show up in the current named policy.
func (e *CachedEnforcer) GetAllNamedSubjects(ptype string) ([]string, error) {
	return e.api.GetAllNamedSubjects(ptype)
}

// GetAllObjects gets the list of objects that show up in the current policy.
func (e *CachedEnforcer) GetAllObjects() ([]string, error) {
	return e.api.GetAllObjects()
}

// GetAllNamedObjects gets the list of objects that show up in the current named policy.
func (e *CachedEnforcer) GetAllNamedObjects(ptype string) ([]string, error) {
	return e.api.GetAllNamedObjects(ptype)
}

// GetAllActions gets the list of actions that show up in the current policy.
func (e *CachedEnforcer) GetAllActions() ([]string, error) {
	return e.api.GetAllActions()
}

// GetAllNamedActions gets the list of actions that show up in the current named policy.
func (e *CachedEnforcer) GetAllNamedActions(ptype string) ([]string, error) {
	return e.api.GetAllNamedActions(ptype)
}

// GetAllRoles gets the list of roles that show up in the current policy.
func (e *CachedEnforcer) GetAllRoles() ([]string, error) {
	return e.api.GetAllRoles()
}

// GetAllNamedRoles gets the list of roles that show up in the current named policy.
func (e *CachedEnforcer) GetAllNamedRoles(ptype string) ([]string, error) {
	return e.api.GetAllNamedRoles(ptype)
}

// GetPolicy gets all the authorization rules in the policy.
func (e *CachedEnforcer) GetPolicy() ([][]string, error) {
	return e.api.GetPolicy()
}

// GetFilteredPolicy gets all the authorization rules in the policy, field filters can be specified.
func (e *CachedEnforcer) GetFilteredPolicy(fieldIndex int, fieldValues ...string) ([][]string, error) {
	return e.api.GetFilteredPolicy(fieldIndex, fieldValues...)
}

// GetNamedPolicy gets all the authorization rules in the named policy.
func (e *CachedEnforcer) GetNamedPolicy(ptype string) ([][]string, error) {
	return e.api.GetNamedPolicy(ptype)
}

// GetFilteredNamedPolicy gets all the authorization rules in the named policy, field filters can be specified.
func (e *CachedEnforcer) GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	return e.api.GetFilteredNamedPolicy(ptype, fieldIndex, fieldValues...)
}

// GetGroupingPolicy gets all the role inheritance rules in the policy.
func (e *CachedEnforcer) GetGroupingPolicy() ([][]string, error) {
	return e.api.GetGroupingPolicy()
}

// GetFilteredGroupingPolicy gets all the role inheritance rules in the policy, field filters can be specified.
func (e *CachedEnforcer) GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) ([][]string, error) {
	return e.api.GetFilteredGroupingPolicy(fieldIndex, fieldValues...)
}

// GetNamedGroupingPolicy gets all the role inheritance rules in the policy.
func (e *CachedEnforcer) GetNamedGroupingPolicy(ptype string) ([][]string, error) {
	return e.api.GetNamedGroupingPolicy(ptype)
}

// GetFilteredNamedGroupingPolicy gets all the role inheritance rules in the policy, field filters can be specified.
func (e *CachedEnforcer) GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	return e.api.GetFilteredNamedGroupingPolicy(ptype, fieldIndex, fieldValues...)
}

// HasPolicy determines whether an authorization rule exists.
func (e *CachedEnforcer) HasPolicy(params ...interface{}) (bool, error) {
	return e.api.HasPolicy(params...)
}

// HasNamedPolicy determines whether a named authorization rule exists.
func (e *CachedEnforcer) HasNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	return e.api.HasNamedPolicy(ptype, params...)
}

// AddPolicy adds an authorization rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (e *CachedEnforcer) AddPolicy(params ...interface{}) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.AddPolicy(params...)
}

// AddNamedPolicy adds an authorization rule to the current named policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (e *CachedEnforcer) AddNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.AddNamedPolicy(ptype, params...)
}

// RemovePolicy removes an authorization rule from the current policy.
func (e *CachedEnforcer) RemovePolicy(params ...interface{}) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.RemovePolicy(params...)
}

// RemoveFilteredPolicy removes an authorization rule from the current policy, field filters can be specified.
func (e *CachedEnforcer) RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.RemoveFilteredPolicy(fieldIndex, fieldValues...)
}

// RemoveNamedPolicy removes an authorization rule from the current named policy.
func (e *CachedEnforcer) RemoveNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.RemoveNamedPolicy(ptype, params...)
}

// RemoveFilteredNamedPolicy removes an authorization rule from the current named policy, field filters can be specified.
func (e *CachedEnforcer) RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.RemoveFilteredNamedPolicy(ptype, fieldIndex, fieldValues...)
}

// HasGroupingPolicy determines whether a role inheritance rule exists.
func (e *CachedEnforcer) HasGroupingPolicy(params ...interface{}) (bool, error) {
	return e.api.HasGroupingPolicy(params...)
}

// HasNamedGroupingPolicy determines whether a named role inheritance rule exists.
func (e *CachedEnforcer) HasNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	return e.api.HasNamedGroupingPolicy(ptype, params...)
}

// AddGroupingPolicy adds a role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (e *CachedEnforcer) AddGroupingPolicy(params ...interface{}) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.AddGroupingPolicy(params...)
}

// AddNamedGroupingPolicy adds a named role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (e *CachedEnforcer) AddNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.AddNamedGroupingPolicy(ptype, params...)
}

// RemoveGroupingPolicy removes a role inheritance rule from the current policy.
func (e *CachedEnforcer) RemoveGroupingPolicy(params ...interface{}) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.RemoveGroupingPolicy(params...)
}

// RemoveFilteredGroupingPolicy removes a role inheritance rule from the current policy, field filters can be specified.
func (e *CachedEnforcer) RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.RemoveFilteredGroupingPolicy(fieldIndex, fieldValues...)
}

// RemoveNamedGroupingPolicy removes a role inheritance rule from the current named policy.
func (e *CachedEnforcer) RemoveNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.RemoveNamedGroupingPolicy(ptype, params...)
}

// RemoveFilteredNamedGroupingPolicy removes a role inheritance rule from the current named policy, field filters can be specified.
func (e *CachedEnforcer) RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.RemoveFilteredNamedGroupingPolicy(ptype, fieldIndex, fieldValues...)
}

// AddFunction adds a customized function.
func (e *CachedEnforcer) AddFunction(name string, function govaluate.ExpressionFunction) {
	e.api.AddFunction(name, function)
}

// GetImplicitPermissionsForUser gets implicit permissions for a user or role.
// Compared to GetPermissionsForUser(), this function retrieves permissions for inherited roles.
// For example:
// p, admin, data1, read
// p, alice, data2, read
// g, alice, admin
//
// GetPermissionsForUser("alice") can only get: [["alice", "data2", "read"]].
// But GetImplicitPermissionsForUser("alice") will get: [["admin", "data1", "read"], ["alice", "data2", "read"]].
func (e *CachedEnforcer) GetImplicitPermissionsForUser(user string, domain ...string) ([][]string, error) {
	return e.api.GetImplicitPermissionsForUser(user, domain...)
}

// GetImplicitRolesForUser gets implicit roles that a user has.
// Compared to GetRolesForUser(), this function retrieves indirect roles besides direct roles.
// For example:
// g, alice, role:admin
// g, role:admin, role:user
//
// GetRolesForUser("alice") can only get: ["role:admin"].
// But GetImplicitRolesForUser("alice") will get: ["role:admin", "role:user"].
func (e *CachedEnforcer) GetImplicitRolesForUser(user string, domain ...string) ([]string, error) {
	return e.api.GetImplicitRolesForUser(user, domain...)
}

// GetImplicitUsersForPermission gets implicit users for a permission.
// For example:
// p, admin, data1, read
// p, bob, data1, read
// g, alice, admin
//
// GetImplicitUsersForPermission("data1", "read") will get: ["alice", "bob"].
// Note: only users will be returned, roles (2nd arg in "g") will be excluded.
func (e *CachedEnforcer) GetImplicitUsersForPermission(permission ...string) ([]string, error) {
	return e.api.GetImplicitUsersForPermission(permission...)
}
