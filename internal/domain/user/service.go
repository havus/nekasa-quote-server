package user

import "context"

type UserService interface {
	SignUp(ctx context.Context, user *User) error
	SignIn(ctx context.Context, email, password string) (*User, error)
}
