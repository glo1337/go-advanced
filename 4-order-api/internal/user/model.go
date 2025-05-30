package user

import (
	"math/rand"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone     string `json:"phone" gorm:"uniqueIndex"`
	SessionId string `json:"session_id"`
	Token     string `json:"token"`
}

func NewUser(phone string) *User {
	user := &User{
		Phone: phone,
	}
	user.GenerateSessionId()
	return user
}

func (user *User) GenerateSessionId() {
	user.SessionId = RandStringRunes(10)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
