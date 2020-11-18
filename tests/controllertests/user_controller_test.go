package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCreateUser(t *testing.T) {

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
			errorMessage: "Email Already Taken",
		},
		{
			inputJSON:    `{"username":"Pet", "email": "grand@gmail.com", "password": "password"}`,
			statusCode:   400,
			errorMessage: "username Already Taken",
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
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

// func TestGetUsers(t *testing.T) {

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = seedUsers()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req, err := http.NewRequest("GET", "/users", nil)
// 	if err != nil {
// 		t.Errorf("this is the error: %v\n", err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(server.GetUsers)
// 	handler.ServeHTTP(rr, req)

// 	var users []models.User
// 	err = json.Unmarshal([]byte(rr.Body.String()), &users)
// 	if err != nil {
// 		log.Fatalf("Cannot convert to json: %v\n", err)
// 	}
// 	assert.Equal(t, rr.Code, http.StatusOK)
// 	assert.Equal(t, len(users), 2)
// }

// func TestGetUserByID(t *testing.T) {

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	user, err := seedOneUser()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	userSample := []struct {
// 		id           string
// 		statusCode   int
// 		nickname     string
// 		email        string
// 		errorMessage string
// 	}{
// 		{
// 			id:         strconv.Itoa(int(user.ID)),
// 			statusCode: 200,
// 			nickname:   user.Username,
// 			email:      user.Email,
// 		},
// 		{
// 			id:         "unknwon",
// 			statusCode: 400,
// 		},
// 	}
// 	for _, v := range userSample {

// 		req, err := http.NewRequest("GET", "/users", nil)
// 		if err != nil {
// 			t.Errorf("This is the error: %v\n", err)
// 		}
// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.GetUser)
// 		handler.ServeHTTP(rr, req)

// 		responseMap := make(map[string]interface{})
// 		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 		if err != nil {
// 			log.Fatalf("Cannot convert to json: %v", err)
// 		}

// 		assert.Equal(t, rr.Code, v.statusCode)

// 		if v.statusCode == 200 {
// 			assert.Equal(t, user.Username, responseMap["username"])
// 			assert.Equal(t, user.Email, responseMap["email"])
// 		}
// 	}
// }

// func TestUpdateUser(t *testing.T) {

// 	var AuthEmail, AuthPassword string
// 	var AuthID uint32

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	users, err := seedUsers() //we need atleast two users to properly check the update
// 	if err != nil {
// 		log.Fatalf("Error seeding user: %v\n", err)
// 	}
// 	// Get only the first user
// 	for _, user := range users {
// 		if user.ID == 2 {
// 			continue
// 		}
// 		AuthID = user.ID
// 		AuthEmail = user.Email
// 		AuthPassword = "password" //Note the password in the database is already hashed, we want unhashed
// 	}

// 	samples := []struct {
// 		id             string
// 		updateJSON     string
// 		statusCode     int
// 		updateNickname string
// 		updateEmail    string
// 		errorMessage   string
// 	}{
// 		{
// 			// Convert int32 to int first before converting to string
// 			id:             strconv.Itoa(int(AuthID)),
// 			updateJSON:     `{"niusernameckname":"Grand", "email": "grand@gmail.com", "password": "password"}`,
// 			statusCode:     200,
// 			updateNickname: "Grand",
// 			updateEmail:    "grand@gmail.com",
// 			errorMessage:   "",
// 		},
// 		{
// 			// When password field is empty
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"username":"Woman", "email": "woman@gmail.com", "password": ""}`,
// 			statusCode:   422,
// 			errorMessage: "Required Password",
// 		},
// 		{
// 			// When incorrect token was passed
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"username":"Woman", "email": "woman@gmail.com", "password": "password"}`,
// 			statusCode:   401,
// 			errorMessage: "Unauthorized",
// 		},
// 		{
// 			// Remember "kenny@gmail.com" belongs to user 2
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"username":"Frank", "email": "kenny@gmail.com", "password": "password"}`,
// 			statusCode:   500,
// 			errorMessage: "Email Already Taken",
// 		},
// 		{
// 			// Remember "Kenny Morris" belongs to user 2
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"username":"Kenny Morris", "email": "grand@gmail.com", "password": "password"}`,
// 			statusCode:   500,
// 			errorMessage: "Nickname Already Taken",
// 		},
// 		{
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"username":"Kan", "email": "kangmail.com", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Invalid Email",
// 		},
// 		{
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"username": "", "email": "kan@gmail.com", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Required username",
// 		},
// 		{
// 			id:           strconv.Itoa(int(AuthID)),
// 			updateJSON:   `{"username": "Kan", "email": "", "password": "password"}`,
// 			statusCode:   422,
// 			errorMessage: "Required Email",
// 		},
// 		{
// 			id:         "unknwon",
// 			statusCode: 400,
// 		},
// 		{
// 			// When user 2 is using user 1 token
// 			id:           strconv.Itoa(int(2)),
// 			updateJSON:   `{"username": "Mike", "email": "mike@gmail.com", "password": "password"}`,
// 			statusCode:   401,
// 			errorMessage: "Unauthorized",
// 		},
// 	}

// 	for _, v := range samples {

// 		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.updateJSON))
// 		if err != nil {
// 			t.Errorf("This is the error: %v\n", err)
// 		}
// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})

// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.UpdateUser)

// 		// req.Header.Set("Authorization", v.tokenGiven)

// 		handler.ServeHTTP(rr, req)

// 		responseMap := make(map[string]interface{})
// 		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 		if err != nil {
// 			t.Errorf("Cannot convert to json: %v", err)
// 		}
// 		assert.Equal(t, rr.Code, v.statusCode)
// 		if v.statusCode == 200 {
// 			assert.Equal(t, responseMap["username"], v.updateNickname)
// 			assert.Equal(t, responseMap["email"], v.updateEmail)
// 		}
// 		if v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
// 			assert.Equal(t, responseMap["error"], v.errorMessage)
// 		}
// 	}
// }

// func TestDeleteUser(t *testing.T) {

// 	var AuthEmail, AuthPassword string
// 	var AuthID uint32

// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	users, err := seedUsers() //we need atleast two users to properly check the update
// 	if err != nil {
// 		log.Fatalf("Error seeding user: %v\n", err)
// 	}
// 	// Get only the first and log him in
// 	for _, user := range users {
// 		if user.ID == 2 {
// 			continue
// 		}
// 		AuthID = user.ID
// 		AuthEmail = user.Email
// 		AuthPassword = "password" ////Note the password in the database is already hashed, we want unhashed
// 	}

// 	userSample := []struct {
// 		id           string
// 		statusCode   int
// 		errorMessage string
// 	}{
// 		{
// 			// Convert int32 to int first before converting to string
// 			id:           strconv.Itoa(int(AuthID)),
// 			statusCode:   204,
// 			errorMessage: "",
// 		},
// 		{
// 			// When incorrect token is given
// 			id:           strconv.Itoa(int(AuthID)),
// 			statusCode:   401,
// 			errorMessage: "Unauthorized",
// 		},
// 		{
// 			id:         "unknwon",
// 			statusCode: 400,
// 		},
// 		{
// 			// User 2 trying to use User 1 token
// 			id:           strconv.Itoa(int(2)),
// 			statusCode:   401,
// 			errorMessage: "Unauthorized",
// 		},
// 	}
// 	for _, v := range userSample {

// 		req, err := http.NewRequest("GET", "/users", nil)
// 		if err != nil {
// 			t.Errorf("This is the error: %v\n", err)
// 		}
// 		req = mux.SetURLVars(req, map[string]string{"id": v.id})
// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(server.DeleteUser)

// 		// req.Header.Set("Authorization", v.tokenGiven)

// 		handler.ServeHTTP(rr, req)
// 		assert.Equal(t, rr.Code, v.statusCode)

// 		if v.statusCode == 401 && v.errorMessage != "" {
// 			responseMap := make(map[string]interface{})
// 			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
// 			if err != nil {
// 				t.Errorf("Cannot convert to json: %v", err)
// 			}
// 			assert.Equal(t, responseMap["error"], v.errorMessage)
// 		}
// 	}
// }
