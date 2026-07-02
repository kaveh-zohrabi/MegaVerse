package service

import (
	"megaverse.dev/services/user-service/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

type UpdateProfileRequest struct {
	Bio       string `json:"bio"`
	Website   string `json:"website"`
	Location  string `json:"location"`
	AvatarURL string `json:"avatar_url"`
	CoverURL  string `json:"cover_url"`
}

type ProfileResponse struct {
	UserID         string `json:"user_id"`
	Bio            string `json:"bio"`
	Website        string `json:"website"`
	Location       string `json:"location"`
	AvatarURL      string `json:"avatar_url"`
	CoverURL       string `json:"cover_url"`
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
	PostsCount     int    `json:"posts_count"`
}

type FollowResponse struct {
	IsFollowing bool `json:"is_following"`
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetProfile(userID string) (*ProfileResponse, error) {
	profile, err := s.userRepo.GetProfile(userID)
	if err != nil {
		// Create profile if not exists
		if err := s.userRepo.CreateProfile(userID); err != nil {
			return nil, err
		}
		profile, err = s.userRepo.GetProfile(userID)
		if err != nil {
			return nil, err
		}
	}

	return &ProfileResponse{
		UserID:         profile.UserID,
		Bio:            profile.Bio,
		Website:        profile.Website,
		Location:       profile.Location,
		AvatarURL:      profile.AvatarURL,
		CoverURL:       profile.CoverURL,
		FollowersCount: profile.FollowersCount,
		FollowingCount: profile.FollowingCount,
		PostsCount:     profile.PostsCount,
	}, nil
}

func (s *UserService) UpdateProfile(userID string, req *UpdateProfileRequest) error {
	return s.userRepo.UpdateProfile(userID, req.Bio, req.Website, req.Location, req.AvatarURL, req.CoverURL)
}

func (s *UserService) Follow(followerID, followingID string) error {
	return s.userRepo.Follow(followerID, followingID)
}

func (s *UserService) Unfollow(followerID, followingID string) error {
	return s.userRepo.Unfollow(followerID, followingID)
}

func (s *UserService) IsFollowing(followerID, followingID string) (bool, error) {
	return s.userRepo.IsFollowing(followerID, followingID)
}

func (s *UserService) GetFollowers(userID string, limit, offset int) ([]string, error) {
	return s.userRepo.GetFollowers(userID, limit, offset)
}

func (s *UserService) GetFollowing(userID string, limit, offset int) ([]string, error) {
	return s.userRepo.GetFollowing(userID, limit, offset)
}
