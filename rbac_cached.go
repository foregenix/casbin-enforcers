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
