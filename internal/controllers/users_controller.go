package controllers

import (
	"Basicgo/internal/models"
	"Basicgo/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/RussellGilmore/potago/api/responses"
	"github.com/RussellGilmore/potago/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID:       4,
		Username: "Nhinh create",
		Email:    "nhinhdt@tmh.vn",
		Password: "123456",
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
		Email:    "nhinhdtaaaaa@tmh.vn",
		Password: "123456890",
	}
	userSelect, err := user.SelectSaveUser(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
	}
	utils.JSON(w, http.StatusCreated, userSelect)
}

// func (server *Server) CreateBatchInsert(w http.ResponseWriter, r *http.Request) {
// 	var users = []models.User{
// 		{Username: "jinzhu1", Email: "Nhinhdt12@tmh-techlab", Password: "12345"},
// 		{Username: "jinzhu2", Email: "Nhinhdt123@tmh-techlab", Password: "123455"},
// 		{Username: "jinzhu3", Email: "Nhinhdt1234@tmh-techlab", Password: "123457"},
// 	}
// 	result, err := users.BatchInsert(server.DB)
// 	if err != nil {
// 		utils.ERROR(w, http.StatusBadRequest, err)
// 	}
// 	utils.JSON(w, http.StatusCreated, result)
// }

// func (server *Server) CreateFormMap(w http.ResponseWriter, r *http.Request) {
// 	var user = []map[string]interface{}{
// 		{
// 			"Username": "map1",
// 			"Email":    "Nhinhdt12map@tmh-techlab",
// 			"Password": "12345map",
// 		},
// 		{
// 			"Username": "map2",
// 			"Email":    "Nhinhdt12map@tmh-techlab",
// 			"Password": "12345map",
// 		},
// 	}
// 	utils.JSON(w, http.StatusCreated, user)
// }

// func (server *Server) GetSingleObject(w http.ResponseWriter, r *http.Request) {
// 	user := models.User{}
// 	result, err := user.SingleObject(server.DB)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, result)
// }

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Loi tai day 3")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		fmt.Println("Loi tai day 4")
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

// DeleteUser...
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	utils.JSON(w, http.StatusNoContent, "Xoa thanh cong")
}
