package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                  int            `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName           string         `gorm:"column:first_name"`
	MiddleName          string         `gorm:"column:middle_name"`
	LastName            string         `gorm:"column:last_name"`
	Email               string         `gorm:"column:email;unique"`
	UnconfirmedEmail    string         `gorm:"column:unconfirmed_email"`
	Status              int            `gorm:"column:status"`
	EmailVerified       bool           `gorm:"column:email_verified"`
	PhoneVerified       bool           `gorm:"column:phone_verified"`
	Username            string         `gorm:"column:username;unique"`
	Phone               string         `gorm:"column:phone"`
	EncryptedPassword   string         `gorm:"column:encrypted_password"`
	PasswordUpdatedAt   *time.Time     `gorm:"column:password_updated_at"`
	ResetPasswordSentAt *time.Time     `gorm:"column:reset_password_sent_at"`
	RememberMeCreatedAt *time.Time     `gorm:"column:remember_me_created_at"`
	ConfirmationToken   string         `gorm:"column:confirmation_token"`
	ConfirmedAt         *time.Time     `gorm:"column:confirmed_at"`
	ConfirmationSentAt  *time.Time     `gorm:"column:confirmation_sent_at"`
	Provider            string         `gorm:"column:provider"`
	UID                 string         `gorm:"column:uid"`
	SignInCount         int            `gorm:"column:sign_in_count"`
	CurrentSignInAt     *time.Time     `gorm:"column:current_sign_in_at"`
	LastSignInAt        *time.Time     `gorm:"column:last_sign_in_at"`
	CurrentSignInIP     string         `gorm:"column:current_sign_in_ip"`
	LastSignInIP        string         `gorm:"column:last_sign_in_ip"`
	SignUpIP            string         `gorm:"column:sign_up_ip"`
	CreatedAt           time.Time      `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}
