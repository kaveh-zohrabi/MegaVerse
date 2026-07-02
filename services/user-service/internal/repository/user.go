package repository

import (
	"database/sql"
	"time"
)

type UserProfile struct {
	UserID         string
	Bio            string
	Website        string
	Location       string
	AvatarURL      string
	CoverURL       string
	FollowersCount int
	FollowingCount int
	PostsCount     int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserFollow struct {
	FollowerID  string
	FollowingID string
	CreatedAt   time.Time
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetProfile(userID string) (*UserProfile, error) {
	profile := &UserProfile{}
	err := r.db.QueryRow(
		`SELECT user_id, COALESCE(bio,''), COALESCE(website,''), COALESCE(location,''), 
		 COALESCE(avatar_url,''), COALESCE(cover_url,''), followers_count, following_count, 
		 posts_count, created_at, updated_at 
		 FROM user_profiles WHERE user_id = ?`,
		userID,
	).Scan(&profile.UserID, &profile.Bio, &profile.Website, &profile.Location,
		&profile.AvatarURL, &profile.CoverURL, &profile.FollowersCount, &profile.FollowingCount,
		&profile.PostsCount, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *UserRepository) CreateProfile(userID string) error {
	_, err := r.db.Exec(
		"INSERT INTO user_profiles (user_id) VALUES (?)",
		userID,
	)
	return err
}

func (r *UserRepository) UpdateProfile(userID string, bio, website, location, avatarURL, coverURL string) error {
	_, err := r.db.Exec(
		`UPDATE user_profiles SET bio=?, website=?, location=?, avatar_url=?, cover_url=? WHERE user_id=?`,
		bio, website, location, avatarURL, coverURL, userID,
	)
	return err
}

func (r *UserRepository) Follow(followerID, followingID string) error {
	_, err := r.db.Exec(
		"INSERT IGNORE INTO user_follows (follower_id, following_id) VALUES (?, ?)",
		followerID, followingID,
	)
	if err != nil {
		return err
	}

	// Update counts
	_, err = r.db.Exec("UPDATE user_profiles SET followers_count = followers_count + 1 WHERE user_id = ?", followingID)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("UPDATE user_profiles SET following_count = following_count + 1 WHERE user_id = ?", followerID)
	return err
}

func (r *UserRepository) Unfollow(followerID, followingID string) error {
	result, err := r.db.Exec(
		"DELETE FROM user_follows WHERE follower_id = ? AND following_id = ?",
		followerID, followingID,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows > 0 {
		_, _ = r.db.Exec("UPDATE user_profiles SET followers_count = GREATEST(followers_count - 1, 0) WHERE user_id = ?", followingID)
		_, _ = r.db.Exec("UPDATE user_profiles SET following_count = GREATEST(following_count - 1, 0) WHERE user_id = ?", followerID)
	}
	return nil
}

func (r *UserRepository) IsFollowing(followerID, followingID string) (bool, error) {
	var count int
	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM user_follows WHERE follower_id = ? AND following_id = ?",
		followerID, followingID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) GetFollowers(userID string, limit, offset int) ([]string, error) {
	rows, err := r.db.Query(
		"SELECT follower_id FROM user_follows WHERE following_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?",
		userID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		followers = append(followers, id)
	}
	return followers, nil
}

func (r *UserRepository) GetFollowing(userID string, limit, offset int) ([]string, error) {
	rows, err := r.db.Query(
		"SELECT following_id FROM user_follows WHERE follower_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?",
		userID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		following = append(following, id)
	}
	return following, nil
}
