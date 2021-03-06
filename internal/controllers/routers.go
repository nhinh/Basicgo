package controllers

func (s *Server) initializeRoutes() {
	//Users routes
	s.Router.HandleFunc("/users", (s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/usersSelect", (s.CreateUserSelect)).Methods("POST")
	s.Router.HandleFunc("/users", (s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", (s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", (s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", (s.DeleteUser)).Methods("DELETE")

	////Posts routes
	//s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	//s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	//s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	//s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	//s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
