package comment

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	comService Service
}

func (s *Server) CreateCom(w http.ResponseWriter, r *http.Request) {
	var newCom Comment
	err := json.NewDecoder(r.Body).Decode(&newCom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	postId := chi.URLParam(r, "id_post")
	id, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment, err := s.comService.CreateCom(newCom, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *Server) GetCom(w http.ResponseWriter, _ *http.Request) {
	comment, err := s.comService.GetCom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetComByID(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id_post")
	_, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commentId := chi.URLParam(r, "id")
	idCom, err := strconv.Atoi(commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment, err := s.comService.GetComByID(idCom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetComByPostID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id_post")
	id, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment, err := s.comService.GetComByPostID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetComByUserID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id_user")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment, err := s.comService.GetComByUserID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetComByDate(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id_post")
	id, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commentDate := chi.URLParam(r, "date")
	date, err := time.Parse("2006-01-02", commentDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment, err := s.comService.GetComByDate(date, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) EditCom(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id_post")
	idPost, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commentId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var editedCom Comment
	err = json.NewDecoder(r.Body).Decode(&editedCom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if editedCom.IDPost != idPost {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment, err := s.comService.EditCom(editedCom, id, idPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeleteCom(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id_post")
	_, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commentId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.comService.DeleteCom(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewServer(comService Service) *Server { return &Server{comService} }
