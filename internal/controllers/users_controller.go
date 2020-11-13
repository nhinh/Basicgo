package controllers

import (
	"Basicgo/internal/models"
	"net/http"
	"Basicgo/pkg/utils"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID: 4,
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

func (server *Server) CreateUserSelect(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Username: "selectName",
		Email: "nhinhdtaaaaa@tmh.vn",
		Password : "123456890",
	}
	userSelect, err := user.SelectSaveUser(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
	}
	utils.JSON(w, http.StatusCreated, userSelect)
}
