package casbin_enforcers

// GetUsersForRoleInDomain gets the users that has a role inside a domain. Add by Gordon
func (e *SyncedEnforcer) GetUsersForRoleInDomain(name string, domain string) []string {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.GetUsersForRoleInDomain(name, domain)
}

// GetRolesForUserInDomain gets the roles that a user has inside a domain.
func (e *SyncedEnforcer) GetRolesForUserInDomain(name string, domain string) []string {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.GetRolesForUserInDomain(name, domain)
}

// GetPermissionsForUserInDomain gets permissions for a user or role inside a domain.
func (e *SyncedEnforcer) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.GetPermissionsForUserInDomain(user, domain)
}

// AddRoleForUserInDomain adds a role for a user inside a domain.
// Returns false if the user already has the role (aka not affected).
func (e *SyncedEnforcer) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.AddRoleForUserInDomain(user, role, domain)
}

// DeleteRoleForUserInDomain deletes a role for a user inside a domain.
// Returns false if the user does not have the role (aka not affected).
func (e *SyncedEnforcer) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.DeleteRoleForUserInDomain(user, role, domain)
}

func (e *SyncedEnforcer) GetRolesForUser(name string, domains ...string) ([]string, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.GetRolesForUser(name, domains...)
}

func (e *SyncedEnforcer) GetUsersForRole(name string, domains ...string) ([]string, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.GetUsersForRole(name, domains...)
}

func (e *SyncedEnforcer) HasRoleForUser(name string, role string, domain ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.HasRoleForUser(name, role, domain...)
}

func (e *SyncedEnforcer) AddRoleForUser(user string, role string, domain ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.AddRoleForUser(user, role, domain...)
}

func (e *SyncedEnforcer) DeleteRoleForUser(user string, role string, domain ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.DeleteRoleForUser(user, role, domain...)
}

func (e *SyncedEnforcer) DeleteRolesForUser(user string, domain ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.DeleteRolesForUser(user, domain...)
}

func (e *SyncedEnforcer) DeleteUser(user string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.DeleteUser(user)
}

func (e *SyncedEnforcer) DeleteRole(role string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.DeleteRole(role)
}

func (e *SyncedEnforcer) DeletePermission(permission ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.DeletePermission(permission...)
}

func (e *SyncedEnforcer) AddPermissionForUser(user string, permission ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.AddPermissionForUser(user, permission...)
}

func (e *SyncedEnforcer) DeletePermissionForUser(user string, permission ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.DeletePermissionForUser(user, permission...)
}

func (e *SyncedEnforcer) DeletePermissionsForUser(user string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.DeletePermissionsForUser(user)
}

func (e *SyncedEnforcer) GetPermissionsForUser(user string, domain ...string) [][]string {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.GetPermissionsForUser(user, domain...)
}

func (e *SyncedEnforcer) HasPermissionForUser(user string, permission ...string) bool {
	e.m.Lock()
	defer e.m.Unlock()
	return e.api.HasPermissionForUser(user, permission...)
}
