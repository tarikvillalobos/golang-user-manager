package repository

type RoleRepository interface {
	// In a real system, roles might be stored in DB.
	IsValid(role string) bool
}

type InMemoryRoleRepo struct{}

func (InMemoryRoleRepo) IsValid(role string) bool {
	switch role {
	case "admin", "client", "support", "user":
		return true
	default:
		return false
	}
}
