package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"megaverse.dev/services/user-service/internal/middleware"
	"megaverse.dev/services/user-service/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	profile, err := h.userService.GetProfile(userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get profile")
		return
	}

	writeJSON(w, http.StatusOK, profile)
}

func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	if userID == "" {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req service.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.userService.UpdateProfile(userID, &req); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update profile")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "updated"})
}

func (h *UserHandler) Follow(w http.ResponseWriter, r *http.Request) {
	followerID := middleware.GetUserID(r)
	if followerID == "" {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	followingID := vars["id"]

	if followerID == followingID {
		writeError(w, http.StatusBadRequest, "Cannot follow yourself")
		return
	}

	if err := h.userService.Follow(followerID, followingID); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to follow user")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "followed"})
}

func (h *UserHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	followerID := middleware.GetUserID(r)
	if followerID == "" {
		writeError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	followingID := vars["id"]

	if err := h.userService.Unfollow(followerID, followingID); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to unfollow user")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "unfollowed"})
}

func (h *UserHandler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 20
	}
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	followers, err := h.userService.GetFollowers(userID, limit, offset)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get followers")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"followers": followers,
		"limit":     limit,
		"offset":    offset,
	})
}

func (h *UserHandler) GetFollowing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 20
	}
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	following, err := h.userService.GetFollowing(userID, limit, offset)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to get following")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"following": following,
		"limit":     limit,
		"offset":    offset,
	})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
