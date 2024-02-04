package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"hristoforovada-project/backend/internal/entity"
	"hristoforovada-project/backend/internal/service"
	"hristoforovada-project/backend/pkg/auth"
	"log"
	"net/http"
)

type Handler struct {
	service *service.Service
	auth    entity.AuthManager
}

func NewHandler(s *service.Service, auth *auth.AuthManager) *Handler {
	return &Handler{service: s, auth: auth}
}

func (h *Handler) NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/posts", h.GetPosts).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/posts/{post_id}", h.GetPostById).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/posts/{post_id}/delete", h.DeletePost).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/create", h.CreatePost).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/posts/{post_id}/update", h.UpdatePost).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/registration", h.RegisterUser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/login", h.LoginUser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/publication/{post_id}/comments", h.GetComments).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/publication/{post_id}/create_comment", h.CreateComment).Methods(http.MethodPost, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(LoggingMiddleware(r))
	return r
}

func LoggingMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Printf("Origin: %s | Forwarded: %s | Method: %s | RequestURI: %s", req.Header.Get("Origin"), req.Header.Get("Forwarded"), req.Method, req.RequestURI)

			next.ServeHTTP(w, req)
		})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!\n")
}
