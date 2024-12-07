package user

type InsertUserRequest struct {
	Email          string
	Username       string
	SaltedPassword string
}
