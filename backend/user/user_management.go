package user

// UserTmp struct provides the user datamodel for temporarily storing in redis
type UserTmp struct {
	UserID       string
	UserName     string
	Password     string
	Email        string
	RegisteredAt int
}
