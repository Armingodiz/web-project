package app

import (
	"fmt"

	"web-project/store"
	"web-project/utils"

	"web-project/config"

	urlcontroller "web-project/controllers/urlController"
	usercontroller "web-project/controllers/userController"
	"web-project/middlewares"

	"web-project/db"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.ForceConsoleColor()
}

type App struct {
	route *gin.Engine
}

func NewApp() *App {
	// Initialize databae
	db, err := initializeDB()
	utils.FailOnError(err, "Database initialization failed, exiting the app with error!")
	r := routing(db)
	return &App{
		route: r,
	}
}

func (a *App) Start(restPort string) error {
	return a.route.Run(restPort)
}

func routing(db *db.DB) *gin.Engine {
	r := gin.Default()
	postgresStore := store.NewStore(db)
	UserController := usercontroller.UserController{Store: postgresStore}
	UrlController := urlcontroller.UrlController{Store: postgresStore}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})
	r.POST("/users/signup", UserController.Signup())
	r.POST("/users/login", UserController.Login())

	//Protected routes
	r.Use(middlewares.JwtAuthorizationMiddleware())
	r.POST("/urls", UrlController.CreateUrl())
	r.GET("/urls", UrlController.GetUrls())
	r.GET("/urls/:id", UrlController.GetUrl())
	return r
}

func initializeDB() (*db.DB, error) {
	host := config.Configs.Database.Host
	port := config.Configs.Database.Port
	user := config.Configs.Database.User
	password := config.Configs.Database.Password
	dbName := config.Configs.Database.DbName
	extras := config.Configs.Database.Extras
	driver := config.Configs.Database.Driver

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s %s",
		host, port, user, password, dbName, extras)
	db, err := db.Connect(driver, connStr)
	return db, err
}
