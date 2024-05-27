package casbin_enforcers

import (
	"errors"

	"github.com/casbin/govaluate"
)

// DummyEnforcer is a dummy implementation of APIEnforcer which simply returns either
// nil or ErrorUnsupported on all function calls.
// DummyEnforcer is used to provide functionality to enforcer wrappers which do not
// implement APIEnforcer.
type DummyEnforcer struct {
}

// ErrorUnsupported is returned to indicate that a specific function is not supported by an enforcer wrapper.
var ErrorUnsupported = errors.New("unsupported function")

// GetRolesForUser returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetRolesForUser(name string, domains ...string) ([]string, error) {
	return nil, ErrorUnsupported
}

// GetUsersForRole returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetUsersForRole(name string, domains ...string) ([]string, error) {
	return nil, ErrorUnsupported
}

// HasRoleForUser returns false and ErrorUnsupported
func (e *DummyEnforcer) HasRoleForUser(name string, role string, domain ...string) (bool, error) {
	return false, ErrorUnsupported
}

// AddRoleForUser returns false and ErrorUnsupported.
func (e *DummyEnforcer) AddRoleForUser(user string, role string, domain ...string) (bool, error) {
	return false, ErrorUnsupported
}

// DeleteRoleForUser returns false and ErrorUnsupported.
func (e *DummyEnforcer) DeleteRoleForUser(user string, role string, domain ...string) (bool, error) {
	return false, ErrorUnsupported
}

// DeleteRolesForUser returns false and ErrorUnsupported.
func (e *DummyEnforcer) DeleteRolesForUser(user string, domain ...string) (bool, error) {
	return false, ErrorUnsupported
}

// DeleteUser returns false and ErrorUnsupported.
func (e *DummyEnforcer) DeleteUser(user string) (bool, error) {
	return false, ErrorUnsupported
}

// DeleteRole returns false and ErrorUnsupported.
func (e *DummyEnforcer) DeleteRole(role string) (bool, error) {
	return false, ErrorUnsupported
}

// DeletePermission returns false and ErrorUnsupported.
func (e *DummyEnforcer) DeletePermission(permission ...string) (bool, error) {
	return false, ErrorUnsupported
}

// AddPermissionForUser returns false and ErrorUnsupported.
func (e *DummyEnforcer) AddPermissionForUser(user string, permission ...string) (bool, error) {
	return false, ErrorUnsupported
}

// DeletePermissionForUser returns false and ErrorUnsupported.
func (e *DummyEnforcer) DeletePermissionForUser(user string, permission ...string) (bool, error) {
	return false, ErrorUnsupported
}

// DeletePermissionsForUser returns false and ErrorUnsupported.
func (e *DummyEnforcer) DeletePermissionsForUser(user string) (bool, error) {
	return false, ErrorUnsupported
}

// GetPermissionsForUser returns nil.
func (e *DummyEnforcer) GetPermissionsForUser(user string, domain ...string) ([][]string, error) {
	return nil, ErrorUnsupported
}

// HasPermissionForUser returns false and ErrorUnsupported.
func (e *DummyEnforcer) HasPermissionForUser(user string, permission ...string) (bool, error) {
	return false, ErrorUnsupported
}

// GetImplicitRolesForUser returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetImplicitRolesForUser(name string, domain ...string) ([]string, error) {
	return nil, ErrorUnsupported
}

// GetImplicitPermissionsForUser returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetImplicitPermissionsForUser(user string, domain ...string) ([][]string, error) {
	return nil, ErrorUnsupported
}

// GetImplicitUsersForPermission returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetImplicitUsersForPermission(permission ...string) ([]string, error) {
	return nil, ErrorUnsupported
}

// GetUsersForRoleInDomain returns nil.
func (e *DummyEnforcer) GetUsersForRoleInDomain(name string, domain string) []string {
	return nil
}

// GetRolesForUserInDomain returns nil.
func (e *DummyEnforcer) GetRolesForUserInDomain(name string, domain string) []string {
	return nil
}

// GetPermissionsForUserInDomain returns nil.
func (e *DummyEnforcer) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	return nil
}

// AddRoleForUserInDomain returns false and ErrorUnsupported.
func (e *DummyEnforcer) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return false, ErrorUnsupported
}

// DeleteRoleForUserInDomain returns false and ErrorUnsupported.
func (e *DummyEnforcer) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return false, ErrorUnsupported
}

// GetAllSubjects returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetAllSubjects() ([]string, error) {
	return nil, ErrorUnsupported
}

// GetAllNamedSubjects returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetAllNamedSubjects(ptype string) ([]string, error) {
	return nil, ErrorUnsupported
}

// GetAllObjects returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetAllObjects() ([]string, error) {
	return nil, ErrorUnsupported
}

// GetAllNamedObjects returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetAllNamedObjects(ptype string) ([]string, error) {
	return nil, ErrorUnsupported
}

// GetAllActions returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetAllActions() ([]string, error) {
	return nil, ErrorUnsupported
}

// GetAllNamedActions returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetAllNamedActions(ptype string) ([]string, error) {
	return nil, ErrorUnsupported
}

// GetAllRoles returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetAllRoles() ([]string, error) {
	return nil, ErrorUnsupported
}

// GetAllNamedRoles returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetAllNamedRoles(ptype string) ([]string, error) {
	return nil, ErrorUnsupported
}

// GetPolicy returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetPolicy() ([][]string, error) {
	return nil, ErrorUnsupported
}

// GetFilteredPolicy returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetFilteredPolicy(fieldIndex int, fieldValues ...string) ([][]string, error) {
	return nil, ErrorUnsupported
}

// GetNamedPolicy returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetNamedPolicy(ptype string) ([][]string, error) {
	return nil, ErrorUnsupported
}

// GetFilteredNamedPolicy returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	return nil, ErrorUnsupported
}

// GetGroupingPolicy returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetGroupingPolicy() ([][]string, error) {
	return nil, ErrorUnsupported
}

// GetFilteredGroupingPolicy returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) ([][]string, error) {
	return nil, ErrorUnsupported
}

// GetNamedGroupingPolicy returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetNamedGroupingPolicy(ptype string) ([][]string, error) {
	return nil, ErrorUnsupported
}

// GetFilteredNamedGroupingPolicy returns nil and ErrorUnsupported.
func (e *DummyEnforcer) GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	return nil, ErrorUnsupported
}

// HasPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) HasPolicy(params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// HasNamedPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) HasNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// AddPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) AddPolicy(params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// AddNamedPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) AddNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// RemovePolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) RemovePolicy(params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// RemoveFilteredPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	return false, ErrorUnsupported
}

// RemoveNamedPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) RemoveNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// RemoveFilteredNamedPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	return false, ErrorUnsupported
}

// HasGroupingPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) HasGroupingPolicy(params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// HasNamedGroupingPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) HasNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// AddGroupingPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) AddGroupingPolicy(params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// AddNamedGroupingPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) AddNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// RemoveGroupingPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) RemoveGroupingPolicy(params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// RemoveFilteredGroupingPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	return false, ErrorUnsupported
}

// RemoveNamedGroupingPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) RemoveNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	return false, ErrorUnsupported
}

// RemoveFilteredNamedGroupingPolicy returns false and ErrorUnsupported.
func (e *DummyEnforcer) RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	return false, ErrorUnsupported
}

// AddFunction does nothing.
func (e *DummyEnforcer) AddFunction(name string, function govaluate.ExpressionFunction) {
	return
}
