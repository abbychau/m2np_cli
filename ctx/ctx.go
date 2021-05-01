package ctx

type User struct {
	ID       int
	Username string

	followings []User
	followers  []User

	numberOfArticles int
	lastLogin        int //timestamp
}

type M2npContext struct {
	User  User
	Token string
}
