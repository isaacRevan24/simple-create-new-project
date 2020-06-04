package users

// User type declaration
type User struct {
	UserID string
	// UserName string
	Email   string
	Manager bool
}

// GenerateMember Return a User
func GenerateMember(member map[string]string) User {
	// Initialize a empty User
	var user User

	// Assign the member map info to a user
	for userID, email := range member {
		user = User{
			UserID:  userID,
			Email:   email,
			Manager: true,
		}
	}
	return user
}
