package application

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/havus/nekasa-quote-server/internal/domain/user"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/logger"
)

type UserServiceImpl struct {
	userRepository user.UserRepository
	logger         *logger.Logger
}

func NewUserService(userRepo user.UserRepository, logger *logger.Logger) user.UserService {
	return &UserServiceImpl{
		userRepository: userRepo,
		logger:         logger,
	}
}

func (s *UserServiceImpl) SignUp(ctx context.Context, u *user.User) error {
	existingUser, err := s.userRepository.FindByEmail(ctx, u.Email)
	if err == nil && existingUser != nil {
		s.logger.GeneralLog("error", "UserService.SignUp", "Email already in use", map[string]interface{}{"email": u.Email})
		return errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.EncryptedPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(hashedPassword)

	// Add additional values
	u.UID = generateUID()
	u.Provider = "email"
	u.SignUpIP = getCurrentIP(ctx)

	if err := s.userRepository.CreateUser(ctx, u); err != nil {
		return err
	}

	s.logger.GeneralLog("info", "UserService.SignUp", "User signed up successfully", map[string]interface{}{"email": u.Email})
	return nil
}

func (s *UserServiceImpl) SignIn(ctx context.Context, email, password string) (*user.User, error) {
	u, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		s.logger.GeneralLog("error", "UserService.SignIn", "User not found", map[string]interface{}{"email": email})
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)); err != nil {
		s.logger.GeneralLog("error", "UserService.SignIn", "Invalid password", map[string]interface{}{"email": email})
		return nil, errors.New("invalid credentials")
	}

	// Update sign-in information
	u.SignInCount++
	now := time.Now().UTC()
	u.LastSignInAt = u.CurrentSignInAt
	u.CurrentSignInAt = &now
	u.LastSignInIP = u.CurrentSignInIP
	u.CurrentSignInIP = getCurrentIP(ctx)

	if err := s.userRepository.UpdateUser(ctx, u); err != nil {
		return nil, err
	}

	s.logger.GeneralLog("info", "UserService.SignIn", "User signed in successfully", map[string]interface{}{"email": email})
	return u, nil
}
