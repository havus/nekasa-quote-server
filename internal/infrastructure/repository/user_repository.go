package repository

import (
	"context"

	"github.com/havus/nekasa-quote-server/internal/domain/user"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) user.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *GormUserRepository) CreateUser(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}
