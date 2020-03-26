package domains

import (
	"math/rand"
	"time"
)

type User struct {
	ID       int64  `db:"id"`
	ComanyID int64  `db:"comany_id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
}

type Users []*User

type UserService struct {
	id int64
}

func NewUserService() *UserService {
	rand.Seed(time.Now().UnixNano())
	return &UserService{}
}

func (u *UserService) NextID() int64 {
	u.id++
	return u.id
}

func (u *UserService) RandomName() string {
	atoz := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 5)
	for i := range b {
		b[i] = atoz[rand.Intn(len(atoz))]
	}
	return string(b)
}

func (u *UserService) RandomAge() int {
	var (
		max = 70
		min = 20
	)
	return rand.Intn(max-min) + min
}
