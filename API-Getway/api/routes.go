package api

func (server *Server) InitRoutes() {
	server.router.POST("/login", server.login)
	server.router.POST("/register", server.register)
	server.router.POST("/renew-token", server.renewToken)

	server.router.Use(server.validateTokenMiddleware)
	{
		server.router.GET("/users/:id", server.getUserById)
		server.router.GET("/users", server.listUsers)

		server.router.GET("/books/:id", server.getBook)
		server.router.GET("/books", server.listBooks)
		server.router.POST("/books", server.addBook)
		server.router.PUT("/books", server.updateBook)

		server.router.GET("/authors/:id", server.getAuthor)
		server.router.POST("/authors", server.addAuthor)

		server.router.GET("/publishers/:id", server.getPublisher)
		server.router.POST("/publishers", server.addPublisher)
	}
}
