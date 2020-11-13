package controllers

import (
	"Basicgo/internal/models"
	"net/http"
	"Basicgo/pkg/utils"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID: 3,
		Username: "Nhinh create",
		Email: "nhinhdt@tmh.vn",
		Password : "123456",
	}
	userCreated, error := user.SaveUser(server.DB)
	if error != nil {
		utils.ERROR(w, http.StatusBadRequest, error)
	}
	utils.JSON(w, http.StatusCreated, userCreated)
}
