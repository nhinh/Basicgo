package controllers

import (
	"Basicgo/internal/models"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"github.com/rs/cors"
	_ "github.com/lib/pq"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	// Switch connect DB
	if Dbdriver == "postgres" {
		//url connect db
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			DbHost,
			DbPort,
			DbUser,
			DbName,
			DbPassword)
		//assignable to server.DB
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		//check err
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.Post{}) //database migration

	// assignable to Router
	server.Router = mux.NewRouter()

	// List Router
	server.initializeRoutes()
}

// Run Serve and integrate Route
func (server *Server) Run(addr string){
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you server is available and allowed for this base url
		AllowedMethods: []string{"GET", "PUT", "POST", "OPTION", "DELETE"},

		AllowedHeaders: []string{
			"*",//or you can your header key values which you are using in your application
		},
	})
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, corsOpts.Handler(server.Router)))
}