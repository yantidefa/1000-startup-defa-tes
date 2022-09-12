package controllers

import "seribu/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddleware(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/users/login", middlewares.SetMiddleware(s.Login)).Methods("POST")
	s.Router.HandleFunc("/admin/login", middlewares.SetMiddleware(s.Logins)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddleware(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	s.Router.HandleFunc("/startups", middlewares.SetMiddleware(s.CreateStartup)).Methods("POST")
	s.Router.HandleFunc("/startups", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.GetStartups))).Methods("GET")
	s.Router.HandleFunc("/startups/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.GetStartup))).Methods("GET")
	s.Router.HandleFunc("/startups/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateStartup))).Methods("PUT")
	s.Router.HandleFunc("/startups/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteStartup)).Methods("DELETE")

	s.Router.HandleFunc("/domiciles", middlewares.SetMiddleware(s.CreateDomicile)).Methods("POST")
	s.Router.HandleFunc("/domiciles", middlewares.SetMiddleware(s.GetDomiciles)).Methods("GET")
	s.Router.HandleFunc("/domiciles/{id}", middlewares.SetMiddleware(s.GetDomicile)).Methods("GET")
	s.Router.HandleFunc("/domiciles/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateDomicile))).Methods("PUT")
	s.Router.HandleFunc("/domiciles/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteDomicile)).Methods("DELETE")

	s.Router.HandleFunc("/blogs_categorys", middlewares.SetMiddleware(s.CreateBlogsCategory)).Methods("POST")
	s.Router.HandleFunc("/blogs_categorys", middlewares.SetMiddleware(s.GetBlogsCategorys)).Methods("GET")
	s.Router.HandleFunc("/blogs_categorys/{id}", middlewares.SetMiddleware(s.GetBlogsCategory)).Methods("GET")
	s.Router.HandleFunc("/blogs_categorys/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateBlogsCategory))).Methods("PUT")
	s.Router.HandleFunc("/blogs_categorys/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteBlogsCategory)).Methods("DELETE")

	s.Router.HandleFunc("/books_categorys", middlewares.SetMiddleware(s.CreateBooksCategory)).Methods("POST")
	s.Router.HandleFunc("/books_categorys", middlewares.SetMiddleware(s.GetBooksCategorys)).Methods("GET")
	s.Router.HandleFunc("/books_categorys/{id}", middlewares.SetMiddleware(s.GetBooksCategory)).Methods("GET")
	s.Router.HandleFunc("/books_categorys/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateBooksCategory))).Methods("PUT")
	s.Router.HandleFunc("/books_categorys/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteBooksCategory)).Methods("DELETE")

	s.Router.HandleFunc("/admins", middlewares.SetMiddleware(s.CreateAdmin)).Methods("POST")
	s.Router.HandleFunc("/admins", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.GetAdmins))).Methods("GET")
	s.Router.HandleFunc("/admins/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.GetAdmin))).Methods("GET")
	s.Router.HandleFunc("/admins/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateAdmin))).Methods("PUT")
	s.Router.HandleFunc("/admins/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteAdmin)).Methods("DELETE")

	s.Router.HandleFunc("/blogs", middlewares.SetMiddleware(s.CreateBlog)).Methods("POST")
	s.Router.HandleFunc("/blogs", middlewares.SetMiddleware(s.GetBlogs)).Methods("GET")
	s.Router.HandleFunc("/blogs/{id}", middlewares.SetMiddleware(s.GetBlog)).Methods("GET")
	s.Router.HandleFunc("/blogs/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateBlog))).Methods("PUT")
	s.Router.HandleFunc("/blogs/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteBlog)).Methods("DELETE")

	s.Router.HandleFunc("/roles", middlewares.SetMiddleware(s.CreateRolesPoint)).Methods("POST")
	s.Router.HandleFunc("/roles", middlewares.SetMiddleware(s.GetRolesPoints)).Methods("GET")
	s.Router.HandleFunc("/roles/{id}", middlewares.SetMiddleware(s.GetRolesPoint)).Methods("GET")
	s.Router.HandleFunc("/roles/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateRolesPoint))).Methods("PUT")
	s.Router.HandleFunc("/roles/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteRolesPoint)).Methods("DELETE")

	s.Router.HandleFunc("/answers", middlewares.SetMiddleware(s.CreateAnswer)).Methods("POST")
	s.Router.HandleFunc("/answers", middlewares.SetMiddleware(s.GetAnswers)).Methods("GET")
	s.Router.HandleFunc("/answers/{id}", middlewares.SetMiddleware(s.GetAnswer)).Methods("GET")
	s.Router.HandleFunc("/answers/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateAnswer))).Methods("PUT")
	s.Router.HandleFunc("/answers/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteAnswer)).Methods("DELETE")

	s.Router.HandleFunc("/questions", middlewares.SetMiddleware(s.CreateQuestion)).Methods("POST")
	s.Router.HandleFunc("/questions", middlewares.SetMiddleware(s.GetQuestions)).Methods("GET")
	s.Router.HandleFunc("/questions/{id}", middlewares.SetMiddleware(s.GetQuestion)).Methods("GET")
	s.Router.HandleFunc("/questions/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateQuestion))).Methods("PUT")
	s.Router.HandleFunc("/questions/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteQuestion)).Methods("DELETE")

	s.Router.HandleFunc("/books", middlewares.SetMiddleware(s.CreateBook)).Methods("POST")
	s.Router.HandleFunc("/books", middlewares.SetMiddleware(s.GetBooks)).Methods("GET")
	s.Router.HandleFunc("/books/{id}", middlewares.SetMiddleware(s.GetBook)).Methods("GET")
	s.Router.HandleFunc("/books/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateBook))).Methods("PUT")
	s.Router.HandleFunc("/books/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteBook)).Methods("DELETE")

	s.Router.HandleFunc("/activitys", middlewares.SetMiddleware(s.CreateActivity)).Methods("POST")
	s.Router.HandleFunc("/activitys", middlewares.SetMiddleware(s.GetActivitys)).Methods("GET")
	s.Router.HandleFunc("/activitys/{id}", middlewares.SetMiddleware(s.GetActivity)).Methods("GET")
	s.Router.HandleFunc("/activitys/{id}", middlewares.SetMiddleware(middlewares.SetMiddlewareAuthentication(s.UpdateActivity))).Methods("PUT")
	s.Router.HandleFunc("/activitys/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteActivity)).Methods("DELETE")
}
