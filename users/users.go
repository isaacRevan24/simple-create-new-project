package users

// User type declaration
type User struct {
	UserID   string
	UserName string
	Email    string
	Manager  bool
}

// GenerateMember Return a User
//! In this case only return the manager
func GenerateMember(member *map[string][2]string) User {
	// Initialize a empty User
	var user User
	var memberBase map[string][2]string
	memberBase = *member

	// Assign the member map info to a user
	for userID, email := range memberBase {
		user = User{
			UserID:   userID,
			UserName: email[0],
			Email:    email[1],
			Manager:  true,
		}
	}
	return user
}
