package post

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	postService Service
}

func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post, err := s.postService.CreatePost(newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetPosts(w http.ResponseWriter, _ *http.Request) {
	posts, err := s.postService.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := s.postService.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetPostByUserID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id_user")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := s.postService.GetPostByUserID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetPostByDate(w http.ResponseWriter, r *http.Request) {
	postDate := chi.URLParam(r, "date")
	date, err := time.Parse("2006-01-02", postDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := s.postService.GetPostByDate(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetPostByTitle(w http.ResponseWriter, r *http.Request) {
	postTitle := chi.URLParam(r, "title")
	post, err := s.postService.GetPostByTitle(postTitle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) EditPost(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var editedPost Post
	err = json.NewDecoder(r.Body).Decode(&editedPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post, err := s.postService.EditPost(editedPost, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.postService.DeletePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewServer(postService Service) *Server {
	return &Server{postService}
}
