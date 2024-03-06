package user

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type Server struct {
	userService Service
}

func (s *Server) GetUsers(w http.ResponseWriter, _ *http.Request) {
	user, err := s.userService.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := s.userService.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := s.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	userEmail := chi.URLParam(r, "email")
	user, err := s.userService.GetUserByEmail(userEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var userUp User
	err = json.NewDecoder(r.Body).Decode(&userUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := s.userService.UpdateUser(userUp, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.userService.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) FollowUser(w http.ResponseWriter, r *http.Request) {
	idFollower := chi.URLParam(r, "id")
	follower, err := strconv.Atoi(idFollower)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idFollowing := chi.URLParam(r, "following_id")
	following, err := strconv.Atoi(idFollowing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.userService.FollowUser(follower, following)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeleteConnection(w http.ResponseWriter, r *http.Request) {
	idFollower := chi.URLParam(r, "id")
	follower, err := strconv.Atoi(idFollower)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idFollowing := chi.URLParam(r, "following_id")
	following, err := strconv.Atoi(idFollowing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.userService.DeleteConnection(follower, following)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetFollowingByUserID(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := s.userService.GetFollowingByUserID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetUserFollowers(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := s.userService.GetUserFollowers(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetFollow(w http.ResponseWriter, r *http.Request) {
	follower := r.URL.Query().Get("follower")
	isFollowing, err := strconv.ParseBool(follower)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if isFollowing {
		s.GetFollowingByUserID(w, r)
	} else {
		s.GetUserFollowers(w, r)
	}
}

func NewServer(userService Service) *Server {
	return &Server{userService}
}
