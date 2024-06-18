package lib

var charLimits = map[string]int{
	"username": 32,
	"password": 32,
	"email":    32,
}

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

func (user User) Validate() bool {
	if user.Username == "" || user.Password == "" {
		return false
	}

	if len(user.Username) > charLimits["username"] ||
		len(user.Password) > charLimits["password"] ||
		len(user.Email) > charLimits["email"] {
		return false
	}
	return true
}
