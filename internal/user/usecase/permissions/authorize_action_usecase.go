package permissions

type AuthorizeActionUseCase struct{}

func (AuthorizeActionUseCase) Can(role, action string) bool {
	if role == "admin" {
		return true
	}
	// very naive policy
	switch action {
	case "user:read":
		return role == "user" || role == "client" || role == "support"
	case "user:write":
		return role == "support"
	default:
		return false
	}
}
