package handler

import (
	"clevergo.tech/jsend"
	"encoding/json"
	"hristoforovada-project/backend/internal/entity"
	"log"
	"net/http"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	userReg := entity.UserRegister{}
	err := json.NewDecoder(r.Body).Decode(&userReg)

	if err != nil {
		log.Printf("Error decoding when register user: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
	}
	log.Printf(userReg.Password)

	err = h.service.User.Register(&userReg)
	if err != nil {
		log.Printf("Error when register user: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	userReg := entity.UserRegister{}
	err := json.NewDecoder(r.Body).Decode(&userReg)
	if err != nil {
		log.Printf("Error decoding when register user: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
	}

	userId, err := h.service.User.Login(&userReg)
	if err != nil {
		log.Printf("Error when register user: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenString, _ := h.auth.MakeAuth(userId)
	w.Header().Add("Authorization", "Bearer "+tokenString)
	json_msg := map[string]string{
		"message": "successful login",
	}
	jsend.Success(w, json_msg, http.StatusOK)
}
