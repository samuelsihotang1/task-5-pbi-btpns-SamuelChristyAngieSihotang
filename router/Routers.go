package router

import (
	"belajar-go/controllers"
	"belajar-go/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	// Register
	r.POST("/users/register", controllers.SignUp)
	r.GET("/register", func(c *gin.Context) {
		c.File("./views/Register.html")
	})

	// Login
	r.POST("/login", controllers.Login)
	r.GET("/users/login", func(c *gin.Context) {
		c.File("./views/Login.html")
	})

	// Middleware for all routes
	r.NoRoute(middleware.RequireAuth)

	m := r.Group("/")
	m.Use(middleware.RequireAuth)
	{
		// Homepage
		m.GET("/", func(c *gin.Context) {
			c.File("./views/Homepage.html")
		})

		/////////////////////////////////////////////////////
		// Get User Info
		m.GET("/users/info", controllers.GetUserInfo)

		// Edit Users
		m.GET("/users", func(c *gin.Context) {
			c.File("./views/Users.html")
		})
		m.PUT("/users/:userId", controllers.EditUsers)

		// Logout
		m.POST("/users/logout", controllers.Logout)

		// Delete User
		m.DELETE("/users/:userId", controllers.DeleteUser)
		/////////////////////////////////////////////////////
	}
}
