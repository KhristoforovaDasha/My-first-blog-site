package handler

import (
	"clevergo.tech/jsend"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"hristoforovada-project/backend/internal/entity"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	posts, err := h.service.Post.GetAll()
	if err != nil {
		log.Printf("Error when getting all posts: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}

	err = jsend.Success(w, posts, http.StatusOK)
	if err != nil {
		log.Printf("Error when encoding posts: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	vars := mux.Vars(r)
	postId, err := strconv.ParseUint(vars["post_id"], 10, 32)
	if err != nil {
		log.Printf("Error when parsing post id: %v\n", err.Error())
		return
	}

	post, err := h.service.Post.Get(uint(postId))
	if err != nil {
		log.Printf("Error when getting post: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		log.Printf("Error when encoding posts: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	post := entity.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	fmt.Println("create a post")
	if err != nil {
		log.Printf("Error decoding when creating post: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
	}

	err = h.service.Post.Create(&post)
	if err != nil {
		log.Printf("Error when creating post: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	vars := mux.Vars(r)
	postId, err := strconv.ParseUint(vars["post_id"], 10, 32)
	if err != nil {
		log.Printf("Error when parsing post id: %v\n", err.Error())
		return
	}

	err = h.service.Post.Delete(uint(postId))
	if err != nil {
		log.Printf("Error when getting post: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	post := entity.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Printf("Error decoding when updating a post: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
	}

	err = h.service.Post.Update(&post)
	if err != nil {
		log.Printf("Error when updating a post: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	jsend.Success(w, post, http.StatusOK)
}
