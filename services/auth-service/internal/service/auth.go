package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"megaverse.dev/services/auth-service/internal/repository"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailExists        = errors.New("email already registered")
	ErrUserNotFound       = errors.New("user not found")
)

type AuthService struct {
	userRepo       *repository.UserRepository
	jwtSecret      []byte
	jwtExpiration  time.Duration
	refreshExpiry  time.Duration
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string, jwtExp, refreshExp time.Duration) *AuthService {
	return &AuthService{
		userRepo:      userRepo,
		jwtSecret:     []byte(jwtSecret),
		jwtExpiration: jwtExp,
		refreshExpiry: refreshExp,
	}
}

func (s *AuthService) Register(email, name, password string) (*repository.User, error) {
	exists, err := s.userRepo.EmailExists(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &repository.User{
		ID:           uuid.New().String(),
		Email:        email,
		Name:         name,
		PasswordHash: string(hash),
		Role:         "user",
		Active:       true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (*TokenPair, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return s.generateTokenPair(user)
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (s *AuthService) generateTokenPair(user *repository.User) (*TokenPair, error) {
	now := time.Now()
	expiresAt := now.Add(s.jwtExpiration)

	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "megaverse",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}

	refreshClaims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshExpiry)),
		IssuedAt:  jwt.NewNumericDate(now),
		Issuer:    "megaverse",
		Subject:   user.ID,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshStr,
		ExpiresAt:    expiresAt.Unix(),
	}, nil
}
