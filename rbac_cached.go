package casbin_enforcers

// GetUsersForRoleInDomain gets the users that has a role inside a domain. Add by Gordon
func (e *CachedEnforcer) GetUsersForRoleInDomain(name string, domain string) []string {
	return e.api.GetUsersForRoleInDomain(name, domain)
}

// GetRolesForUserInDomain gets the roles that a user has inside a domain.
func (e *CachedEnforcer) GetRolesForUserInDomain(name string, domain string) []string {
	return e.api.GetRolesForUserInDomain(name, domain)
}

// GetPermissionsForUserInDomain gets permissions for a user or role inside a domain.
func (e *CachedEnforcer) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	return e.api.GetPermissionsForUserInDomain(user, domain)
}

// AddRoleForUserInDomain adds a role for a user inside a domain.
// Returns false if the user already has the role (aka not affected).
func (e *CachedEnforcer) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return e.api.AddRoleForUserInDomain(user, role, domain)
}

// DeleteRoleForUserInDomain deletes a role for a user inside a domain.
// Returns false if the user does not have the role (aka not affected).
func (e *CachedEnforcer) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return e.api.DeleteRoleForUserInDomain(user, role, domain)
}

func (e *CachedEnforcer) GetRolesForUser(name string, domains ...string) ([]string, error) {
	return e.api.GetRolesForUser(name, domains...)
}

func (e *CachedEnforcer) GetUsersForRole(name string, domains ...string) ([]string, error) {
	return e.api.GetUsersForRole(name, domains...)
}

func (e *CachedEnforcer) HasRoleForUser(name string, role string, domain ...string) (bool, error) {
	return e.api.HasRoleForUser(name, role, domain...)
}

func (e *CachedEnforcer) AddRoleForUser(user string, role string, domain ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.AddRoleForUser(user, role, domain...)
}

func (e *CachedEnforcer) DeleteRoleForUser(user string, role string, domain ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.DeleteRoleForUser(user, role, domain...)
}

func (e *CachedEnforcer) DeleteRolesForUser(user string, domain ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.DeleteRolesForUser(user, domain...)
}

func (e *CachedEnforcer) DeleteUser(user string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.DeleteUser(user)
}

func (e *CachedEnforcer) DeleteRole(role string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.DeleteRole(role)
}

func (e *CachedEnforcer) DeletePermission(permission ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.DeletePermission(permission...)
}

func (e *CachedEnforcer) AddPermissionForUser(user string, permission ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.AddPermissionForUser(user, permission...)
}

func (e *CachedEnforcer) DeletePermissionForUser(user string, permission ...string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.DeletePermissionForUser(user, permission...)
}

func (e *CachedEnforcer) DeletePermissionsForUser(user string) (bool, error) {
	if e.autoClear {
		defer e.InvalidateCache()
	}
	return e.api.DeletePermissionsForUser(user)
}

func (e *CachedEnforcer) GetPermissionsForUser(user string, domain ...string) [][]string {
	return e.api.GetPermissionsForUser(user, domain...)
}

func (e *CachedEnforcer) HasPermissionForUser(user string, permission ...string) bool {
	return e.api.HasPermissionForUser(user, permission...)
}
