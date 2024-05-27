package casbin_enforcers

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/effector"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/casbin/casbin/v2/rbac"
	"github.com/casbin/govaluate"
)

type (
	// BasicEnforcer is the interface that describes the minimal set of functions required for an enforcer
	//
	// An object implements BasicEnforcer to enable it to be used as a wrapper around Enforcer.
	BasicEnforcer interface {
		InitWithFile(modelPath string, policyPath string) error
		InitWithAdapter(modelPath string, adapter persist.Adapter) error
		InitWithModelAndAdapter(m model.Model, adapter persist.Adapter) error
		LoadModel() error
		GetModel() model.Model
		SetModel(m model.Model)
		GetAdapter() persist.Adapter
		SetAdapter(adapter persist.Adapter)
		SetWatcher(watcher persist.Watcher) error
		GetRoleManager() rbac.RoleManager
		SetRoleManager(rm rbac.RoleManager)
		SetEffector(eft effector.Effector)
		ClearPolicy()
		LoadPolicy() error
		LoadFilteredPolicy(filter interface{}) error
		IsFiltered() bool
		SavePolicy() error
		EnableEnforce(enable bool)
		EnableLog(enable bool)
		EnableAutoSave(autoSave bool)
		EnableAutoBuildRoleLinks(autoBuildRoleLinks bool)
		BuildRoleLinks() error
		Enforce(rvals ...interface{}) (bool, error)
		EnforceWithMatcher(matcher string, rvals ...interface{}) (bool, error)
	}

	ChainedEnforcer interface {
		BasicEnforcer
		GetParentEnforcer() ChainedEnforcer
		GetRootEnforcer() *casbin.Enforcer
	}

	// APIEnforcer is the interface which provides the management and RBAC API functions
	//
	// Enforcer wrappers must implement this interface in order to expose the RBAC and
	// management APIs from lower level wrappers or the root Enforcer.
	APIEnforcer interface {
		GetRolesForUser(name string, domains ...string) ([]string, error)
		GetUsersForRole(name string, domains ...string) ([]string, error)
		HasRoleForUser(name string, role string, domain ...string) (bool, error)
		AddRoleForUser(user string, role string, domain ...string) (bool, error)
		DeleteRoleForUser(user string, role string, domain ...string) (bool, error)
		DeleteRolesForUser(user string, domain ...string) (bool, error)
		DeleteUser(user string) (bool, error)
		DeleteRole(role string) (bool, error)
		DeletePermission(permission ...string) (bool, error)
		AddPermissionForUser(user string, permission ...string) (bool, error)
		DeletePermissionForUser(user string, permission ...string) (bool, error)
		DeletePermissionsForUser(user string) (bool, error)
		GetPermissionsForUser(user string, domain ...string) ([][]string, error)
		HasPermissionForUser(user string, permission ...string) (bool, error)
		GetImplicitRolesForUser(name string, domain ...string) ([]string, error)
		GetImplicitPermissionsForUser(user string, domain ...string) ([][]string, error)
		GetImplicitUsersForPermission(permission ...string) ([]string, error)
		GetUsersForRoleInDomain(name string, domain string) []string
		GetRolesForUserInDomain(name string, domain string) []string
		GetPermissionsForUserInDomain(user string, domain string) [][]string
		AddRoleForUserInDomain(user string, role string, domain string) (bool, error)
		DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error)
		GetAllSubjects() ([]string, error)
		GetAllNamedSubjects(ptype string) ([]string, error)
		GetAllObjects() ([]string, error)
		GetAllNamedObjects(ptype string) ([]string, error)
		GetAllActions() ([]string, error)
		GetAllNamedActions(ptype string) ([]string, error)
		GetAllRoles() ([]string, error)
		GetAllNamedRoles(ptype string) ([]string, error)
		GetPolicy() ([][]string, error)
		GetFilteredPolicy(fieldIndex int, fieldValues ...string) ([][]string, error)
		GetNamedPolicy(ptype string) ([][]string, error)
		GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error)
		GetGroupingPolicy() ([][]string, error)
		GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) ([][]string, error)
		GetNamedGroupingPolicy(ptype string) ([][]string, error)
		GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error)
		HasPolicy(params ...interface{}) (bool, error)
		HasNamedPolicy(ptype string, params ...interface{}) (bool, error)
		AddPolicy(params ...interface{}) (bool, error)
		AddNamedPolicy(ptype string, params ...interface{}) (bool, error)
		RemovePolicy(params ...interface{}) (bool, error)
		RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error)
		RemoveNamedPolicy(ptype string, params ...interface{}) (bool, error)
		RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error)
		HasGroupingPolicy(params ...interface{}) (bool, error)
		HasNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error)
		AddGroupingPolicy(params ...interface{}) (bool, error)
		AddNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error)
		RemoveGroupingPolicy(params ...interface{}) (bool, error)
		RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) (bool, error)
		RemoveNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error)
		RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error)
		AddFunction(name string, function govaluate.ExpressionFunction)
	}

	// FullEnforcer is the interface which describes the full featured Enforcer interface
	FullEnforcer interface {
		BasicEnforcer
		APIEnforcer
	}
)

// GetRootEnforcer locates and returns the root instance of Enforcer from a arbitrarily wrapped instance
func GetRootEnforcer(e BasicEnforcer) *casbin.Enforcer {
	var ce ChainedEnforcer
	var ok bool
	for {
		if ce, ok = e.(ChainedEnforcer); !ok {
			return e.(*casbin.Enforcer)
		}
		if ne := ce.GetParentEnforcer(); ne != nil {
			e = ne
		} else {
			return ce.GetRootEnforcer()
		}
	}
}
