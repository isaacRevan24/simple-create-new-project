package users

// User type declaration
type User struct {
	Email    string
	UserID   string
	UserName string
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
	for email, info := range memberBase {
		user = User{
			Email:    email,
			UserName: info[0],
			UserID:   info[1],
			Manager:  true,
		}
	}
	return user
}
