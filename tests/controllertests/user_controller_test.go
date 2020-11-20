package controllertests

import (
	"Basicgo/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

// Link tham khoan: https://blog.vietnamlab.vn/cach-viet-unit-test-cho-rest-api-trong-golang/

func TestCreateUser(t *testing.T) {
	// Xoa table
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	// Test case
	samples := []struct {
		inputJSON    string
		statusCode   int
		username     string
		email        string
		errorMessage string
	}{
		{
			inputJSON:    `{"username":"Pet", "email": "pet@gmail.com", "password": "password"}`,
			statusCode:   201,
			username:     "Pet",
			email:        "pet@gmail.com",
			errorMessage: "",
		},
		{
			inputJSON:    `{"username":"Frank", "email": "pet@gmail.com", "password": "password"}`,
			statusCode:   400,
			errorMessage: "pq: duplicate key value violates unique constraint \"users_email_key\"",
		},
		{
			inputJSON:    `{"username":"Pet", "email": "grand@gmail.com", "password": "password"}`,
			statusCode:   400,
			errorMessage: "pq: duplicate key value violates unique constraint \"users_username_key\"",
		},
		{
			inputJSON:    `{"username":"Kan", "email": "kangmail.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Invalid Email",
		},
		{
			inputJSON:    `{"username": "", "email": "kan@gmail.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required username",
		},
		{
			inputJSON:    `{"username": "Kan", "email": "", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required email",
		},
		{
			inputJSON:    `{"username": "Kan", "email": "kan@gmail.com", "password": ""}`,
			statusCode:   422,
			errorMessage: "Required password",
		},
	}

	for _, v := range samples {
		// tạo một request
		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}

		// Gọi phương thức cần test.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}

		// So sánh status thực tế nhận được và status mong muốn.
		assert.Equal(t, rr.Code, v.statusCode)

		// So sánh response thực tế và response mong muốn
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["username"], v.username)
			assert.Equal(t, responseMap["email"], v.email)
		}

		if v.statusCode == 422 || v.statusCode == 500 || v.statusCode == 400 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestGetUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	// Gọi phương thức cần test.
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUsers)
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal([]byte(rr.Body.String()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	// So sánh status thực tế nhận được và status mong muốn.
	assert.Equal(t, rr.Code, http.StatusOK)

	// So sánh response thực tế và response mong muốn
	assert.Equal(t, len(users), 2)
}

func TestGetUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}
	userSample := []struct {
		id           string
		statusCode   int
		username     string
		email        string
		errorMessage string
	}{
		{
			id:         strconv.Itoa(int(user.ID)),
			statusCode: 200,
			username:   user.Username,
			email:      user.Email,
		},
		{
			id:         "unknwon",
			statusCode: 400,
		},
	}
	for _, v := range userSample {

		req, err := http.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}

		// able to read variables from a url using mux (/users/{id})
		req = mux.SetURLVars(req, map[string]string{"id": v.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, user.Username, responseMap["username"])
			assert.Equal(t, user.Email, responseMap["email"])
		}
	}
}

func TestUpdateUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	users, err := seedUsers() //we need atleast two users to properly check the update
	if err != nil {
		log.Fatalf("Error seeding user: %v\n", err)
	}
	samples := []struct {
		id             string
		updateJSON     string
		statusCode     int
		updateNickname string
		updateEmail    string
		errorMessage   string
	}{
		{
			// Convert int32 to int first before converting to string
			id:             strconv.Itoa(int(users[0].ID)),
			updateJSON:     `{"username":"Grand", "email": "grand@gmail.com", "password": "password"}`,
			statusCode:     200,
			updateNickname: "Grand",
			updateEmail:    "grand@gmail.com",
			errorMessage:   "",
		},
		{
			// When password field is empty
			id:           strconv.Itoa(int(users[0].ID)),
			updateJSON:   `{"username":"Woman", "email": "woman@gmail.com", "password": ""}`,
			statusCode:   422,
			errorMessage: "Required password",
		},
		{
			// Remember "kenny@gmail.com" belongs to user 2
			id:           strconv.Itoa(int(1)),
			updateJSON:   `{"username":"Frank", "email": "kenny@gmail.com", "password": "password"}`,
			statusCode:   500,
			errorMessage: "pq: duplicate key value violates unique constraint \"users_email_key\"",
		},
		{
			// Remember "Kenny Morris" belongs to user 2
			id:           strconv.Itoa(int(1)),
			updateJSON:   `{"username":"Kenny Morris", "email": "grand@gmail.com", "password": "password"}`,
			statusCode:   500,
			errorMessage: "pq: duplicate key value violates unique constraint \"users_username_key\"",
		},
		{
			id:           strconv.Itoa(int(1)),
			updateJSON:   `{"username":"Kan", "email": "kangmail.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Invalid Email",
		},
		{
			id:           strconv.Itoa(int(1)),
			updateJSON:   `{"username": "", "email": "kan@gmail.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required username",
		},
		{
			id:           strconv.Itoa(int(1)),
			updateJSON:   `{"username": "Kan", "email": "", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required email",
		},
		{
			id:         "unknwon",
			statusCode: 400,
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.updateJSON))
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.UpdateUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			assert.Equal(t, responseMap["username"], v.updateNickname)
			assert.Equal(t, responseMap["email"], v.updateEmail)
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	users, err := seedUsers() //we need atleast two users to properly check the update
	if err != nil {
		log.Fatalf("Error seeding user: %v\n", err)
	}
	// Get only the first and log him in

	userSample := []struct {
		id           string
		statusCode   int
		errorMessage string
	}{
		{
			// Convert int32 to int first before converting to string
			id:           strconv.Itoa(int(users[0].ID)),
			statusCode:   204,
			errorMessage: "",
		},
		{
			id:         "unknwon",
			statusCode: 400,
		},
	}
	for _, v := range userSample {

		req, err := http.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Errorf("This is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.DeleteUser)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 401 && v.errorMessage != "" {
			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}
