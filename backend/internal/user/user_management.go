package user

// TmpUser struct provides the user datamodel for temporarily storing in redis
type TmpUser struct {
	UserID       string
	UserName     string
	Password     string
	Email        string
	RegisteredAt int
}
