package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"hristoforovada-project/backend/internal/entity"
	"log"
	"net/http"
	"strconv"
)

type CommentWithAuthor struct {
	userName string
	comment  entity.Comment
}

func (h *Handler) GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, err := strconv.ParseUint(vars["post_id"], 10, 32)
	if err != nil {
		log.Printf("Error when parsing post id: %v\n", err.Error())
		return
	}

	var comments_with_author []CommentWithAuthor
	comments, err := h.service.Comment.GetCommentsByPostId(uint(postId))
	for _, comment := range *comments {
		user, err := h.service.User.Get(comment.UserId)
		if err != nil {
			log.Printf("Error when getting a comment author: %v\n", err.Error())
		} else {
			comments_with_author = append(comments_with_author, CommentWithAuthor{user.Login, comment})
		}
		//fmt.Printf("%+v\n", CommentWithAuthor{user.Login, comment})
	}
	if err != nil {
		log.Printf("Error when getting comments: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	err = json.NewEncoder(w).Encode(comments_with_author)
	if err != nil {
		log.Printf("Error when encoding post's comments: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, err := strconv.ParseUint(vars["post_id"], 10, 32)
	comment := entity.Comment{}
	err = json.NewDecoder(r.Body).Decode(&comment)
	comment.PostId = uint(postId)
	if err != nil {
		log.Printf("Error decoding when adding a comment: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
	}
	fmt.Printf("%+v\n", comment)

	err = h.service.Comment.Create(&comment)
	if err != nil {
		log.Printf("Error when adding a comment: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
